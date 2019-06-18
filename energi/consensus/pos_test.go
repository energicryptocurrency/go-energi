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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/assert"
)

func TestPoSChain(t *testing.T) {
	t.Parallel()
	log.Root().SetHandler(log.StdoutHandler)
	log.Trace("prevent unused")

	tmpKey, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	tmpAddress := crypto.PubkeyToAddress(tmpKey.PublicKey)

	results := make(chan *types.Block, 1)
	stop := make(chan struct{})

	var (
		testdb = ethdb.NewMemDatabase()
		engine = New(new(params.EnergiConfig), testdb)
	)

	engine.SetSigner(func(addr common.Address, hash []byte) ([]byte, error) {
		assert.Equal(t, tmpAddress, addr)
		return crypto.Sign(hash, tmpKey)
	})

	var (
		gspec = &core.Genesis{
			Config:    params.TestChainConfig,
			Timestamp: 1000,
			Coinbase:  tmpAddress,
		}
		genesis = gspec.MustSignCommit(testdb, func(b *types.Block) (*types.Block, error) {
			err := engine.Seal(nil, b, results, stop)
			assert.Empty(t, err)
			b = <-results
			err = engine.VerifySeal(nil, b.Header())
			assert.Empty(t, err)
			return b, err
		})

		now = engine.now()
	)

	chain, err := core.NewBlockChain(testdb, nil, params.TestChainConfig, engine, vm.Config{}, nil)
	assert.Empty(t, err)
	defer chain.Stop()

	//--
	_, err = chain.InsertChain([]*types.Block{genesis})
	assert.Empty(t, err)

	parent := chain.GetHeaderByHash(genesis.Hash())
	assert.NotEmpty(t, parent)
	err = engine.VerifySeal(nil, parent)
	assert.Empty(t, err)

	iterCount := 500
	iterMid := iterCount * 2 / 3

	for i := 1; i < iterCount; i++ {
		number := new(big.Int).Add(parent.Number, common.Big1)
		stdb, err := chain.GetStateDB()
		blstate, err := state.New(parent.Root, stdb)
		assert.Empty(t, err)

		if i <= iterMid {
			blstate.AddBalance(parent.Coinbase, minStake)
		} else {
			blstate.SubBalance(parent.Coinbase, minStake)
		}

		//---
		header := &types.Header{
			ParentHash: parent.Hash(),
			Coinbase:   parent.Coinbase,
			GasLimit:   parent.GasLimit,
			Number:     number,
			Time:       parent.Time,
		}
		err = engine.Prepare(chain, header)
		assert.Empty(t, err)
		assert.NotEmpty(t, header.Difficulty)
		block, err := engine.Finalize(
			chain, header, blstate, []*types.Transaction{}, nil, []*types.Receipt{})

		blstate.Commit(true)
		stdb.TrieDB().Commit(block.Root(), true)

		//---
		err = engine.Seal(chain, block, results, stop)
		assert.Empty(t, err)

		block = <-results
		header = block.Header()
		err = engine.VerifySeal(nil, header)
		assert.Empty(t, err)

		// Time tests
		//---
		tt := engine.calcTimeTarget(chain, parent)
		assert.True(t, tt.max_time >= now)
		assert.True(t, tt.max_time <= engine.now()+30)

		if i < 60 {
			assert.Equal(t, header.Time, parent.Time+30)

			assert.Equal(t, tt.min_time, header.Time)
			assert.Equal(t, tt.target_time, header.Time+30)
		} else if i < 61 {
			assert.Equal(t, header.Time, genesis.Time()+3600)
			assert.Equal(t, header.Time, parent.Time+1830)

			assert.Equal(t, tt.min_time, header.Time)
			assert.Equal(t, tt.target_time, parent.Time+60)
		} else if i < 62 {
			assert.Equal(t, header.Time, genesis.Time()+3630)
		}

		assert.True(t, parent.Time < tt.min_time, "Header %v", i)

		// Stake modifier tests
		//---
		if i > 59 {
			assert.NotEqual(t, header.MixDigest.Hex(), parent.MixDigest.Hex(), "Header %v", i)
		} else if i > 1 {
			assert.Equal(t, header.MixDigest.Hex(), parent.MixDigest.Hex(), "Header %v", i)
		}
		//---

		_, err = chain.InsertHeaderChain([]*types.Header{header}, 1)
		assert.Empty(t, err)

		// Stake amount tests
		//---

		expminbal := new(big.Int).Mul(header.Number, minStake)

		if i > iterMid {
			expminbal = new(big.Int).Mul(big.NewInt(int64(iterMid-(i-iterMid))), minStake)
		}

		minbal, err := engine.lookupMinBalance(chain, header.Time+1, header, header.Coinbase)
		assert.Empty(t, err)
		assert.Equal(t, expminbal.String(), minbal.String())

		minbal, err = engine.lookupMinBalance(chain, header.Time, header, header.Coinbase)
		assert.Empty(t, err)
		assert.Equal(t, expminbal.String(), minbal.String())

		minbal, err = engine.lookupMinBalance(chain, parent.Time, header, header.Coinbase)
		assert.Empty(t, err)
		assert.Equal(t, expminbal.String(), minbal.String())

		minbal, err = engine.lookupMinBalance(chain, parent.Time-1, header, header.Coinbase)
		assert.Empty(t, err)
		if i <= iterMid {
			assert.Equal(t, new(big.Int).Sub(expminbal, minStake).String(), minbal.String())

		} else {
			assert.Equal(t, expminbal.String(), minbal.String())
		}

		if i == iterCount-20 {
			minbal, err = engine.lookupMinBalance(chain, 0, header, header.Coinbase)
			assert.Empty(t, err)
			assert.Equal(t, common.Big0.String(), minbal.String())

			minbal, err = engine.lookupMinBalance(chain, header.Time-3600, header, header.Coinbase)
			assert.Empty(t, err)
			assert.Equal(t, expminbal.String(), minbal.String())
		}
		//---

		parent = header
	}
}
