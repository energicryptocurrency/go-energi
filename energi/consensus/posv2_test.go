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
	"flag"
	"math/big"
	"testing"

	"github.com/energicryptocurrency/energi/common"
	eth_consensus "github.com/energicryptocurrency/energi/consensus"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/core/vm"
	"github.com/energicryptocurrency/energi/crypto"
	"github.com/energicryptocurrency/energi/ethdb"
	"github.com/energicryptocurrency/energi/log"
	// "github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/params"
	"github.com/stretchr/testify/assert"

	energi_params "github.com/energicryptocurrency/energi/energi/params"
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

		now = engine.now()
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
		if !assert.True(t, tt.max <= engine.now()+30) {
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

// TODO: this test case needs to be fixed so it can properly calculate
// the PID difficulty, currently it lacks the data to do so
/*
 * Run multiple test cases
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

	type TC struct {
		Parent     int64
		Drift      int64
		Integral   int64
		Derivative int64
		Result     uint64
	}

	// the numbers below create an example with 10 second segments both
	// where target is before progressing to target is after and the
	// first and last ones are there to show the limit
	tests := []TC{
		{ Parent :343768608, Drift :34400, Integral :-9236968, Derivative :77803, Result :344003095 },
		{ Parent :344003095, Drift :-59405, Integral :-14277173, Derivative :-21953, Result :343933274 },
		{ Parent :343933274, Drift :32056, Integral :-11686883, Derivative :55532, Result :344100661 },
		{ Parent :344100661, Drift :-1025942, Integral :-4376900, Derivative :-1169771, Result :340539746 },
		{ Parent :340539746, Drift :-1239734, Integral :-10845931, Derivative :-137940, Result :340063186},
		{ Parent :340063186, Drift :-1308853, Integral :-12406416, Derivative :-65131, Result :339801488 },
		{ Parent :339801488, Drift :-2063008, Integral :-14612285, Derivative :-746167, Result :337458821 },
		{ Parent :337458821, Drift :-2606988, Integral :-12700753, Derivative :-603860, Result :335516009 },
		{ Parent :335516009, Drift :-2479996, Integral :-11595164, Derivative :75098, Result :335616497 },
		{ Parent :335616497, Drift :-2355453, Integral :-14054440, Derivative :124543, Result :335871377 },
		{ Parent :335871377, Drift :-2368705, Integral :-20693478, Derivative :46623, Result :335891373 },
		{ Parent :335891373, Drift :-2345595, Integral :-29718287, Derivative :118917, Result :336128780 },
		{ Parent :336128780, Drift :-2776293, Integral :-25209135, Derivative :-530494, Result :334396732 },
		{ Parent :334396732, Drift :-2530362, Integral :-22598439, Derivative :170086, Result :334778902 },
		{ Parent :334778902, Drift :-2737470, Integral :-32332974, Derivative :-107312, Result :334317847 },
		{ Parent :334317847, Drift :-2719512, Integral :-43440964, Derivative :137714, Result :334591996 },
		{ Parent :334591996, Drift :-2766783, Integral :-41923727, Derivative :-107148, Result :334129301 },
		{ Parent :334129301, Drift :-2519335, Integral :-39505900, Derivative :175596, Result :334527378 },
		{ Parent :334527378, Drift :-2482388, Integral :-41184081, Derivative :24972, Result :334475314 },
		{ Parent :334475314, Drift :-2673713, Integral :-45788199, Derivative :-163377, Result :333848317 },
		{ Parent :333848317, Drift :-2629883, Integral :-54615153, Derivative :131643, Result :334107959 },
		{ Parent :334107959, Drift :-3069606, Integral :-57520300, Derivative :-439723, Result :332631315 },
		{ Parent :332631315, Drift :-3076595, Integral :-66265696, Derivative :72851, Result :332691436 },
		{ Parent :332691436, Drift :-3148453, Integral :-75275662, Derivative :11970, Result :332564695 },
		{ Parent :332564695, Drift :-3317809, Integral :-78723129, Derivative :-165360, Result :331897257 },
		{ Parent :83004075, Drift :590120, Integral :39700447, Derivative :152702, Result :83494443 },
		{ Parent :83494443, Drift :473297, Integral :31794414, Derivative :2935, Result :83529120 },
		{ Parent :83529120, Drift :363760, Integral :30459748, Derivative :-85591, Result :83292650 },
		{ Parent :83292650, Drift :389924, Integral :23752471, Derivative :125962, Result :83691681 },
		{ Parent :83691681, Drift :281412, Integral :17273845, Derivative :-12709, Result :83668824 },
		{ Parent :83668824, Drift :135460, Integral :9666906, Derivative :-34180, Result :83573728 },
		{ Parent :83573728, Drift :231466, Integral :14327736, Derivative :32137, Result :83682707 },
		{ Parent :83682707, Drift :249115, Integral :13427618, Derivative :37610, Result :83808925 },
		{ Parent :83808925, Drift :469733, Integral :40373117, Derivative :-158605, Result :83359400 },
		{ Parent :83359400, Drift :446190, Integral :32793233, Derivative :92222, Result :83660652 },
		{ Parent :83660652, Drift :535198, Integral :43628343, Derivative :-58691, Result :83514368 },
	}
	// look through tests and assert result
	for i, tc := range tests {
		res := CalcPoSDifficultyV2(0, &types.Header{Difficulty: big.NewInt(tc.Parent)}, &TimeTarget{Drift: tc.Drift, Integral: tc.Integral, Derivative: tc.Derivative})
		assert.Equal(t, tc.Result, res.Uint64(), "TC %v", i)
	}
}
