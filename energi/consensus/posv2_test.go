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
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)
	flag.Parse()
	log.Root().SetHandler(
		log.LvlFilterHandler(
			5,
			log.StreamHandler(
				os.Stderr,
				log.TerminalFormat(false),
			),
		),
	)
	// this enables code location printing
	log.PrintOrigins(true)

	results := make(chan *eth_consensus.SealResult, 1)
	stop := make(chan struct{})

	addresses, signers, alloc, migrationSigner := generateAddresses(60)

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
	)

	chain, err := core.NewBlockChain(testdb, nil, &chainConfig, engine, vm.Config{}, nil)
	if !assert.Empty(t, err) {
		t.FailNow()
	}
	defer chain.Stop()

	// --
	_, err = chain.InsertChain([]*types.Block{genesis})
	if !assert.Empty(t, err) {
		t.FailNow()
	}

	parent := chain.GetHeaderByHash(genesis.Hash())
	if !assert.NotEmpty(t, parent) {
		t.FailNow()
	}

	iterCount := 150

	engine.diffFn = func(ChainReader, uint64, *types.Header,
		*timeTarget) *big.Int {
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
			t.FailNow()
		}
		log.Debug("preparing engine")
		err = engine.Prepare(chain, header)
		if !assert.Empty(t, err) {
			t.FailNow()
		}
		if !assert.NotEmpty(t, header.Difficulty) {
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
			receipt, _, err := core.ApplyTransaction(
				&chainConfig, chain, &header.Coinbase,
				new(core.GasPool).AddGas(header.GasLimit),
				blstate, header, tx,
				&header.GasUsed, *chain.GetVMConfig())
			if !assert.Empty(t, err) {
				t.FailNow()
			}
			txs = append(txs, tx)
			receipts = append(receipts, receipt)
		}
		block, finalizedReceipts, err := engine.Finalize(
			chain, header, blstate, txs, nil, receipts)
		if !assert.Empty(t, err) {
			t.FailNow()
		}

		if i == 1 {
			if !assert.Equal(t, 1, len(finalizedReceipts)) {
				t.FailNow()
			}
		} else {
			if !assert.Empty(t, finalizedReceipts) {
				t.FailNow()
			}
		}

		// ---
		log.Debug("sealing migration block")
		err = engine.Seal(chain, block, results, stop)
		if !assert.Empty(t, err) {
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
			t.FailNow()
		}
		if !assert.NotEmpty(t, blstate) {
			log.Debug("state was empty")
			t.FailNow()
		}
		if !assert.NotEmpty(t, finalizedReceipts) {
			log.Debug("receipts were empty")
			t.FailNow()
		}
		header = block.Header()
		// assert.NotEqual(t, parent.Coinbase, header.Coinbase, "Header %v", i)
		if !assert.NotEqual(t, parent.Coinbase, common.Address{},
			"Header %v", i) {
			t.FailNow()
		}
		err = engine.VerifySeal(chain, header)
		if !assert.Empty(t, err) {
			t.FailNow()
		}

		// Test consensus tx check during block processing
		// ---
		if i == 2 {
			tmptxs := block.Transactions()
			tmpheader := *header

			if !assert.Equal(t, len(tmptxs), 1) {
				t.FailNow()
			}

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(), tmptxs, nil, finalizedReceipts)
			if !assert.Empty(t, err) {
				t.FailNow()
			}

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(), append(tmptxs, tmptxs[len(tmptxs)-1]), nil, finalizedReceipts)
			if !assert.Equal(t, eth_consensus.ErrInvalidConsensusTx,
				err) {
				t.FailNow()
			}

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(),
				append(tmptxs[:len(tmptxs)-1], tmptxs[len(tmptxs)-1].WithConsensusSender(common.Address{})),
				nil, finalizedReceipts)
			if !assert.Equal(t, eth_consensus.ErrInvalidConsensusTx,
				err) {
				t.FailNow()
			}
		}

		// Time tests
		// ---
		tt := engine.calcTimeTargetV2(chain, parent)
		if !assert.True(t, tt.maxTime >= now) {
			t.FailNow()
		}
		if !assert.True(t, tt.maxTime <= engine.now()+30) {
			t.FailNow()
		}

		if i < 60 {
			// parent header and current header must be minTime apart(30s)
			if !assert.Equal(t, header.Time,
				parent.Time+30) {
				t.FailNow()
			}
			if !assert.Equal(t, tt.minTime,
				header.Time) {
				t.FailNow()
			}
			if !assert.Equal(t, tt.target,
				header.Time+30) {
				t.FailNow()
			}
		} else if i < 61 {
			if !assert.Equal(t, header.Time, genesis.Time()+1800) {
				t.FailNow()
			}
			if !assert.Equal(t, header.Time, parent.Time+30) {
				t.FailNow()
			}
			if !assert.Equal(t, tt.minTime, header.Time) {
				t.FailNow()
			}
			if !assert.Equal(t, tt.target,
				parent.Time+60) {
				t.FailNow()
			}
		} else if i < 62 {
			log.Debug("time test", "header.Time", header.Time,
				"genesis.Time()", genesis.Time(),
				"difference", header.Time-genesis.Time(),
			)
			if !assert.Equal(t, header.Time,
				genesis.Time()+1830) {
				t.FailNow()
			}
		}

		if !assert.True(t, parent.Time < tt.minTime, "Header %v",
			i) {
			t.FailNow()
		}

		//		assert.Empty(t, engine.enforceTime(header, tt))
		//		assert.Empty(t, engine.checkTime(header, tt))

		_, err = chain.WriteBlockWithState(block, finalizedReceipts, blstate)
		if !assert.Empty(t, err) {
			t.FailNow()
		}

		parent = header
	}
}

