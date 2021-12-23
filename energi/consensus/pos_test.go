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

package consensus

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/energicryptocurrency/energi/common"
	eth_consensus "github.com/energicryptocurrency/energi/consensus"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/core/state"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/core/vm"
	"github.com/energicryptocurrency/energi/crypto"
	"github.com/energicryptocurrency/energi/ethdb"

	// "github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/params"
	"github.com/stretchr/testify/assert"

	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

type mockChainReader struct {
	current *types.Header
	stateDB *state.StateDB
	headers map[common.Hash]*types.Header
}

func (cr *mockChainReader) Config() *params.ChainConfig {
	panic("Not impl")
}
func (cr *mockChainReader) CurrentHeader() *types.Header {
	return cr.current
}
func (cr *mockChainReader) GetHeader(hash common.Hash, number uint64) *types.Header {
	return cr.headers[hash]
}
func (cr *mockChainReader) GetHeaderByNumber(number uint64) *types.Header {
	panic("Not impl")
}
func (cr *mockChainReader) GetHeaderByHash(hash common.Hash) *types.Header {
	panic("Not impl")
}
func (cr *mockChainReader) GetBlock(hash common.Hash, number uint64) *types.Block {
	panic("Not impl")
}
func (cr *mockChainReader) CalculateBlockState(hash common.Hash, number uint64) *state.StateDB {
	return cr.stateDB
}

func generateAddresses(len int) ([]common.Address, map[common.Address]*ecdsa.PrivateKey, core.GenesisAlloc, common.Address) {
	signers := make(map[common.Address]*ecdsa.PrivateKey, len)
	addresses := make([]common.Address, 0, len)
	alloc := make(core.GenesisAlloc, len)

	for i := 0; i < len; i++ {
		k, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		a := crypto.PubkeyToAddress(k.PublicKey)

		signers[a] = k
		addresses = append(addresses, a)
		alloc[a] = core.GenesisAccount{
			Balance: minStake,
		}
	}

	alloc[energi_params.Energi_MigrationContract] = core.GenesisAccount{
		Balance: minStake,
	}

	migrationSigner := addresses[len-1]
	signers[energi_params.Energi_MigrationContract] = signers[migrationSigner]

	return addresses, signers, alloc, migrationSigner
}

func TestPoSChainV1(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

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
	assert.Empty(t, err)
	defer chain.Stop()

	// --
	_, err = chain.InsertChain([]*types.Block{genesis})
	assert.Empty(t, err)

	parent := chain.GetHeaderByHash(genesis.Hash())
	assert.NotEmpty(t, parent)

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
		blstate := chain.CalculateBlockState(header.ParentHash, parent.Number.Uint64())
		assert.NotEmpty(t, blstate)

		err = engine.Prepare(chain, header)
		assert.Empty(t, err)
		assert.NotEmpty(t, header.Difficulty)
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
				}, engine)
			receipt, _, err := core.ApplyTransaction(
				&chainConfig, chain, &header.Coinbase,
				new(core.GasPool).AddGas(header.GasLimit),
				blstate, header, tx,
				&header.GasUsed, *chain.GetVMConfig())
			assert.Empty(t, err)
			txs = append(txs, tx)
			receipts = append(receipts, receipt)
		}
		block, receipts, err := engine.Finalize(
			chain, header, blstate, txs, nil, receipts)
		assert.Empty(t, err)

		if i == 1 {
			assert.Equal(t, 1, len(receipts))
		} else {
			assert.Empty(t, receipts)
		}

		// ---
		err = engine.Seal(chain, block, results, stop)
		assert.Empty(t, err)

		seal_res := <-results
		block = seal_res.Block
		blstate = seal_res.NewState
		receipts = seal_res.Receipts
		assert.NotEmpty(t, block)
		assert.NotEmpty(t, blstate)
		assert.NotEmpty(t, receipts)
		header = block.Header()
		// assert.NotEqual(t, parent.Coinbase, header.Coinbase, "Header %v", i)
		assert.NotEqual(t, parent.Coinbase, common.Address{}, "Header %v", i)
		err = engine.VerifySeal(chain, header)
		assert.Empty(t, err)

		// Test consensus tx check during block processing
		// ---
		if i == 2 {
			tmptxs := block.Transactions()
			tmpheader := *header

			assert.Equal(t, len(tmptxs), 1)

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(), tmptxs, nil, receipts)
			assert.Empty(t, err)

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(), append(tmptxs, tmptxs[len(tmptxs)-1]), nil, receipts)
			assert.Equal(t, eth_consensus.ErrInvalidConsensusTx, err)

			_, _, err = engine.Finalize(
				chain, &tmpheader, blstate.Copy(),
				append(tmptxs[:len(tmptxs)-1], tmptxs[len(tmptxs)-1].WithConsensusSender(common.Address{})),
				nil, receipts)
			assert.Equal(t, eth_consensus.ErrInvalidConsensusTx, err)
		}

		// Time tests
		// ---
		tt := engine.calcTimeTargetV1(chain, parent)
		assert.True(t, tt.max >= now)
		assert.True(t, tt.max <= engine.now()+30)

		if i < 60 {
			assert.Equal(t, header.Time, parent.Time+30)

			assert.Equal(t, tt.min, header.Time)
			assert.Equal(t, tt.blockTarget, header.Time+30)
		} else if i < 61 {
			assert.Equal(t, header.Time, genesis.Time()+3570)
			assert.Equal(t, header.Time, parent.Time+1800)

			assert.Equal(t, tt.min, header.Time)
			assert.Equal(t, tt.blockTarget, parent.Time+60)
		} else if i < 62 {
			assert.Equal(t, header.Time, genesis.Time()+3600)
		}

		assert.True(t, parent.Time < tt.min, "Header %v", i)

		assert.Empty(t, engine.enforceMinTime(header, tt))
		assert.Empty(t, engine.checkTime(header, tt))

		_, err = chain.WriteBlockWithState(block, receipts, blstate)
		assert.Empty(t, err)

		parent = header
	}
}

