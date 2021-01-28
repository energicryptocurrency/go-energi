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
	"errors"
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

const (
	// maxHardforkNameSize defines the max length in bytes a hf name can have.
	maxHardforkNameSize = 32
	// minHardforkNameSize defines the min length in bytes a hf name can have.
	minHardforkNameSize = 1
)

// HardforkRegistryAPI is holds the data required to access the API. It has a
// cache that temporarily holds regularly accessed data.
type HardforkRegistryAPI struct {
	backend   Backend
	hfCache   *energi_common.CacheStorage
	proxyAddr common.Address
}

// NewHardforkRegistryAPI returns a new HardforkRegistryAPI instance. It also
// pre-fetches the latest list of the hardforks available in the system.
func NewHardforkRegistryAPI(b Backend) *HardforkRegistryAPI {
	r := &HardforkRegistryAPI{
		backend:   b,
		hfCache:   energi_common.NewCacheStorage(),
		proxyAddr: b.ChainConfig().HardforkRegistryProxyAddress,
	}

	// use the default proxy address if we don't have it from ChainConfig
	emptyAddr := common.Address{}
	if r.proxyAddr == emptyAddr {
		r.proxyAddr = energi_params.Energi_HardforkRegistry
	}

	b.OnSyncedHeadUpdates(func() {
		r.ListHardforks()
	})
	return r
}

func registrySession(
	backend Backend,
	dst, proxyAddr common.Address,
	password *string,
) (session *energi_abi.IHardforkRegistrySession, err error) {
	contract, err := energi_abi.NewIHardforkRegistry(proxyAddr,
		backend.(bind.ContractBackend))
	if err != nil {
		log.Error("Creating NewIHardforkRegistryCaller Failed", "err", err)
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

func registryCaller(
	backend Backend,
	proxyAddr common.Address,
) (*energi_abi.IHardforkRegistryCaller, *bind.CallOpts, error) {
	registry, err := energi_abi.NewIHardforkRegistryCaller(proxyAddr,
		backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Creating NewIHardforkRegistryCaller Failed", "err", err)
		return nil, nil, err
	}

	callOpts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}

	return registry, callOpts, nil
}

// HardforkInfo defines the hardfork payload information returned.
type HardforkInfo struct {
	BlockNumber    *hexutil.Big
	Name       string
	BlockHash  common.Hash
	SWFeatures *hexutil.Big
	SWVersion  string
}
//
// ListHardforks returns a list of the latest hardfork payload information.
// It caches the last fetched data till a new block is mined.
func (hf *HardforkRegistryAPI) ListHardforks() (res []*HardforkInfo, err error) {
	data, err := hf.hfCache.Get(hf.backend, hf.listHardforks)
	if err != nil || data == nil {
		log.Error("ListHardforks failed", "err", err)
		return nil, err
	}

	res = data.([]*HardforkInfo)
	return
}

// listHardforks queries the actual hardforks information from the contracts.
func (hf *HardforkRegistryAPI) listHardforks(num *big.Int) (interface{}, error) {
	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return nil, err
	}

	HfNames, err := registry.Enumerate(callOpts)
	if err != nil {
		log.Error("Running Enumerate failed", "err", err)
		return nil, err
	}

	return processHfListings(HfNames, registry, callOpts)
	return nil, nil
}

// ListPendingHardforks returns a list of the latest pending hardfork payload.
// It caches the last fetched data till a new block is mined.
func (hf *HardforkRegistryAPI) ListPendingHardforks() (res []*HardforkInfo, err error) {
	data, err := hf.hfCache.Get(hf.backend, hf.listPendingHardforks)
	if err != nil || data == nil {
		log.Error("Running ListPendingHardforks failed", "err", err)
		return nil, err
	}

	res = data.([]*HardforkInfo)
	return
}

// listPendingHardforks returns a list of the current pending hardforks.
func (hf *HardforkRegistryAPI) listPendingHardforks(num *big.Int) (interface{}, error) {
	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return nil, err
	}

	HfNames, err := registry.EnumeratePending(callOpts)
	if err != nil {
		log.Error("Running EnumeratePending failed", "err", err)
		return nil, err
	}

	return processHfListings(HfNames, registry, callOpts)
}

// ListActiveHardforks returns a list of the latest active hardfork payload.
// It caches the last fetched data till a new block is mined.
func (hf *HardforkRegistryAPI) ListActiveHardforks() (res []*HardforkInfo, err error) {
	data, err := hf.hfCache.Get(hf.backend, hf.listActiveHardforks)
	if err != nil || data == nil {
		log.Error("ListActiveHardforks failed", "err", err)
		return nil, err
	}

	res = data.([]*HardforkInfo)
	return
}

// listActiveHardforks returns a list of the current Active hardforks.
func (hf *HardforkRegistryAPI) listActiveHardforks(num *big.Int) (interface{}, error) {
	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return nil, err
	}

	HfNames, err := registry.EnumerateActive(callOpts)
	if err != nil {
		log.Error("Running EnumerateActive failed", "err", err)
		return nil, err
	}

	return processHfListings(HfNames, registry, callOpts)
}

