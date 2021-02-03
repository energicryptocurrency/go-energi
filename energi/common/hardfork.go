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

package common

import (
	"errors"
	"math/big"
	"sync"

	"energi.world/core/gen3/common"
)

var (
	errInvalidHfName = errors.New("empty hardfork name not supported")
	errEmptyHash     = errors.New("empty hardfork block hash not supported")

	hfInfo = latestHardfork{
		hfs: make(map[string]hardforkInfo),
	}
)

type hardforkInfo struct {
	name       string
	blockNumber    *big.Int
	blockHash  common.Hash
	swFeatures *big.Int
}

// latestHardfork holds information about the latest hardfork
type latestHardfork struct {
	mtx sync.RWMutex
	hfs map[string]hardforkInfo
}

// UpdateHfActive sets the list of finalized hardforks i.e. a hardfork with an empty
// blockhash is rejected. A hardfork with an empty name is also rejected.
func UpdateHfActive(name string, blockNumber *big.Int, blockHash common.Hash, swFeatures *big.Int) error {
	hfInfo.mtx.Lock()
	defer hfInfo.mtx.Unlock()

	switch {
	case name == "":
		return errInvalidHfName
	case blockHash == common.Hash{}:
		return errEmptyHash

	default:
		hfInfo.hfs[name] = hardforkInfo{
			name:       name,
			blockNumber:    blockNumber,
			blockHash:  blockHash,
			swFeatures: swFeatures,
		}
		return nil
	}
}

// IsHfActive returns true if and only if the provided hardfork name exists and
// has been finalised (block hash has been set).
func IsHfActive(name string) (isactive bool) {
	hfInfo.mtx.RLock()
	defer hfInfo.mtx.RUnlock()

	for _, info := range hfInfo.hfs {
		if info.name == name && info.blockHash != (common.Hash{}) {
			return true
		}
	}

	return
}
