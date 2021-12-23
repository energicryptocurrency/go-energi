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
	"fmt"
	"math/big"

	"github.com/energicryptocurrency/energi/params"
)

// SWVersionToInt converts the current semantic software version to an integer.
func SWVersionToInt() *big.Int {
	semver := int64(params.VersionMajor<<24 | params.VersionMinor<<16 | params.VersionPatch<<8)
	return new(big.Int).SetInt64(semver)
}

// SWVersionIntToString returns a string from the provided software version int.
func SWVersionIntToString(semver *big.Int) string {
	val := semver.Int64()
	major, minor, patch := byte(val>>24), byte(val>>16), byte(val>>8)
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
