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
	"io/ioutil"

	"github.com/energicryptocurrency/energi/common"
)

const (
	// SimnetMigrationTx defines the special simnet migration tx identified set
	// as the migration file name.
	SimnetMigrationTx = "simnet-migration"

	// This is the default simnet genesis block. Its written to the simnet datadir
	// for custom editting on consecutive node restarts.
	energiSimnetChainConfig = `
	{
		"alloc": {
			"0x0000000000000000000000000000000000000308": { "balance": "10000000000000000000" },
			"0x0000000000000000000000000000000000000310": { "balance": "120000000000000000000" },
			"0x844621b803a2a4f2b4a0d5b1cd43ffa7c2f20f94": { "balance": "150000000000000000000000000000"}
		},
		"difficulty": "0x2c",
		"gasLimit"  : "0x2fefd8",
		"timestamp" : "1586995092",
		"nonce"		: "0x00",
		"extraData"	: "",
		"coinbase"	: "0x0000000000000000000000000000000000000301",
		"number"	: "0x00",
		"gasUsed"	: "0x3D0900",
		"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000400",
		"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000500",
		"config"    : {
			"chainId" 				: 59797,
			"homesteadBlock"		: 0,
			"eip150Block"			: 0,
			"eip150Hash"			: "0x0000000000000000000000000000000000000000000000000000000000000600",
			"eip155Block"			: 0,
			"eip158Block"			: 0,
			"byzantiumBlock"		: 0,
			"constantinopleBlock"	: 0,
			"petersburgBlock"		: 0,
			"energi"  : {
				"backboneAddress": "0xbbf49b4e3e363b5cbf1074cc52f4330764d5cf91",
				"migrationSigner": "0x7d33f22bc04fd4f5041f13eb1183eff9ae7c7712",
				"ebiSigner"      : "0x128eaec174d59a7a77be8a77899efe8ff8469e76",
				"cppSigner"      : "0x21902f8414b0b4810d75ba40c651191bcf311552",
				"hfSigner"       : "0xfadfcbc05aa1f56d9ee0ee162ea9e77f49d5e45d"
			},
			"superblockCycle"    	: 100,
			"mnRequireValidation"	: 2,
			"mnValidationPeriod" 	: 2,
			"mnCleanupPeriod"    	: 86400,
			"mnEverCollateral"   	: 3000000000000000000000,
			"mnRewardsPerBlock"  	: 10,
			"hfFinalizationPeriod"  : 30
		}
	}`
)

// CreateEnergiSimnetGenesisBlock creates a simnet genesis block. It only creates
// the file if it doesn't exist in the provided file path.
func CreateEnergiSimnetGenesisBlock(filepath string) error {
	if !common.FileExist(filepath) {
		if err := ioutil.WriteFile(filepath, []byte(energiSimnetChainConfig), 0600); err != nil {
			return errors.New("Failed to write the simnet genesis block json file")
		}
	}
	return nil
}
