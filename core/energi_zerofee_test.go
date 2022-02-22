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
	"testing"
	"time"

	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/consensus/ethash"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/core/vm"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/ethdb"
	"github.com/energicryptocurrency/energi/params"

	"github.com/stretchr/testify/assert"
)

func TestIsValidZeroFee(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	migration_abi, err := abi.JSON(strings.NewReader(energi_abi.Gen2MigrationABI))
	assert.Empty(t, err)
	claimCall, err := migration_abi.Pack(
		"claim", common.Big0, common.Address{}, uint8(0), common.Hash{}, common.Hash{})
	assert.Empty(t, err)

	mnreg_abi, err := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryV2ABI))
	assert.Empty(t, err)
	heartbeatCall, err := mnreg_abi.Pack("heartbeat", common.Big1, common.Hash{}, common.Big0)
	assert.Empty(t, err)
	invalidateCall, err := mnreg_abi.Pack("invalidate", common.Address{})
	assert.Empty(t, err)
	cpreg_abi, err := abi.JSON(strings.NewReader(energi_abi.ICheckpointRegistryABI))
	assert.Empty(t, err)
	cpsignCall, err := cpreg_abi.Pack("sign", common.Address{}, []byte{})
	assert.Empty(t, err)

	res := false

	//
	res = IsValidZeroFee(types.NewTransaction(
		0,
		common.Address{},
		common.Big0,
		500000,
		common.Big0,
		[]byte{},
	))
	assert.False(t, res, "Simple zero")
	//
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		500000,
		common.Big0,
		claimCall,
	))
	assert.True(t, res, "Valid migration claim")
	//
	res = IsValidZeroFee(types.NewTransaction(
		2,
		energi_params.Energi_MigrationContract,
		common.Big0,
		500000,
		common.Big0,
		claimCall,
	))
	assert.True(t, res, "Valid migration claim - other nonce")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big1,
		500000,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid migration claim - amount")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		500001,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid migration claim - gas")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		500000,
		common.Big1,
		claimCall,
	))
	assert.False(t, res, "Invalid migration claim - gas price")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MasternodeRegistry,
		common.Big0,
		500000,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid migration claim - dst")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		500000,
		common.Big0,
		heartbeatCall,
	))
	assert.False(t, res, "Invalid migration claim - call")
	//
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MasternodeRegistry,
		common.Big0,
		10000,
		common.Big0,
		heartbeatCall,
	))
	assert.True(t, res, "Valid MN heartbeat")
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MasternodeRegistry,
		common.Big0,
		10000,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid MN heartbeat - data")
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		40000,
		common.Big0,
		heartbeatCall,
	))
	assert.False(t, res, "Invalid MN heartbeat - dst")
	//
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MasternodeRegistry,
		common.Big0,
		10000,
		common.Big0,
		invalidateCall,
	))
	assert.True(t, res, "Valid MN validate")
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MasternodeRegistry,
		common.Big0,
		10000,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid MN validate - data")
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		40000,
		common.Big0,
		invalidateCall,
	))
	assert.False(t, res, "Invalid MN validate - dst")
	//
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_CheckpointRegistry,
		common.Big0,
		10000,
		common.Big0,
		cpsignCall,
	))
	assert.True(t, res, "Valid MN CP sign")
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_CheckpointRegistry,
		common.Big0,
		10000,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid MN CP sign - data")
}

//---

type fakeSigner struct {
	sender common.Address
}

func (s *fakeSigner) Sender(tx *types.Transaction) (common.Address, error) {
	return s.sender, nil
}
func (sg *fakeSigner) SignatureValues(tx *types.Transaction, sig []byte) (r, s, v *big.Int, err error) {
	return common.Big0, common.Big0, common.Big0, nil
}
func (s *fakeSigner) Hash(tx *types.Transaction) common.Hash {
	return tx.Hash()
}
func (s *fakeSigner) Equal(types.Signer) bool {
	return false
}

