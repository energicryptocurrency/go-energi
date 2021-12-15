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

package common

import (
	"net"

	"github.com/energicryptocurrency/energi/crypto"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/p2p/enode"
	"github.com/energicryptocurrency/energi/params"
)

func MastenodeEnode(ipv4address uint32, pubkey [2][32]byte, cfg *params.ChainConfig) *enode.Node {
	ip := net.IPv4(
		byte(ipv4address>>24),
		byte(ipv4address>>16),
		byte(ipv4address>>8),
		byte(ipv4address),
	)

	pubkey_buf := make([]byte, 33)
	copy(pubkey_buf[:32], pubkey[0][:])
	copy(pubkey_buf[32:33], pubkey[1][:])
	pk, err := crypto.DecompressPubkey(pubkey_buf)
	if err != nil {
		log.Error("Failed to unmarshal Masternode pubkey")
		return nil
	}

	return enode.NewV4(pk, ip, int(cfg.ChainID.Int64()), int(cfg.ChainID.Int64()))
}
