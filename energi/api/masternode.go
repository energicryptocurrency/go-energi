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
	"fmt"
	"math/big"

	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/common/hexutil"
	"github.com/energicryptocurrency/energi/crypto"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/p2p/enode"

	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_common "github.com/energicryptocurrency/energi/energi/common"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
)

const (
	mntokenCallGas    uint64 = energi_params.MasternodeCallGas
	masternodeCallGas uint64 = energi_params.MasternodeCallGas
)

type MasternodeAPI struct {
	backend    Backend
	nodesCache *energi_common.CacheStorage
	statsCache *energi_common.CacheStorage
}

func NewMasternodeAPI(b Backend) *MasternodeAPI {
	r := &MasternodeAPI{
		backend:    b,
		nodesCache: energi_common.NewCacheStorage(),
		statsCache: energi_common.NewCacheStorage(),
	}
	b.OnSyncedHeadUpdates(func() {
		r.ListMasternodes()
		r.Stats()
	})
	return r
}

type MasternodeStats struct {
	Active           uint64
	Total            uint64
	ActiveCollateral *hexutil.Big
	TotalCollateral  *hexutil.Big
	MaxOfAllTimes    *hexutil.Big
}

func (m *MasternodeAPI) token(
	password *string,
	dst common.Address,
) (session *energi_abi.IMasternodeTokenSession, err error) {
	contract, err := energi_abi.NewIMasternodeToken(
		energi_params.Energi_MasternodeToken, m.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IMasternodeTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending:  true,
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		TransactOpts: bind.TransactOpts{
			From:     dst,
			Signer:   createSignerCallback(m.backend, password),
			Value:    common.Big0,
			GasLimit: mntokenCallGas,
		},
	}
	return
}

