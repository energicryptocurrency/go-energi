// Copyright 2018 The Energi Core Authors
// Copyright 2014 The go-ethereum Authors
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

package types

import (
	"bytes"
	"math/big"
	"reflect"
	"testing"

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/rlp"
)

// from bcValidBlockTest.json, "SimpleTx"
func TestBlockEncoding(t *testing.T) {
	txTo := common.HexToAddress("095e7baea6a6c7c4c2dfeb977efac326af552d87")
	var b = Block{
		header: &Header{
			UncleHash:   common.HexToHash("1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
			Coinbase:    common.HexToAddress("8888f1f195afa192cfee860698584c030f4c9db1"),
			Root:        common.HexToHash("ef1552a40b7165c3cd773806b9e0c165b75356e0314bf0706f279c729f51e017"),
			TxHash:      common.HexToHash("ac81275d7a81a720240146377982939179218535bfcfa9469c8bdd3e264ef179"),
			ReceiptHash: common.HexToHash("37aee413d1d4f1335720ed1ee882bac4b2d41c8072453b23445020c951682054"),
			Bloom:       BytesToBloom(common.Hex2Bytes("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
			Difficulty:  big.NewInt(0x020000),
			Number:      big.NewInt(0x01),
			GasLimit:    0x2fefd8,
			GasUsed:     0x5208,
			Extra:       common.Hex2Bytes("0x"),
			Time:        1426516743,
			MixDigest:   common.HexToHash("bd4472abb6659ebe3ee06ee4d7b72a00a9f4d001caca51342001075469aff498"),
			Nonce:       EncodeNonce(11617697748499542468),
			Signature:   common.Hex2Bytes("f90231f901faa00000000000000000000000000000000000000000000000000000000000000000a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347948888f1f195afa192cfee860698584c030f4c9db1a0ef1552a40b7165c3cd773806b9e0c165b75356e0314bf0706f279c729f51e017a0ac81275d7a81a720240146377982939179218535bfcfa9469c8bdd3e264ef179a037aee413d1d4f1335720ed1ee882bac4b2d41c8072453b23445020c951682054b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008302000001832fefd8825208845506eb0780a0bd4472abb6659ebe3ee06ee4d7b72a00a9f4d001caca51342001075469aff49888a13a5a8c8f2bb1c480f2f1800182753094095e7baea6a6c7c4c2dfeb977efac326af552d8785012a05f200801c86057cb46c0b708705ba7c83d99be6c0"),
		},
		uncles: []*Header{},
		transactions: Transactions{
			{
				data: txdata{
					AccountNonce: 0,
					GasLimit:     50000,
					Price:        big.NewInt(10),
					Recipient:    &txTo,
					Amount:       big.NewInt(10),
					V:            big.NewInt(1),
					R:            big.NewInt(7400161059665675640),
					S:            big.NewInt(3482875120068134972),
				},
			},
		},
	}

	blockEnc, err := rlp.EncodeToBytes(&b)
	if err != nil {
		t.Fatal("encode error: ", err)
	}

	var block Block
	if err := rlp.DecodeBytes(blockEnc, &block); err != nil {
		t.Fatal("decode error: ", err)
	}

	check := func(f string, got, want interface{}) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s mismatch: got %v, want %v", f, got, want)
		}
	}
	check("Difficulty", block.Difficulty(), big.NewInt(131072))
	check("GasLimit", block.GasLimit(), uint64(3141592))
	check("GasUsed", block.GasUsed(), uint64(21000))
	check("Coinbase", block.Coinbase(), common.HexToAddress("8888f1f195afa192cfee860698584c030f4c9db1"))
	check("MixDigest", block.MixDigest(), common.HexToHash("bd4472abb6659ebe3ee06ee4d7b72a00a9f4d001caca51342001075469aff498"))
	check("Root", block.Root(), common.HexToHash("ef1552a40b7165c3cd773806b9e0c165b75356e0314bf0706f279c729f51e017"))
	check("Hash", block.Hash(), common.HexToHash("0eebd1275d0eaa172e34ea388739b4db0cd9307b0f65f56467da8fa568ed35be"))
	check("Nonce", block.Nonce(), uint64(0xa13a5a8c8f2bb1c4))
	check("Time", block.Time(), uint64(1426516743))
	check("Size", block.Size(), common.StorageSize(len(blockEnc)))
	check("len(Transactions)", len(block.Transactions()), 1)
	check("Transactions[0].Hash", block.Transactions()[0].Hash(), common.HexToHash("1f0ebbb093b13188cdd7c667fc9132fdc8331829ad6262b101046b4e1932e99c"))

	ourBlockEnc, err := rlp.EncodeToBytes(&block)
	if err != nil {
		t.Fatal("encode error: ", err)
	}
	if !bytes.Equal(ourBlockEnc, blockEnc) {
		t.Errorf("encoded block mismatch:\ngot:  %x\nwant: %x", ourBlockEnc, blockEnc)
	}
}
