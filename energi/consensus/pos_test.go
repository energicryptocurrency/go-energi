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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/assert"
)

func TestPoSChain(t *testing.T) {
	t.Parallel()
	log.Root().SetHandler(log.StdoutHandler)
	log.Trace("prevent unused")

	var (
		testdb = ethdb.NewMemDatabase()
		gspec  = &core.Genesis{
			Config: params.TestChainConfig,
		}
		genesis = gspec.MustCommit(testdb)

		engine = New(new(params.EnergiConfig), testdb)
		now    = engine.now()
	)

	chain, err := core.NewBlockChain(testdb, nil, params.TestChainConfig, engine, vm.Config{}, nil)
	assert.Empty(t, err)
	defer chain.Stop()

	_, err = chain.InsertChain([]*types.Block{genesis})
	assert.Empty(t, err)

	parent := chain.GetHeaderByHash(genesis.Hash())
	assert.NotEmpty(t, parent)

	for i := 1; i < 1000; i++ {
		header := &types.Header{
			Root:       parent.Hash(),
			ParentHash: parent.Hash(),
			Coinbase:   parent.Coinbase,
			GasLimit:   parent.GasLimit,
			Number:     new(big.Int).Add(parent.Number, common.Big1),
			Time:       parent.Time,
		}
		err = engine.Prepare(chain, header)
		assert.Empty(t, err)

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

		if i > 60 {
			assert.NotEqual(t, header.MixDigest.Hex(), parent.MixDigest.Hex(), "Header %v", i)
		}

		_, err = chain.InsertHeaderChain([]*types.Header{header}, 1)
		assert.Empty(t, err)

		parent = header
	}
}
