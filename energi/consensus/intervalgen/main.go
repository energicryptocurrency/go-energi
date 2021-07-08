package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

func main() {
	const  sampleNum =  60
	// This will generate a consistent set of random values
	rand.Seed(32)
	// // This will generate an always changing set of random values
	// rand.Seed(time.Now().UnixNano())
	// create 60 samples as per needed for the interval test
	samples := make([]uint64, sampleNum)
	samples[0] = uint64(30 + rand.Int63n(sampleNum))
	for i := range samples {
		if i > 0 {
			samples[i] = samples[i-1] + uint64(30+rand.Int63n(
				60))
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
