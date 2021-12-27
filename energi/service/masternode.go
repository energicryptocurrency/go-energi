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
	"errors"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/crypto"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_common "github.com/energicryptocurrency/energi/energi/common"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/eth"
	"github.com/energicryptocurrency/energi/eth/downloader"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/node"
	"github.com/energicryptocurrency/energi/p2p"
	"github.com/energicryptocurrency/energi/rpc"
)

var (
	heartbeatInterval = time.Duration(5) * time.Minute
	recheckInterval   = time.Duration(2) * time.Minute
)

const (
	masternodeCallGas uint64 = 500000

	// cpChanBufferSize defines the number of checkpoint to be pushed into the
	// checkpoints channel before it can be considered to be full.
	cpChanBufferSize  = 16
	chainHeadChanSize = 10
)

type checkpointVote struct {
	address   common.Address
	signature []byte
}

type MasternodeService struct {
	server *p2p.Server
	eth    *eth.Ethereum

	inSync int32

	address  common.Address
	owner    common.Address
	registry *energi_abi.IMasternodeRegistryV2Session

	cpRegistry *energi_abi.ICheckpointRegistrySession
	cpVoteChan chan *checkpointVote

	nextHB   time.Time
	features *big.Int

	validator *peerValidator
}

func NewMasternodeService(ethServ *eth.Ethereum, owner common.Address) (node.Service, error) {
	r := &MasternodeService{
		eth:      ethServ,
		inSync:   1,
		features: energi_common.SWVersionToInt(),
		owner:    owner,
		// NOTE: we need to avoid triggering DoS on restart.
		// There is no reliable way to check blockchain and all pools in the network.
		nextHB: time.Now().Add(recheckInterval),

		cpVoteChan: make(chan *checkpointVote, cpChanBufferSize),
	}
	go r.listenDownloader()
	return r, nil
}

func (m *MasternodeService) Protocols() []p2p.Protocol {
	return nil
}

func (m *MasternodeService) APIs() []rpc.API {
	return nil
}

func (m *MasternodeService) Start(server *p2p.Server) error {
	address := crypto.PubkeyToAddress(server.PrivateKey.PublicKey)
	m.address = address

	//---
	m.eth.TxPool().RemoveBySender(address)

	//---
	contract, err := energi_abi.NewIMasternodeRegistryV2(
		energi_params.Energi_MasternodeRegistry, m.eth.APIBackend)
	if err != nil {
		return err
	}

	m.registry = &energi_abi.IMasternodeRegistryV2Session{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From: address,
		},
		TransactOpts: bind.TransactOpts{
			From: address,
			Signer: func(
				signer types.Signer,
				addr common.Address,
				tx *types.Transaction,
			) (*types.Transaction, error) {
				if addr != address {
					log.Error("Invalid Masternode address!", "addr", addr)
					return nil, errors.New("Invalid MN address")
				}

				return types.SignTx(tx, signer, server.PrivateKey)
			},
			Value:    common.Big0,
			GasLimit: masternodeCallGas,
			GasPrice: common.Big0,
		},
	}

	//---

	cpContract, err := energi_abi.NewICheckpointRegistry(
		energi_params.Energi_CheckpointRegistry, m.eth.APIBackend)
	if err != nil {
		return err
	}

	m.cpRegistry = &energi_abi.ICheckpointRegistrySession{
		Contract:     cpContract,
		CallOpts:     m.registry.CallOpts,
		TransactOpts: m.registry.TransactOpts,
	}

	m.server = server
	m.validator = newPeerValidator(common.Address{}, m)
	go m.loop()

	log.Info("Started Energi Masternode", "addr", address)
	return nil
}

func (m *MasternodeService) Stop() error {
	log.Info("Shutting down Energi Masternode", "addr", m.address)
	m.validator.cancel()
	return nil
}

func (m *MasternodeService) listenDownloader() {
	events := m.eth.EventMux().Subscribe(
		downloader.StartEvent{},
		downloader.DoneEvent{},
		downloader.FailedEvent{},
	)
	defer events.Unsubscribe()

	for ev := range events.Chan() {
		if ev == nil {
			return
		}
		switch ev.Data.(type) {
		case downloader.StartEvent:
			atomic.StoreInt32(&m.inSync, 0)
			log.Debug("Masternode is not in sync")
		case downloader.DoneEvent, downloader.FailedEvent:
			atomic.StoreInt32(&m.inSync, 1)
			log.Debug("Masternode is in sync")
			return
		}
	}
}

