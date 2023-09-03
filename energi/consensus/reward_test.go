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

	"github.com/energicryptocurrency/go-energi/common"
	"github.com/energicryptocurrency/go-energi/consensus/ethash"
	"github.com/energicryptocurrency/go-energi/core"
	"github.com/energicryptocurrency/go-energi/core/state"
	"github.com/energicryptocurrency/go-energi/core/types"
	"github.com/energicryptocurrency/go-energi/core/vm"
	energi_params "github.com/energicryptocurrency/go-energi/energi/params"
	"github.com/energicryptocurrency/go-energi/ethdb"
	"github.com/energicryptocurrency/go-energi/params"

	"github.com/stretchr/testify/assert"
)

func TestBlockRewards(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)
	var (
		testdb = ethdb.NewMemDatabase()
		gspec  = &core.Genesis{
			Config: params.EnergiTestnetChainConfig,
			Xfers:  core.DeployEnergiGovernance(params.EnergiTestnetChainConfig),
		}
		parent = gspec.MustCommit(testdb)

		engine = New(new(params.EnergiConfig), testdb)
	)

	chain, _ := core.NewBlockChain(testdb, nil, params.EnergiTestnetChainConfig, ethash.NewFaker(), vm.Config{}, nil)
	defer chain.Stop()

	statedb, _ := state.New(parent.Root(), state.NewDatabase(testdb))

	header := &types.Header{
		Root:       statedb.IntermediateRoot(chain.Config().IsEIP158(parent.Number())),
		ParentHash: parent.Hash(),
		Coinbase:   parent.Coinbase(),
		Difficulty: engine.CalcDifficulty(chain, parent.Time()+10, &types.Header{
			Number:     parent.Number(),
			Time:       parent.Time(),
			Difficulty: parent.Difficulty(),
			UncleHash:  parent.UncleHash(),
		}),
		GasLimit: parent.GasLimit(),
		Number:   new(big.Int).Add(parent.Number(), common.Big1),
		Time:     parent.Time(),
	}

	err := engine.processConsensusGasLimits(chain, header, statedb)
	assert.Empty(t, err)

	for i := 0; i < 5; i++ {
		// TODO: check balance changes
		txs, receipts, err := engine.processBlockRewards(chain, header, statedb, nil, nil)
		assert.Equal(t, 1, len(txs))
		assert.Equal(t, 1, len(receipts))

		if err != nil {
			panic(err)
		}
	}

}

func TestStakerReward(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)
	var (
		testdb = ethdb.NewMemDatabase()
		gspec  = &core.Genesis{
			Config: params.EnergiTestnetChainConfig,
			Xfers:  core.DeployEnergiGovernance(params.EnergiTestnetChainConfig),
		}
		parent = gspec.MustCommit(testdb)
		engine = New(new(params.EnergiConfig), testdb)
		// define arbitrary number for header gas used
		HeaderGas = uint64(100)
	)

	chain, _ := core.NewBlockChain(testdb, nil, params.EnergiTestnetChainConfig, ethash.NewFaker(), vm.Config{}, nil)
	defer chain.Stop()

	statedb, _ := state.New(parent.Root(), state.NewDatabase(testdb))

	header := &types.Header{
		Root:       statedb.IntermediateRoot(chain.Config().IsEIP158(parent.Number())),
		ParentHash: parent.Hash(),
		Coinbase:   parent.Coinbase(),
		Difficulty: engine.CalcDifficulty(chain, parent.Time()+10, &types.Header{
			Number:     parent.Number(),
			Time:       parent.Time(),
			Difficulty: parent.Difficulty(),
			UncleHash:  parent.UncleHash(),
		}),
		GasLimit: parent.GasLimit(),
		GasUsed:  HeaderGas,
		Number:   new(big.Int).Add(parent.Number(), common.Big1),
		Time:     parent.Time(),
	}

	// process consensus gas limits
	err := engine.processConsensusGasLimits(chain, header, statedb)
	assert.Empty(t, err)

	// run staker rewarding
	txs, receipts, err := engine.processFeeReward(chain, header, statedb, nil, nil)

	// check it returns the only rewarding transaction
	assert.Equal(t, 2, len(txs))

	// check the rewarded amount is correct
	assert.Equal(t, (HeaderGas*energi_params.StakerReward)/100, txs[0].Value().Uint64())

	// check rewarded address is coinbase
	assert.Equal(t, header.Coinbase, *txs[1].To())

	// check rewarded is system faucet
	assert.Equal(t, energi_params.Energi_SystemFaucet, txs[0].ConsensusSender())
	assert.Equal(t, 2, len(receipts))
}
