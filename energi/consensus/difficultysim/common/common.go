package common

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/blockchain"
)

const (
	DAY   = 86400
	WEEK  = DAY * 7
	MONTH = DAY * 30 //approximately 716 hours
)

// Set flags (command line args)
var stakeConfig = flag.String(
	"stakeConfig",
	"",
	"The filename of the JSON file that contains the simulation configuration to use",
)

func GetConfigFile() []byte {
	// Parse the arguments
	flag.Parse()

	// Sanity check the arguments
	stakers := *stakeConfig
	if stakers == "" {
		panic("The simulation requires a stake configuration file")
	} else {
		return loadJSON(*stakeConfig)
	}

}

// loadJson Opens our jsonFile
func loadJSON(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

// prints how much time passed since genesis
func PrintPassedTime(chain *blockchain.Blockchain) {
	initialTime := chain.Now()
	for {
		time.Sleep(20 * time.Second)
		fmt.Println("Passed virtual time - ", time.Duration(chain.Now()-initialTime)*time.Second)
	}
}
