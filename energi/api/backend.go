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

// Package ethapi implements the general Ethereum API functions.
package api

import (
	"context"
	"math/big"

	"energi.world/core/gen3/accounts"
	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/core"
	"energi.world/core/gen3/core/state"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/core/vm"
	"energi.world/core/gen3/eth/downloader"
	"energi.world/core/gen3/ethdb"
	"energi.world/core/gen3/event"
	"energi.world/core/gen3/params"
	"energi.world/core/gen3/rpc"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// General Ethereum API
	Downloader() *downloader.Downloader
	ProtocolVersion() int
	SuggestPrice(ctx context.Context) (*big.Int, error)
	ChainDb() ethdb.Database
	EventMux() *event.TypeMux
	AccountManager() *accounts.Manager
	RPCGasCap() *big.Int // global gas cap for eth_call over rpc: DoS protection

	// BlockChain API
	SetHead(number uint64)
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error)
	BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error)
	StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error)
	GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)
	GetTd(blockHash common.Hash) *big.Int
	GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header) (*vm.EVM, func() error, error)
	SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription
	BlockChain() *core.BlockChain

	// TxPool API
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription

	ChainConfig() *params.ChainConfig
	CurrentBlock() *types.Block

	AddLocalCheckpoint(num uint64, hash common.Hash) error
	ListCheckpoints() []core.CheckpointInfo
	CheckpointSignatures(cp core.Checkpoint) []core.CheckpointSignature

	IsPublicService() bool
	OnSyncedHeadUpdates(cb func())
}

func createSignerCallback(
	backend Backend,
	password *string,
) bind.SignerFn {
	return func(
		signer types.Signer,
		addr common.Address,
		tx *types.Transaction,
	) (*types.Transaction, error) {
		account := accounts.Account{Address: addr}
		wallet, err := backend.AccountManager().Find(account)
		if err != nil {
			return nil, err
		}

		chain_id := backend.ChainConfig().ChainID

		if password == nil {
			return wallet.SignTx(
				account, tx, chain_id)
		}

		return wallet.SignTxWithPassphrase(
			account, *password, tx, chain_id)
	}
}
