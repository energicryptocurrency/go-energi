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
	"crypto/sha256"
	"io"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"

	"github.com/shengdoushi/base58"
	"golang.org/x/crypto/ripemd160"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	base54PrivateKeyLen int = 52
	privateKeyLen       int = 32
)

type MigrationAPI struct {
	backend ethapi.Backend
}

func NewMigrationAPI(b ethapi.Backend) *MigrationAPI {
	return &MigrationAPI{b}
}

type Gen2Coin struct {
	ItemID   uint64
	RawOwner common.Address
	Owner    string
	Amount   *big.Int
}

type Gen2Key struct {
	RawOwner common.Address
	Key      *ecdsa.PrivateKey
}

func (m *MigrationAPI) ListGen2Coins() (coins []Gen2Coin) {
	mgrt_contract, err := energi_abi.NewGen2MigrationCaller(
		energi_params.Energi_MigrationContract, m.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed to create contract face", "err", err)
		return []Gen2Coin{}
	}

	call_opts := &bind.CallOpts{}
	bigItems, err := mgrt_contract.ItemCount(call_opts)
	if err != nil {
		log.Error("Failed to get coin count", "err", err)
		return []Gen2Coin{}
	}

	items := bigItems.Int64()
	coins = make([]Gen2Coin, 0, items)

	prefix := byte(33)
	if m.backend.ChainConfig().ChainID.Int64() == 49797 {
		prefix = byte(127)
	}

	for i := int64(0); i < items; i++ {
		res, err := mgrt_contract.Coins(call_opts, big.NewInt(i))
		if err != nil {
			log.Error("Failed to get coin info", "err", err)
			return []Gen2Coin{}
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
			Amount:   res.Amount,
		})
	}

	return
}

func (m *MigrationAPI) SearchGen2Coins(
	owners []string,
	include_empty bool,
) (coins []Gen2Coin) {
	rawOwners := make([]common.Address, len(owners))
	for i, o := range owners {
		ro, err := base58.Decode(o, base58.BitcoinAlphabet)
		if err != nil || len(ro) < 20 {
			log.Error("Failed to decode owner", "err", err, "owner", o)
			continue
		}
		rawOwners[i] = common.BytesToAddress(ro[1 : len(ro)-4])
	}
	return m.searchGen2Coins(rawOwners, m.ListGen2Coins(), include_empty)
}

func (m *MigrationAPI) SearchRawGen2Coins(
	rawOwners []common.Address,
	include_empty bool,
) (coins []Gen2Coin) {
	return m.searchGen2Coins(rawOwners, m.ListGen2Coins(), include_empty)
}

func (m *MigrationAPI) searchGen2Coins(
	owners []common.Address,
	all_coins []Gen2Coin,
	include_empty bool,
) (coins []Gen2Coin) {
	coins = make([]Gen2Coin, 0, len(owners))

	owners_map := make(map[common.Address]bool)
	for _, o := range owners {
		owners_map[o] = true
	}

	for _, c := range all_coins {
		if _, ok := owners_map[c.RawOwner]; ok {
			if include_empty || c.Amount.Cmp(common.Big0) > 0 {
				coins = append(coins, c)
			}
		}
	}

	return coins
}

func (m *MigrationAPI) loadGen2Dump(file string) (keys []Gen2Key) {
	f, err := os.Open(file)
	if err != nil {
		log.Error("Failed to open dump file", "err", err)
		return nil
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Error("Failed to stat file", "err", err)
		return nil
	}

	buf := make([]byte, fi.Size())
	len, err := io.ReadFull(f, buf)
	if err != nil {
		log.Error("Failed to read file", "err", err)
		return nil
	}

	return m.parseGen2Dump(string(buf[:len]))
}

func (m *MigrationAPI) parseGen2Dump(data string) (keys []Gen2Key) {
	lines := strings.Split(data, "\n")
	keys = make([]Gen2Key, 0, len(lines))

	for i, l := range lines {
		lp := strings.Split(l, " ")
		if len(lp) < 3 {
			continue
		}

		tkey := lp[0]
		if len(tkey) != base54PrivateKeyLen {
			continue
		}

		rkey, err := base58.Decode(tkey, base58.BitcoinAlphabet)
		if err != nil {
			log.Error("Failed to decode key", "err", err, "line", i)
			continue
		}

		// There is prefix + key + [magic +] checksum
		key, err := crypto.ToECDSA(rkey[1 : 1+privateKeyLen])
		if err != nil {
			log.Error("Failed to create key", "err", err, "len", len(rkey)*8, "line", i)
			continue
		}

		var owner common.Address

		basehash := sha256.Sum256(crypto.CompressPubkey(&key.PublicKey))
		ripemd := ripemd160.New()
		ripemd.Write(basehash[:])
		owner.SetBytes(ripemd.Sum(nil))

		keys = append(keys, Gen2Key{
			RawOwner: owner,
			Key:      key,
		})
	}

	return
}
