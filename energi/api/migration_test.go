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
	"math/big"
	"testing"

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/common/hexutil"
	"github.com/energicryptocurrency/energi/crypto"

	"github.com/stretchr/testify/assert"
)

const (
	testWalletDump = `
# bla bla bla bla

cPiXpw35kL9diQjw3phzUJFtJGouiHZY822KSkGg6u2xwwdz2VvL 2019-05-17T23:08:56Z change=1 # addr=tRGs7s5rYCPkAiVnCKmse8ZRpZ9BBFvtSL
cQ8ChCLsNbvWZwGDrTXQVqPtM1srL5GgmGYxJkDYZuFp2n3E54nH 2019-06-13T17:26:00Z reserve=1 # addr=tSvHLh3KR1KwwtgthaaPT3r1teec36NVKz

# ----
`
)

func TestGen2Parse(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	m := NewMigrationAPI(nil)
	res := m.parseGen2Dump(testWalletDump)
	assert.Equal(t, 2, len(res))
	assert.Equal(t,
		"0xC94729d0212C2D1074d858EB6c9ee44Fb19D76e6",
		res[0].RawOwner.String())
	assert.Equal(t,
		"0x3fb403e4227a129129e9c0f01ba3ed79294f1ae18bbd961e7ad7fa0996680c40",
		hexutil.Encode(crypto.FromECDSA(res[0].Key)))
	assert.Equal(t,
		"0xDB52E60435e09e998b6077eE65e3719836fA0d2e",
		res[1].RawOwner.String())
	assert.Equal(t,
		"0x4be146dcb88089732cdbea785bcf5c2c188d67152f0b80329e275c2b553174bc",
		hexutil.Encode(crypto.FromECDSA(res[1].Key)))
}

func TestSearchCoins(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	m := NewMigrationAPI(nil)

	listCoins := func() ([]Gen2Coin, error) {
		return []Gen2Coin{
			{
				ItemID:   77,
				RawOwner: common.HexToAddress("0xC94729d0212C2D1074d858EB6c9ee44Fb19D76e6"),
				Amount:   (*hexutil.Big)(big.NewInt(0)),
			},
			{
				ItemID:   78,
				RawOwner: common.HexToAddress("0xC94729d0212C2D1074d858EB6c9ee44Fb19D76e6"),
				Amount:   (*hexutil.Big)(big.NewInt(10)),
			},
			{
				ItemID:   79,
				RawOwner: common.HexToAddress("0xDB52E60435e09e998b6077eE65e3719836fA0d2e"),
				Amount:   (*hexutil.Big)(big.NewInt(10)),
			},
		}, nil
	}

	res, err := m.searchGen2Coins(
		[]common.Address{
			common.HexToAddress("0xC94729d0212C2D1074d858EB6c9ee44Fb19D76e6"),
		},
		listCoins,
		true,
	)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, err, nil)
	assert.Equal(t, uint64(77), res[0].ItemID)
	assert.Equal(t, uint64(78), res[1].ItemID)

	listCoins = func() ([]Gen2Coin, error) {
		return []Gen2Coin{
			{
				ItemID:   77,
				RawOwner: common.HexToAddress("0xC94729d0212C2D1074d858EB6c9ee44Fb19D76e6"),
				Amount:   (*hexutil.Big)(big.NewInt(0)),
			},
			{
				ItemID:   78,
				RawOwner: common.HexToAddress("0xC94729d0212C2D1074d858EB6c9ee44Fb19D76e6"),
				Amount:   (*hexutil.Big)(big.NewInt(10)),
			},
			{
				ItemID:   79,
				RawOwner: common.HexToAddress("0xDB52E60435e09e998b6077eE65e3719836fA0d2e"),
				Amount:   (*hexutil.Big)(big.NewInt(10)),
			},
		}, nil
	}

	res, err = m.searchGen2Coins(
		[]common.Address{
			common.HexToAddress("0xC94729d0212C2D1074d858EB6c9ee44Fb19D76e6"),
		},
		listCoins,
		false,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, uint64(78), res[0].ItemID)
}
