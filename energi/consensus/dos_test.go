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
	"testing"

	"github.com/energicryptocurrency/energi/common"
	eth_consensus "github.com/energicryptocurrency/energi/consensus"
	"github.com/energicryptocurrency/energi/core/state"
	"github.com/energicryptocurrency/energi/core/types"

	// "github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/params"

	"github.com/stretchr/testify/assert"

	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

type fakeDoSChain struct {
	current *types.Header
	parent  *types.Header
}

func (fc *fakeDoSChain) Config() *params.ChainConfig {
	panic("Not impl")
}
func (fc *fakeDoSChain) CurrentHeader() *types.Header {
	return fc.current
}
func (fc *fakeDoSChain) GetHeader(hash common.Hash, number uint64) *types.Header {
	return fc.parent
}
func (fc *fakeDoSChain) GetHeaderByNumber(number uint64) *types.Header {
	panic("Not impl")
}
func (fc *fakeDoSChain) GetHeaderByHash(hash common.Hash) *types.Header {
	panic("Not impl")
}
func (fc *fakeDoSChain) GetBlock(hash common.Hash, number uint64) *types.Block {
	panic("Not impl")
}
func (fc *fakeDoSChain) CalculateBlockState(hash common.Hash, number uint64) *state.StateDB {
	panic("Not impl")
}

func KnownStakesTestCount(ks *KnownStakes) (ret int) {
	ks.Range(func(_, _ interface{}) bool {
		ret++
		return true
	})
	return
}

func TestPoSDoS(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	h := &types.Header{}
	p := &types.Header{}
	c := &types.Header{}
	fc := &fakeDoSChain{
		parent:  p,
		current: c,
	}

	base := uint64(1000000)
	curr_time := base
	engine := New(nil, nil)
	engine.now = func() uint64 { return curr_time }
	
	// POS-8: old fork protection
	//============================

	// log.Trace("Regular grow")
	curr_time = base
	p.Time = base
	c.Time = base
	h.Time = base + energi_params.MinBlockGap
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))

	// log.Trace("Side chain as new fork")

	curr_time = base
	p.Time = base - energi_params.OldForkPeriod
	c.Time = base
	h.Time = base + energi_params.MinBlockGap
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))

	// log.Trace("Side chain as old fork")
	curr_time = base + 1
	p.Time = base - energi_params.OldForkPeriod
	c.Time = base
	h.Time = base + energi_params.MinBlockGap
	h1 := &types.Header{
		Time: base + energi_params.MinBlockGap,
		Nonce: types.BlockNonce{
			12, 15, 35,
		},
	}
	err := engine.checkDoS(fc, h1, p)
	assert.Equal(t, eth_consensus.ErrDoSThrottle, err)

	// these are commented out because POS-8 is disabled
	/*
		// log.Trace("Side chain as old fork")
			curr_time = base
			p.Time = base - energi_params.OldForkPeriod - 1
			c.Time = base
			h.Time = base + energi_params.MinBlockGap
			err = engine.checkDoS(fc, h, p)
			assert.Equal(t, eth_consensus.ErrDoSThrottle, err)

			// log.Trace("Side chain as old fork an near old current")
			curr_time = base + energi_params.OldForkPeriod - 1
			p.Time = base - energi_params.OldForkPeriod - 1
			c.Time = base
			h.Time = base + energi_params.MinBlockGap
			assert.Equal(t, eth_consensus.ErrDoSThrottle, engine.checkDoS(fc, h, p))

			// log.Trace("Side chain as old fork an old current - allow old forks")
			curr_time = base + energi_params.OldForkPeriod
			p.Time = base - energi_params.OldForkPeriod - 1
			c.Time = base
			h.Time = base + energi_params.MinBlockGap
			assert.Equal(t, nil, engine.checkDoS(fc, h, p))

	*/
	// POS-9: stake throttling
	//============================

	// log.Trace("Another variation")
	curr_time += energi_params.StakeThrottle
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))
	p.Time = base
	c.Time = base
	h.Time = base + energi_params.MinBlockGap + 1
	assert.Equal(t, eth_consensus.ErrDoSThrottle, engine.checkDoS(fc, h, p))
	assert.Equal(t, 1, KnownStakesTestCount(&engine.knownStakes))

	// log.Trace("Another coinbase")
	h.Coinbase = common.HexToAddress("0x1234")
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))
	assert.Equal(t, 2, KnownStakesTestCount(&engine.knownStakes))

	// log.Trace("Another variation")
	h.Root = common.HexToHash("0x1234")
	assert.Equal(t, eth_consensus.ErrDoSThrottle, engine.checkDoS(fc, h, p))

	// log.Trace("Should reset")
	curr_time += energi_params.StakeThrottle
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))

	// log.Trace("Check correct cleanup")
	h.Coinbase = common.HexToAddress("0x2345")
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))
	h.Coinbase = common.HexToAddress("0x3456")
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))
	assert.Equal(t, 4, KnownStakesTestCount(&engine.knownStakes))

	curr_time += energi_params.StakeThrottle / 2
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))
	h.Time += 1
	assert.Equal(t, eth_consensus.ErrDoSThrottle, engine.checkDoS(fc, h, p))
	assert.Equal(t, 3, KnownStakesTestCount(&engine.knownStakes))

	curr_time += energi_params.StakeThrottle
	assert.Equal(t, nil, engine.checkDoS(fc, h, p))
	assert.Equal(t, 3, KnownStakesTestCount(&engine.knownStakes))
}
