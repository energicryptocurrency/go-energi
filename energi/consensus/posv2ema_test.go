package consensus

import (
	"testing"
)

func TestCalculateBlockTimeEMA(t *testing.T) {
	emaCalculated := CalculateBlockTimeEMA(emaSamples)
	emaExpected := uint64(59161280)
	if emaCalculated != emaExpected {
		t.Log("EMA mismatch - expected", emaExpected, "got", emaCalculated)
		t.FailNow()
	}
}
