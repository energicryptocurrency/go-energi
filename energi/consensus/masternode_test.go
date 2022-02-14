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
	"math/big"
	"strings"
	"testing"

	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/core/vm"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/ethdb"
	"github.com/energicryptocurrency/energi/params"

	"github.com/stretchr/testify/assert"
)

func TestMasternodeList(t *testing.T) {
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

	header := chain.GetHeaderByHash(genesis.Hash())
	assert.NotEmpty(t, header)

	blstate, err := chain.StateAt(header.Root)
	assert.Empty(t, err)

	err = engine.processConsensusGasLimits(chain, header, blstate)
	assert.Empty(t, err)

	owner_addr1 := common.HexToAddress("0x0000000000000000000000000000000012345678")
	owner_addr2 := common.HexToAddress("0x0000000000000000000000000000000012345679")
	mn_addr1 := common.HexToAddress("0x0000000000000000000000000000000022345678")
	mn_addr2 := common.HexToAddress("0x0000000000000000000000000000000022345679")

	collateral := new(big.Int).Mul(big.NewInt(100000), big.NewInt(1e18))
	blstate.AddBalance(owner_addr1, collateral)
	blstate.AddBalance(owner_addr2, collateral)
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
		owner_addr1,
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
	msg = types.NewMessage(
		owner_addr2,
		&energi_params.Energi_MasternodeToken,
		0,
		collateral,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	evm = engine.createEVM(msg, chain, header, blstate)
	gp = new(core.GasPool).AddGas(engine.callGas)
	// log.Trace("depositCollateral")
	_, _, _, _ = core.ApplyMessage(evm, msg, gp)
	//---
	mnreg_abi, _ := abi.JSON(strings.NewReader(energi_abi.IMasternodeRegistryV2ABI))
	callData, err = mnreg_abi.Pack("announce", mn_addr1, uint32(130<<24), [2][32]byte{})
	assert.Empty(t, err)
	msg = types.NewMessage(
		owner_addr1,
		&energi_params.Energi_MasternodeRegistry,
		0,
		common.Big0,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	evm = engine.createEVM(msg, chain, header, blstate)
	gp.AddGas(engine.callGas)
	// log.Trace("announce")
	_, _, _, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)

	callData, err = mnreg_abi.Pack("announce", mn_addr2, uint32(130<<24), [2][32]byte{})
	assert.Empty(t, err)
	msg = types.NewMessage(
		owner_addr2,
		&energi_params.Energi_MasternodeRegistry,
		0,
		common.Big0,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	evm = engine.createEVM(msg, chain, header, blstate)
	gp.AddGas(engine.callGas)
	// log.Trace("announce")
	_, _, _, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)

	//---
	err = engine.processMasternodes(chain, header, blstate)
	assert.Empty(t, err)
	header.Root, err = blstate.Commit(true)
	assert.Empty(t, err)

	//---
	empty := common.Hash{}
	assert.True(t, blstate.GetState(energi_params.Energi_MasternodeList, mn_addr1.Hash()) != empty)
	assert.True(t, blstate.GetState(energi_params.Energi_MasternodeList, mn_addr2.Hash()) != empty)
	assert.True(t, blstate.GetState(energi_params.Energi_MasternodeList, owner_addr1.Hash()) == empty)

	//---
	callData, err = mnreg_abi.Pack("denounce", mn_addr2)
	assert.Empty(t, err)
	msg = types.NewMessage(
		owner_addr2,
		&energi_params.Energi_MasternodeRegistry,
		0,
		common.Big0,
		engine.callGas,
		common.Big0,
		callData,
		false,
	)
	evm = engine.createEVM(msg, chain, header, blstate)
	gp.AddGas(engine.callGas)
	// log.Trace("denounce")
	_, _, _, err = core.ApplyMessage(evm, msg, gp)
	assert.Empty(t, err)

	err = engine.processMasternodes(chain, header, blstate)
	assert.Empty(t, err)
	header.Root, err = blstate.Commit(true)
	assert.Empty(t, err)

	//---
	assert.True(t, blstate.GetState(energi_params.Energi_MasternodeList, mn_addr1.Hash()) != empty)
	assert.True(t, blstate.GetState(energi_params.Energi_MasternodeList, mn_addr2.Hash()) == empty)
	assert.True(t, blstate.GetState(energi_params.Energi_MasternodeList, owner_addr1.Hash()) == empty)
}
