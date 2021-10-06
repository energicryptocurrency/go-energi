// Copyright 2021 The Energi Core Authors
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

package consensus

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"energi.world/core/gen3/common"
	eth_consensus "energi.world/core/gen3/consensus"
	"energi.world/core/gen3/core"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/core/vm"
	"energi.world/core/gen3/crypto"
	"energi.world/core/gen3/ethdb"
	"energi.world/core/gen3/log"
	// "energi.world/core/gen3/log"
	"energi.world/core/gen3/params"
	"github.com/stretchr/testify/assert"

	energi_params "energi.world/core/gen3/energi/params"
)

func TestCalculateBlockTimeEMA(t *testing.T) {
	t.Parallel()
	emaCalculated := CalculateBlockTimeEMA(testDataBlockTimes, energi_params.BlockTimeEMAPeriod)

	// check a known value
	emaExpected58 := uint64(59808819)
	if emaCalculated[58] != emaExpected58 {
		t.Log("EMA mismatch - expected", emaExpected58, "got", emaCalculated[58])
		t.FailNow()
	}

	// check the entire series
	for i := range emaCalculated {
		if emaCalculated[i] != testDataBlockTimeEMA[i] {
			t.Log("EMA mismatch at index", i, "- expected", testDataBlockTimeEMA[i], "got", emaCalculated[i])
			t.FailNow()
		}
	}
}

func TestCalculateBlockTimeDrift(t *testing.T) {
	t.Parallel()
	blockDrift := CalculateBlockTimeDrift(testDataBlockTimeEMA)

	// check a known value
	blockDriftExpected58 := int64(191181)
	if blockDrift[58] != blockDriftExpected58 {
		t.Log("Block Time Drift mismatch - expected", blockDriftExpected58, "got", blockDrift[58])
		t.FailNow()
	}

	// check the entire series
	for i := range blockDrift {
		if blockDrift[i] != testDataBlockTimeDrift[i] {
			t.Log("Block Time Drift mismatch at index", i, "- expected", testDataBlockTimeDrift[i], "got", blockDrift[i])
			t.FailNow()
		}
	}
}

func TestCalculateBlockTimeIntegral(t *testing.T) {
	t.Parallel()
	integral := CalculateBlockTimeIntegral(testDataBlockTimeDrift)
	integralExpected := int64(363392749)
	// check a known value
	if integral != integralExpected {
		t.Log("Block Time Integral mismatch - expected", integralExpected, "got", integral)
		t.FailNow()
	}
}

func TestCalculateBlockTimeDerivative(t *testing.T) {
	t.Parallel()
	derivative := CalculateBlockTimeDerivative(testDataBlockTimeDrift)
	// check a known value
	derivativeExpected58 := int64(testDataBlockTimeDrift[59] - testDataBlockTimeDrift[58])
	if derivative[58] != derivativeExpected58 {
		t.Log("Block Time Drift mismatch - expected", derivativeExpected58, "got", derivative[58])
		t.FailNow()
	}

	// check the entire series
	for i := range derivative {
		if derivative[i] != testDataBlockTimeDerivative[i] {
			t.Log("Block Time Drift mismatch at index", i, "- expected", testDataBlockTimeDerivative[i], "got", derivative[i])
			t.FailNow()
		}
	}
}

/*
 * Create a mock chain
 * For 150 iterations, create a block
 * After each block is Finalized, call CalcTimeTargetV2
 * Analyze the populated TimeTargetV2 struct
 * Assertions:
 * - Target is correct
 * - Block time is correct (current header time - parent header time)
 */