func TestPoSDiffV1(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	type TC struct {
		parent  int64
		time    uint64
		min     uint64
		btarget uint64
		ptarget uint64
		result  uint64
	}

	tests := []TC{
		{
			parent:  100,
			time:    61,
			min:     31,
			btarget: 61,
			ptarget: 61,
			result:  100,
		},
		{
			parent:  100,
			time:    31,
			min:     31,
			btarget: 61,
			ptarget: 61,
			result:  1744,
		},
		{
			parent:  100,
			time:    31,
			min:     31,
			btarget: 51,
			ptarget: 71,
			result:  1744,
		},
		{
			parent:  100,
			time:    31,
			min:     61,
			btarget: 31,
			ptarget: 31,
			result:  1744,
		},
		{
			parent:  100,
			time:    31,
			min:     31,
			btarget: 31,
			ptarget: 31,
			result:  100,
		},
		{
			parent:  1744,
			time:    91,
			min:     31,
			btarget: 61,
			ptarget: 61,
			result:  403,
		},
		{
			parent:  1744,
			time:    121,
			min:     31,
			btarget: 61,
			ptarget: 61,
			result:  93,
		},
		{
			parent:  1744,
			time:    200,
			min:     31,
			btarget: 61,
			ptarget: 61,
			result:  4,
		},
		{
			parent:  1744,
			time:    181,
			min:     31,
			btarget: 61,
			ptarget: 61,
			result:  4,
		},
	}

	for i, tc := range tests {
		parent := &types.Header{
			Difficulty: big.NewInt(tc.parent),
		}
		tt := &TimeTarget{
			min:          tc.min,
			blockTarget:  tc.btarget,
			periodTarget: tc.ptarget,
		}

		res := calcPoSDifficultyV1(tc.time, parent, tt)
		assert.Equal(t, tc.result, res.Uint64(), "TC %v", i)
	}
}

