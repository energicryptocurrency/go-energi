// Copyright 2020 The Energi Core Authors
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

package miner

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi "energi.world/core/gen3/energi/consensus"
	energi_params "energi.world/core/gen3/energi/params"
)

const maxAutoCollateralBlockAge = time.Duration(time.Minute)

func (w *worker) tryAutocollateral() {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if _, ok := w.engine.(*energi.Energi); !ok {
		// Energi consensus engine not running.
		log.Debug("energi consensus engine not running")
		return
	}

	block := w.eth.BlockChain().CurrentBlock()

	// MN-17 - 2
	// Check for block timeout.
	blockTime := time.Unix(int64(block.Time()), 0)
	timeNow := time.Now().UTC()
	if timeNow.After(blockTime.Add(maxAutoCollateralBlockAge)) {
		// if block older is older than maxAutoCollateralBlockAge, exit.
		log.Debug("block is older than maxAutoCollateralBlockAge")
		return
	}

	// Skip superblocks
	blockReward, err := w.superBlockCheck(block.Number())
	if err != nil {
		log.Debug(err.Error())
		return
	}

	log.Debug("Auto-Collateralize loop")

	for _, wallet := range w.eth.AccountManager().Wallets() {
		for _, account := range wallet.Accounts() {
			if wallet.IsUnlockedForStaking(account) {
				log.Debug("Auto-Collateralize checking", "account", account)

				amount, err := w.hasJustReceivedRewards(account.Address, block, blockReward)
				if err != nil {
					log.Debug(err.Error())
					// TODO: conditional disable for rapid live testing
					continue
				}

				if _, coins, err := w.doAutocollateral(account.Address, amount); err != nil {
					// Most likely, an invalid amount to deposit was found in the account.
					log.Debug("Auto-Collateralize failed", "err", err.Error())
				} else {
					log.Info("Auto-Collateralize successful", "coins deposited",
						coins.Uint64(), "account", account.Address.String())
				}
			}
		}
	}
}

func (w *worker) superBlockCheck(blockNo *big.Int) (*big.Int, error) {
	reward, err := w.getBlockReward(blockNo)
	if err != nil {
		return nil, err
	}

	// MN-17 - 4
	if reward.Cmp(common.Big0) == 0 {
		return nil, errors.New("Skipping super block for auto-collateral")
	}

	return reward, nil
}

func (w *worker) getBalanceAtBlock(block *types.Block, address common.Address) (*big.Int, error) {
	stateDb, err := w.eth.BlockChain().StateAt(block.Root())
	if err != nil {
		return nil, err
	}
	return stateDb.GetBalance(address), nil
}

// isMNPayouts checks if the provided block in relation to the previous block has
// a masternode block payout.
func (w *worker) isMNPayouts(currentblock *types.Block, mnOwner common.Address, blockReward *big.Int) (bool, *big.Int, error) {
	blockNo := currentblock.Number().Uint64()
	if blockNo <= 0 {
		return false, nil, errors.New("Invalid block cannot posses MN payouts")
	}

	balanceNow, err := w.getBalanceAtBlock(currentblock, mnOwner)
	if err != nil {
		return false, nil, err
	}

	prevBlock := w.eth.BlockChain().GetBlockByNumber(blockNo - 1)
	balancePrev, err := w.getBalanceAtBlock(prevBlock, mnOwner)
	if err != nil {
		return false, nil, err
	}

	diff := new(big.Int).Sub(balanceNow, balancePrev)

	// Should be: 0 < diff <= mnPayout
	status := diff.Cmp(blockReward) <= 0 && diff.Cmp(common.Big0) > 0
	return status, balanceNow, nil
}

func (w *worker) hasJustReceivedRewards(account common.Address, block *types.Block, blockReward *big.Int) (*big.Int, error) {
	// MN-17 - 5
	// (a) Confirm no MN payouts in the current block.
	isCurrentMNPayout, balanceNow, err := w.isMNPayouts(block, account, blockReward)
	if err != nil {
		return nil, err
	}

	if isCurrentMNPayout {
		return balanceNow, errors.New("Current block masternode payout is active")
	}

	// MN-17 - 5
	// (b) Confirm atleast 1 masternode payout in the previous block.
	prevBlock := w.eth.BlockChain().GetBlockByNumber(block.Number().Uint64() - 1)
	isPrevPayout, _, err := w.isMNPayouts(prevBlock, account, blockReward)
	if err != nil {
		return balanceNow, err
	}

	if !isPrevPayout {
		return balanceNow, errors.New("Expected at least one payout from a previous block")
	}

	return balanceNow, nil
}

