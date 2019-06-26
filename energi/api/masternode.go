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
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	masternodeCallGas uint64 = 300000
)

type MasternodeAPI struct {
	backend ethapi.Backend
}

func NewMasternodeAPI(b ethapi.Backend) *MasternodeAPI {
	return &MasternodeAPI{b}
}

func (m *MasternodeAPI) getAddress(
	dst common.Address,
) (account accounts.Account, wallet accounts.Wallet, err error) {
	account = accounts.Account{Address: dst}
	wallet, err = m.backend.AccountManager().Find(accounts.Account{Address: dst})
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
			GasLimit: masternodeCallGas,
		},
	}
	return
}

func (m *MasternodeAPI) CollateralBalance(
	dst common.Address,
) struct {
	Balance   *hexutil.Big
	LastBlock *hexutil.Big
} {
	var ret struct {
		Balance   *hexutil.Big
		LastBlock *hexutil.Big
	}

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
	_, err = token.DepositCollateral()

	log.Info("Note: please wait until collateral TX gets into a block!")

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

	_, err = token.WithdrawCollateral(amount.ToInt())

	log.Info("Note: please wait until collateral TX gets into a block!")

	return err
}
