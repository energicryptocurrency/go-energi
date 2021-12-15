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

package consensus

import (
	"math/big"
	"strings"
	"testing"

	"github.com/energicryptocurrency/energi/common"

	"github.com/stretchr/testify/assert"
)

const (
	testSnapshotData = `
{
  "snapshot_utxos": [
    {
      "owner": "t6vtJKxdjaJdofaUrx7w4xUs5bMcjDq5R2",
      "amount": 10228000000,
      "type": "pubkeyhash"
    },
    {
      "owner": "tWFyUdwGxEkcam2aikVsDMPDpvMNKfP2XV",
      "amount": 1000010,
      "type": "pubkeyhash"
    }
  ],
  "snapshot_blacklist": [
    "tWFyUdwGxEkcam2aikVsDMPDpvMNKfP2XV"
  ],
  "snapshot_hash": "778d7a438e3b86e0e754c4e46af802f852eb7c051d268c8599aa17c0cb9ce819"
}
`
)

func TestSnapshotParser(t *testing.T) {
	t.Parallel()

	ret, err := parseSnapshot(strings.NewReader(testSnapshotData))
	assert.Empty(t, err)
	assert.Equal(
		t,
		ret.Hash,
		"778d7a438e3b86e0e754c4e46af802f852eb7c051d268c8599aa17c0cb9ce819")
	assert.Equal(
		t,
		len(ret.Txouts),
		2)
	assert.Equal(
		t,
		ret.Txouts[0].Owner,
		"t6vtJKxdjaJdofaUrx7w4xUs5bMcjDq5R2")
	assert.Equal(
		t,
		ret.Txouts[0].Atype,
		"pubkeyhash")
	assert.Equal(
		t,
		ret.Txouts[0].Amount.String(),
		big.NewInt(10228000000).String())
	assert.Equal(
		t,
		ret.Txouts[1].Owner,
		"tWFyUdwGxEkcam2aikVsDMPDpvMNKfP2XV")
	assert.Equal(
		t,
		ret.Txouts[1].Atype,
		"pubkeyhash")
	assert.Equal(
		t,
		ret.Txouts[1].Amount.String(),
		big.NewInt(1000010).String())
	assert.Equal(
		t,
		ret.Blacklist[0],
		"tWFyUdwGxEkcam2aikVsDMPDpvMNKfP2XV")
}

func TestSnapshotParams(t *testing.T) {
	t.Parallel()

	snapshot, err := parseSnapshot(strings.NewReader(testSnapshotData))
	assert.Empty(t, err)
	owners, amounts, blacklist := createSnapshotParams(snapshot)
	assert.NotEmpty(t, owners)
	assert.Equal(
		t,
		owners[0].String(),
		common.HexToAddress("0x000D90BA0EFF81760202C28B40563B9636C1CCD4").String())
	assert.Equal(
		t,
		amounts[0].String(),
		"102280000000000000000")
	assert.Equal(
		t,
		owners[1].String(),
		common.HexToAddress("0xFFF4AF0E7421838CDB2F14134F74AA4B5B0A816E").String())
	assert.Equal(
		t,
		amounts[1].String(),
		"10000100000000000")
	assert.Equal(
		t,
		blacklist[0].String(),
		common.HexToAddress("0xFFF4AF0E7421838CDB2F14134F74AA4B5B0A816E").String())
}
