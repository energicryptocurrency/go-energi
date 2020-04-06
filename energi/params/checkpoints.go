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

package params

import (
	"energi.world/core/gen3/common"
	eth_params "energi.world/core/gen3/params"
)

// map Genesis to map of checkpoints
var EnergiCheckpoints = map[common.Hash]map[uint64]common.Hash{
	eth_params.MainnetGenesisHash: map[uint64]common.Hash{
		uint64(38283): common.HexToHash("0xbcdaf97b5fa9041b34afcee7ffa85ce88e8e1a9eada2855b5e3ee55b8153f70d"),
		uint64(38284): common.HexToHash("0x2a724c7ae36f26b61f5702b0b2099c37061b105b619f250991b573e8d32d63a0"),
	},
	eth_params.TestnetGenesisHash: {},
}
