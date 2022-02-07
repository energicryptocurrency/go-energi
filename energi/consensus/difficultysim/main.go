// Copyright 2022 The Energi Core Authors
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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/blockchain"
	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/common"
	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/miner"
	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/stats"
)

// SimConfig configuration of the simulation
type SimConfig struct {
	Stakers []miner.Staker `json:"stakers"`
}

// SimulationPeriod is the virtual period simulation needs to run
var SimulationPeriod = flag.String(
	"simulationPeriod",
	"day",
	"The virtual time period for running blockchain (day, week or month)",
)

// Start here
func main() {
	// validate passed arguments
	var config SimConfig
	inputBytes := common.GetConfigFile()
	err := json.Unmarshal(inputBytes, &config)
	if err != nil {
		fmt.Println(err)
	}

	// create blockchain parameters
	blockchain := blockchain.CreateBlockchain()

	// run miners and start mining of the blockchain
	for _, staker := range config.Stakers {
		s := staker
		go (&s).Mine(blockchain)
	}

	// periodically prints passed time in simulated world
	go common.PrintPassedTime(blockchain)

	// wait for the specific period
	switch *SimulationPeriod {
	case "day":
		fmt.Println("collecting a day statistics")
		time.Sleep(time.Duration(common.DAY) * time.Millisecond)
	case "week":
		fmt.Println("collecting a week statistics")
		time.Sleep(time.Duration(common.WEEK) * time.Millisecond)
	case "month":
		fmt.Println("collecting a month statistics")
		time.Sleep(time.Duration(common.MONTH) * time.Millisecond)
	default:
		fmt.Println("collecting a day statistics")
		time.Sleep(time.Duration(common.DAY) * time.Millisecond)
	}

	// collect and print statistics
	stats.PrintStats(blockchain, config.Stakers)

}
