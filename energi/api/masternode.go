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
	"net"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"

	energi_abi "energi.world/core/gen3/energi/abi"
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

func (m *MasternodeAPI) getAddress(
	dst common.Address,
) (account accounts.Account, wallet accounts.Wallet, err error) {
	account = accounts.Account{Address: dst}
	wallet, err = m.backend.AccountManager().Find(account)
	return
}

func (m *MasternodeAPI) token(
	password string,
	dst common.Address,
) (session *energi_abi.IMasternodeTokenSession, err error) {
	account, wallet, err := m.getAddress(dst)
	if err != nil {
		return nil, err
	}

	contract, err := energi_abi.NewIMasternodeToken(
		energi_params.Energi_MasternodeToken, m.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IMasternodeTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From: dst,
		},
		TransactOpts: bind.TransactOpts{
			From: dst,
			Signer: func(
				signer types.Signer,
				addr common.Address,
				tx *types.Transaction,
			) (*types.Transaction, error) {
				return wallet.SignTxWithPassphrase(
					account, password, tx, m.backend.ChainConfig().ChainID)
			},
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
			From: dst,
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
	password string,
) error {
	token, err := m.token(password, dst)
	if err != nil {
		return err
	}

	token.TransactOpts.Value = amount.ToInt()
	tx, err := token.DepositCollateral()

	if tx != nil {
		log.Info("Note: please wait until collateral TX gets into a block!", "tx", tx.Hash().Hex())
	}

	return err
}

func (m *MasternodeAPI) WithdrawCollateral(
	dst common.Address,
	amount *hexutil.Big,
	password string,
) error {
	token, err := m.token(password, dst)
	if err != nil {
		return err
	}

	tx, err := token.WithdrawCollateral(amount.ToInt())

	if tx != nil {
		log.Info("Note: please wait until collateral TX gets into a block!", "tx", tx.Hash().Hex())
	}

	return err
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

	call_opts := &bind.CallOpts{}
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

	call_opts := &bind.CallOpts{}
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

	call_opts := &bind.CallOpts{}
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
	ip := net.IPv4(
		byte(ipv4address>>24),
		byte(ipv4address>>16),
		byte(ipv4address>>8),
		byte(ipv4address),
	)

	pubkey_buf := make([]byte, 33)
	copy(pubkey_buf[:32], pubkey[0][:])
	copy(pubkey_buf[32:33], pubkey[1][:])
	pk, err := crypto.DecompressPubkey(pubkey_buf)
	if err != nil {
		log.Error("Failed to unmarshal pubkey")
		return ""
	}

	cfg := m.backend.ChainConfig()

	return enode.NewV4(pk, ip, int(cfg.ChainID.Int64()), int(cfg.ChainID.Int64())).String()
}

func (m *MasternodeAPI) registry(
	password string,
	dst common.Address,
) (session *energi_abi.IMasternodeRegistrySession, err error) {
	account, wallet, err := m.getAddress(dst)
	if err != nil {
		return nil, err
	}

	contract, err := energi_abi.NewIMasternodeRegistry(
		energi_params.Energi_MasternodeRegistry, m.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IMasternodeRegistrySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From: dst,
		},
		TransactOpts: bind.TransactOpts{
			From: dst,
			Signer: func(
				signer types.Signer,
				addr common.Address,
				tx *types.Transaction,
			) (*types.Transaction, error) {
				return wallet.SignTxWithPassphrase(
					account, password, tx, m.backend.ChainConfig().ChainID)
			},
			Value:    common.Big0,
			GasLimit: masternodeCallGas,
		},
	}
	return
}

func (m *MasternodeAPI) Announce(owner common.Address, enode_url, password string) error {
	registry, err := m.registry(password, owner)
	if err != nil {
		return err
	}

	var (
		ipv4address uint32
		pubkey      [2][32]byte
	)

	//---
	res, err := enode.ParseV4(enode_url)
	if err != nil {
		return err
	}

	//---
	ip := res.IP().To4()
	if ip == nil {
		return errors.New("Invalid IPv4")
	}

	if ip[0] == byte(127) || ip[0] == byte(10) ||
		(ip[0] == byte(192) && ip[1] == byte(168)) ||
		(ip[0] == byte(172) && (ip[1]&0xF0) == byte(16)) {
		return errors.New("Wrong enode IP")
	}

	cfg := m.backend.ChainConfig()

	if res.UDP() != int(cfg.ChainID.Int64()) || res.TCP() != int(cfg.ChainID.Int64()) {
		return errors.New("Wrong enode port")
	}

	ipv4address = uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])

	//---
	pk := crypto.CompressPubkey(res.Pubkey())
	if len(pk) != 33 {
		log.Error("Wrong public key length", "pklen", len(pk))
		return errors.New("Wrong public key")
	}

	copy(pubkey[0][:], pk[:32])
	copy(pubkey[1][:], pk[32:33])

	//---
	masternode := crypto.PubkeyToAddress(*res.Pubkey())

	//---
	tx, err := registry.Announce(masternode, ipv4address, pubkey)

	if tx != nil {
		log.Info("Note: please wait until TX gets into a block!", "tx", tx.Hash().Hex())
	}

	return err
}

func (m *MasternodeAPI) Denounce(owner common.Address, password string) (err error) {
	registry, err := m.registry(password, owner)
	if err != nil {
		return err
	}

	ownerinfo, err := registry.OwnerInfo(owner)
	if err != nil {
		log.Error("Not found", "owner", owner)
		return err
	}

	tx, err := registry.Denounce(ownerinfo.Masternode)

	if tx != nil {
		log.Info("Note: please wait until TX gets into a block!", "tx", tx.Hash().Hex())
	}

	return err
}
