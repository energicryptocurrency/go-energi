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
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"io"
	"math/big"
	"os"
	"strings"

	"energi.world/core/gen3/accounts"
	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/accounts/keystore"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/common/hexutil"
	"energi.world/core/gen3/crypto"
	"energi.world/core/gen3/log"
	"energi.world/core/gen3/rpc"

	"github.com/shengdoushi/base58"
	"golang.org/x/crypto/ripemd160"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_common "energi.world/core/gen3/energi/common"
	energi_consensus "energi.world/core/gen3/energi/consensus"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	base54PrivateKeyLen int    = 52
	privateKeyLen       int    = 32
	migrationGas        uint64 = 100000
	ownerSafetyLimit    int    = 10000
)

type MigrationAPI struct {
	backend    Backend
	coinsCache *energi_common.CacheStorage

	lastCoins   interface{}
	lastBalance *big.Int
}

func NewMigrationAPI(b Backend) *MigrationAPI {
	r := &MigrationAPI{
		backend:    b,
		coinsCache: energi_common.NewCacheStorage(),
	}
	b.OnSyncedHeadUpdates(func() {
		r.listGen2Coins()
	})
	return r
}

type Gen2Coin struct {
	ItemID   uint64
	RawOwner common.Address
	Owner    string
	Amount   *hexutil.Big
}

type Gen2Key struct {
	RawOwner common.Address
	Key      *ecdsa.PrivateKey
}

func (m *MigrationAPI) ListGen2Coins() (coins []Gen2Coin, err error) {
	if m.backend.IsPublicService() {
		return nil, errors.New("This API is disabled for security reasons")
	}

	return m.listGen2Coins()
}

func (m *MigrationAPI) listGen2Coins() (coins []Gen2Coin, err error) {
	data, err := m.coinsCache.Get(m.backend, m.listGen2CoinsUncached)
	if err != nil || data == nil {
		log.Error("listGen2Coins failed", "err", err)
		return
	}

	coins = data.([]Gen2Coin)

	return
}

func (m *MigrationAPI) listGen2CoinsUncached(num *big.Int) (interface{}, error) {
	// Check if it makes sense to re-create the list at all
	//---
	state, _, err := m.backend.StateAndHeaderByNumber(context.Background(), rpc.BlockNumber(num.Int64()))
	if err != nil {
		log.Error("Failed to get migration state", "err", err)
		return nil, err
	}

	currBalance := state.GetBalance(energi_params.Energi_MigrationContract)

	if m.lastCoins != nil && m.lastBalance.Cmp(currBalance) == 0 {
		return m.lastCoins, nil
	}

	//---
	log.Info("Preparing a new migration coin list")

	mgrt_contract, err := energi_abi.NewGen2MigrationCaller(
		energi_params.Energi_MigrationContract, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed to create contract face", "err", err)
		return nil, err
	}

	call_opts := &bind.CallOpts{
		Pending:  true,
		GasLimit: energi_params.UnlimitedGas,
	}
	bigItems, err := mgrt_contract.ItemCount(call_opts)
	if err != nil {
		log.Error("Failed to get coin count", "err", err)
		return nil, err
	}

	items := bigItems.Int64()
	coins := make([]Gen2Coin, 0, items)

	prefix := byte(33)
	if m.backend.ChainConfig().ChainID.Int64() == 49797 {
		prefix = byte(127)
	}

	for i := int64(0); i < items; i++ {
		res, err := mgrt_contract.Coins(call_opts, big.NewInt(i))
		if err != nil {
			log.Error("Failed to get coin info", "err", err)
			return nil, err
		}

		owner := make([]byte, 25)
		owner[0] = prefix
		copy(owner[1:], res.Owner[:])
		ownerhash := sha256.Sum256(owner[:21])
		ownerhash = sha256.Sum256(ownerhash[:])
		copy(owner[21:], ownerhash[:4])

		coins = append(coins, Gen2Coin{
			ItemID:   uint64(i),
			RawOwner: common.BytesToAddress(res.Owner[:]),
			Owner:    base58.Encode(owner, base58.BitcoinAlphabet),
			Amount:   (*hexutil.Big)(res.Amount),
		})
	}

	m.lastBalance = currBalance
	m.lastCoins = coins

	return coins, nil
}

