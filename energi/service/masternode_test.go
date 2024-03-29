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
	"errors"
	"fmt"
	"math/big"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/energicryptocurrency/go-energi/accounts/abi/bind"
	"github.com/energicryptocurrency/go-energi/accounts/keystore"
	"github.com/energicryptocurrency/go-energi/common"
	"github.com/energicryptocurrency/go-energi/core"
	"github.com/energicryptocurrency/go-energi/core/types"
	"github.com/energicryptocurrency/go-energi/crypto"
	energi_abi "github.com/energicryptocurrency/go-energi/energi/abi"
	energi_testutils "github.com/energicryptocurrency/go-energi/energi/common/testutils"
	energi_params "github.com/energicryptocurrency/go-energi/energi/params"
	"github.com/energicryptocurrency/go-energi/eth"
	"github.com/energicryptocurrency/go-energi/node"
	"github.com/energicryptocurrency/go-energi/p2p"
	"github.com/energicryptocurrency/go-energi/p2p/nat"
	"github.com/energicryptocurrency/go-energi/params"

	"github.com/stretchr/testify/assert"
)

const (
	mnIndex           = 0 // Main Masternode Index. Index 0-4 are occupied by masternodes.
	cpSignerIndex     = 5 // Checkpoint signer node index.
	txsSenderIndex    = 6 // Sends the extra regular txs.
	totalNetworkNodes = 7 // Network has this nodes in total.

	accountPass     = "secret-pass"
	peerConInterval = 1 * time.Minute
	miningInterval  = 5 * time.Minute // Max interval which the test can run before its seen as failed.
	peerSyncDelay   = 20 * time.Second

	gasLimit          = 40000000
	mnPrepareGasLimit = 3900000
)

var (
	migrationFile  string
	nodesInfo      []nodeConfig
	signers        map[common.Address]*ecdsa.PrivateKey
	mnAddrToOwners map[common.Address]*ecdsa.PrivateKey

	mgSigner, cpSigner common.Address

	gasPrice   = big.NewInt(10000000000)
	collateral = new(big.Int).Mul(big.NewInt(100000), big.NewInt(1e18))
	balance    = new(big.Int).Mul(big.NewInt(1000000), big.NewInt(1e18))

	errTimeout = errors.New("time expired")

	txDesc txDescription
)

type nodeConfig struct {
	isMN    bool
	stack   *node.Node
	address common.Address
}

// txDescription helps to identify by description the respective txs mined.
type txDescription struct {
	mtx     sync.RWMutex
	descMap map[common.Hash]string
}

// addTxDesc adds a description to the tx description map.
func addTxDesc(tx *types.Transaction, desc string) {
	txDesc.mtx.Lock()
	defer txDesc.mtx.Unlock()

	txDesc.descMap[tx.Hash()] = desc
}

// checkTxDesc checks if the provided tx has its description set.
func checkTxDesc(tx *types.Transaction) string {
	txDesc.mtx.RLock()
	defer txDesc.mtx.RUnlock()

	desc := txDesc.descMap[tx.Hash()]
	if desc == "" {
		desc = "-------"
	}
	return desc
}

// accountGen returns the private key and address account information.
func accountGen() (*ecdsa.PrivateKey, common.Address) {
	privKey, _ := crypto.GenerateKey()
	walletAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	return privKey, walletAddr
}

func TestMasternodeService(t *testing.T) {
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

	masternodeCPSigningTest(t)

	// Clean Up
	_ = migrations.CleanUp()

	// Stop the entire protocol for all nodesInfo.
	for _, data := range nodesInfo {
		err = data.stack.Stop()
		withErr("Failed to stop the protocol stack", err)
	}
}

// serviceConfig generates the ethereum service configuration.
func serviceConfig(presale core.GenesisAlloc) *eth.Config {
	ethConfig := &eth.DefaultConfig
	ethConfig.MinerMigration = migrationFile
	ethConfig.MinerRecommit = 3 * time.Second
	ethConfig.Genesis = core.DefaultEnergiTestnetGenesisBlock()
	ethConfig.Genesis.GasLimit = gasLimit
	ethConfig.Genesis.Config.ChainID = big.NewInt(1)
	ethConfig.Genesis.Config.Energi.CPPSigner = cpSigner
	ethConfig.Genesis.Config.Energi.MigrationSigner = mgSigner
	ethConfig.Genesis.Alloc = presale
	ethConfig.Genesis.Xfers = core.DeployEnergiGovernance(ethConfig.Genesis.Config)
	ethConfig.Genesis.Difficulty = big.NewInt(1)
	ethConfig.Genesis.Coinbase = energi_params.Energi_Treasury
	ethConfig.Genesis.Timestamp = 12900000
	ethConfig.Genesis.Coinbase = mgSigner
	return ethConfig
}

