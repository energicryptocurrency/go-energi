package calculus

import (
	"encoding/binary"
	"math/big"

	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/params"
	"github.com/energicryptocurrency/energi/crypto"
)

var (
	// Bigs defined for simplicity
	big1          = big.NewInt(1)
	big0          = big.NewInt(0)
	diffV1_BTable []*big.Int
	diffV1_ATable []*big.Int
	diffV1_Div    = new(big.Int).SetUint64(diffV1_DivPlain)
)

const (
	// we compute in nanoseconds so that we can do integer math with higher precision
	microseconds uint64 = 1000000

	// just here to avoid extra typecasts in calculation
	two             uint64 = 2
	diffV1_BMax     uint64 = 30
	diffV1_AMax     uint64 = 120
	diffV1_DivPlain uint64 = 100

	// Roughly get 2x difficulty decrease
	//diffV1_MigrationStakerDelay  uint64 = 15
	//diffV1_MigrationStakerTarget uint64 = 0xFFFF
)

// TimeTarget used for next block calculation
type TimeTarget struct {
	Min, Max, BlockTarget, PeriodTarget uint64
	Drift, Integral, Derivative         int64
}

func initDiffTable(l uint64, c float64) []*big.Int {
	t := make([]*big.Int, l+1)
	t[0] = big1
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

//CalcPoSHash  Implements hash consensus
func CalcPoSHash(
	time uint64,
	coinbase string,
	target *big.Int,
) (usedWeight uint64) {
	serializedTime := [8]byte{}
	binary.BigEndian.PutUint64(serializedTime[:], time)

	posHash := new(big.Int).SetBytes(
		crypto.Keccak256(
			serializedTime[:],
			[]byte(coinbase),
		),
	)

	if posHash.Cmp(target) > 0 {
		count, mod := new(big.Int).DivMod(posHash, target, new(big.Int))
		usedWeight = count.Uint64()
		if mod.Cmp(big0) > 0 {
			usedWeight++
		}
	} else {
		usedWeight = 1
	}

	return usedWeight
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
// See htimeTargetps://en.wikipedia.org/wiki/PID_controller#Mathematical_form for more
// information.
func CalcPoSDifficultyV2(
	newBlockTime uint64, // TODO: maybe we should be recalculating ema/drift/etc with the new time included?
	diff uint64,
	timeTarget *TimeTarget,
) *big.Int {
	// set tuning parameters
	gain := big.NewInt(params.Gain)
	integralTime := big.NewInt(params.IntegralTime)
	derivativeTime := big.NewInt(params.DerivativeTime)

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

	difficulty := big.NewInt(0).Set(big.NewInt(int64(diff)))
	difficulty.Add(difficulty, difficultyAdjustment)

	// ensure the difficulty does not fall below 1
	if difficulty.Cmp(big1) < 0 {
		difficulty = big1
	}

	return difficulty
}

func CalcPoSDifficultyV1(
	newBlockTime uint64, // TODO: maybe we should be recalculating ema/drift/etc with the new time included?
	diff uint64,
	timeTarget *TimeTarget,
) uint64 {
	res := uint64(0)

	// Find the target anchor
	target := (timeTarget.BlockTarget + timeTarget.PeriodTarget) / 2
	if target < timeTarget.Min {
		target = timeTarget.Min
	}

	if newBlockTime < target {
		targetDelta := target - newBlockTime
		if targetDelta > diffV1_BMax {
			targetDelta = diffV1_BMax
		}

		B := diffV1_BTable[targetDelta]
		res = new(big.Int).Div(new(big.Int).Mul(new(big.Int).SetUint64(diff), B), diffV1_Div).Uint64()

	} else if newBlockTime > target {
		targetDelta := newBlockTime - target
		// clamp the target delta to max
		if targetDelta > diffV1_AMax {
			targetDelta = diffV1_AMax
		}
		A := diffV1_ATable[targetDelta]
		res = new(big.Int).Div(
			new(big.Int).Mul(new(big.Int).SetUint64(diff), diffV1_Div),
			A,
		).Uint64()

	} else {
		return diff
	}

	if new(big.Int).SetUint64(res).Cmp(big1) < 0 {
		res = big1.Uint64()
	}
	return res
}

// CalculateBlockTimeEMA computes the exponential moving average of block times
// this will return the EMA of block times as microseconds
// for a description of the EMA algorithm, please see:
// see htimeTargetps://www.itl.nist.gov/div898/handbook/pmc/section4/pmc431.htm
func CalculateBlockTimeEMA(blockTimeDifferences []uint64, emaPeriod uint64, TargetBlockGap uint64) (ema []uint64) {
	sampleSize := len(blockTimeDifferences)
	N := emaPeriod + 1
	ema = make([]uint64, sampleSize)

	// choice of initial condition is important for an EMA. We could use the first
	// block time difference, but instead we'll set it to the target value so our
	// EMA will tend toward the target. However we don't include this value in our
	// EMA series data that we return, we only use it to calculate the first EMA
	emaPrev := TargetBlockGap * microseconds
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
func CalculateBlockTimeDrift(ema []uint64, TargetBlockGap uint64) (drift []int64) {
	target := int64(TargetBlockGap * microseconds)
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
	for i := 1; i < sampleSize-1; i++ {
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
	derivative = make([]int64, sampleSize-1)

	for i := 1; i < sampleSize; i++ {
		derivative[i-1] = (drift[i] - drift[i-1])
	}
	return
}
