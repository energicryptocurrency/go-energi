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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_common "energi.world/core/gen3/energi/common"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	cppChanBufferSize = 10
)

type CheckpointProposalEvent struct {
	core.Checkpoint
	Proposal common.Address
}

type CheckpointService struct {
	server *p2p.Server
	eth    *eth.Ethereum

	cpRegistry *energi_abi.ICheckpointRegistry
	callOpts   *bind.CallOpts
}

func NewCheckpointService(ethServ *eth.Ethereum) (node.Service, error) {
	r := &CheckpointService{
		eth:      ethServ,
		callOpts: &bind.CallOpts{},
	}
	return r, nil
}

func (c *CheckpointService) Protocols() []p2p.Protocol {
	return nil
}

func (c *CheckpointService) APIs() []rpc.API {
	return nil
}

func (c *CheckpointService) Start(server *p2p.Server) (err error) {
	c.cpRegistry, err = energi_abi.NewICheckpointRegistry(
		energi_params.Energi_CheckpointRegistry, c.eth.APIBackend)
	if err != nil {
		return err
	}

	c.server = server

	//---
	oldCheckpoints, err := c.cpRegistry.Checkpoints(c.callOpts)
	if err != nil {
		log.Error("Failed to get old checkpoints (startup)", "err", err)
	} else if lc := len(oldCheckpoints); lc > 0 {
		// NOTE: enable the latest checkpoint immediately for safety reasons
		c.onCheckpoint(oldCheckpoints[lc-1], false)
	}

	//---
	go c.loop()
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

	for {
		select {
		case ev := <-events.Chan():
			if ev == nil {
				return false
			}
			switch ev.Data.(type) {
			case downloader.StartEvent:
				continue
			case downloader.DoneEvent:
				return true
			case downloader.FailedEvent:
				return c.eth.BlockChain().IsRunning()
			}
		}
	}
}

func (c *CheckpointService) loop() {
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

	oldCheckpoints, err := c.cpRegistry.Checkpoints(c.callOpts)
	if err != nil {
		log.Error("Failed to get old checkpoints", "err", err)
	} else {
		// NOTE: we should feed for recent first
		for i := len(oldCheckpoints) - 1; i >= 0; i-- {
			c.onCheckpoint(oldCheckpoints[i], false)
		}
	}

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

func (c *CheckpointService) onCheckpoint(cpAddr common.Address, live bool) {
	backend := c.eth.APIBackend
	cppSigner := backend.ChainConfig().Energi.CPPSigner

	cp, err := energi_abi.NewICheckpointCaller(cpAddr, backend)
	if err != nil {
		log.Warn("Failed to create CP contract caller", "addr", cpAddr, "err", err)
		return
	}

	info, err := cp.Info(c.callOpts)
	if err != nil {
		log.Warn("Failed to get CP info", "addr", cpAddr, "err", err)
		return
	}

	cpp_sig, err := cp.Signature(c.callOpts, cppSigner)
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
