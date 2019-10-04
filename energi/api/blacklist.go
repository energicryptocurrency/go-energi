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
	"github.com/ethereum/go-ethereum/log"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_common "energi.world/core/gen3/energi/common"
	energi_params "energi.world/core/gen3/energi/params"
)

type BlacklistAPI struct {
	backend   Backend
	infoCache *energi_common.CacheStorage
}

func NewBlacklistAPI(b Backend) *BlacklistAPI {
	return &BlacklistAPI{
		backend:   b,
		infoCache: energi_common.NewCacheStorage(),
	}
}

const (
	blacklistCallGas uint64 = 3000000
)

func (b *BlacklistAPI) registry(
	password *string,
	dst common.Address,
) (session *energi_abi.IBlacklistRegistrySession, err error) {
	contract, err := energi_abi.NewIBlacklistRegistry(
		energi_params.Energi_BlacklistRegistry, b.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IBlacklistRegistrySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		TransactOpts: bind.TransactOpts{
			From:     dst,
			Signer:   createSignerCallback(b.backend, password),
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
	data, err := b.infoCache.Get(b.backend, b.blacklistInfo)
	if err != nil || data == nil {
		log.Error("BlacklistInfo failed", "err", err)
		return
	}

	res = data.([]BLInfo)

	return
}

func (b *BlacklistAPI) blacklistInfo(blockhash common.Hash) (interface{}, error) {
	registry, err := energi_abi.NewIBlacklistRegistryCaller(
		energi_params.Energi_BlacklistRegistry, b.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed", "err", err)
		return nil, err
	}

	call_opts := &bind.CallOpts{
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
