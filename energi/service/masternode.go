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
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	heartbeatInterval uint64 = 35 * 60
	recheckInterval   uint64 = 5 * 60
	masternodeCallGas uint64 = 500000

	txChanSize        = 4096
	chainHeadChanSize = 10
)

type MasternodeService struct {
	server *p2p.Server
	eth    *eth.Ethereum

	quitCh chan struct{}

	address  common.Address
	registry *energi_abi.IMasternodeRegistrySession
	lastHB   uint64
	features *big.Int
}

func NewMasternodeService(ethServ *eth.Ethereum) (node.Service, error) {
	return &MasternodeService{
		eth:      ethServ,
		quitCh:   make(chan struct{}),
		features: big.NewInt(0),
	}, nil
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

	m.server = server
	go m.loop()

	log.Info("Started Energi Masternode", "addr", address)
	return nil
}

func (m *MasternodeService) Stop() error {
	return nil
}

func (m *MasternodeService) isActive() bool {
	res, err := m.registry.IsValid(m.address)

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
		now := uint64(time.Now().Unix())
		nextHB := m.lastHB + heartbeatInterval

		if now > nextHB {
			if m.isActive() {
				current := bc.CurrentHeader()
				tx, err := m.registry.Heartbeat(current.Number, current.Hash(), m.features)

				if err == nil {
					log.Info("Masternode Heartbeat", "tx", tx.Hash())
					m.lastHB = now
					nextHB = now + heartbeatInterval
				} else {
					log.Error("Failed to send Masternode Heartbeat", "err", err)
					nextHB = now + recheckInterval
				}
			} else {
				log.Error("Masternode is not active!")
				nextHB = now + recheckInterval
			}
		}

		select {
		case <-m.quitCh:
			return
		case <-chainHeadCh:
			break
		case <-txEventCh:
			break
		case <-time.After(time.Duration(nextHB-now) * time.Second):
			break

		// Shutdown
		case <-headSub.Err():
			return
		case <-txSub.Err():
			return
		}
	}
}
