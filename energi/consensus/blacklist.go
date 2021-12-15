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

var (
	blacklistValue = common.BytesToHash([]byte{0x01})
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
	rev_id := statedb.Snapshot()
	evm := e.createEVM(msg, chain, header, statedb)
	gp := core.GasPool(e.unlimitedGas)
	output, gas_used, _, err := core.ApplyMessage(evm, msg, &gp)
	statedb.RevertToSnapshot(rev_id)
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
	// if account is in whitelist don't set it as blocked in Energi_Blacklist account
	whitelist := e.createWhitelist(statedb)
	empty_addr := common.Address{}

	// clear out the account storage
	statedb.ForEachStorage(energi_params.Energi_Blacklist, func(key, value common.Hash) bool {
		statedb.SetState(energi_params.Energi_Blacklist, key, common.BytesToHash([]byte{0x00}))
		return true
	})

	// set only blocked accounts
	for _, addr := range *address_list {
		if addr != empty_addr && !whitelist[addr] {
			log.Trace("Blacklisting account", "addr", addr)
			statedb.SetState(energi_params.Energi_Blacklist, addr.Hash(), common.BytesToHash([]byte{0x01}))
		}
	}
	// commit tree changes
	statedb.GetOrNewStateObject(energi_params.Energi_Blacklist).CommitTrie(statedb.Database())

	// clear out the account storage
	statedb.ForEachStorage(energi_params.Energi_Whitelist, func(key, value common.Hash) bool {
		statedb.SetState(energi_params.Energi_Whitelist, key, common.BytesToHash([]byte{0x00}))
		return true
	})

	// set whitelisted accounts in the Energi_Whitelist storage
	for addr := range whitelist {
		// log.Trace("Whitelisting account", "addr", addr)
		statedb.SetState(energi_params.Energi_Whitelist, addr.Hash(), common.BytesToHash([]byte{0x01}))
	}
	// commit tree changes
	statedb.GetOrNewStateObject(energi_params.Energi_Whitelist).CommitTrie(statedb.Database())

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
	txs types.Transactions,
	receipts types.Receipts,
) (types.Transactions, types.Receipts, error) {
	blregistry := energi_params.Energi_BlacklistRegistry
	var comp_fund common.Address

	txhash := common.Hash{}
	bhash := header.Hash()
	statedb.Prepare(txhash, bhash, len(txs))

	// 1. List drainable addresses address
	// ---
	enumerateData, err := e.blacklistAbi.Pack("enumerateDrainable")
	if err != nil {
		log.Error("Fail to prepare enumerateDrainable() call", "err", err)
		return nil, nil, err
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
	rev_id := statedb.Snapshot()
	evm := e.createEVM(msg, chain, header, statedb)
	gp := core.GasPool(e.unlimitedGas)
	output, gas_used, _, err := core.ApplyMessage(evm, msg, &gp)
	statedb.RevertToSnapshot(rev_id)
	if err != nil {
		log.Error("Failed in enumerateDrainable() call", "err", err)
		return nil, nil, err
	}

	if gas_used > e.callGas {
		log.Warn("BlacklistRegistry::enumerateDrainable() took excessive gas",
			"gas", gas_used, "limit", e.callGas)
	}

	address_list := new([]common.Address)
	err = e.blacklistAbi.Unpack(&address_list, "enumerateDrainable", output)
	if err != nil {
		log.Error("Failed to unpack enumerateDrainable() call", "err", err)
		return nil, nil, err
	}

	log.Debug("Address drain list", "address_list", address_list)

	// 2. Get current compensation fund address
	if len(*address_list) > 0 {
		enumerateData, err := e.blacklistAbi.Pack("compensation_fund")
		if err != nil {
			log.Error("Fail to prepare compensation_fund() call", "err", err)
			return nil, nil, err
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
		rev_id := statedb.Snapshot()
		evm := e.createEVM(msg, chain, header, statedb)
		gp = core.GasPool(e.callGas)
		output, _, _, err := core.ApplyMessage(evm, msg, &gp)
		statedb.RevertToSnapshot(rev_id)
		if err != nil {
			log.Error("Failed in compensation_fund() call", "err", err)
			return nil, nil, err
		}

		err = e.blacklistAbi.Unpack(&comp_fund, "compensation_fund", output)
		if err != nil {
			log.Error("Failed to unpack compensation_fund() call", "err", err)
			return nil, nil, err
		}
	}

	// 3. Drain
	// ---
	empty_addr := common.Address{}

	for _, addr := range *address_list {
		if addr == empty_addr {
			continue
		}

		// --
		bal := statedb.GetBalance(addr)

		// Skip, if nothing
		if bal.Cmp(common.Big0) == 0 {
			continue
		}

		// Skip whitelisted
		if core.CanTransfer(statedb, addr, bal) {
			continue
		}

		log.Trace("Draining account", "fund", comp_fund, "addr", addr, "bal", bal)

		// ====================================
		contributeData, err := e.treasuryAbi.Pack("contribute")
		if err != nil {
			log.Error("Fail to prepare contribute() call", "err", err)
			return nil, nil, err
		}

		tx := types.NewTransaction(
			statedb.GetNonce(addr),
			comp_fund,
			bal,
			e.xferGas,
			common.Big0,
			contributeData)
		tx = tx.WithConsensusSender(addr)

		statedb.Prepare(tx.Hash(), bhash, len(txs))

		msg, err = tx.AsMessage(&ConsensusSigner{})
		if err != nil {
			log.Error("Failed in onDrain() msg", "err", err)
			return nil, nil, err
		}

		statedb.SetState(energi_params.Energi_Blacklist, addr.Hash(), common.Hash{})
		evm = e.createEVM(msg, chain, header, statedb)
		gp = core.GasPool(msg.Gas())
		_, gas1, failed, err := core.ApplyMessage(evm, msg, &gp)
		statedb.SetState(energi_params.Energi_Blacklist, addr.Hash(), blacklistValue)

		if err != nil {
			log.Error("Failed in onDrain() call", "err", err)
			return nil, nil, err
		}

		// NOTE: it should be Byzantium finalization here...
		root := statedb.IntermediateRoot(chain.Config().IsEIP158(header.Number))
		receipt := types.NewReceipt(root.Bytes(), failed, header.GasUsed)
		receipt.TxHash = tx.Hash()
		receipt.GasUsed = gas1
		receipt.Logs = statedb.GetLogs(tx.Hash())
		receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

		txs = append(txs, tx)
		receipts = append(receipts, receipt)

		// ====================================
		collectData, err := e.blacklistAbi.Pack("onDrain", addr)
		if err != nil {
			log.Error("Fail to prepare onDrain() call", "err", err, "addr", addr)
			return nil, nil, err
		}

		tx = types.NewTransaction(
			statedb.GetNonce(blregistry),
			blregistry,
			common.Big0,
			e.xferGas,
			common.Big0,
			collectData)
		tx = tx.WithConsensusSender(blregistry)

		statedb.Prepare(tx.Hash(), bhash, len(txs))

		msg, err = tx.AsMessage(&ConsensusSigner{})
		if err != nil {
			log.Error("Failed in onDrain() msg", "err", err)
			return nil, nil, err
		}

		evm = e.createEVM(msg, chain, header, statedb)
		gp = core.GasPool(msg.Gas())
		_, gas2, failed, err := core.ApplyMessage(evm, msg, &gp)
		if err != nil {
			log.Error("Failed in onDrain() call", "err", err)
			return nil, nil, err
		}

		// NOTE: it should be Byzantium finalization here...
		root = statedb.IntermediateRoot(chain.Config().IsEIP158(header.Number))
		receipt = types.NewReceipt(root.Bytes(), failed, header.GasUsed)
		receipt.TxHash = tx.Hash()
		receipt.GasUsed = gas2
		receipt.Logs = statedb.GetLogs(tx.Hash())
		receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

		txs = append(txs, tx)
		receipts = append(receipts, receipt)
	}

	return txs, receipts, nil
}
