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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_common "energi.world/core/gen3/energi/common"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	mntokenCallGas    uint64 = 300000
	masternodeCallGas uint64 = 500000
)

type MasternodeAPI struct {
	backend Backend
}

func NewMasternodeAPI(b Backend) *MasternodeAPI {
	return &MasternodeAPI{b}
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
}) {
	token, err := energi_abi.NewIMasternodeTokenCaller(
		energi_params.Energi_MasternodeToken, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return ret
	}

	res, err := token.BalanceInfo(
		&bind.CallOpts{
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		dst,
	)
	if err != nil {
		log.Error("Failed", "err", err)
		return ret
	}

	ret.Balance = (*hexutil.Big)(res.Balance)
	ret.LastBlock = (*hexutil.Big)(res.LastBlock)
	return ret
}

func (m *MasternodeAPI) DepositCollateral(
	dst common.Address,
	amount *hexutil.Big,
	password *string,
) (txhash common.Hash, err error) {
	token, err := m.token(password, dst)
	if err != nil {
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
	token, err := m.token(password, dst)
	if err != nil {
		return
	}

	tx, err := token.WithdrawCollateral(amount.ToInt())

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the collateral TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

type MNInfo struct {
	Masternode     common.Address
	Owner          common.Address
	Enode          string
	Collateral     *hexutil.Big
	AnnouncedBlock uint64
}

func (m *MasternodeAPI) ListMasternodes() (res []MNInfo) {
	registry, err := energi_abi.NewIMasternodeRegistryCaller(
		energi_params.Energi_MasternodeRegistry, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return nil
	}

	call_opts := &bind.CallOpts{
		GasLimit: energi_params.UnlimitedGas,
	}
	masternodes, err := registry.Enumerate(call_opts)
	if err != nil {
		log.Error("Failed", "err", err)
		return nil
	}

	res = make([]MNInfo, 0, len(masternodes))

	for _, mn := range masternodes {
		mninfo, err := registry.Info(call_opts, mn)
		if err != nil {
			log.Warn("Info error", "mn", mn, "err", err)
			continue
		}

		res = append(res, MNInfo{
			Masternode:     mn,
			Owner:          mninfo.Owner,
			Enode:          m.enode(mninfo.Ipv4address, mninfo.Enode),
			Collateral:     (*hexutil.Big)(mninfo.Collateral),
			AnnouncedBlock: mninfo.AnnouncedBlock.Uint64(),
		})
	}

	return
}

func (m *MasternodeAPI) MasternodeInfo(owner_or_mn common.Address) (res MNInfo) {
	registry, err := energi_abi.NewIMasternodeRegistryCaller(
		energi_params.Energi_MasternodeRegistry, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	call_opts := &bind.CallOpts{
		GasLimit: energi_params.UnlimitedGas,
	}
	mninfo, err := registry.Info(call_opts, owner_or_mn)

	if err == nil {
		res.Masternode = owner_or_mn
		res.Owner = mninfo.Owner
		res.Enode = m.enode(mninfo.Ipv4address, mninfo.Enode)
		res.Collateral = (*hexutil.Big)(mninfo.Collateral)
		res.AnnouncedBlock = mninfo.AnnouncedBlock.Uint64()
		return
	}

	ownerinfo, err := registry.OwnerInfo(call_opts, owner_or_mn)
	if err != nil {
		log.Error("Not found", "mn", owner_or_mn)
		return
	}

	res.Masternode = ownerinfo.Masternode
	res.Owner = owner_or_mn
	res.Enode = m.enode(ownerinfo.Ipv4address, ownerinfo.Enode)
	res.Collateral = (*hexutil.Big)(ownerinfo.Collateral)
	res.AnnouncedBlock = ownerinfo.AnnouncedBlock.Uint64()
	return
}

func (m *MasternodeAPI) Stats() (res struct {
	Active           uint64
	Total            uint64
	ActiveCollateral *hexutil.Big
	TotalCollateral  *hexutil.Big
	MaxOfAllTimes    *hexutil.Big
}) {
	registry, err := energi_abi.NewIMasternodeRegistryCaller(
		energi_params.Energi_MasternodeRegistry, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	call_opts := &bind.CallOpts{
		GasLimit: energi_params.UnlimitedGas,
	}
	count, err := registry.Count(call_opts)
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	res.Active = count.Active.Uint64()
	res.Total = count.Total.Uint64()
	res.ActiveCollateral = (*hexutil.Big)(count.ActiveCollateral)
	res.TotalCollateral = (*hexutil.Big)(count.TotalCollateral)
	res.MaxOfAllTimes = (*hexutil.Big)(count.MaxOfAllTimes)

	return
}

func (m *MasternodeAPI) enode(ipv4address uint32, pubkey [2][32]byte) string {
	cfg := m.backend.ChainConfig()
	res := energi_common.MastenodeEnode(ipv4address, pubkey, cfg)

	if res == nil {
		return ""
	}

	return res.String()
}

func (m *MasternodeAPI) registry(
	password *string,
	dst common.Address,
) (session *energi_abi.IMasternodeRegistrySession, err error) {
	contract, err := energi_abi.NewIMasternodeRegistry(
		energi_params.Energi_MasternodeRegistry, m.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IMasternodeRegistrySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		TransactOpts: bind.TransactOpts{
			From:     dst,
			Signer:   createSignerCallback(m.backend, password),
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
	registry, err := m.registry(password, owner)
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
	registry, err := m.registry(password, owner)
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
