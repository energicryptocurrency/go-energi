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
	energiBLProposeID    types.MethodID
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

	bl_abi, err := abi.JSON(strings.NewReader(energi_abi.IBlacklistRegistryABI))
	if err != nil {
		panic(err)
	}
	copy(energiBLProposeID[:], bl_abi.Methods["propose"].Id())
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

//=============================================================================

var (
	pbCleanupTimeout = time.Minute
	pbPeriod         = time.Hour

	ErrPreBlacklist = errors.New("preliminary blacklisted")
)

type preBlacklist struct {
	proposed    map[common.Address]time.Time
	nextCleanup time.Time
	timeNow     func() time.Time
}

func newPreBlacklist() *preBlacklist {
	return &preBlacklist{
		proposed:    make(map[common.Address]time.Time),
		nextCleanup: time.Now().Add(pbCleanupTimeout),
		timeNow:     time.Now,
	}
}

func (pb *preBlacklist) cleanupRoutine(now time.Time) {
	if pb.nextCleanup.After(now) {
		return
	}

	pb.nextCleanup = now.Add(pbCleanupTimeout)
	//---
	for k, v := range pb.proposed {
		if now.Sub(v) > pbPeriod {
			delete(pb.proposed, k)
		}
	}
}

func (pb *preBlacklist) processTx(pool *TxPool, tx *types.Transaction) error {
	now := pb.timeNow()

	pb.cleanupRoutine(now)

	// Check if know preliminary blacklist
	//---
	sender, err := types.Sender(pool.signer, tx)
	if err != nil {
		log.Debug("Pre-blacklist sender error", "err", err)
		return err
	}

	if pb.isActive(sender, now) {
		log.Debug("Pre-blacklisted sender", "sender", sender)
		return ErrPreBlacklist
	}

	// Check if the call is valid
	//---
	pb.processProposal(pool, sender, now, tx)

	return nil
}

func (pb *preBlacklist) isActive(sender common.Address, now time.Time) bool {
	if t, ok := pb.proposed[sender]; ok && now.Sub(t) <= pbPeriod {
		return true
	}

	return false
}

func (pb *preBlacklist) processProposal(
	pool *TxPool,
	sender common.Address,
	now time.Time,
	tx *types.Transaction,
) {
	// Check if a new blacklist proposal
	//---
	if to := tx.To(); to == nil || *to != energi_params.Energi_BlacklistRegistry {
		return
	}
	if method := tx.MethodID(); method != energiBLProposeID {
		return
	}
	// DBL-10 - only enable for EBI proposals which have zero cost by design
	if tx.GasPrice().Cmp(common.Big0) != 0 {
		return
	}

	//---
	var target common.Address
	callData := tx.Data()
	copy(target[:], callData[16:36])

	// Do not reset timeout, if already known!
	if _, ok := pb.proposed[target]; ok {
		return
	}

	statedb := pool.currentState.Copy()

	if IsWhitelisted(statedb, target) {
		log.Warn("Skipping preliminary blacklist for whitelisted target",
			"target", target.Hex(), "sender", sender.Hex())
		return
	}

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

	bc := pool.chain.(*BlockChain)
	if bc == nil {
		log.Debug("PreBlacklist on missing blockchain")
		return
	}
	vmc := bc.GetVMConfig()
	ctx := NewEVMContext(msg, bc.CurrentHeader(), bc, &sender)
	ctx.GasLimit = ZeroFeeGasLimit
	evm := vm.NewEVM(ctx, statedb, bc.Config(), *vmc)

	gp := new(GasPool).AddGas(tx.Gas())
	output, _, failed, err := ApplyMessage(evm, msg, gp)
	if failed || err != nil {
		log.Debug("PreBlacklist failure at execution",
			"sender", sender, "target", target.Hex(), "err", err, "output", output)
		return
	}

	if len(output) != len(common.Hash{}) {
		log.Debug("PreBlacklist at unpack",
			"sender", sender, "target", target.Hex(), "output", output)
		return
	}

	// New pre-blacklist item
	//---
	log.Warn("New preliminary blacklist", "target", target.Hex(), "sender", sender.Hex())
	pb.proposed[target] = now
	pool.removeBySenderLocked(target)
}

func (pb *preBlacklist) filterBlocks(blocks types.Blocks) types.Blocks {
	pb.cleanupRoutine(pb.timeNow())

	for i, b := range blocks {
		if _, ok := pb.proposed[b.Coinbase()]; ok {
			return blocks[:i]
		}
	}

	return blocks
}

//=============================================================================

func (pool *TxPool) PreBlacklistHook(blocks types.Blocks) types.Blocks {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	return pool.preBlacklist.filterBlocks(blocks)
}

func (pool *TxPool) RemoveBySender(sender common.Address) bool {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	return pool.removeBySenderLocked(sender)
}

func (pool *TxPool) removeBySenderLocked(sender common.Address) bool {
	res := false

	if txs, ok := pool.pending[sender]; ok {
		for _, tx := range txs.Flatten() {
			txhash := tx.Hash()
			log.Trace("Removing by sender", "txhash", txhash, "sender", sender)
			pool.removeTx(txhash, true)
		}

		res = true
	}

	if txs, ok := pool.queue[sender]; ok {
		for _, tx := range txs.Flatten() {
			txhash := tx.Hash()
			log.Trace("Removing by sender", "txhash", txhash, "sender", sender)
			pool.removeTx(txhash, true)
		}

		res = true
	}

	return res
}
