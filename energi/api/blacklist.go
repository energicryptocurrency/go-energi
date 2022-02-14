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

package api

import (
	"errors"
	"math/big"

	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/common/hexutil"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_common "github.com/energicryptocurrency/energi/energi/common"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/log"
)

type BlacklistAPI struct {
	backend   Backend
	infoCache *energi_common.CacheStorage
	compCache *energi_common.CacheStorage
}

func NewBlacklistAPI(b Backend) *BlacklistAPI {
	r := &BlacklistAPI{
		backend:   b,
		infoCache: energi_common.NewCacheStorage(),
		compCache: energi_common.NewCacheStorage(),
	}
	b.OnSyncedHeadUpdates(func() {
		r.BlacklistInfo()
		r.CompensationInfo()
	})
	return r
}

const (
	blacklistCallGas uint64 = 3000000
)

func (b *BlacklistAPI) registry(
	password *string,
	dst common.Address,
) (session *energi_abi.IBlacklistRegistrySession, err error) {
	return blacklistRegistry(b.backend, password, dst)
}

func blacklistRegistry(
	backend Backend,
	password *string,
	dst common.Address,
) (session *energi_abi.IBlacklistRegistrySession, err error) {
	contract, err := energi_abi.NewIBlacklistRegistry(
		energi_params.Energi_BlacklistRegistry, backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IBlacklistRegistrySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending:  true,
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		TransactOpts: bind.TransactOpts{
			From:     dst,
			Signer:   createSignerCallback(backend, password),
			Value:    common.Big0,
			GasLimit: blacklistCallGas,
		},
	}
	return
}

type BLInfo struct {
	Target  common.Address
	Enforce *ProposalInfo
	Revoke  *ProposalInfo
	Drain   *ProposalInfo
	Blocked bool
}

func (b *BlacklistAPI) BlacklistInfo() (res []BLInfo, err error) {
	data, err := b.infoCache.Get(b.backend, b.blacklistInfo)
	if err != nil || data == nil {
		log.Error("BlacklistInfo failed", "err", err)
		return
	}

	res = data.([]BLInfo)

	return
}

func (b *BlacklistAPI) blacklistInfo(num *big.Int) (interface{}, error) {
	registry, err := energi_abi.NewIBlacklistRegistryCaller(
		energi_params.Energi_BlacklistRegistry, b.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return nil, err
	}

	call_opts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}
	addresses, err := registry.EnumerateAll(call_opts)
	if err != nil {
		log.Error("Failed", "err", err)
		return nil, err
	}

	res := make([]BLInfo, 0, len(addresses))

	for _, addr := range addresses {
		blocked, err := registry.IsBlacklisted(call_opts, addr)
		if err != nil {
			log.Debug("IsBlacklisted error", "addr", addr, "err", err)
			continue
		}

		proposals, err := registry.Proposals(call_opts, addr)
		if err != nil {
			log.Debug("Proposals error", "addr", addr, "err", err)
			continue
		}

		enforceInfo, err := proposalInfo(b.backend, proposals.Enforce)
		if err != nil {
			log.Debug("Enforce info error", "addr", addr, "err", err)
		}

		revokeInfo, err := proposalInfo(b.backend, proposals.Revoke)
		if err != nil {
			log.Debug("Revoke info error", "addr", addr, "err", err)
		}

		drainInfo, err := proposalInfo(b.backend, proposals.Drain)
		if err != nil {
			log.Debug("Drain info error", "addr", addr, "err", err)
		}

		res = append(res, BLInfo{
			Target:  addr,
			Enforce: enforceInfo,
			Revoke:  revokeInfo,
			Drain:   drainInfo,
			Blocked: blocked,
		})
	}

	return res, nil
}