func (m *MigrationAPI) SearchGen2Coins(
	owners []string,
	include_empty bool,
) (coins []Gen2Coin, err error) {
	if m.backend.IsPublicService() && len(owners) > ownerSafetyLimit {
		return nil, errors.New("Too many owners requests.")
	}

	rawOwners := make([]common.Address, len(owners))
	for i, o := range owners {
		ro, err := base58.Decode(o, base58.BitcoinAlphabet)
		if err != nil || len(ro) < 20 {
			log.Error("Failed to decode owner", "err", err, "owner", o)
			continue
		}
		rawOwners[i] = common.BytesToAddress(ro[1 : len(ro)-4])
	}
	return m.searchGen2Coins(rawOwners, m.listGen2Coins, include_empty)
}

type Gen3Dest struct {
	ItemID      uint64
	Gen3Address common.Address
	Amount      *big.Int
}

// SearchGen3DestinationByGen2Address returns the gen3 destination address(s) searched
// by the gen2 address owner(s).
func (m *MigrationAPI) SearchGen3DestinationByGen2Address(
	gen2Owners []string,
	includeEmpty bool,
) ([]Gen3Dest, error) {
	mgrtContract, err := energi_abi.NewGen2Migration(
		energi_params.Energi_MigrationContract, m.backend.(bind.ContractBackend))
	if err != nil {
		log.Error("Failed to create contract face", "err", err)
		return nil, err
	}

	currBlockNo := m.backend.CurrentBlock().NumberU64()
	filterOps := &bind.FilterOpts{
		Context: context.WithValue(
			context.Background(),
			energi_params.GeneralProxyCtxKey,
			energi_common.GeneralProxyHashGen(m.backend.BlockChain()),
		),
		End: &currBlockNo,
	}

	migrations, err := mgrtContract.FilterMigrated(filterOps)
	if err != nil {
		log.Error("Failed to fetch all migrations", "err", err)
		return nil, err
	}

	// Release the resources on exit.
	defer migrations.Close()

	coins, err := m.SearchGen2Coins(gen2Owners, includeEmpty)
	if err != nil {
		return nil, err
	}

	data := make([]Gen3Dest, 0, len(gen2Owners))
	for migrations.Next() {
		mEvent := migrations.Event

		for _, coin := range coins {
			if mEvent.ItemId.Cmp(new(big.Int).SetUint64(coin.ItemID)) == 0 {
				dst := Gen3Dest{
					ItemID:      mEvent.ItemId.Uint64(),
					Amount:      mEvent.Amount,
					Gen3Address: mEvent.Destination,
				}
				data = append(data, dst)
				break
			}
		}
	}

	if err = migrations.Error(); err != nil {
		log.Error("Migrations fetch error", "err", err)
		return nil, err
	}

	return data, err
}

func (m *MigrationAPI) SearchRawGen2Coins(
	rawOwners []common.Address,
	include_empty bool,
) (coins []Gen2Coin, err error) {
	if m.backend.IsPublicService() && len(rawOwners) > ownerSafetyLimit {
		return nil, errors.New("Too many owners requests.")
	}

	return m.searchGen2Coins(rawOwners, m.listGen2Coins, include_empty)
}

type listCoins func() (coins []Gen2Coin, err error)

func (m *MigrationAPI) searchGen2Coins(
	owners []common.Address,
	all_coins listCoins,
	include_empty bool,
) (coins []Gen2Coin, err error) {
	coins = make([]Gen2Coin, 0, len(owners))

	owners_map := make(map[common.Address]bool)
	for _, o := range owners {
		owners_map[o] = true
	}

	list, err := all_coins()
	if err != nil {
		log.Error("Failed to get all coins", "err", err)
		return
	}

	for _, c := range list {
		if _, ok := owners_map[c.RawOwner]; ok {
			if include_empty || c.Amount.ToInt().Cmp(common.Big0) > 0 {
				coins = append(coins, c)
			}
		}
	}

	return coins, nil
}

