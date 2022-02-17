// Copyright 2021 The Energi Core Authors
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
	"bytes"
	"math/big"

	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/log"

	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_common "github.com/energicryptocurrency/energi/energi/common"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

// HardforkRegistryAPI holds the data required to access the API. It has a
// cache that temporarily holds regularly accessed data.
type HardforkRegistryAPI struct {
	backend   Backend
	proxyAddr common.Address
}

// HardforkInfo defines the hardfork payload information returned.
type HardforkInfo struct {
	Name        string      `json:"name"`
	BlockNumber *big.Int    `json:"block_number"`
	BlockHash   common.Hash `json:"block_hash,omitempty"`
	SWFeatures  *big.Int    `json:"sw_features"`
	SWVersion   string      `json:"sw_version"`
}

// NewHardforkRegistryAPI returns a new HardforkRegistryAPI instance. It also
// pre-fetches the latest list of the hardforks available in the system.
func NewHardforkRegistryAPI(b Backend) *HardforkRegistryAPI {
	r := &HardforkRegistryAPI{
		backend:   b,
		proxyAddr: b.ChainConfig().Energi.HardforkRegistryProxyAddress,
	}

	// use the default proxy address if we don't have it from ChainConfig
	// this is likely to be the case on networks besides mainnet / testnet
	emptyAddr := common.Address{}
	if r.proxyAddr == emptyAddr {
		r.proxyAddr = energi_params.Energi_HardforkRegistry
	}

	return r
}

func (hf *HardforkRegistryAPI) HardforkGet(name string) (info *HardforkInfo, err error) {
	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return nil, err
	}
	data, err := registry.Get(callOpts, encodeName(name))
	if err != nil {
		log.Error("HardforkRegsitryAPI::HardforkGet", "err", err)
		return nil, err
	}

	info = &HardforkInfo{
		BlockNumber: data.BlockNumber,
		Name:        name,
		BlockHash:   common.BytesToHash(data.BlockHash[:]),
		SWFeatures:  data.SwFeatures,
		SWVersion:   energi_common.SWVersionIntToString(data.SwFeatures),
	}

	return
}

func (hf *HardforkRegistryAPI) HardforkEnumerate() (hardforks []*HardforkInfo, err error) {
	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return nil, err
	}
	names, err := registry.Enumerate(callOpts)
	if err != nil {
		if err != bind.ErrNoCode {
			log.Error("HardforkRegsitryAPI::EnumeratePending", "err", err)
		}
		return nil, err
	}

	return hf.processHfListings(names)
}

func (hf *HardforkRegistryAPI) HardforkEnumeratePending() (hardforks []*HardforkInfo, err error) {
	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return nil, err
	}
	names, err := registry.EnumeratePending(callOpts)
	if err != nil {
		if err != bind.ErrNoCode {
			log.Error("HardforkRegsitryAPI::EnumeratePending", "err", err)
		}
		return nil, err
	}

	return hf.processHfListings(names)
}

func (hf *HardforkRegistryAPI) HardforkEnumerateActive() (hardforks []*HardforkInfo, err error) {
	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return nil, err
	}
	names, err := registry.EnumerateActive(callOpts)
	if err != nil {
		if err != bind.ErrNoCode {
			log.Error("HardforkRegsitryAPI::EnumerateActive", "err", err)
		}
		return nil, err
	}

	return hf.processHfListings(names)
}

func (hf *HardforkRegistryAPI) HardforkIsActive(name string) (bool, error) {
	registry, callOpts, err := registryCaller(hf.backend, hf.proxyAddr)
	if err != nil {
		return false, err
	}

	isActive, err := registry.IsActive(callOpts, encodeName(name))
	if err != nil {
		log.Error("HardforkRegistryAPI::IsActive", "err", err)
		return false, err
	}

	return isActive, nil
}

func encodeName(data string) [32]byte {
	value := [32]byte{}
	copy(value[:], []byte(data))
	return value
}

func decodeName(data [32]byte) string {
	return string(bytes.Trim(data[:], "\x00"))
}

func registryCaller(backend Backend, proxyAddr common.Address) (*energi_abi.IHardforkRegistryCaller, *bind.CallOpts, error) {
	registry, err := energi_abi.NewIHardforkRegistryCaller(proxyAddr, backend.(bind.ContractCaller))
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

func (hf *HardforkRegistryAPI) processHfListings(HfNames [][32]byte) ([]*HardforkInfo, error) {
	resp := make([]*HardforkInfo, 0, len(HfNames))
	for _, name := range HfNames {
		hf, err := hf.HardforkGet(decodeName(name))
		if err != nil {
			log.Error("HardforkRegistryAPI::HardforkGet", "err", err)
			return nil, err
		}
		resp = append(resp, hf)
	}

	return resp, nil
}
