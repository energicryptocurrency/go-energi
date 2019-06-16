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

package consensus

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	BigZero    = big.NewInt(0)
	BigOne     = big.NewInt(1)
	BigBalance = new(big.Int).Div(math.MaxBig256, big.NewInt(2))
)

func (e *Energi) processBlockRewards(
	chain ChainReader,
	header *types.Header,
	statedb *state.StateDB,
) error {
	gp := new(core.GasPool)
	systemFaucet := e.systemFaucet

	// Temporary balance setup & clean up
	statedb.SetBalance(systemFaucet, BigBalance)
	defer statedb.SetBalance(systemFaucet, BigZero)

	// Common get reward call
	getRewardData, err := e.rewardAbi.Pack("getReward", header.Number)
	if err != nil {
		log.Error("Fail to prepare getReward() call", err)
		return err
	}

	rewardData, err := e.rewardAbi.Pack("reward")
	if err != nil {
		log.Error("Fail to prepare reward() call", err)
		return err
	}

	for _, caddr := range e.rewardGov {
		// GetReward()
		msg := types.NewMessage(
			systemFaucet,
			&caddr,
			0,
			BigZero,
			e.callGas,
			BigOne,
			getRewardData,
			false,
		)
		evm := e.createEVM(msg, chain, header, statedb)
		gp.AddGas(e.callGas)
		output, _, _, err := core.ApplyMessage(evm, msg, gp)
		if err != nil {
			log.Error("Failed in getReward() call", err)
			return err
		}

		//
		value := big.NewInt(0)
		err = e.rewardAbi.Unpack(&value, "getReward", output)
		if err != nil {
			log.Error("Failed to unpack getReward() call", err)
			return err
		}

		// Reward
		msg = types.NewMessage(
			systemFaucet,
			&caddr,
			0,
			value,
			e.xferGas,
			BigOne,
			rewardData,
			false,
		)
		evm = e.createEVM(msg, chain, header, statedb)
		gp.AddGas(e.xferGas)
		_, _, _, err = core.ApplyMessage(evm, msg, gp)
		if err != nil {
			log.Error("Failed in reward() call", err)
			return err
		}
	}

	return nil
}