func (b *BlacklistAPI) BlacklistEnforce(
	address common.Address,
	fee *hexutil.Big,
	payer common.Address,
	password *string,
) (txhash common.Hash, err error) {
	registry, err := b.registry(password, payer)
	if err != nil {
		return
	}

	registry.TransactOpts.Value = fee.ToInt()
	tx, err := registry.Propose(address)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (b *BlacklistAPI) BlacklistRevoke(
	address common.Address,
	fee *hexutil.Big,
	payer common.Address,
	password *string,
) (txhash common.Hash, err error) {
	registry, err := b.registry(password, payer)
	if err != nil {
		return
	}

	is_blacklisted, err := registry.IsBlacklisted(address)
	if err != nil {
		return
	}
	if !is_blacklisted {
		err = errors.New("Not blocklisted")
		return
	}

	registry.TransactOpts.Value = fee.ToInt()
	tx, err := registry.ProposeRevoke(address)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (b *BlacklistAPI) BlacklistDrain(
	address common.Address,
	fee *hexutil.Big,
	payer common.Address,
	password *string,
) (txhash common.Hash, err error) {
	registry, err := b.registry(password, payer)
	if err != nil {
		return
	}

	is_blacklisted, err := registry.IsBlacklisted(address)
	if err != nil {
		return
	}
	if !is_blacklisted {
		err = errors.New("Not blocklisted")
		return
	}

	registry.TransactOpts.Value = fee.ToInt()
	tx, err := registry.ProposeDrain(address)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (b *BlacklistAPI) BlacklistCollect(
	target common.Address,
	payer common.Address,
	password *string,
) (txhash common.Hash, err error) {
	registry, err := b.registry(password, payer)
	if err != nil {
		return
	}

	tx, err := registry.Collect(target)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the collect TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (b *BlacklistAPI) CompensationInfo() (*BudgetInfo, error) {
	data, err := b.compCache.Get(b.backend, b.compensationInfo)
	if err != nil || data == nil {
		log.Error("CompensationInfo failed", "err", err)
		return nil, err
	}

	return data.(*BudgetInfo), nil
}

func (b *BlacklistAPI) compensationInfo(num *big.Int) (interface{}, error) {
	comp_fund, err := b.compensationFundAddress()
	if err != nil {
		return nil, err
	}

	return treasuryInfo(comp_fund, b.backend)
}

func (b *BlacklistAPI) CompensationPropose(
	amount *hexutil.Big,
	ref_uuid string,
	period uint64,
	fee *hexutil.Big,
	payer common.Address,
	password *string,
) (txhash common.Hash, err error) {
	comp_fund, err := b.compensationFundAddress()
	if err != nil {
		return common.Hash{}, err
	}

	return treasuryPropose(
		b.backend, comp_fund,
		amount, ref_uuid,
		period, fee,
		payer, password,
	)
}

func (b *BlacklistAPI) compensationFundAddress() (common.Address, error) {
	contract, err := energi_abi.NewIBlacklistRegistry(
		energi_params.Energi_BlacklistRegistry, b.backend.(bind.ContractBackend))
	if err != nil {
		return common.Address{}, err
	}

	call_opts := &bind.CallOpts{
		Pending: true,
	}

	return contract.CompensationFund(call_opts)
}

// NOTE: use of MigrationAPI is to simplify things and to minimize refactoring
func (m *MigrationAPI) CompensationProcess(
	payer common.Address,
	password *string,
) error {
	reward := false

	// Process drainable migrations
	//---
	registry, err := blacklistRegistry(m.backend, password, payer)
	if err != nil {
		log.Error("Failed BLRegistry", "err", err)
		return err
	}

	addresses, err := registry.EnumerateDrainable()
	if err != nil {
		log.Error("Failed EnumerateDrainable", "err", err)
		return err
	}

	found, _ := m.SearchRawGen2Coins(addresses, false)

	for _, fa := range found {
		tx, err := registry.DrainMigration(new(big.Int).SetUint64(fa.ItemID), fa.RawOwner)

		if err != nil {
			log.Error("Failed DrainMigration", "err", err)
			return err
		} else {
			log.Info("Sent drain transaction", "tx", tx.Hash().Hex(), "coins", fa.Owner)
		}

		reward = true
	}

	//---
	comp_fund_addr, err := registry.CompensationFund()
	if err != nil {
		return err
	}

	comp_fund, err := treasury(m.backend, comp_fund_addr, password, payer)
	if err != nil {
		return err
	}

	proposals, err := comp_fund.ListProposals()
	if err != nil {
		log.Error("Failed ListProposals", "err", err)
		return err
	}

	for _, p := range proposals {
		contract, err := energi_abi.NewIProposal(p, m.backend.(bind.ContractBackend))
		if err != nil {
			return err
		}

		if yes, _ := contract.IsAccepted(&comp_fund.CallOpts); yes {
			reward = true
			break
		}
	}

	if !reward {
		return nil
	}

	reward_comp_fund, err := energi_abi.NewIBlockReward(
		comp_fund_addr, m.backend.(bind.ContractBackend))
	if err != nil {
		return err
	}

	tx, err := reward_comp_fund.Reward(&comp_fund.TransactOpts)
	log.Info("Sent distribute transaction", "tx", tx.Hash().Hex())

	return err
}
