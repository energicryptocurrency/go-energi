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
	"context"
	"errors"
	"math/big"

	ethereum "energi.world/core/gen3"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/core"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/core/vm"
	"energi.world/core/gen3/event"
	"energi.world/core/gen3/rpc"

	energi_common "energi.world/core/gen3/energi/common"
	energi_params "energi.world/core/gen3/energi/params"
)

func (b *EthAPIBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	rpcBlockNumber := rpc.LatestBlockNumber

	if blockNumber != nil {
		rpcBlockNumber = rpc.BlockNumber(blockNumber.Int64())
	}

	state, _, err := b.StateAndHeaderByNumber(ctx, rpcBlockNumber)
	if err != nil {
		return nil, err
	}

	return state.GetCode(contract), nil
}

func (b *EthAPIBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	rpcBlockNumber := rpc.LatestBlockNumber

	if blockNumber != nil {
		rpcBlockNumber = rpc.BlockNumber(blockNumber.Int64())
	}

	state, header, err := b.StateAndHeaderByNumber(ctx, rpcBlockNumber)
	if err != nil {
		return nil, err
	}

	if call.Gas == 0 {
		call.Gas = 100000
	}

	msg := types.NewMessage(
		energi_params.Energi_SystemFaucet,
		call.To,
		0,
		common.Big0,
		call.Gas,
		common.Big0,
		call.Data,
		false,
	)

	evmctx := core.NewEVMContext(msg, header, b.eth.blockchain, &header.Coinbase)
	vmenv := vm.NewEVM(evmctx, state, b.eth.chainConfig, *b.eth.blockchain.GetVMConfig())
	gaspool := new(core.GasPool).AddGas(call.Gas)

	ret, _, _, err := core.NewStateTransition(vmenv, msg, gaspool).TransitionDb()
	return ret, err
}

func (b *EthAPIBackend) PendingCodeAt(
	ctx context.Context,
	account common.Address,
) ([]byte, error) {
	return b.CodeAt(ctx, account, new(big.Int).SetInt64(int64(rpc.PendingBlockNumber)))
}

func (b *EthAPIBackend) PendingNonceAt(
	ctx context.Context,
	account common.Address,
) (uint64, error) {
	return b.GetPoolNonce(ctx, account)
}

func (b *EthAPIBackend) PendingCallContract(
	ctx context.Context,
	call ethereum.CallMsg,
) ([]byte, error) {
	return b.CallContract(ctx, call, new(big.Int).SetInt64(int64(rpc.PendingBlockNumber)))
}

func (b *EthAPIBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *EthAPIBackend) EstimateGas(
	ctx context.Context,
	call ethereum.CallMsg,
) (gas uint64, err error) {
	return 0, errors.New("Not implemented")
}

func (b *EthAPIBackend) SendTransaction(
	ctx context.Context,
	tx *types.Transaction,
) error {
	return b.SendTx(ctx, tx)
}

// FilterLogs is a less efficient method of fetching the logs in a given block.
func (b *EthAPIBackend) FilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
) ([]types.Log, error) {
	toBlock := rpc.LatestBlockNumber
	if query.ToBlock != nil {
		toBlock = rpc.BlockNumber(query.ToBlock.Int64())
	}

	rpcBlockNumber := toBlock
	if query.FromBlock != nil {
		rpcBlockNumber = rpc.BlockNumber(query.FromBlock.Int64())
	}

	requiredLogs := make([]types.Log, 0, int(toBlock-rpcBlockNumber))
	for i := rpcBlockNumber; i <= toBlock; i++ {
		header, err := b.HeaderByNumber(ctx, i)
		if err != nil {
			return nil, err
		}

		// Fetch txs in the block with the provided block hash
		allLogs, err := b.GetLogs(ctx, header.Hash())
		if err != nil {
			return nil, err
		}

		blockNo := uint64(i)
		for _, logs := range allLogs {
			for _, log := range logs {
				if b.isFilteredLog(ctx, query, log, &blockNo) {
					requiredLogs = append(requiredLogs, *log)
				}
			}
		}

		// When fetching the latest block, do not loop more than once
		if i == rpc.LatestBlockNumber {
			break
		}
	}

	return requiredLogs, nil
}

// SubscribeFilterLogs returns the logs that are created after subscription.
func (b *EthAPIBackend) SubscribeFilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	// Subscribe to all contract events
	sinkLogs := make(chan []*types.Log)

	sub := b.SubscribeLogsEvent(sinkLogs)
	// Since we're getting logs in batches, we need to flatten them into a plain stream
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case logs := <-sinkLogs:
				for _, log := range logs {
					// Select the required logs only.
					if !b.isFilteredLog(ctx, query, log, nil) {
						continue
					}

					select {
					case ch <- *log:
					case err := <-sub.Err():
						if err != nil {
							return err
						}
					case <-quit:
						return nil
					}
				}
			case err := <-sub.Err():
				if err != nil {
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (b *EthAPIBackend) isFilteredLog(
	ctx context.Context,
	q ethereum.FilterQuery,
	log *types.Log,
	blockNo *uint64,
) bool {

	for _, addr := range q.Addresses {
		generalProxyHash := energi_common.GeneralProxyHashExtractor(ctx, addr, blockNo)
		if generalProxyHash != nil && log.Address.Hash() == *generalProxyHash {
			return true
		}

		if addr == log.Address {
			return true
		}
	}

	for _, queryTopics := range q.Topics {
		if len(queryTopics) > 0 {
			for _, foundTopic := range log.Topics {
				// Check if missed event name topic was returned.
				if len(foundTopic) > 0 && queryTopics[0] == foundTopic {
					return true
				}
			}
		}
	}

	return false
}
