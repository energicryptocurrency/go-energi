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
	"math/big"
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
	energiClaimID        types.MethodID
	energiVerifyClaimID  types.MethodID
	energiMNHeartbeatID  types.MethodID
	energiMNInvalidateID types.MethodID
)

func init() {
	migration_abi, err := abi.JSON(strings.NewReader(energi_abi.Gen2MigrationABI))
	if err != nil {
		panic(err)
	}

	copy(energiClaimID[:], migration_abi.Methods["claim"].Id())
	copy(energiVerifyClaimID[:], migration_abi.Methods["verifyClaim"].Id())

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
		(tx.MethodID() == energiClaimID)
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

func IsWhitelisted(db vm.StateDB, addr common.Address) bool {
	return db.GetState(energi_params.Energi_Whitelist, addr.Hash()) != common.Hash{}
}

//=============================================================================

var (
	zfCleanupTimeout        = time.Minute
	zfMinHeartbeatPeriod    = time.Duration(30) * time.Minute
	zfMinInvalidationPeriod = time.Duration(2) * time.Minute
	zfMinCoinClaimPeriod    = time.Duration(3) * time.Minute

	ErrZeroFeeDoS = errors.New("zero-fee DoS")
)

type zeroFeeProtector struct {
	mnHeartbeats    map[common.Address]time.Time
	mnInvalidations map[common.Address]time.Time
	coinClaims      map[uint32]time.Time
	nextCleanup     time.Time
	timeNow         func() time.Time
}

func newZeroFeeProtector() *zeroFeeProtector {
	return &zeroFeeProtector{
		mnHeartbeats:    make(map[common.Address]time.Time),
		mnInvalidations: make(map[common.Address]time.Time),
		coinClaims:      make(map[uint32]time.Time),
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

	for k, v := range z.coinClaims {
		if now.Sub(v) > zfMinCoinClaimPeriod {
			delete(z.coinClaims, k)
		}
	}
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
	log.Debug("ZeroFee masternode", "sender", sender, "now", now)
	return nil
}

func (z *zeroFeeProtector) checkMigration(
	pool *TxPool,
	sender common.Address,
	now time.Time,
	tx *types.Transaction,
) error {
	callData := tx.Data()
	item_id := uint32(new(big.Int).SetBytes(callData[4:36]).Uint64())

	if v, ok := z.coinClaims[item_id]; ok && now.Sub(v) < zfMinCoinClaimPeriod {
		log.Debug("ZeroFee DoS by time", "item_id", item_id, "interval", now.Sub(v))
		return ErrZeroFeeDoS
	}

	// Check if call is valid
	//---
	copy(callData[:], energiVerifyClaimID[:])

	msg := types.NewMessage(
		sender,
		tx.To(),
		tx.Nonce(),
		tx.Value(),
		tx.Gas(),
		tx.GasPrice(),
		callData,
		false,
	)

	// Just in case: safety measure
	statedb := pool.currentState.Copy()

	bc := pool.chain.(*BlockChain)
	if bc == nil {
		log.Debug("ZeroFee DoS on missing blockchain")
		return ErrZeroFeeDoS
	}
	vmc := bc.GetVMConfig()
	ctx := NewEVMContext(msg, bc.CurrentHeader(), bc, &sender)
	ctx.GasLimit = ZeroFeeGasLimit
	evm := vm.NewEVM(ctx, statedb, bc.Config(), *vmc)

	gp := new(GasPool).AddGas(tx.Gas())
	output, _, failed, err := ApplyMessage(evm, msg, gp)
	if failed || err != nil {
		log.Debug("ZeroFee DoS by execution",
			"item", item_id, "err", err, "output", output)
		return ErrZeroFeeDoS
	}

	if len(output) != len(common.Hash{}) {
		log.Debug("ZeroFee DoS by unpack", "item", item_id, "output", output)
		return ErrZeroFeeDoS
	}

	amount := new(big.Int).SetBytes(output)

	if amount.Cmp(common.Big0) <= 0 {
		log.Debug("ZeroFee DoS by already claimed", "item", item_id)
		return ErrZeroFeeDoS
	}

	//---
	z.coinClaims[item_id] = now
	log.Debug("ZeroFee migration", "item_id", item_id, "now", now)
	return nil
}

func (z *zeroFeeProtector) checkDoS(pool *TxPool, tx *types.Transaction) error {
	now := z.timeNow()

	defer z.cleanupRoutine(now)

	sender, err := types.Sender(pool.signer, tx)
	if err != nil {
		log.Debug("ZeroFee DoS sender error", "err", err)
		return err
	}

	// NOTE: assumed to be called only on zero fee
	if method := tx.MethodID(); method == energiMNHeartbeatID {
		return z.checkMasternode(pool, sender, now, z.mnHeartbeats, zfMinHeartbeatPeriod)
	} else if method == energiMNInvalidateID {
		return z.checkMasternode(pool, sender, now, z.mnInvalidations, zfMinInvalidationPeriod)
	} else if method == energiClaimID {
		return z.checkMigration(pool, sender, now, tx)
	}
	return nil
}

