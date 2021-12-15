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

	"github.com/energicryptocurrency/energi/common/hexutil"
)

func TestTopic(t *testing.T) {
	related, _ := hexutil.Decode("0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789")
	topicName := "test-topic"
	topic, _ := NewTopic(topicName, related)
	hex := topic.Hex()
	expectedHex := "0xdfa89c750e3108f9c2aeef0123456789abcdef0123456789abcdef0123456789"
	if hex != expectedHex {
		t.Fatalf("Expected %s, got %s", expectedHex, hex)
	}

	var topic2 Topic
	topic2.FromHex(hex)
	if topic2 != topic {
		t.Fatal("Expected recovered topic to be equal to original one")
	}

	if topic2.Name(related) != topicName {
		t.Fatal("Retrieved name does not match")
	}

	bytes, err := topic2.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	expectedJSON := `"0xdfa89c750e3108f9c2aeef0123456789abcdef0123456789abcdef0123456789"`
	equal, err := areEqualJSON(expectedJSON, string(bytes))
	if err != nil {
		t.Fatal(err)
	}
	if !equal {
		t.Fatalf("Expected JSON to be %s, got %s", expectedJSON, string(bytes))
	}

	err = topic2.UnmarshalJSON(bytes)
	if err != nil {
		t.Fatal(err)
	}
	if topic2 != topic {
		t.Fatal("Expected recovered topic to be equal to original one")
	}

}
