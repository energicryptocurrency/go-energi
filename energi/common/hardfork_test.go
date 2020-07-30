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
	"math/big"
	"testing"

	"energi.world/core/gen3/common"
)

func TestHardforkManagement(t *testing.T) {
	var hfName = "Ba Sing Se"
	var hfBlockNo = big.NewInt(120)
	var hfBlockHash = common.BigToHash(big.NewInt(0x111234))
	var hfSwFeatures = big.NewInt(3000600)

	// Update with accurate data
	if err := UpdateHf(hfName, hfBlockNo, hfBlockHash, hfSwFeatures); err != nil {
		t.Fatalf("expected no error but got: (%v)", err)
	}

	// Update with empty Hardfork name.
	if err := UpdateHf("", hfBlockNo, hfBlockHash, hfSwFeatures); err != errInvalidHfName {
		t.Fatalf("expected (%v) but got: (%v)", errInvalidHfName, err)
	}

	// Update with very old hardfork info
	if err := UpdateHf(hfName, big.NewInt(100), hfBlockHash, hfSwFeatures); err != errTooOldHfInfo {
		t.Fatalf("expected (%v) but got: (%v)", errTooOldHfInfo, err)
	}

	// Update with empty empty block hash
	if err := UpdateHf(hfName, hfBlockNo, common.Hash{}, hfSwFeatures); err != errEmptyHash {
		t.Fatalf("expected (%v) but got: (%v)", errEmptyHash, err)
	}

	var newHfName = "Omaha"
	// Updating a pre-existing instance should still work.
	if err := UpdateHf(newHfName, hfBlockNo, hfBlockHash, hfSwFeatures); err != nil {
		t.Fatalf("expected no error but got: (%v)", err)
	}

	blockNo := LastSetHfBlock()
	if blockNo.Cmp(hfBlockNo) != 0 {
		t.Fatalf("expected  block number to (%v) but got: (%v)", hfBlockNo, blockNo)
	}

	isActive := IsHfActive(hfName)
	if isActive {
		t.Fatalf("expected Hf name %v not to be active but it was found to be", hfName)
	}

	isActive = IsHfActive(newHfName)
	if !isActive {
		t.Fatalf("expected Hf name %v to be active but it was found not to be", newHfName)
	}

}
