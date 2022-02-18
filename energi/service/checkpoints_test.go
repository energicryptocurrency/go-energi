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
	"crypto/ecdsa"
	"fmt"
	"net"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/energicryptocurrency/energi/accounts/keystore"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/crypto"
	"github.com/energicryptocurrency/energi/eth"

	// "github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/node"
	"github.com/energicryptocurrency/energi/p2p"
	"github.com/energicryptocurrency/energi/p2p/nat"
	"github.com/energicryptocurrency/energi/params"
	"github.com/stretchr/testify/assert"

	energi_testutils "github.com/energicryptocurrency/energi/energi/common/testutils"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

var sentCheckPoints = &testCheckPoints{}

type testCheckPoints struct {
	mtx sync.RWMutex
	cps []*core.CheckpointInfo
}

func (c *testCheckPoints) add(info *core.CheckpointInfo) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if len(c.cps) == 0 {
		c.cps = []*core.CheckpointInfo{info}
		return
	}

	// check for duplicates
	for _, oldCps := range c.cps {
		if reflect.DeepEqual(oldCps, info) {
			return
		}
	}
	c.cps = append(c.cps, info)
}

func (c *testCheckPoints) find(cp core.Checkpoint) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	for _, oldCps := range c.cps {
		if oldCps.Checkpoint.Number == cp.Number && oldCps.Checkpoint.Hash == cp.Hash {
			return true
		}
	}
	return false
}

func TestCheckpointsService(t *testing.T) {
	// log.Root().SetHandler(log.StdoutHandler)

	// initialize tx Description Map
	txDesc = txDescription{
		descMap: make(map[common.Hash]string),
	}

	withErr := func(msg string, err error) {
		if err != nil {
			panic(fmt.Errorf("%v error: %v", msg, err))
		}
	}

	delegatedPOS := []common.Address{
		energi_params.Energi_MigrationContract,
		params.EnergiTestnetChainConfig.Energi.CPPSigner,
	}

	nodesInfo = make([]nodeConfig, 0, totalNetworkNodes)
	signers = make(map[common.Address]*ecdsa.PrivateKey, totalNetworkNodes)
	mnAddrToOwners = make(map[common.Address]*ecdsa.PrivateKey, 2)
	allocs := core.DefaultPrealloc()

	// generate private keys for all nodes.
	for index := 0; index < totalNetworkNodes; index++ {
		key, accAddr := accountGen()
		signers[accAddr] = key
		allocs[accAddr] = core.GenesisAccount{Balance: balance}

		var isMasternode bool
		// select masternodes
		switch index {
		case 0, 1, 2, 3, 4: // accounts at indexes 0 to 4 are masternodes.
			isMasternode = true

			// Create the masternode owners and pre-assign them a balance.
			mnOwnerKey, mnOwnerAddr := accountGen()
			mnAddrToOwners[accAddr] = mnOwnerKey
			allocs[mnOwnerAddr] = core.GenesisAccount{Balance: balance}

		default: // rest of the account belong to enodes.
			isMasternode = false
		}

		nodesInfo = append(nodesInfo, nodeConfig{
			isMN:    isMasternode,
			address: accAddr,
		})
	}

	delPoSKeys := make([]*ecdsa.PrivateKey, 0, len(delegatedPOS))
	delPoSAddr := make([]common.Address, 0, len(delegatedPOS))
	// Map signer addresses to existing node private keys for signer accounts.
	for _, addr := range delegatedPOS {
		privKey, accAddr := accountGen()
		allocs[accAddr] = core.GenesisAccount{Balance: balance}

		delPoSAddr = append(delPoSAddr, accAddr)
		delPoSKeys = append(delPoSKeys, privKey)

		switch addr {
		case energi_params.Energi_MigrationContract:
			mgSigner = accAddr

		case params.EnergiTestnetChainConfig.Energi.CPPSigner:
			cpSigner = accAddr
		}
	}

	for index := 0; index < totalNetworkNodes; index++ {
		var err error
		var node *node.Node
		nConfig := nodesInfo[index]
		key := signers[nConfig.address]

		switch nConfig.isMN {
		case true:
			node, err = energiServices(key, allocs)

		default: // rest of the account belong to enodes.
			node, err = newNode(key, allocs)
		}

		msg := fmt.Sprintf("Creating node with Address: %v failed", nConfig.address.Hash().String())
		withErr(msg, err)

		// Now assign the node to the node config.
		nodesInfo[index].stack = node
	}

	// Add the delegetedPoS Addresses to the signer map
	for i, addr := range delPoSAddr {
		signers[addr] = delPoSKeys[i]
	}

	// Add the masternode owners
	for _, ownerKey := range mnAddrToOwners {
		ownerAddr := crypto.PubkeyToAddress(ownerKey.PublicKey)
		signers[ownerAddr] = ownerKey
	}

	migrations := energi_testutils.NewTestGen2Migration()
	// Create a gen2 migration tempfile
	err := migrations.PrepareTestGen2Migration(params.EnergiTestnetChainConfig.ChainID.Uint64())
	withErr("Creating the Gen2 snapshot failed", err)

	migrationFile = migrations.TempFileName()

	injectAccount := func(store *keystore.KeyStore, privKey *ecdsa.PrivateKey) {
		account, err := store.ImportECDSA(privKey, accountPass)
		withErr("Failed to Inject new account", err)

		// Unlock the account for staking
		err = store.Unlock(account, accountPass, true)
		withErr("Failed to Unlock new account for staking", err)
	}

	// Boot up the entire protocol while importing the accounts into respective nodes.
	for _, data := range nodesInfo {
		err = data.stack.Start()
		withErr("Failed to start the protocol stack", err)

		srv := data.stack.Server()
		addr, _ := net.ResolveUDPAddr("udp", srv.ListenAddr)
		conn, _ := net.ListenUDP("udp", addr)
		realAdr := conn.LocalAddr().(*net.UDPAddr)
		quit := make(chan struct{})
		if !realAdr.IP.IsLoopback() && srv.NAT != nil {
			go nat.Map(srv.NAT, quit, "udp", realAdr.Port, realAdr.Port, "ethereum discovery")
		}

		// trigger external IP Address to be set.
		_, _ = srv.NAT.ExternalIP()

		var ethService *eth.Ethereum
		_ = data.stack.Service(&ethService)

		store := data.stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
		// inject the main node personal account
		injectAccount(store, signers[data.address])

		// Add delegated POS accounts to every node.
		for i, addr := range delPoSAddr {
			injectAccount(store, signers[addr])

			contractAddr := delegatedPOS[i]
			ethService.AddDPoS(contractAddr, crypto.PubkeyToAddress(signers[addr].PublicKey))
		}
	}

	listenToCheckpointsTest(t)

	// Clean Up
	_ = migrations.CleanUp()

	// Stop the entire protocol for all nodesInfo.
	for _, data := range nodesInfo {
		err = data.stack.Stop()
		withErr("Failed to stop the protocol stack", err)
	}
}