// newNode create a new ethereum node service and registers it.
func newNode(privKey *ecdsa.PrivateKey, presale core.GenesisAlloc) (*node.Node, error) {
	config := &node.Config{
		P2P: p2p.Config{
			ListenAddr:  "0.0.0.0:0",
			NAT:         nat.Any(),
			NoDiscovery: true,
			MaxPeers:    25,
			PrivateKey:  privKey,
		},
		NoUSB:             true,
		UseLightweightKDF: true,
	}

	stack, err := node.New(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create network node: %v", err)
	}

	ethConstructor := func(ctx *node.ServiceContext) (node.Service, error) {
		return eth.New(ctx, serviceConfig(presale))
	}

	// Register the ethereum(energi pos engine) service
	if err := stack.Register(ethConstructor); err != nil {
		return nil, fmt.Errorf("Failed to register Ethereum service: %v", err)
	}

	return stack, nil
}

// energiServices creates a new network node to the protocol with the default
// values and registers the ethereum, masternode and checkpoint services into the network.
func energiServices(privKey *ecdsa.PrivateKey, presale core.GenesisAlloc) (*node.Node, error) {
	stack, err := newNode(privKey, presale)
	if err != nil {
		return nil, err
	}

	MNConstructor := func(ctx *node.ServiceContext) (node.Service, error) {
		var ethServ *eth.Ethereum
		if err := ctx.Service(&ethServ); err != nil {
			return nil, err
		}
		return NewMasternodeService(ethServ, common.Address{})
	}

	// Register the masternode service.
	if err := stack.Register(MNConstructor); err != nil {
		return nil, fmt.Errorf("Failed to register MN service: %v", err)
	}

	CPConstructor := func(ctx *node.ServiceContext) (node.Service, error) {
		var ethServ *eth.Ethereum
		if err := ctx.Service(&ethServ); err != nil {
			return nil, err
		}
		return NewCheckpointService(ethServ)
	}

	// Register the checkpoint service.
	if err := stack.Register(CPConstructor); err != nil {
		return nil, fmt.Errorf("Failed to register CP service: %v", err)
	}

	return stack, nil
}

func SignerCallback(node *node.Node, signerAddrs map[common.Address]*ecdsa.PrivateKey) bind.SignerFn {
	signerfunc := func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signerKey := signerAddrs[addr]

		var ethServ *eth.Ethereum
		_ = node.Service(&ethServ)
		chainID := ethServ.APIBackend.ChainConfig().ChainID

		return types.SignTx(tx, types.NewEIP155Signer(chainID), signerKey)
	}
	return signerfunc
}

// cpPropose uses the CPPSigner to propose a checkpoint to the network.
// It creates a checkpoint on the current header.
func cpPropose() (*core.CheckpointInfo, error) {
	signer := nodesInfo[cpSignerIndex]

	var ethServ *eth.Ethereum
	_ = signer.stack.Service(&ethServ)

	// Governed Proxy Contract
	cpRegistry, err := energi_abi.NewCheckpointRegistryV2(
		energi_params.Energi_CheckpointRegistry, ethServ.APIBackend,
	)
	if err != nil {
		return nil, err
	}

	callOpts := bind.CallOpts{
		From:     cpSigner,
		GasLimit: mnPrepareGasLimit,
	}

	block := ethServ.BlockChain().CurrentHeader()
	hashRaw, err := cpRegistry.SignatureBase(&callOpts, block.Number, block.Hash())
	if err != nil {
		return nil, err
	}

	signature, err := crypto.Sign(hashRaw[:], signers[cpSigner])
	if err != nil {
		return nil, err
	}

	signature[64] += 27

	transactOpts := bind.TransactOpts{
		From:     cpSigner,
		Signer:   SignerCallback(signer.stack, signers),
		Value:    common.Big0,
		GasLimit: mnPrepareGasLimit,
	}

	tx, err := cpRegistry.Propose(&transactOpts, block.Number, block.Hash(), signature)
	if tx != nil {
		addTxDesc(tx, "checkpoint")
	}

	cp := core.CheckpointInfo{
		Checkpoint: core.Checkpoint{
			Number: block.Number.Uint64(),
			Hash:   block.Hash(),
		},
		CppSignature: signature,
	}

	return &cp, err
}

