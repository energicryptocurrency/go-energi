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

package p2p

import "github.com/energicryptocurrency/energi/p2p/enode"

// IsPeerActive checks if the specified node is already connected as a peer.
func (srv *Server) IsPeerActive(node *enode.Node) bool {
	var isActive bool

	select {
	// Note: We'd love to put this function into a variable but
	// that seems to cause a weird compiler error in some
	// environments.
	case srv.peerOp <- func(peers map[enode.ID]*Peer) {
		for ID := range peers {
			if ID == node.ID() {
				isActive = true
				break
			}
		}
	}:
		<-srv.peerOpDone
	case <-srv.quit:
	}
	return isActive
}
