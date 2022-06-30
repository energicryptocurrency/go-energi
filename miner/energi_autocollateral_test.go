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

package miner

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	ethereum "github.com/energicryptocurrency/energi"
	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/consensus"
	"github.com/energicryptocurrency/energi/core/types"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi "github.com/energicryptocurrency/energi/energi/consensus"
	"github.com/energicryptocurrency/energi/ethdb"
	"github.com/energicryptocurrency/energi/params"
)

type MockContractBackend struct {
	onCodeAt func(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)
	onCallContract func(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	onPendingCodeAt func(ctx context.Context, account common.Address) ([]byte, error)
	onPendingNonceAt func(ctx context.Context, account common.Address) (uint64, error)
	onSuggestGasPrice func(ctx context.Context) (*big.Int, error)
	onEstimateGas func(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
	onSendTransaction func(ctx context.Context, tx *types.Transaction) error
	onFilterLogs func(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	onSubscribeFilterLogs func(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
}

func (backend *MockContractBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return backend.onCodeAt(ctx, contract, blockNumber)
}

func (backend *MockContractBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return backend.onCallContract(ctx, call, blockNumber)
}

func (backend *MockContractBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return backend.onPendingCodeAt(ctx, account)
}

func (backend *MockContractBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return backend.onPendingNonceAt(ctx, account)
}

func (backend *MockContractBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return backend.onSuggestGasPrice(ctx)
}

func (backend *MockContractBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return backend.onEstimateGas(ctx, call)
}

func (backend *MockContractBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return backend.onSendTransaction(ctx, tx)
}

func (backend *MockContractBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return backend.onFilterLogs(ctx, query)
}

func (backend *MockContractBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return backend.onSubscribeFilterLogs(ctx, query, ch)
}

func newMockIMasterNodeTokenSession(backend bind.ContractBackend) (*energi_abi.IMasternodeTokenSession, error) {
	contract, err := energi_abi.NewIMasternodeToken(common.HexToAddress("0x0"), backend)
	if err != nil {
		return nil, err
	}

	session := &energi_abi.IMasternodeTokenSession{
		Contract: contract,
	}
	return session, nil
}

func TestCanAutocollateralizeEnergi(t *testing.T) {
	testCanAutocollateralize(t, energiChainConfig, energi.New(energiChainConfig.Energi, ethdb.NewMemDatabase()))
}

func testCanAutocollateralize(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, 0)
	w.current.header.Difficulty = new(big.Int).SetInt64(0)
	defer func() {
		w.close()
		_ = b.CleanUp()
	}()

	test_cases := []struct{
		min_limit, max_limit, token_balance, amount, amount_to_deposit *big.Int
	}{
		{ new(big.Int).SetUint64(10), new(big.Int).SetUint64(100), new(big.Int).SetUint64(70), new(big.Int).SetUint64(100), new(big.Int).SetUint64(30) },
		{ new(big.Int).SetUint64(10), new(big.Int).SetUint64(1000), new(big.Int).SetUint64(100), new(big.Int).SetUint64(200), new(big.Int).SetUint64(200) },
		{ new(big.Int).SetUint64(5), new(big.Int).SetUint64(20), new(big.Int).SetUint64(15), new(big.Int).SetUint64(10), new(big.Int).SetUint64(5) },
		{ new(big.Int).SetUint64(5), new(big.Int).SetUint64(20), new(big.Int).SetUint64(2), new(big.Int).SetUint64(10), new(big.Int).SetUint64(10) },
	}

	for _, tc := range test_cases {
		// Create a mock IMasternodeRegistery contract backend instance for testing
		w.apiBackend = &MockContractBackend{
			onCallContract: func(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
				parsedAbi, err := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryV2ABI))
				if err != nil {
					t.Error(err)
				}
	
				method, err := parsedAbi.MethodById(call.Data[:4])
				if err != nil {
					return nil, err
				}
	
				switch method.Name {
				// When the worker calls for collateral limits for master nodes
				case "collateralLimits":
					return method.Outputs.PackValues([]interface{}{tc.min_limit, tc.max_limit})
				default:
					return nil, fmt.Errorf("unsupported contract method %s", method.Name)
				}
			},
		}

		// Create a mock IMasternodetokenSession instance for testing
		mock_api, err := newMockIMasterNodeTokenSession(&MockContractBackend{
			onCallContract: func(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
				parsedAbi, err := abi.JSON(strings.NewReader(energi_abi.IMasternodeTokenABI))
				if err != nil {
					t.Error(err)
				}
	
				method, err := parsedAbi.MethodById(call.Data[:4])
				if err != nil {
					return nil, err
				}
	
				switch method.Name {
				// When the worker calls for balance of 0x0 address
				// Return the amount to be passed to the auto collateralize method
				case "balanceOf":
					return method.Outputs.PackValues([]interface{}{tc.token_balance})
				default:
					return nil, fmt.Errorf("unsupported contract method %s", method.Name)
				}
			},
		})
		if err != nil {
			t.Error(err)
		}

		// Check canautocollateralize
		res, err := w.canAutocollateralize(common.HexToAddress("0x0"), tc.amount, mock_api)
		if err != nil {
			t.Error(err)
		}

		// Confirm the result
		if res.Cmp(tc.amount_to_deposit) != 0 {
			t.Errorf("expected %v got %v", tc.amount_to_deposit.Uint64(), res.Uint64())
		}
	}
}
