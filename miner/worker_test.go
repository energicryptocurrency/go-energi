// Copyright 2018 The Energi Core Authors
// Copyright 2018 The go-ethereum Authors
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

package miner

import (
	"math/big"
	"sync/atomic"
	"testing"
	"time"

	"github.com/energicryptocurrency/energi/accounts"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/consensus"
	"github.com/energicryptocurrency/energi/consensus/clique"
	"github.com/energicryptocurrency/energi/consensus/ethash"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/core/vm"
	"github.com/energicryptocurrency/energi/crypto"
	energi_testutils "github.com/energicryptocurrency/energi/energi/common/testutils"
	energi "github.com/energicryptocurrency/energi/energi/consensus"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/ethdb"
	"github.com/energicryptocurrency/energi/event"
	"github.com/energicryptocurrency/energi/params"
)

var (
	// Test chain configurations
	testTxPoolConfig  core.TxPoolConfig
	ethashChainConfig *params.ChainConfig
	cliqueChainConfig *params.ChainConfig
	energiChainConfig *params.ChainConfig

	// Test accounts
	testBankKey, _  = crypto.GenerateKey()
	testBankAddress = crypto.PubkeyToAddress(testBankKey.PublicKey)
	testBankFunds   = big.NewInt(1000000000000000000)

	testUserKey, _  = crypto.GenerateKey()
	testUserAddress = crypto.PubkeyToAddress(testUserKey.PublicKey)

	// Test transactions
	pendingTxs []*types.Transaction
	newTxs     []*types.Transaction

	// Energi signers private keys
	migrationSigner, _ = crypto.GenerateKey()
)

func init() {
	testTxPoolConfig = core.DefaultTxPoolConfig
	testTxPoolConfig.Journal = ""
	ethashChainConfig = params.TestChainConfig
	cliqueChainConfig = params.TestChainConfig
	cliqueChainConfig.Clique = &params.CliqueConfig{
		Period: 10,
		Epoch:  30000,
	}
	energiChainConfig = params.EnergiTestnetChainConfig
	tx1, _ := types.SignTx(types.NewTransaction(0, testUserAddress, big.NewInt(1000), params.TxGas, nil, nil), types.HomesteadSigner{}, testBankKey)
	pendingTxs = append(pendingTxs, tx1)
	tx2, _ := types.SignTx(types.NewTransaction(1, testUserAddress, big.NewInt(1000), params.TxGas, nil, nil), types.HomesteadSigner{}, testBankKey)
	newTxs = append(newTxs, tx2)
}

// testWorkerBackend implements worker.Backend interfaces and wraps all information needed during the testing.
type testWorkerBackend struct {
	db         ethdb.Database
	txPool     *core.TxPool
	chain      *core.BlockChain
	uncleBlock *types.Block
	migration  *energi_testutils.TestGen2Migration
}

func newTestWorkerBackend(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine, n int) *testWorkerBackend {
	var (
		migrations *energi_testutils.TestGen2Migration

		db    = ethdb.NewMemDatabase()
		gspec = core.Genesis{
			Config: chainConfig,
			Alloc:  core.GenesisAlloc{testBankAddress: {Balance: testBankFunds}},
		}
	)

	switch engine := engine.(type) {
	case *clique.Clique:
		gspec.ExtraData = make([]byte, 32+common.AddressLength+65)
		copy(gspec.ExtraData[32:], testBankAddress[:])
	case *ethash.Ethash:
	case *energi.Energi:
		// Create a gen2 migration tempfile
		migrations = energi_testutils.NewTestGen2Migration()
		if err := migrations.PrepareTestGen2Migration(chainConfig.ChainID.Uint64()); err != nil {
			t.Fatalf("Creating the Gen2 snapshot failed: %v", err)
		}

		// Set the migrations signer.
		gspec.Alloc[energi_params.Energi_MigrationContract] = core.GenesisAccount{Balance: testBankFunds}
		engine.SetMinerCB(
			func() []common.Address {
				return []common.Address{energi_params.Energi_MigrationContract}
			},
			func(addr common.Address, hash []byte) ([]byte, error) {
				return crypto.Sign(hash, migrationSigner)
			},
			func() int { return 1 },
			func() bool { return true },
		)
		chainConfig.Energi = &params.EnergiConfig{
			MigrationSigner: crypto.PubkeyToAddress(migrationSigner.PublicKey),
		}

		// Update genesis block config.
		gspec.GasLimit = 8000000
		gspec.Timestamp = 1000
		gspec.Difficulty = big.NewInt(1)
		gspec.Coinbase = energi_params.Energi_Treasury
		gspec.Xfers = core.DeployEnergiGovernance(chainConfig)

	default:
		t.Fatalf("unexpected consensus engine type: %T", engine)
	}
	genesis := gspec.MustCommit(db)

	chain, _ := core.NewBlockChain(db, nil, gspec.Config, engine, vm.Config{}, nil)
	txpool := core.NewTxPool(testTxPoolConfig, chainConfig, chain)

	// Generate a small n-block chain and an uncle block for it
	if n > 0 {
		blocks, _ := core.GenerateChain(chainConfig, genesis, engine, db, n, func(i int, gen *core.BlockGen) {
			gen.SetCoinbase(testBankAddress)
		})
		if _, err := chain.InsertChain(blocks); err != nil {
			t.Fatalf("failed to insert origin chain: %v", err)
		}
	}
	parent := genesis
	if n > 0 {
		parent = chain.GetBlockByHash(chain.CurrentBlock().ParentHash())
	}
	blocks, _ := core.GenerateChain(chainConfig, parent, engine, db, 1, func(i int, gen *core.BlockGen) {
		gen.SetCoinbase(testUserAddress)
	})

	return &testWorkerBackend{
		db:         db,
		chain:      chain,
		txPool:     txpool,
		uncleBlock: blocks[0],
		migration:  migrations,
	}
}