func TestZeroFeeProtectorMasternode(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	now := time.Now() // It can be fixed
	adjust_time := time.Duration(0)

	protector := newZeroFeeProtector()
	protector.timeNow = func() time.Time {
		return now.Add(adjust_time)
	}

	pool := &TxPool{}

	signer := &fakeSigner{}
	pool.signer = signer

	testdb := ethdb.NewMemDatabase()
	gspec := &Genesis{
		Config: params.TestnetChainConfig,
	}
	gspec.MustCommit(testdb)
	engine := ethash.NewFaker()

	chain, err := NewBlockChain(
		testdb, nil, gspec.Config,
		engine, vm.Config{}, nil)
	assert.Empty(t, err)
	defer chain.Stop()
	pool.chain = chain
	pool.currentState, _ = chain.State()

	mn_active1 := common.HexToAddress("0x0000000000000000000000000000000022345678")
	mn_active2 := common.HexToAddress("0x0000000000000000000000000000000022345679")
	mn_inactive := common.HexToAddress("0x0000000000000000000000000000000022345680")

	pool.currentState.SetState(
		energi_params.Energi_MasternodeList,
		mn_active1.Hash(),
		mn_inactive.Hash(),
	)
	pool.currentState.SetState(
		energi_params.Energi_MasternodeList,
		mn_active2.Hash(),
		mn_inactive.Hash(),
	)
	pool.currentState.SetCode(
		energi_params.Energi_MasternodeRegistry,
		// PUSH1 0 PUSH1 0 RETURN
		[]byte{0x60, 0x00, 0x60, 0x00, 0xF3},
	)

	mnreg_abi, err := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryV2ABI))
	assert.Empty(t, err)
	heartbeatCall, err := mnreg_abi.Pack("heartbeat", common.Big1, common.Hash{}, common.Big0)
	assert.Empty(t, err)
	invalidateCall, err := mnreg_abi.Pack("invalidate", common.Address{})
	assert.Empty(t, err)
	cpreg_abi, _ := abi.JSON(strings.NewReader(energi_abi.ICheckpointRegistryABI))
	cpsignCall, err := cpreg_abi.Pack("sign", common.Address{}, []byte{})
	assert.Empty(t, err)

	hbtx0 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, heartbeatCall)
	invtx0 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, invalidateCall)
	sigtx0 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, cpsignCall)
	hbtx1 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, heartbeatCall)
	invtx1 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, invalidateCall)
	sigtx1 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, cpsignCall)
	hbtx2 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, heartbeatCall)
	invtx2 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, invalidateCall)
	sigtx2 := types.NewTransaction(
		1, energi_params.Energi_MasternodeRegistry, common.Big0, 100000, common.Big0, cpsignCall)

	// Inactive MN
	signer.sender = mn_inactive
	errMsg := errors.New("err: zero-fee DoS desc: inactive MN found")
	err = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, invtx0)
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, sigtx0)
	assert.Equal(t, errMsg, err)

	// Active MN
	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, sigtx1)
	assert.Equal(t, nil, err)

	// Active MN repeat interval
	signer.sender = mn_active1
	errMsg = errors.New("err: zero-fee DoS desc: time interval is less than the required")
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, sigtx1)
	assert.Equal(t, errMsg, err)

	// Active another MV
	signer.sender = mn_active2
	err = protector.checkDoS(pool, hbtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, sigtx2)
	assert.Equal(t, nil, err)

	// Active MN after invalidation period is over
	adjust_time = time.Duration(5) * time.Minute

	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, sigtx1)
	assert.Equal(t, errMsg, err)

	signer.sender = mn_active2
	err = protector.checkDoS(pool, hbtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, sigtx2)
	assert.Equal(t, errMsg, err)

	// Active MN after checkpoint period is over
	adjust_time = time.Duration(10) * time.Minute

	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, sigtx1)
	assert.Equal(t, nil, err)

	signer.sender = mn_active2
	err = protector.checkDoS(pool, hbtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, sigtx2)
	assert.Equal(t, nil, err)

	// Active MN after heartbeat period is over
	adjust_time = time.Duration(20) * time.Minute

	pool.currentState.SetCode(
		energi_params.Energi_MasternodeRegistry,
		// PUSH1 0 PUSH1 0 REVERT
		[]byte{0x60, 0x00, 0x60, 0x00, 0xFD},
	)

	errMsg = errors.New("err: zero-fee DoS desc: <nil>")
	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, sigtx1)
	assert.Equal(t, errMsg, err)

	signer.sender = mn_active2
	err = protector.checkDoS(pool, hbtx2)
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, invtx2)
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, sigtx2)
	assert.Equal(t, errMsg, err)

	// Restore correct
	pool.currentState.SetCode(
		energi_params.Energi_MasternodeRegistry,
		// PUSH1 0 PUSH1 0 RETURN
		[]byte{0x60, 0x00, 0x60, 0x00, 0xF3},
	)

	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, sigtx1)
	assert.Equal(t, nil, err)

	signer.sender = mn_active2
	err = protector.checkDoS(pool, hbtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, sigtx2)
	assert.Equal(t, nil, err)

	// Test automatic cleanup
	signer.sender = mn_inactive

	_ = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, 2, len(protector.mnHeartbeats))
	assert.Equal(t, 2, len(protector.mnInvalidations))
	assert.Equal(t, 2, len(protector.mnCheckpoints))

	adjust_time = time.Duration(21)*time.Minute + time.Second
	_ = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, 0, len(protector.mnHeartbeats))
	assert.Equal(t, 0, len(protector.mnInvalidations))
	assert.Equal(t, 2, len(protector.mnCheckpoints))

	adjust_time = time.Duration(30) * time.Minute
	_ = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, 0, len(protector.mnHeartbeats))
	assert.Equal(t, 0, len(protector.mnInvalidations))
	assert.Equal(t, 2, len(protector.mnCheckpoints))

	adjust_time = time.Duration(30)*time.Minute + time.Second
	protector.nextCleanup = now
	_ = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, 0, len(protector.mnHeartbeats))
	assert.Equal(t, 0, len(protector.mnInvalidations))
	assert.Equal(t, 0, len(protector.mnCheckpoints))
}

