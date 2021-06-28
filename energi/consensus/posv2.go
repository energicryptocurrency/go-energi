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
	"math"
	"math/big"
	"sort"
	"time"

	"energi.world/core/gen3/common"
	eth_consensus "energi.world/core/gen3/consensus"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/energi/diffv1"
	energi_params "energi.world/core/gen3/energi/params"
	"energi.world/core/gen3/log"
)

const (
	targetWindow          = energi_params.TargetWindow
	maxTimeDifferenceDrop = energi_params.MaxTimeDifferenceDrop
	difficultyChangeBase  = energi_params.DifficultyChangeBase

	diffV2MigrationStakerTimeDelay  = energi_params.DiffV2MigrationStakerTimeDelay
	diffV2MigrationStakerBlockDelay = energi_params.DiffV2MigrationStakerBlockDelay
	diffV2MigrationStakerTarget     = energi_params.DiffV2MigrationStakerTarget
)

type timeTargetV2 struct {
	minTime uint64
	maxTime uint64
	target  uint64
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
	return t.minTime
}

func (t *timeTargetV2) setMinTime(minTime uint64) {
	t.minTime = minTime
}

func (t *timeTargetV2) getMaxTime() uint64 {
	return t.maxTime
}

func (t *timeTargetV2) setMaxTime(maxTime uint64) {
	t.maxTime = maxTime
}

func (t *timeTargetV2) getTarget() uint64 {
	return t.target
}

