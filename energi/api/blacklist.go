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

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"energi.world/core/gen3/internal/ethapi"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

type BlacklistAPI struct {
	backend ethapi.Backend
}

func NewBlacklistAPI(b ethapi.Backend) *BlacklistAPI {
	return &BlacklistAPI{b}
}

const (
	blacklistCallGas uint64 = 500000
)

func (b *BlacklistAPI) registry(
	password string,
	dst common.Address,
) (session *energi_abi.IBlacklistRegistrySession, err error) {
	account := accounts.Account{Address: dst}
	wallet, err := b.backend.AccountManager().Find(account)
	if err != nil {
		return nil, err
	}

	contract, err := energi_abi.NewIBlacklistRegistry(
		energi_params.Energi_BlacklistRegistry, b.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IBlacklistRegistrySession{
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
					account, password, tx, b.backend.ChainConfig().ChainID)
			},
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

func (b *BlacklistAPI) BlacklistInfo() (res []BLInfo) {
	registry, err := energi_abi.NewIBlacklistRegistryCaller(
		energi_params.Energi_BlacklistRegistry, b.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return nil
	}

	call_opts := &bind.CallOpts{}
	addresses, err := registry.EnumerateAll(call_opts)
	if err != nil {
		log.Error("Failed", "err", err)
		return nil
	}

	res = make([]BLInfo, 0, len(addresses))

	for _, addr := range addresses {
		blocked, err := registry.IsBlacklisted(call_opts, addr)
		if err != nil {
			log.Warn("IsBlacklisted error", "addr", addr, "err", err)
			continue
		}

		proposals, err := registry.Proposals(call_opts, addr)
		if err != nil {
			log.Warn("Proposals error", "addr", addr, "err", err)
			continue
		}

		res = append(res, BLInfo{
			Target:  addr,
			Enforce: proposalInfo(b.backend, proposals.Enforce),
			Revoke:  proposalInfo(b.backend, proposals.Revoke),
			Drain:   proposalInfo(b.backend, proposals.Drain),
			Blocked: blocked,
		})
	}

	return
}

func (b *BlacklistAPI) BlacklistEnforce(
	address common.Address,
	fee *hexutil.Big,
	payer common.Address,
	password string,
) error {
	registry, err := b.registry(password, payer)
	if err != nil {
		return err
	}

	registry.TransactOpts.Value = fee.ToInt()
	tx, err := registry.Propose(address)

	log.Info("Note: please wait until proposal TX gets into a block!", "tx", tx.Hash())

	return err
}

func (b *BlacklistAPI) BlacklistRevoke(
	address common.Address,
	fee *hexutil.Big,
	payer common.Address,
	password string,
) error {
	registry, err := b.registry(password, payer)
	if err != nil {
		return err
	}

	is_blacklisted, err := registry.IsBlacklisted(address)
	if err != nil {
		return err
	}
	if !is_blacklisted {
		return errors.New("Not blocklisted")
	}

	registry.TransactOpts.Value = fee.ToInt()
	tx, err := registry.ProposeRevoke(address)

	log.Info("Note: please wait until proposal TX gets into a block!", "tx", tx.Hash())

	return err
}

func (b *BlacklistAPI) BlacklistDrain(
	address common.Address,
	fee *hexutil.Big,
	payer common.Address,
	password string,
) error {
	registry, err := b.registry(password, payer)
	if err != nil {
		return err
	}

	is_blacklisted, err := registry.IsBlacklisted(address)
	if err != nil {
		return err
	}
	if !is_blacklisted {
		return errors.New("Not blocklisted")
	}

	registry.TransactOpts.Value = fee.ToInt()
	tx, err := registry.ProposeDrain(address)

	log.Info("Note: please wait until proposal TX gets into a block!", "tx", tx.Hash())

	return err
}

func (b *BlacklistAPI) BlacklistCollect(
	target common.Address,
	payer common.Address,
	password string,
) error {
	registry, err := b.registry(password, payer)
	if err != nil {
		return err
	}

	tx, err := registry.Collect(target)

	log.Info("Note: please wait until collect TX gets into a block!", "tx", tx.Hash())

	return err
}
