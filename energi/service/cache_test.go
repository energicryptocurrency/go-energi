package eth

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// TestDataCache tests the cache's setter and getter methods.
func TestDataCache(t *testing.T) {
	type testData struct {
		name      string
		key       cacheKey
		blockhash *common.Hash
		entry     interface{}
		isError   bool
	}

	dummyBlockHash := common.BytesToHash([]byte{1, 3, 4, 5})

	td := []testData{
		{
			name:      "Test_adding_a_string",
			key:       ListGen2CoinsReq,
			blockhash: &dummyBlockHash,
			entry:     "ListGen2CoinsReq",
		},
		{
			name:      "Test_addding_an_array",
			key:       BlacklistInfoReq,
			blockhash: &dummyBlockHash,
			entry:     []int{1, 2, 3, 5, 6, 6},
		},
		{
			name:      "Test_adding_a_nil_struct",
			key:       CheckpointInfoReq,
			blockhash: &dummyBlockHash,
			entry:     struct{}{},
			isError:   true,
		},
		{
			name:    "Testing_adding_another_nil_data_case_scenario",
			key:     BudgetInfoReq,
			isError: true,
		},
	}

	for _, data := range td {
		t.Run(data.name, func(t *testing.T) {
			cacheInstance := NewCache()

			err := cacheInstance.Set(data.key, data.blockhash, data.entry)
			if err == nil && data.isError {
				t.Fatal("expected an error but found none")
			}

			if err != nil && !data.isError {
				t.Fatalf("expected no error but found %v", err)
			}

			newData, lastUpdated, er := cacheInstance.Get(data.key, data.blockhash)

			if data.isError && er != ErrMissingKeyData {
				// when an error is expected during cache update, no data was updated.
				t.Fatalf("expected the error returned to be (%v) but found (%v)",
					ErrExpiredData, err)
			}

			if !data.isError && er != nil {
				// If setting the cache was successful, fetching should not return
				// any error.
				t.Fatalf("expected no error but found %v", err)
			}

			if data.isError && lastUpdated != 0 {
				t.Fatal("expected the last update not to have been set but it was")
			}

			if !data.isError && lastUpdated == 0 {
				t.Fatal("expected the last update to have been set but it wasn't")
			}

			if data.isError && newData != nil {
				t.Fatal("expected no entry to have been set but it was set")
			}

			if !data.isError && newData == nil {
				t.Fatal("expected the entry to have been set but it wasn't")
			}
		})
	}
}
