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
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	eth_consensus "github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
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
type AccountsFn func() []common.Address
type SignerFn func(common.Address, []byte) ([]byte, error)
type DiffFn func(ChainReader, uint64, *types.Header, *timeTarget) *big.Int

type Energi struct {
	config       *params.EnergiConfig
	db           ethdb.Database
	rewardAbi    abi.ABI
	rewardGov    []common.Address
	dposAbi      abi.ABI
	systemFaucet common.Address
	xferGas      uint64
	callGas      uint64
	signerFn     SignerFn
	accountsFn   AccountsFn
	diffFn       DiffFn
	testing      bool
}

func New(config *params.EnergiConfig, db ethdb.Database) *Energi {
	reward_abi, err := abi.JSON(strings.NewReader(energi_abi.IBlockRewardABI))
	if err != nil {
		panic(err)
		return nil
	}

	dpos_abi, err := abi.JSON(strings.NewReader(energi_abi.IDelegatedPoSABI))
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
		dposAbi:      dpos_abi,
		systemFaucet: energi_params.Energi_SystemFaucet,
		xferGas:      2000000,
		callGas:      1000000,
		diffFn:       calcPoSDifficultyV1,
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
	is_migration := header.IsGen2Migration()

	// Ensure that the header's extra-data section is of a reasonable size
	if uint64(len(header.Extra)) > params.MaximumExtraDataSize && !is_migration {
		return fmt.Errorf("extra-data too long: %d > %d",
			len(header.Extra), params.MaximumExtraDataSize)
	}

	// A special Migration block #1
	if is_migration && (header.Coinbase != energi_params.Energi_MigrationContract) {
		log.Error("PoS migration mismatch",
			"signer", header.Coinbase,
			"required", energi_params.Energi_MigrationContract)
		return errors.New("Invalid Migration")
	}

	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)

	if parent == nil {
		if header.Number.Cmp(common.Big0) != 0 {
			log.Trace("Not found parent", "number", header.Number,
				"hash", header.Hash, "parent", header.ParentHash)
			return eth_consensus.ErrUnknownAncestor
		}

		return nil
	}

	time_target := e.calcTimeTarget(chain, parent)
	err = e.checkTime(header, time_target)
	if err != nil {
		return err
	}

	modifier := e.calcPoSModifier(chain, header.Time, parent)
	if header.MixDigest != modifier {
		return fmt.Errorf("invalid modifier: have %v, want %v",
			header.MixDigest, modifier)
	}

	difficulty := e.calcPoSDifficulty(chain, header.Time, parent, time_target)

	if header.Difficulty.Cmp(difficulty) != 0 {
		return fmt.Errorf("invalid difficulty: have %v, want %v",
			header.Difficulty, difficulty)
	}

	cap := uint64(0x7fffffffffffffff)
	if header.GasLimit > cap {
		return fmt.Errorf("invalid gasLimit: have %v, max %v",
			header.GasLimit, cap)
	}

	// Verify that the gasUsed is <= gasLimit, except for migration
	if (header.GasUsed > header.GasLimit) && !is_migration {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d",
			header.GasUsed, header.GasLimit)
	}

	// Verify that the gas limit remains within allowed bounds
	diff := int64(parent.GasLimit) - int64(header.GasLimit)
	if diff < 0 {
		diff *= -1
	}
	limit := parent.GasLimit / params.GasLimitBoundDivisor

	if (uint64(diff) >= limit) && !is_migration && !parent.IsGen2Migration() {
		return fmt.Errorf("invalid gas limit: have %d, want %d += %d",
			header.GasLimit, parent.GasLimit, limit)
	}

	if header.GasLimit < params.MinGasLimit {
		return fmt.Errorf("invalid gas limit: have %d, minimum %d",
			header.GasLimit, params.MinGasLimit)
	}

	// Verify that the block number is parent's +1
	if diff := new(big.Int).Sub(header.Number, parent.Number); diff.Cmp(big.NewInt(1)) != 0 {
		return eth_consensus.ErrInvalidNumber
	}

	// Verify the engine specific seal securing the block
	err = e.VerifySeal(chain, header)
	if err != nil {
		return err
	}

	err = e.verifyPoSHash(chain, header)
	if err != nil {
		return err
	}

	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
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
	chan<- struct{}, <-chan error, chan<- bool,
) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))
	ready := make(chan bool, len(headers))

	go func() {
		for i, header := range headers {
			// NOTE: unlike Ethash with DAG, there is little sense of this
			//       batch async routine overhead
			select {
			case <-abort:
				return
			case <-ready:
			}

			err := e.VerifyHeader(chain, header, seals[i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()

	return abort, results, ready
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

	sighash := e.SignatureHash(header)
	log.Trace("PoS verify signature hash", "sighash", sighash)

	pubkey, err := crypto.Ecrecover(sighash.Bytes(), header.Signature)
	if err != nil {
		return err
	}

	var addr common.Address
	copy(addr[:], crypto.Keccak256(pubkey[1:])[12:])

	if addr != header.Coinbase {
		// POS-5: Delegated PoS
		//--
		stdb, err := chain.GetStateDB()
		if err != nil {
			log.Error("PoS seal is called without state database", "err", err)
			return err
		}

		parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
		if parent == nil {
			return eth_consensus.ErrUnknownAncestor
		}

		blockst, err := state.New(parent.Root, stdb)
		if err != nil {
			log.Error("PoS state root failure", "err", err)
			return err
		}

		if blockst.GetCodeSize(header.Coinbase) > 0 {
			signerData, err := e.dposAbi.Pack("signerAddress")
			if err != nil {
				log.Error("Fail to prepare signerAddress() call", "err", err)
				return err
			}

			blockst.SetBalance(e.systemFaucet, new(big.Int).SetUint64(e.callGas))
			msg := types.NewMessage(
				e.systemFaucet,
				&header.Coinbase,
				0,
				common.Big0,
				e.callGas,
				common.Big0,
				signerData,
				false,
			)

			evm := e.createEVM(msg, chain, parent, blockst)
			gp := new(core.GasPool).AddGas(e.callGas)
			output, _, _, err := core.ApplyMessage(evm, msg, gp)
			if err != nil {
				log.Trace("Fail to get signerAddress()", "err", err)
				return err
			}

			//
			signer := common.Address{}
			err = e.dposAbi.Unpack(&signer, "signerAddress", output)
			if err != nil {
				log.Error("Failed to unpack signerAddress() call", "err", err)
				return err
			}

			if signer == addr {
				return nil
			}

			log.Trace("PoS seal compare", "addr", addr, "signer", signer)
		} else {
			log.Trace("PoS seal compare", "addr", addr, "coinbase", header.Coinbase)
		}

		return errInvalidSig
	}

	return nil
}

// Prepare initializes the consensus fields of a block header according to the
// rules of a particular engine. The changes are executed inline.
func (e *Energi) Prepare(chain ChainReader, header *types.Header) error {
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)

	if parent == nil {
		log.Error("Fail to find parent", "header", header)
		return eth_consensus.ErrUnknownAncestor
	}

	_, err := e.posPrepare(chain, header, parent)
	return err
}

func (e *Energi) posPrepare(
	chain ChainReader,
	header *types.Header,
	parent *types.Header,
) (time_target *timeTarget, err error) {
	// Clear field to be set in mining
	header.Coinbase = common.Address{}
	header.Nonce = types.BlockNonce{}

	time_target = e.calcTimeTarget(chain, parent)

	err = e.enforceTime(header, time_target)

	// Repurpose the MixDigest field
	header.MixDigest = e.calcPoSModifier(chain, header.Time, parent)

	// Diff
	header.Difficulty = e.calcPoSDifficulty(chain, header.Time, parent, time_target)

	return time_target, err
}

// Finalize runs any post-transaction state modifications (e.g. block rewards)
// and assembles the final block.
// Note: The block header and state database might be updated to reflect any
// consensus rules that happen at finalization (e.g. block rewards).
func (e *Energi) Finalize(
	chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
	uncles []*types.Header, receipts []*types.Receipt,
) (
	block *types.Block, err error,
) {
	// Do not finalize too early in mining
	if (header.Coinbase != common.Address{}) {
		err = e.govFinalize(chain, header, state, txs)
	}

	header.UncleHash = uncleHash

	return types.NewBlock(header, txs, nil, receipts), err
}

func (e *Energi) govFinalize(
	chain ChainReader,
	header *types.Header,
	state *state.StateDB,
	txs types.Transactions,
) (err error) {
	err = e.processBlockRewards(chain, header, state)
	if err == nil {
		err = e.finalizeMigration(chain, header, state, txs)
	}
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	return
}

// Seal generates a new sealing request for the given input block and pushes
// the result into the given channel.
//
// Note, the method returns immediately and will send the result async. More
// than one result may also be returned depending on the consensus algorithm.
func (e *Energi) Seal(
	chain ChainReader,
	block *types.Block,
	results chan<- *eth_consensus.SealResult,
	stop <-chan struct{},
) (err error) {
	go func() {
		header := block.Header()
		result := eth_consensus.NewSealResult(block, nil)

		if header.Number.Cmp(common.Big0) != 0 {
			success, err := e.mine(chain, header, stop)

			// NOTE: due to the fact that PoS mining may change Coinbase
			//       it is required to reprocess all transaction with correct
			//       state of the block (input parameters). This is essential
			//       for consensus and correct distribution of gas.
			if success && err == nil {
				result, err = e.recreateBlock(chain, header, block.Transactions())
			}

			if err != nil {
				log.Error("PoS miner error", "err", err)
				success = false
			}

			if !success {
				select {
				case results <- eth_consensus.NewSealResult(nil, nil):
				default:
				}
				return
			}

			header = result.Block.Header()
		}

		sighash := e.SignatureHash(header)
		log.Trace("PoS seal hash", "sighash", sighash)

		header.Signature, err = e.signerFn(header.Coinbase, sighash.Bytes())
		if err != nil {
			log.Error("PoS miner error", "err", err)
			return
		}

		result.Block = result.Block.WithSeal(header)

		select {
		case results <- result:
			log.Info("PoS seal has submitted solution", "block", result.Block.Hash())
		default:
			log.Warn("PoS seal is not read by miner", "sealhash", e.SealHash(header))
		}
	}()

	return nil
}

func (e *Energi) recreateBlock(
	chain ChainReader,
	header *types.Header,
	txs types.Transactions,
) (
	result *eth_consensus.SealResult, err error,
) {
	var (
		usedGas = new(uint64)
		gp      = new(core.GasPool).AddGas(header.GasLimit)

		bc *core.BlockChain
		ok bool
	)

	stdb, err := chain.GetStateDB()
	if err != nil {
		return nil, err
	}

	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return nil, eth_consensus.ErrUnknownAncestor
	}

	blstate, err := state.New(parent.Root, stdb)
	if err != nil {
		return nil, err
	}

	vmc := &vm.Config{}
	if bc, ok = chain.(*core.BlockChain); ok {
		vmc = bc.GetVMConfig()
	}

	receipts := make(types.Receipts, 0, len(txs))

	for i, tx := range txs {
		blstate.Prepare(tx.Hash(), common.Hash{}, i)
		receipt, _, err := core.ApplyTransaction(
			chain.Config(), bc, nil,
			gp, blstate, header, tx, usedGas, *vmc)
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, receipt)
	}

	header.GasUsed = *usedGas
	header.Bloom = types.Bloom{}

	block, err := e.Finalize(chain, header, blstate, txs, nil, receipts)
	return eth_consensus.NewSealResult(block, blstate), err
}

// SealHash returns the hash of a block prior to it being sealed.
func (e *Energi) SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	rlp.Encode(hasher, []interface{}{
		// NOTE: commented parts are part of "mining" process
		header.ParentHash,
		header.UncleHash,
		//header.Coinbase,
		//header.Root,
		header.TxHash,
		//header.ReceiptHash,
		//header.Bloom,
		//header.Difficulty,
		header.Number,
		header.GasLimit,
		//header.GasUsed,
		//header.Time,
		//header.Extra,
		//header.MixDigest,
		//header.Nonce,
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

func (e *Energi) SetMinerCB(accountsFn AccountsFn, signerFn SignerFn) {
	if e.signerFn != nil {
		panic("Callbacks must be set only once!")
	}

	e.signerFn = signerFn
	e.accountsFn = accountsFn
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have.
func (e *Energi) CalcDifficulty(chain ChainReader, time uint64, parent *types.Header) *big.Int {
	time_target := e.calcTimeTarget(chain, parent)
	return e.calcPoSDifficulty(chain, time, parent, time_target)
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