// canAutocollateralize returns the maximum amount that can be deposited as the
// collateral if the maximum collateral amount is not yet reached.
func (w *worker) canAutocollateralize(
	account common.Address,
	amount *big.Int,
	api *energi_abi.IMasternodeTokenSession,
) (*big.Int, error) {
	minLimit, maxLimit, err := w.collateralLimits()
	if err != nil {
		return nil, err
	}

	// MN-17 - 5
	// (c) Ensures that available balance is at least one minimal collateral.
	if amount.Cmp(minLimit) < 0 {
		return nil, errors.New("Amount found is less than the minimum required")
	}

	tokenBalance, err := api.BalanceOf(account)
	if err != nil {
		return nil, err
	}

	// MN-17 - 5
	// (d) Ensure that the current collateral is below the maximum allowed.
	if tokenBalance.Cmp(maxLimit) >= 0 {
		return nil, errors.New("Maximum collateral supported already achieved")
	}

	modAmount := new(big.Int).Mod(amount, minLimit)
	amountToDeposit := new(big.Int).Sub(amount, modAmount)

	totalAmount := new(big.Int).Add(tokenBalance, amountToDeposit)
	if totalAmount.Cmp(maxLimit) == 1 {
		// Gets the maximum amount to deposit since all the available amount
		// could breach the max collateral limit if deposited in full.
		amountToDeposit = new(big.Int).Sub(maxLimit, tokenBalance)
	}

	return amountToDeposit, nil
}

func (w *worker) doAutocollateral(account common.Address, amount *big.Int) (common.Hash, *big.Int, error) {
	tokenAPI, err := w.tokenRegistry(account)
	if err != nil {
		return common.Hash{}, nil, err
	}

	// Returns the maximum amount that can be deposited if the collateral max
	// amount hasn't been reached.
	newAmount, err := w.canAutocollateralize(account, amount, tokenAPI)
	if err != nil {
		return common.Hash{}, nil, err
	}

	// MN-17 - 4
	// (e) Perform MNReg.depositCollataral
	tokenAPI.TransactOpts.Value = newAmount
	tx, err := tokenAPI.DepositCollateral()
	if tx == nil || err != nil {
		return common.Hash{}, nil, err
	}

	coinsDeposited := new(big.Int).Div(newAmount, big.NewInt(params.Ether))

	return tx.Hash(), coinsDeposited, nil
}

func (w *worker) getBlockReward(blockNumber *big.Int) (*big.Int, error) {
	contract, err := energi_abi.NewIBlockReward(
		energi_params.Energi_MasternodeRegistry, w.apiBackend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{
		GasLimit: energi_params.UnlimitedGas,
	}

	resp, err := contract.GetReward(callOpts, blockNumber)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *worker) tokenRegistry(dst common.Address) (*energi_abi.IMasternodeTokenSession, error) {
	contract, err := energi_abi.NewIMasternodeToken(
		energi_params.Energi_MasternodeToken, w.apiBackend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}
	session := &energi_abi.IMasternodeTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		TransactOpts: bind.TransactOpts{
			From:     dst,
			Signer:   w.createStakeTxSignerCallback(),
			Value:    common.Big0,
			GasLimit: energi_params.MasternodeCallGas,
		},
	}
	return session, nil
}

func (w *worker) collateralLimits() (minCollateral, maxCollateral *big.Int, err error) {
	registry, err := energi_abi.NewIMasternodeRegistryV2(
		energi_params.Energi_MasternodeRegistry, w.apiBackend.(bind.ContractBackend))
	if err != nil {
		return nil, nil, err
	}

	callOpts := &bind.CallOpts{
		GasLimit: energi_params.UnlimitedGas,
	}

	limits, err := registry.CollateralLimits(callOpts)
	if err != nil {
		return nil, nil, err
	}

	return limits.Min, limits.Max, nil
}

// CreateStakeTxSignerCallback uses unlocked accounts to sign transactions without
// a password.
func (w *worker) createStakeTxSignerCallback() bind.SignerFn {
	return func(
		signer types.Signer,
		addr common.Address,
		tx *types.Transaction,
	) (*types.Transaction, error) {
		account := accounts.Account{Address: addr}
		wallet, err := w.eth.AccountManager().Find(account)
		if err != nil {
			return nil, err
		}

		if !wallet.IsUnlockedForStaking(account) {
			return nil, errors.New("Auto-collateral requires an unlocked account")
		}

		chainID := w.eth.BlockChain().Config().ChainID
		return wallet.SignTx(account, tx, chainID)
	}
}
