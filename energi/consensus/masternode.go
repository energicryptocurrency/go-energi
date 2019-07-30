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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	energi_params "energi.world/core/gen3/energi/params"
)

func (e *Energi) processMasternodes(
	chain ChainReader,
	header *types.Header,
	statedb *state.StateDB,
) (err error) {
	gp := new(core.GasPool)
	mnregistry := energi_params.Energi_MasternodeRegistry

	enumerateData, err := e.mnregAbi.Pack("enumerateActive")
	if err != nil {
		log.Error("Fail to prepare enumerateActive() call", "err", err)
		return err
	}

	msg := types.NewMessage(
		mnregistry,
		&mnregistry,
		0,
		common.Big0,
		e.callGas,
		common.Big0,
		enumerateData,
		false,
	)
	evm := e.createEVM(msg, chain, header, statedb)
	gp.AddGas(e.callGas)
	output, _, _, err := core.ApplyMessage(evm, msg, gp)
	if err != nil {
		log.Error("Failed in enumerateActive() call", "err", err)
		return err
	}

	masternodes := new([]common.Address)
	err = e.mnregAbi.Unpack(&masternodes, "enumerateActive", output)
	if err != nil {
		log.Error("Failed to unpack enumerateActive() call", "err", err)
		return err
	}

	log.Debug("Masternode list", "masternodes", masternodes)

	state_obj := statedb.GetOrNewStateObject(energi_params.Energi_MasternodeList)
	db := statedb.Database()
	value := common.BytesToHash([]byte{0x01})
	keep := make(state.KeepStorage, len(*masternodes))

	for _, addr := range *masternodes {
		addr_hash := addr.Hash()

		if (state_obj.GetState(db, addr_hash) == common.Hash{}) {
			log.Debug("New masternode", "addr", addr)
		}

		state_obj.SetState(db, addr_hash, value)
		keep[addr_hash] = true
	}

	state_obj.CleanupStorage(keep)

	return nil
}
