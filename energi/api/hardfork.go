// Copyright 2020 The Energi Core Authors
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

// Package ethapi implements the general Ethereum API functions.

package api

import (
	"context"
	"fmt"
	"math/big"

	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/common/hexutil"
	"energi.world/core/gen3/log"
	"energi.world/core/gen3/rpc"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_common "energi.world/core/gen3/energi/common"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	// maxHardforkNameSize defines the max length in bytes a hf name can have.
	maxHardforkNameSize = 32
	// minHardforkNameSize defines the max length in bytes a hf name can have.
	minHardforkNameSize = 5
)

// HardforkRegistryAPI is holds the data required to access the API. It has a
// cache that temporarily holds regularly accessed data.
type HardforkRegistryAPI struct {
	backend Backend
	hfCache *energi_common.CacheStorage
}

// NewHardforkRegistryAPI returns a new HardforkRegistryAPI instance. It also
// pre-fetches the latest list of the hardforks available in the system.
func NewHardforkRegistryAPI(b Backend) *HardforkRegistryAPI {
	r := &HardforkRegistryAPI{
		backend: b,
		hfCache: energi_common.NewCacheStorage(),
	}

	b.OnSyncedHeadUpdates(func() {
		r.ListHardforks()
	})
	return r
}

func registry(
	backend Backend,
	dst common.Address,
	password *string,
) (session *energi_abi.IHardforkRegistrySession, err error) {
	contract, err := energi_abi.NewIHardforkRegistry(
		energi_params.Energi_HardforkRegistryV1, backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IHardforkRegistrySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending:  true,
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		TransactOpts: bind.TransactOpts{
			From:     dst,
			Signer:   createSignerCallback(backend, password),
			Value:    common.Big0,
			GasLimit: blacklistCallGas,
		},
	}
	return
}

// HardforkInfo defines the hardfork payload information returned.
type HardforkInfo struct {
	BlockNo    *hexutil.Big
	Name       string
	BlockHash  common.Hash
	SWFeatures *hexutil.Big
	SWVersion  string
}

// ListHardforks returns a list of the latest hardfork payload information.
// It caches the last fetched data till a new block is mined.
func (hf *HardforkRegistryAPI) ListHardforks() (res []*HardforkInfo, err error) {
	data, err := hf.hfCache.Get(hf.backend, hf.listHardforks)
	if err != nil || data == nil {
		log.Error("ListHardforks failed", "err", err)
		return
	}

	res = data.([]*HardforkInfo)
	return
}

// listHardforks queries the actual hardforks information from the contracts.
func (hf *HardforkRegistryAPI) listHardforks(num *big.Int) (interface{}, error) {
	registry, err := energi_abi.NewIHardforkRegistryCaller(
		energi_params.Energi_HardforkRegistry, hf.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Creating NewIHardforkRegistryCaller Failed", "err", err)
		return nil, err
	}

	callOpts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}

	blockNos, err := registry.Enumerate(callOpts)
	if err != nil {
		log.Error("Running Enumerate Failed", "err", err)
		return nil, err
	}

	resp := make([]*HardforkInfo, 0, len(blockNos))
	for _, blockNo := range blockNos {
		data, err := registry.GetByBlockNo(callOpts, blockNo)
		if err != nil {
			log.Error("Running GetByBlockNo Failed", "err", err)
			return nil, err
		}

		resp = append(resp, &HardforkInfo{
			BlockNo:    (*hexutil.Big)(blockNo),
			Name:       decodeToString(data.Name),
			BlockHash:  common.BytesToHash(data.BlockHash[:]),
			SWFeatures: (*hexutil.Big)(data.SwFeatures),
			SWVersion:  energi_common.SWVersionIntToString(data.SwFeatures),
		})
	}

	return resp, nil
}

// GenerateHardfork creates and updates the hardfork information.
// It validates the block number and the hardfork name used as parameters.
func (hf *HardforkRegistryAPI) GenerateHardfork(
	blockNo *hexutil.Big,
	name string,
	dst common.Address,
	password *string,
) error {
	switch {
	case blockNo.ToInt().Cmp(common.Big0) < 1:
		return fmt.Errorf("Hardfork on the genesis block is not supportted")

	case len([]byte(name)) > maxHardforkNameSize:
		return fmt.Errorf("Hardfork name is too long")

	case len([]byte(name)) < minHardforkNameSize:
		return fmt.Errorf("Hardfork name is too long")

	default:
		swFeatures := (*hexutil.Big)(energi_common.SWVersionToInt())
		return hf.generateHardfork(blockNo, name, swFeatures, dst, password)
	}
}

