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
	// "sync/atomic"

	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/common"
	// "energi.world/core/gen3/core"
	// "energi.world/core/gen3/core/types"
	"energi.world/core/gen3/eth"
	// "energi.world/core/gen3/eth/downloader"
	"energi.world/core/gen3/log"
	"energi.world/core/gen3/p2p"
	"energi.world/core/gen3/rpc"

	energi_api "energi.world/core/gen3/energi/api"
	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
	energi_common "energi.world/core/gen3/energi/common"


)

var (
	logAllHfs = int32(1)
	// logIntervals defines the block interval in which pending blocks will be logged.
	logIntervals = big.NewInt(20)
)


const (
	//event channel default site
	EventChanBufferSize = 10
)

// HardforkService defines the hardfork service type.
type HardforkService struct {
	eth *eth.Ethereum
	ctx context.Context

	ctxCancel func()

	inSync int32
	callOpts   *bind.CallOpts

	hfAPI *energi_api.HardforkRegistryAPI
	hfRegistry *energi_abi.IHardforkRegistry

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

	//initialize Ihardforkregistry for further calls
	var err error
	hf.hfRegistry, err = energi_abi.NewIHardforkRegistry(hf.eth.APIBackend.ChainConfig().HardforkRegistryProxyAddress, hf.eth.APIBackend)
	if err != nil {
		log.Error("Failed to get create NewIHardforkRegistry (startup)", "err", err);
		return err
	}

	/*
	Upon startup retrieve all active hardforks and check if the
	active hardfork version parameter is higher than running core node version
	then log error to let user know the core node version is behind
	*/
	activeHardforks, err := hf.hfAPI.ListActiveHardforks()
	if err != nil {
		log.Error("Failed to get active hardforks (startup)", "err", err);
	} else if lc := len(activeHardforks); lc > 0 {
		hf.LogHardForks(activeHardforks);
	}

	//routine will listen to events thrown when hardfork is created
	go hf.listenHardforkCreatedEvents();
	// //routine will listen to hardfork finalization event
	// go hf.listenHardforkFinalizedEvents();
	// //routine will listen to events thrown when hardfork is removed
	// go hf.listenHardforkRemovedEvents();
	return nil

}

//function is watches newly created hardfork events and logs them
func (hf *HardforkService) listenHardforkCreatedEvents() {

	//create chan for subscribing to  HardforkCreated events
	hdCreatedChan := make(chan *energi_abi.IHardforkRegistryHardforkCreated, EventChanBufferSize)

	//create Opts for call
	watchOpts := &bind.WatchOpts{
		Context: context.WithValue(
			context.Background(),
			energi_params.GeneralProxyCtxKey,
			energi_common.GeneralProxyHashGen(hf.eth.BlockChain()),
		),
	}

	//subscribe to event
	subscribe, err := hf.hfRegistry.WatchHardforkCreated(watchOpts, hdCreatedChan, [][32]byte{})
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

		case hardfork := <-hdCreatedChan:
			log.Warn("New Hardfork  created: ",
							"block Number",
							hardfork.BlockNumber.String(),
							"hardfork Name",
							hardfork.Name,
							"hardfork SwFeatures",
							hardfork.SwFeatures.String())
		}
	}
}

func (hf *HardforkService) LogHardForks(hardforks []*energi_api.HardforkInfo)  {

	//atomically read the pointer to the most recent block header
	currentBlockHeader := hf.eth.BlockChain().CurrentBlock().Header()
	currentBlockNumber := currentBlockHeader.Number

	//get the hf finalization period parameter from config
	hfFinalizationPeriod := hf.eth.BlockChain().Config().HFFinalizationPeriod

	//for each hardfork name 	log the information considering the current block number
	for _, hardfork := range hardforks {
			 logHardforkInfo(currentBlockNumber, hfFinalizationPeriod, hardfork)
	}

}

// Stop terminates all goroutines belonging to the service, blocking until they
// are all terminated.
func (hf *HardforkService) Stop() error {
	log.Info("Shutting down Energi Hardforks")

	hf.ctxCancel()
	return nil
}

func (hf *HardforkService) listenDownloader() {
	// events := hf.eth.EventMux().Subscribe(
	// 	downloader.StartEvent{},
	// 	downloader.DoneEvent{},
	// 	downloader.FailedEvent{},
	// )
	// defer events.Unsubscribe()
	//
	// for {
	// 	select {
	// 	case <-hf.ctx.Done(): // Triggers immediate shutdown.
	// 		return
	// 	case ev := <-events.Chan():
	// 		if ev == nil {
	// 			return
	// 		}
	// 		switch ev.Data.(type) {
	// 		case downloader.StartEvent:
	// 			atomic.StoreInt32(&hf.inSync, 0)
	// 			log.Debug("Hardfork service is not in sync")
	// 		case downloader.DoneEvent, downloader.FailedEvent:
	// 			atomic.StoreInt32(&hf.inSync, 1)
	// 			log.Debug("Hardfork service is in sync")
	// 			return
	// 		}
	// 	}
	// }
}
//
// func (hf *HardforkService) onChainHead(block *types.Block) {
// 	hardforks, err := hf.hfAPI.ListHardforks()
// 	if err != nil {
// 		log.Warn("ListHardforks error", "err", err)
// 		return
// 	}
//
// 	if len(hardforks) < 1 {
// 		log.Debug("No hardforks currently available in the system")
// 		return
// 	}
//
// 	period := hf.eth.BlockChain().Config().HFFinalizationPeriod
//
// 	// The first time log the last loggedCount hardforks otherwise log only the last one.
// 	if atomic.CompareAndSwapInt32(&logAllHfs, 1, 0) {
// 		offset := len(hardforks) - loggedCount
// 		if offset < 0 {
// 			offset = 0
// 		}
//
// 		for _, hfInfo := range hardforks[offset:] {
// 			logHardforkInfo(block.Number(), period, hfInfo)
// 		}
//
// 		log.Info("Initial hardforks listing on startup", "logged", loggedCount,
// 			"remaining", offset)
// 	} else {
// 		pendingHardforks, er := hf.hfAPI.ListPendingHardforks()
// 		if er != nil {
// 			log.Warn("ListPendingHardforks", "err", err)
// 		}
//
// 		if len(pendingHardforks) < 1 && er == nil {
// 			log.Debug("No pending hardforks currently available in the system")
// 		}
//
//
// 		//check pendingHardforks not to be nil
// 		if er == nil {
// 			// Otherwise only log information about the pending hardforks.
// 			for _, hfInfo := range pendingHardforks {
// 				// log this data at intervals of logIntervals.
// 				mod := new(big.Int).Mod(block.Number(), logIntervals)
// 				if mod.Cmp(common.Big0) == 0 {
// 					logHardforkInfo(block.Number(), period, hfInfo)
// 				}
// 			}
// 		}
//
// 	}
//
// 	for _, fork := range hardforks {
// 		// Updates the current list of Active(finalized) Hardforks.
// 		energi_common.UpdateHfActive(fork.Name, fork.BlockNo.ToInt(),
// 			fork.BlockHash, fork.SWFeatures.ToInt())
// 	}
// }

//logHardfork logs the information about the provided hardforks.
func logHardforkInfo(currentBlockNo, period *big.Int, hfInfo *energi_api.HardforkInfo) {
	logFunc := log.Debug
	emptyHash := [32]byte{}
	hfBlockNo := hfInfo.BlockNumber.ToInt()
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
