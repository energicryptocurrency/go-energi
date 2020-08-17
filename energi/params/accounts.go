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
	"math/big"

	"energi.world/core/gen3/common"
	eth_params "energi.world/core/gen3/params"
)

type newProxyContract int

const (
	// EnergiHardforkRegistry proxy contract address identifier.
	EnergiHardforkRegistry newProxyContract = iota
)

var (
	Energi_BlockReward        = common.BigToAddress(big.NewInt(0x300))
	Energi_Treasury           = common.BigToAddress(big.NewInt(0x301))
	Energi_MasternodeRegistry = common.BigToAddress(big.NewInt(0x302))
	Energi_StakerReward       = common.BigToAddress(big.NewInt(0x303))
	Energi_BackboneReward     = common.BigToAddress(big.NewInt(0x304))
	Energi_SporkRegistry      = common.BigToAddress(big.NewInt(0x305))
	Energi_CheckpointRegistry = common.BigToAddress(big.NewInt(0x306))
	Energi_BlacklistRegistry  = common.BigToAddress(big.NewInt(0x307))
	Energi_MigrationContract  = common.BigToAddress(big.NewInt(0x308))
	Energi_MasternodeToken    = common.BigToAddress(big.NewInt(0x309))
	Energi_Blacklist          = common.BigToAddress(big.NewInt(0x30A))
	Energi_Whitelist          = common.BigToAddress(big.NewInt(0x30B))
	Energi_MasternodeList     = common.BigToAddress(big.NewInt(0x30C))
	Energi_HardforkRegistry   = common.BigToAddress(big.NewInt(0x30D)) // Only used in simnet and devnet

	Energi_BlockRewardV1        = common.BigToAddress(big.NewInt(0x310))
	Energi_TreasuryV1           = common.BigToAddress(big.NewInt(0x311))
	Energi_MasternodeRegistryV1 = common.BigToAddress(big.NewInt(0x312))
	Energi_StakerRewardV1       = common.BigToAddress(big.NewInt(0x313))
	Energi_BackboneRewardV1     = common.BigToAddress(big.NewInt(0x314))
	Energi_SporkRegistryV1      = common.BigToAddress(big.NewInt(0x315))
	Energi_CheckpointRegistryV1 = common.BigToAddress(big.NewInt(0x316))
	Energi_BlacklistRegistryV1  = common.BigToAddress(big.NewInt(0x317))
	Energi_CompensationFundV1   = common.BigToAddress(big.NewInt(0x318))
	Energi_MasternodeTokenV1    = common.BigToAddress(big.NewInt(0x319))
	Energi_HardforkRegistryV1   = common.BigToAddress(big.NewInt(0x321)) // Only used in simnet and devnet

	Energi_SystemFaucet = common.BigToAddress(big.NewInt(0x320))
	Energi_Ephemeral    = common.HexToAddress("0x457068656d6572616c")

	// NOTE: this is NOT very safe, but it optimizes significantly
	Storage_ProxyImpl = common.BigToHash(big.NewInt(0x01))
)

// newProxyContracts defines network specific proxy contract addresses that are
// set after the new proxy contracts are deployed
var newProxyContracts = map[common.Hash]map[newProxyContract]common.Address{
	eth_params.MainnetGenesisHash: {
		EnergiHardforkRegistry: common.Address{},
	},
	eth_params.TestnetGenesisHash: {
		EnergiHardforkRegistry: common.Address{},
	},
}

// defaultAddresses holds the default proxy address if the required network
// network configuration doest exist.
var defaultAddresses = map[newProxyContract]common.Address{
	// EnergiHardforkRegistry: Energi_HardforkRegistry,
	EnergiHardforkRegistry: common.HexToAddress("0xf280007b154e48adb309c278f6515f358025e484"),
}

// GetProxyContractAddress returns the required proxy address if it exists.
func GetProxyContractAddress(name newProxyContract, net common.Hash) common.Address {
	info, ok := newProxyContracts[net]
	if ok {
		addr, ok := info[name]
		if !ok {
			return common.Address{}
		}

		if addr != (common.Address{}) {
			return addr
		}
	}

	return defaultAddresses[name]
}