// networkEvents receives all new changes that are mdae to the network.
func networkEvents(
	quitCh chan struct{},
	isSignedCPP chan struct{},
	ethService *eth.Ethereum,
	cpService *CheckpointService,
) {
	bc := ethService.BlockChain()
	txpool := ethService.TxPool()

	chainHeadCh := make(chan core.ChainHeadEvent, chainHeadChanSize)
	headSub := bc.SubscribeChainHeadEvent(chainHeadCh)
	defer headSub.Unsubscribe()

	txEventCh := make(chan core.NewTxsEvent, 10)
	txSub := txpool.SubscribeNewTxsEvent(txEventCh)
	defer txSub.Unsubscribe()

	evt := cpService.eth.EventMux().Subscribe(CheckpointProposalEvent{})
	defer evt.Unsubscribe()

	//---
	for {
		select {
		case <-quitCh:
			return
		case ev := <-chainHeadCh:
			fmt.Println(" _____ New Block Mined _____")

			for _, tx := range ev.Block.Transactions() {
				fmt.Printf("\t BlockNo: %v, Tx Desc: %v, Tx Hash: %v, Nonce: %v, GasPrice: %v, Gas: %v, To Address: %v \n",
					ev.Block.Number(), checkTxDesc(tx), tx.Hash().String(), tx.Nonce(), tx.GasPrice(), tx.Gas(), tx.To().Hash().String())
			}
		case txEvent := <-txEventCh:
			for _, tx := range txEvent.Txs {
				fmt.Printf("\t\t _____ (%s) Tx Announced  %v _____ \n", checkTxDesc(tx), tx.Hash().String())
			}

		case ev := <-evt.Chan():
			if ev == nil {
				return
			}
			switch ev.Data.(type) {
			case CheckpointProposalEvent:
				if sentCheckPoints.find(ev.Data.(CheckpointProposalEvent).Checkpoint) {
					isSignedCPP <- struct{}{}
				}
			}

			break

		// Shutdown
		case <-headSub.Err():
			return
		case <-txSub.Err():
			return
		}
	}
}

