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

package lookup_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/swarm/storage/feed/lookup"
)

func TestMarshallers(t *testing.T) {

	for i := uint64(1); i < lookup.MaxTime; i *= 3 {
		e := lookup.Epoch{
			Time:  i,
			Level: uint8(i % 20),
		}
		b, err := e.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		var e2 lookup.Epoch
		if err := e2.UnmarshalBinary(b); err != nil {
			t.Fatal(err)
		}
		if e != e2 {
			t.Fatal("Expected unmarshalled epoch to be equal to marshalled onet.Fatal(err)")
		}
	}

}

func TestAfter(t *testing.T) {
	a := lookup.Epoch{
		Time:  5,
		Level: 3,
	}
	b := lookup.Epoch{
		Time:  6,
		Level: 3,
	}
	c := lookup.Epoch{
		Time:  6,
		Level: 4,
	}

	if !b.After(a) {
		t.Fatal("Expected 'after' to be true, got false")
	}

	if b.After(b) {
		t.Fatal("Expected 'after' to be false when both epochs are identical, got true")
	}

	if !b.After(c) {
		t.Fatal("Expected 'after' to be true when both epochs have the same time but the level is lower in the first one, but got false")
	}

}
