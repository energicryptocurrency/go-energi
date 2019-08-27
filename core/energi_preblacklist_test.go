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
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/assert"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

/*
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
*/

func TestPreBlacklist(t *testing.T) {
	t.Parallel()
	log.Root().SetHandler(log.StdoutHandler)

	now := time.Now() // It can be fixed
	adjust_time := time.Duration(0)

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

	pool := NewTxPool(TxPoolConfig{}, gspec.Config, chain)
	prebl := pool.preBlacklist
	prebl.timeNow = func() time.Time {
		return now.Add(adjust_time)
	}

	bladdr1 := common.HexToAddress("0x1111")
	bladdr2 := common.HexToAddress("0x2222")
	sender := common.HexToAddress("0x3333")
	wladdr1 := common.HexToAddress("0x4444")

	signer := &fakeSigner{}
	pool.signer = signer
	signer.sender = sender

	defer chain.Stop()
	defer pool.Stop()
	pool.chain = chain
	pool.currentState, _ = chain.State()

	blreg_abi, err := abi.JSON(strings.NewReader(energi_abi.IBlacklistRegistryABI))
	assert.Empty(t, err)
	propose1Call, err := blreg_abi.Pack("propose", bladdr1)
	assert.Empty(t, err)
	propose2Call, err := blreg_abi.Pack("propose", bladdr2)
	assert.Empty(t, err)
	proposeWLCall, err := blreg_abi.Pack("propose", wladdr1)
	assert.Empty(t, err)
	revokeCall, err := blreg_abi.Pack("proposeRevoke", bladdr1)
	assert.Empty(t, err)

	propose1 := types.NewTransaction(
		1, energi_params.Energi_BlacklistRegistry, common.Big0, 1000000, common.Big0, propose1Call)
	propose2 := types.NewTransaction(
		2, energi_params.Energi_BlacklistRegistry, common.Big0, 1000000, common.Big0, propose2Call)
	proposeWL := types.NewTransaction(
		3, energi_params.Energi_BlacklistRegistry, common.Big0, 1000000, common.Big0, proposeWLCall)
	revoke := types.NewTransaction(
		4, energi_params.Energi_BlacklistRegistry, common.Big0, 1000000, common.Big0, revokeCall)
	proposePaid := types.NewTransaction(
		1, energi_params.Energi_BlacklistRegistry, common.Big0, 1000000, common.Big1, propose1Call)

	pool.currentState.SetCode(
		energi_params.Energi_BlacklistRegistry,
		// PUSH1 1 PUSH1 0 MSTORE PUSH1 1 PUSH1 0 RETURN
		[]byte{0x60, 0x01, 0x60, 0x00, 0x52, 0x60, 0x20, 0x60, 0x00, 0xF3},
	)
	pool.currentState.SetState(
		energi_params.Energi_Whitelist,
		wladdr1.Hash(),
		common.BytesToHash([]byte{0x01}),
	)
	assert.True(t, IsWhitelisted(pool.currentState, wladdr1))

	log.Trace("Make sure removed in pool")
	signer.sender = bladdr1
	err = pool.AddLocal(revoke)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(pool.queue[bladdr1].Flatten()))
	signer.sender = sender

	log.Trace("Initial")
	err = prebl.processTx(pool, revoke)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(prebl.proposed))
	assert.Equal(t, 1, len(pool.queue[bladdr1].Flatten()))

	err = prebl.processTx(pool, propose1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(prebl.proposed))
	if qb, ok := pool.queue[bladdr1]; ok {
		assert.Equal(t, 1, len(qb.Flatten()))
	}

	adjust_time = time.Duration(10) * time.Minute

	err = prebl.processTx(pool, propose2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(prebl.proposed))

	err = prebl.processTx(pool, proposeWL)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(prebl.proposed))
	assert.True(t, IsWhitelisted(pool.currentState, wladdr1))

	// Non-EBI
	err = prebl.processTx(pool, proposePaid)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(prebl.proposed))

	log.Trace("Check if filtered properly")
	signer.sender = bladdr1
	err = prebl.processTx(pool, revoke)
	assert.Equal(t, ErrPreBlacklist, err)

	signer.sender = sender
	err = prebl.processTx(pool, revoke)
	assert.Equal(t, nil, err)

	signer.sender = bladdr2
	err = prebl.processTx(pool, revoke)
	assert.Equal(t, ErrPreBlacklist, err)

	signer.sender = sender

	log.Trace("Check block filter hook")
	blocks := make(types.Blocks, 0, 3)
	blocks = append(blocks, types.NewBlockWithHeader(&types.Header{Coinbase: sender}))
	blocks = append(blocks, types.NewBlockWithHeader(&types.Header{Coinbase: bladdr1}))
	blocks = append(blocks, types.NewBlockWithHeader(&types.Header{Coinbase: sender}))
	blocks = pool.PreBlacklistHook(blocks)
	assert.Equal(t, 1, len(blocks))
	assert.Equal(t, sender, blocks[0].Coinbase())

	log.Trace("After timeout")
	adjust_time = pbPeriod + time.Minute

	err = prebl.processTx(pool, revoke)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(prebl.proposed))

	err = prebl.processTx(pool, propose1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(prebl.proposed))

	err = prebl.processTx(pool, propose2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(prebl.proposed))

	adjust_time *= time.Duration(2)
	err = prebl.processTx(pool, revoke)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(prebl.proposed))

	log.Trace("Ignore not valid proposals")

	pool.currentState.SetCode(
		energi_params.Energi_BlacklistRegistry,
		// PUSH1 1 PUSH1 0 MSTORE PUSH1 1 PUSH1 0 REVERT
		[]byte{0x60, 0x01, 0x60, 0x00, 0x52, 0x60, 0x20, 0x60, 0x00, 0xFD},
	)

	err = prebl.processTx(pool, propose1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(prebl.proposed))
}