func (m *MasternodeService) isActive() bool {
	if atomic.LoadInt32(&m.inSync) == 0 {
		return false
	}

	if m.owner != (common.Address{}) {
		mninfo, err := m.registry.Info(m.address)
		if err != nil {
			log.Error("Masternode info fetch Err: %v", err)
			return false
		}

		if mninfo.Owner != m.owner {
			log.Error("Masternode owner mismatch", " needed=", m.owner, " got=", mninfo.Owner)
			return false
		}
	}

	res, err := m.registry.IsActive(m.address)

	if err != nil {
		log.Error("Masternode check failed", "err", err)
		return false
	}

	return res
}

func (m *MasternodeService) loop() {
	bc := m.eth.BlockChain()

	chainHeadCh := make(chan core.ChainHeadEvent, chainHeadChanSize)
	headSub := bc.SubscribeChainHeadEvent(chainHeadCh)
	defer headSub.Unsubscribe()

	events := m.eth.EventMux().Subscribe(
		CheckpointProposalEvent{},
	)
	defer events.Unsubscribe()

	//---
	for {
		select {
		case ev := <-events.Chan():
			if ev == nil {
				return
			}
			switch ev.Data.(type) {
			case CheckpointProposalEvent:
				m.onCheckpoint(ev.Data.(CheckpointProposalEvent))
			}
			break

		case ev := <-chainHeadCh:
			m.onChainHead(ev.Block)
			break

		// Shutdown
		case <-headSub.Err():
			return
		}
	}
}

func (m *MasternodeService) onCheckpoint(cpe CheckpointProposalEvent) {
	cpAddr := cpe.Proposal

	cp, err := energi_abi.NewICheckpointV2Caller(cpAddr, m.eth.APIBackend)
	if err != nil {
		log.Error("Failed to create the checkpoint iface", "cp", cpAddr, "err", err)
		return
	}

	callOpts := &m.cpRegistry.CallOpts

	// Check if the current masternode has voted and vote on it if not yet.
	can_vote, err := cp.CanVote(callOpts, m.address)
	if err != nil {
		log.Warn("Failed at Checkpoint.canVote()", "cp", cpAddr, "err", err)
		return
	}

	if !can_vote {
		return
	}

	if h := m.eth.BlockChain().GetHeaderByNumber(cpe.Number); h == nil {
		log.Warn("Checkpoint Proposal is ahead of this MN blockchain",
			"number", cpe.Number, "cp", cpe.Hash)
		return
	} else if h.Hash() != cpe.Hash {
		log.Warn("Checkpoint Proposal is not aligned with this MN blockchain",
			"number", cpe.Number, "header", h.Hash(), "cp", cpe.Hash)
		return
	}

	log.Info("MN checkpoint signature not found, now generating a new one", "cp", cpAddr)

	baseHash, err := cp.SignatureBase(callOpts)
	if err != nil {
		log.Error("Failed to get base hash", "cp", cpAddr, "err", err)
		return
	}

	signature, err := crypto.Sign(baseHash[:], m.server.PrivateKey)
	if err != nil {
		log.Error("Failed to sign base hash", "cp", cpAddr, "err", err)
		return
	}

	signature[64] += 27

	m.cpVoteChan <- &checkpointVote{
		address:   cpAddr,
		signature: signature,
	}
}

// voteOnCheckpoints recieves the identified checkpoints vote information and
// attempts to vote them in.
func (m *MasternodeService) voteOnCheckpoints() {
	var cpVote *checkpointVote

	for {
		select {
		case cpVote = <-m.cpVoteChan:
			// Only vote on the latest due to zero fee triggers
			break

		default:
			if cpVote != nil {
				tx, err := m.cpRegistry.Sign(cpVote.address, cpVote.signature)
				if tx != nil {
					txhash := tx.Hash()
					log.Warn("Voting on checkpoint", "addr", cpVote.address, "tx", txhash.Hex())
				}

				if err != nil {
					log.Error("Checkpoint vote failed", "checkpoint", cpVote.address, "err", err)
				}
			}

			return
		}
	}
}

