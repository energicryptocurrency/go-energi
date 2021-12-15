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

package api

import (
	"context"
	"math/big"

	"github.com/energicryptocurrency/energi/accounts"
	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/rpc"

	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_common "github.com/energicryptocurrency/energi/energi/common"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

type CheckpointAdminAPI struct {
	backend Backend
}

func NewCheckpointAdminAPI(b Backend) *CheckpointAdminAPI {
	return &CheckpointAdminAPI{b}
}

func (b *CheckpointAdminAPI) CheckpointLocal(
	number uint64,
	hash common.Hash,
) error {
	return b.backend.AddLocalCheckpoint(number, hash)
}

type CheckpointRegistryAPI struct {
	backend   Backend
	cpCache   *energi_common.CacheStorage
	proxyAddr common.Address
}

func NewCheckpointRegistryAPI(b Backend) *CheckpointRegistryAPI {
	r := &CheckpointRegistryAPI{
		backend: b,
		cpCache: energi_common.NewCacheStorage(),
		proxyAddr: energi_params.Energi_CheckpointRegistry,
	}
	b.OnSyncedHeadUpdates(func() {
		r.CheckpointInfo()
	})
	return r
}

const (
	checkpointCallGas uint64 = 3000000
)

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

// returns all the checkpoints from cache (or contract)
func (b *CheckpointRegistryAPI) CheckpointInfo() (res *AllCheckpointInfo, err error) {
	var data interface{}
	data, err = b.cpCache.Get(b.backend, b.checkpointInfo)
	if err != nil || data == nil {
		log.Error("CheckpointInfo failed", "err", err)
		return
	}

	res = data.(*AllCheckpointInfo)
	return
}

// lists the existing checkpoint addresses from contract
func (b *CheckpointRegistryAPI) Checkpoints() ([]common.Address, error) {
	// initialize contract caller
	registry, callOpts, err := checkpointRegistryCaller(b.backend, b.proxyAddr)
	if err != nil {
		return nil, err
	}

	// call "checkpoins" function on contract
	checkpointAddresses, err := registry.Checkpoints(callOpts)
	if err != nil {
		log.Error("CheckpointRegsitryAPI::Checkpoints", "err", err)
		return nil, err
	}

	return checkpointAddresses, nil
}

// executes command to propose checkpoint
func (b *CheckpointRegistryAPI) CheckpointRemove(
	number uint64,
	hash common.Hash,
	password *string,
) (txhash common.Hash, err error) {
	// request registry caller and signer
	registry, _, err := b.registry(password, b.backend.ChainConfig().Energi.CPPSigner)
	if err != nil {
		return
	}

	tx, err := registry.Remove(new(big.Int).SetUint64(number), hash)
	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}
	return
}


// returns existing checkpoints' info
func (b *CheckpointRegistryAPI) checkpointInfo(num *big.Int) (interface{}, error) {
	// get existing checkpoint addresses from contract
	addresses, err := b.Checkpoints()
	if err != nil {
		log.Error("Failed", "err", err)
		return nil, err
	}

	res := &AllCheckpointInfo{}
	res.Registry = make([]CheckpointInfo, 0, len(addresses))

	for _, addr := range addresses {
		cp, err := energi_abi.NewICheckpointV2Caller(
			addr, b.backend.(bind.ContractCaller))
		if err != nil {
			log.Error("Failed", "err", err)
			continue
		}

		info, err := cp.Info(&bind.CallOpts{Pending: true, GasLimit: energi_params.UnlimitedGas})
		if err != nil {
			log.Warn("Info error", "cp", addr, "err", err)
			continue
		}

		sigs, err := cp.Signatures(&bind.CallOpts{Pending: true, GasLimit: energi_params.UnlimitedGas})
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
			SigCount: cp.SigCount,
		})
	}

	return res, nil
}

// initializes registry which is used to propose checkpoint
func (b *CheckpointRegistryAPI) registry(password *string, from common.Address) (
	session *energi_abi.ICheckpointRegistryV2Session,
	hashsig func(common.Hash) ([]byte, error),
	err error,
) {
	contract, err := energi_abi.NewICheckpointRegistryV2(energi_params.Energi_CheckpointRegistry, b.backend.(bind.ContractBackend))
	if err != nil {
		return nil, nil, err
	}

	hashsig = func(h common.Hash) ([]byte, error) {
		account := accounts.Account{Address: from}
		wallet, err := b.backend.AccountManager().Find(account)
		if err != nil {
			return nil, err
		}

		if password == nil {
			return wallet.SignHash(account, h.Bytes())
		}

		return wallet.SignHashWithPassphrase(account, *password, h.Bytes())
	}

	session = &energi_abi.ICheckpointRegistryV2Session{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending:  true,
			From:     from,
			GasLimit: energi_params.UnlimitedGas,
		},
		TransactOpts: bind.TransactOpts{
			From:     from,
			Signer:   createSignerCallback(b.backend, password),
			GasLimit: checkpointCallGas,
		},
	}
	return
}

// executes command to propose checkpoint
func (b *CheckpointRegistryAPI) CheckpointPropose(
	number uint64,
	hash common.Hash,
	password *string,
) (txhash common.Hash, err error) {
	// check if proposed block number hash corresponds to the block in current chain
	if h, _ := b.backend.HeaderByNumber(context.Background(), rpc.BlockNumber(number)); h == nil {
		log.Error("Block not found on local node", "number", number)
		return
	} else if h.Hash() != hash {
		log.Error("Block mismatch on local node", "number", number, "hash", hash, "block", h.Hash())
		return
	}

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

	// NOTE: compatibility with ecrecover opcode.
	sig[64] += 27

	tx, err := registry.Propose(bnum, hash, sig)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

// function returns already initialized checkpoint Icheckpoint registry caller
func checkpointRegistryCaller(backend Backend, proxyAddr common.Address) (*energi_abi.ICheckpointRegistryV2Caller, *bind.CallOpts, error) {
	registry, err := energi_abi.NewICheckpointRegistryV2Caller(proxyAddr, backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Creating NewICheckpointRegistryV2Caller Failed", "err", err)
		return nil, nil, err
	}

	callOpts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}

	return registry, callOpts, nil
}
