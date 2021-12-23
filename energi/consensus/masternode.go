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
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/core/state"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/log"

	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

func (e *Energi) processMasternodes(
	chain ChainReader,
	header *types.Header,
	statedb *state.StateDB,
) (err error) {
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
		e.unlimitedGas,
		common.Big0,
		enumerateData,
		false,
	)
	rev_id := statedb.Snapshot()
	evm := e.createEVM(msg, chain, header, statedb)
	gp := core.GasPool(e.unlimitedGas)
	output, gas_used, _, err := core.ApplyMessage(evm, msg, &gp)
	statedb.RevertToSnapshot(rev_id)
	if err != nil {
		log.Error("Failed in enumerateActive() call", "err", err)
		return err
	}

	if gas_used > e.callGas {
		log.Warn("MasternodeRegistry::enumerateActive() took excessive gas",
			"gas", gas_used, "limit", e.callGas)
	}

	masternodes := new([]common.Address)
	err = e.mnregAbi.Unpack(&masternodes, "enumerateActive", output)
	if err != nil {
		log.Error("Failed to unpack enumerateActive() call", "err", err)
		return err
	}

	log.Debug("Masternode list", "masternodes", masternodes)

	mnGauge.Update(int64(len(*masternodes)))

	//clear out the account storage
	statedb.ForEachStorage(energi_params.Energi_MasternodeList, func(key, value common.Hash) bool {
		statedb.SetState(energi_params.Energi_MasternodeList, key, common.BytesToHash([]byte{0x00}))
		return true
	})

	//set active masternodes
	for _, addr := range *masternodes {
		statedb.SetState(energi_params.Energi_MasternodeList, addr.Hash(), common.BytesToHash([]byte{0x01}))
	}

	//commit tree changes
	statedb.GetOrNewStateObject(energi_params.Energi_MasternodeList).CommitTrie(statedb.Database())

	return nil
}
