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
	"errors"
	"strings"
	"time"

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
	energiMigrateID      types.MethodID
	energiMNHeartbeatID  types.MethodID
	energiMNInvalidateID types.MethodID
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
	copy(energiMNInvalidateID[:], mnreg_abi.Methods["invalidate"].Id())
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
	} else if method == energiMNInvalidateID {
		return true
	}

	return false
}

func IsBlacklisted(db vm.StateDB, addr common.Address) bool {
	return db.GetState(energi_params.Energi_Blacklist, addr.Hash()) != common.Hash{}
}

//=============================================================================

var (
	zfCleanupTimeout        = time.Minute
	zfMinHeartbeatPeriod    = time.Duration(30) * time.Minute
	zfMinInvalidationPeriod = time.Duration(2) * time.Minute

	ErrZeroFeeDoS = errors.New("zero-fee DoS")
)

type zeroFeeProtector struct {
	mnHeartbeats    map[common.Address]time.Time
	mnInvalidations map[common.Address]time.Time
	nextCleanup     time.Time
	timeNow         func() time.Time
}

func newZeroFeeProtector() *zeroFeeProtector {
	return &zeroFeeProtector{
		mnHeartbeats:    make(map[common.Address]time.Time),
		mnInvalidations: make(map[common.Address]time.Time),
		nextCleanup:     time.Now().Add(zfCleanupTimeout),
		timeNow:         time.Now,
	}
}

func (z *zeroFeeProtector) cleanupTimeout(
	now time.Time,
	timeMap map[common.Address]time.Time,
	timeout time.Duration,
) {
	for k, v := range timeMap {
		if now.Sub(v) > timeout {
			delete(timeMap, k)
		}
	}
}

func (z *zeroFeeProtector) cleanupRoutine(now time.Time) {
	if z.nextCleanup.After(now) {
		return
	}

	z.nextCleanup = now.Add(zfCleanupTimeout)
	//---

	z.cleanupTimeout(now, z.mnHeartbeats, zfMinHeartbeatPeriod)
	z.cleanupTimeout(now, z.mnInvalidations, zfMinInvalidationPeriod)
}

func (z *zeroFeeProtector) checkMasternode(
	pool *TxPool,
	sender common.Address,
	now time.Time,
	timeMap map[common.Address]time.Time,
	timeout time.Duration,
) error {
	if v, ok := timeMap[sender]; ok && now.Sub(v) < timeout {
		log.Debug("ZeroFee DoS by time", "sender", sender, "interval", now.Sub(v))
		return ErrZeroFeeDoS
	}

	// NOTE: potential issue with nonce gap
	mn_indicator := pool.currentState.GetState(
		energi_params.Energi_MasternodeList, sender.Hash())
	if (mn_indicator == common.Hash{}) {
		log.Debug("ZeroFee DoS by inactive MN", "sender", sender)
		return ErrZeroFeeDoS
	}

	timeMap[sender] = now
	return nil
}

func (z *zeroFeeProtector) checkDoS(pool *TxPool, tx *types.Transaction) error {
	now := z.timeNow()

	defer z.cleanupRoutine(now)

	sender, err := types.Sender(pool.signer, tx)
	if err != nil {
		return err
	}

	// NOTE: assumed to be called only on zero fee
	if method := tx.MethodID(); method == energiMNHeartbeatID {
		return z.checkMasternode(pool, sender, now, z.mnHeartbeats, zfMinHeartbeatPeriod)
	} else if method == energiMNInvalidateID {
		return z.checkMasternode(pool, sender, now, z.mnInvalidations, zfMinInvalidationPeriod)
	} else if method == energiMigrateID {
		return nil
	}
	return nil
}