/*
 * Run multiple test cases
 * Call CalcPoSDifficultyV2, analyzing the result
 * Assertions:
 * - Difficulty is correct
 */
func TestPoSDiffV2(t *testing.T) {
	// t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)
	flag.Parse()
	log.Root().SetHandler(
		log.LvlFilterHandler(
			5,
			log.StreamHandler(
				os.Stderr,
				log.TerminalFormat(false),
			),
		),
	)
	// this enables code location printing
	log.PrintOrigins(true)

	type TC struct {
		parent int64
		time   uint64
		min    uint64
		target uint64
		result uint64
	}

	tests := []TC{
		{
			parent: 100,
			time:   61,
			min:    31,
			target: 61,
			result: 100,
		},
		{
			parent: 100,
			time:   31,
			min:    31,
			target: 61,
			result: 100,
		},
		{
			parent: 100,
			time:   31,
			min:    31,
			target: 51,
			result: 100,
		},
		{
			parent: 100,
			time:   31,
			min:    61,
			target: 31,
			result: 100,
		},
		{
			parent: 100,
			time:   31,
			min:    31,
			target: 31,
			result: 100,
		},
		{
			parent: 1744,
			time:   91,
			min:    31,
			target: 61,
			result: 1738,
		},
		{
			parent: 1744,
			time:   121,
			min:    31,
			target: 61,
			result: 1738,
		},
		{
			parent: 1744,
			time:   200,
			min:    31,
			target: 61,
			result: 1738,
		},
		{
			parent: 1744,
			time:   181,
			min:    31,
			target: 61,
			result: 1738,
		},
	}

	log.Debug("looping over tests")
	for i, tc := range tests {
		parent := &types.Header{
			Difficulty: big.NewInt(tc.parent),
		}
		tt := &timeTargetV2{
			minTime: tc.min,
			target:  tc.target,
		}

		res := calcPoSDifficultyV2(tc.time, parent, tt)
		assert.Equal(t, tc.result, res.Uint64(), "TC %v", i)
	}
}

/*
 * Create 5 addresses
 * For each address call minev2, analyzing the result
 * Assertions:
 * - The function returns a Success
 * - No error is returned
 */
