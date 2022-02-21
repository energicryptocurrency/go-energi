package hfcache

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "math/big"
)

func TestAddHardfork(t *testing.T) {
	t.Parallel()

  var name [32]byte
	copy(name[:], string("Asgard")[:])

  AddHardfork(&Hardfork{Name: string(name[:]), BlockNumber: new(big.Int).SetUint64(2)})
  AddHardfork(&Hardfork{Name: string(name[:]), BlockNumber: new(big.Int).SetUint64(3)})
  assert.Equal(t, IsHardforkActive("Asgard",4),true)
}


func TestRemoveHardfork(t *testing.T) {
	t.Parallel()

  var asgard [32]byte
	copy(asgard[:], string("Asgard")[:])

  var banana [32]byte
	copy(banana[:], string("Banana")[:])

  AddHardfork(&Hardfork{Name: string(asgard[:]), BlockNumber: new(big.Int).SetUint64(2)})
  AddHardfork(&Hardfork{Name: string(banana[:]), BlockNumber: new(big.Int).SetUint64(3)})

  RemoveHardfork(banana)
  assert.Equal(t, IsHardforkActive("Asgard",4),true)
  assert.Equal(t, IsHardforkActive("Banana",4),false)
}
