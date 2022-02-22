package stats

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/blockchain"
	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/miner"
)

type AccountStat struct {
	Address      string
	WonBlocks    uint64
	AverageNonce uint64
}

// AccountStats collects account mining results
func PrintAccountStats(chain *blockchain.Blockchain, stakers []miner.Staker) {
	chain.Locker.Lock()
	defer chain.Locker.Unlock()

	// associating address to owner
	addressOwner := make(map[string]string)
	for _, staker := range stakers {
		for _, account := range staker.Accounts {
			addressOwner[account.Address] = staker.Name
		}
	}

	// collect won blocks and average nonce used
	// initial csv file fields
	difficultyData := [][]string{
		{"Block", "Difficulty", "Coinbase", "Nonce", "BlockTime"},
	}
	accountStats := make(map[string]*AccountStat)
	totalTime := uint64(0)
	for i, block := range chain.Chain {
		var timeDifference uint64
		if i != 0 {
			timeDifference = block.Time - chain.Chain[i-1].Time
			totalTime += timeDifference
		}
		difficultyData = append(difficultyData, []string{strconv.FormatUint(block.Height, 10), strconv.FormatUint(block.Difficulty, 10), block.Coinbase, strconv.FormatUint(block.Nonce, 10), strconv.FormatUint(timeDifference, 10)})
		if accountStat, ok := accountStats[block.Coinbase]; ok {
			accountStat.WonBlocks++
			accountStat.AverageNonce += block.Nonce
		} else {
			accountStats[block.Coinbase] = &AccountStat{
				WonBlocks:    1,
				AverageNonce: block.Nonce,
			}
		}
	}
	//print average block creation time
	fmt.Println("average block creation time ", totalTime/(uint64(len(chain.Chain))-1))
	for _, v := range accountStats {
		v.AverageNonce = v.AverageNonce / v.WonBlocks
	}

	for k := 0; k <= len(stakers)-2; k++ {
		for j := k + 1; j <= len(stakers)-1; j++ {
			if stakers[k].Accounts[0].Balance < stakers[j].Accounts[0].Balance {
				s := stakers[k]
				stakers[k] = stakers[j]
				stakers[j] = s
			}
		}
	}

	// initial csv file fields
	stakingData := [][]string{
		{"Name", "Address", "Balance", "Blocks", "NonceCap", "AverageWeight"},
	}

	// run miners and start mining of the blockchain
	for _, staker := range stakers {
		for _, account := range staker.Accounts {
			wonBlock := uint64(0)
			averageNonce := uint64(0)
			if _, ok := accountStats[account.Address]; ok {
				wonBlock = accountStats[account.Address].WonBlocks
				averageNonce = accountStats[account.Address].AverageNonce
			}
			stakingData = append(stakingData, []string{staker.Name, account.Address, strconv.FormatUint(account.Balance, 10), strconv.FormatUint(wonBlock, 10), strconv.FormatUint(account.NonceCap, 10), strconv.FormatUint(averageNonce, 10)})
		}
	}

	// generate csv with the data
	createCSV(stakingData, "stakindData")
	createCSV(difficultyData, "difficultyAdjustment")

}

/*
   getFileTime simply takes current time precisely and
   creates string date-time prefix for each file
*/
func getFileTime() string {
	var logTime string
	now := time.Now()
	year, month, day := now.Date()
	logTime += strconv.Itoa(year) + " "
	logTime += month.String() + " "
	logTime += strconv.Itoa(day) + " "
	hour, min, _ := now.Clock()
	logTime += strconv.Itoa(hour) + ":"
	logTime += strconv.Itoa(min) + "."
	return logTime
}

func createCSV(output [][]string, name string) {
	if _, err := os.Stat("csv/"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("csv/", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	csvFile, err := os.Create("csv/" + strings.ReplaceAll(getFileTime()+name, " ", "-") + ".csv")

	if err != nil {
		fmt.Println("failed creating file:", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range output {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}

// PrintStats prints various stats of the blockchain after specific time
func PrintStats(chain *blockchain.Blockchain, stakers []miner.Staker) {
	// print account mining results
	PrintAccountStats(chain, stakers)
}