func depositCollateral(
	index int,
	ethService *eth.Ethereum,
	transactOpts bind.TransactOpts,
) error {
	mnRegistry, err := energi_abi.NewIMasternodeTokenTransactor(
		energi_params.Energi_MasternodeToken, ethService.APIBackend)
	if err != nil {
		return err
	}

	transactOpts.Value = collateral // Set Collateral amount.

	tx, err := mnRegistry.DepositCollateral(&transactOpts)
	if err != nil {
		return fmt.Errorf("MN collateral deposit failed error: %v", err)
	}

	if tx != nil {
		addTxDesc(tx, fmt.Sprintf("collateral deposit-%v", index))
	}
	return nil
}

func announceMN(
	index int,
	ethService *eth.Ethereum,
	transactOpts bind.TransactOpts,
	mnAddr common.Address,
	ip net.IP,
) error {
	registry, err := energi_abi.NewIMasternodeRegistryV2Transactor(
		energi_params.Energi_MasternodeRegistry, ethService.APIBackend)
	if err != nil {
		return err
	}

	var ipv4address = uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
	if ip[0] == byte(127) || ip[0] == byte(10) ||
		(ip[0] == byte(192) && ip[1] == byte(168)) ||
		(ip[0] == byte(172) && (ip[1]&0xF0) == byte(16)) {

		// Use a mocked IP address if a proper external IP address cannot be obtained.
		ipv4address = uint32(130 << 24)
	}

	tx, err := registry.Announce(&transactOpts, mnAddr, ipv4address, [2][32]byte{})
	if err != nil {
		return fmt.Errorf("MN announcing failed error: %v", err)
	}
	if tx != nil {
		addTxDesc(tx, fmt.Sprintf("masternode announced-%v", index))
	}
	return nil
}

// mnPrepare deposits collateral and announces each masternode to the network.
func mnPrepare(nodes []nodeConfig) error {
	for i, n := range nodes {
		if !n.isMN {
			continue
		}

		var ethService *eth.Ethereum
		_ = n.stack.Service(&ethService)

		ownerAddr := crypto.PubkeyToAddress(mnAddrToOwners[n.address].PublicKey)
		transactOpts := bind.TransactOpts{
			From:     ownerAddr,
			Signer:   SignerCallback(n.stack, signers),
			Value:    common.Big0,
			GasLimit: mnPrepareGasLimit,
		}

		fmt.Printf("\t _______ DEPOSIT COLLATERAL for %v _____ \n", n.address.Hash().String())
		if err := depositCollateral(i, ethService, transactOpts); err != nil {
			return err
		}

		ip := n.stack.Server().Self().IP().To4()
		fmt.Printf("\t _______ ANNOUNCE %v TO THE NETWORK_____ \n", n.address.Hash().String())
		if err := announceMN(i, ethService, transactOpts, n.address, ip); err != nil {
			return err
		}
	}
	return nil
}

func sendMoreTxs(mnPrivKey *ecdsa.PrivateKey) error {
	fromNodeInfo := nodesInfo[txsSenderIndex]
	var ethService *eth.Ethereum
	_ = fromNodeInfo.stack.Service(&ethService)

	txsCount := 5
	txs := make([]*types.Transaction, 0, txsCount)
	fromKey := fromNodeInfo.stack.Server().PrivateKey
	toAddr := crypto.PubkeyToAddress(mnPrivKey.PublicKey)

	for i := 0; i < txsCount; i++ {
		chainID := ethService.APIBackend.ChainConfig().ChainID
		amount := new(big.Int)

		// send to the mn
		rawtx := types.NewTransaction(uint64(i), toAddr, amount, mnPrepareGasLimit, gasPrice, nil)
		tx, err := types.SignTx(rawtx, types.NewEIP155Signer(chainID), fromKey)
		if err != nil {
			return err
		}

		if tx != nil {
			addTxDesc(tx, fmt.Sprintf("regular tx-%v", i))
			txs = append(txs, tx)
		}
	}

	for _, err := range ethService.TxPool().AddRemotes(txs) {
		if err != nil {
			return err
		}
	}
	return nil
}

