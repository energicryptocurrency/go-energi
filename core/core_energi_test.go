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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

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
	validateCall, err := mnreg_abi.Pack("validate", common.Address{})
	assert.Empty(t, err)

	res := false

	//
	res = IsValidZeroFee(types.NewTransaction(
		0,
		common.Address{},
		common.Big0,
		50000,
		common.Big0,
		[]byte{},
	))
	assert.False(t, res, "Simple zero")
	//
	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		50000,
		common.Big0,
		claimCall,
	))
	assert.True(t, res, "Valid migration claim")
	//
	res = IsValidZeroFee(types.NewTransaction(
		2,
		energi_params.Energi_MigrationContract,
		common.Big0,
		50000,
		common.Big0,
		claimCall,
	))
	assert.True(t, res, "Valid migration claim - other nonce")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big1,
		50000,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid migration claim - amount")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		50001,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid migration claim - gas")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		50000,
		common.Big1,
		claimCall,
	))
	assert.False(t, res, "Invalid migration claim - gas price")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MasternodeRegistry,
		common.Big0,
		50000,
		common.Big0,
		claimCall,
	))
	assert.False(t, res, "Invalid migration claim - dst")

	res = IsValidZeroFee(types.NewTransaction(
		0,
		energi_params.Energi_MigrationContract,
		common.Big0,
		50000,
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
		validateCall,
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
		validateCall,
	))
	assert.False(t, res, "Invalid MN validate - dst")
}
