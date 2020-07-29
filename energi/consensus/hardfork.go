// Copyright 20 The Energi Core Authors
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
	"bytes"
	"math/big"

	"energi.world/core/gen3/common"
	"energi.world/core/gen3/core"
	"energi.world/core/gen3/core/state"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/log"

	energi_common "energi.world/core/gen3/energi/common"
	energi_params "energi.world/core/gen3/energi/params"
)

func (e *Energi) checkLatestHardforks(
	chain ChainReader,
	header *types.Header,
	statedb *state.StateDB,
) (err error) {
	enumerateData, err := e.hardforkAbi.Pack("enumerate")
	if err != nil {
		log.Error("Fail to prepare enumerate() call", "err", err)
		return err
	}

	msg := types.NewMessage(
		energi_params.Energi_HardforkRegistry,
		&energi_params.Energi_HardforkRegistry,
		0,
		common.Big0,
		e.unlimitedGas,
		common.Big0,
		enumerateData,
		false,
	)

	revID := statedb.Snapshot()
	evm := e.createEVM(msg, chain, header, statedb)
	gp := core.GasPool(e.unlimitedGas)
	output, gasUsed, _, err := core.ApplyMessage(evm, msg, &gp)
	statedb.RevertToSnapshot(revID)
	if err != nil {
		log.Error("Failed in enumerate() call", "err", err)
		return err
	}

	if gasUsed > e.callGas {
		log.Warn("HardforkRegistry::enumerate() took excessive gas",
			"gas", gasUsed, "limit", e.callGas)
	}

	hardforkBlocks := make([]*big.Int, 0, 100)
	err = e.hardforkAbi.Unpack(&hardforkBlocks, "enumerate", output)
	if err != nil {
		log.Error("Failed to unpack enumerate() call", "err", err)
		return err
	}

	hardforks := make([]*big.Int, 0, len(hardforkBlocks))
	for i, hfBlock := range hardforkBlocks {
		num := new(big.Int).Add(hfBlock, chain.Config().HFFinalizationPeriod)
		if num.Cmp(header.Number) < 1 {
			// Selects all the hardforks that are within the current hardfork finalization period.
			hardforks = append(hardforks, hardforkBlocks[i])
		}
	}

	// Fetchs the hardfork information identif
	for _, hfblock := range hardforks {
		queryData, err := e.hardforkAbi.Pack("getByBlockNo", hfblock)
		if err != nil {
			log.Error("Fail to prepare getByBlockNo() call", "err", err)
			return err
		}

		msg := types.NewMessage(
			energi_params.Energi_HardforkRegistry,
			&energi_params.Energi_HardforkRegistry,
			0,
			common.Big0,
			e.unlimitedGas,
			common.Big0,
			queryData,
			false,
		)

		revID := statedb.Snapshot()
		evm := e.createEVM(msg, chain, header, statedb)
		gp := core.GasPool(e.unlimitedGas)
		output, gasUsed, _, err := core.ApplyMessage(evm, msg, &gp)
		statedb.RevertToSnapshot(revID)
		if err != nil {
			log.Error("Failed in getByBlockNo() call", "err", err)
			return err
		}

		if gasUsed > e.callGas {
			log.Warn("HardforkRegistry::getByBlockNo() took excessive gas",
				"gas", gasUsed, "limit", e.callGas)
		}

		type HFInfo struct {
			Name       [32]byte
			BlockHash  [32]byte
			SwFeatures *big.Int
		}

		hfInfo := new(HFInfo)
		err = e.hardforkAbi.Unpack(&hfInfo, "getByBlockNo", output)
		if err != nil {
			log.Error("Failed to unpack getByBlockNo() call", "err", err)
			return err
		}

		emptyBytes := [32]byte{}
		if bytes.Compare(hfInfo.BlockHash[:], emptyBytes[:]) == 0 {
			// BlockHash not set.
			log.Info("Hardfork could be finalized", "block", hfblock,
				"hardforkName", energi_common.DecodeToString(hfInfo.Name),
				"remainingBlocks", new(big.Int).Sub(header.Number, hfblock),
			)
		} else {
			// BlockHash already. Hardfork already finalised.
			log.Info("Hardfork already finalized", "block", hfblock,
				"hardforkName", energi_common.DecodeToString(hfInfo.Name),
				"blockHash", common.BytesToHash(hfInfo.BlockHash[:]),
			)
		}
	}
	return
}