// generateHardfork generates the actual hardfork. It also checks if the block
// number is within its block finalization period where if affirmative the
// hardfork is finalized.
func (hf *HardforkRegistryAPI) generateHardfork(
	blockNo *hexutil.Big,
	name string,
	swFeatures *hexutil.Big,
	dst common.Address,
	password *string,
) error {
	registry, err := registry(hf.backend, dst, password)
	if err != nil {
		return err
	}

	blockHash := common.Hash{}
	block, err := hf.backend.BlockByNumber(context.Background(),
		rpc.BlockNumber(blockNo.ToInt().Int64()))
	if err == nil && block != nil {
		// Its time to finalize the hardfork now.
		blockHash = block.Hash()
	}

	tx, err := registry.Propose(blockNo.ToInt(), encodeToString(name),
		blockHash, swFeatures.ToInt())
	if err != nil {
		return err
	}

	if tx != nil {
		if tx != nil {
			log.Info("hardfork creation tx: %v will be processed in the next block",
				tx.Hash().String())
		}
	}
	return nil
}

// GetHardforkByName returns the hardfork info associated with the provided name.
func (hf *HardforkRegistryAPI) GetHardforkByName(name string) (*HardforkInfo, error) {
	registry, err := energi_abi.NewIHardforkRegistryCaller(
		energi_params.Energi_HardforkRegistry, hf.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Creating NewIHardforkRegistryCaller Failed", "err", err)
		return nil, err
	}
	callOpts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}

	data, err := registry.GetByName(callOpts, encodeToString(name))
	if err != nil {
		log.Error("Running GetName Failed", "err", err)
		return nil, err
	}

	resp := &HardforkInfo{
		BlockNo:    (*hexutil.Big)(data.BlockNo),
		Name:       name,
		BlockHash:  common.BytesToHash(data.BlockHash[:]),
		SWFeatures: (*hexutil.Big)(data.SwFeatures),
		SWVersion:  energi_common.SWVersionIntToString(data.SwFeatures),
	}

	return resp, nil
}

// GetHardforkByBlockNo returns the hardfork info identified by the provided blockno.
func (hf *HardforkRegistryAPI) GetHardforkByBlockNo(blockNo *big.Int) (*HardforkInfo, error) {
	registry, err := energi_abi.NewIHardforkRegistryCaller(
		energi_params.Energi_HardforkRegistry, hf.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Creating NewIHardforkRegistryCaller Failed", "err", err)
		return nil, err
	}

	callOpts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}

	data, err := registry.GetByBlockNo(callOpts, blockNo)
	if err != nil {
		log.Error("Running GetName Failed", "err", err)
		return nil, err
	}

	resp := &HardforkInfo{
		BlockNo:    (*hexutil.Big)(blockNo),
		Name:       decodeToString(data.Name),
		BlockHash:  common.BytesToHash(data.BlockHash[:]),
		SWFeatures: (*hexutil.Big)(data.SwFeatures),
		SWVersion:  energi_common.SWVersionIntToString(data.SwFeatures),
	}

	return resp, nil
}

// encodeToString converts the string provided to a bytes32 bytes array.
func encodeToString(data string) [32]byte {
	value := [32]byte{}
	copy(value[:], []byte(data))
	return value
}

// decodeToString converts the bytes32 bytes array back to the original string.
func decodeToString(data [32]byte) string {
	return string(data[:])
}

// DropHardfork drops any hardfork that is yet to be finalized.
func (hf *HardforkRegistryAPI) DropHardfork(
	blockNo *hexutil.Big,
	dst common.Address,
	password *string,
) error {
	registry, err := registry(hf.backend, dst, password)
	if err != nil {
		return err
	}

	tx, err := registry.Remove(blockNo.ToInt())
	if err != nil {
		return fmt.Errorf("Dropping the hardfork at block %v failed. Error: %v",
			blockNo.String(), err)
	}

	if tx != nil {
		log.Info("hardfork removal tx: %v will be processed in the next block",
			tx.Hash().String())
	}
	return nil
}
