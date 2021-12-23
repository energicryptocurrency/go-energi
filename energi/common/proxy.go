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

package common

import (
	"context"

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/core/state"

	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

// GeneralProxyHashFunc provides the function that helps check for governed
// proxy contracts filtered logs.
type GeneralProxyHashFunc func(addr common.Address, blockheight *uint64) *common.Hash

// GeneralProxyHashGen returns a function that helps retrieve the governed proxy contract hash.
func GeneralProxyHashGen(blockchain *core.BlockChain) GeneralProxyHashFunc {
	return func(addr common.Address, blockheight *uint64) *common.Hash {
		var err error
		var statedb *state.StateDB

		if blockheight == nil || *blockheight < 1 {
			statedb, err = blockchain.State()
		} else {
			header := blockchain.GetHeaderByNumber(*blockheight)
			statedb, err = blockchain.StateAt(header.Hash())
		}
		if err != nil {
			return nil
		}

		// Get a hash that allows logs from governed proxy checkpoint contract to be filtered.
		prxyHash := statedb.GetState(addr, energi_params.Storage_ProxyImpl)
		return &prxyHash
	}
}

// GeneralProxyHashExtractor retrieves if it exists the proxy hash func passed
// through the context.
func GeneralProxyHashExtractor(ctx context.Context, qAddr common.Address, blockNo *uint64) *common.Hash {
	proxyHashFunc := ctx.Value(energi_params.GeneralProxyCtxKey).(GeneralProxyHashFunc)
	if proxyHashFunc == nil {
		return nil
	}

	return proxyHashFunc(qAddr, blockNo)
}
