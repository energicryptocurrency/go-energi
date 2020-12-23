// Copyright 2020 The Energi Core Authors
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

package service

import (
	"bytes"
	"context"
	"math/big"
	"sync/atomic"

	"energi.world/core/gen3/common"
	"energi.world/core/gen3/core"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/eth"
	"energi.world/core/gen3/eth/downloader"
	"energi.world/core/gen3/log"
	"energi.world/core/gen3/p2p"
	"energi.world/core/gen3/rpc"

	energi_api "energi.world/core/gen3/energi/api"
	energi_common "energi.world/core/gen3/energi/common"
)

var logAllHfs = int32(1)

// logIntervals defines the block interval in which pending blocks will be logged.
var logIntervals = big.NewInt(20)

// HardforkService defines the hardfork service type.
type HardforkService struct {
	eth *eth.Ethereum
	ctx context.Context

	ctxCancel func()

	inSync int32

	hfAPI *energi_api.HardforkRegistryAPI
}

// NewHardforkService returns a new HardforkService instance.
func NewHardforkService(ethServ *eth.Ethereum) (*HardforkService, error) {
	hf := &HardforkService{
		eth:   ethServ,
		hfAPI: energi_api.NewHardforkRegistryAPI(ethServ.APIBackend),
	}

	hf.ctx, hf.ctxCancel = context.WithCancel(context.Background())
	go hf.listenDownloader()

	return hf, nil
}

// Protocols retrieves the P2P protocols the service wishes to start.
func (hf *HardforkService) Protocols() []p2p.Protocol {
	return nil
}

// APIs retrieves the list of RPC descriptors the service provides
func (hf *HardforkService) APIs() []rpc.API {
	return nil
}

// Start is called after all services have been constructed and the networking
// layer was also initialized to spawn any goroutines required by the service.
func (hf *HardforkService) Start(server *p2p.Server) error {
	go func() {
		chainHeadCh := make(chan core.ChainHeadEvent, chainHeadChanSize)
		headSub := hf.eth.BlockChain().SubscribeChainHeadEvent(chainHeadCh)
		defer headSub.Unsubscribe()

		//---
		for {
			select {
			case <-hf.ctx.Done(): // Triggers immediate shutdown.
				return

			case ev := <-chainHeadCh:
				hf.onChainHead(ev.Block)
				break

			// Shutdown
			case <-headSub.Err():
				return
			}
		}
	}()

	return nil
}

// Stop terminates all goroutines belonging to the service, blocking until they
// are all terminated.
func (hf *HardforkService) Stop() error {
	log.Info("Shutting down Energi Hardforks")

	hf.ctxCancel()
	return nil
}

func (hf *HardforkService) listenDownloader() {
	events := hf.eth.EventMux().Subscribe(
		downloader.StartEvent{},
		downloader.DoneEvent{},
		downloader.FailedEvent{},
	)
	defer events.Unsubscribe()

	for {
		select {
		case <-hf.ctx.Done(): // Triggers immediate shutdown.
			return
		case ev := <-events.Chan():
			if ev == nil {
				return
			}
			switch ev.Data.(type) {
			case downloader.StartEvent:
				atomic.StoreInt32(&hf.inSync, 0)
				log.Debug("Hardfork service is not in sync")
			case downloader.DoneEvent, downloader.FailedEvent:
				atomic.StoreInt32(&hf.inSync, 1)
				log.Debug("Hardfork service is in sync")
				return
			}
		}
	}
}

func (hf *HardforkService) onChainHead(block *types.Block) {
	hardforks, err := hf.hfAPI.ListHardforks()
	if err != nil {
		log.Warn("ListHardforks error", "err", err)
		return
	}

	if len(hardforks) < 1 {
		log.Debug("No hardforks currently available in the system")
		return
	}

	period := hf.eth.BlockChain().Config().HFFinalizationPeriod

	// The first time log the last 10 hardforks otherwise log only the last one.
	if atomic.CompareAndSwapInt32(&logAllHfs, 1, 0) {
		loggedCount := 10 // limiting logged hardforks allows proper real estate usage.
		offset := len(hardforks) - loggedCount
		if offset < 0 {
			offset = 0
		}

		for _, hfInfo := range hardforks[offset:] {
			logHardforkInfo(block.Number(), period, hfInfo)
		}

		log.Info("Initial hardforks listing on startup", "logged", loggedCount,
			"remaining", offset)
	} else {
		pendingHardforks, er := hf.hfAPI.ListPendingHardforks()
		if er != nil {
			log.Warn("ListPendingHardforks", "err", err)
		}

		if len(pendingHardforks) < 1 && er == nil {
			log.Debug("No pending hardforks currently available in the system")
		}


		//check pendingHardforks not to be nil
		if er == nil {
			// Otherwise only log information about the pending hardforks.
			for _, hfInfo := range pendingHardforks {
				// log this data at intervals of logIntervals.
				mod := new(big.Int).Mod(block.Number(), logIntervals)
				if mod.Cmp(common.Big0) == 0 {
					logHardforkInfo(block.Number(), period, hfInfo)
				}
			}
		}

	}

	for _, fork := range hardforks {
		// Updates the current list of Active(finalized) Hardforks.
		energi_common.UpdateHfActive(fork.Name, fork.BlockNo.ToInt(),
			fork.BlockHash, fork.SWFeatures.ToInt())
	}
}

// logHardfork logs information about the information about the provided hardfork.
func logHardforkInfo(currentBlockNo, period *big.Int, hfInfo *energi_api.HardforkInfo) {
	logFunc := log.Debug
	emptyHash := [32]byte{}
	hfBlockNo := hfInfo.BlockNo.ToInt()
	diff := new(big.Int).Sub(currentBlockNo, hfBlockNo)

	if bytes.Compare(hfInfo.BlockHash[:], emptyHash[:]) == 0 {
		if diff.Cmp(big.NewInt(-10)) > 0 && diff.Cmp(period) <= 0 {
			// -10 < Currentblock - hfblock <= hfPeriod
			logFunc = log.Warn
		}

		desc := "blocks To Hardfork"
		if diff.Cmp(common.Big0) > 0 {
			desc = "blocks after Hardfork"
		}

		// BlockHash not yet set.
		logFunc("Hardfork almost being finalized", "block Number", hfBlockNo,
			"hardfork Name", hfInfo.Name, desc, new(big.Int).Abs(diff),
		)
	} else {
		if diff.Cmp(common.Big0) > 0 && diff.Cmp(period) <= 0 {
			// 0 < Currentblock - hfblock <= hfPeriod
			logFunc = log.Info
		}

		// BlockHash already set. Hardfork already finalized.
		logFunc("Hardfork already finalized", "block Number", hfBlockNo,
			"hardfork Name", hfInfo.Name, "block Hash", hfInfo.BlockHash.String(),
		)
	}
}
