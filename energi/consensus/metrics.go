package consensus

import "energi.world/core/gen3/metrics"

var (
	mnGauge = metrics.NewRegisteredGauge("energi/consensus/masternodes", nil)
)
