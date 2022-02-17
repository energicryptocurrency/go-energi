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

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/consensus"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/crypto"
	"github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/energi/api/hfcache"
)

var (
	minStake    = big.NewInt(1e18) // 1 NRG
	diff1Target = new(big.Int).Exp(
		big.NewInt(2), big.NewInt(256), big.NewInt(0),
	)

	errBlockMinTime    = errors.New("block is before minimum time")
	errInvalidPoSHash  = errors.New("invalid PoS hash")
	errInvalidPoSNonce = errors.New("invalid stake weight")
)

type TimeTarget struct {
	min, max, blockTarget, periodTarget uint64
	Drift, Integral, Derivative         int64
}

/**
 * Implements block time consensus
 *
 * POS-11: Block time restrictions
 * POS-12: Block interval enforcement
 */
func (e *Energi) calcTimeTargetV1(
	chain ChainReader, parent *types.Header,
) (ret *TimeTarget) {
	ret = new(TimeTarget)
	now := e.now()
	parentNumber := parent.Number.Uint64()
	blockNumber := parentNumber + 1

	// POS-11: Block time restrictions
	ret.max = now + params.MaxFutureGap

	// POS-11: Block time restrictions
	ret.min = parent.Time + params.MinBlockGap
	ret.blockTarget = parent.Time + params.TargetBlockGap
	ret.periodTarget = ret.blockTarget

	// POS-12: Block interval enforcement
	// ---
	if blockNumber >= params.AveragingWindow {
		// TODO: LRU cache here for extra DoS mitigation
		past := parent

		// NOTE: we have to do this way as parent may be not part of canonical
		//       chain. As no mutex is held, we cannot do checks for canonical.
		for i := params.AveragingWindow - 1; i > 0; i-- {
			past = chain.GetHeader(past.ParentHash, past.Number.Uint64()-1)

			if past == nil {
				log.Trace("Inconsistent tree, shutdown?")
				return
			}
		}

		ret.periodTarget = past.Time + params.TargetPeriodGap
		periodMinTime := ret.periodTarget - params.MinBlockGap

		if periodMinTime > ret.min {
			ret.min = periodMinTime
		}
	}

	log.Trace(
		"PoS time", "block", blockNumber,
		"min", ret.min, "max", ret.max,
		"blockTarget", ret.blockTarget,
		"periodTarget", ret.periodTarget,
	)
	return
}

func (e *Energi) enforceMinTime(
	header *types.Header, timeTarget *TimeTarget,
) error {
	// NOTE: allow Miner to hint already tried period by
	if header.Time < timeTarget.min {
		header.Time = timeTarget.min
	}

	return nil
}

func (e *Energi) checkTime(header *types.Header, timeTarget *TimeTarget) error {

	if header.Time < timeTarget.min {
		return errBlockMinTime
	}

	// Check if able to mine
	if header.Time > timeTarget.max {
		return consensus.ErrFutureBlock
	}

	return nil
}

/**
 * Implements check modifier consensus
 *
 * POS-14: Stake modifier
 */
