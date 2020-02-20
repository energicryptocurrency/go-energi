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
	"encoding/binary"
	"errors"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	eth_consensus "github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	energi_params "energi.world/core/gen3/energi/params"
)

const (
	MaturityPeriod    uint64 = energi_params.MaturityPeriod
	AverageTimeBlocks uint64 = energi_params.AverageTimeBlocks
	TargetBlockGap    uint64 = energi_params.TargetBlockGap
	MinBlockGap       uint64 = energi_params.MinBlockGap
	MaxFutureGap      uint64 = energi_params.MaxFutureGap
	TargetPeriodGap   uint64 = energi_params.TargetPeriodGap
)

var (
	minStake = big.NewInt(1e18)

	diff1Target = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

	errBlockMinTime = errors.New("Minimal time gap is not obeyed")

	errInvalidPoSHash  = errors.New("Invalid PoS hash")
	errInvalidPoSNonce = errors.New("Invalid Stake weight")
)

type timeTarget struct {
	min_time      uint64
	max_time      uint64
	block_target  uint64
	period_target uint64
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
	parent_number := parent.Number.Uint64()
	block_number := parent_number + 1

	// POS-11: Block time restrictions
	ret.max_time = now + MaxFutureGap

	// POS-11: Block time restrictions
	ret.min_time = parent.Time + MinBlockGap
	ret.block_target = parent.Time + TargetBlockGap
	ret.period_target = ret.block_target

	// POS-12: Block interval enforcement
	//---
	if block_number >= AverageTimeBlocks {
		// TODO: LRU cache here for extra DoS mitigation
		past := parent

		// NOTE: we have to do this way as parent may be not part of canonical
		//       chain. As no mutex is held, we cannot do checks for canonical.
		for i := AverageTimeBlocks - 1; i > 0; i-- {
			past = chain.GetHeader(past.ParentHash, past.Number.Uint64()-1)

			if past == nil {
				log.Trace("Inconsistent tree, shutdown?")
				return ret
			}
		}

		ret.period_target = past.Time + TargetPeriodGap
		period_min_time := ret.period_target - MinBlockGap

		if period_min_time > ret.min_time {
			ret.min_time = period_min_time
		}
	}

	log.Trace("PoS time", "block", block_number,
		"min", ret.min_time, "max", ret.max_time,
		"block_target", ret.block_target,
		"period_target", ret.period_target,
	)
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

	return nil
}

