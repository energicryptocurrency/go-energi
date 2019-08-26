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

package api

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

type CheckpointAPI struct {
	backend Backend
}

func NewCheckpointAPI(b Backend) *CheckpointAPI {
	return &CheckpointAPI{b}
}

const (
	checkpointCallGas uint64 = 3000000
)

func (b *CheckpointAPI) registry(
	password string,
	from common.Address,
) (
	session *energi_abi.ICheckpointRegistrySession,
	hashsig func(common.Hash) ([]byte, error),
	err error,
) {
	account := accounts.Account{Address: from}
	wallet, err := b.backend.AccountManager().Find(account)
	if err != nil {
		return nil, nil, err
	}

	contract, err := energi_abi.NewICheckpointRegistry(
		energi_params.Energi_CheckpointRegistry, b.backend.(bind.ContractBackend))
	if err != nil {
		return nil, nil, err
	}

	hashsig = func(h common.Hash) ([]byte, error) {
		return wallet.SignHashWithPassphrase(account, password, h.Bytes())
	}

	session = &energi_abi.ICheckpointRegistrySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From: from,
		},
		TransactOpts: bind.TransactOpts{
			From: from,
			Signer: func(
				signer types.Signer,
				addr common.Address,
				tx *types.Transaction,
			) (*types.Transaction, error) {
				return wallet.SignTxWithPassphrase(
					account, password, tx, b.backend.ChainConfig().ChainID)
			},
			GasLimit: checkpointCallGas,
		},
	}
	return
}

type CheckpointInfo struct {
	Number   uint64
	Hash     common.Hash
	Since    uint64
	SigCount uint64
}

type AllCheckpointInfo struct {
	Registry []CheckpointInfo
	Active   []CheckpointInfo
}

func (b *CheckpointAPI) CheckpointInfo() (res AllCheckpointInfo, err error) {
	registry, err := energi_abi.NewICheckpointRegistryCaller(
		energi_params.Energi_CheckpointRegistry, b.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return res, err
	}

	call_opts := &bind.CallOpts{}
	addresses, err := registry.Checkpoints(call_opts)
	if err != nil {
		log.Error("Failed", "err", err)
		return res, err
	}

	res.Registry = make([]CheckpointInfo, 0, len(addresses))

	for _, addr := range addresses {
		cp, err := energi_abi.NewICheckpointCaller(
			addr, b.backend.(bind.ContractCaller))
		if err != nil {
			log.Error("Failed", "err", err)
			continue
		}

		info, err := cp.Info(call_opts)
		if err != nil {
			log.Warn("Info error", "cp", addr, "err", err)
			continue
		}

		sigs, err := cp.Signatures(call_opts)
		if err != nil {
			log.Warn("Proposals error", "addr", addr, "err", err)
			continue
		}

		res.Registry = append(res.Registry, CheckpointInfo{
			Number:   info.Number.Uint64(),
			Hash:     info.Hash,
			Since:    info.Since.Uint64(),
			SigCount: uint64(len(sigs)),
		})
	}

	local := b.backend.ListCheckpoints()
	res.Active = make([]CheckpointInfo, 0, len(local))

	for _, cp := range local {
		res.Active = append(res.Active, CheckpointInfo{
			Number:   cp.Number,
			Hash:     cp.Hash,
			Since:    cp.Since,
			SigCount: uint64(0),
		})
	}

	return
}

func (b *CheckpointAPI) CheckpointPropose(
	number uint64,
	hash common.Hash,
	password string,
) (txhash common.Hash, err error) {
	registry, hashsig, err := b.registry(password, b.backend.ChainConfig().Energi.CPPSigner)
	if err != nil {
		return
	}

	bnum := new(big.Int).SetUint64(number)
	tosig, err := registry.SignatureBase(bnum, hash)
	if err != nil {
		return
	}

	sig, err := hashsig(tosig)
	if err != nil {
		return
	}

	tx, err := registry.Propose(bnum, hash, sig)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (b *CheckpointAPI) CheckpointLocal(
	number uint64,
	hash common.Hash,
) error {
	return b.backend.AddLocalCheckpoint(number, hash)
}
