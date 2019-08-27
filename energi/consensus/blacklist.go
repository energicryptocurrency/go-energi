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

func (e *Energi) processBlacklists(
	chain ChainReader,
	header *types.Header,
	statedb *state.StateDB,
) (err error) {
	blregistry := energi_params.Energi_BlacklistRegistry

	enumerateData, err := e.blacklistAbi.Pack("enumerateBlocked")
	if err != nil {
		log.Error("Fail to prepare enumerateBlocked() call", "err", err)
		return err
	}

	msg := types.NewMessage(
		blregistry,
		&blregistry,
		0,
		common.Big0,
		e.unlimitedGas,
		common.Big0,
		enumerateData,
		false,
	)
	evm := e.createEVM(msg, chain, header, statedb)
	gp := core.GasPool(e.unlimitedGas)
	output, gas_used, _, err := core.ApplyMessage(evm, msg, &gp)
	if err != nil {
		log.Error("Failed in enumerateBlocked() call", "err", err)
		return err
	}

	if gas_used > e.callGas {
		log.Warn("BlacklistRegistry::enumerateDrainable() took excessive gas",
			"gas", gas_used, "limit", e.callGas)
	}

	address_list := new([]common.Address)
	err = e.blacklistAbi.Unpack(&address_list, "enumerateBlocked", output)
	if err != nil {
		log.Error("Failed to unpack enumerateBlocked() call", "err", err)
		return err
	}

	log.Debug("Address blacklist", "address_list", address_list)
	empty_addr := common.Address{}
	state_obj := statedb.GetOrNewStateObject(energi_params.Energi_Blacklist)
	db := statedb.Database()
	value := common.BytesToHash([]byte{0x01})
	keep := make(state.KeepStorage, len(*address_list))
	whitelist := e.createWhitelist(statedb)

	for _, addr := range *address_list {
		if addr != empty_addr && !whitelist[addr] {
			addr_hash := addr.Hash()

			if (state_obj.GetState(db, addr_hash) == common.Hash{}) {
				log.Debug("New blacklisted account", "addr", addr)
			}

			log.Trace("Blacklisting account", "addr", addr)
			state_obj.SetState(db, addr_hash, value)
			keep[addr_hash] = true
		}
	}

	state_obj.CleanupStorage(keep)

	//
	wl_state_obj := statedb.GetOrNewStateObject(energi_params.Energi_Whitelist)
	wl_keep := make(state.KeepStorage, len(whitelist))
	for addr := range whitelist {
		addr_hash := addr.Hash()

		if (wl_state_obj.GetState(db, addr_hash) == common.Hash{}) {
			log.Debug("New whitelisted account", "addr", addr)
		}

		log.Trace("Whitelisting account", "addr", addr)
		wl_state_obj.SetState(db, addr_hash, value)
		wl_keep[addr_hash] = true
	}

	wl_state_obj.CleanupStorage(wl_keep)

	return nil
}

var (
	consensusProxies = []common.Address{
		energi_params.Energi_Treasury,
		energi_params.Energi_MasternodeRegistry,
		energi_params.Energi_StakerReward,
		energi_params.Energi_BackboneReward,
		energi_params.Energi_SporkRegistry,
		energi_params.Energi_CheckpointRegistry,
		energi_params.Energi_BlacklistRegistry,
		energi_params.Energi_MasternodeToken,
	}

	consensusStandalone = []common.Address{
		energi_params.Energi_MigrationContract,
		energi_params.Energi_SystemFaucet,
	}
)

func (e *Energi) createWhitelist(
	statedb *state.StateDB,
) map[common.Address]bool {
	whitelist := map[common.Address]bool{}

	for _, addr := range consensusProxies {
		whitelist[addr] = true
		impl := statedb.GetState(addr, energi_params.Storage_ProxyImpl)
		whitelist[common.BytesToAddress(impl[:])] = true
	}

	for _, addr := range consensusStandalone {
		whitelist[addr] = true
	}

	return whitelist
}