func (b *testWorkerBackend) BlockChain() *core.BlockChain { return b.chain }
func (b *testWorkerBackend) TxPool() *core.TxPool         { return b.txPool }
func (b *testWorkerBackend) PostChainEvents(events []interface{}) {
	b.chain.PostChainEvents(events, nil)
}
func (b *testWorkerBackend) AccountManager() *accounts.Manager {
	return nil
}
func (b *testWorkerBackend) CleanUp() error { return b.migration.CleanUp() }

func newTestWorker(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine, blocks int) (*worker, *testWorkerBackend) {
	backend := newTestWorkerBackend(t, chainConfig, engine, blocks)
	backend.txPool.AddLocals(pendingTxs)
	w := newWorker(chainConfig, engine, backend, new(event.TypeMux), time.Second, params.GenesisGasLimit, params.GenesisGasLimit, nil)
	w.setEtherbase(testBankAddress)
	w.setMigration(backend.migration.TempFileName())
	return w, backend
}

// func TestPendingStateAndBlockEthash(t *testing.T) {
// 	testPendingStateAndBlock(t, ethashChainConfig, ethash.NewFaker())
// }
// func TestPendingStateAndBlockClique(t *testing.T) {
// 	testPendingStateAndBlock(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, ethdb.NewMemDatabase()))
// }

func TestPendingStateAndBlockEnergi(t *testing.T) {
	testPendingStateAndBlock(t, energiChainConfig, energi.New(energiChainConfig.Energi, ethdb.NewMemDatabase()))
}

func testPendingStateAndBlock(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, 0)
	defer func() {
		w.close()
		_ = b.CleanUp()
	}()

	// Trigger processing of the migration tx at block number 1
	atomic.StoreInt32(&w.running, 1)

	// Ensure snapshot has been updated.
	time.Sleep(1000 * time.Millisecond)
	block, state := w.pending()

	var wantBlockheight uint64
	switch engine.(type) {
	case *energi.Energi:
		// Block count increase because of the migration block that is mined too.
		wantBlockheight = 2
	default:
		wantBlockheight = 1
	}

	if block.NumberU64() != wantBlockheight {
		t.Errorf("block number mismatch: have %d, want %d", block.NumberU64(), wantBlockheight)
	}
	if balance := state.GetBalance(testUserAddress); balance.Cmp(big.NewInt(1000)) != 0 {
		t.Errorf("account balance mismatch: have %d, want %d", balance, 1000)
	}
	b.txPool.AddLocals(newTxs)

	// Ensure the new tx events has been processed
	time.Sleep(1000 * time.Millisecond)
	_, state = w.pending()
	if balance := state.GetBalance(testUserAddress); balance.Cmp(big.NewInt(2000)) != 0 {
		t.Errorf("account balance mismatch: have %d, want %d", balance, 2000)
	}
}

