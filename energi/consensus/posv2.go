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

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/log"
)

const (
	// we compute in microseconds so that we can do integer math with higher precision
	microseconds uint64 = 1000000

	// just here to avoid extra typecasts in calculation
	two uint64 = 2
)

// CalculateBlockTimeEMA computes the exponential moving average of block times
// this will return the EMA of block times as microseconds
// for a description of the EMA algorithm, please see:
// see https://www.itl.nist.gov/div898/handbook/pmc/section4/pmc431.htm
func CalculateBlockTimeEMA(blockTimeDifferences []uint64, emaPeriod uint64) (ema []uint64) {
	sampleSize := len(blockTimeDifferences)
	N := emaPeriod + 1
	ema = make([]uint64, sampleSize)

	// choice of initial condition is important for an EMA. We could use the first
	// block time difference, but instead we'll set it to the target value so our
	// EMA will tend toward the target. However we don't include this value in our
	// EMA series data that we return, we only use it to calculate the first EMA
	emaPrev := params.TargetBlockGap * microseconds
	for i := 0; i < sampleSize; i++ {
		// this formula has a factor of 2/(emaPeriod+1) in a couple places. This is our
		// smoothing coefficient for the EMA, often referred to as alpha. We have
		// not precomputed this value so we don't lose precision on early division
		ema[i] = ((two * blockTimeDifferences[i] * microseconds) + (emaPrev * (N - two))) / N
		emaPrev = ema[i]
	}
	return
}

// CalculateBlockTimeDrift calculates the difference between the target block time
// and the EMA block time. Drift should be a negative value if blocks are too slow
// and a positive value if blocks are too fast, representing the direction
// to adjust the difficulty
func CalculateBlockTimeDrift(ema []uint64) (drift []int64) {
	target := int64(params.TargetBlockGap * microseconds)
	drift = make([]int64, len(ema))
	for i := range ema {
		drift[i] = target - int64(ema[i])
	}
	return
}

// CalculateBlockTimeIntegral calculates the integral of the block drift function
// This provides us with some idea fo historical "error", how far the block time
// has been from the target value for the duration of the period
// We use the trapezoidal rule here for integration
func CalculateBlockTimeIntegral(drift []int64) (integral int64) {
	sampleSize := len(drift)
	integral = 0
	// this is a simplification of the trapezoid rule based on uniform spacing
	for i := 1; i < sampleSize - 1; i++ {
		integral += drift[i]
	}
	integral += (drift[0] + drift[sampleSize-1]) / 2
	return
}

// CalculateBlockTimeDerivative computes the derivative series of a data series
// Here we use the central difference formula, for some small step h (each block)
// f'(x) = 1/2h * (f(x+h) - f(x-h))
func CalculateBlockTimeDerivative(drift []int64) (derivative []int64) {
	sampleSize := len(drift)
	derivative = make([]int64, sampleSize - 1)

	for i := 1; i < sampleSize; i++ {
		derivative[i-1] = (drift[i] - drift[i-1])
	}
	return
}

/*
 * Block Time Target Calculation V2
 * @chain Current Chain
 * @parent Parent Block Header
 * @ret Time Target structure
 * Populates ret with an updated Time Target
 * Calculates a Target Block Time based on previous block times in order to maintain a 60 second average time
 * Implements the Exponential Moving Average in calculating the block target time
 * Based on the last 60 elapsed block times
 * A block cannot be created with a time greater than 3 seconds in the future
 * ~~The minimum block time is 30 seconds~~ - This should not be enforced
here as an early or late target is for difficulty adjustment not the block
timestamp
*/
func (e * Energi) calcTimeTargetV2(chain ChainReader, parent *types.Header) *TimeTarget {
	// check if we have already calculated
	if parent.Hash() == e.calculatedBlockHash {
		timeTarget := e.calculatedTimeTarget
		return &timeTarget
	}

	ret := &TimeTarget{}
	parentBlockTime := parent.Time // Defines the original parent block time.
	parentNumber := parent.Number.Uint64()

	// POS-11: Block time restrictions
	ret.max = e.now() + params.MaxFutureGap

	// POS-11: Block time restrictions
	ret.min = parentBlockTime + params.MinBlockGap
	ret.blockTarget = parentBlockTime + params.TargetBlockGap
	ret.periodTarget = ret.blockTarget

	// Block interval enforcement
	// TODO: LRU cache here for extra DoS mitigation
	timeDiffs := make([]uint64, params.BlockTimeEMAPeriod)

	// compute block time differences
	// note that the most recent time difference will be the most
	// weighted by the EMA, and the oldest time difference will be the least
	for i := params.BlockTimeEMAPeriod; i > 0; i-- {
		past := chain.GetHeader(parent.ParentHash, parent.Number.Uint64()-1)
		if past == nil {
			// this normally can't happen because there is more
			// than enough blocks before the hard fork to always
			// get params.BlockTimeEMAPeriod timestamps
			log.Trace("Inconsistent tree, shutdown?")
			return ret
		}
		timeDiffs[i-1] = parent.Time - past.Time
		parent = past
	}

	ema := CalculateBlockTimeEMA(timeDiffs, params.BlockTimeEMAPeriod)

	ret.periodTarget = ema[len(ema)-1]

	// set up the parameters for PID control (diffV2)
	drift := CalculateBlockTimeDrift(ema)
	integral := CalculateBlockTimeIntegral(drift)
	derivative := CalculateBlockTimeDerivative(drift)
	ret.Drift = drift[len(drift)-1]
	ret.Integral = integral
	ret.Derivative = derivative[len(derivative)-1]

	log.Trace("PoS time", "block", parentNumber+1,
		"min", ret.min, "max", ret.max,
		"TimeTarget", ret.blockTarget,
		"averageBlockTimeMicroseconds", ret.periodTarget,
	)

	// set calculated results for optimization
	e.calculatedBlockHash = parent.Hash()
	e.calculatedTimeTarget = *ret

	return ret
}

