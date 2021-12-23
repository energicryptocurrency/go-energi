package miner

import "github.com/energicryptocurrency/energi/metrics"

var (
	difficultyGauge  = metrics.NewRegisteredGauge("energi/consensus/difficulty", nil)
	blocksizeGauge   = metrics.NewRegisteredGauge("energi/consensus/blocksize", nil)
	transactionGauge = metrics.NewRegisteredGauge("energi/consensus/transactions", nil)
	blocktimeGauge   = metrics.NewRegisteredGauge("energi/consensus/blocktime", nil)
)