func (e *Energi) checkTime(
	header *types.Header,
	time_target *timeTarget,
) error {
	if header.Time < time_target.min_time {
		return errBlockMinTime
	}

	// Check if allowed to mine
	if header.Time > time_target.max_time {
		return eth_consensus.ErrFutureBlock
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
	// TODO: LRU cache here for extra DoS mitigation

	// Find maturity period border
	maturity_border := time

	if maturity_border < MaturityPeriod {
		// This should happen only in testing
		maturity_border = 0
	} else {
		maturity_border -= MaturityPeriod
	}

	// Find the oldest inside maturity period
	// NOTE: we have to do this walk as parent may not be part of the canonical chain
	parent_height := parent.Number.Uint64()
	oldest := parent

	for header, num := oldest, oldest.Number.Uint64(); (header.Time > maturity_border) && (num > 0); {

		oldest = header
		num--
		header = chain.GetHeader(header.ParentHash, num)
	}

	// Create Stake Modifier
	ret = crypto.Keccak256Hash(
		parent.Coinbase.Bytes(),
		oldest.Root.Bytes(),
	)

	//
	log.Trace("PoS modifier", "block", parent_height+1,
		"modifier", ret, "oldest", oldest.Number.Uint64())
	return ret
}

/**
 * Implements difficulty consensus
 */
func (e *Energi) calcPoSDifficulty(
	chain ChainReader,
	time uint64,
	parent *types.Header,
	tt *timeTarget,
) (ret *big.Int) {
	ret = e.diffFn(chain, time, parent, tt)
	log.Trace("PoS difficulty", "block", parent.Number.Uint64()+1, "time", time, "diff", ret)
	return ret
}

/**
 * POS-13: Difficulty algorithm (Proposal v1)
 */
const (
	diffV1_BMax     uint64 = 30
	diffV1_AMax     uint64 = 120
	diffV1_DivPlain uint64 = 100

	// Roughly get 2x difficulty decrease
	diffV1_MigrationStakerDelay  uint64 = 15
	diffV1_MigrationStakerTarget uint64 = 0xFFFF
)

var (
	diffV1_BTable []*big.Int
	diffV1_ATable []*big.Int
	diffV1_Div    = new(big.Int).SetUint64(diffV1_DivPlain)
)

func initDiffTable(l uint64, c float64) []*big.Int {
	t := make([]*big.Int, l+1)
	t[0] = common.Big1
	var acc float64 = 1
	for i := 1; i < len(t); i++ {
		acc *= c
		t[i] = big.NewInt(int64(acc * float64(diffV1_DivPlain)))
	}
	return t
}
func init() {
	diffV1_BTable = initDiffTable(diffV1_BMax, 1.1)
	diffV1_ATable = initDiffTable(diffV1_AMax, 1.05)
}
func calcPoSDifficultyV1(
	chain ChainReader,
	time uint64,
	parent *types.Header,
	tt *timeTarget,
) (D *big.Int) {
	// Find out our target anchor
	target := (tt.block_target + tt.period_target) / 2
	if target < tt.min_time {
		target = tt.min_time
	}

	if time < target {
		S := target - time
		if S > diffV1_BMax {
			S = diffV1_BMax
		}
		B := diffV1_BTable[S]
		D = new(big.Int).Div(
			new(big.Int).Mul(parent.Difficulty, B),
			diffV1_Div,
		)
		log.Trace("Diff multiplier", "before", S, "mult", B)
	} else if time > target {
		S := time - target
		if S > diffV1_AMax {
			S = diffV1_AMax
		}
		A := diffV1_ATable[S]
		D = new(big.Int).Div(
			new(big.Int).Mul(parent.Difficulty, diffV1_Div),
			A,
		)
		log.Trace("Diff multiplier", "after", S, "div", A)
	} else {
		log.Trace("No difficulty change", "parent", parent.Difficulty)
		return parent.Difficulty
	}

	if D.Cmp(common.Big1) < 0 {
		D = common.Big1
	}

	log.Trace("Difficulty change",
		"parent", parent.Difficulty, "new", D,
		"time", time, "target", target)
	return D
}

/**
 * Implements hash consensus
 *
 * POS-18: PoS hash generation
 * POS-22: Partial stake amount
 */
func (e *Energi) calcPoSHash(
	header *types.Header,
	target *big.Int,
	weight uint64,
) (poshash *big.Int, used_weight uint64) {
	betime64 := [8]byte{}
	binary.BigEndian.PutUint64(betime64[:], header.Time)

	poshash = new(big.Int).SetBytes(crypto.Keccak256(
		betime64[:],
		header.MixDigest.Bytes(),
		header.Coinbase.Bytes(),
	))

	if poshash.Cmp(target) > 0 {
		mod := new(big.Int)
		count, mod := new(big.Int).DivMod(poshash, target, mod)
		used_weight = count.Uint64()

		if mod.Cmp(common.Big0) > 0 {
			used_weight += 1
		}
	} else {
		used_weight = 1
	}

	if weight < used_weight {
		return nil, 0
	}

	log.Trace("PoS hash",
		"target", target,
		"poshash", poshash,
		"used_weight", used_weight,
		"weight", weight)
	return poshash, used_weight
}

func (e *Energi) verifyPoSHash(
	chain ChainReader,
	header *types.Header,
) error {
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	weight, err := e.lookupStakeWeight(chain, header.Time, parent, header.Coinbase)
	if err != nil {
		return err
	}

	target := new(big.Int).Div(diff1Target, header.Difficulty)

	poshash, used_weight := e.calcPoSHash(header, target, weight)

	if poshash == nil {
		return errInvalidPoSHash
	}

	if used_weight != header.Nonce.Uint64() {
		return errInvalidPoSNonce
	}

	return nil
}

/**
 * Implements stake amount calculation.
 *
 * POS-3: Stake maturity period
 * POS-4: Stake amount
 * POS-22: Partial stake amount
 *
 * This is a basic helper for stake amount calculation.
 * There are ways to optimize it for high load, but we need something
 * to start with.
 */
func (e *Energi) lookupStakeWeight(
	chain ChainReader,
	now uint64,
	till *types.Header,
	addr common.Address,
) (weight uint64, err error) {
	var since uint64

	if now > MaturityPeriod {
		since = now - MaturityPeriod
	} else {
		since = 0
	}

	// NOTE: Do not set to high initial value due to defensive coding approach!
	weight = 0
	total_staked := uint64(0)
	first_run := true
	blockst := chain.CalculateBlockState(till.Hash(), till.Number.Uint64())

	// NOTE: we need to ensure at least one iteration with the balance condition
	for (till.Time > since) || first_run {
		if blockst == nil {
			log.Warn("PoS state root failure", "header", till.Hash())
			return 0, eth_consensus.ErrMissingState
		}

		weight_at_block := new(big.Int).Div(
			blockst.GetBalance(addr),
			minStake,
		).Uint64()

		if first_run {
			weight = weight_at_block
			first_run = false
		}

		// Find the minimum balance
		if weight > weight_at_block {
			weight = weight_at_block
		}

		// No need to lookup further
		if weight < 1 {
			break
		}

		// POS-22: partial stake amount
		if till.Coinbase == addr {
			total_staked += till.Nonce.Uint64()
		}

		curr := till
		parent_number := curr.Number.Uint64() - 1
		till = chain.GetHeader(curr.ParentHash, parent_number)

		if till == nil {
			if curr.Number.Cmp(common.Big0) == 0 {
				break
			}

			log.Error("PoS state missing parent", "parent", curr.ParentHash)
			return 0, eth_consensus.ErrUnknownAncestor
		}

		blockst = chain.CalculateBlockState(curr.ParentHash, parent_number)
	}

	if weight < total_staked {
		log.Debug("Nothing to stake",
			"addr", addr, "since", since, "weight", weight, "total_staked", total_staked)
		weight = 0
	} else {
		weight -= total_staked
	}

	//log.Trace("PoS stake weight", "addr", addr, "weight", weight)
	return weight, nil
}

/**
 * POS-19: PoS miner implementation
 */
func (e *Energi) mine(
	chain ChainReader,
	header *types.Header,
	stop <-chan struct{},
) (success bool, err error) {
	type Candidates struct {
		addr   common.Address
		weight uint64
	}

	accounts := e.accountsFn()
	if len(accounts) == 0 {
		select {
		case <-stop:
			return false, nil
		}
	}

	candidates := make([]Candidates, 0, len(accounts))
	migration_dpos := false
	for _, a := range accounts {
		candidates = append(candidates, Candidates{
			addr:   a,
			weight: 0,
		})
		//log.Trace("PoS miner candidate found", "address", a)

		if a == energi_params.Energi_MigrationContract {
			migration_dpos = true
		}
	}

	//---
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)

	if parent == nil {
		return false, eth_consensus.ErrUnknownAncestor
	}

	time_target := e.calcTimeTarget(chain, parent)

	blockTime := time_target.min_time

	// Special case due to expected very large gap between Genesis and Migration
	if header.IsGen2Migration() && !e.testing {
		blockTime = e.now()
	}

	// A special workaround to obey target time when migration contract is used
	// for mining to prevent any difficult bombs.
	if migration_dpos && !e.testing {
		// Obey block target
		if blockTime < time_target.block_target {
			blockTime = time_target.block_target
		}

		// Also, obey period target
		if blockTime < time_target.period_target {
			blockTime = time_target.period_target
		}

		// Decrease difficulty, if it got bumped
		if header.Difficulty.Uint64() > diffV1_MigrationStakerTarget {
			blockTime += diffV1_MigrationStakerDelay
		}
	}

	//---
	for ; ; blockTime++ {
		if max_time := e.now() + MaxFutureGap; blockTime > max_time {
			log.Trace("PoS miner is sleeping")
			select {
			case <-stop:
				// NOTE: it's very important to ignore stop until all variants are tried
				//       to prevent rogue stakers taking the initiative.
				return false, nil
			case <-time.After(time.Duration(blockTime-max_time) * time.Second):
			}
		}

		if e.peerCountFn() == 0 {
			log.Trace("Skipping PoS miner due to missing peers")
			continue
		}

		header.Time = blockTime
		time_target, err = e.posPrepare(chain, header, parent)
		if err != nil {
			return false, err
		}

		target := new(big.Int).Div(diff1Target, header.Difficulty)
		log.Trace("PoS miner time", "time", blockTime)

		// It could be done once, but then there is a chance to miss blocks.
		// Some significant algo optimizations are possible, but we start with simplicity.
		for i := range candidates {
			v := &candidates[i]
			v.weight, err = e.lookupStakeWeight(
				chain, blockTime, parent, v.addr)
			if err != nil {
				return false, err
			}
		}
		// Try smaller amounts first
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].weight < candidates[j].weight
		})
		// Try to match target
		for i := range candidates {
			v := &candidates[i]
			if v.weight < 1 {
				continue
			}

			//log.Trace("PoS stake candidate", "addr", v.addr, "weight", v.weight)
			header.Coinbase = v.addr
			poshash, used_weight := e.calcPoSHash(header, target, v.weight)

			nonceCap := e.GetMinerNonceCap()
			if nonceCap != 0 && nonceCap < used_weight {
				continue
			} else if poshash != nil {
				log.Trace("PoS stake", "addr", v.addr, "weight", v.weight, "used_weight", used_weight)
				header.Nonce = types.EncodeNonce(used_weight)
				return true, nil
			}
		}
	}

	return false, nil
}
