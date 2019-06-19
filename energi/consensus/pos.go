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
	"errors"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"golang.org/x/crypto/sha3"
)

const (
	MaturityPeriod    uint64 = 60 * 60
	AverageTimeBlocks uint64 = 60
	TargetBlockGap    uint64 = 60
	MinBlockGap       uint64 = 30
	MaxFutureGap      uint64 = 3
	TargetPeriodGap   uint64 = AverageTimeBlocks * TargetBlockGap

	maturityGuessBlocks uint64 = MaturityPeriod / TargetBlockGap
)

var (
	minStake = big.NewInt(1e18)

	baseDifficulty = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

	errBlockMinTime  = errors.New("Minimal time gap is not obeyed")
	errBlockInFuture = errors.New("Too much in future")
	errMissingParent = errors.New("Missing parent")

	errInvalidPoSHash     = errors.New("Invalid PoS hash")
	errInvalidStakeAmount = errors.New("Invalid Stake amount")
)

type timeTarget struct {
	min_time    uint64
	target_time uint64
	max_time    uint64
}

func (e *Energi) now() uint64 {
	return uint64(time.Now().Unix())
}

/**
 * Implements block time consensus
 *
 * POS-11: Block time restrictions
 * POS-12: Block interval enforcement
 */
func (e *Energi) calcTimeTarget(
	chain ChainReader,
	parent *types.Header,
) *timeTarget {
	ret := &timeTarget{}
	now := e.now()
	block_number := parent.Number.Uint64() + 1

	// POS-11: Block time restrictions
	ret.max_time = now + MaxFutureGap

	// POS-11: Block time restrictions
	ret.min_time = parent.Time + MinBlockGap
	ret.target_time = parent.Time + TargetBlockGap

	// POS-12: Block interval enforcement
	//---
	if block_number >= AverageTimeBlocks {
		past := chain.GetHeaderByNumber(block_number - AverageTimeBlocks)
		actual := parent.Time - past.Time
		expected := TargetPeriodGap - TargetBlockGap

		if expected > actual {
			ret.min_time = past.Time + TargetPeriodGap
		}
	}

	log.Trace("PoS time", "block", block_number,
		"min", ret.min_time, "target", ret.target_time, "max", ret.max_time)
	return ret
}

func (e *Energi) enforceTime(
	header *types.Header,
	time_target *timeTarget,
) error {
	// NOTE: allow Miner to hint already tried period by
	if header.Time < time_target.min_time {
		header.Time = time_target.min_time
	}

	// Check if allowed to mine
	if header.Time > time_target.max_time {
		return errBlockInFuture
	}

	return nil
}

func (e *Energi) checkTime(
	header *types.Header,
	time_target *timeTarget,
) error {
	// NOTE: allow Miner to hint already tried period by
	if header.Time < time_target.min_time {
		return errBlockMinTime
	}

	// Check if allowed to mine
	if header.Time > time_target.max_time {
		return errBlockInFuture
	}

	return nil
}

/**
 * Implements check modifier consensus
 *
 * POS-14: Stake modifier
 */
func (e *Energi) calcPoSModifier(
	chain ChainReader,
	time uint64,
	parent *types.Header,
) (ret common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	// Add coinbase
	hasher.Write(parent.Coinbase.Bytes())

	// Find maturity period border
	maturity_border := time

	if maturity_border < MaturityPeriod {
		// This should happen only in testing
		maturity_border = 0
	} else {
		maturity_border -= MaturityPeriod
	}

	// Find the oldest inside maturity period
	parent_height := parent.Number.Uint64()
	guess := parent_height

	if guess < maturityGuessBlocks {
		guess = 0
	} else {
		guess -= maturityGuessBlocks
	}

	// NOTE: the logic below can go into if-clauses, but we always run both
	//       cases

	// If we hit inside the period
	oldest := chain.GetHeaderByNumber(guess)

	for (oldest.Time > maturity_border) && (guess > 0) {
		guess--
		oldest = chain.GetHeaderByNumber(guess)
	}

	// If we hit outside the period
	for (oldest.Time <= maturity_border) && (oldest.Number.Uint64() < parent_height) {
		guess++
		oldest = chain.GetHeaderByNumber(guess)
	}

	// Hash it
	hasher.Write(oldest.Root.Bytes())

	// Sum together
	ret = common.BytesToHash(hasher.Sum(nil))
	log.Trace("PoS modifier", "block", parent_height+1,
		"modifier", ret, "oldest", oldest.Number.Uint64())
	return ret
}

/**
 * Implements difficulty consensus
 *
 * POS-13: Difficulty algorithm
 */
