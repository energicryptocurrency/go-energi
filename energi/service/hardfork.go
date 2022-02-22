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
	"strconv"
	"sync/atomic"

	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/eth"
	"github.com/energicryptocurrency/energi/eth/downloader"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/p2p"
	"github.com/energicryptocurrency/energi/rpc"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_api "github.com/energicryptocurrency/energi/energi/api"
	energi_common "github.com/energicryptocurrency/energi/energi/common"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/energi/api/hfcache"
)

const (
	//event channel default site
	EventChanBufferSize = 10
	//remaining number of blocks where we start logging/notifying user about upcoming pending hardfork
	lastBlockNumToLogPendingHardforks = int64(100)
)

// HardforkService defines the hardfork service type.
type HardforkService struct {
	eth       *eth.Ethereum
	ctx       context.Context
	ctxCancel func()

	inSync int32

	hfAPI      *energi_api.HardforkRegistryAPI
	hfRegistry *energi_abi.IHardforkRegistry
}

// NewHardforkService returns a new HardforkService instance.
func NewHardforkService(ethServ *eth.Ethereum) (*HardforkService, error) {
	hf := &HardforkService{
		eth:   ethServ,
		hfAPI: energi_api.NewHardforkRegistryAPI(ethServ.APIBackend),
	}
	hf.ctx, hf.ctxCancel = context.WithCancel(context.Background())

	//initialize Ihardforkregistry for further calls
	var err error
	hf.hfRegistry, err = energi_abi.NewIHardforkRegistry(hf.eth.APIBackend.ChainConfig().Energi.HardforkRegistryProxyAddress, hf.eth.APIBackend)
	if err != nil {
		log.Error("Failed to get create NewIHardforkRegistry (startup)", "err", err)
		return nil, err
	}

	//listen and log downloading/syncing status
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
	/*
		Upon startup retrieve all active hardforks and check if the
		active hardfork version parameter is higher than running core node version
		then log error to let user know the core node version is behind
	*/
	activeHardforks, err := hf.hfAPI.HardforkEnumerateActive()
	if err != nil {
		if err != bind.ErrNoCode {
			log.Error("Failed to get active hardforks (startup)", "err", err);
		} else {
			log.Debug("Hardfork contract hasn't been deployed yet", "err", err);
		}
	} else if lc := len(activeHardforks); lc > 0 {
		hf.LogHardForks(activeHardforks)
	}

	/*
	Upon startup retrieve all hardforks and store them in cache
	*/
	allHardforks, err := hf.hfAPI.HardforkEnumerate()
	if err != nil {
		if err != bind.ErrNoCode {
			log.Error("Failed to get hardforks (startup)", "err", err);
		} else {
			log.Debug("Hardfork contract hasn't been deployed yet", "err", err);
		}
	} else if lc := len(allHardforks); lc > 0 {
		for _, hardfork := range allHardforks {
			hfcache.AddHardfork(&hfcache.Hardfork{Name: hardfork.Name, BlockNumber: hardfork.BlockNumber})
		}
	}

	//routine will listen to events thrown when hardfork is created
	go hf.listenHardforkCreatedEvents()
	// //routine will listen to hardfork finalization event
	go hf.listenHardforkFinalizedEvents()
	//routine will listen to events thrown when hardfork is removed
	go hf.listenHardforkRemovedEvents()
	//logs upcoming pending hardforks notifying users about version change
	go hf.logUpcomingHardforks()

	return nil
}

//logUpcomingHardforks periodically logs upcoming (pending) hardforks
func (hf *HardforkService) logUpcomingHardforks() {
	//create channel and subscribe for new chain events
	chainHeadCh := make(chan core.ChainHeadEvent, chainHeadChanSize)
	headSub := hf.eth.BlockChain().SubscribeChainHeadEvent(chainHeadCh)
	defer headSub.Unsubscribe()

	//listen for blockchain change and notify users about upcoming hardforks
	for {
		select {

		case <-hf.ctx.Done(): // Triggers immediate shutdown.
			return

		case ev := <-chainHeadCh:
			pendingHardforks, err := hf.hfAPI.HardforkEnumeratePending()
			if err != nil {
				if err != bind.ErrNoCode {
					log.Error("Failed to get pending hardforks from api", "err", err)
				}
				break
			}

			// for each hardfork name log the information considering the current block number
			for _, hardfork := range pendingHardforks {
				logHardforkInfo(ev.Block.Header().Number, hardfork)
			}
		// Shutdown
		case <-headSub.Err():
			return
		}
	}
}

//function is watches newly created hardfork events and logs them
func (hf *HardforkService) listenHardforkCreatedEvents() {

	//create chan for subscribing to  HardforkCreated events
	hfCreatedChan := make(chan *energi_abi.IHardforkRegistryHardforkCreated, EventChanBufferSize)

	//create Opts for call
	watchOpts := &bind.WatchOpts{
		Context: context.WithValue(
			context.Background(),
			energi_params.GeneralProxyCtxKey,
			energi_common.GeneralProxyHashGen(hf.eth.BlockChain()),
		),
	}

	//subscribe to event
	subscribe, err := hf.hfRegistry.WatchHardforkCreated(watchOpts, hfCreatedChan, [][32]byte{})
	if err != nil {
		log.Error("Failed HardforkCreated subscription", "err", err)
		return
	}
	defer subscribe.Unsubscribe()

	//listen to events and log accordingly
	for {
		select {
		case err = <-subscribe.Err():
			log.Debug("HardforkCreated subscription error", "err", err)
			return

		case hardfork := <-hfCreatedChan:
			hfcache.AddHardfork(&hfcache.Hardfork{Name: string(bytes.Trim(hardfork.Name[:], "\x00")), BlockNumber: hardfork.BlockNumber})
			log.Warn("New Hardfork  created: ",
				"block Number",
				hardfork.BlockNumber.String(),
				"hardfork Name",
				string(hardfork.Name[:]),
				"hardfork SwFeatures",
				hardfork.SwFeatures.String())
		}
	}
}

