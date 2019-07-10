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

package core

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	ZeroFeeGasLimit uint64 = 500000
)

var (
	energiMigrateID     types.MethodID
	energiMNHeartbeatID types.MethodID
	energiMNValidateID  types.MethodID
)

func init() {
	migration_abi, err := abi.JSON(strings.NewReader(energi_abi.Gen2MigrationABI))
	if err != nil {
		panic(err)
	}

	copy(energiMigrateID[:], migration_abi.Methods["claim"].Id())

	mnreg_abi, err := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryABI))
	if err != nil {
		panic(err)
	}
	copy(energiMNHeartbeatID[:], mnreg_abi.Methods["heartbeat"].Id())
	copy(energiMNValidateID[:], mnreg_abi.Methods["validate"].Id())
}

/**
 * SC-7: Zero-fee transactions
 *
 * Check if Energi consensus allows the transaction to be processed as
 * zero-fee.
 */
func IsValidZeroFee(tx *types.Transaction) bool {
	// Skip check for non-zero price
	if tx.Cost().Cmp(common.Big0) != 0 {
		return false
	}

	if tx.Gas() > ZeroFeeGasLimit {
		log.Trace("Zero-fee gas is over limit", "hash", tx.Hash(), "limit", tx.Gas())
		return false
	}

	if IsGen2Migration(tx) {
		return true
	}

	if IsMasternodeCall(tx) {
		return true
	}

	return false
}

func IsGen2Migration(tx *types.Transaction) bool {
	to := tx.To()

	return (to != nil) &&
		(*to == energi_params.Energi_MigrationContract) &&
		(tx.MethodID() == energiMigrateID)
}

func IsMasternodeCall(tx *types.Transaction) bool {
	to := tx.To()

	if (to == nil) || (*to != energi_params.Energi_MasternodeRegistry) {
		return false
	}

	if method := tx.MethodID(); method == energiMNHeartbeatID {
		return true
	} else if method == energiMNValidateID {
		return true
	}

	return false
}

func IsBlacklisted(db vm.StateDB, addr common.Address) bool {
	return db.GetState(energi_params.Energi_Blacklist, addr.Hash()) != common.Hash{}
}
