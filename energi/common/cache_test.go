// Copyright 2019 The Energi Core Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

package common

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
)

var dummyBlockHash = common.BytesToHash([]byte{1, 3, 4, 5})

type fakeChain struct{}

func (f *fakeChain) CurrentBlock() *types.Block {
	return types.NewBlock(&types.Header{
		ParentHash: dummyBlockHash,
	}, nil, nil, nil)
}

func (f *fakeChain) IsPublicService() bool {
	return true
}

// TestDataCache tests the cache's setter and getter methods.
func TestDataCache(t *testing.T) {
	chain := new(fakeChain)
	cacheInstance := NewCacheStorage()

	var newData interface{}
	cacheQueryfunc := func(num *big.Int) (interface{}, error) {
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

		if reflect.DeepEqual(data, newData) {
			t.Fatalf("expected the returned data to be old on the first call")
		}

		time.Sleep(100 * time.Millisecond)

		data, err = cacheInstance.Get(chain, cacheQueryfunc)
		if err != nil {
			t.Fatalf("expected no error but found %v", err)
		}

		if !reflect.DeepEqual(data, newData) {
			t.Fatalf("expected the returned data to match it didn't")
		}
	})

	t.Run("Test_stress", func(t *testing.T) {
		dummyBlockHash = common.BytesToHash([]byte{120, 23, 90, 7})
		newData = map[string]int{
			"rsc": 3711,
		}

		data, err := cacheInstance.Get(chain, cacheQueryfunc)
		if err != nil {
			t.Fatalf("expected no error but found %v", err)
		}

		if reflect.DeepEqual(data, newData) {
			t.Fatalf("expected the returned data to be old on the first call")
		}

		time.Sleep(1 * time.Second)

		tf := func() {
			data, err := cacheInstance.Get(chain, cacheQueryfunc)
			if err != nil {
				t.Fatalf("expected no error but found %v", err)
			}

			if !reflect.DeepEqual(data, newData) {
				t.Fatalf("expected the returned data to match it didn't")
			}
		}

		for i := 100; i > 0; i-- {
			tf()
		}
	})
}
