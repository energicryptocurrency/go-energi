package miner

import "energi.world/core/gen3/metrics"

var (
	difficultyGauge  = metrics.NewRegisteredGauge("energi/consensus/difficulty", nil)
	blocksizeGauge   = metrics.NewRegisteredGauge("energi/consensus/blocksize", nil)
	transactionGauge = metrics.NewRegisteredGauge("energi/consensus/transactions", nil)
	blocktimeGauge   = metrics.NewRegisteredGauge("energi/consensus/blocktime", nil)
)
