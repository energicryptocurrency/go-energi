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

package eth

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// cacheKey defines the type of unique keys to be used to index the respective
// data entries.
type cacheKey int

const (
	// ListGen2CoinsReq is cache key for ListGen2Coins data.
	ListGen2CoinsReq cacheKey = iota

	// BlacklistInfoReq is cache key for BlacklistInfo data.
	BlacklistInfoReq

	// CheckpointInfoReq is cache key for CheckpointInfo data.
	CheckpointInfoReq

	// UpgradeInfoReq is cache key for UpgradeInfo data.
	UpgradeInfoReq

	// BudgetInfoReq is cache key for BudgetInfo data.
	BudgetInfoReq

	// ListMasternodesReq is cache key for ListMasternodes data.
	ListMasternodesReq

	// MasternodesStatsReq is cache key for masternode stats data.
	MasternodesStatsReq

	// validityInterval is the interval after the last update time when cache data
	// entry is considered to have expired.
	validityInterval = time.Minute * 2
)

var (
	// ErrMissingKeyData is returned when attempting to fetch data that is
	// non-existent in the cache.
	ErrMissingKeyData = errors.New("Data for the provided key is missing")

	// ErrExpiredData is returned when attempting to fetch data that has already
	// expired.
	ErrExpiredData = errors.New("Data found has already expired")
)

// CacheStorage defines the interface required to access and retrieved cachec data.
type CacheStorage interface {
	Set(key cacheKey, blockhash *common.Hash, data interface{}) error
	Get(key cacheKey, blockhash *common.Hash) (data interface{}, lastUpdated uint64, err error)
}

// NewCache returns an empty cache instance that is ready to hold data till it expires.
func NewCache() CacheStorage {
	// returns a cache instance with a map whose default size is 10 entries.
	return &cache{data: make(map[cacheKey]*cacheEntry, 10)}
}

// cache is thread-safe to access since its mutex protected.
type cache struct {
	mtx  sync.RWMutex
	data map[cacheKey]*cacheEntry
}

// cacheEntry defines the individual entry that is mapped to a cache key.
type cacheEntry struct {
	updated   time.Time
	blockHash *common.Hash
	entry     interface{}
}

// Set append the newly provided entry to the cache. Nil data cannot not be
// added to the cache as an entry.
func (c *cache) Set(key cacheKey, blockhash *common.Hash, data interface{}) error {
	if data == nil {
		return fmt.Errorf("Invalid data found cannot be pushed to the cache")
	}

	if blockhash == nil {
		return fmt.Errorf("Invalid blockhash found")
	}

	c.mtx.RLocker()
	defer c.mtx.Unlock()

	c.data[key] = &cacheEntry{
		blockHash: blockhash,
		updated:   time.Now(),
		entry:     data,
	}

	return nil
}

// Get fetches the data entry mapped to the provided key. An error is returned if
// data to be mapped to the key is missing or has expired. The change in blockhash
// also helps in determining when the cache has expired.
func (c *cache) Get(key cacheKey, blockhash *common.Hash) (data interface{}, lastUpdated uint64, err error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	entry, ok := c.data[key]
	if !ok {
		err = ErrMissingKeyData
		return
	}

	timeAtExpiry := entry.updated.Add(validityInterval)
	if time.Now().After(timeAtExpiry) || !bytes.Equal(blockhash.Bytes(), entry.blockHash.Bytes()) {
		err = ErrExpiredData
		return
	}

	data = entry.entry
	lastUpdated = uint64(entry.updated.UnixNano())

	return
}