func (m *MasternodeService) onChainHead(block *types.Block) {
	if !m.isActive() {
		do_cleanup := m.validator.target != common.Address{}
		m.validator.cancel()

		if do_cleanup {
			m.eth.TxPool().RemoveBySender(m.address)
		}
		return
	}

	// MN-4 - Heartbeats
	now := time.Now()

	if now.After(m.nextHB) {
		// It is more important than invalidation duty.
		// Some chance of race is still left, but at acceptable probability.
		m.validator.cancel()

		// Only present in IMasternodeRegistryV2
		if ok, err := m.registry.CanHeartbeat(m.address); err == nil && !ok {
			return
		}

		// Ensure heartbeat on clean queue
		if !m.eth.TxPool().RemoveBySender(m.address) {
			current := m.eth.BlockChain().CurrentHeader()
			tx, err := m.registry.Heartbeat(current.Number, current.Hash(), m.features)

			if err == nil {
				log.Info("Masternode Heartbeat", "tx", tx.Hash())
				m.nextHB = now.Add(heartbeatInterval)
			} else {
				log.Error("Failed to send Masternode Heartbeat", "err", err)
				m.nextHB = now.Add(recheckInterval)
			}
		} else {
			// NOTE: we need to recover from Nonce mismatch to enable heartbeats
			//       as soon as possible.
			log.Warn("Delaying Masternode Heartbeat due to pending zero-fee tx")
		}

		return
	}

	// Vote on the identified checkpoints.
	m.voteOnCheckpoints()

	//
	target, err := m.registry.ValidationTarget(m.address)
	if err != nil {
		log.Warn("MNTarget error", "mn", m.address, "err", err)
		m.validator.cancel()
		return
	}

	// MN-14: validation duty
	if old_target := m.validator.target; old_target != target {
		m.validator.cancel()
		m.validator = newPeerValidator(target, m)

		// Only present in IMasternodeRegistryV2
		if ok, err := m.registry.CanInvalidate(m.address); err == nil && !ok {
			return
		}

		// Skip the first validation cycle to prevent possible DoS trigger on restart
		if (old_target != common.Address{}) {
			go m.validator.validate()
		}
	}
}

type peerValidator struct {
	target   common.Address
	mnsvc    *MasternodeService
	cancelCh chan struct{}
}

func newPeerValidator(
	target common.Address,
	mnsvc *MasternodeService,
) *peerValidator {
	return &peerValidator{
		target:   target,
		mnsvc:    mnsvc,
		cancelCh: make(chan struct{}),
	}
}

func (v *peerValidator) cancel() {
	if v.mnsvc != nil {
		close(v.cancelCh)
		v.mnsvc = nil
	}
}

func (v *peerValidator) validate() {
	log.Debug("Masternode validation started", "target", v.target.Hex())
	defer log.Debug("Masternode validation stopped", "target", v.target.Hex())

	mnsvc := v.mnsvc
	if mnsvc == nil {
		return
	}
	server := mnsvc.server

	//---
	mninfo, err := mnsvc.registry.Info(v.target)
	if err != nil {
		log.Warn("MNInfo error", "mn", v.target, "err", err)
		return
	}

	cfg := mnsvc.eth.BlockChain().Config()
	enode := energi_common.MastenodeEnode(mninfo.Ipv4address, mninfo.Enode, cfg)

	if enode == nil {
		log.Debug("Invalid ipv4address or public key was used")
		return
	}

	// Check if the node is already connected as a peer and
	// skip MN Validation if true.
	if isFound := server.IsPeerActive(enode); isFound {
		log.Debug("Masternode validation skipped since peer is already connected",
			"target", v.target.Hex())
		return
	}

	peerCh := make(chan *p2p.PeerEvent)
	defer close(peerCh)

	peerSub := server.SubscribeEvents(peerCh)

	// EnableMsg Events.
	server.EnableMsgEvents = true
	defer peerSub.Unsubscribe()

	server.AddPeer(enode)

	defer func() {
		// Disconnect this peer if more than half of the max peers are connected.
		if server.PeerCount() > server.MaxPeers/2 {
			server.RemovePeer(enode)
		}
	}()

	//---
	deadline := time.Now().Add(time.Minute)

	for {
		select {
		case <-v.cancelCh:
			return
		case pe := <-peerCh:
			if pe.Peer != enode.ID() || pe.Type != p2p.PeerEventTypeMsgRecv {
				break
			}
			// TODO: validate block availability as per MN-14
			return
		case <-time.After(time.Until(deadline)):
			log.Info("MN Invalidation", "mn", v.target)

			// TODO: an excepted, but not seen before problem on scale got identified.
			log.Warn("Invalidations are temporarily disabled.")
			return

			//_, err := mnsvc.registry.Invalidate(v.target)
			//if err != nil {
			//	log.Warn("MN Invalidate error", "mn", v.target, "err", err)
			//}
			//return
		}
	}
}
