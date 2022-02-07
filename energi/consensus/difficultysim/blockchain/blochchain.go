package blockchain

import (
	"sync"
	"time"

	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/calculus"
	"github.com/energicryptocurrency/energi/energi/consensus/difficultysim/params"
)

// Block represents each block in the blockchain
type Block struct {
	Difficulty uint64 `json:"difficultyStart"`
	Height     uint64 `json:"height"`
	Coinbase   string `json:"coinbase"`
	Time       uint64 `json:"time"`
	Nonce      uint64 `json:"nonce"`
}

// Blockchain represents current state of the chain
type Blockchain struct {
	Locker *sync.Mutex
	Chain  []Block

	// time function
	Now func() uint64
}

// CreateBlockchain creates initial blockchain state
func CreateBlockchain() *Blockchain {
	// create blockchain parameters
	chain := Blockchain{
		Now: func() uint64 {
			return uint64(time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)))
		},
		Locker: &sync.Mutex{},
	}

	// insert genesis block
	chain.Chain = []Block{{Difficulty: params.InitialDifficulty, Height: 1,
		Coinbase: "", Nonce: 0, Time: chain.Now()}}

	return &chain
}

// GetBlock return height corresponding mined block
func (chain *Blockchain) GetBlock(height uint64) *Block {
	chain.Locker.Lock()
	defer chain.Locker.Unlock()
	if len(chain.Chain) < int(height) || height <= 0 {
		return nil
	}
	return &chain.Chain[height-1]
}

// LastBlock returns last mined block in chain (the blockchain has genesis block)
func (chain *Blockchain) LastBlock() *Block {
	chain.Locker.Lock()
	defer chain.Locker.Unlock()
	return &chain.Chain[len(chain.Chain)-1]
}

// NewBlock appends new mined block into the chain
func (chain *Blockchain) NewBlock(block *Block) {
	chain.Locker.Lock()
	defer chain.Locker.Unlock()
	if chain.Chain[len(chain.Chain)-1].Height+1 != block.Height {
		return
	}
	chain.Chain = append(chain.Chain, *block)

}

/**
 * Implements block time consensus
 *
 * POS-11: Block time restrictions
 * POS-12: Block interval enforcement
 */
func (chain *Blockchain) CalcTimeTargetV1(parent *Block) *calculus.TimeTarget {
	ret := new(calculus.TimeTarget)
	now := chain.Now()
	parentNumber := parent.Height
	blockNumber := parentNumber + 1

	// POS-11: Block time restrictions
	ret.Max = now + params.MaxFutureGap

	// POS-11: Block time restrictions
	ret.Min = parent.Time + params.MinBlockGap
	ret.BlockTarget = parent.Time + params.TargetBlockGap
	ret.PeriodTarget = ret.BlockTarget

	// POS-12: Block interval enforcement
	// ---
	if blockNumber >= params.AveragingWindow {
		// TODO: LRU cache here for extra DoS mitigation
		past := parent

		// NOTE: we have to do this way as parent may be not part of canonical
		//       chain. As no mutex is held, we cannot do checks for canonical.
		for i := params.AveragingWindow - 1; i > 0; i-- {
			past = chain.GetBlock(past.Height - 1)
			if past == nil {
				return ret
			}
		}

		ret.PeriodTarget = past.Time + params.TargetPeriodGap
		periodMinTime := ret.PeriodTarget - params.MinBlockGap
		if periodMinTime > ret.Min {
			ret.Min = periodMinTime
		}
	}

	return ret
}

// CalcTimeTargetV2 calculates asgard time target
func (chain *Blockchain) CalcTimeTargetV2(parent *Block) *calculus.TimeTarget {
	ret := &calculus.TimeTarget{}
	parentBlockTime := parent.Time // Defines the original parent block time.

	// POS-11: Block time restrictions
	ret.Max = chain.Now() + params.MaxFutureGap

	// POS-11: Block time restrictions
	ret.Min = parentBlockTime + params.MinBlockGap
	ret.BlockTarget = parentBlockTime + params.TargetBlockGap
	ret.PeriodTarget = ret.BlockTarget

	// Block interval enforcement
	// TODO: LRU cache here for extra DoS mitigation
	timeDiffs := make([]uint64, params.BlockTimeEMAPeriod)

	// compute block time differences
	// note that the most recent time difference will be the most
	// weighted by the EMA, and the oldest time difference will be the least
	for i := params.BlockTimeEMAPeriod; i > 0; i-- {
		past := chain.GetBlock(parent.Height - 1)
		if past == nil {
			break
		}
		timeDiffs[i-1] = parent.Time - past.Time
		parent = past
	}

	ema := calculus.CalculateBlockTimeEMA(timeDiffs, params.BlockTimeEMAPeriod, params.TargetBlockGap)

	ret.PeriodTarget = ema[len(ema)-1]

	// set up the parameters for PID control (diffV2)
	drift := calculus.CalculateBlockTimeDrift(ema, params.TargetBlockGap)
	integral := calculus.CalculateBlockTimeIntegral(drift)
	derivative := calculus.CalculateBlockTimeDerivative(drift)
	ret.Drift = drift[len(drift)-1]
	ret.Integral = integral
	ret.Derivative = derivative[len(derivative)-1]

	return ret
}
