package hfcache

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddHardfork(t *testing.T) {
	t.Parallel()
	AddHardfork(&Hardfork{Name: "Asgard", BlockNumber: new(big.Int).SetUint64(2)})
	AddHardfork(&Hardfork{Name: "Asgard", BlockNumber: new(big.Int).SetUint64(3)})
	assert.Equal(t, IsHardforkActive("Asgard", 4), true)
}

func TestRemoveHardfork(t *testing.T) {
	t.Parallel()
	AddHardfork(&Hardfork{Name: "Asgard", BlockNumber: new(big.Int).SetUint64(2)})
	AddHardfork(&Hardfork{Name: "Banana-txfee", BlockNumber: new(big.Int).SetUint64(3)})

	RemoveHardfork("Banana-txfee")
	assert.Equal(t, IsHardforkActive("Asgard", 4), true)
	assert.Equal(t, IsHardforkActive("Banana-txfee", 4), false)
}
