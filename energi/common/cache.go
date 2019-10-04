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
	"errors"
	"sync"

	eth_common "github.com/ethereum/go-ethereum/common"
	eth_types "github.com/ethereum/go-ethereum/core/types"
)

// ErrInvalidData is returned if the CacheQuery function returns a null result
var ErrInvalidData = errors.New("Invalid data returned by the CacheQuery func")

// CacheQuery is the function that allow a fresh data query if the previous data
// held is considered to have expired or the cache was empty in the first place.
type CacheQuery func(blockhash eth_common.Hash) (interface{}, error)

// CacheStorage is a storage that is held by the client that wants to cache specific
// data.
type CacheStorage struct {
	mtx       sync.RWMutex
	blockHash eth_common.Hash
	entry     interface{}
}

// CacheChain defines the method(s) needed by the cache implementation to access
// the chain data.
type CacheChain interface {
	CurrentBlock() *eth_types.Block
}

// NewCacheStorage creates a new CacheStorage instance.
func NewCacheStorage() *CacheStorage {
	return new(CacheStorage)
}

// Get returns the cached data entry if it hasn't expired(new blockhash generated).
// An error is returned if a nil cache instance is used or the cache query function
// returns nil data.
func (c *CacheStorage) Get(chain CacheChain, source CacheQuery) (interface{}, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	blockhash := chain.CurrentBlock().Hash()
	if c.entry == nil || blockhash != c.blockHash {
		var err error
		if c.entry, err = source(blockhash); err != nil {
			return nil, err
		}

		if c.entry == nil {
			return nil, ErrInvalidData
		}
	}

	return c.entry, nil
}
