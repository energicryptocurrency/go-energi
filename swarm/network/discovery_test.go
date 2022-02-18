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

package network

import (
	"testing"

	p2ptest "github.com/energicryptocurrency/energi/p2p/testing"
)

/***
 *
 * - after connect, that outgoing subpeersmsg is sent
 *
 */
func TestDiscovery(t *testing.T) {
	params := NewHiveParams()
	s, pp := newHiveTester(t, params, 1, nil)

	node := s.Nodes[0]
	raddr := NewAddr(node)
	_ = pp.Register(raddr)

	// start the hive and wait for the connection
	_ = pp.Start(s.Server)
	defer pp.Stop()

	// send subPeersMsg to the peer
	err := s.TestExchanges(p2ptest.Exchange{
		Label: "outgoing subPeersMsg",
		Expects: []p2ptest.Expect{
			{
				Code: 1,
				Msg:  &subPeersMsg{Depth: 0},
				Peer: node.ID(),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
