// Copyright 2018 The Energi Core Authors
// Copyright 2015 The go-ethereum Authors
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

package eth

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/energicryptocurrency/energi/eth/downloader"
	// "github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/p2p"
	"github.com/energicryptocurrency/energi/p2p/enode"
)

// Tests that fast sync gets disabled as soon as a real block is successfully
// imported into the blockchain.
func TestFastSyncDisabling(t *testing.T) {
	// log.Root().SetHandler(log.StdoutHandler)
	// Create a pristine protocol manager, check that fast sync is left enabled
	pmEmpty, _ := newTestProtocolManagerMust(t, downloader.FastSync, 0, nil, nil)
	if atomic.LoadUint32(&pmEmpty.fastSync) == 0 {
		t.Fatalf("fast sync disabled on pristine blockchain")
	}
	// Blocks count to sync to must be a multiple of MaxHeaderFetch to avoid the
	// block fetcher from blocking forever.
	blocksToSync := 1152 // = MaxHeaderFetch * 6
	// Create a full protocol manager, check that fast sync gets disabled
	pmFull, _ := newTestProtocolManagerMust(t, downloader.FastSync, blocksToSync, nil, nil)
	if atomic.LoadUint32(&pmFull.fastSync) == 1 {
		t.Fatalf("fast sync not disabled on non-empty blockchain")
	}
	// Sync up the two peers
	io1, io2 := p2p.MsgPipe()

	go pmFull.handle(pmFull.newPeer(nrg70, p2p.NewPeer(enode.ID{}, "empty", nil), io2))
	go pmEmpty.handle(pmEmpty.newPeer(nrg70, p2p.NewPeer(enode.ID{}, "full", nil), io1))

	time.Sleep(250 * time.Millisecond)

	// This discards the checkpoint request from the writer clearing way for others
	// to use the writer freely.
	if err := p2p.ExpectMsg(io1, GetCheckpointsMsg, nil); err != nil {
		t.Fatalf(" GetCheckpointsMsg returned an error")
	}

	pmEmpty.synchronise(pmEmpty.peers.BestPeer())

	// Check that fast sync was disabled
	if atomic.LoadUint32(&pmEmpty.fastSync) == 1 {
		t.Fatalf("fast sync not disabled after successful synchronisation")
	}
}