func (e *Energi) calcPoSModifier(
	chain ChainReader, time uint64, parent *types.Header,
) (ret common.Hash) {
	// TODO: LRU cache here for extra DoS mitigation

	// Find maturity period border
	maturityBorder := time

	// maturity period is reduced to 30m in Asgard
	maturityPeriod := params.MaturityPeriod
	if !e.testing {
		// check if Asgard hardfork is activated use new difficulty algorithm
		// check if Asgard hardfork is activated use new difficulty algorithm
		isAsgardActive := hfcache.IsHardforkActive("Asgard", parent.Number.Uint64())
		log.Debug("hf check", "isAsgardActive", isAsgardActive)
		// don't check for hard forks being active if we're testing
		if e.testing {
			isAsgardActive = false
		}
		if isAsgardActive {
			maturityPeriod = params.MaturityPeriodAsgard
		}
	}

	if maturityBorder < maturityPeriod {
		// This should happen only in testing
		maturityBorder = 0
	} else {
		maturityBorder -= maturityPeriod
	}

	// Find the oldest inside maturity period
	// NOTE: we have to do this walk as parent may not be part of the canonical chain
	parentHeight := parent.Number.Uint64()
	oldest := parent

	for header, num := oldest, oldest.Number.Uint64(); (header.Time > maturityBorder) && (num > 0); {

		oldest = header
		num--
		header = chain.GetHeader(header.ParentHash, num)
	}

	// Create Stake Modifier
	//
	// The stake modifier is computed by hashing the parent coinbase and the root state of the block nearest to the
	// maturityBorder
	ret = crypto.Keccak256Hash(
		parent.Coinbase.Bytes(),
		oldest.Root.Bytes(),
	)

	log.Trace(
		"PoS modifier", "block", parentHeight+1,
		"modifier", ret, "oldest", oldest.Number.Uint64(),
	)
	return
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
	time uint64,
	parent *types.Header,
	tt *TimeTarget,
) (D *big.Int) {

	// Find the target anchor
	target := (tt.blockTarget + tt.periodTarget) / 2
	if target < tt.min {
		target = tt.min
	}

	if time < target {
		targetDelta := target - time
		if targetDelta > diffV1_BMax {
			targetDelta = diffV1_BMax
		}

		B := diffV1_BTable[targetDelta]
		D = new(big.Int).Div(new(big.Int).Mul(parent.Difficulty, B), diffV1_Div)
		log.Trace("Diff multiplier", "before", targetDelta, "mult", B)

	} else if time > target {
		targetDelta := time - target
		// clamp the target delta to max
		if targetDelta > diffV1_AMax {
			targetDelta = diffV1_AMax
		}
		A := diffV1_ATable[targetDelta]
		D = new(big.Int).Div(
			new(big.Int).Mul(parent.Difficulty, diffV1_Div),
			A,
		)
		log.Trace("Diff multiplier", "after", targetDelta, "div", A)

	} else {
		log.Trace("No difficulty change", "parent", parent.Difficulty)
		return parent.Difficulty
	}

	if D.Cmp(common.Big1) < 0 {
		D = common.Big1
	}

	log.Trace(
		"Difficulty change",
		"parent", parent.Difficulty, "new", D, "time", time, "target", target,
	)
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
) (posHash *big.Int, usedWeight uint64) {
	serializedTime := [8]byte{}
	binary.BigEndian.PutUint64(serializedTime[:], header.Time)

	posHash = new(big.Int).SetBytes(
		crypto.Keccak256(
			serializedTime[:],
			header.MixDigest.Bytes(),
			header.Coinbase.Bytes(),
		),
	)

	if posHash.Cmp(target) > 0 {
		count, mod := new(big.Int).DivMod(posHash, target, new(big.Int))
		usedWeight = count.Uint64()

		if mod.Cmp(common.Big0) > 0 {
			usedWeight += 1
		}

	} else {
		usedWeight = 1
	}

	if weight < usedWeight {
		return nil, 0
	}

	log.Trace(
		"PoS hash",
		"target", target,
		"posHash", posHash,
		"used_weight", usedWeight,
		"weight", weight,
	)
	return posHash, usedWeight
}

