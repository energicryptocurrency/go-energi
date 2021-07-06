package diffv1

import (
	"math/big"

	"energi.world/core/gen3/common"
)

/**
 * POS-13: Difficulty algorithm (Proposal v1)
 */
const (
	ReduceMax             uint64 = 30
	IncreaseMax           uint64 = 120
	DivPlain              uint64 = 100
	MigrationStakerDelay  uint64 = 15
	MigrationStakerTarget uint64 = 0xFFFF
)

var (
	ReduceTable   []*big.Int
	IncreaseTable []*big.Int
	Div           = new(big.Int).SetUint64(DivPlain)
	Target        = new(big.Int).Exp(
		big.NewInt(2), big.NewInt(256), big.NewInt(0),
	)
)

// InitDiffTable generates the difficulty adjustment tables
func InitDiffTable(l uint64, c float64) []*big.Int {
	t := make([]*big.Int, l+1)
	t[0] = common.Big1
	var acc float64 = 1
	for i := 1; i < len(t); i++ {
		acc *= c
		t[i] = big.NewInt(int64(acc * float64(DivPlain)))
	}
	return t
}

func init() {
	ReduceTable = InitDiffTable(ReduceMax, 1.1)
	IncreaseTable = InitDiffTable(IncreaseMax, 1.05)
}
