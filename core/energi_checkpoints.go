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

package core

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/crypto"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/event"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/params"
)

// max number of checkpoints stored in validated checkpoint map
const MaxCachedCheckpoints int = 10

type CheckpointValidateChain interface {
	GetHeaderByNumber(number uint64) *types.Header
	CurrentHeader() *types.Header
}

type CheckpointChain interface {
	CheckpointValidateChain

	EnforceCheckpoint(cp Checkpoint) error
	Config() *params.ChainConfig
}

type Checkpoint struct {
	Since  uint64
	Number uint64
	Hash   common.Hash
}

// Format implements fmt.Formatter, forcing the Checkpoint to be formatted as is,
// without going through the stringer interface used for logging.
func (cp Checkpoint) Format(s fmt.State, c rune) {
	cpStr := struct {
		Number uint64
		Hash   string
	}{
		cp.Number,
		cp.Hash.String(),
	}
	fmt.Fprintf(s, "%+"+string(c), cpStr)
}

type CheckpointSignature []byte

type CheckpointInfo struct {
	Checkpoint
	CppSignature CheckpointSignature
	SigCount     uint64
}

type NewCheckpointEvent struct {
	CheckpointInfo
}

type validCheckpoint struct {
	Checkpoint
	signatures []CheckpointSignature
}

type futureCheckpoint struct {
	Checkpoint
}

type checkpointManager struct {
	validated map[uint64]validCheckpoint
	latest    uint64
	future    map[uint64]futureCheckpoint
	mtx       sync.RWMutex
	newCpFeed event.Feed
}

func newCheckpointManager() *checkpointManager {
	return &checkpointManager{
		validated: make(map[uint64]validCheckpoint),
		future:    make(map[uint64]futureCheckpoint),
	}
}

func (cm *checkpointManager) setup(chain CheckpointChain) {
	genesis_hash := chain.GetHeaderByNumber(0).Hash()
	if checkpoints, ok := energi_params.EnergiCheckpoints[genesis_hash]; ok {
		for k, v := range checkpoints {
			cm.addCheckpoint(
				chain,
				Checkpoint{
					Number: k,
					Hash:   v,
				},
				[]CheckpointSignature{},
				true,
			)
		}
	}
}

func (cm *checkpointManager) validate(chain CheckpointValidateChain, num uint64, hash common.Hash) error {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	// Check against validated checkpoints & mismatch
	if cp, ok := cm.validated[num]; ok {
		if cp.Hash != hash {
			return ErrCheckpointMismatch
		}

		cm.updateLatest(chain, &cp.Checkpoint)

		return nil
	}

	// Check if before the latest checkpoint & mismatch
	if num < cm.latest {
		header := chain.GetHeaderByNumber(num)

		if header != nil && header.Hash() != hash {
			return ErrCheckpointMismatch
		}

		return nil
	}

	// TODO: proper future checkpoint processing
	if cp, ok := cm.future[num]; ok {
		if cp.Hash != hash {
			return ErrCheckpointMismatch
		}

		return nil
	}

	return nil
}

// returns the smallest key (blockHeight)
func oldestCheckpoint(validated map[uint64]validCheckpoint) uint64 {
	minHeight := uint64(math.MaxUint64)
	for k := range validated {
		if k < minHeight {
			minHeight = k
		}
	}
	return minHeight
}

func (bc *BlockChain) AddCheckpoint(
	cp Checkpoint,
	sigs []CheckpointSignature,
	local bool,
) error {
	return bc.checkpoints.addCheckpoint(bc, cp, sigs, local)
}

