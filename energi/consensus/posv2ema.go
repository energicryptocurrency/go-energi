package consensus

import (
	"energi.world/core/gen3/energi/params"
)

// CalculateBlockTimeEMA computes the exponential moving average of block times
// this will return the EMA of block times as microseconds
// for a description of the EMA algorithm, please see:
// see https://www.itl.nist.gov/div898/handbook/pmc/section4/pmc431.htm
func CalculateBlockTimeEMA(blockTimeDifferences []uint64) (ema uint64) {
	sampleSize := len(blockTimeDifferences)

	// we use a scaling factor due to entirely integer calculation for this function
	// scaling up lets us calculate an EMA at a higher resolution that 1 second
	scalingFactor := 1000000 // after scaling the units will be microseconds

	// choice of initial condition is important for an EMA. We could use the first
	// block time difference, but instead we'll set it to the target value so our
	// EMA will tend toward the target
	ema = params.TargetBlockGap * scalingFactor
	for i := 1; i < sampleSize; i++ {
		blockTimeDifferences[i-1] *= scalingFactor
		// this formula has a factor of 2/(N+1) in a couple places. This is our
		// smoothing coefficient for the EMA, often referred to as alpha. We have
		// not precomputed this value so we don't lose precision on early division
		ema = ((2 * blockTimeDifferences[i-1])/(sampleSize + 1)) + (ema - ((ema * 2)/(sampleSize + 1))
	}
	return
}
