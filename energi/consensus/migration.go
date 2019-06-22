// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package consensus

import (
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/shengdoushi/base58"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	gasPerMigrationEntry uint64 = 100000
)

func (e *Energi) finalizeMigration(
	chain ChainReader,
	header *types.Header,
	statedb *state.StateDB,
	txs types.Transactions,
) error {
	if !header.IsGen2Migration() {
		return nil
	}

	if len(txs) != 1 {
		err := errors.New("Wrong number of migration block txs")
		log.Error("Failed to finalize migration", "err", err)
		return err
	}

	migration_abi, err := abi.JSON(strings.NewReader(energi_abi.Gen2MigrationABI))
	if err != nil {
		return err
	}

	callData, err := migration_abi.Pack("totalAmount")
	if err != nil {
		log.Error("Failed to prepare totalAmount() call", "err", err)
		return err
	}

	// totalAmount()
	msg := types.NewMessage(
		e.systemFaucet,
		&energi_params.Energi_MigrationContract,
		0,
		common.Big0,
		e.callGas,
		common.Big0,
		callData,
		false,
	)
	evm := e.createEVM(msg, chain, header, statedb)
	gp := new(core.GasPool).AddGas(e.callGas)
	output, _, _, err := core.ApplyMessage(evm, msg, gp)
	if err != nil {
		log.Error("Failed in totalAmount() call", "err", err)
		return err
	}

	//
	totalAmount := big.NewInt(0)
	err = migration_abi.Unpack(&totalAmount, "totalAmount", output)
	if err != nil {
		log.Error("Failed to unpack totalAmount() call", "err", err)
		return err
	}

	statedb.SetBalance(energi_params.Energi_MigrationContract, totalAmount)
	log.Warn("Setting Migration contract balance", "amount", totalAmount)

	return nil
}

func MigrationTx(
	signer types.Signer,
	header *types.Header,
	migration_file string,
	engine consensus.Engine,
) (res *types.Transaction) {
	e, ok := engine.(*Energi)
	if !ok {
		log.Error("Not Energi consensus engine")
		return nil
	}

	file, err := os.Open(migration_file)
	if err != nil {
		log.Error("Failed to open snapshot", "err", err)
		return nil
	}
	defer file.Close()

	snapshot, err := parseSnapshot(file)
	if err != nil {
		log.Error("Failed to parse snapshot", "err", err)
		return nil
	}

	owners, amounts := createSnapshotParams(snapshot)
	if owners == nil || amounts == nil {
		log.Error("Failed to create arguments")
		return nil
	}

	migration_abi, err := abi.JSON(strings.NewReader(energi_abi.Gen2MigrationABI))
	if err != nil {
		panic(err)
	}

	callData, err := migration_abi.Pack("setSnapshot", owners, amounts)
	if err != nil {
		panic(err)
	}

	gasLimit := gasPerMigrationEntry * uint64(len(owners))
	header.GasLimit = gasLimit
	header.Extra = make([]byte, 64)
	copy(header.Extra[32:], common.HexToHash(snapshot.Hash).Bytes())

	res = types.NewTransaction(
		uint64(0), // it should be the first transaction
		energi_params.Energi_MigrationContract,
		common.Big0,
		gasLimit,
		common.Big0,
		callData,
	)

	if e.signerFn == nil {
		log.Error("Signer is not set")
		return nil
	}

	if e.config == nil {
		log.Error("Engine config is not set")
		return nil
	}

	tx_hash := signer.Hash(res)
	tx_sig, err := e.signerFn(e.config.MigrationSigner, tx_hash.Bytes())
	if err != nil {
		log.Error("Failed to sign migration tx")
		return nil
	}

	res, err = res.WithSignature(signer, tx_sig)
	if err != nil {
		log.Error("Failed to pack migration tx")
		return nil
	}
	return
}

func createSnapshotParams(ss *snapshot) (owners []common.Address, amounts []*big.Int) {
	owners = make([]common.Address, len(ss.Txouts))
	amounts = make([]*big.Int, len(ss.Txouts))

	// NOTE: Gen 2 precision is 8, but Gen 3 is 18
	multiplier := big.NewInt(1e10)

	for i, info := range ss.Txouts {
		owner, err := base58.Decode(info.Owner, base58.BitcoinAlphabet)

		if err != nil {
			log.Error("Failed to decode address", "err", err, "address", info.Owner)
			return nil, nil
		}

		owner = owner[1 : len(owner)-4]
		owners[i] = common.BytesToAddress(owner)
		amounts[i] = new(big.Int).Mul(info.Amount, multiplier)
	}

	return
}

func parseSnapshot(reader io.Reader) (*snapshot, error) {
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()
	ret := &snapshot{}
	err := dec.Decode(ret)
	return ret, err
}

type snapshotItem struct {
	Owner  string   `json:"owner"`
	Amount *big.Int `json:"amount"`
	Atype  string   `json:"type"`
}

type snapshot struct {
	Txouts []snapshotItem `json:"snapshot_utxos"`
	Hash   string         `json:"snapshot_hash"`
}
