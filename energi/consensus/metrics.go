package consensus

import "github.com/energicryptocurrency/energi/metrics"

var (
	mnGauge = metrics.NewRegisteredGauge("energi/consensus/masternodes", nil)
)