func (t *timeTargetV2) setTarget(target uint64) {
	t.target = target
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
 * The minimum block time is 30 seconds
 */
func (e *Energi) calcTimeTargetV2(chain ChainReader, parent *types.Header) *timeTargetV2 {
	ret := &timeTargetV2{}
	parentBlockTime := parent.Time // Defines the original parent block time.
	parentNumber := parent.Number.Uint64()
	smoothingFactor := 2.0 / float64(targetWindow+1)

	// POS-11: Block time restrictions
	ret.maxTime = e.now() + energi_params.MaxFutureGap

	// POS-11: Block time restrictions
	ret.minTime = parentBlockTime + energi_params.MinBlockGap
	ret.target = parentBlockTime + energi_params.TargetBlockGap
	ret.pHash = parent.Hash()

	// Block interval enforcement
	// ---
	if parentNumber > energi_params.AverageTimeBlocks {
		// TODO: LRU cache here for extra DoS mitigation
		timeDiff := make([]float64, energi_params.AverageTimeBlocks)

		// NOTE: we have to do this way as parent may be not part of canonical
		//       chain. As no mutex is held, we cannot do checks for canonical.
		for i := energi_params.AverageTimeBlocks; i > 0; i-- {
			past := chain.GetHeader(parent.ParentHash, parent.Number.Uint64()-1)
			if past == nil {
				log.Trace("Inconsistent tree, shutdown?")
				return ret
			}
			timeDiff[i-1] = float64(parent.Time - past.Time)
			parent = past
		}

		// Holds the simple moving average of blocktime difference calculated
		// at the moving average interval window.
		SMA := make([]float64, (energi_params.AverageTimeBlocks - targetWindow + 1))
		var sum float64
		for i := 0; i < len(SMA); i++ {
			if sum == 0 {
				for _, val := range timeDiff[:targetWindow] {
					sum += val
				}
			} else {
				sum = sum - timeDiff[i-1]
				sum = sum + timeDiff[i+int(targetWindow)-1]
			}
			// Obtain average at the specified target window.
			SMA[i] = sum / float64(targetWindow)
		}

		// Holds the exponential moving average calculated from the simple moving
		// list average.
		EMA := make([]float64, (energi_params.AverageTimeBlocks - targetWindow + 1))
		for i, val := range SMA {
			forecastedDiff := SMA[0]
			if i > 0 {
				forecastedDiff = (val * smoothingFactor) + (1-smoothingFactor)*EMA[i-1]
			}
			EMA[i] = forecastedDiff
		}

		forecastTimeDiff := uint64(EMA[len(EMA)-1])
		if forecastTimeDiff > energi_params.TargetBlockGap {
			// Max block gap should not exceed value defined in TargetBlockGap.
			forecastTimeDiff = energi_params.TargetBlockGap
		}

		ret.target = parentBlockTime + forecastTimeDiff
	}

	log.Trace("PoS time", "block", parentNumber+1,
		"min", ret.minTime, "max", ret.maxTime,
		"timeTarget", ret.target,
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
	// Find out our target anchor
	target := timeTarget.target
	if target < timeTarget.minTime {
		target = timeTarget.minTime
	}

	if newBlockTime == target {
		log.Trace("No difficulty change", "parent", parent.Difficulty)
		return parent.Difficulty
	}

	// TimeDifference = timeTarget - newBlockTime
	timeDiff := float64(target) - float64(newBlockTime)
	if timeDiff < maxTimeDifferenceDrop {
		timeDiff = maxTimeDifferenceDrop
	}

	preMultiplier := big.NewFloat(math.Pow(difficultyChangeBase, timeDiff))

	// To reduce the possibility of difficulty collision likely to happen between
	// simultaneously mining nodes, add the current nanoseconds as the
	// mantissa to the multiplier. Divide by 10^15 to make a mantissa with
	// a small to negligible effect to the overall difficulty value.
	// Max mantissa added => 999999999.0/1000000000000000.0 = 0.000000999999999 = 9.99999999e-7
	salt := big.NewFloat(float64(time.Now().Nanosecond()) / 1000000000000000.0) // mantissa
	multiplier := new(big.Float).Add(preMultiplier, salt)

	parentDiff := new(big.Float).SetInt(parent.Difficulty)
	difficulty := new(big.Int)
	new(big.Float).Mul(parentDiff, multiplier).Int(difficulty)

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
		case <-time.After(10 * time.Second):
			// If no mining accounts that are found in 10 seconds quit.
			accounts = e.accountsFn()
			if len(accounts) == 0 {
				log.Error("No mining candidate accounts found: timeout")
				return false, nil
			}
		case <-stop:
			log.Error("No mining candidate accounts found")
			return false, nil
		}
	}

	candidates := make([]Candidates, 0, len(accounts))
	migrationDPOS := false
	for _, a := range accounts {
		candidates = append(candidates, Candidates{
			addr:   a,
			weight: 0,
		})
		// log.Trace("PoS miner candidate found", "address", a)

		if a == energi_params.Energi_MigrationContract {
			migrationDPOS = true
		}
	}

	// ---
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)

	if parent == nil {
		return false, eth_consensus.ErrUnknownAncestor
	}

	blockTarget := e.calcTimeTargetV2(chain, parent)

	blockTime := blockTarget.minTime

	// Special case due to expected very large gap between Genesis and Migration
	if header.IsGen2Migration() && !e.testing {
		blockTime = e.now()
	}

	// A special workaround to obey target time when migration contract is used
	// for mining to prevent any difficult bombs.
	if migrationDPOS && !e.testing && header.Number.Uint64(
		) < energi_params.DiffV2MigrationStakerBlockDelay {
		// Obey block target
		if blockTime < blockTarget.target {
			blockTime = blockTarget.target
		}

		// Decrease difficulty, if it got bumped
		if header.Difficulty.Uint64() > energi_params.
			DiffV2MigrationStakerTarget {
			blockTime += energi_params.
				DiffV2MigrationStakerTimeDelay
		}
	}

	// ---
	for ; ; blockTime++ {
		select {
		case <-stop:
			// Ensure that a shutdown request is handled as fast as possible.
			return false, nil
		default:
			if maxTime := e.now() + energi_params.
				MaxFutureGap; blockTime > maxTime {
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
		if err = e.posPrepareV2(chain, header, parent, blockTarget); err != nil {
			return false, err
		}

		target := new(big.Int).Div(diffv1.Target, header.Difficulty)
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

			// log.Trace("PoS stake candidate", "addr", v.addr, "weight", v.weight)
			header.Coinbase = v.addr
			poshash, usedWeight := e.calcPoSHash(
				header,
				target,
				v.weight,
			)
			header.Nonce = types.EncodeNonce(usedWeight)

			nonceCap := e.GetMinerNonceCap()
			if nonceCap != 0 && e.nonceCap < usedWeight {
				continue
			} else if poshash != nil {
				log.Trace("PoS stake", "addr", v.addr, "weight", v.weight, "usedWeight", usedWeight)

				return true, nil
			}
		}
	}
}