func (e *Energi) processDrainable(
	chain ChainReader,
	header *types.Header,
	statedb *state.StateDB,
) (err error) {
	blregistry := energi_params.Energi_BlacklistRegistry
	var comp_fund common.Address

	txhash := common.Hash{}
	statedb.Prepare(txhash, common.Hash{}, 0)

	// 1. List drainable addresses address
	//---
	enumerateData, err := e.blacklistAbi.Pack("enumerateDrainable")
	if err != nil {
		log.Error("Fail to prepare enumerateDrainable() call", "err", err)
		return err
	}

	msg := types.NewMessage(
		blregistry,
		&blregistry,
		0,
		common.Big0,
		e.unlimitedGas,
		common.Big0,
		enumerateData,
		false,
	)
	evm := e.createEVM(msg, chain, header, statedb)
	gp := core.GasPool(e.unlimitedGas)
	output, gas_used, _, err := core.ApplyMessage(evm, msg, &gp)
	if err != nil {
		log.Error("Failed in enumerateDrainable() call", "err", err)
		return err
	}

	if gas_used > e.callGas {
		log.Warn("BlacklistRegistry::enumerateDrainable() took excessive gas",
			"gas", gas_used, "limit", e.callGas)
	}

	address_list := new([]common.Address)
	err = e.blacklistAbi.Unpack(&address_list, "enumerateDrainable", output)
	if err != nil {
		log.Error("Failed to unpack enumerateDrainable() call", "err", err)
		return err
	}

	log.Debug("Address drain list", "address_list", address_list)

	// 2. Get current compensation fund address
	if len(*address_list) > 0 {
		enumerateData, err := e.blacklistAbi.Pack("compensation_fund")
		if err != nil {
			log.Error("Fail to prepare compensation_fund() call", "err", err)
			return err
		}

		msg := types.NewMessage(
			blregistry,
			&blregistry,
			0,
			common.Big0,
			e.callGas,
			common.Big0,
			enumerateData,
			false,
		)
		evm := e.createEVM(msg, chain, header, statedb)
		gp = core.GasPool(e.callGas)
		output, _, _, err := core.ApplyMessage(evm, msg, &gp)
		if err != nil {
			log.Error("Failed in compensation_fund() call", "err", err)
			return err
		}

		err = e.blacklistAbi.Unpack(&comp_fund, "compensation_fund", output)
		if err != nil {
			log.Error("Failed to unpack compensation_fund() call", "err", err)
			return err
		}
	}

	// 3. Drain
	//---
	empty_addr := common.Address{}

	for _, addr := range *address_list {
		if addr == empty_addr {
			continue
		}

		//--
		bal := statedb.GetBalance(addr)

		// Skip, if nothing
		if bal.Cmp(common.Big0) == 0 {
			continue
		}

		// Skip whitelisted
		if core.CanTransfer(statedb, addr, bal) {
			continue
		}

		statedb.AddBalance(comp_fund, bal)
		statedb.SetBalance(addr, common.Big0)
		log.Trace("Draining account", "fund", comp_fund, "addr", addr, "bal", bal)

		//--
		collectData, err := e.blacklistAbi.Pack("onDrain", addr)
		if err != nil {
			log.Error("Fail to prepare onDrain() call", "err", err, "addr", addr)
			return err
		}

		msg := types.NewMessage(
			blregistry,
			&blregistry,
			0,
			common.Big0,
			e.xferGas,
			common.Big0,
			collectData,
			false,
		)
		evm = e.createEVM(msg, chain, header, statedb)
		gp = core.GasPool(e.xferGas)
		_, _, _, err = core.ApplyMessage(evm, msg, &gp)
		if err != nil {
			log.Trace("Failed in onDrain() call", "err", err, "addr", addr)
			return err
		}
	}

	return nil
}