func (m *MigrationAPI) loadGen2Dump(file string) (keys []Gen2Key, err error) {
	f, err := os.Open(file)
	if err != nil {
		log.Error("Failed to open dump file", "err", err)
		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Error("Failed to stat file", "err", err)
		return nil, err
	}

	buf := make([]byte, fi.Size())
	len, err := io.ReadFull(f, buf)
	if err != nil {
		log.Error("Failed to read file", "err", err)
		return nil, err
	}

	return m.parseGen2Dump(string(buf[:len])), nil
}

func (m *MigrationAPI) parseGen2Dump(data string) (keys []Gen2Key) {
	lines := strings.Split(data, "\n")
	keys = make([]Gen2Key, 0, len(lines))

	for i, l := range lines {
		lp := strings.Split(l, " ")
		if len(lp) < 3 || lp[0] == "#" {
			continue
		}

		key, err := m.parseGen2Key(lp[0])
		if err != nil {
			log.Error("Failed to parse key", "err", err, "line", i)
			continue
		}

		keys = append(keys, *key)
	}

	return
}

func (m *MigrationAPI) parseGen2Key(tkey string) (*Gen2Key, error) {
	if len(tkey) != base54PrivateKeyLen {
		return nil, errors.New("Invalid private key length")
	}

	rkey, err := base58.Decode(tkey, base58.BitcoinAlphabet)
	if err != nil {
		return nil, err
	}

	// There is prefix + key + [magic +] checksum
	key_obj, err := crypto.ToECDSA(rkey[1 : 1+privateKeyLen])
	if err != nil {
		return nil, err
	}

	var owner common.Address

	basehash := sha256.Sum256(crypto.CompressPubkey(&key_obj.PublicKey))
	ripemd := ripemd160.New()
	ripemd.Write(basehash[:])
	owner.SetBytes(ripemd.Sum(nil))

	return &Gen2Key{
		RawOwner: owner,
		Key:      key_obj,
	}, nil
}

func (m *MigrationAPI) ClaimGen2CoinsDirect(
	password *string,
	dst common.Address,
	tkey string,
) (txhash common.Hash, err error) {
	key, err := m.parseGen2Key(tkey)
	if err != nil {
		log.Error("Failed to parse key", "err", err)
		return
	}

	coins, err := m.SearchRawGen2Coins([]common.Address{key.RawOwner}, false)
	if err != nil {
		return
	}

	if len(coins) != 1 {
		log.Error("Unable to find coins")
		err = errors.New("No coins found")
		return
	}

	txhash, err = m.claimGen2Coins(password, dst, &coins[0], key)
	if err != nil {
		log.Error("Failed to claim", "err", err)
	}

	return
}

func (m *MigrationAPI) ClaimGen2CoinsCombined(
	password *string,
	dst common.Address,
	file string,
) (txhashes []common.Hash, err error) {
	keys, err := m.loadGen2Dump(file)
	if err != nil {
		return nil, err
	}

	raw_owners := make([]common.Address, len(keys))
	owner2key := make(map[common.Address]*Gen2Key, len(keys))
	for i, k := range keys {
		raw_owners[i] = k.RawOwner
		owner2key[k.RawOwner] = &keys[i]
	}

	coins, err := m.SearchRawGen2Coins(raw_owners, false)
	if err != nil {
		return
	}

	txhashes = make([]common.Hash, len(coins))
	for _, c := range coins {
		txhash, err := m.claimGen2Coins(password, dst, &c, owner2key[c.RawOwner])
		if err != nil {
			return nil, err
		}

		txhashes = append(txhashes, txhash)
	}

	return txhashes, nil
}