//function is watches newly created hardfork events and logs them
func (hf *HardforkService) listenHardforkFinalizedEvents() {

	//create chan for subscribing to  HardforkCreated events
	hfFinalizedChan := make(chan *energi_abi.IHardforkRegistryHardforkFinalized, EventChanBufferSize)

	//create Opts for call
	watchOpts := &bind.WatchOpts{
		Context: context.WithValue(
			context.Background(),
			energi_params.GeneralProxyCtxKey,
			energi_common.GeneralProxyHashGen(hf.eth.BlockChain()),
		),
	}

	//subscribe to event
	subscribe, err := hf.hfRegistry.WatchHardforkFinalized(watchOpts, hfFinalizedChan, [][32]byte{})
	if err != nil {
		log.Error("Failed HardforkCreated subscription", "err", err)
		return
	}
	defer subscribe.Unsubscribe()

	//listen to events and log accordingly
	for {
		select {
		case err = <-subscribe.Err():
			log.Debug("HardforkCreated subscription error", "err", err)
			return

		case hardfork := <-hfFinalizedChan:
			log.Warn("New Hardfork Finalized: ",
				"block Number",
				hardfork.BlockNumber.String(),
				"block Hash",
				common.BytesToHash(hardfork.BlockHash[:]).String(),
				"hardfork Name",
				string(hardfork.Name[:]),
				"hardfork SwFeatures",
				hardfork.SwFeatures.String())
		}
	}
}

//function is watches newly created hardfork events and logs them
func (hf *HardforkService) listenHardforkRemovedEvents() {

	//create chan for subscribing to  HardforkCreated events
	hfRemovedChan := make(chan *energi_abi.IHardforkRegistryHardforkRemoved, EventChanBufferSize)

	//create Opts for call
	watchOpts := &bind.WatchOpts{
		Context: context.WithValue(
			context.Background(),
			energi_params.GeneralProxyCtxKey,
			energi_common.GeneralProxyHashGen(hf.eth.BlockChain()),
		),
	}

	//subscribe to event
	subscribe, err := hf.hfRegistry.WatchHardforkRemoved(watchOpts, hfRemovedChan, [][32]byte{})
	if err != nil {
		log.Error("Failed HardforkCreated subscription", "err", err)
		return
	}
	defer subscribe.Unsubscribe()

	//listen to events and log accordingly
	for {
		select {
		case err = <-subscribe.Err():
			log.Debug("HardforkCreated subscription error", "err", err)
			return

		case hardfork := <-hfRemovedChan:
			// remove hardfork from active hardfork cache
			hfcache.RemoveHardfork(string(bytes.Trim(hardfork.Name[:], "\x00")))
			log.Warn("Hardfork Removed: ",
				"Hardfork Name",
				string(hardfork.Name[:]))
		}
	}
}

func (hf *HardforkService) LogHardForks(hardforks []*energi_api.HardforkInfo) {

	//atomically read the pointer to the most recent block header
	currentBlockHeader := hf.eth.BlockChain().CurrentBlock().Header()
	currentBlockNumber := currentBlockHeader.Number

	//for each hardfork name 	log the information considering the current block number
	for _, hardfork := range hardforks {
		logHardforkInfo(currentBlockNumber, hardfork)
	}

}

// Stop terminates all goroutines belonging to the service, blocking until they
// are all terminated.
func (hf *HardforkService) Stop() error {
	log.Info("Shutting down Energi Hardforks")

	hf.ctxCancel()
	return nil
}

//log downloading status
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

//logHardfork logs the information about the provided hardforks.
func logHardforkInfo(currentBlockNo *big.Int, hfInfo *energi_api.HardforkInfo) {
	diff := new(big.Int).Sub(hfInfo.BlockNumber, currentBlockNo)
	emptyArray := [32]byte{}
	// check if hf is finalized (block hash set)
	if bytes.Equal(hfInfo.BlockHash[:], emptyArray[:]) {
		// check if hf is active (current block passed hf block)
		if diff.Cmp(common.Big0) <= 0 {
			// BlockHash not yet set but hardfork is active
			log.Warn("Active hard fork", "Name", hfInfo.Name, "activated at", hfInfo.BlockNumber.String())
		} else {
			// hardfork is to be activated in the future
			hours := strconv.FormatInt(diff.Int64()/60, 10)
			minutes := strconv.FormatInt(diff.Int64()%60, 10)
			if diff.Int64() < lastBlockNumToLogPendingHardforks {
				log.Warn("Hard fork will activate in approximately "+hours+
					" hours and "+minutes+" minutes", "hardfork Name", hfInfo.Name)
			} else {
				log.Debug("Hard fork will activate in approximately "+hours+
					" hours and "+minutes+" minutes", "hardfork Name", hfInfo.Name)
			}
		}
	} else {
		// BlockHash already set. Hardfork already finalized.
		log.Warn("Hardfork already finalized", "block Number", hfInfo.BlockNumber,
			"hardfork Name", hfInfo.Name, "block Hash", hfInfo.BlockHash.String())
	}
}
