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
	"math/big"
	"sync"
	"sync/atomic"
	"unsafe"

	eth_common "github.com/ethereum/go-ethereum/common"
	eth_types "github.com/ethereum/go-ethereum/core/types"
)

// ErrInvalidData is returned if the CacheQuery function returns a null result
var ErrInvalidData = errors.New("Invalid data returned by the CacheQuery func")

// CacheQuery is the function that allow a fresh data query if the previous data
// held is considered to have expired or the cache was empty in the first place.
type CacheQuery func(block_num *big.Int) (interface{}, error)

type cacheState struct {
	blockHash eth_common.Hash
	entry     interface{}
}

// CacheStorage is a storage that is held by the client that wants to cache specific
// data.
type CacheStorage struct {
	mtx      sync.RWMutex
	state    unsafe.Pointer
	updating int32
}

// CacheChain defines the method(s) needed by the cache implementation to access
// the chain data.
type CacheChain interface {
	CurrentBlock() *eth_types.Block
}

// NewCacheStorage creates a new CacheStorage instance.
func NewCacheStorage() *CacheStorage {
	c := new(CacheStorage)
	state := &cacheState{}
	atomic.StorePointer(&c.state, unsafe.Pointer(state))
	return c
}

// Get returns the cached data entry if it hasn't expired(new blockhash generated).
// An error is returned if a nil cache instance is used or the cache query function
// returns nil data.
func (c *CacheStorage) Get(chain CacheChain, source CacheQuery) (interface{}, error) {
	block := chain.CurrentBlock()
	blockhash := block.Hash()

	state := (*cacheState)(atomic.LoadPointer(&c.state))

	do_update := func() (*cacheState, error) {
		state := (*cacheState)(atomic.LoadPointer(&c.state))

		// Concurrent update could happened
		if state.entry == nil || state.blockHash != blockhash {
			entry, err := source(block.Number())

			if err != nil {
				return nil, err
			}

			state = &cacheState{blockhash, entry}
			atomic.StorePointer(&c.state, unsafe.Pointer(state))
		}

		return state, nil
	}

	// First run or error recovery
	if state.entry == nil {
		c.mtx.Lock()
		defer c.mtx.Unlock()
		var err error
		state, err = do_update()
		if err != nil {
			return nil, err
		}

		// Update needed
	} else if state.blockHash != blockhash {
		if atomic.CompareAndSwapInt32(&c.updating, 0, 1) {
			go func() {
				defer atomic.StoreInt32(&c.updating, 0)
				defer func() {
					// Workaround for shutdown
					// and force to go blocking on any other error.
					if recover() != nil {
						state := &cacheState{}
						atomic.StorePointer(&c.state, unsafe.Pointer(state))
					}
				}()
				do_update()
			}()
		}
	}

	if state.entry == nil {
		return nil, ErrInvalidData
	}

	return state.entry, nil
}