func processHfListings(
	HfNames [][32]byte,
	registry *energi_abi.IHardforkRegistryCaller,
	callOpts *bind.CallOpts,
) ([]*HardforkInfo, error) {
	resp := make([]*HardforkInfo, 0, len(HfNames))
	for _, name := range HfNames {
		data, err := registry.Get(callOpts, name)
		if err != nil {
			log.Error("Running GetHardfork failed", "err", err)
			return nil, err
		}

		resp = append(resp, &HardforkInfo{
			BlockNumber:    (*hexutil.Big)(data.BlockNumber),
			Name:       energi_common.DecodeToString(name),
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
	BlockNumber *hexutil.Big,
	name string,
	password *string,
) (common.Hash, error) {
	switch {
	case BlockNumber.ToInt().Cmp(common.Big0) < 1:
		return (common.Hash{}), fmt.Errorf("Hardfork on the genesis block is not supportted")

	case len([]byte(name)) > maxHardforkNameSize:
		return (common.Hash{}), fmt.Errorf("Hardfork name is too long")

	case len([]byte(name)) < minHardforkNameSize:
		return (common.Hash{}), fmt.Errorf("Hardfork name is too short")

	default:
		swFeatures := (*hexutil.Big)(energi_common.SWVersionToInt())
		return hf.generateHardfork(BlockNumber, name, swFeatures, password)
	}
}



// generateHardfork generates the actual hardfork.
func (hf *HardforkRegistryAPI) generateHardfork(
	BlockNumber *hexutil.Big,
	name string,
	swFeatures *hexutil.Big,
	password *string,
) (common.Hash, error) {
	txHash := common.Hash{}
	dst := hf.backend.ChainConfig().Energi.HFSigner
	registry, err := registrySession(hf.backend, dst, hf.proxyAddr, password)
	if err != nil {
		return txHash, err
	}

	tx, err := registry.Add(energi_common.EncodeToString(name), BlockNumber.ToInt(), swFeatures.ToInt())
	if err != nil {
		return txHash, err
	}

	if tx != nil {
		txHash = tx.Hash()
		log.Info("Note: please wait till HF create TX gets into a block!", "tx", txHash.Hex())
	}

	return txHash, nil
}



// FinalizeHardfork validates hardfork name parameter and calls contract finalize function
func (hf *HardforkRegistryAPI) FinalizeHardfork(name string, password *string) (common.Hash, error) {
	switch {
	case len([]byte(name)) > maxHardforkNameSize || len(name) == 0:
		return (common.Hash{}), fmt.Errorf("incorrect Hardfork name size")

	default:
		return hf.finalizeHardfork(name, password)
	}
}


/*
finalizeHardfork calls finalize contract function that checks if hardfork
with given name is finalizable and if so sets blockhash parameter for the hardfork
*/
func (hf *HardforkRegistryAPI) finalizeHardfork(name string, password *string) (common.Hash, error) {
	txHash := common.Hash{}
	dst := hf.backend.ChainConfig().Energi.HFSigner
	registry, err := registrySession(hf.backend, dst, hf.proxyAddr, password)
	if err != nil {
		return txHash, err
	}

	tx, err := registry.Finalize(energi_common.EncodeToString(name))
	if err != nil {
		return txHash, err
	}

	if tx != nil {
		txHash = tx.Hash()
		log.Info("Note: please wait till HF create TX gets into a block!", "tx", txHash.Hex())
	}

	return txHash, nil
}

// GetHardfork returns the hardfork info associated with the provided name.
func (hf *HardforkRegistryAPI) GetHardfork(name string) (*HardforkInfo, error) {
	if name == "" {
		return nil, errors.New("Empty hardfork name is not allowed")
	}

	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return nil, err
	}

	data, err := registry.Get(callOpts, energi_common.EncodeToString(name))
	if err != nil {
		log.Error("Running GetName Failed", "err", err)
		return nil, err
	}

	// If hardfork data was not found, return an error.
	if data.BlockNumber.Cmp(common.Big0) == 0 {
		return nil, fmt.Errorf("hardfork with name: (%v) was not found", name)
	}

	resp := &HardforkInfo{
		BlockNumber:    (*hexutil.Big)(data.BlockNumber),
		Name:           name,
		BlockHash:      common.BytesToHash(data.BlockHash[:]),
		SWFeatures:     (*hexutil.Big)(data.SwFeatures),
		SWVersion:      energi_common.SWVersionIntToString(data.SwFeatures),
	}

	return resp, nil
}

// DropHardfork drops any hardfork that is yet to be finalized.
func (hf *HardforkRegistryAPI) DropHardfork(
	name string,
	password *string,
) (common.Hash, error) {

	//hardfork name cannot be empty
	if name == "" {
		return common.Hash{}, errors.New("Empty hardfork name is not allowed")
	}

	txHash := common.Hash{}
	dst := hf.backend.ChainConfig().Energi.HFSigner
	registry, err := registrySession(hf.backend, dst, hf.proxyAddr, password)
	if err != nil {
		return txHash, err
	}

	tx, err := registry.Remove(energi_common.EncodeToString(name))
	if err != nil {
		return txHash, fmt.Errorf("Dropping the hardfork (%v) failed. Error: %v",
			name, err)
	}

	if tx != nil {
		txHash = tx.Hash()
		log.Info("Note: please wait till HF drop TX gets into a block!", "tx", txHash.Hex())
	}

	return txHash, nil
}

// IsHardforkActive returns true if the hardfork block has been processed otherwise
// it returns false.
func (hf *HardforkRegistryAPI) IsHardforkActive(name string) bool {
	if name == "" {
		return false
	}

	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return false
	}

	isActive, err := registry.IsActive(callOpts, energi_common.EncodeToString(name))
	if err != nil {
		log.Error("Running IsActive Failed", "err", err)
	}
	return isActive
}