func listenToCheckpointsTest(t *testing.T) {
	miningTimeout := time.After(miningInterval)

	// masternode mn node picked is at index 1.
	mn := nodesInfo[mnIndex]
	var mnEthService *eth.Ethereum
	_ = mn.stack.Service(&mnEthService)

	var cpServ *CheckpointService
	_ = mn.stack.Service(&cpServ)

	quitChan := make(chan struct{}, 1)
	isCPPChan := make(chan struct{}, 1)
	// Listen to the network events
	go networkEvents(quitChan, isCPPChan, mnEthService, cpServ)

	mnServer := mn.stack.Server()
	peerCh := make(chan *p2p.PeerEvent)
	peerSub := mnServer.SubscribeEvents(peerCh)
	defer close(peerCh)
	defer peerSub.Unsubscribe()

	// EnableMsg Events.
	mnServer.EnableMsgEvents = true

	fmt.Println(" _______ ADDING PEERS _____")
	var peers int
	// Add all nodes as peers then start mining in each peer
	for count, data := range nodesInfo {
		if count == mnIndex {
			// Do not add the main masternode as a peer to itself.
			continue
		}

		mnServer.AddPeer(data.stack.Server().Self())

		// This is a blocking operation that requires all nodes to be fully added
		// as peers before further progress can be made.
	peerConLoop:
		for {
			select {
			case event := <-peerCh:
				if event.Type == p2p.PeerEventTypeMsgRecv {
					// Allow some delay for the peer to sync.
					time.Sleep(peerSyncDelay)
					break peerConLoop
				}
			case <-time.After(peerConInterval):
				t.Fatal(errTimeout)
				break peerConLoop
			}
		}

		peers++
	}

	// On subscription, peerCh has to always be read when full to allow other txs
	// to be announced.
	go func() {
	waitLoop:
		for {
			select {
			case <-peerCh:
			case <-quitChan:
				break waitLoop
			}
		}
	}()

	// Confirm that all the peers were added to the network.
	assert.Equal(t, peers, mnServer.PeerCount())

	fmt.Println(" _______ START MINING _____")
	// Add all nodes as peers then start mining in each peer
	for _, data := range nodesInfo {
		var ethService *eth.Ethereum
		_ = data.stack.Service(&ethService)

		go func() {
			err := ethService.StartMining(2)
			assert.Equal(t, nil, err)
			if err != nil {
				return
			}

			// If shutting down, exit this goroutine.
			for range quitChan {
				return
			}
		}()
	}

	fmt.Println(" _______ ACTIVATE MASTERNODES _____")
	err := mnPrepare(nodesInfo)
	assert.Equal(t, nil, err)

	// The cpp signer node proposes a checkpoint.
	fmt.Println(" _______ PROPOSE CHECKPOINT-1 _____")
	cpInfo, err := cpPropose()
	sentCheckPoints.add(cpInfo)
	assert.Equal(t, nil, err)

	// Send more txs.
	fmt.Println(" _______ SEND MORE TXS _____")
	err = sendMoreTxs(mn.stack.Server().PrivateKey)
	assert.Equal(t, nil, err)

	cppsigner := nodesInfo[cpSignerIndex]
	var ethServ *eth.Ethereum
	_ = cppsigner.stack.Service(&ethServ)

	fmt.Println(" _______ CHECK TX POOL BEFORE WAITING _____")
	{
		// Tx pool according to the main masternode before test.
		pending, queued := ethServ.TxPool().Content()
		txPoolContents(pending, "(BEFORE) ___ CPP Signer Pending")
		txPoolContents(queued, "(BEFORE) ____ CPP Signer queued")
	}

	{
		// Tx pool according to the cpp signer node before test.
		pending, queued := mnEthService.TxPool().Content()
		txPoolContents(pending, "(BEFORE) ____ MN Pending")
		txPoolContents(queued, "(BEFORE) ____ MN queued")
	}

	fmt.Println(" _______ PROPOSE CHECKPOINT-2 _____")
	cpInfo, err = cpPropose()
	sentCheckPoints.add(cpInfo)
	assert.Equal(t, nil, err)

	// Wait for a checkpoint signed by a masternode to be discovered or the max
	// mining interval to expire.

	select {
	case <-isCPPChan:
		// Test Passed
		fmt.Println(" _______ A checkpoint event by the checkpoint service was found _____")

	case <-miningTimeout:
		// Test Failed
		t.Fatalf(" _______ Checkpoint event NOT found: checkpoint service failed to send event on time _____")
	}

	// Now quit the listening of network events.
	quitChan <- struct{}{}

	fmt.Println(" _______ CHECK TX POOL AFTER WAITING _____")
	{
		// Tx pool according to the main masternode after test.
		pending, queued := mnEthService.TxPool().Content()
		txPoolContents(pending, "(AFTER) ____ MN Pending")
		txPoolContents(queued, "(AFTER) ____ MN queued")
	}

	{
		// Tx pool according to the cpp signer node after test.
		pending, queued := ethServ.TxPool().Content()
		txPoolContents(pending, "(AFTER) ____ CPP Signer Pending")
		txPoolContents(queued, "(AFTER) ____ CPP Signer queued")
	}

	fmt.Println(" >>>>>>>> Checkpoint Service Test Complete <<<<<<<<<<")
}
