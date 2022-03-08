package hfcache

import (
	"sync"
	"math/big"
)

const (
	HfNameLenght = 32
)

var (
	hardforkCache = &HardforkCache{cacheLock: &sync.Mutex{}}
)

// AddActiveHardfork adds a new active hardfork
func AddHardfork(hardfork *Hardfork) {
	hardforkCache.cacheLock.Lock()
	defer hardforkCache.cacheLock.Unlock()
	hardforkCache.hardforks = append(hardforkCache.hardforks, hardfork)
}

// RemoveActiveHardfork removes hardfork
func RemoveHardfork(hfName string) {
	hardforkCache.cacheLock.Lock()
	defer hardforkCache.cacheLock.Unlock()
	for i, activeHardfork := range hardforkCache.hardforks {
		if hfName == activeHardfork.Name {
			hardforkCache.hardforks[i] = hardforkCache.hardforks[len(hardforkCache.hardforks)-1] // Copy last element to index i.
			hardforkCache.hardforks[len(hardforkCache.hardforks)-1] = nil   // Erase last element (write zero value).
			hardforkCache.hardforks = hardforkCache.hardforks[:len(hardforkCache.hardforks)-1]   // Truncate slice.
			return
		}
	}
}

// IsHardforkActive checks if given hardfork is active
func IsHardforkActive(hardforkName string, blockNum uint64) bool {
	hardforkCache.cacheLock.Lock()
	defer hardforkCache.cacheLock.Unlock()
	if len(hardforkName) > HfNameLenght {
		return false
	}

	for _, hardfork := range hardforkCache.hardforks {
		if hardfork.Name == hardforkName && blockNum >= hardfork.BlockNumber.Uint64() {
			return true
		}
	}
	return false
}



// HardforkInfo defines the hardfork payload information returned.
type Hardfork struct {
	Name        string      `json:"name"`
	BlockNumber *big.Int    `json:"block_number"`
}

// HardforkCache caches currently active hardforks
type HardforkCache struct {
	hardforks []*Hardfork
	cacheLock *sync.Mutex
}
