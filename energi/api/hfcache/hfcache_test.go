package hfcache

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "math/big"
)

func TestAddHardfork(t *testing.T) {
  t.Parallel()
  AddHardfork(&Hardfork{Name: "Asgard", BlockNumber: new(big.Int).SetUint64(2)})
  AddHardfork(&Hardfork{Name: "Asgard", BlockNumber: new(big.Int).SetUint64(3)})
  assert.Equal(t, IsHardforkActive("Asgard",4),true)
}


func TestRemoveHardfork(t *testing.T) {
  t.Parallel()
  AddHardfork(&Hardfork{Name: "Asgard", BlockNumber: new(big.Int).SetUint64(2)})
  AddHardfork(&Hardfork{Name: "Banana", BlockNumber: new(big.Int).SetUint64(3)})

  RemoveHardfork("Banana")
  assert.Equal(t, IsHardforkActive("Asgard",4),true)
  assert.Equal(t, IsHardforkActive("Banana",4),false)
}
