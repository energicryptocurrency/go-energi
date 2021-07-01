package consensus

import (
	"testing"

	"energi.world/core/gen3/energi/params"
)

func TestCalcEMAUint64(t *testing.T) {
	o := CalcEMAUint64(emaSamples, 2, params.SMAPeriod+1, params.SMAPeriod)
	if o != 142 {
		t.Log("EMA calculation did not produce expected result")
		t.Log("expected 142, got", o)
		t.FailNow()
	}
}
