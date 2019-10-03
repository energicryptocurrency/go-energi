package common

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var dummyBlockHash = common.BytesToHash([]byte{1, 3, 4, 5})

type fakeChain struct{}

func (f *fakeChain) CurrentBlock() *types.Block {
	return types.NewBlock(&types.Header{
		ParentHash: dummyBlockHash,
	}, nil, nil, nil)
}

// TestDataCache tests the cache's setter and getter methods.
func TestDataCache(t *testing.T) {
	chain := new(fakeChain)
	cacheInstance := new(CacheStorage)

	var newData interface{}
	cacheQueryfunc := func(blockhash common.Hash) (interface{}, error) {
		return newData, nil
	}

	t.Run("Test_adding_new_data_nil_data", func(t *testing.T) {
		data, err := cacheInstance.Get(chain, cacheQueryfunc)
		if err != ErrInvalidData {
			t.Fatalf("expected error (%v) but found (%v)", ErrInvalidData, err)
		}

		if !reflect.DeepEqual(data, newData) {
			t.Fatalf("expected the returned data to match but it didn't")
		}
	})

	t.Run("Test_adding_cache_data", func(t *testing.T) {
		newData = []int{1, 2, 3, 5, 6, 6}
		data, err := cacheInstance.Get(chain, cacheQueryfunc)
		if err != nil {
			t.Fatalf("expected no error but found %v", err)
		}

		if !reflect.DeepEqual(data, newData) {
			t.Fatalf("expected the returned data to match but it didn't")
		}
	})

	t.Run("Test_adding_new_cache_data_but_with_old_hash", func(t *testing.T) {
		newData = "This is some random new data"
		data, err := cacheInstance.Get(chain, cacheQueryfunc)
		if err != nil {
			t.Fatalf("expected no error but found %v", err)
		}

		if reflect.DeepEqual(data, newData) {
			t.Fatalf("expected the returned data not to match but it did")
		}
	})

	t.Run("Test_adding_new_cache_data_but_with_new_hash", func(t *testing.T) {
		dummyBlockHash = common.BytesToHash([]byte{120, 23, 90, 5})
		newData = map[string]int{
			"rsc": 3711,
			"r":   2138,
			"gri": 1908,
			"adg": 912,
		}

		data, err := cacheInstance.Get(chain, cacheQueryfunc)
		if err != nil {
			t.Fatalf("expected no error but found %v", err)
		}

		if !reflect.DeepEqual(data, newData) {
			t.Fatalf("expected the returned data to match it didn't")
		}
	})
}
