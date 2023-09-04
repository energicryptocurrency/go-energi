package consensus

import "github.com/energicryptocurrency/go-energi/metrics"

var (
	mnGauge = metrics.NewRegisteredGauge("energi/consensus/masternodes", nil)
)