func TestPoSChainV2(t *testing.T) {
	// t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)
	flag.Parse()
	//log.Root().SetHandler(
	//	log.LvlFilterHandler(
	//		3,
	//		log.StreamHandler(
	//			os.Stderr,
	//			log.TerminalFormat(false),
	//		),
	//	),
	//)
	// this enables code location printing
	log.PrintOrigins(true)

	results := make(chan *eth_consensus.SealResult, 1)
	stop := make(chan struct{})

	addresses, signers, alloc, migrationSigner := generateAddresses(120)

	testdb := ethdb.NewMemDatabase()
	engine := New(&params.EnergiConfig{MigrationSigner: migrationSigner}, testdb)
	var header *types.Header

	engine.testing = true
	engine.SetMinerCB(
		func() []common.Address {
			if header.Number.Uint64() == 1 {
				return []common.Address{
					energi_params.Energi_MigrationContract,
				}
			}

			return addresses
		},
		func(addr common.Address, hash []byte) ([]byte, error) {
			if header.Number.Uint64() == 1 {
				return crypto.Sign(hash,
					signers[migrationSigner])
			}
			return crypto.Sign(hash, signers[addr])
		},
		func() int { return 1 },
		func() bool { return true },
	)

	chainConfig := *params.EnergiTestnetChainConfig
	chainConfig.Energi = &params.EnergiConfig{
		MigrationSigner: migrationSigner,
	}

	var (
		gspec = &core.Genesis{
			Config:     &chainConfig,
			GasLimit:   8000000,
			Timestamp:  1000,
			Difficulty: big.NewInt(1),
			Coinbase:   energi_params.Energi_Treasury,
			Alloc:      alloc,
			Xfers:      core.DeployEnergiGovernance(&chainConfig),
		}
		genesis = gspec.MustCommit(testdb)

		now = uint64(time.Now().Unix())
		_   = now
	)
	var err error
	_ = err
	var chain *core.BlockChain
	chain, err = core.NewBlockChain(testdb, nil, &chainConfig, engine, vm.Config{}, nil)
	if !assert.Empty(t, err) {
		log.Debug("failed")
		t.FailNow()
	}
	defer chain.Stop()

	// --
	_, err = chain.InsertChain([]*types.Block{genesis})
	if !assert.Empty(t, err) {
		log.Debug("failed")

		t.FailNow()
	}

	parent := chain.GetHeaderByHash(genesis.Hash())
	if !assert.NotEmpty(t, parent) {
		log.Debug("failed")

		t.FailNow()
	}

	iterCount := 150

	engine.diffFn = func(uint64, *types.Header,
		*TimeTarget) *big.Int {
		return common.Big1
	}

	for i := 1; i < iterCount; i++ {
		number := new(big.Int).Add(parent.Number, common.Big1)

		// ---
		header = &types.Header{
			ParentHash: parent.Hash(),
			Coinbase:   common.Address{},
			GasLimit:   parent.GasLimit,
			Number:     number,
			Time:       parent.Time,
		}
		log.Debug("calculating state")
		blstate := chain.CalculateBlockState(header.ParentHash, parent.Number.Uint64())
		if !assert.NotEmpty(t, blstate) {
			log.Debug("failed")

			t.FailNow()
		}
		log.Debug("preparing engine")
		err = engine.Prepare(chain, header)
		if !assert.Empty(t, err) {
			log.Debug("failed")

			t.FailNow()
		}
		if !assert.NotEmpty(t, header.Difficulty) {
			log.Debug("failed")

			t.FailNow()
		}
		txs := types.Transactions{}
		receipts := []*types.Receipt{}
		if i == 1 {
			tx := migrationTx(
				types.NewEIP155Signer(chainConfig.ChainID), header,
				&snapshot{
					Txouts: []snapshotItem{
						{
							Owner:  "t6vtJKxdjaJdofaUrx7w4xUs5bMcjDq5R2",
							Amount: big.NewInt(10228000000),
							Atype:  "pubkeyhash",
						},
					},
				}, engine,
			)
			var receipt *types.Receipt
			receipt, _, err = core.ApplyTransaction(
				&chainConfig, chain, &header.Coinbase,
				new(core.GasPool).AddGas(header.GasLimit),
				blstate, header, tx,
				&header.GasUsed, *chain.GetVMConfig())
			if !assert.Empty(t, err) {
				log.Debug("failed")

				t.FailNow()
			}
			txs = append(txs, tx)
			receipts = append(receipts, receipt)
		}
		var block *types.Block
		var finalizedReceipts []*types.Receipt
		block, finalizedReceipts, err = engine.Finalize(
			chain, header, blstate, txs, nil, receipts)
		if !assert.Empty(t, err) {
			log.Debug("failed")

			t.FailNow()
		}

		if i == 1 {
			if !assert.Equal(t, 1, len(finalizedReceipts)) {
				log.Debug("failed")

				t.FailNow()
			}
		} else {
			if !assert.Empty(t, finalizedReceipts) {
				log.Debug("failed")

				t.FailNow()
			}
		}

		// ---
		log.Debug("sealing migration block")
		err = engine.Seal(chain, block, results, stop)
		if !assert.Empty(t, err) {
			log.Debug("failed")

			t.FailNow()
		}

		log.Debug("waiting for results", "number", block.Number())
		seal_res := <-results
		log.Debug("received results")
		block = seal_res.Block
		blstate = seal_res.NewState
		finalizedReceipts = seal_res.Receipts
		if !assert.NotEmpty(t, block) {
			log.Debug("block was empty")
			log.Debug("failed")

			t.FailNow()
		}
		if !assert.NotEmpty(t, blstate) {
			log.Debug("state was empty")
			log.Debug("failed")

			t.FailNow()
		}

		if !assert.NotEmpty(t, finalizedReceipts) {
			log.Debug("receipts were empty")
			log.Debug("failed")

			t.FailNow()
		}
		header = block.Header()
		// todo: this next assert fails sometimes
		// assert.NotEqual(t, parent.Coinbase, header.Coinbase, "Header %v", i)
		if !assert.NotEqual(t, parent.Coinbase, common.Address{},
			"Header %v", i) {
			log.Debug("failed")

			t.FailNow()
		}
		err = engine.VerifySeal(chain, header)
		if !assert.Empty(t, err) {
			log.Debug("failed")

			t.FailNow()
		}

		// Test consensus tx check during block processing
		// ---
		if i == 2 {
			tmptxs := block.Transactions()
			tmpheader := *header

			if !assert.Equal(t, len(tmptxs), 1) {
				log.Debug("failed")

				t.FailNow()
			}

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(), tmptxs, nil, finalizedReceipts)
			if !assert.Empty(t, err) {
				log.Debug("failed")

				t.FailNow()
			}

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(), append(tmptxs, tmptxs[len(tmptxs)-1]), nil, finalizedReceipts)
			if !assert.Equal(t, eth_consensus.ErrInvalidConsensusTx,
				err) {
				log.Debug("failed")

				t.FailNow()
			}

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(),
				append(tmptxs[:len(tmptxs)-1], tmptxs[len(tmptxs)-1].WithConsensusSender(common.Address{})),
				nil, finalizedReceipts)
			if !assert.Equal(t, eth_consensus.ErrInvalidConsensusTx,
				err) {
				log.Debug("failed")

				t.FailNow()
			}
		}

		// Time tests
		// ---
		tt := engine.calcTimeTargetV2(chain, parent)
		_ = tt
		if !assert.True(t, tt.max >= now) {
			log.Debug("failed")

			t.FailNow()
		}
		if !assert.True(t, tt.max <= uint64(time.Now().Unix()) + 30) {
			log.Debug("failed")

			t.FailNow()
		}

		if i < 60 {
			// parent header and current header must be TimeTarget.min apart(30s)
			if !assert.Equal(t, header.Time, parent.Time+30) {
				t.FailNow()
			}
			if !assert.Equal(t, tt.min, header.Time) {
				t.FailNow()
			}
			// assert.Equal(t, tt.target, header.Time+30)
		} else if i < 61 {
			if !assert.Equal(t, header.Time, genesis.Time()+3570) {
				t.FailNow()
			}
			if !assert.Equal(t, header.Time, parent.Time+1800) {
				t.FailNow()
			}
			if !assert.Equal(t, tt.min, header.Time-1770) {
				t.FailNow()
			}
			// todo: this test is getting different numbers for
			//  each block
			// if !assert.Equal(t, tt.target, parent.Time-122){
			// 	log.Debug(fmt.Sprintln(tt.target,
			// 		parent.Time-120))
			// 	t.FailNow()
			// }
		} else if i < 62 {
			if !assert.Equal(t, header.Time, genesis.Time()+3600) {
				t.FailNow()
			}
		}

		assert.True(t, parent.Time < tt.min, "Header %v", i)

		//		assert.Empty(t, engine.enforceTime(header, tt))
		//		assert.Empty(t, engine.checkTime(header, tt))

		_, err = chain.WriteBlockWithState(block, finalizedReceipts, blstate)
		if !assert.Empty(t, err) {
			log.Debug("failed")

			t.FailNow()
		}

		parent = header
	}
}

