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
	errTooOldHfInfo  = errors.New("Hardfork cannot be created in the past")
	errEmptyHash     = errors.New("empty hardfork block hash not supported")

	hfInfo = latestHardfork{
		hfs: make([]hardforkInfo, 0, 10),
	}
)

type hardforkInfo struct {
	name       string
	blockNo    *big.Int
	blockHash  common.Hash
	swFeatures *big.Int
}

// latestHardfork holds information about the latest hardfork
type latestHardfork struct {
	mtx sync.RWMutex
	hfs []hardforkInfo
}

// UpdateHf sets the latest supported hardfork. A hardfork with an empty
// blockhash, lesser block number in relation to the previoulsy set and software
// version lesser than the previously set is rejected. A hardwork with an empty
// name is also rejected.
func UpdateHf(name string, blockNo *big.Int, blockHash common.Hash, swFeatures *big.Int) error {
	hfInfo.mtx.Lock()
	defer hfInfo.mtx.Unlock()

	latestHF := hardforkInfo{}
	if len(hfInfo.hfs) > 0 {
		// pick the last hardfork added
		latestHF = hfInfo.hfs[len(hfInfo.hfs)-1]
	}

	switch {
	case name == "":
		return errInvalidHfName

	case latestHF.blockNo.Cmp(blockNo) < 0, latestHF.swFeatures.Cmp(swFeatures) < 0:
		return errTooOldHfInfo

	case blockHash == (common.Hash{}):
		return errEmptyHash

	default:
		for i, info := range hfInfo.hfs {
			// This case should never happen but if it happens just update the
			// pre-existing instance.
			if info.blockNo.Cmp(blockNo) == 0 {
				hfInfo.hfs[i].blockHash = blockHash
				hfInfo.hfs[i].name = name
				hfInfo.hfs[i].swFeatures = swFeatures
				return nil
			}
		}

		// Append this as new HF instance.
		hfInfo.hfs = append(hfInfo.hfs, hardforkInfo{
			name:       name,
			blockNo:    blockNo,
			blockHash:  blockHash,
			swFeatures: swFeatures,
		})
		return nil
	}
}

// LastSetHfBlock returns the last supported hardfork.
func LastSetHfBlock() *big.Int {
	hfInfo.mtx.RLock()
	defer hfInfo.mtx.RUnlock()

	if len(hfInfo.hfs) == 0 {
		return big.NewInt(0)
	}
	return hfInfo.hfs[len(hfInfo.hfs)-1].blockNo
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
