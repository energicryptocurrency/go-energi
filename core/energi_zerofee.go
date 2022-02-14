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
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/core/vm"
	"github.com/energicryptocurrency/energi/log"

	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

const (
	ZeroFeeGasLimit uint64 = 500000
)

var (
	energiClaimID        types.MethodID
	energiVerifyClaimID  types.MethodID
	energiMNHeartbeatID  types.MethodID
	energiMNInvalidateID types.MethodID
	energiCPSignID       types.MethodID
)

func init() {
	migration_abi, err := abi.JSON(strings.NewReader(energi_abi.Gen2MigrationABI))
	if err != nil {
		panic(err)
	}

	copy(energiClaimID[:], migration_abi.Methods["claim"].Id())
	copy(energiVerifyClaimID[:], migration_abi.Methods["verifyClaim"].Id())

	mnreg_abi, err := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryV2ABI))
	if err != nil {
		panic(err)
	}
	copy(energiMNHeartbeatID[:], mnreg_abi.Methods["heartbeat"].Id())
	copy(energiMNInvalidateID[:], mnreg_abi.Methods["invalidate"].Id())

	cpreg_abi, err := abi.JSON(strings.NewReader(energi_abi.ICheckpointRegistryABI))
	if err != nil {
		panic(err)
	}
	copy(energiCPSignID[:], cpreg_abi.Methods["sign"].Id())
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

	if IsCheckpointCall(tx) {
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

func IsCheckpointCall(tx *types.Transaction) bool {
	to := tx.To()

	if (to == nil) || (*to != energi_params.Energi_CheckpointRegistry) {
		return false
	}

	return tx.MethodID() == energiCPSignID
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
	zfMinHeartbeatPeriod    = time.Duration(1) * time.Minute
	zfMinInvalidationPeriod = time.Duration(1) * time.Minute
	zfMinCoinClaimPeriod    = time.Duration(3) * time.Minute
	zfMinCheckpointPeriod   = time.Duration(10) * time.Minute

	ErrZeroFeeDoS = errors.New("zero-fee DoS")
)

type zeroFeeProtector struct {
	mnHeartbeats    map[common.Address]time.Time
	mnInvalidations map[common.Address]time.Time
	mnCheckpoints   map[common.Address]time.Time
	coinClaims      map[uint32]time.Time
	nextCleanup     time.Time
	timeNow         func() time.Time
}

func newZeroFeeProtector() *zeroFeeProtector {
	return &zeroFeeProtector{
		mnHeartbeats:    make(map[common.Address]time.Time),
		mnInvalidations: make(map[common.Address]time.Time),
		mnCheckpoints:   make(map[common.Address]time.Time),
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

func (z *zeroFeeProtector) cleanupBySender(
	sender common.Address,
	timeMap map[common.Address]time.Time,
) {
	delete(timeMap, sender)
}

func (z *zeroFeeProtector) cleanupAllBySender(sender common.Address) {
	z.cleanupBySender(sender, z.mnHeartbeats)
	z.cleanupBySender(sender, z.mnInvalidations)
	z.cleanupBySender(sender, z.mnCheckpoints)
}

func (z *zeroFeeProtector) cleanupAllByTimeout(now time.Time) {
	if z.nextCleanup.After(now) {
		return
	}

	z.nextCleanup = now.Add(zfCleanupTimeout)
	//---

	z.cleanupTimeout(now, z.mnHeartbeats, zfMinHeartbeatPeriod)
	z.cleanupTimeout(now, z.mnInvalidations, zfMinInvalidationPeriod)
	z.cleanupTimeout(now, z.mnCheckpoints, zfMinCheckpointPeriod)

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
	tx *types.Transaction,
	timeMap map[common.Address]time.Time,
	timeout time.Duration,
) error {
	if v, ok := timeMap[sender]; ok && now.Sub(v) < timeout {
		log.Debug("ZeroFee DoS by time", "sender", sender, "interval", now.Sub(v))
		return fmt.Errorf("err: %v desc: time interval is less than the required", ErrZeroFeeDoS)
	}

	// NOTE: potential issue with nonce gap
	mn_indicator := pool.currentState.GetState(
		energi_params.Energi_MasternodeList, sender.Hash())
	if (mn_indicator == common.Hash{}) {
		log.Debug("ZeroFee DoS by inactive MN", "sender", sender)
		return fmt.Errorf("err: %v desc: inactive MN found", ErrZeroFeeDoS)
	}

	// Check if call is valid
	//---
	msg := types.NewMessage(
		sender,
		tx.To(),
		tx.Nonce(),
		tx.Value(),
		tx.Gas(),
		tx.GasPrice(),
		tx.Data(),
		false,
	)

	statedb := pool.currentState.Copy()

	bc, ok := pool.chain.(*BlockChain)
	if bc == nil || !ok {
		log.Debug("ZeroFee DoS on missing blockchain")
		return fmt.Errorf("err: %v desc: missing blockchain", ErrZeroFeeDoS)
	}
	vmc := bc.GetVMConfig()
	hdr := types.CopyHeader(bc.CurrentHeader())
	hdr.ParentHash = hdr.Hash()
	hdr.Number = new(big.Int).Add(hdr.Number, common.Big1)
	ctx := NewEVMContext(msg, hdr, bc, &sender)
	ctx.GasLimit = ZeroFeeGasLimit
	evm := vm.NewEVM(ctx, statedb, bc.Config(), *vmc)

	gp := new(GasPool).AddGas(tx.Gas())
	output, _, failed, err := ApplyMessage(evm, msg, gp)
	if failed || err != nil {
		strOutput := ""
		if len(output) > 4 {
			strOutput = string(output[4:])
		}
		log.Debug("ZeroFee DoS MN by execution",
			"sender", sender, "err", err, "output", strOutput)
		return fmt.Errorf("err: %v desc: %v", ErrZeroFeeDoS, err)
	}

	//---
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
	if len(callData) <= 36 {
		log.Debug("Missing tx data")
		return fmt.Errorf("invalid tx: missing tx data length")
	}

	item_id := uint32(new(big.Int).SetBytes(callData[4:36]).Uint64())

	if v, ok := z.coinClaims[item_id]; ok && now.Sub(v) < zfMinCoinClaimPeriod {
		log.Debug("ZeroFee DoS by time", "item_id", item_id, "interval", now.Sub(v))
		return fmt.Errorf("migrationErr: %v desc: time interval is less than the required", ErrZeroFeeDoS)
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

	bc, ok := pool.chain.(*BlockChain)
	if bc == nil || !ok {
		log.Debug("ZeroFee DoS on missing blockchain")
		return fmt.Errorf("migrationErr: %v desc: missing blockchain", ErrZeroFeeDoS)
	}
	vmc := bc.GetVMConfig()
	ctx := NewEVMContext(msg, bc.CurrentHeader(), bc, &sender)
	ctx.GasLimit = ZeroFeeGasLimit
	evm := vm.NewEVM(ctx, statedb, bc.Config(), *vmc)

	gp := new(GasPool).AddGas(tx.Gas())
	output, _, failed, err := ApplyMessage(evm, msg, gp)
	strOutput := ""
	if len(output) > 4 {
		strOutput = string(output[4:])
	}
	if failed || err != nil {
		log.Debug("ZeroFee DoS by execution",
			"item", item_id, "err", err, "output", strOutput)
		return fmt.Errorf("migrationErr: %v desc: %v", ErrZeroFeeDoS, err)
	}

	if len(output) != len(common.Hash{}) {
		log.Debug("ZeroFee DoS by unpack", "item", item_id, "output", strOutput)
		return fmt.Errorf("migrationErr: %v desc: %v", ErrZeroFeeDoS, err)
	}

	amount := new(big.Int).SetBytes(output)

	if amount.Cmp(common.Big0) <= 0 {
		log.Debug("ZeroFee DoS by already claimed", "item", item_id)
		return fmt.Errorf("migrationErr: %v desc: already claimed", ErrZeroFeeDoS)
	}

	//---
	z.coinClaims[item_id] = now
	log.Debug("ZeroFee migration", "item_id", item_id, "now", now)
	return nil
}

func (z *zeroFeeProtector) checkDoS(pool *TxPool, tx *types.Transaction) error {
	now := z.timeNow()

	defer z.cleanupAllByTimeout(now)

	sender, err := types.Sender(pool.signer, tx)
	if err != nil {
		log.Debug("ZeroFee DoS sender error", "err", err)
		return err
	}

	// NOTE: assumed to be called only on zero fee
	if method := tx.MethodID(); method == energiMNHeartbeatID {
		return z.checkMasternode(pool, sender, now, tx, z.mnHeartbeats, zfMinHeartbeatPeriod)
	} else if method == energiMNInvalidateID {
		return z.checkMasternode(pool, sender, now, tx, z.mnInvalidations, zfMinInvalidationPeriod)
	} else if method == energiClaimID {
		return z.checkMigration(pool, sender, now, tx)
	} else if method == energiCPSignID {
		return z.checkMasternode(pool, sender, now, tx, z.mnCheckpoints, zfMinCheckpointPeriod)
	}
	return nil
}
