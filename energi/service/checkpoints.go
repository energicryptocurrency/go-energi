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

package service

import (
	"context"
	"math/big"

	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_api "github.com/energicryptocurrency/energi/energi/api"
	energi_common "github.com/energicryptocurrency/energi/energi/common"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/eth"
	"github.com/energicryptocurrency/energi/eth/downloader"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/node"
	"github.com/energicryptocurrency/energi/p2p"
	"github.com/energicryptocurrency/energi/rpc"
)

const (
	cppChanBufferSize = 10
)

type CheckpointProposalEvent struct {
	core.Checkpoint
	Proposal common.Address
}

type CheckpointService struct {
	eth        *eth.Ethereum
	cpRegistry *energi_abi.ICheckpointRegistry
	cpAPI      *energi_api.CheckpointRegistryAPI
}

func NewCheckpointService(ethServ *eth.Ethereum) (node.Service, error) {
	c := &CheckpointService{
		eth:   ethServ,
		cpAPI: energi_api.NewCheckpointRegistryAPI(ethServ.APIBackend),
	}

	//initialize Icheckpointregistry for further calls
	var err error
	c.cpRegistry, err = energi_abi.NewICheckpointRegistry(energi_params.Energi_CheckpointRegistry, c.eth.APIBackend)
	if err != nil {
		log.Error("Failed to get create NewICheckpointkRegistry (startup)", "err", err)
		return nil, err
	}

	return c, nil
}

func (c *CheckpointService) Protocols() []p2p.Protocol {
	return nil
}

func (c *CheckpointService) APIs() []rpc.API {
	return nil
}

func (c *CheckpointService) Start(server *p2p.Server) (err error) {
	// retrieve last checkpoints and ensure that  the last one if valid for the current chain
	oldCheckpoints, err := c.cpAPI.Checkpoints()

	if err != nil {
		log.Error("Failed to get old checkpoints (startup)", "err", err)
	} else if lc := len(oldCheckpoints); lc > 0 {
		// NOTE: enable the latest checkpoint immediately for safety reasons
		c.onCheckpoint(oldCheckpoints[lc-1], false)
	}

	// watch for new checkpoints
	go c.watchCheckpoints()
	return nil
}

func (c *CheckpointService) Stop() error {
	return nil
}

func (c *CheckpointService) waitDownloader() bool {
	events := c.eth.EventMux().Subscribe(
		downloader.StartEvent{},
		downloader.DoneEvent{},
		downloader.FailedEvent{},
	)
	defer events.Unsubscribe()

	for ev := range events.Chan() {
		if ev == nil {
			return false
		}
		switch ev.Data.(type) {
		case downloader.StartEvent:
			log.Debug("Checkpoint service is not in sync")
			continue
		case downloader.DoneEvent:
			log.Debug("Checkpoint service is in sync")
			return true
		case downloader.FailedEvent:
			return c.eth.BlockChain().IsRunning()
		}
	}
	return false
}

func (c *CheckpointService) watchCheckpoints() {
	if !c.waitDownloader() {
		return
	}
	cpChan := make(chan *energi_abi.ICheckpointRegistryCheckpoint, cppChanBufferSize)
	watchOpts := &bind.WatchOpts{
		Context: context.WithValue(
			context.Background(),
			energi_params.GeneralProxyCtxKey,
			energi_common.GeneralProxyHashGen(c.eth.BlockChain()),
		),
	}

	subscribe, err := c.cpRegistry.WatchCheckpoint(watchOpts, cpChan, []*big.Int{})
	if err != nil {
		log.Error("Failed checkpoint subscription", "err", err)
		return
	}
	defer subscribe.Unsubscribe()

	// process existing checkpoints first
	oldCheckpoints, err := c.cpAPI.Checkpoints()
	if err != nil {
		log.Error("Failed to get old checkpoints", "err", err)
	} else {
		// NOTE: we should feed for recent first
		for i := len(oldCheckpoints) - 1; i >= 0; i-- {
			c.onCheckpoint(oldCheckpoints[i], false)
		}
	}

	// listen for incoming new checkpoints
	for {
		select {
		case err = <-subscribe.Err():
			log.Debug("Checkpoint subscription error", "err", err)
			return

		case cpData := <-cpChan:
			c.onCheckpoint(cpData.Checkpoint, true)
		}
	}
}

// process checkpoint
func (c *CheckpointService) onCheckpoint(cpAddr common.Address, live bool) {
	backend := c.eth.APIBackend
	cppSigner := backend.ChainConfig().Energi.CPPSigner

	cp, err := energi_abi.NewICheckpointV2Caller(cpAddr, backend)
	if err != nil {
		log.Warn("Failed to create CP contract caller", "addr", cpAddr, "err", err)
		return
	}

	info, err := cp.Info(&bind.CallOpts{})
	if err != nil {
		log.Warn("Failed to get CP info", "addr", cpAddr, "err", err)
		return
	}

	cpp_sig, err := cp.Signature(&bind.CallOpts{}, cppSigner)
	if err != nil {
		log.Debug("Skipping checkpoint with no CPP sig", "addr", cpAddr, "err", err)
		return
	}

	if len(cpp_sig) >= 65 {
		// Drop Ecrecover workaround
		cpp_sig[64] -= 27
	}

	sigs := []core.CheckpointSignature{
		core.CheckpointSignature(cpp_sig),
	}

	// rebroadcast the received checkpoint
	backend.AddDynamicCheckpoint(info.Since.Uint64(), info.Number.Uint64(), info.Hash, sigs)
	if live {
		log.Warn("Found new dynamic checkpoint", "num", info.Number, "hash", common.Hash(info.Hash).Hex())
		c.eth.EventMux().Post(CheckpointProposalEvent{
			core.Checkpoint{
				Since:  info.Since.Uint64(),
				Number: info.Number.Uint64(),
				Hash:   info.Hash,
			},
			cpAddr,
		})
	}
}