func (e *Energi) verifyPoSHash(chain ChainReader, header *types.Header) error {
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	weight, err := e.lookupStakeWeight(
		chain, header.Time, parent, header.Coinbase,
	)
	if err != nil {
		return err
	}

	target := new(big.Int).Div(diff1Target, header.Difficulty)

	posHash, usedWeight := e.calcPoSHash(header, target, weight)

	if posHash == nil {
		return errInvalidPoSHash
	}

	if usedWeight != header.Nonce.Uint64() {
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
	until *types.Header,
	addr common.Address,
) (weight uint64, err error) {
	var since uint64

	// maturity period is reduced to 30m in Asgard
	maturityPeriod := params.MaturityPeriod
	if !e.testing {
		// check if Asgard hardfork is activated use new difficulty algorithm
		isAsgardActive := hfcache.IsHardforkActive("Asgard", until.Number.Uint64())
		log.Debug("hf check", "isAsgardActive", isAsgardActive)
		// don't check for hard forks being active if we're testing
		if e.testing {
			isAsgardActive = false
		}
		if isAsgardActive {
			maturityPeriod = params.MaturityPeriodAsgard
		}
	}

	if now > maturityPeriod {
		since = now - maturityPeriod
	} else {
		since = 0
	}

	// NOTE: Do not set to high initial value due to defensive coding approach!
	weight = 0
	totalStaked := uint64(0)
	firstRun := true
	blockState := chain.CalculateBlockState(until.Hash(), until.Number.Uint64())

	// NOTE: we need to ensure at least one iteration with the balance condition
	for (until.Time > since) || firstRun {

		if blockState == nil {
			log.Warn("PoS state root failure", "header", until.Hash())
			return 0, consensus.ErrMissingState
		}

		weightAtBlock := new(big.Int).Div(
			blockState.GetBalance(addr),
			minStake,
		).Uint64()

		if firstRun {
			weight = weightAtBlock
			firstRun = false
		}

		// Find the minimum balance
		if weight > weightAtBlock {
			weight = weightAtBlock
		}

		// No need to lookup further
		if weight < 1 {
			break
		}

		// POS-22: partial stake amount
		if until.Coinbase == addr {
			totalStaked += until.Nonce.Uint64()
		}

		curr := until
		parentNumber := curr.Number.Uint64() - 1
		until = chain.GetHeader(curr.ParentHash, parentNumber)

		if until == nil {

			if curr.Number.Cmp(common.Big0) == 0 {
				break
			}

			log.Error("PoS state missing parent", "parent", curr.ParentHash)
			return 0, consensus.ErrUnknownAncestor
		}

		blockState = chain.CalculateBlockState(curr.ParentHash, parentNumber)
	}

	if weight < totalStaked {
		log.Debug(
			"Nothing to stake",
			"addr", addr, "since", since, "weight", weight, "total_staked",
			totalStaked,
		)
		weight = 0
	} else {
		weight -= totalStaked
	}

	// log.Trace("PoS stake weight", "addr", addr, "weight", weight)
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
		<-stop
		return false, nil
	}

	candidates := make([]Candidates, 0, len(accounts))
	migration_dpos := false
	for _, a := range accounts {
		candidates = append(candidates, Candidates{
			addr:   a,
			weight: 0,
		})
		//log.Trace("PoS miner candidate found", "address", a)

		if a == params.Energi_MigrationContract {
			migration_dpos = true
		}
	}

	//---
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)

	if parent == nil {
		return false, consensus.ErrUnknownAncestor
	}

	// check if Asgard hardfork is activated use new difficulty algorithm
	isAsgardActive := hfcache.IsHardforkActive("Asgard", header.Number.Uint64())
	log.Debug("hf check", "isAsgardActive", isAsgardActive)
	// don't check for hard forks being active if we're testing
	if e.testing {
		isAsgardActive = false
	}

	// make time target calculation depending on asgard status
	var timeTarget *TimeTarget
	if isAsgardActive {
		timeTarget = e.calcTimeTargetV2(chain, parent)
	} else {
		timeTarget = e.calcTimeTargetV1(chain, parent)
	}
	blockTime := timeTarget.min

	// Special case due to expected very large gap between Genesis and Migration
	if header.IsGen2Migration() && !e.testing {
		blockTime = e.now()
	}

	// A special workaround to obey target time when migration contract is used
	// for mining to prevent any difficult bombs.
	if migration_dpos && !isAsgardActive && !e.testing {
		// Obey block target
		if blockTime < timeTarget.blockTarget {
			blockTime = timeTarget.blockTarget
		}

		// Also, obey period target
		if blockTime < timeTarget.periodTarget {
			blockTime = timeTarget.periodTarget
		}

		// Decrease difficulty, if it got bumped
		if header.Difficulty.Uint64() > diffV1_MigrationStakerTarget {
			blockTime += diffV1_MigrationStakerDelay
		}
	}

	//---
	for ; ; blockTime++ {
		if maxTime := e.now() + params.MaxFutureGap; blockTime > maxTime {
			log.Trace("PoS miner is sleeping")
			select {
			case <-stop:
				// NOTE: it's very important to ignore stop until all variants are tried
				//       to prevent rogue stakers taking the initiative.
				return false, nil
			case <-time.After(time.Duration(blockTime-maxTime) * time.Second):
			}
		}

		if e.peerCountFn() == 0 {
			log.Trace("Skipping PoS miner due to missing peers")
			continue
		}

		header.Time = blockTime
		err = e.Prepare(chain, header)
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
			poshash, usedWeight := e.calcPoSHash(header, target, v.weight)

			nonceCap := e.GetMinerNonceCap()
			if nonceCap != 0 && nonceCap < usedWeight {
				continue
			} else if poshash != nil {
				log.Trace("PoS stake", "addr", v.addr, "weight", v.weight, "used_weight", usedWeight)
				header.Nonce = types.EncodeNonce(usedWeight)
				return true, nil
			}
		}
	}
}
