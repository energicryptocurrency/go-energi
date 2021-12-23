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
	"bytes"
	"sort"

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/log"
)

type EngineAPI struct {
	chain  ChainReader
	engine *Energi
}

func NewEngineAPI(chain ChainReader, engine *Energi) *EngineAPI {
	return &EngineAPI{
		chain:  chain,
		engine: engine,
	}
}

type StakingStatusInfo struct {
	Hash        common.Hash
	Height      uint64
	Miner       bool
	NonceCap    uint64
	Staking     bool
	TotalWeight uint64
	Accounts    []StakingAccount
}

type StakingAccount struct {
	Account common.Address
	Weight  uint64
}

func (a *EngineAPI) StakingStatus() *StakingStatusInfo {
	res := &StakingStatusInfo{}

	chain := a.chain
	engine := a.engine

	parent := chain.CurrentHeader()
	res.Hash = parent.Hash()
	res.Height = parent.Number.Uint64()

	res.Miner = engine.isMiningFn()
	res.NonceCap = engine.GetMinerNonceCap()

	raw_accounts := engine.accountsFn()
	sort.Slice(raw_accounts, func(a, b int) bool {
		return bytes.Compare(raw_accounts[a][:], raw_accounts[b][:]) < 0
	})
	res.Accounts = make([]StakingAccount, 0, len(raw_accounts))

	for _, acct := range raw_accounts {
		weight, err := engine.lookupStakeWeight(
			chain,
			engine.now(),
			parent,
			acct,
		)
		if err != nil {
			log.Warn("PoS weight lookup failed", "err", err)
			continue
		}
		res.TotalWeight += weight
		res.Accounts = append(res.Accounts, StakingAccount{
			Account: acct,
			Weight:  weight,
		})
	}

	res.Staking = (res.TotalWeight > 0) && res.Miner

	return res
}

func (a *EngineAPI) SetNonceCap(nonce *uint64) (oldNonce uint64) {
	oldNonce = a.engine.GetMinerNonceCap()
	if nonce == nil {
		return
	}

	a.engine.SetMinerNonceCap(*nonce)
	return
}
