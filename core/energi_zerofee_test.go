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
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/assert"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

func TestIsValidZeroFee(t *testing.T) {
	t.Parallel()
	log.Root().SetHandler(log.StdoutHandler)

	migration_abi, err := abi.JSON(strings.NewReader(energi_abi.Gen2MigrationABI))
	assert.Empty(t, err)
	claimCall, err := migration_abi.Pack(
		"claim", common.Big0, common.Address{}, uint8(0), common.Hash{}, common.Hash{})
	assert.Empty(t, err)

	mnreg_abi, err := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryABI))
	assert.Empty(t, err)
	heartbeatCall, err := mnreg_abi.Pack("heartbeat", common.Big1, common.Hash{}, common.Big0)
	assert.Empty(t, err)
	invalidateCall, err := mnreg_abi.Pack("invalidate", common.Address{})
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
	log.Root().SetHandler(log.StdoutHandler)

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
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(testdb))
	pool.currentState = statedb

	mn_active1 := common.HexToAddress("0x0000000000000000000000000000000022345678")
	mn_active2 := common.HexToAddress("0x0000000000000000000000000000000022345679")
	mn_inactive := common.HexToAddress("0x0000000000000000000000000000000022345680")

	statedb.SetState(
		energi_params.Energi_MasternodeList,
		mn_active1.Hash(),
		mn_inactive.Hash(),
	)
	statedb.SetState(
		energi_params.Energi_MasternodeList,
		mn_active2.Hash(),
		mn_inactive.Hash(),
	)

	mnreg_abi, err := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryABI))
	assert.Empty(t, err)
	heartbeatCall, err := mnreg_abi.Pack("heartbeat", common.Big1, common.Hash{}, common.Big0)
	assert.Empty(t, err)
	invalidateCall, err := mnreg_abi.Pack("invalidate", common.Address{})
	assert.Empty(t, err)

	hbtx0 := types.NewTransaction(
		1, common.Address{}, common.Big0, 100000, common.Big0, heartbeatCall)
	invtx0 := types.NewTransaction(
		1, common.Address{}, common.Big0, 100000, common.Big0, invalidateCall)
	hbtx1 := types.NewTransaction(
		1, common.Address{}, common.Big0, 100000, common.Big0, heartbeatCall)
	invtx1 := types.NewTransaction(
		1, common.Address{}, common.Big0, 100000, common.Big0, invalidateCall)
	hbtx2 := types.NewTransaction(
		1, common.Address{}, common.Big0, 100000, common.Big0, heartbeatCall)
	invtx2 := types.NewTransaction(
		1, common.Address{}, common.Big0, 100000, common.Big0, invalidateCall)

	// Inactive MN
	signer.sender = mn_inactive
	err = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, ErrZeroFeeDoS, err)
	err = protector.checkDoS(pool, invtx0)
	assert.Equal(t, ErrZeroFeeDoS, err)

	// Active MN
	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, nil, err)

	// Active MN repeat interval
	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, ErrZeroFeeDoS, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, ErrZeroFeeDoS, err)

	// Active another MV
	signer.sender = mn_active2
	err = protector.checkDoS(pool, hbtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx2)
	assert.Equal(t, nil, err)

	// Active MN after first period is over
	adjust_time = time.Duration(5) * time.Minute

	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, ErrZeroFeeDoS, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, nil, err)

	signer.sender = mn_active2
	err = protector.checkDoS(pool, hbtx2)
	assert.Equal(t, ErrZeroFeeDoS, err)
	err = protector.checkDoS(pool, invtx2)
	assert.Equal(t, nil, err)

	// Active MN after second period is over
	adjust_time = time.Duration(30) * time.Minute

	signer.sender = mn_active1
	err = protector.checkDoS(pool, hbtx1)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx1)
	assert.Equal(t, nil, err)

	signer.sender = mn_active2
	err = protector.checkDoS(pool, hbtx2)
	assert.Equal(t, nil, err)
	err = protector.checkDoS(pool, invtx2)
	assert.Equal(t, nil, err)

	// Test automatic cleanup
	signer.sender = mn_inactive

	err = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, 2, len(protector.mnHeartbeats))
	assert.Equal(t, 2, len(protector.mnInvalidations))

	adjust_time = time.Duration(35) * time.Minute
	err = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, 2, len(protector.mnHeartbeats))
	assert.Equal(t, 0, len(protector.mnInvalidations))

	adjust_time = time.Duration(59) * time.Minute
	err = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, 2, len(protector.mnHeartbeats))
	assert.Equal(t, 0, len(protector.mnInvalidations))

	adjust_time = time.Duration(61) * time.Minute
	err = protector.checkDoS(pool, hbtx0)
	assert.Equal(t, 0, len(protector.mnHeartbeats))
	assert.Equal(t, 0, len(protector.mnInvalidations))
}

func TestZeroFeeProtectorMigration(t *testing.T) {
	t.Parallel()
	log.Root().SetHandler(log.StdoutHandler)

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
	assert.Equal(t, ErrZeroFeeDoS, err)
	err = protector.checkDoS(pool, claim2)
	assert.Equal(t, ErrZeroFeeDoS, err)

	// After reset + return 0
	adjust_time = time.Duration(3) * time.Minute

	err = protector.checkDoS(pool, claim1)
	assert.Equal(t, nil, err)

	pool.currentState.SetCode(
		energi_params.Energi_MigrationContract,
		// PUSH1 0 PUSH1 0 MSTORE PUSH1 1 PUSH1 0 RETURN
		[]byte{0x60, 0x00, 0x60, 0x00, 0x52, 0x60, 0x20, 0x60, 0x00, 0xF3},
	)

	err = protector.checkDoS(pool, claim2)
	assert.Equal(t, ErrZeroFeeDoS, err)

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

	err = protector.checkDoS(pool, claim2)
	assert.Equal(t, ErrZeroFeeDoS, err)
}