// func TestEmptyWorkEthash(t *testing.T) {
// 	testEmptyWork(t, ethashChainConfig, ethash.NewFaker())
// }
// func TestEmptyWorkClique(t *testing.T) {
// 	testEmptyWork(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, ethdb.NewMemDatabase()))
// }
func TestEmptyWorkEnergi(t *testing.T) {
	testEmptyWork(t, energiChainConfig, energi.New(energiChainConfig.Energi, ethdb.NewMemDatabase()))
}

func testEmptyWork(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, 0)
	defer func() {
		w.close()
		_ = b.CleanUp()
	}()

	var (
		taskCh    = make(chan struct{}, 2)
		taskIndex int
	)

	checkEqual := func(t *testing.T, task *task, index int) {
		receiptLen, balance := 0, big.NewInt(0)
		if index == 1 {
			receiptLen, balance = 1, big.NewInt(1000)
		}
		if len(task.receipts) != receiptLen {
			t.Errorf("receipt number mismatch: have %d, want %d", len(task.receipts), receiptLen)
		}
		if task.state.GetBalance(testUserAddress).Cmp(balance) != 0 {
			t.Errorf("account balance mismatch: have %d, want %d", task.state.GetBalance(testUserAddress), balance)
		}
	}

	w.newTaskHook = func(task *task) {
		if task.block.NumberU64() == 2 {
			checkEqual(t, task, taskIndex)
			taskIndex++
			taskCh <- struct{}{}
		}
	}
	w.fullTaskHook = func() {
		time.Sleep(1000 * time.Millisecond)
	}

	// Trigger processing of the migration tx at block number 1
	atomic.StoreInt32(&w.running, 1)

	// Ensure worker has finished initialization
	for {
		b := w.pendingBlock()
		if b != nil && b.NumberU64() >= 1 {
			break
		}
	}

	w.start()
	for i := 0; i < 2; i++ {
		select {
		case <-taskCh:
		case <-time.NewTimer(2 * time.Second).C:
			t.Error("new task timeout")
		}
	}
}

func TestStreamUncleBlock(t *testing.T) {
	t.Skip("Uncle block implementation is not supported on Energi")

	ethash := ethash.NewFaker()
	defer ethash.Close()

	w, b := newTestWorker(t, ethashChainConfig, ethash, 1)
	defer func() {
		w.close()
		_ = b.CleanUp()
	}()

	var taskCh = make(chan struct{})

	taskIndex := 0
	w.newTaskHook = func(task *task) {
		if task.block.NumberU64() == 2 {
			if taskIndex == 2 {
				have := task.block.Header().UncleHash
				want := types.CalcUncleHash([]*types.Header{b.uncleBlock.Header()})
				if have != want {
					t.Errorf("uncle hash mismatch: have %s, want %s", have.Hex(), want.Hex())
				}
			}
			taskCh <- struct{}{}
			taskIndex++
		}
	}
	w.skipSealHook = func(task *task) bool {
		return true
	}
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}

	// Ensure worker has finished initialization
	for {
		b := w.pendingBlock()
		if b != nil && b.NumberU64() == 2 {
			break
		}
	}
	w.start()

	// Ignore the first two works
	for i := 0; i < 2; i++ {
		select {
		case <-taskCh:
		case <-time.NewTimer(time.Second).C:
			t.Error("new task timeout")
		}
	}
	b.PostChainEvents([]interface{}{core.ChainSideEvent{Block: b.uncleBlock}})

	select {
	case <-taskCh:
	case <-time.NewTimer(time.Second).C:
		t.Error("new task timeout")
	}
}

// func TestRegenerateMiningBlockEthash(t *testing.T) {
// 	testRegenerateMiningBlock(t, ethashChainConfig, ethash.NewFaker())
// }

// func TestRegenerateMiningBlockClique(t *testing.T) {
// 	testRegenerateMiningBlock(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, ethdb.NewMemDatabase()))
// }

func TestRegenerateMiningBlockEnergi(t *testing.T) {
	testRegenerateMiningBlock(t, energiChainConfig, energi.New(energiChainConfig.Energi, ethdb.NewMemDatabase()))
}

