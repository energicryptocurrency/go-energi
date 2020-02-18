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
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_api "energi.world/core/gen3/energi/api"
	energi "energi.world/core/gen3/energi/consensus"
	energi_params "energi.world/core/gen3/energi/params"
)

const maxAutoCollateralBlockAge = time.Duration(time.Minute)

var mnPayout *big.Int

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

	for _, wallet := range w.APIBackend.AccountManager().Wallets() {
		for _, account := range wallet.Accounts() {
			if wallet.IsUnlockedForStaking(account) {
				// If unlocked, proceed.
				if err := w.superBlockCheck(account.Address, block.Number()); err != nil {
					// Skip if a superblock was found.
					log.Debug(err.Error())
					return
				}

				status, amount, err := w.hasJustReceivedRewards(account.Address, block)
				if err != nil {
					log.Debug(err.Error())
					continue
				}

				if status || amount == nil {
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

func (w *worker) superBlockCheck(dst common.Address, blockNo *big.Int) error {
	reward, err := w.getBlockReward(dst, blockNo)
	if err != nil {
		return err
	}

	mnPayout = reward.ToInt()
	// MN-17 - 4
	// Check for SuperBlock: mnPayout == 0
	if mnPayout.Cmp(common.Big0) == 0 {
		return fmt.Errorf("Possible super block was found")
	}

	return nil
}

func (w *worker) getBalanceAtBlock(block *types.Block, address common.Address) (*hexutil.Big, error) {
	stateDb, err := w.eth.BlockChain().StateAt(block.Root())
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(stateDb.GetBalance(address)), nil
}

// isMNPayouts checks if the provided block in relation to the previous block has
// a masternode block payout.
func (w *worker) isMNPayouts(currentblock *types.Block, mnOwner common.Address) (bool, *hexutil.Big, error) {
	blockNo := currentblock.Number().Uint64()
	if blockNo <= 0 {
		return false, nil, fmt.Errorf("Invalid block cannot posses MN payouts")
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

	diff := new(big.Int).Sub(balanceNow.ToInt(), balancePrev.ToInt())

	// Should be: 0 < diff <= mnPayout
	status := diff.Cmp(mnPayout) < 1 && diff.Cmp(common.Big0) == 1
	return status, balanceNow, nil
}

func (w *worker) hasJustReceivedRewards(account common.Address, block *types.Block) (bool, *hexutil.Big, error) {
	// MN-17 - 5
	// (a) Confirm no MN payouts in the current block.
	isCurrentMNPayout, balanceNow, err := w.isMNPayouts(block, account)
	if err != nil {
		return true, nil, err
	}

	if isCurrentMNPayout {
		return true, nil, fmt.Errorf("Current block masternode payout is active")
	}

	// MN-17 - 5
	// (b) Confirm atleast 1 masternode payout in the previous block.
	prevBlock := w.eth.BlockChain().GetBlockByNumber(block.Number().Uint64() - 1)
	isPrevPayout, _, err := w.isMNPayouts(prevBlock, account)
	if err != nil {
		return true, nil, err
	}

	if !isPrevPayout {
		return true, nil, fmt.Errorf("Expected at least one payout from a previous block")
	}

	return false, balanceNow, nil
}

// canAutocollateralize returns the maximum amount that can be deposited as the
// collateral if the maximum collateral amount is not yet reached.
func (w *worker) canAutocollateralize(
	account common.Address,
	amount *hexutil.Big,
	api *energi_abi.IMasternodeTokenSession,
) (*hexutil.Big, error) {
	minLimit, maxLimit, err := w.CollateralLimits(account)
	if err != nil {
		return nil, err
	}

	// MN-17 - 5
	// (c) Ensures that available balance is at least one minimal collateral.
	if amount.ToInt().Cmp(minLimit.ToInt()) < 0 {
		return nil, fmt.Errorf("Amount found is less than the minimum required")
	}

	tokenBalance, err := api.BalanceOf(account)
	if err != nil {
		return nil, err
	}

	// MN-17 - 5
	// (d) Ensure that the current collateral is below the maximum allowed.
	if tokenBalance.Cmp(maxLimit.ToInt()) >= 0 {
		return nil, fmt.Errorf("Maximum collateral supported already achieved")
	}

	modAmount := new(big.Int).Mod(amount.ToInt(), minLimit.ToInt())
	amountToDeposit := new(big.Int).Sub(amount.ToInt(), modAmount)

	totalAmount := new(big.Int).Add(tokenBalance, amountToDeposit)
	if totalAmount.Cmp(maxLimit.ToInt()) == 1 {
		// Gets the maximum amount to deposit since all the available amount
		// could breach the max collateral limit if deposited in full.
		amountToDeposit = new(big.Int).Sub(maxLimit.ToInt(), tokenBalance)
	}

	return (*hexutil.Big)(amountToDeposit), nil
}

func (w *worker) doAutocollateral(account common.Address, amount *hexutil.Big) (common.Hash, *big.Int, error) {
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
	tokenAPI.TransactOpts.Value = newAmount.ToInt()
	tx, err := tokenAPI.DepositCollateral()
	if tx == nil || err != nil {
		return common.Hash{}, nil, err
	}

	coinsDeposited := new(big.Int).Div(newAmount.ToInt(), big.NewInt(params.Ether))

	return tx.Hash(), coinsDeposited, nil
}

func (w *worker) getBlockReward(dst common.Address, blockNumber *big.Int) (*hexutil.Big, error) {
	contract, err := energi_abi.NewIBlockReward(
		energi_params.Energi_MasternodeRegistry, w.APIBackend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{
		From:     dst,
		GasLimit: energi_params.UnlimitedGas,
	}

	resp, err := contract.GetReward(callOpts, blockNumber)
	if err != nil {
		err = fmt.Errorf("Fetching Block reward failed: %v", err)
		return nil, err
	}

	return (*hexutil.Big)(resp), nil
}

func (w *worker) tokenRegistry(dst common.Address) (*energi_abi.IMasternodeTokenSession, error) {
	contract, err := energi_abi.NewIMasternodeToken(
		energi_params.Energi_MasternodeToken, w.APIBackend.(bind.ContractBackend))
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
			Signer:   energi_api.CreateStakeTxSignerCallback(w.APIBackend),
			Value:    common.Big0,
			GasLimit: energi_api.MasternodeCallGas,
		},
	}
	return session, nil
}

func (w *worker) CollateralLimits(dst common.Address) (minCollateral, maxCollateral *hexutil.Big, err error) {
	registry, err := energi_abi.NewIMasternodeRegistryV2(
		energi_params.Energi_MasternodeRegistry, w.APIBackend.(bind.ContractBackend))
	if err != nil {
		return nil, nil, err
	}

	callOpts := &bind.CallOpts{
		From:     dst,
		GasLimit: energi_params.UnlimitedGas,
	}

	limits, err := registry.CollateralLimits(callOpts)
	if err != nil {
		return nil, nil, err
	}

	minCollateral = (*hexutil.Big)(limits.Min)
	maxCollateral = (*hexutil.Big)(limits.Max)
	return minCollateral, maxCollateral, nil
}
