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

package eth

import (
	"errors"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

var (
	heartbeatInterval = time.Duration(35) * time.Minute
	recheckInterval   = time.Duration(5) * time.Minute

	// errInvalidSubscription is returned if valid subscription cannot be established.
	errInvalidSubscription = errors.New("Invalid subscription")
)

const (
	masternodeCallGas uint64 = 500000

	// cpChanBufferSize defines the number of checkpoint to be pushed into the
	// checkpoints channel before it can be considered to be full.
	cpChanBufferSize = 500

	txChanSize        = 4096
	chainHeadChanSize = 10
)

type checkpointVote struct {
	address   common.Address
	signature []byte
}

type MasternodeService struct {
	server *p2p.Server
	eth    *eth.Ethereum

	quitCh chan struct{}
	inSync int32

	address  common.Address
	registry *energi_abi.IMasternodeRegistrySession

	cpRegistry  *energi_abi.ICheckpointRegistrySession
	lastCPBlock uint64
	cpVoteChan  chan *checkpointVote

	nextHB   time.Time
	features *big.Int

	validator *peerValidator
}

func NewMasternodeService(ethServ *eth.Ethereum) (node.Service, error) {
	r := &MasternodeService{
		eth:      ethServ,
		quitCh:   make(chan struct{}),
		inSync:   1,
		features: big.NewInt(0),
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
	contract, err := energi_abi.NewIMasternodeRegistry(
		energi_params.Energi_MasternodeRegistry, m.eth.APIBackend)
	if err != nil {
		return err
	}

	m.registry = &energi_abi.IMasternodeRegistrySession{
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

	// fetch all checkpoints since block number zero.
	go m.getMNCheckpoint()

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

	for {
		select {
		case ev := <-events.Chan():
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
		case <-m.quitCh:
			return
		}
	}
}

func (m *MasternodeService) isActive() bool {
	if atomic.LoadInt32(&m.inSync) == 0 {
		return false
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
	txpool := m.eth.TxPool()

	chainHeadCh := make(chan core.ChainHeadEvent, chainHeadChanSize)
	headSub := bc.SubscribeChainHeadEvent(chainHeadCh)
	defer headSub.Unsubscribe()

	txEventCh := make(chan core.NewTxsEvent, txChanSize)
	txSub := txpool.SubscribeNewTxsEvent(txEventCh)
	defer txSub.Unsubscribe()

	//---
	for {
		select {
		case <-m.quitCh:
			return
		case ev := <-chainHeadCh:
			m.onChainHead(ev.Block)
			break
		case <-txEventCh:
			break

		// Shutdown
		case <-headSub.Err():
			return
		case <-txSub.Err():
			return
		}
	}
}

// getMNCheckpoint identifies the checkpoint indexed by the lastCPBlock value.
// The identified checkpoints are signed and the pushed to the cpVoteChan for
// further voting after heartbeat operations finish.
func (m *MasternodeService) getMNCheckpoint() error {
	log.Info("Running checkpoints fetch from", "block", m.lastCPBlock)

	cpChan := make(chan *energi_abi.ICheckpointRegistryCheckpoint, cpChanBufferSize)
	watchOpts := &bind.WatchOpts{Start: &m.lastCPBlock}

	subscribe, err := m.cpRegistry.Contract.WatchCheckpoint(watchOpts, cpChan, []*big.Int{})
	if err != nil {
		log.Error("Failed checkpoint subscription", "block", m.lastCPBlock, "err", err)

		return err
	}

	defer func() {
		subscribe.Unsubscribe()
		log.Debug("Exiting checkpoint subscription", "block", m.lastCPBlock, "err", err)
	}()

	for {
		select {
		case <-m.quitCh:
			return nil

		case err = <-subscribe.Err():
			log.Error("checkpoint subscription error", "err", err)
			return err

		case cpData := <-cpChan:
			cpAddr := cpData.Checkpoint
			cp, err := energi_abi.NewICheckpointCaller(cpAddr, m.eth.APIBackend)
			if err != nil {
				log.Error("Failed to retrieve the checkpoint", "err", err)
				continue
			}

			// Update the current checkpoint height.
			m.lastCPBlock = cpData.Number.Uint64()

			// Check if the current masternode has voted and vote on it if not yet.
			_, err = cp.Signature(&m.cpRegistry.CallOpts, m.address)
			if err != nil {
				log.Debug("MN signature fetch failed", "err", err)

				baseHash, err := m.cpRegistry.SignatureBase(cpData.Number, cpAddr.Hash())
				if err != nil {
					log.Error("Failed to get base hash", "err", err)
					continue
				}

				signature, err := crypto.Sign(baseHash[:], m.server.PrivateKey)
				if err != nil {
					log.Error("Failed to sign base hash", "err", err)
					continue
				}

				// Push the processed signature and the checkpoint address to channel
				// for later processing.
				m.cpVoteChan <- &checkpointVote{
					address:   cpAddr,
					signature: signature,
				}
			}

			log.Debug("MN has already voted", "MN", m.address, "checkpoint", cpAddr)
		}
	}
}

// voteOnCheckpoints recieves the identified checkpoints vote information and
// attempts to vote them in.
func (m *MasternodeService) voteOnCheckpoints() {
	for {
		select {
		case <-m.quitCh:
			return

		case cpVote := <-m.cpVoteChan:
			tx, err := m.cpRegistry.Sign(cpVote.address, cpVote.signature)
			if tx != nil {
				txhash := tx.Hash()
				log.Info("Note: please wait until the vote TX gets into a block!", "tx", txhash.Hex())
			}

			if err != nil {
				log.Error("Checkpoint vote failed", "checkpoint", cpVote.address, "err", err)
			}

		default:
			// This triggers loop ending to allow the priority zero-fee txs to be
			// processed before attempting checkpoint voting again.
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
		case <-mnsvc.quitCh:
			return
		case <-v.cancelCh:
			return
		case pe := <-peerCh:
			if pe.Peer != enode.ID() || pe.Type != p2p.PeerEventTypeMsgRecv {
				break
			}
			// TODO: validate block availability as per MN-14
			return
		case <-time.After(deadline.Sub(time.Now())):
			if mnsvc.eth.TxPool().RemoveBySender(mnsvc.address) {
				log.Warn("Skipping MN invalidation due to tx queue", "mn", v.target)
				return
			}

			log.Info("MN Invalidation", "mn", v.target)
			_, err := mnsvc.registry.Invalidate(v.target)
			if err != nil {
				log.Warn("MN Invalidate error", "mn", v.target, "err", err)
			}
			return
		}
	}
}
