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
package common

import (
	"errors"
	"io/ioutil"

	"energi.world/core/gen3/common"
)

// The is the default simnet genesis block. Its written to the simnet datadir
// for custom editting consecutive node restarts.
const energiSimnetChainConfig = `{
	"alloc": {},
	"difficulty": "0x20000",
	"gasLimit"  : "0x2fefd8",
	"timestamp" : "1586995092",
	"nonce"		: "0x00",
	"extraData"	: "",
	"coinbase"	: "0x0000000000000000000000000000000000522222",
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
			"backboneAddress": "0x0000000000000000000000000000000000500021",
			"migrationSigner": "0x0000000000000000000000000000000000500022",
			"ebiSigner"      : "0x0000000000000000000000000000000000500023",
			"cppSigner"      : "0x0000000000000000000000000000000000500024"
		},
		"superblockCycle"    	: 600,
		"mnRequireValidation"	: 5,
		"mnValidationPeriod" 	: 5,
		"mnCleanupPeriod"    	: 86400,
		"mnEverCollateral"   	: 30000000000000000000000,
		"mnRewardsPerBlock"  	: 10
	}
}`

// CreateEnergiSimnetGenesisBlock creates a simnet genesis block. It only creates
// the file if it doesn't exist in the predefined file.
func CreateEnergiSimnetGenesisBlock(filepath string) error {
	if !common.FileExist(filepath) {
		if err := ioutil.WriteFile(filepath, []byte(energiSimnetChainConfig), 0600); err != nil {
			return errors.New("Failed to write the simnet genesis block json file")
		}
	}
	return nil
}
