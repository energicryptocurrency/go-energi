package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

func main() {
	const minute, sampleNum = 60, 60
	// This will generate a consistent set of random values
	rand.Seed(69)
	// // This will generate an always changing set of random values
	// rand.Seed(time.Now().UnixNano())
	// create 60 samples as per needed for the interval test
	samples := make([]uint64, sampleNum)
	samples[0] = uint64(30 + rand.Int63n(minute*2/3))
	for i := range samples {
		if i > 0 {
			// allow up to 10 minutes long between blocks, we always
			// add something to the previous block time creating the
			// consensus monotonic timestamps with minimum 30 second
			// difference
			samples[i] = samples[i-1] + uint64(30+rand.Int63n(
				minute*2/3))
		}
	}
	output := `package consensus
	
`+`//go:generate go run ./intervalgen/.

var emaSamples = []uint64{
`
	for i := range samples {
		output += fmt.Sprint("\t",samples[i], ",\n")
	}
	output+="}\n"
	ioutil.WriteFile("posv2emasamples.go", []byte(output), 0660)
}
