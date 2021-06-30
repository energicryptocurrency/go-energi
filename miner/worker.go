commit 1f31c45c6dd5431d138e6067f7f5f3ebdb952d79
Author: Ryan Lucchese <ryan@energi.team>
Date:   Wed Jun 30 08:55:39 2021 -0600

    removing changes that are not relevant to this branch

diff --git a/miner/worker.go b/miner/worker.go
index 20e260e9a..fc63fd4c8 100644
--- a/miner/worker.go
+++ b/miner/worker.go
@@ -30,13 +30,12 @@ import (
 	"energi.world/core/gen3/core"
 	"energi.world/core/gen3/core/state"
 	"energi.world/core/gen3/core/types"
-	energi_consensus "energi.world/core/gen3/energi/consensus"
 	"energi.world/core/gen3/event"
 	"energi.world/core/gen3/log"
 	"energi.world/core/gen3/params"
 	mapset "github.com/deckarep/golang-set"
-	energi_params "energi.world/core/gen3/energi/params"
 
+	energi_consensus "energi.world/core/gen3/energi/consensus"
 )
 
 const (
@@ -235,7 +234,7 @@ func newWorker(config *params.ChainConfig, engine consensus.Engine, eth Backend,
 		Number:     num.Add(num, common.Big1),
 		GasLimit:   core.CalcGasLimit(parent, worker.gasFloor, worker.gasCeil),
 		Extra:      worker.extra,
-		Time:       uint64(parent.Time() + energi_params.MinBlockGap),
+		Time:       uint64(parent.Time() + energi_consensus.MinBlockGap),
 	}
 	if err := worker.makeCurrent(parent, header); err != nil {
 		panic(err)
@@ -894,7 +893,7 @@ func (w *worker) commitNewWork(interrupt *int32, noempty bool, timestamp int64)
 		Number:     num.Add(num, common.Big1),
 		GasLimit:   core.CalcGasLimit(parent, w.gasFloor, w.gasCeil),
 		Extra:      w.extra,
-		Time:       uint64(parent.Time() + energi_params.MinBlockGap),
+		Time:       uint64(parent.Time() + energi_consensus.MinBlockGap),
 	}
 	// Only set the coinbase if our consensus engine is running (avoid spurious block rewards)
 	if w.isRunning() {