func testRegenerateMiningBlock(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, 0)
	defer func() {
		w.close()
		_ = b.CleanUp()
	}()

	var taskCh = make(chan struct{})

	taskIndex := 0
	w.newTaskHook = func(task *task) {
		if task.block.NumberU64() == 2 {
			if taskIndex == 2 {
				receiptLen, balance := 2, big.NewInt(2000)
				if len(task.receipts) != receiptLen {
					t.Errorf("receipt number mismatch: have %d, want %d", len(task.receipts), receiptLen)
				}
				if task.state.GetBalance(testUserAddress).Cmp(balance) != 0 {
					t.Errorf("account balance mismatch: have %d, want %d", task.state.GetBalance(testUserAddress), balance)
				}
			}
			taskCh <- struct{}{}
			taskIndex++
		}
	}
	// w.skipSealHook = func(task *task) bool {
	// 	return true
	// }
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}

	// Trigger processing of the migration tx at block number 1
	atomic.StoreInt32(&w.running, 1)

	// Ensure worker has finished initialization
	for {
		b := w.pendingBlock()
		if b != nil && b.NumberU64() == 1 {
			break
		}
	}

	w.start()
	// Ignore the first two works
	for i := 0; i < 2; i++ {
		select {
		case <-taskCh:
		case <-time.NewTimer(2 * time.Second).C:
			t.Error("new task timeout ..")
		}
	}
	b.txPool.AddLocals(newTxs)
	time.Sleep(time.Second)

	select {
	case <-taskCh:
	case <-time.NewTimer(2 * time.Second).C:
		t.Error("new task timeout")
	}
}

// func TestAdjustIntervalEthash(t *testing.T) {
// 	testAdjustInterval(t, ethashChainConfig, ethash.NewFaker())
// }

// func TestAdjustIntervalClique(t *testing.T) {
// 	testAdjustInterval(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, ethdb.NewMemDatabase()))
// }

func TestAdjustIntervalEnergi(t *testing.T) {
	testAdjustInterval(t, energiChainConfig, energi.New(energiChainConfig.Energi, ethdb.NewMemDatabase()))
}

func testAdjustInterval(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, 0)
	defer func() {
		w.close()
		_ = b.CleanUp()
	}()

	w.skipSealHook = func(task *task) bool {
		return true
	}
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}
	var (
		progress = make(chan struct{}, 10)
		result   = make([]float64, 0, 10)
		index    = 0
		start    = false
	)
	w.resubmitHook = func(minInterval time.Duration, recommitInterval time.Duration) {
		// Short circuit if interval checking hasn't started.
		if !start {
			return
		}
		var wantMinInterval, wantRecommitInterval time.Duration

		switch index {
		case 0:
			wantMinInterval, wantRecommitInterval = 3*time.Second, 3*time.Second
		case 1:
			origin := float64(3 * time.Second.Nanoseconds())
			estimate := origin*(1-intervalAdjustRatio) + intervalAdjustRatio*(origin/0.8+intervalAdjustBias)
			wantMinInterval, wantRecommitInterval = 3*time.Second, time.Duration(estimate)*time.Nanosecond
		case 2:
			estimate := result[index-1]
			min := float64(3 * time.Second.Nanoseconds())
			estimate = estimate*(1-intervalAdjustRatio) + intervalAdjustRatio*(min-intervalAdjustBias)
			wantMinInterval, wantRecommitInterval = 3*time.Second, time.Duration(estimate)*time.Nanosecond
		case 3:
			wantMinInterval, wantRecommitInterval = time.Second, time.Second
		}

		// Check interval
		if minInterval != wantMinInterval {
			t.Errorf("resubmit min interval mismatch: have %v, want %v ", minInterval, wantMinInterval)
		}
		if recommitInterval != wantRecommitInterval {
			t.Errorf("resubmit interval mismatch: have %v, want %v", recommitInterval, wantRecommitInterval)
		}
		result = append(result, float64(recommitInterval.Nanoseconds()))
		index++
		progress <- struct{}{}
	}

	// Trigger processing of the migration tx at block number 1
	atomic.StoreInt32(&w.running, 1)

	// Ensure worker has finished initialization
	for {
		b := w.pendingBlock()
		if b != nil && b.NumberU64() == 1 {
			break
		}
	}

	w.start()

	time.Sleep(time.Second)

	start = true
	w.setRecommitInterval(3 * time.Second)
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.resubmitAdjustCh <- &intervalAdjust{inc: true, ratio: 0.8}
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.resubmitAdjustCh <- &intervalAdjust{inc: false}
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.setRecommitInterval(500 * time.Millisecond)
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}
}
