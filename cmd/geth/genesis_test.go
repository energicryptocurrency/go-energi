// Copyright 2018 The Energi Core Authors
// Copyright 2016 The go-ethereum Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var customGenesisTests = []struct {
	genesis string
	query   string
	result  string
}{
	// Plain genesis file without anything extra
	// {
	// 	genesis: `{
	// 		"alloc"      : {},
	// 		"coinbase"   : "0x0000000000000000000000000000000000000000",
	// 		"difficulty" : "0x20000",
	// 		"extraData"  : "",
	// 		"gasLimit"   : "0x2fefd8",
	// 		"nonce"      : "0x0000000000000042",
	// 		"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
	// 		"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
	// 		"timestamp"  : "0x00"
	// 	}`,
	// 	query:  "eth.getBlock(0).nonce",
	// 	result: "0x0000000000000000",
	// },
	// Genesis file with an empty chain configuration (ensure missing fields work)
	{
		genesis: `{
			"alloc" : {
				"0x0000000000000000000000000000000000000001": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000002": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000003": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000004": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000005": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000006": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000007": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000008": {"balance": "0x1"}
			},
			"coinbase"   : "0x0000000000000000000000000000000000000009",
			"difficulty" : "0x20000",
			"extraData"  : "",
			"gasLimit"   : "0x2fefda",
			"nonce"      : "0x0000000000000345",
			"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000006",
			"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000007",
			"timestamp"  : "0x00",
			"config"     : {
				"chainId" : 1,
				"energi"  : {
					"backboneAddress": "0x0000000000000000000000000000000000400001",
					"migrationSigner": "0x0000000000000000000000000000000000400002",
					"ebiSigner"      : "0x0000000000000000000000000000000000400003",
					"cppSigner"      : "0x0000000000000000000000000000000000500004",
					"hfSigner"       : "0x0000000000000000000000000000000000500005"
				},
				"constantinopleBlock": 0,
				"superblockCycle"    : 10,
				"mnRequireValidation": 5,
				"mnValidationPeriod" : 5,
				"mnCleanupPeriod"    : 1209600,
				"mnEverCollateral"   : 30000000000000000000000,
				"mnRewardsPerBlock"  : 10,
				"hfFinalizationPeriod": 10
			}
		}`,
		query:  "eth.getBlock(0).nonce",
		result: "0x0000000000000345",
	},
	// Genesis file with specific chain configurations
	{
		genesis: `{
			"alloc" : {
				"0x0000000000000000000000000000000000000001": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000002": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000003": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000004": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000005": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000006": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000007": {"balance": "0x1"},
				"0x0000000000000000000000000000000000000008": {"balance": "0x1"}
			},
			"coinbase"   : "0x0000000000000000000000000000000000000002",
			"difficulty" : "0x20000",
			"extraData"  : "",
			"gasLimit"   : "0x2fefd8",
			"nonce"      : "0x0000000000000246",
			"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000400",
			"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000500",
			"timestamp"  : "0x00",
			"config"     : {
				"chainId" : 1,
				"energi"  : {
					"backboneAddress": "0x0000000000000000000000000000000000500001",
					"migrationSigner": "0x0000000000000000000000000000000000500002",
					"ebiSigner"      : "0x0000000000000000000000000000000000500003",
					"cppSigner"      : "0x0000000000000000000000000000000000500004",
					"hfSigner"       : "0x0000000000000000000000000000000000500005"
				},
				"constantinopleBlock": 0,
				"superblockCycle"    : 10,
				"mnRequireValidation": 5,
				"mnValidationPeriod" : 5,
				"mnCleanupPeriod"    : 1209600,
				"mnEverCollateral"   : 30000000000000000000000,
				"mnRewardsPerBlock"  : 10,
				"hfFinalizationPeriod": 10,
				"homesteadBlock"     : 314,
				"daoForkBlock"       : 141,
				"daoForkSupport"     : true
			}
		}`,
		query:  "eth.getBlock(0).nonce",
		result: "0x0000000000000246",
	},
}

// Tests that initializing Geth with a custom genesis block and chain definitions
// work properly.
func TestCustomGenesis(t *testing.T) {
	for i, tt := range customGenesisTests {
		// Create a temporary data directory to use and inspect later
		datadir := tmpdir(t)
		defer os.RemoveAll(datadir)

		// Initialize the data directory with the custom genesis block.
		json := filepath.Join(datadir, "genesis.json")
		if err := ioutil.WriteFile(json, []byte(tt.genesis), 0600); err != nil {
			t.Fatalf("test %d: failed to write genesis file: %v", i, err)
		}
		runGeth(t, "--datadir", datadir, "init", json).WaitExit()

		// Query the custom genesis block WITHOUT --init geth flag.
		geth := runGeth(t,
			"--datadir", datadir, "--maxpeers", "0", "--port", "0",
			"--nodiscover", "--nat", "none", "--ipcdisable",
			"--exec", tt.query, "console")

		errFormat := "Fatal: Error starting protocol stack: database already contains an incompatible genesis block...*"
		geth.ExpectRegexp(errFormat)
		geth.ExpectExit()

		// Query the custom genesis block WITH --init geth flag.
		geth = runGeth(t,
			"--datadir", datadir, "--maxpeers", "0", "--port", "0",
			"--nodiscover", "--nat", "none", "--ipcdisable",
			"--exec", tt.query, "--init", json, "console")
		geth.ExpectRegexp(tt.result)
		geth.ExpectExit()
	}
}