func TestStakeWeightLookup(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	addresses, _, _, _ := generateAddresses(5)
	engine := New(nil, nil)
	engine.testing = true

	genesis := types.NewBlock(&types.Header{
		Number:     big.NewInt(0),
		Time:       1000,
		GasLimit:   8000000,
		GasUsed:    0,
		Difficulty: big.NewInt(1),
		Coinbase:   energi_params.Energi_Treasury,
	}, nil, nil, nil)

	stateCache := state.NewDatabaseWithCache(ethdb.NewMemDatabase(), 256)
	stateDB, err := state.New(genesis.Root(), stateCache)
	assert.Empty(t, err)

	fakeChain := new(mockChainReader)
	fakeChain.stateDB = stateDB
	fakeChain.headers = make(map[common.Hash]*types.Header)
	fakeChain.headers[genesis.Hash()] = genesis.Header()

	balance, _ := new(big.Int).SetString("3280000000000000000", 10)
	expectedWeights := []uint64{3, 6, 19, 78, 393}
	parent := genesis.Header()
	for i, address := range addresses {
		number := new(big.Int).Add(parent.Number, common.Big1)
		header := &types.Header{
			ParentHash: parent.Hash(),
			Coinbase:   address,
			GasLimit:   parent.GasLimit,
			Number:     number,
			Time:       parent.Time,
		}
		fakeChain.current = header
		fakeChain.headers[header.Hash()] = header
		fakeChain.headers[parent.Hash()] = parent

		multiplier := big.NewInt(int64(i))
		multiplier = multiplier.Add(multiplier, big.NewInt(1))
		stateDB.SetBalance(header.Coinbase, balance.Mul(balance, multiplier))

		// Test CalculateBlockState returns nil handled correctly
		fakeChain.stateDB = nil
		weight, err := engine.lookupStakeWeight(fakeChain, header.Time, parent, header.Coinbase)
		assert.Equal(t, err, eth_consensus.ErrMissingState)
		assert.Equal(t, weight, uint64(0))

		// Test when Coinbase addresses are the same
		fakeChain.stateDB = stateDB
		parentCoinbase := parent.Coinbase
		parent.Coinbase = header.Coinbase
		weight, err = engine.lookupStakeWeight(fakeChain, header.Time, parent, header.Coinbase)
		assert.Empty(t, err)
		assert.Equal(t, weight, expectedWeights[i])

		// Test total_steaked > weight returns 0
		parentNonce := parent.Nonce
		parent.Nonce = types.BlockNonce{255, 255, 255, 255, 255, 255, 255, 255}
		weight, err = engine.lookupStakeWeight(fakeChain, header.Time, parent, header.Coinbase)
		assert.Empty(t, err)
		assert.Equal(t, weight, uint64(0))

		parent.Coinbase = parentCoinbase
		parent.Nonce = parentNonce

		// Test weights match expected
		weight, err = engine.lookupStakeWeight(fakeChain, header.Time, parent, header.Coinbase)
		assert.Empty(t, err)
		assert.Equal(t, weight, expectedWeights[i])

		parent = header
	}
}

func TestPoSMine(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	addresses, signers, alloc, migrationSigner := generateAddresses(5)
	testdb := ethdb.NewMemDatabase()
	var header *types.Header

	engine := New(&params.EnergiConfig{MigrationSigner: migrationSigner}, testdb)
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
		success, err := engine.mine(fakeChain, header, stop)
		assert.Empty(t, err)
		assert.False(t, success)
		close(stop)

		engine.accountsFn = engineAccountsFn

		// Test if chain header returns nil
		parentHeader := fakeChain.headers[header.ParentHash]
		fakeChain.headers[header.ParentHash] = nil
		success, err = engine.mine(fakeChain, header, make(chan struct{}))
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
		success, err = engine.mine(fakeChain, header, stop)
		assert.Empty(t, err)
		assert.False(t, success)
		close(stop)

		engine.now = engineNow

		// Test missing state
		fakeChain.stateDB = nil
		success, err = engine.mine(fakeChain, header, make(chan struct{}))
		assert.Equal(t, err, eth_consensus.ErrMissingState)
		assert.False(t, success)

		fakeChain.stateDB = stateDB

		// Test PoS mining
		success, err = engine.mine(fakeChain, header, make(chan struct{}))
		assert.Empty(t, err)
		assert.True(t, success)

		parent = header
	}
}

func TestVerifyPoSHash(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	addresses, _, alloc, migrationSigner := generateAddresses(5)
	testdb := ethdb.NewMemDatabase()
	var header *types.Header

	engine := New(nil, nil)
	engine.testing = true

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

	for i := 1; i < len(addresses)-1; i++ {
		number := new(big.Int).Add(parent.Number, common.Big1)

		header = &types.Header{
			ParentHash: parent.Hash(),
			Coinbase:   addresses[i],
			Difficulty: big.NewInt(1),
			GasLimit:   parent.GasLimit,
			Number:     number,
			Time:       parent.Time,
			Nonce:      types.BlockNonce{0, 0, 0, 0, 0, 0, 0, 1},
		}

		fakeChain.current = header
		fakeChain.headers[header.Hash()] = header
		fakeChain.headers[parent.Hash()] = parent

		// Test if chain get header fails
		headerParent := fakeChain.headers[header.ParentHash]
		fakeChain.headers[header.ParentHash] = nil
		err := engine.verifyPoSHash(fakeChain, header)
		assert.Equal(t, err, eth_consensus.ErrUnknownAncestor)

		fakeChain.headers[header.ParentHash] = headerParent

		// Test if state is empty
		fakeChain.stateDB = nil
		err = engine.verifyPoSHash(fakeChain, header)
		assert.Equal(t, err, eth_consensus.ErrMissingState)

		fakeChain.stateDB = stateDB

		// Test invalid PoS hash
		headerDifficulty := header.Difficulty
		header.Difficulty = big.NewInt(1000000)
		err = engine.verifyPoSHash(fakeChain, header)
		assert.Equal(t, err, errInvalidPoSHash)

		header.Difficulty = headerDifficulty

		// Test valid arguments
		err = engine.verifyPoSHash(fakeChain, header)
		assert.Empty(t, err)

		parent = header
	}
}