func (e *Energi) calcPoSDifficulty(
	chain ChainReader,
	time uint64,
	parent *types.Header,
) *big.Int {
	//time_target := e.calcTimeTarget(chain, parent)
	ret := big.NewInt(1)
	log.Trace("PoS difficulty", "block", parent.Number.Uint64()+1, "time", time, "diff", ret)
	return ret
}

/**
 * Implements hash consensus
 *
 * POS-18: PoS hash generation
 */
func (e *Energi) calcPoSHash(
	header *types.Header,
	stake *big.Int,
) *big.Int {
	// new(big.Int).SetBytes(poshash.Bytes())
	return baseDifficulty
}

func (e *Energi) verifyPoSHash(
	chain ChainReader,
	header *types.Header,
) error {
	parent := chain.GetHeaderByHash(header.ParentHash)
	stake, err := e.lookupMinBalance(chain, header.Time-MaturityPeriod, parent, header.Coinbase)
	if err != nil {
		return err
	}
	if stake.Cmp(minStake) < 0 {
		return errInvalidStakeAmount
	}
	poshash := e.calcPoSHash(header, stake)
	target := new(big.Int).Div(baseDifficulty, header.Difficulty)

	if poshash.Cmp(target) > 0 {
		return errInvalidPoSHash
	}

	return nil
}

/**
 * Implements stake amount calculation.
 *
 * POS-3: Stake maturity period
 * POS-4: Stake amount
 *
 * This is a basic helper for stake amount calculation.
 * There are ways to optimize it for high load, but we need something
 * to start with.
 */
func (e *Energi) lookupMinBalance(
	chain ChainReader,
	since uint64,
	till *types.Header,
	addr common.Address,
) (min_balance *big.Int, err error) {
	stdb, err := chain.GetStateDB()

	if err != nil {
		log.Error("PoS stake amount is called without state database", "err", err)
		return nil, err
	}

	// NOTE: we need to ensure at least one iteration with the balance condition
	for (till.Time > since) || (min_balance == nil) {
		/*if (till.Coinbase == addr) {
			// Found block resets maturity period
			min_balance = common.Big0
			break
		}*/

		blockst, err := state.New(till.Root, stdb)

		if err != nil {
			log.Error("PoS state root failure", "err", err)
			return nil, err
		}

		bl_balance := blockst.GetBalance(addr)

		if (min_balance == nil) || (min_balance.Cmp(bl_balance) > 0) {
			min_balance = bl_balance

			// No need to lookup further
			if min_balance.Cmp(minStake) < 0 {
				break
			}
		}

		curr := till
		till = chain.GetHeaderByHash(till.ParentHash)

		if till == nil {
			if curr.Number.Cmp(common.Big0) == 0 {
				break
			}

			log.Error("PoS state missing parent")
			return nil, errMissingParent
		}
	}

	log.Trace("PoS stake amount", "addr", addr, "amount", min_balance)
	return min_balance, nil
}

func (e *Energi) mine(
	chain ChainReader,
	header *types.Header,
	stop <-chan struct{},
) (success bool, err error) {
	type Candidates struct {
		addr  common.Address
		stake *big.Int
	}

	accounts := e.accountsFn()
	candidates := make([]Candidates, 0, len(accounts))
	for _, a := range accounts {
		candidates = append(candidates, Candidates{
			addr:  a,
			stake: common.Big0,
		})
		log.Trace("PoS miner candidate found", "address", a)
	}

	//---
	parent := chain.GetHeaderByHash(header.ParentHash)
	tt := e.calcTimeTarget(chain, parent)
	target := new(big.Int).Div(baseDifficulty, header.Difficulty)

	blockTime := tt.min_time

	//---
	for ; ; blockTime++ {
		header.Time = blockTime
		log.Trace("PoS miner time", "time", blockTime)

		// It could be done once, but then there is a chance to miss blocks.
		// Some significant algo optimizations are possible, but we start with simplicity.
		for i, v := range candidates {
			candidates[i].stake, err = e.lookupMinBalance(
				chain, blockTime-MaturityPeriod, parent, v.addr)
			if err != nil {
				return false, err
			}
		}
		// Try smaller amounts first
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].stake.Cmp(candidates[j].stake) < 0
		})
		// Try to match target
		for _, v := range candidates {
			if v.stake.Cmp(minStake) < 0 {
				log.Trace("PoS miner skipping small amount",
					"stake", v.stake, "minstake", minStake)
				continue
			}

			header.Coinbase = v.addr
			poshash := e.calcPoSHash(header, v.stake)

			if poshash.Cmp(target) <= 0 {
				return true, nil
			}
		}

		if now := e.now(); blockTime > now {
			log.Trace("PoS miner is sleeping")
			select {
			case <-stop:
				return false, nil
			case <-time.After(time.Duration(blockTime-now) * time.Second):
			}
		}
		// else try to find a better block in any case!
	}

	return false, nil
}
