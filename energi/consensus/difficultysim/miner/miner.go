package miner

import (
	"math/big"
	"time"

	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/blockchain"
	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/calculus"
	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/params"
)

// Staker is individual person mining with accoount(s)
type Staker struct {
	Name     string    `json:"name"`
	Accounts []Account `json:"accounts"`
}

// Account represents individual account of a particular person mining
type Account struct {
	Address  string `json:"address"`
	Balance  uint64 `json:"balance"`
	NonceCap uint64 `json:"nonceCap"`
}

var (
	// constant
	diff1Target = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))
)

// Mine starts separate mining routines for each account
func (staker *Staker) Mine(chain *blockchain.Blockchain) {
	// start separate mining for each account of the staker
	for _, account := range staker.Accounts {
		go account.mine(chain)
	}
}

// mine function simulates real mining with it's difficulty calculation and pos hash generation based winning
func (account *Account) mine(chain *blockchain.Blockchain) {
	for {
		// last mined block of the blockchain
		parent := chain.LastBlock()

		// make time target calculation depending on asgard status
		var timeTarget *calculus.TimeTarget
		if params.AsgardIsActive {
			timeTarget = chain.CalcTimeTargetV2(parent)
		} else {
			timeTarget = chain.CalcTimeTargetV1(parent)
		}

		// new potential block
		block := &blockchain.Block{
			Coinbase: account.Address,
			Time:     timeTarget.Min,
			// NOTE: currently first parameter TIME isn't used in difficulty calculation
			Difficulty: calculus.CalcPoSDifficultyV2(timeTarget.Min, parent.Difficulty, timeTarget).Uint64(),
			Height:     parent.Height + 1,
		}

		//mining loop
		for ; ; block.Time++ {
			// if asgard is not active we need to calculate difficulty with V1 algorithm
			if !params.AsgardIsActive {
				block.Difficulty = calculus.CalcPoSDifficultyV1(block.Time, parent.Difficulty, timeTarget)
			}

			if maxTime := chain.Now() + params.MaxFutureGap; block.Time > maxTime {
				<-time.After(time.Duration(block.Time-maxTime) * time.Millisecond)
			}

			// check if chain changed
			if chain.LastBlock().Height != parent.Height {
				break
			}

			// check available stake weight
			availableWeight := account.LookupStakeWeight(block.Time, parent, account.NonceCap, chain)

			// check if we cal mine a block with given available weight for this specific second(millisecond)
			block.Nonce = calculus.CalcPoSHash(block.Time, block.Coinbase, new(big.Int).Div(diff1Target, new(big.Int).SetUint64(block.Difficulty)))

			// check if we can mine a block
			if block.Nonce <= availableWeight {
				chain.NewBlock(block)
				break
			}
		}
	}
}

/**
 * Implements stake amount calculation.
 *
 * POS-3: Stake maturity period
 * POS-4: Stake amount
 * POS-22: Partial stake amount
 *
 * This is a basic helper for stake amount calculation.
 * There are ways to optimize it for high load, but we need something
 * to start with.
 */
func (account *Account) LookupStakeWeight(
	now uint64,
	until *blockchain.Block,
	nonceCap uint64,
	chain *blockchain.Blockchain,
) (weight uint64) {
	// NOTE: Do not set to high initial value due to defensive coding approach!
	totalStaked := uint64(0)
	var maturityPeriod uint64
	if params.AsgardIsActive {
		maturityPeriod = params.MaturityPeriodAsgard
	} else {
		maturityPeriod = params.MaturityPeriod
	}
	// NOTE: we need to ensure at least one iteration with the balance condition
	for until.Time > now-maturityPeriod {

		// POS-22: partial stake amount
		if until.Coinbase == account.Address {
			totalStaked += until.Nonce
		}

		// check the genesis block
		if until.Height == 1 {
			break
		}
		until = chain.GetBlock(until.Height - 1)
	}

	if nonceCap != 0 && nonceCap < account.Balance-totalStaked {
		return nonceCap
	}
	return account.Balance - totalStaked
}
