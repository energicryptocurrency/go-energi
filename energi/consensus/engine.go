// Copyright 2019 The Energi Core Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

package consensus

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	eth_consensus "github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"

	"golang.org/x/crypto/sha3"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

var (
	sealLen   = 65
	uncleHash = types.CalcUncleHash(nil)

	errMissingSig = errors.New("Signature is missing")
	errInvalidSig = errors.New("Invalid signature")
)

type ChainReader = eth_consensus.ChainReader
type SignerFn func(common.Address, []byte) ([]byte, error)

type Energi struct {
	config       *params.EnergiConfig
	db           ethdb.Database
	rewardAbi    abi.ABI
	rewardGov    []common.Address
	systemFaucet common.Address
	xferGas      uint64
	callGas      uint64
	signerFn     SignerFn
}

func New(config *params.EnergiConfig, db ethdb.Database) *Energi {
	reward_abi, err := abi.JSON(strings.NewReader(energi_abi.IBlockRewardABI))
	if err != nil {
		panic(err)
		return nil
	}

	return &Energi{
		config:    config,
		db:        db,
		rewardAbi: reward_abi,
		rewardGov: []common.Address{
			energi_params.Energi_Treasury,
			energi_params.Energi_MasternodeRegistry,
			energi_params.Energi_BackboneReward,
			energi_params.Energi_StakerReward,
		},
		systemFaucet: energi_params.Energi_SystemFaucet,
		xferGas:      2000000,
		callGas:      1000000,
	}
}

func (e *Energi) createEVM(
	msg types.Message,
	chain ChainReader,
	header *types.Header,
	statedb *state.StateDB,
) *vm.EVM {
	vmc := &vm.Config{}

	if bc, ok := chain.(*core.BlockChain); ok {
		vmc = bc.GetVMConfig()
	}

	// Only From() is used by fact
	ctx := core.NewEVMContext(msg, header, chain.(core.ChainContext), &header.Coinbase)
	ctx.GasLimit = e.xferGas
	return vm.NewEVM(ctx, statedb, chain.Config(), *vmc)
}

// Author retrieves the Ethereum address of the account that minted the given
// block, which may be different from the header's coinbase if a consensus
// engine is based on signatures.
func (e *Energi) Author(header *types.Header) (common.Address, error) {
	return common.Address{}, nil
}

// VerifyHeader checks whether a header conforms to the consensus rules of a
// given engine. Verifying the seal may be done optionally here, or explicitly
// via the VerifySeal method.
func (e *Energi) VerifyHeader(chain ChainReader, header *types.Header, seal bool) error {
	var err error

	err = e.VerifySeal(chain, header)
	if err != nil {
		return err
	}

	return nil
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
// concurrently. The method returns a quit channel to abort the operations and
// a results channel to retrieve the async verifications (the order is that of
// the input slice).
func (e *Energi) VerifyHeaders(
	chain ChainReader, headers []*types.Header, seals []bool,
) (
	chan<- struct{}, <-chan error,
) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			// TODO: optimize to use headers as history
			err := e.VerifyHeader(chain, header, seals[i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()

	return abort, results

}

// VerifyUncles verifies that the given block's uncles conform to the consensus
// rules of a given engine.
func (e *Energi) VerifyUncles(chain ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// VerifySeal checks whether the crypto seal on a header is valid according to
// the consensus rules of the given engine.
func (e *Energi) VerifySeal(chain ChainReader, header *types.Header) error {
	// Retrieve the signature from the header extra-data
	if len(header.Signature) != sealLen {
		return errMissingSig
	}

	sealhash := e.SealHash(header)
	log.Trace("PoS verify seal hash", "hash", sealhash)

	pubkey, err := crypto.Ecrecover(sealhash.Bytes(), header.Signature)
	if err != nil {
		return err
	}

	var addr common.Address
	copy(addr[:], crypto.Keccak256(pubkey[1:])[12:])

	if addr != header.Coinbase {
		log.Trace("PoS seal compare", "addr", addr, "coinbase", header.Coinbase)
		return errInvalidSig
	}

	return nil
}

// Prepare initializes the consensus fields of a block header according to the
// rules of a particular engine. The changes are executed inline.
func (e *Energi) Prepare(chain ChainReader, header *types.Header) error {
	// Clear out unused
	header.Nonce = types.BlockNonce{}

	parent := chain.GetHeaderByHash(header.ParentHash)

	if parent == nil {
		return errors.New("Unknown parent")
	}

	time_target := e.calcTimeTarget(chain, parent)

	err := e.enforceTime(header, time_target)

	// Repurpose the MixDigest field
	header.MixDigest = e.calcPoSModifier(chain, header.Time, parent)

	// TODO: trim Extra

	// Diff
	header.Difficulty = e.calcPoSDifficulty(chain, header.Time, parent)

	return err
}

// Finalize runs any post-transaction state modifications (e.g. block rewards)
// and assembles the final block.
// Note: The block header and state database might be updated to reflect any
// consensus rules that happen at finalization (e.g. block rewards).
func (e *Energi) Finalize(
	chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
	uncles []*types.Header, receipts []*types.Receipt,
) (
	*types.Block, error,
) {
	err := e.processBlockRewards(chain, header, state)

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = uncleHash

	// NOTE: the code does not check result!
	return types.NewBlock(header, txs, nil, receipts), err

}

// Seal generates a new sealing request for the given input block and pushes
// the result into the given channel.
//
// Note, the method returns immediately and will send the result async. More
// than one result may also be returned depending on the consensus algorithm.
func (e *Energi) Seal(chain ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) (err error) {
	header := block.Header()

	sealhash := e.SealHash(header)
	log.Trace("PoS seal hash", "hash", sealhash)

	// TODO: implement Coinbase+time mining
	header.Signature, err = e.signerFn(header.Coinbase, sealhash.Bytes())
	if err != nil {
		return err
	}

	select {
	case results <- block.WithSeal(header):
	default:
		log.Warn("PoS seal is not ready by miner", "sealhash", e.SealHash(header))
	}

	return nil
}

// SealHash returns the hash of a block prior to it being sealed.
func (e *Energi) SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		// This part is for "mining"
		//header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		// This part is for "mining"
		//header.Time,
		header.Extra,
		header.MixDigest,
		header.Nonce,
		// This part is to be added afterwards
		//header.Signature,
	})
	hasher.Sum(hash[:0])
	return hash
}

func (e *Energi) SignatureHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra,
		header.MixDigest,
		header.Nonce,
		//header.Signature,
	})
	hasher.Sum(hash[:0])
	return hash
}

func (e *Energi) SetSigner(signerFn SignerFn) {
	if e.signerFn != nil {
		panic("Signer must be set only once!")
	}

	e.signerFn = signerFn
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have.
func (e *Energi) CalcDifficulty(chain ChainReader, time uint64, parent *types.Header) *big.Int {
	return e.calcPoSDifficulty(chain, time, parent)
}

// APIs returns the RPC APIs this consensus engine provides.
func (e *Energi) APIs(chain ChainReader) []rpc.API {
	return make([]rpc.API, 0)
}

// Close terminates any background threads maintained by the consensus engine.
func (e *Energi) Close() error {
	return nil
}

func (e *Energi) Hashrate() float64 {
	return 0
}
