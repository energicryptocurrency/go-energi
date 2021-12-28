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

	"github.com/energicryptocurrency/energi/accounts"
	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi "github.com/energicryptocurrency/energi/energi/consensus"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/params"
)

const maxAutoCollateralBlockAge = time.Duration(time.Minute)

const (
	acDisabled uint64 = 0
	acRapid    uint64 = 2
)

func (w *worker) tryAutoCompound() {
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

	// Get rewards
	mnReward, err := w.getBlockReward(energi_params.Energi_MasternodeRegistry, block.Number())
	if err != nil {
		log.Error(err.Error())
		return
	}

	// Skip superblocks
	// MN-17 - 4
	if mnReward.Cmp(common.Big0) == 0 {
		log.Debug("Skipping super block for auto-collateral")
		return
	}

	log.Debug("Auto-Collateralize loop")

	for _, wallet := range w.eth.AccountManager().Wallets() {
		for _, account := range wallet.Accounts() {
			if wallet.IsUnlockedForStaking(account) {
				log.Debug("Auto-Collateralize checking", "account", account)

				amount, err := w.hasJustReceivedRewards(account.Address, block, mnReward)
				if err != nil {
					log.Debug(err.Error())
					if amount == nil || w.autocollateral != acRapid {
						continue
					}
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

	// MN-17 - 5
	// (a).ii Ensure the change was not due to stake reward
	if currentblock.Coinbase() == mnOwner {
		return false, balanceNow, errors.New("Stake reward")
	}

	prevBlock := w.eth.BlockChain().GetBlockByNumber(blockNo - 1)
	balancePrev, err := w.getBalanceAtBlock(prevBlock, mnOwner)
	if err != nil {
		return false, balanceNow, err
	}

	diff := new(big.Int).Sub(balanceNow, balancePrev)

	// Should be: 0 < diff <= mnPayout
	status := diff.Cmp(blockReward) <= 0 && diff.Cmp(common.Big0) > 0
	return status, balanceNow, nil
}

func (w *worker) hasJustReceivedRewards(account common.Address, block *types.Block, mnReward *big.Int) (*big.Int, error) {
	// MN-17 - 5
	// (a).i Confirm no MN payouts in the current block.
	isCurrentMNPayout, balanceNow, err := w.isMNPayouts(block, account, mnReward)
	if err != nil {
		return nil, err
	}

	if isCurrentMNPayout {
		return balanceNow, errors.New("Current block masternode payout is active")
	}

	// MN-17 - 5
	// (a).ii Confirm atleast 1 masternode payout in the previous block.
	prevBlock := w.eth.BlockChain().GetBlockByNumber(block.Number().Uint64() - 1)
	isPrevPayout, _, err := w.isMNPayouts(prevBlock, account, mnReward)
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
	// (b) Ensures that available balance is at least one minimal collateral.
	if amount.Cmp(minLimit) < 0 {
		return nil, errors.New("Amount found is less than the minimum required")
	}

	tokenBalance, err := api.BalanceOf(account)
	if err != nil {
		return nil, err
	}

	// MN-17 - 5
	// (c) Ensure that the current collateral is below the maximum allowed and more than zero.
	if tokenBalance.Cmp(common.Big0) <= 0 {
		return nil, errors.New("No collateral exists")
	} else if tokenBalance.Cmp(maxLimit) >= 0 {
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

	// MN-17 - 5
	// (d) Perform MNReg.depositCollataral
	tokenAPI.TransactOpts.Value = newAmount
	tx, err := tokenAPI.DepositCollateral()
	if tx == nil || err != nil {
		return common.Hash{}, nil, err
	}

	coinsDeposited := new(big.Int).Div(newAmount, big.NewInt(params.Ether))

	return tx.Hash(), coinsDeposited, nil
}

func (w *worker) getBlockReward(proxy common.Address, blockNumber *big.Int) (
	*big.Int, error) {
	contract, err := energi_abi.NewIBlockReward(proxy, w.apiBackend)
	if err != nil {
		return nil, err
	}

	resp, err := contract.GetReward(&bind.CallOpts{}, blockNumber)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *worker) tokenRegistry(dst common.Address) (*energi_abi.IMasternodeTokenSession, error) {
	contract, err := energi_abi.NewIMasternodeToken(
		energi_params.Energi_MasternodeToken, w.apiBackend)
	if err != nil {
		return nil, err
	}
	session := &energi_abi.IMasternodeTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending:  true,
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
		energi_params.Energi_MasternodeRegistry, w.apiBackend)
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

		// MN-17: force transaction creation even when unlocked for staking only
		h := signer.Hash(tx)
		sig, err := wallet.SignHash(account, h[:])
		if err != nil {
			return nil, err
		}

		return tx.WithSignature(signer, sig)
	}
}