func (m *MigrationAPI) ClaimGen2CoinsImport(
	password string,
	file string,
) (txhashes []common.Hash, err error) {
	keys, err := m.loadGen2Dump(file)
	if err != nil {
		return
	}

	raw_owners := make([]common.Address, len(keys))
	owner2key := make(map[common.Address]*Gen2Key, len(keys))
	for i, k := range keys {
		raw_owners[i] = k.RawOwner
		owner2key[k.RawOwner] = &keys[i]
	}

	coins, err := m.SearchRawGen2Coins(raw_owners, false)
	if err != nil {
		return
	}

	am := m.backend.AccountManager()
	ks := am.Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	txhashes = make([]common.Hash, len(coins))
	for _, c := range coins {
		key := owner2key[c.RawOwner]
		dst := crypto.PubkeyToAddress(key.Key.PublicKey)

		//----
		sink := make(chan accounts.WalletEvent)
		evtsub := am.Subscribe(sink)
		defer evtsub.Unsubscribe()

		if _, err := ks.ImportECDSA(key.Key, password); err != nil {
			log.Warn("Failed to import private key", "err", err)
			// Most likely key exists
		} else {
			select {
			case <-sink:
			}
		}

		evtsub.Unsubscribe()
		//----

		txhash, err := m.claimGen2Coins(&password, dst, &c, key)
		if err != nil {
			return nil, err
		}

		txhashes = append(txhashes, txhash)
	}

	return txhashes, nil
}

func (m *MigrationAPI) claimGen2Coins(
	password *string,
	dst common.Address,
	coin *Gen2Coin,
	key *Gen2Key,
) (txhash common.Hash, err error) {
	mgrt_contract_obj, err := energi_abi.NewGen2Migration(
		energi_params.Energi_MigrationContract, m.backend.(bind.ContractBackend))
	if err != nil {
		return
	}

	mgrt_contract := energi_abi.Gen2MigrationSession{
		Contract: mgrt_contract_obj,
		CallOpts: bind.CallOpts{
			Pending:  true,
			From:     dst,
			GasLimit: energi_params.UnlimitedGas,
		},
		TransactOpts: bind.TransactOpts{
			From:     dst,
			Signer:   createSignerCallback(m.backend, password),
			Value:    common.Big0,
			GasPrice: common.Big0,
			GasLimit: migrationGas,
		},
	}

	hts, err := mgrt_contract.HashToSign(dst)
	if err != nil {
		return
	}

	sig, err := crypto.Sign(hts[:], key.Key)
	if err != nil {
		return
	}

	if len(sig) != 65 {
		err = errors.New("Wrong signature size")
		return
	}

	item := new(big.Int).SetUint64(coin.ItemID)
	r := [32]byte{}
	copy(r[:], sig[:32])
	s := [32]byte{}
	copy(s[:], sig[32:64])
	v := uint8(sig[64])

	amt, err := mgrt_contract.VerifyClaim(item, dst, v, r, s)
	if err != nil {
		return
	}

	if amt.Cmp(common.Big0) == 0 {
		log.Warn("Already claimed", "coins", coin.Owner)
		return
	}

	tx, err := mgrt_contract.Claim(item, dst, v, r, s)
	if tx != nil {
		txhash = tx.Hash()
		log.Info("Sent migration transaction", "tx", tx.Hash().Hex(), "coins", coin.Owner)
	}

	return
}

type MigrationAdminAPI struct {
	backend Backend
}

func NewMigrationAdminAPI(b Backend) *MigrationAdminAPI {
	return &MigrationAdminAPI{b}
}

func (b *MigrationAdminAPI) ValidateMigration(
	migration_file string,
) bool {
	block, err := b.backend.BlockByNumber(context.Background(), 1)
	if block == nil || err != nil {
		log.Error("Migration block is missing")
		return false
	}

	return energi_consensus.ValidateMigration(block, migration_file)
}