// CalcPoSDifficultyV2 is our v2 difficulty algorithm
// this algorithm is a PID controlled difficulty
// first we take an Exponential Moving Average of
// the last 60 elapsed block times. EMA was chosen because it
// favors more recent block times, and so should be more responsive.
// Then we compute the drift, which is the difference between EMA
// block time, and the target time of 60 seconds.
// Finally, we take the integral and derivative of the drift.
// This gives us 3 terms for PID control:
// proportional (drift)
// integral
// derivative
//
// A PID controller is an excellent way to remove oscillation when
// approaching a target value. To describe the difficulty algorithm
// as a PID controller we need a set point, a process variable,
// and a control variable.
//
// The set point is our 60 second block time. Block time EMA is our
// process variable. The difficulty itself is the control variable.
// We calculate a new difficulty as a weighted sum of the difference
// between the set point and process variable,
// the integral of this difference, and the derivative of this difference.
//
// The proportional term accounts for current error in block time.
// The integral term accounts for past error in block time.
// The derivative term accounts for future error in block time.
// By carefully weighting these 3, we can quickly approach the set point
// without much oscillation.
//
// The PID control implemented here is generally called the "standard form"
// which has only a single gain, and the derivative and integral terms are
// scaled by time.
//
// See https://en.wikipedia.org/wiki/PID_controller#Mathematical_form for more
// information.
func CalcPoSDifficultyV2(
	newBlockTime uint64, // TODO: maybe we should be recalculating ema/drift/etc with the new time included?
	parent *types.Header,
	timeTarget *TimeTarget,
) *big.Int {
	// set tuning parameters
	gain := big.NewInt(50000)
	integralTime := big.NewInt(720)
	derivativeTime := big.NewInt(60)

	difficultyAdjustmentProportional := big.NewInt(timeTarget.Drift)
	difficultyAdjustmentIntegral := big.NewInt(timeTarget.Integral)
	difficultyAdjustmentIntegral.Div(difficultyAdjustmentIntegral, integralTime)
	difficultyAdjustmentDerivative := big.NewInt(timeTarget.Derivative)
	difficultyAdjustmentDerivative.Mul(difficultyAdjustmentDerivative, derivativeTime)

	difficultyAdjustment := big.NewInt(0)
	difficultyAdjustment.Add(difficultyAdjustment, difficultyAdjustmentProportional)
	difficultyAdjustment.Add(difficultyAdjustment, difficultyAdjustmentIntegral)
	difficultyAdjustment.Add(difficultyAdjustment, difficultyAdjustmentDerivative)
	difficultyAdjustment.Mul(difficultyAdjustment, gain)
	difficultyAdjustment.Div(difficultyAdjustment, big.NewInt(int64(microseconds)))

	difficulty := big.NewInt(0).Set(parent.Difficulty)
	difficulty.Add(difficulty, difficultyAdjustment)

	// ensure the difficulty does not fall below 1
	if difficulty.Cmp(common.Big1) < 0 {
		difficulty = common.Big1
	}

	log.Trace("Difficulty adjustment",
		"parent", parent.Difficulty, "new difficulty", difficulty,
		"block time", newBlockTime, "target time", timeTarget)
	return difficulty
}
