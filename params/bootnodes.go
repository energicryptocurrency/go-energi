// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://6f7d6401de7be03fe172777590e97a2e5198b6a9ad4986ae3b7fd53a03961c640d4e88e1816b9be91671e07ba63c2e64d712425e4a7eec510e70f653efd17f9d@18.188.121.119:39797",
	"enode://4a9f03b85219dc587a763ac0eb282f13bcfafed821703fb62d57854b4cd153866fb872f7461eb9db5e0cd1147e23aa37212db567edd076ef3d959fa96bee8a08@18.130.217.124:39797",
	"enode://1d9219f99fdfec678e4ab8dc9846cd4838a944ddfefbaabe10038a4e2e45ec28f43fd11ad8b60b4026165b9a4020f685c722ea86e3b9d9808bf92c3651c4f685@13.209.138.91:39797",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://ccb54af42913ed8fa23b890c26f13e305057f8036aab94dd6e884da26c5853fe60079edf6a180cee72f395602c18bd115bbc06e41353d62589db52f26ac25964@34.211.246.66:49797",
	"enode://439da3ce5873d7afb3f4516ecc8bf3fd3ff3772f00fdbe9248d13978ab9f2e5a95aeb20e808ef5248ca5e071853202c0bdaf9aae8c5bf907423de5d575f34665@52.36.22.29:49797",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
