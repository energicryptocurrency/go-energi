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
	"github.com/energicryptocurrency/energi/common"
	eth_params "github.com/energicryptocurrency/energi/params"
)

// map Genesis to map of checkpoints
var EnergiCheckpoints = map[common.Hash]map[uint64]common.Hash{
	eth_params.MainnetGenesisHash: map[uint64]common.Hash{
		uint64(38283):  common.HexToHash("0xbcdaf97b5fa9041b34afcee7ffa85ce88e8e1a9eada2855b5e3ee55b8153f70d"),
		uint64(38284):  common.HexToHash("0x2a724c7ae36f26b61f5702b0b2099c37061b105b619f250991b573e8d32d63a0"),
		uint64(324911): common.HexToHash("0x6c4d2ec49ebe49135793733b8fb43adf5e0d883d96a8d65dbe98e1debcc77c3a"),
		uint64(324913): common.HexToHash("0x7ea1dd2b9b737ea22232dc55378ea9c11aa8e5d1712ad650f1cf2fa53fe820c8"),
		uint64(325054): common.HexToHash("0x9d33948d3411b8276696e2a4422b412830a4083082d523463bd79ab572803ddd"),
		uint64(563400): common.HexToHash("0xfc5725db1869a1ca7a19769e376c826792719fe8d361f5fc9d8407725faae3f5"),
		uint64(765780): common.HexToHash("0x265f3b69f31a6f077db1acd84680dd8177eedbd7167b510997cfaa92fd31358c"),
	},
	eth_params.TestnetGenesisHash: {
		uint64(100000):  common.HexToHash("0x581e691b64ffbec38b484796377581443ada724db6fbeb6e611573fbea04ecd2"),
		uint64(200000):  common.HexToHash("0xe791464c99d0921c58d2fdc3b93d0e4d42b4d5f2a3b499f0982cb695483b4b82"),
		uint64(300000):  common.HexToHash("0x3bab5c007ede9d287e3ccccbb1534c3f580b7ea52ed4f1097500bc6338bea7ff"),
		uint64(400000):  common.HexToHash("0x563f3fd48dc9b6b87a49ba0f69a9ab6bbccbc3a1394b6bc74794bf3377d8210b"),
		uint64(500000):  common.HexToHash("0xed8b151dd83abcac55246eb2447cb70444b00cff9b25212947b4670b01844888"),
		uint64(600000):  common.HexToHash("0x1b35b6a6ae757157401fe326fd87d183b6cf1e5dd490cc58e15b9038dedd8faf"),
		uint64(700000):  common.HexToHash("0x10b39c50ac7e0a7f4c24c471d70a20ad64c860ee4f79ffaf7cfa6e3dba11008b"),
		uint64(800000):  common.HexToHash("0x4bc90008915235989ef86ebcc7668dd67d83f05460d248b59fe6c2b45e0be778"),
	},
}
