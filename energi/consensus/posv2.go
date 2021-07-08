// Copyright 2021 The Energi Core Authors
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
	"sort"
	"time"

	"energi.world/core/gen3/common"
	eth_consensus "energi.world/core/gen3/consensus"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/energi/diffv1"
	"energi.world/core/gen3/energi/params"
	"energi.world/core/gen3/log"
)

const (
	targetWindow          = params.TargetWindow
	maxTimeDifferenceDrop = params.MaxTimeDifferenceDrop
	difficultyChangeBase  = params.DifficultyChangeBase

	diffV2MigrationStakerTimeDelay  = params.DiffV2MigrationStakerTimeDelay
	diffV2MigrationStakerBlockDelay = params.DiffV2MigrationStakerBlockDelay
	diffV2MigrationStakerTarget     = params.DiffV2MigrationStakerTarget
)

type timeTargetV2 struct {
	min uint64
	max uint64
	blockTarget  uint64
	pHash   common.Hash
}

type mineTimeTarget interface {
	// getters
	getMinTime() uint64
	getMaxTime() uint64
	getTarget() uint64
	getPeriodTarget() interface{}

	// setters
	setMinTime(uint64)
	setMaxTime(uint64)
	setTarget(uint64)
}

func (t *timeTargetV2) getMinTime() uint64 {
	return t.min
}

func (t *timeTargetV2) setMinTime(minTime uint64) {
	t.min = minTime
}

func (t *timeTargetV2) getMaxTime() uint64 {
	return t.max
}

func (t *timeTargetV2) setMaxTime(maxTime uint64) {
	t.max = maxTime
}

func (t *timeTargetV2) getTarget() uint64 {
	return t.blockTarget
}

func (t *timeTargetV2) setTarget(target uint64) {
	t.blockTarget = target
}

func (t *timeTargetV2) getPeriodTarget() interface{} {
	return t.pHash
}

/*
 * Block Time Target Calculation V2
 * @chain Current Chain
 * @parent Parent Block Header
 * @ret Time Target structure
 * Populates ret with an updated Time Target
 * Calculates a Target Block Time based on previous block times in order to maintain a 60 second average time
 * Implements the Exponential Moving Average in calculating the block target time
 * Based on the last 60 blocks
 * A block cannot be created with a time greater than 3 seconds in the future
 * ~~The minimum block time is 30 seconds~~ - This should not be enforced
here as an early or late target is for difficulty adjustment not the block
timestamp
*/
func (e *Energi) calcTimeTargetV2(chain ChainReader, parent *types.Header) *timeTargetV2 {

	ret := &timeTargetV2{}
	parentBlockTime := parent.Time // Defines the original parent block time.
	parentNumber := parent.Number.Uint64()

	// POS-11: Block time restrictions
	ret.max = e.now() + params.MaxFutureGap

	// POS-11: Block time restrictions
	ret.min = parentBlockTime + params.MinBlockGap
	ret.blockTarget = parentBlockTime + params.TargetBlockGap

	ret.pHash = parent.Hash()

	// Block interval enforcement
	// TODO: LRU cache here for extra DoS mitigation
	timeDiffs := make([]uint64, params.AveragingWindow)

	// NOTE: we have to do this way as parent may be not part of canonical
	//       chain. As no mutex is held, we cannot do checks for canonical.
	for i := params.AveragingWindow - 1; i > 0 && parent.Number.Uint64() > 3; i-- {
		past := chain.GetHeader(parent.ParentHash, parent.Number.Uint64()-1)
		if past == nil {
			// this normally can't happen because there is more
			// than enough blocks before the hard fork to always
			// get params.AveragingWindow timestamps
			log.Trace("Inconsistent tree, shutdown?")
			return ret
		}
		timeDiffs[i] = parent.Time - past.Time
		parent = past
	}

	emaLast := CalcEMAUint64(timeDiffs, 2, params.SMAPeriod+1, params.SMAPeriod)
	if emaLast > params.TargetBlockGap {
		// Max block gap should not exceed value defined in TargetBlockGap.
		emaLast = params.TargetBlockGap
	}

	ret.blockTarget = parentBlockTime + emaLast

	log.Trace("PoS time", "block", parentNumber+1,
		"min", ret.min, "max", ret.max,
		"timeTarget", ret.blockTarget,
	)
	return ret
}

/*
 * Difficulty algorithm V2
 * Returns a difficulty value to be used in the next Block
 * @newBlockTime Last Block Time
 * @parent Parent Block Header
 * @timeTarget Target Block Time
 * If the block time is less than the minimum time, the difficulty must be increased
 * If the block time is the target time, the difficulty should stay the same
 * If the block time is more than the target time the difficulty must be reduced
 * New Difficulty = Parent Difficulty * (1.0001 ^ Block Time)
 * NB To reduce the possibility of a difficulty collision, a mantissa is added in the calculation
 */