// networkEventsLoop receives all new changes that are mdae to the network.
func networkEventsLoop(
	quitCh chan struct{},
	isSignedCPP chan struct{},
	ethService *eth.Ethereum,
) {
	bc := ethService.BlockChain()
	txpool := ethService.TxPool()

	chainHeadCh := make(chan core.ChainHeadEvent, chainHeadChanSize)
	headSub := bc.SubscribeChainHeadEvent(chainHeadCh)
	defer headSub.Unsubscribe()

	txEventCh := make(chan core.NewTxsEvent, 10)
	txSub := txpool.SubscribeNewTxsEvent(txEventCh)
	defer txSub.Unsubscribe()

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

				// Waits until the tx from the mn signed checkpoint to be mined into a block.
				if isSignedCheckpointTx(tx) {
					isSignedCPP <- struct{}{}
				}
			}
		case txEvent := <-txEventCh:
			for _, tx := range txEvent.Txs {
				fmt.Printf("\t\t _____ (%s) Tx Announced  %v _____ \n", checkTxDesc(tx), tx.Hash().String())
			}

		// Shutdown
		case <-headSub.Err():
			return
		case <-txSub.Err():
			return
		}
	}
}

// Detect the checkpoint signed by the current masternode and check if its valid.
func isSignedCheckpointTx(tx *types.Transaction) bool {
	if *tx.To() == energi_params.Energi_CheckpointRegistry {
		for _, node := range nodesInfo {
			if !node.isMN {
				// expected only masternodes to be used
				continue
			}

			var ethServ *eth.Ethereum
			_ = node.stack.Service(&ethServ)

			signer := types.NewEIP155Signer(ethServ.BlockChain().Config().ChainID)
			from, err := types.Sender(signer, tx)
			if err != nil {
				continue
			}

			if node.address == from {
				// found the recovered address in one of the masternodes.
				fmt.Println("found the recovered address in one of the masternodes")
				return true
			}
		}
		// could not match the recovered address to a mn address: not required tx.
		fmt.Println("could not match the recovered address to a mn address: not required tx")
		return false
	}

	// invalid recipient address found: Not Required tx
	// fmt.Println("invalid recipient address found: Not Required tx")
	return false
}

func masternodeCPSigningTest(t *testing.T) {
	miningTimeout := time.After(miningInterval)

	// masternode mn node picked is at index 1.
	mn := nodesInfo[mnIndex]
	var mnEthService *eth.Ethereum
	_ = mn.stack.Service(&mnEthService)

	quitChan := make(chan struct{}, 1)
	isCPPChan := make(chan struct{}, 1)
	// Listen to the network events
	go networkEventsLoop(quitChan, isCPPChan, mnEthService)

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
	fmt.Println(" _______ PROPOSE CHECKPOINT _____")
	_, err = cpPropose()
	assert.Equal(t, nil, err)

	// Send more txs.
	fmt.Println(" _______ SEND MORE TXS _____")
	err = sendMoreTxs(mn.stack.Server().PrivateKey)
	assert.Equal(t, nil, err)

	fmt.Println(" _______ CHECK TX POOL BEFORE WAITING _____")
	cppsigner := nodesInfo[cpSignerIndex]
	var ethServ *eth.Ethereum
	_ = cppsigner.stack.Service(&ethServ)

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

	// Wait for a checkpoint signed by a masternode to be discovered or the max
	// mining interval to expire.

	select {
	case <-isCPPChan:
		// Test Passed
		fmt.Println(" _______ A checkpoint signed by a masternode was found _____")

	case <-miningTimeout:
		// Test Failed
		t.Fatalf(" _______ signed Checkpoint NOT found: mining interval expired _____")
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

	fmt.Println(" >>>>>>>> MN Checkpoint Signing Test Complete <<<<<<<<<<")
}

func txPoolContents(addrMap map[common.Address]types.Transactions, descr string) {
	for addr, pTx := range addrMap {
		for _, tx := range pTx {
			fmt.Printf("\t %s: Addr: %v  Tx Desc: %v Tx Hash: %v \n",
				descr, addr.Hash().String(), checkTxDesc(tx), tx.Hash().String())
		}
	}
}
