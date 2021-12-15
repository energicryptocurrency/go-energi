package main

import (
	"encoding/json"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/energi/consensus"
	"github.com/energicryptocurrency/energi/energi/params"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
)

const (
	sampleNum = 360
)

type PoSDiffV2TestCase struct {
	Time       uint64
	Parent     int64
	Drift      int64
	Integral   int64
	Derivative int64
	Result     int64
}

func addBlockTimes(output *string) (samples []uint64) {
	// generate random block times that average about 60 seconds
	rand.Seed(32)
	samples = make([]uint64, sampleNum)
	for i := range samples {
		samples[i] = uint64(int64(params.MinBlockGap) + rand.Int63n(int64(params.TargetBlockGap)))
	}

	*output += "\nvar testDataBlockTimes = []uint64{\n  "
	for i := range samples {
		*output += fmt.Sprint(samples[i])
		if (i+1)%30 == 0 {
			*output += ",\n  "
		} else {
			*output += ","
		}
	}
	*output += "}\n"
	return
}

func addBlockTimeEMA(samples []uint64, output *string) (ema []uint64) {
	ema = consensus.CalculateBlockTimeEMA(samples, params.BlockTimeEMAPeriod)
	*output += "\nvar testDataBlockTimeEMA = []uint64{\n  "
	for i := range ema {
		*output += fmt.Sprint(ema[i])
		if (i+1)%10 == 0 {
			*output += ",\n  "
		} else {
			*output += ","
		}
	}
	*output += "}\n"
	return
}

func addBlockTimeDrift(ema []uint64, output *string) (drift []int64) {
	drift = consensus.CalculateBlockTimeDrift(ema)
	*output += "\nvar testDataBlockTimeDrift = []int64{\n  "
	for i := range drift {
		*output += fmt.Sprint(drift[i])
		if (i+1)%10 == 0 {
			*output += ",\n  "
		} else {
			*output += ","
		}
	}
	*output += "}\n"
	return
}

func addBlockTimeIntegral(drift []int64, output *string) (integral int64) {
	integral = consensus.CalculateBlockTimeIntegral(drift)
	*output += fmt.Sprint("\nvar testDataBlockTimeIntegral int64 = ", integral, "\n")
	return
}

func addBlockTimeDerivative(drift []int64, output *string) (derivative []int64) {
	derivative = consensus.CalculateBlockTimeDerivative(drift)
	*output += "\nvar testDataBlockTimeDerivative = []int64{\n  "
	for i := range derivative {
		*output += fmt.Sprint(derivative[i])
		if (i+1)%10 == 0 {
			*output += ",\n  "
		} else {
			*output += ","
		}
	}
	*output += "}\n"
	return
}

func simulateStaking(
	blockTimesInitial []uint64,
	ema []uint64,
	drift []int64,
	integral int64,
	derivative []int64,
	output *string,
) {
	const (
		initialDifficulty    int64  = 343768608    // mainnet difficulty number
		simulationBlockCount        = 60 * 24 * 21 // 21 days
		maxStakeTime         uint64 = 10000
	)

	var result []PoSDiffV2TestCase

	nrgStaking := int64(5500000)
	simulationStartBlock := len(blockTimesInitial)
	totalBlocks := simulationStartBlock + simulationBlockCount

	blockTimes := make([]uint64, totalBlocks)
	difficulty := make([]*big.Int, totalBlocks)

	for i := 0; i < simulationStartBlock; i++ {
		blockTimes[i] = blockTimesInitial[i]
		difficulty[i] = big.NewInt(initialDifficulty)
	}

	// initialDifficulty is just used for some deterministic number to seed the rand source
	s := rand.NewSource(initialDifficulty)
	r := rand.New(s)

	// P(blockFound) is nrgStaking / difficulty
	blockFound := func(r *rand.Rand, diff *big.Int) bool {
		return r.Int63n(diff.Int64()) <= nrgStaking
	}

	for blockCount := simulationStartBlock; blockCount < totalBlocks; blockCount++ {
		// double the amount of NRG at stake at block 10000
		if blockCount == 10000 {
			nrgStaking *= 2
		}

		// cut the amount of NRG at stake to 1/4 at block 20000
		if blockCount == 20000 {
			nrgStaking /= 4
		}

		// simulated mining - starting from 30 seconds to some timeout value
		for t := uint64(30); t < maxStakeTime; t++ {
			// the odds of finding a block with any timestamp are nrgAtStake/difficulty
			if blockFound(r, difficulty[blockCount-1]) {
				// rather than initialize a whole engine let's just build a time target
				timeTarget := &consensus.TimeTarget{}
				timeSlice := blockTimes[blockCount-61 : blockCount-1]
				ema := consensus.CalculateBlockTimeEMA(timeSlice, params.BlockTimeEMAPeriod)
				drift := consensus.CalculateBlockTimeDrift(ema)
				integral := consensus.CalculateBlockTimeIntegral(drift)
				derivative := consensus.CalculateBlockTimeDerivative(drift)
				timeTarget.Drift = drift[len(drift)-1]
				timeTarget.Integral = integral
				timeTarget.Derivative = derivative[len(derivative)-1]

				// create a header with the previous difficulty
				parentHeader := &types.Header{}
				parentHeader.Difficulty = difficulty[blockCount-1]

				// calculate the difficulty
				blockTimes[blockCount] = t
				difficulty[blockCount] = consensus.CalcPoSDifficultyV2(t, parentHeader, timeTarget)
				result = append(result, PoSDiffV2TestCase{
					Time:       t,
					Parent:     difficulty[blockCount-1].Int64(),
					Drift:      timeTarget.Drift,
					Integral:   timeTarget.Integral,
					Derivative: timeTarget.Derivative,
					Result:     difficulty[blockCount].Int64(),
				})
				break
			}
			// timeout - prevent infinite mining loop if simulation difficulty gets too high
			if t+1 == maxStakeTime {
				panic("[ERROR]: Stake timeout, something wrong?")
			}
		}
	}
	emaTimes := consensus.CalculateBlockTimeEMA(blockTimes, params.BlockTimeEMAPeriod)

	// write the final sampuleNum (block time, difficulty) pairs for unit tests
	if len(blockTimes) != len(difficulty) {
		panic("[ERROR]: inconsistent difficulty data!")
	}
	if len(blockTimes) != len(emaTimes) {
		panic("[ERROR]: inconsistent EMA data")
	}

	// write simulated data to CSV
	csvData := "time,emaTime,difficulty\n"
	for i := range blockTimes {
		csvData += fmt.Sprint(blockTimes[i], ",", emaTimes[i], ",", difficulty[i].Uint64(), "\n")
	}
	ioutil.WriteFile("staking_simulation.csv", []byte(csvData), 0660)
	jsonCases, _ := json.Marshal(result)
	ioutil.WriteFile("energi/consensus/intervalgen/PoSV2_test_cases.json", jsonCases, 0660)
}

func main() {
	output := `// Copyright 2021 The Energi Core Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

// IMPORTANT: this file is code generated, DO NOT EDIT

package consensus

//go:generate go run ./intervalgen/.
`
	samples := addBlockTimes(&output)
	ema := addBlockTimeEMA(samples, &output)
	drift := addBlockTimeDrift(ema, &output)
	integral := addBlockTimeIntegral(drift, &output)
	derivative := addBlockTimeDerivative(drift, &output)
	simulateStaking(samples, ema, drift, integral, derivative, &output)
	ioutil.WriteFile("posv2emasamples_test.go", []byte(output), 0660)
}