func (m *MasternodeAPI) CollateralBalance(
	dst common.Address,
) (ret struct {
	Balance   *hexutil.Big
	LastBlock *hexutil.Big
}, err error) {
	token, err := energi_abi.NewIMasternodeTokenCaller(
		energi_params.Energi_MasternodeToken, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return ret, err
	}

	res, err := token.BalanceInfo(
		&bind.CallOpts{
			Pending:  true,
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		dst,
	)
	if err != nil {
		log.Error("Failed", "err", err)
		return ret, err
	}

	ret.Balance = (*hexutil.Big)(res.Balance)
	ret.LastBlock = (*hexutil.Big)(res.LastBlock)
	return ret, nil
}

func (m *MasternodeAPI) DepositCollateral(
	dst common.Address,
	amount *hexutil.Big,
	password *string,
) (txhash common.Hash, err error) {
	registry, err := masternodeRegistry(password, dst, m.backend)
	if err != nil {
		return
	}

	limits, err := registry.CollateralLimits()
	if err != nil {
		return
	}

	if err = m.validateAmount("Deposit", amount.ToInt(), limits.Min); err != nil {
		return
	}

	token, err := m.token(password, dst)
	if err != nil {
		return
	}

	balance, err := token.BalanceOf(dst)
	if err != nil {
		// Possible new masternode detected.
		log.Warn("Fetching masternode collateral failed: %v", err)
	}

	newTotalAmount := new(big.Int).Add(balance, amount.ToInt())

	// Expected total amount should not be more than the maximum collateral value allowed.
	if newTotalAmount.Cmp(limits.Max) > 0 {
		err = fmt.Errorf("Total expected deposits should not exceed max allowed deposit")
		return
	}

	token.TransactOpts.Value = amount.ToInt()
	tx, err := token.DepositCollateral()

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the collateral TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (m *MasternodeAPI) WithdrawCollateral(
	dst common.Address,
	amount *hexutil.Big,
	password *string,
) (txhash common.Hash, err error) {
	registry, err := masternodeRegistry(password, dst, m.backend)
	if err != nil {
		return
	}

	limits, err := registry.CollateralLimits()
	if err != nil {
		return
	}

	if err = m.validateAmount("Withdrawal", amount.ToInt(), limits.Min); err != nil {
		return
	}

	token, err := m.token(password, dst)
	if err != nil {
		return
	}

	balance, err := token.BalanceOf(dst)
	if err != nil {
		err = fmt.Errorf("Fetching masternode collateral failed: %v", err)
		return
	}

	// Amount to withdraw should be less than or equal to the collateral amount.
	if amount.ToInt().Cmp(balance) > 0 {
		err = fmt.Errorf("Withdrawal amount is greater than the collateral balance amount")
		return
	}

	tx, err := token.WithdrawCollateral(amount.ToInt())

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the collateral TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (m *MasternodeAPI) validateAmount(validateType string, amount, minColl *big.Int) error {
	// Amount to should be greater than zero.
	if amount.Cmp(common.Big0) < 1 {
		return fmt.Errorf("%v amount should be greater than zero", validateType)
	}

	// Amount should be a multiple of the minimum collateral amount allowed.
	if new(big.Int).Mod(amount, minColl).Cmp(common.Big0) != 0 {
		return fmt.Errorf("%v amount should be a multiple of the minimum collateral amount", validateType)
	}

	return nil
}

type MNInfo struct {
	Masternode     common.Address
	Owner          common.Address
	Enode          string
	Collateral     *hexutil.Big
	AnnouncedBlock uint64
	IsActive       bool
	IsAlive        bool
	SWFeatures     *hexutil.Big
	SWVersion      string
}

func (m *MasternodeAPI) ListMasternodes() (res []MNInfo, err error) {
	data, err := m.nodesCache.Get(m.backend, m.listMasternodes)
	if err != nil || data == nil {
		log.Error("ListMasternodes failed", "err", err)
		return
	}

	res = data.([]MNInfo)

	return
}

func (m *MasternodeAPI) listMasternodes(num *big.Int) (interface{}, error) {
	registry, err := energi_abi.NewIMasternodeRegistryV2Caller(
		energi_params.Energi_MasternodeRegistry, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return nil, err
	}

	call_opts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}
	prev_call_opts := &bind.CallOpts{
		BlockNumber: new(big.Int).Sub(num, common.Big3),
		GasLimit:    energi_params.UnlimitedGas,
	}
	if prev_call_opts.BlockNumber.Cmp(common.Big0) < 0 {
		prev_call_opts.BlockNumber = new(big.Int).Set(common.Big0)
	}

	masternodes, err := registry.Enumerate(call_opts)
	if err != nil {
		log.Error("Failed", "err", err)
		return nil, err
	}

	res := make([]MNInfo, 0, len(masternodes))
	for _, mn := range masternodes {
		mninfo, err := registry.Info(call_opts, mn)
		if err != nil {
			log.Debug("Info error", "mn", mn, "err", err)
			continue
		}

		isActive, err := registry.IsActive(call_opts, mn)
		if err != nil {
			log.Debug("IsActive error", "mn", mn, "err", err)
			continue
		}

		canHeartbeat, err := registry.CanHeartbeat(call_opts, mn)
		if err != nil {
			log.Debug("CanHeartbeat error", "mn", mn, "err", err)
			continue
		}

		prevCanHeartbeat, err := registry.CanHeartbeat(prev_call_opts, mn)
		if err != nil {
			// missing trie node may appear on non-full node
			log.Debug("Prev CanHeartbeat error", "mn", mn, "err", err)
			// assume the same as atm
			prevCanHeartbeat = canHeartbeat
		}

		res = append(res, MNInfo{
			Masternode:     mn,
			Owner:          mninfo.Owner,
			Enode:          m.enode(mninfo.Ipv4address, mninfo.Enode),
			Collateral:     (*hexutil.Big)(mninfo.Collateral),
			AnnouncedBlock: mninfo.AnnouncedBlock.Uint64(),
			IsActive:       isActive,
			IsAlive:        isActive && (!canHeartbeat || !prevCanHeartbeat),
			SWFeatures:     (*hexutil.Big)(mninfo.SwFeatures),
			SWVersion:      energi_common.SWVersionIntToString(mninfo.SwFeatures),
		})
	}

	return res, err
}

func (m *MasternodeAPI) MasternodeInfo(owner_or_mn common.Address) (res MNInfo, err error) {
	Mns, err := m.ListMasternodes()
	if err != nil {
		log.Error("Failed at m.ListMasternodes", "err", err)
		return
	}

	for _, node := range Mns {
		if node.Masternode == owner_or_mn || node.Owner == owner_or_mn {
			res = node
			break
		}
	}

	return
}

func (m *MasternodeAPI) Stats() (res *MasternodeStats, err error) {
	data, err := m.statsCache.Get(m.backend, m.stats)

	if err != nil || data == nil {
		log.Error("Stats failed", "err", err)
		return
	}

	res = data.(*MasternodeStats)
	return
}

func (m *MasternodeAPI) stats(num *big.Int) (interface{}, error) {
	registry, err := energi_abi.NewIMasternodeRegistryV2Caller(
		energi_params.Energi_MasternodeRegistry, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return nil, err
	}

	call_opts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}
	count, err := registry.Count(call_opts)
	if err != nil {
		log.Error("Failed", "err", err)
		return nil, err
	}

	res := &MasternodeStats{}
	res.Active = count.Active.Uint64()
	res.Total = count.Total.Uint64()
	res.ActiveCollateral = (*hexutil.Big)(count.ActiveCollateral)
	res.TotalCollateral = (*hexutil.Big)(count.TotalCollateral)
	res.MaxOfAllTimes = (*hexutil.Big)(count.MaxOfAllTimes)

	return res, nil
}

func (m *MasternodeAPI) enode(ipv4address uint32, pubkey [2][32]byte) string {
	cfg := m.backend.ChainConfig()
	res := energi_common.MastenodeEnode(ipv4address, pubkey, cfg)

	if res == nil {
		return ""
	}

	return res.String()
}

func masternodeRegistry(
	password *string,
	dst common.Address,
	backend Backend,
) (session *energi_abi.IMasternodeRegistryV2Session, err error) {
	contract, err := energi_abi.NewIMasternodeRegistryV2(
		energi_params.Energi_MasternodeRegistry, backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IMasternodeRegistryV2Session{
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
			GasLimit: masternodeCallGas,
		},
	}
	return
}

func (m *MasternodeAPI) Announce(
	owner common.Address,
	enode_url string,
	password *string,
) (txhash common.Hash, err error) {
	registry, err := masternodeRegistry(password, owner, m.backend)
	if err != nil {
		return
	}

	var (
		ipv4address uint32
		pubkey      [2][32]byte
	)

	//---
	res, err := enode.ParseV4(enode_url)
	if err != nil {
		return
	}

	//---
	ip := res.IP().To4()
	if ip == nil {
		err = errors.New("Invalid IPv4")
		return
	}

	if ip[0] == byte(127) || ip[0] == byte(10) ||
		(ip[0] == byte(192) && ip[1] == byte(168)) ||
		(ip[0] == byte(172) && (ip[1]&0xF0) == byte(16)) {
		err = errors.New("Wrong enode IP")
		return
	}

	cfg := m.backend.ChainConfig()

	if res.UDP() != int(cfg.ChainID.Int64()) || res.TCP() != int(cfg.ChainID.Int64()) {
		err = errors.New("Wrong enode port")
		return
	}

	ipv4address = uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])

	//---
	pk := crypto.CompressPubkey(res.Pubkey())
	if len(pk) != 33 {
		log.Error("Wrong public key length", "pklen", len(pk))
		err = errors.New("Wrong public key")
		return
	}

	copy(pubkey[0][:], pk[:32])
	copy(pubkey[1][:], pk[32:33])

	//---
	masternode := crypto.PubkeyToAddress(*res.Pubkey())

	//---
	tx, err := registry.Announce(masternode, ipv4address, pubkey)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (m *MasternodeAPI) Denounce(
	owner common.Address,
	password *string,
) (txhash common.Hash, err error) {
	registry, err := masternodeRegistry(password, owner, m.backend)
	if err != nil {
		return
	}

	ownerinfo, err := registry.OwnerInfo(owner)
	if err != nil {
		log.Error("Not found", "owner", owner)
		return
	}

	tx, err := registry.Denounce(ownerinfo.Masternode)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the TX gets into a block!", "tx", txhash.Hex())
	}

	return
}
