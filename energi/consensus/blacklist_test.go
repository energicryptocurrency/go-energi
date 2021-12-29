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
	"crypto/ecdsa"
	crand "crypto/rand"
	"math/big"
	"strings"
	"testing"

	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/core/rawdb"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/core/vm"
	"github.com/energicryptocurrency/energi/crypto"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/ethdb"
	"github.com/energicryptocurrency/energi/params"

	"github.com/stretchr/testify/assert"
)

func TestBlacklist(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	testdb := ethdb.NewMemDatabase()
	engine := New(&params.EnergiConfig{}, testdb)

	engine.testing = true

	chainConfig := *params.EnergiTestnetChainConfig
	chainConfig.Energi = &params.EnergiConfig{}

	var (
		gspec = &core.Genesis{
			Config:     &chainConfig,
			GasLimit:   8000000,
			Timestamp:  1000,
			Difficulty: big.NewInt(1),
			Coinbase:   energi_params.Energi_Treasury,
			Xfers:      core.DeployEnergiGovernance(&chainConfig),
		}
		genesis = gspec.MustCommit(testdb)
	)

	chain, err := core.NewBlockChain(testdb, nil, &chainConfig, engine, vm.Config{}, nil)
	assert.Empty(t, err)
	defer chain.Stop()

	//--
	_, err = chain.InsertChain([]*types.Block{genesis})
	assert.Empty(t, err)

	header := &types.Header{
		Number:     new(big.Int).Add(genesis.Number(), common.Big1),
		ParentHash: genesis.Hash(),
		Root:       genesis.Root(),
		GasLimit:   genesis.GasLimit(),
		Time:       genesis.Time(),
		Difficulty: genesis.Difficulty(),
	}
	assert.NotEmpty(t, header)

	blstate, err := chain.StateAt(header.Root)
	assert.Empty(t, err)

	err = engine.processConsensusGasLimits(chain, header, blstate)
	assert.Empty(t, err)

	blacklist_key1, _ := ecdsa.GenerateKey(crypto.S256(), crand.Reader)

	blacklist_addr1 := crypto.PubkeyToAddress(blacklist_key1.PublicKey)
	blacklist_addr2 := common.HexToAddress("0x0000000000000000000000000000000012345679")
	owner_addr := common.HexToAddress("0x0000000000000000000000000000000022345678")

	amt := big.NewInt(100)
	collateral := new(big.Int).Mul(big.NewInt(100000), big.NewInt(1e18))
	blstate.SetBalance(owner_addr, collateral)
	blstate.SetBalance(blacklist_addr1, amt)
	blstate.SetBalance(blacklist_addr2, amt)
	header.Root, err = blstate.Commit(true)
	assert.Empty(t, err)
	err = blstate.Database().TrieDB().Commit(header.Root, true)
	assert.Empty(t, err)
	blstate, err = chain.StateAt(header.Root)
	assert.Empty(t, err)

	//---
	mntoken_abi, _ := abi.JSON(strings.NewReader(energi_abi.IMasternodeTokenABI))
	callData, err := mntoken_abi.Pack("depositCollateral")
	assert.Empty(t, err)
	msg := types.NewMessage(
		owner_addr,
		&energi_params.Energi_MasternodeToken,
		0,
		collateral,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	evm := engine.createEVM(msg, chain, header, blstate)
	gp := new(core.GasPool).AddGas(engine.callGas)
	// log.Trace("depositCollateral")
	_, _, _, _ = core.ApplyMessage(evm, msg, gp)
	//---
	mnreg_abi, _ := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryV2ABI))
	callData, err = mnreg_abi.Pack("announce", blacklist_addr1, uint32(130<<24), [2][32]byte{})
	assert.Empty(t, err)
	msg = types.NewMessage(
		owner_addr,
		&energi_params.Energi_MasternodeRegistry,
		0,
		common.Big0,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.callGas)
	// log.Trace("announce")
	_, _, _, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)

	header.Number.Add(header.Number, common.Big1)
	header.Time += 2*24*60*60 + 1
	_ = engine.createEVM(msg, chain, header, blstate)
	//---

	//====================================
	// log.Info("Test: no change")
	err = engine.processBlacklists(chain, header, blstate)
	assert.Empty(t, err)
	assert.True(t, core.CanTransfer(blstate, blacklist_addr1, common.Big1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr2, common.Big1))
	txs, receipts, err := engine.processDrainable(chain, header, blstate, nil, nil)
	assert.Empty(t, err)
	assert.Empty(t, txs)
	assert.Empty(t, receipts)
	assert.Equal(t, blstate.GetBalance(blacklist_addr1).String(), amt.String())
	assert.Equal(t, blstate.GetBalance(blacklist_addr2).String(), amt.String())
	header.Root, err = blstate.Commit(true)
	assert.Empty(t, err)
	err = blstate.Database().TrieDB().Commit(header.Root, true)
	assert.Empty(t, err)
	blstate, err = chain.StateAt(header.Root)
	assert.Empty(t, err)
	evm = engine.createEVM(msg, chain, header, blstate)

	//====================================
	// log.Info("Test: blacklist")
	blacklist_abi, _ := abi.JSON(strings.NewReader(energi_abi.IBlacklistRegistryABI))
	callData, err = blacklist_abi.Pack("propose", blacklist_addr1)
	assert.Empty(t, err)
	fee := new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18))
	blstate.SetBalance(owner_addr, fee)
	msg = types.NewMessage(
		owner_addr,
		&energi_params.Energi_BlacklistRegistry,
		0,
		fee,
		engine.xferGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.xferGas)
	// log.Trace("propose")
	output, _, failed, err := core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)
	assert.Empty(t, failed)

	var enforce_address common.Address
	err = blacklist_abi.Unpack(&enforce_address, "propose", output)
	assert.Empty(t, err)

	proposal_abi, _ := abi.JSON(strings.NewReader(energi_abi.IProposalABI))
	callData, err = proposal_abi.Pack("voteAccept")
	assert.Empty(t, err)
	msg = types.NewMessage(
		owner_addr,
		&enforce_address,
		0,
		common.Big0,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.callGas)
	// log.Trace("voteAccept")
	output, _, _, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)
	assert.Empty(t, output)

	err = engine.processBlacklists(chain, header, blstate)
	assert.Empty(t, err)
	assert.True(t, core.IsBlacklisted(blstate, blacklist_addr1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr1, common.Big0))
	assert.False(t, core.CanTransfer(blstate, blacklist_addr1, common.Big1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr2, common.Big1))
	txs, receipts, err = engine.processDrainable(chain, header, blstate, nil, nil)
	assert.Empty(t, err)
	assert.Empty(t, txs)
	assert.Empty(t, receipts)
	assert.Equal(t, blstate.GetBalance(blacklist_addr1).String(), amt.String())
	assert.Equal(t, blstate.GetBalance(blacklist_addr2).String(), amt.String())
	header.Root, err = blstate.Commit(true)
	assert.Empty(t, err)
	err = blstate.Database().TrieDB().Commit(header.Root, true)
	assert.Empty(t, err)
	blstate, err = chain.StateAt(header.Root)
	assert.Empty(t, err)

	// log.Info("Test Bug: in cleanup untouched when just referenced")
	blstate.AddBalance(owner_addr, common.Big1)
	assert.True(t, core.CanTransfer(blstate, blacklist_addr1, common.Big0))
	assert.False(t, core.CanTransfer(blstate, blacklist_addr1, common.Big1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr2, common.Big1))
	header.Root, err = blstate.Commit(true)
	assert.Empty(t, err)
	blstate.Database().TrieDB().Reference(header.Root, common.Hash{})

	blstate, err = chain.StateAt(header.Root)
	assert.Empty(t, err)
	err = engine.processBlacklists(chain, header, blstate)
	assert.Empty(t, err)
	assert.True(t, core.CanTransfer(blstate, blacklist_addr1, common.Big0))
	assert.False(t, core.CanTransfer(blstate, blacklist_addr1, common.Big1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr2, common.Big1))
	blstate.Database().TrieDB().Dereference(header.Root)

	//====================================
	// log.Trace("coinbase blacklist")
	rawdb.WriteHeader(testdb, header)

	header2 := header
	header2.ParentHash = header.Hash()
	header2.Number = new(big.Int).Add(header.Number, common.Big1)
	header2.Coinbase = blacklist_addr1

	sighash := engine.SignatureHash(header2)
	header2.Signature, err = crypto.Sign(sighash.Bytes(), blacklist_key1)
	assert.Empty(t, err)

	assert.True(t, core.IsBlacklisted(blstate, blacklist_addr1))
	assert.Equal(t, errBlacklistedCoinbase, engine.VerifySeal(chain, header2))

	//====================================
	// log.Info("Test: drain")
	evm = engine.createEVM(msg, chain, header, blstate)
	callData, err = blacklist_abi.Pack("proposeDrain", blacklist_addr1)
	assert.Empty(t, err)
	fee = new(big.Int).Mul(big.NewInt(100), big.NewInt(1e18))
	blstate.SetBalance(owner_addr, fee)
	msg = types.NewMessage(
		owner_addr,
		&energi_params.Energi_BlacklistRegistry,
		0,
		fee,
		engine.xferGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.xferGas)
	// log.Trace("proposeDrain")
	output, _, failed, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)
	assert.Empty(t, failed)

	var drain_address common.Address
	err = blacklist_abi.Unpack(&drain_address, "proposeDrain", output)
	assert.Empty(t, err)

	callData, err = proposal_abi.Pack("voteAccept")
	assert.Empty(t, err)
	msg = types.NewMessage(
		owner_addr,
		&drain_address,
		0,
		common.Big0,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.callGas)
	// log.Trace("voteAccept")
	output, _, _, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)
	assert.Empty(t, output)

	err = engine.processBlacklists(chain, header, blstate)
	assert.Empty(t, err)
	assert.True(t, core.IsBlacklisted(blstate, blacklist_addr1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr1, common.Big0))
	assert.False(t, core.CanTransfer(blstate, blacklist_addr1, common.Big1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr2, common.Big1))
	txs, receipts, err = engine.processDrainable(chain, header, blstate, nil, nil)
	assert.Empty(t, err)
	assert.Equal(t, 2, len(txs))
	assert.Equal(t, 2, len(receipts))
	assert.Equal(t, blstate.GetBalance(blacklist_addr1).String(), common.Big0.String())
	assert.Equal(t, blstate.GetBalance(blacklist_addr2).String(), amt.String())
	header.Root, err = blstate.Commit(true)
	assert.Empty(t, err)
	err = blstate.Database().TrieDB().Commit(header.Root, true)
	assert.Empty(t, err)
	blstate, err = chain.StateAt(header.Root)
	assert.Empty(t, err)
	evm = engine.createEVM(msg, chain, header, blstate)

	//====================================
	// log.Info("Test: no change")
	err = engine.processBlacklists(chain, header, blstate)
	assert.Empty(t, err)
	assert.False(t, core.IsBlacklisted(blstate, blacklist_addr1))
	assert.False(t, core.IsBlacklisted(blstate, blacklist_addr2))
	assert.False(t, core.CanTransfer(blstate, blacklist_addr1, common.Big1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr2, common.Big1))
	txs, receipts, err = engine.processDrainable(chain, header, blstate, nil, nil)
	assert.Empty(t, err)
	assert.Empty(t, txs)
	assert.Empty(t, receipts)
	assert.Equal(t, blstate.GetBalance(blacklist_addr1).String(), common.Big0.String())
	assert.Equal(t, blstate.GetBalance(blacklist_addr2).String(), amt.String())

	//====================================
	// log.Info("Test: whitelist")
	callData, err = blacklist_abi.Pack("propose", energi_params.Energi_TreasuryV1)
	assert.Empty(t, err)
	fee = new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18))
	blstate.SetBalance(owner_addr, fee)
	msg = types.NewMessage(
		owner_addr,
		&energi_params.Energi_BlacklistRegistry,
		0,
		fee,
		engine.xferGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.xferGas)
	// log.Trace("propose")
	output, _, failed, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)
	assert.Empty(t, failed)

	err = blacklist_abi.Unpack(&enforce_address, "propose", output)
	assert.Empty(t, err)

	callData, err = proposal_abi.Pack("voteAccept")
	assert.Empty(t, err)
	msg = types.NewMessage(
		owner_addr,
		&enforce_address,
		0,
		common.Big0,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.callGas)
	// log.Trace("voteAccept")
	output, _, _, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)
	assert.Empty(t, output)

	evm = engine.createEVM(msg, chain, header, blstate)
	callData, err = blacklist_abi.Pack("proposeDrain", energi_params.Energi_TreasuryV1)
	assert.Empty(t, err)
	fee = new(big.Int).Mul(big.NewInt(100), big.NewInt(1e18))
	blstate.SetBalance(owner_addr, fee)
	msg = types.NewMessage(
		owner_addr,
		&energi_params.Energi_BlacklistRegistry,
		0,
		fee,
		engine.xferGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.xferGas)
	// log.Trace("proposeDrain")
	output, _, failed, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)
	assert.Empty(t, failed)

	err = blacklist_abi.Unpack(&drain_address, "proposeDrain", output)
	assert.Empty(t, err)

	callData, err = proposal_abi.Pack("voteAccept")
	assert.Empty(t, err)
	msg = types.NewMessage(
		owner_addr,
		&drain_address,
		0,
		common.Big0,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	gp.AddGas(engine.callGas)
	// log.Trace("voteAccept")
	output, _, _, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)
	assert.Empty(t, output)

	blstate.AddBalance(energi_params.Energi_TreasuryV1, amt)

	err = engine.processBlacklists(chain, header, blstate)
	assert.Empty(t, err)
	assert.False(t, core.IsBlacklisted(blstate, blacklist_addr1))
	assert.False(t, core.CanTransfer(blstate, blacklist_addr1, common.Big1))
	assert.True(t, core.CanTransfer(blstate, blacklist_addr2, common.Big1))
	assert.True(t, core.CanTransfer(blstate, energi_params.Energi_TreasuryV1, common.Big1))
	txs, receipts, err = engine.processDrainable(chain, header, blstate, nil, nil)
	assert.Empty(t, err)
	assert.Empty(t, txs)
	assert.Empty(t, receipts)
	assert.Equal(t, blstate.GetBalance(blacklist_addr1).String(), common.Big0.String())
	assert.Equal(t, blstate.GetBalance(blacklist_addr2).String(), amt.String())
	// NOTE: whitelisted addresses must not be drainable!
	assert.NotEqual(t, blstate.GetBalance(energi_params.Energi_TreasuryV1).String(), common.Big0.String())
	header.Root, err = blstate.Commit(true)
	assert.Empty(t, err)
	err = blstate.Database().TrieDB().Commit(header.Root, true)
	assert.Empty(t, err)
	_, err = chain.StateAt(header.Root)
	assert.Empty(t, err)
}