func calcPoSDifficultyV2(
	newBlockTime uint64,
	parent *types.Header,
	timeTarget *timeTargetV2,
) *big.Int {

	target := timeTarget.blockTarget
	// if the target is the new block time we use the parent difficulty
	if newBlockTime == target {
		log.Trace("No difficulty change", "parent", parent.Difficulty)
		return parent.Difficulty
	}
	// The divergence from the target time to the new block time
	// determines the new difficulty
	targetDivergence := int(newBlockTime) - int(target)
	// clamp to minimum -30
	if targetDivergence < params.MaxTimeDifferenceDrop {
		targetDivergence = params.MaxTimeDifferenceDrop
	}
	// clamp to maximum 60
	if targetDivergence > int(params.TargetBlockGap) {
		targetDivergence = int(params.TargetBlockGap)
	}
	log.Debug(">>>","target", targetDivergence)
	const factorInverse = 10000               // 0.0001 is the same as 1/10000
	const precision = 1000000                 // we want 2 decimal places precision lower
	var scaledPreMultiplier = precision + 100 // this levels it to 1 by
	// dividing the result back by precision

	negative := false
	if targetDivergence < 0 {
		targetDivergence = -targetDivergence
		negative = true
	}
	for i := 0; i < targetDivergence; i++ {
		// the function of 1.0001 ^ timeDiff means the same as
		// repeatedly add 1/10000th to the previous result value as many
		// times as timeDiff, starting with an initial (scaled) value
		if !negative {
			scaledPreMultiplier += scaledPreMultiplier / factorInverse
		} else {
			scaledPreMultiplier -= scaledPreMultiplier / factorInverse
		}
	}
	// multiply the parent difficulty by the multiplier and divide back
	// by the precision value, applying the difficulty change without using
	// floating point numbers
	difficulty := big.NewInt(0).Mul(parent.Difficulty,
		big.NewInt(int64(scaledPreMultiplier)))
	difficulty = difficulty.Div(difficulty, big.NewInt(int64(precision)))

	log.Trace("Difficulty change",
		"parent", parent.Difficulty, "new difficulty", difficulty,
		"block time", newBlockTime, "target time", target)
	return difficulty
}

// MineV2 ...
//
// PoS V2 miner implementation
//
func (e *Energi) MineV2(
	chain ChainReader, header *types.Header, stop <-chan struct{},
) (success bool, err error) {

	type Candidates struct {
		addr   common.Address
		weight uint64
	}

	accounts := e.accountsFn()
	// if no accounts are found, just pause and wait for the stop signal
	//
	// todo: is this what is intended? Is there a case where this value can
	//  change but this thread is then dead? I think that very likely this
	//  should be a repeating loop that retries every minute or something like
	//  this in case an account is created that can be used, while the server is
	//  running.
	//  further thoughts: generally it seems like this function returns quite
	//  quickly so probably this is fine
	if len(accounts) == 0 {
		select {
		case <-stop:
			return
		}
	}

	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		err = eth_consensus.ErrUnknownAncestor
		return
	}

	blockTarget := e.calcTimeTargetV2(chain, parent)
	blockTime := blockTarget.min

	// Special case due to expected very large gap between Genesis and Migration
	if header.IsGen2Migration() && !e.testing {
		blockTime = e.now()
	}

	candidates := make([]Candidates, 0, len(accounts))
	migrationDPoS := false

	for _, a := range accounts {

		if a == params.Energi_MigrationContract {
			migrationDPoS = true
		}

		candidate := Candidates{a, 0}
		candidates = append(candidates, candidate)
	}

	// A special workaround to obey target time when migration contract is used
	// for mining to prevent any difficult bombs.
	if migrationDPoS && !e.testing && header.Number.Uint64() < params.DiffV2MigrationStakerBlockDelay {
		// Obey block target
		if blockTime < blockTarget.blockTarget {
			blockTime = blockTarget.blockTarget
		}

		// Decrease difficulty, if it got bumped
		if header.Difficulty.Uint64() > params.DiffV2MigrationStakerTarget {
			blockTime += params.DiffV2MigrationStakerTimeDelay
		}
	}

	// ---
	for ; ; blockTime++ {
		select {
		case <-stop:
			// Ensure that a shutdown request is handled as fast as possible.
			return false, nil
		default:
			if maxTime := e.now() + params.MaxFutureGap; blockTime > maxTime {
				// NOTE: it's very important to ignore stop until all variants are tried
				//       to prevent rogue stakers taking the initiative.
				log.Trace("PoS miner is sleeping", "seconds", blockTime-maxTime)
				<-time.After(time.Duration(blockTime-maxTime) * time.Second)
			}
		}

		if e.peerCountFn() == 0 {
			log.Trace("Skipping PoS miner due to missing peers")
			continue
		}

		header.Time = blockTime
		if blockTarget, err = e.posPrepareV2(chain, header, parent); err != nil {
			return false, err
		}

		target := new(big.Int).Div(diffv1.Target, header.Difficulty)
		log.Trace("PoS miner time", "time", blockTime)

		// It could be done once, but then there is a chance to miss blocks.
		// Some significant algo optimizations are possible, but we start with simplicity.
		for i := range candidates {
			candidate := &candidates[i]
			candidate.weight, err = e.lookupStakeWeight(
				chain, blockTime, parent, candidate.addr)
			if err != nil {
				return false, err
			}
		}
		// Try smaller amounts first
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].weight < candidates[j].weight
		})

		// This tries each candidate for each timestamp before progressing the
		// timestamp. If the reverse order was desired, the block time needs to
		// be saved and reset here. Since older is better, this is probably the
		// better sequence to work in.
		for _, candidate := range candidates {

			if candidate.weight < 1 {
				continue
			}

			// log.Trace("PoS stake candidate", "addr", candidate.addr, "weight", candidate.weight)
			header.Coinbase = candidate.addr
			poshash, usedWeight := e.calcPoSHash(header, target, candidate.weight)
			nonceCap := e.GetMinerNonceCap()

			if nonceCap != 0 && e.nonceCap < usedWeight {
				continue
			} else if poshash != nil {
				log.Trace("PoS stake", "addr", candidate.addr, "weight", candidate.weight, "usedWeight", usedWeight)
				header.Nonce = types.EncodeNonce(usedWeight)
				return true, nil
			}
		}
	}
	// this doesn't need to be strictly stated but this is a long function
	return success, err
}