func (cm *checkpointManager) addCheckpoint(
	chain CheckpointChain,
	cp Checkpoint,
	sigs []CheckpointSignature,
	local bool,
) (err error) {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	if curr, ok := cm.validated[cp.Number]; ok {
		if curr.Checkpoint == cp {
			return nil
		}

		if curr.Since > cp.Since {
			return nil
		}
	}

	if !local {
		// ignore checkpoints which occur before the latest local checkpoint
		var maxHardcodedCheckpoint uint64
		genesis_hash := chain.GetHeaderByNumber(0).Hash()
		for maxHardcodedCheckpoint = range energi_params.EnergiCheckpoints[genesis_hash] {
			break
		}
		for n := range energi_params.EnergiCheckpoints[genesis_hash] {
			if n > maxHardcodedCheckpoint {
				maxHardcodedCheckpoint = n
			}
		}

		if cp.Number <= maxHardcodedCheckpoint {
			//log.Info("Ignoring checkpoint which occurs before latest checkpoint at", "block", maxHardcodedCheckpoint)
			return nil
		}

		// TODO: proper validation and use of future checkpoints
		if len(sigs) == 0 {
			log.Warn("Checkpoint: missing signatures",
				"num", cp.Number, "hash", cp.Hash)
			return errors.New("missing checkpoint signatures")
		}

		// The first one must always be CPP_signer
		pubkey, err := crypto.Ecrecover(cm.hashToSign(&cp), sigs[0][:])
		if err != nil {
			log.Warn("Checkpoint: failed to extract signature",
				"num", cp.Number, "hash", cp.Hash, "err", err)
			return err
		}

		// Check the primary signature
		var signer common.Address
		copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
		if nrgconf := chain.Config().Energi; nrgconf == nil || signer != nrgconf.CPPSigner {
			log.Warn("Checkpoint: invalid CPP signature", "num", cp.Number, "hash", cp.Hash)
			return errors.New("invalid CPP signature")
		}

	}

	//only received(non-hardcoded) checkpoints will be stored in validated map
	if len(cm.validated) == MaxCachedCheckpoints {
		oldestCheckpointHeight := oldestCheckpoint(cm.validated)
		if cp.Number > oldestCheckpointHeight {
			delete(cm.validated, oldestCheckpointHeight)
			cm.validated[cp.Number] = validCheckpoint{
				Checkpoint: cp,
				signatures: append([]CheckpointSignature{}, sigs...),
			}
		}
	} else {
		cm.validated[cp.Number] = validCheckpoint{
			Checkpoint: cp,
			signatures: append([]CheckpointSignature{}, sigs...),
		}
	}

	log.Info("Added new checkpoint", "checkpoint", cp, "local", local)

	err = chain.EnforceCheckpoint(cp)

	cm.updateLatest(chain, &cp)

	if !local {
		// Send regardless of enforcement success
		cm.newCpFeed.Send(NewCheckpointEvent{CheckpointInfo{cp, sigs[0], uint64(len(sigs))}})
	}

	return err
}

func (cm *checkpointManager) hashToSign(cp *Checkpoint) []byte {
	data := []byte("||Energi Blockchain Checkpoint||")
	data = append(data, common.BigToHash(new(big.Int).SetUint64(cp.Number)).Bytes()...)
	data = append(data, cp.Hash.Bytes()...)
	return crypto.Keccak256(data)
}

func (cm *checkpointManager) updateLatest(chain CheckpointValidateChain, cp *Checkpoint) {
	if cp.Number > cm.latest && cp.Number <= chain.CurrentHeader().Number.Uint64() {
		cm.latest = cp.Number
		log.Info("Latest checkpoint", "height", cp.Number, "hash", cp.Hash.Hex())
	}
}

func (bc *BlockChain) EnforceCheckpoint(cp Checkpoint) error {
	header := bc.GetHeaderByNumber(cp.Number)

	if header != nil && header.Hash() != cp.Hash {
		log.Error("Side chain is detected as canonical", "number", cp.Number, "hash", cp.Hash, "old", header.Hash())

		if cp_block := bc.GetBlock(cp.Hash, cp.Number); cp_block != nil {
			// Known block
			bc.mu.Lock()
			defer bc.mu.Unlock()

			if err := bc.reorg(bc.GetBlock(header.Hash(), cp.Number), cp_block); err != nil {
				log.Crit("Failed to reorg", "err", err)
				// should terminate
				return err
			}

			log.Warn("Chain reorg was successful, resuming normal operation")
		} else {
			// Unknown block
			if err := bc.SetHead(cp.Number - 1); err != nil {
				log.Crit("Failed to rewind before fork point", "err", err)
				// should terminate
				return err
			}
			log.Warn("Chain rewind was successful, resuming normal operation")
		}
	}

	return nil
}

func (bc *BlockChain) ListCheckpoints() []CheckpointInfo {
	cm := bc.checkpoints

	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	res := make([]CheckpointInfo, 0, len(cm.validated))

	for _, v := range cm.validated {
		if len(v.signatures) > 0 {
			res = append(res, CheckpointInfo{v.Checkpoint, v.signatures[0], uint64(len(v.signatures))})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Since > res[j].Since
	})

	return res
}

func (bc *BlockChain) CheckpointSignatures(cp Checkpoint) []CheckpointSignature {
	cm := bc.checkpoints

	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	if vcp, ok := cm.validated[cp.Number]; ok && vcp.Hash == cp.Hash {
		return append([]CheckpointSignature{}, vcp.signatures...)
	}

	return nil
}

func (bc *BlockChain) SubscribeNewCheckpointEvent(ch chan<- NewCheckpointEvent) event.Subscription {
	return bc.scope.Track(bc.checkpoints.newCpFeed.Subscribe(ch))
}

func (bc *BlockChain) IsRunning() bool {
	return atomic.LoadInt32(&bc.running) == 0
}
