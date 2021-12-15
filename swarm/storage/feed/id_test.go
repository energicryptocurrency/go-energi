// Copyright 2018 The Energi Core Authors
// Copyright 2018 The go-ethereum Authors
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

package feed

import (
	"testing"

	"github.com/energicryptocurrency/energi/swarm/storage/feed/lookup"
)

func getTestID() *ID {
	return &ID{
		Feed:  *getTestFeed(),
		Epoch: lookup.GetFirstEpoch(1000),
	}
}

func TestIDAddr(t *testing.T) {
	id := getTestID()
	updateAddr := id.Addr()
	compareByteSliceToExpectedHex(t, "updateAddr", updateAddr, "0x8b24583ec293e085f4c78aaee66d1bc5abfb8b4233304d14a349afa57af2a783")
}

func TestIDSerializer(t *testing.T) {
	testBinarySerializerRecovery(t, getTestID(), "0x776f726c64206e657773207265706f72742c20657665727920686f7572000000876a8936a7cd0b79ef0735ad0896c1afe278781ce803000000000019")
}

func TestIDLengthCheck(t *testing.T) {
	testBinarySerializerLengthCheck(t, getTestID())
}
