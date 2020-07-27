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
	"fmt"
	"math/big"

	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/common/hexutil"
	"energi.world/core/gen3/log"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_common "energi.world/core/gen3/energi/common"
	energi_params "energi.world/core/gen3/energi/params"
)

type HardforkRegistryAPI struct {
	backend Backend
	hfCache *energi_common.CacheStorage
}

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
	password *string,
	dst common.Address,
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

type HardforkInfo struct {
	BlockNo    *hexutil.Big
	Name       string
	BlockHash  common.Hash
	SWFeatures *hexutil.Big
	SWVersion  string
}

func (hf *HardforkRegistryAPI) ListHardforks() (res []*HardforkInfo, err error) {
	data, err := hf.hfCache.Get(hf.backend, hf.listHardforks)
	if err != nil || data == nil {
		log.Error("ListHardforks failed", "err", err)
		return
	}

	res = data.([]*HardforkInfo)
	return
}

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

func (hf *HardforkRegistryAPI) GenerateHardfork() error {
	return nil
}

func (hf *HardforkRegistryAPI) GetHardforkByNameOrBlockNo(searchValue interface{}) (*HardforkInfo, error) {
	registry, err := energi_abi.NewIHardforkRegistryCaller(
		energi_params.Energi_HardforkRegistry, hf.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Creating NewIHardforkRegistryCaller Failed", "err", err)
		return nil, err
	}

	switch searchValue.(type) {
	case uint64:
		blockNo := new(big.Int).SetUint64((searchValue).(uint64))
		return hf.getHardforkByBlockNo(blockNo, registry)
	case string:
		name := (searchValue).(string)
		return hf.getHardforkByName(name, registry)
	default:
		return nil, fmt.Errorf("unknown search value used")
	}
}

func (hf *HardforkRegistryAPI) getHardforkByName(name string,
	registry *energi_abi.IHardforkRegistryCaller) (*HardforkInfo, error) {
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

func encodeToString(data string) [32]byte {
	value := [32]byte{}
	copy(value[:], []byte(data))
	return value
}

func decodeToString(data [32]byte) string {
	return string(data[:])
}

func (hf *HardforkRegistryAPI) getHardforkByBlockNo(blockNo *big.Int,
	registry *energi_abi.IHardforkRegistryCaller) (*HardforkInfo, error) {
	return nil, nil
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

func (hf *HardforkRegistryAPI) DropDirtyHardfork() error {
	return nil
}