/*
func TestPoSMineV2(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	addresses, signers, alloc, migrationSigner := generateAddresses(5)
	testdb := ethdb.NewMemDatabase()
	var header *types.Header

	engine := New(&params.EnergiConfig{MigrationSigner: migrationSigner}, testdb)
	engine.testing = true
	engine.diffFn = func(ChainReader, uint64, *types.Header, *timeTarget) *big.Int {
		return common.Big1
	}
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
			return crypto.Sign(hash, signers[addr])
		},
		func() int { return 1 },
		func() bool { return true },
	)

	chainConfig := *params.EnergiTestnetChainConfig
	chainConfig.Energi = &params.EnergiConfig{
		MigrationSigner: migrationSigner,
	}
	gspec := &core.Genesis{
		Config:     &chainConfig,
		GasLimit:   8000000,
		Timestamp:  1000,
		Difficulty: big.NewInt(1),
		Coinbase:   energi_params.Energi_Treasury,
		Alloc:      alloc,
		Xfers:      core.DeployEnergiGovernance(&chainConfig),
	}
	genesis := gspec.MustCommit(testdb)

	stateCache := state.NewDatabaseWithCache(testdb, 256)
	stateDB, err := state.New(genesis.Root(), stateCache)
	assert.Empty(t, err)

	fakeChain := new(mockChainReader)
	fakeChain.stateDB = stateDB
	fakeChain.headers = make(map[common.Hash]*types.Header)
	fakeChain.headers[genesis.Hash()] = genesis.Header()

	parent := genesis.Header()
	balance, _ := new(big.Int).SetString("3280000000000000000", 10)

	for i, address := range addresses {
		multiplier := big.NewInt(int64(i))
		multiplier = multiplier.Add(multiplier, big.NewInt(1))
		stateDB.SetBalance(address, balance.Mul(balance, multiplier))
	}

	for i := 1; i < len(addresses)-1; i++ {
		number := new(big.Int).Add(parent.Number, common.Big1)

		header = &types.Header{
			ParentHash: parent.Hash(),
			Coinbase:   addresses[i],
			Difficulty: big.NewInt(1),
			GasLimit:   parent.GasLimit,
			Number:     number,
			Time:       parent.Time,
		}

		fakeChain.current = header
		fakeChain.headers[header.Hash()] = header
		fakeChain.headers[parent.Hash()] = parent

		// Test if accounts function returns no addresses
		engineAccountsFn := engine.accountsFn
		engine.accountsFn = func() []common.Address {
			return []common.Address{}
		}
		stop := make(chan struct{})
		go func() {
			stop <- struct{}{}
		}()
		success, err := engine.MineV2(fakeChain, header, stop)
		assert.Empty(t, err)
		assert.False(t, success)
		close(stop)

		engine.accountsFn = engineAccountsFn

		// Test if chain header returns nil
		parentHeader := fakeChain.headers[header.ParentHash]
		fakeChain.headers[header.ParentHash] = nil
		success, err = engine.MineV2(fakeChain, header, make(chan struct{}))
		assert.Equal(t, err, eth_consensus.ErrUnknownAncestor)
		assert.False(t, success)

		fakeChain.headers[header.ParentHash] = parentHeader

		// Test stop works when PoS miner is sleeping
		engineNow := engine.now
		timeNow := uint64(1000)
		engine.now = func() uint64 { timeNow -= 50; return timeNow }
		stop = make(chan struct{})
		go func() {
			stop <- struct{}{}
		}()
		success, err = engine.MineV2(fakeChain, header, stop)
		assert.Empty(t, err)
		assert.False(t, success)
		close(stop)

		engine.now = engineNow

		// Test missing state
		fakeChain.stateDB = nil
		success, err = engine.MineV2(fakeChain, header, make(chan struct{}))
		assert.Equal(t, err, eth_consensus.ErrMissingState)
		assert.False(t, success)

		fakeChain.stateDB = stateDB

		// Test PoS mining
		success, err = engine.MineV2(fakeChain, header, make(chan struct{}))
		assert.Empty(t, err)
		assert.True(t, success)

		parent = header
	}
}*/