func TestZeroFeeProtectorMigration(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	now := time.Now() // It can be fixed
	adjust_time := time.Duration(0)

	protector := newZeroFeeProtector()
	protector.timeNow = func() time.Time {
		return now.Add(adjust_time)
	}

	pool := &TxPool{}

	signer := &fakeSigner{}
	pool.signer = signer
	signer.sender = common.HexToAddress("0x0000000000000000000000000000000022345678")

	testdb := ethdb.NewMemDatabase()
	gspec := &Genesis{
		Config: params.TestnetChainConfig,
	}
	gspec.MustCommit(testdb)
	engine := ethash.NewFaker()

	chain, err := NewBlockChain(
		testdb, nil, gspec.Config,
		engine, vm.Config{}, nil)
	assert.Empty(t, err)
	defer chain.Stop()
	pool.chain = chain
	pool.currentState, _ = chain.State()

	migration_abi, err := abi.JSON(strings.NewReader(energi_abi.Gen2MigrationABI))
	assert.Empty(t, err)
	claim1Call, err := migration_abi.Pack(
		"claim", big.NewInt(1), common.Address{}, uint8(0), common.Hash{}, common.Hash{})
	assert.Empty(t, err)
	claim2Call, err := migration_abi.Pack(
		"claim", big.NewInt(2), common.Address{}, uint8(0), common.Hash{}, common.Hash{})
	assert.Empty(t, err)

	claim1 := types.NewTransaction(
		1, energi_params.Energi_MigrationContract, common.Big0, 100000, common.Big0, claim1Call)
	claim2 := types.NewTransaction(
		1, energi_params.Energi_MigrationContract, common.Big0, 100000, common.Big0, claim2Call)

	pool.currentState.SetCode(
		energi_params.Energi_MigrationContract,
		// PUSH1 1 PUSH1 0 MSTORE PUSH1 1 PUSH1 0 RETURN
		[]byte{0x60, 0x01, 0x60, 0x00, 0x52, 0x60, 0x20, 0x60, 0x00, 0xF3},
	)

	// Initial
	err = protector.checkDoS(pool, claim1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, claim2)
	assert.Equal(t, nil, err)

	// Before reset
	adjust_time = time.Duration(2) * time.Minute

	err = protector.checkDoS(pool, claim1)
	errMsg := errors.New("migrationErr: zero-fee DoS desc: time interval is less than the required")
	assert.Equal(t, errMsg, err)
	err = protector.checkDoS(pool, claim2)
	assert.Equal(t, errMsg, err)

	// After reset + return 0
	adjust_time = time.Duration(3) * time.Minute

	err = protector.checkDoS(pool, claim1)
	assert.Equal(t, nil, err)

	pool.currentState.SetCode(
		energi_params.Energi_MigrationContract,
		// PUSH1 0 PUSH1 0 MSTORE PUSH1 1 PUSH1 0 RETURN
		[]byte{0x60, 0x00, 0x60, 0x00, 0x52, 0x60, 0x20, 0x60, 0x00, 0xF3},
	)

	errMsg = errors.New("migrationErr: zero-fee DoS desc: already claimed")
	err = protector.checkDoS(pool, claim2)
	assert.Equal(t, errMsg, err)

	// After reset + revert 1
	adjust_time = time.Duration(6) * time.Minute

	pool.currentState.SetCode(
		energi_params.Energi_MigrationContract,
		// PUSH1 1 PUSH1 0 MSTORE PUSH1 1 PUSH1 0 RETURN
		[]byte{0x60, 0x01, 0x60, 0x00, 0x52, 0x60, 0x20, 0x60, 0x00, 0xF3},
	)

	err = protector.checkDoS(pool, claim1)
	assert.Equal(t, nil, err)

	pool.currentState.SetCode(
		energi_params.Energi_MigrationContract,
		// PUSH1 1 PUSH1 0 MSTORE PUSH1 1 PUSH1 0 REVERT
		[]byte{0x60, 0x01, 0x60, 0x00, 0x52, 0x60, 0x20, 0x60, 0x00, 0xFD},
	)

	errMsg = errors.New("migrationErr: zero-fee DoS desc: <nil>")
	err = protector.checkDoS(pool, claim2)
	assert.Equal(t, errMsg, err)
}
