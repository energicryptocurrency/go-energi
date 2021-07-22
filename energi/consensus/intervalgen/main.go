package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"energi.world/core/gen3/energi/params"
	"energi.world/core/gen3/energi/consensus"
)

const (
	sampleNum = 360
	averagingWindow = 60
)

func addBlockTimes(output *string) (samples []uint64) {
	// generate random block times that average about 60 seconds
	rand.Seed(32)
	samples = make([]uint64, sampleNum)
	for i := range samples {
		samples[i] = uint64(int64(params.MinBlockGap)+rand.Int63n(int64(params.TargetBlockGap)))
	}

	*output += "\nvar testDataBlockTimes = []uint64{\n  "
	for i := range samples {
		*output += fmt.Sprint(samples[i])
		if (i+1) % 30 == 0 {
			*output += ",\n  "
		} else {
			*output += ","
		}
	}
	*output += "}\n"
	return
}

func addBlockTimeEMA(samples []uint64, output *string) (ema []uint64) {
	ema = consensus.CalculateBlockTimeEMA(samples, averagingWindow)
	*output += "\nvar testDataBlockTimeEMA = []uint64{\n  "
	for i := range ema {
		*output += fmt.Sprint(ema[i])
		if (i+1) % 10 == 0 {
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
		if (i+1) % 10 == 0 {
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
	_ = integral
	ioutil.WriteFile("posv2emasamples_test.go", []byte(output), 0660)
}