/*
 * Run multiple test cases which are found in intervalgen/PoSV2_test_cases.json
	* In order to generate a new testcase set
			cd intervalgen
			go build main.go
 * Call CalcPoSDifficultyV2, analyzing the result
 * Assertions:
 * - Difficulty is correct
*/

func TestPoSDiffV2(t *testing.T) {
	// t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)
	// log.Root().SetHandler(
	// 	log.LvlFilterHandler(
	// 		5,
	// 		log.StreamHandler(
	// 			os.Stderr,
	// 			log.TerminalFormat(false),
	// 		),
	// 	),
	// )
	// this enables code location printing
	log.PrintOrigins(true)

	cases := readJsonPoSV2TestCases()

	for i, tc := range cases {
		parent := &types.Header{
			Difficulty: big.NewInt(tc.Parent),
		}
		tt := &TimeTarget{
			Drift:      tc.Drift,
			Integral:   tc.Integral,
			Derivative: tc.Derivative,
		}

		res := CalcPoSDifficultyV2(tc.Time, parent, tt).Int64()
		assert.Equal(t, tc.Result, res, "TC %v", i)
	}
}

type PoSDiffV2TestCase struct {
	Time       uint64
	Parent     int64
	Drift      int64
	Integral   int64
	Derivative int64
	Result     int64
}

func readJsonPoSV2TestCases() (result []PoSDiffV2TestCase) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(dir + "/intervalgen/PoSV2_test_cases.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	return
}
