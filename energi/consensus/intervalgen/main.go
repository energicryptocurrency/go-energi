package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"energi.world/core/gen3/energi/params"
	"energi.world/core/gen3/energi/consensus"
)

func main() {
	const sampleNum = 1440
	const averagingWindow uint64 = 60
	// This will generate a consistent set of random values
	rand.Seed(32)
	// // This will generate an always changing set of random values
	// rand.Seed(time.Now().UnixNano()) which will have a typical average
	// around the same as params.TargetBlockGap (60 second block interval)
	samples := make([]uint64, sampleNum)
	for i := range samples {
		samples[i] = uint64(int64(params.MinBlockGap)+
				rand.Int63n(int64(params.TargetBlockGap)))
	}

	// calculate the EMA from samples
	samplesEMA := consensus.CalculateBlockTimeEMA(samples, averagingWindow)
	output := `package consensus
	
//go:generate go run ./intervalgen/.

var emaSamples = []uint64{
`
	for i := range samples {
		output += fmt.Sprint("\t", samples[i], ",\n")
	}
	output += `}

var emaSamplesExpected = []uint64{
`
	for i := range samplesEMA {
		output += fmt.Sprint("\t", samplesEMA[i], ",\n")
	}
	output += "}\n"
	ioutil.WriteFile("posv2emasamples_test.go", []byte(output), 0660)
}
