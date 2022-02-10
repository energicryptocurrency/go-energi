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
	"crypto/ecdsa"
	crand "crypto/rand"
	"math/big"

	ethereum "github.com/energicryptocurrency/energi"
	"github.com/energicryptocurrency/energi/accounts"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/crypto"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/event"
)

type EphemeralWallet struct{}

func (ew *EphemeralWallet) URL() accounts.URL {
	return accounts.URL{Scheme: "ephemeral"}
}

func (ew *EphemeralWallet) Status() (string, error) {
	return "Unlocked", nil
}

func (ew *EphemeralWallet) Open(passphrase string) (err error) {
	return
}

func (ew *EphemeralWallet) Close() error {
	return nil
}

func (ew *EphemeralWallet) Accounts() []accounts.Account {
	return []accounts.Account{
		{Address: energi_params.Energi_Ephemeral, URL: ew.URL()},
	}
}

func (ew *EphemeralWallet) Contains(account accounts.Account) bool {
	return account.Address == energi_params.Energi_Ephemeral
}

func (ew *EphemeralWallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	return accounts.Account{}, accounts.ErrNotSupported
}

func (ew *EphemeralWallet) SelfDerive(bases []accounts.DerivationPath, chain ethereum.ChainStateReader) {
}

func (ew *EphemeralWallet) SignData(account accounts.Account, mimeType string, data []byte) ([]byte, error) {
	return ew.SignHash(account, crypto.Keccak256(data))
}

// SignDataWithPassphrase signs keccak256(data). The mimetype parameter describes the type of data being signed.
func (ew *EphemeralWallet) SignDataWithPassphrase(account accounts.Account, passphrase, mimeType string, data []byte) ([]byte, error) {
	return nil, nil
}

// SignText implements accounts.Wallet, attempting to sign the hash of
// the given text with the given account.
func (ew *EphemeralWallet) SignText(account accounts.Account, text []byte) ([]byte, error) {
	return ew.SignHash(account, accounts.TextHash(text))
}

// SignTextWithPassphrase implements accounts.Wallet, attempting to sign the
// hash of the given text with the given account using passphrase as extra authentication.
func (ew *EphemeralWallet) SignTextWithPassphrase(account accounts.Account, passphrase string, text []byte) ([]byte, error) {
	return nil, nil
}

func (ew *EphemeralWallet) SignHash(account accounts.Account, hash []byte) ([]byte, error) {
	if !ew.Contains(account) {
		return nil, accounts.ErrUnknownAccount
	}

	privateKey, err := ecdsa.GenerateKey(crypto.S256(), crand.Reader)
	if err != nil {
		return nil, err
	}

	return crypto.Sign(hash, privateKey)
}

func (ew *EphemeralWallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if !ew.Contains(account) {
		return nil, accounts.ErrUnknownAccount
	}

	privateKey, err := ecdsa.GenerateKey(crypto.S256(), crand.Reader)
	if err != nil {
		return nil, err
	}

	if chainID != nil {
		return types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	}
	return types.SignTx(tx, types.HomesteadSigner{}, privateKey)
}

func (ew *EphemeralWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	return ew.SignHash(account, hash)
}

func (ew *EphemeralWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return ew.SignTx(account, tx, chainID)
}

func (ew *EphemeralWallet) IsUnlockedForStaking(account accounts.Account) bool {
	return false
}

type EphemeralAccount struct {
	wallet      EphemeralWallet
	updateFeed  event.Feed
	updateScope event.SubscriptionScope
}

func NewEphemeralAccount() (*EphemeralAccount, error) {
	ea := &EphemeralAccount{}
	return ea, ea.wallet.Open("")
}

func (ea *EphemeralAccount) Wallets() []accounts.Wallet {
	return []accounts.Wallet{&ea.wallet}
}

func (ea *EphemeralAccount) Subscribe(sink chan<- accounts.WalletEvent) event.Subscription {
	return ea.updateScope.Track(ea.updateFeed.Subscribe(sink))
}
