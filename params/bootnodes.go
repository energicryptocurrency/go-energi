// Copyright 2023 The Energi Core Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// gb-node1.backbone.energi.network
	"enode://4a9f03b85219dc587a763ac0eb282f13bcfafed821703fb62d57854b4cd153866fb872f7461eb9db5e0cd1147e23aa37212db567edd076ef3d959fa96bee8a08@18.130.217.124:39797",
	// de_node2_backbone2_energi_network
	"enode://19d706e65a2e87d636c8f21788f7d7342aacfcbe3ee1e10b4f5681c3ba15c879d7bac4e24586f00d85fba655a33922b814f95006d648bba46af2bea485e6a1ef@31.220.93.144:39797",
	// gb-node3.backbone2.energi.network
	"enode://8056fdb7d17f54aa914aa8cc4b55df86266c2ac6b56c18a0ffd15437e78a40b68cc9edfdddb8e9dae0f4bfe9c9d7e929f05954b8f15d43c76848d88a1bcf1ec3@35.178.54.163:39797",
	// kr-node1.backbone2.energi.network
	"enode://1d9219f99fdfec678e4ab8dc9846cd4838a944ddfefbaabe10038a4e2e45ec28f43fd11ad8b60b4026165b9a4020f685c722ea86e3b9d9808bf92c3651c4f685@13.209.138.91:39797",
	// sg_node2_backbone2_energi_network
	"enode://98d35115395531b6079096ccfc218aa206453ec3d60174670e82552d00b3ddc61bc0e0b3f254e81e8d0a06386406ff33948ebdda3c3e9ff3a2f43c2397726644@46.250.226.242:39797",
	// kr-node3.backbone2.energi.network
	"enode://eead2c8dfc7eb7fe58bd866f9e6a48802f889e4fb5e7c68210375a4c4ff71541153b7ef86f51dc28909de375ee57387a0aeba1072ea7aaf8513f17922df788cb@13.209.220.197:39797",
	// us-node1.backbone2.energi.network
	"enode://6f7d6401de7be03fe172777590e97a2e5198b6a9ad4986ae3b7fd53a03961c640d4e88e1816b9be91671e07ba63c2e64d712425e4a7eec510e70f653efd17f9d@18.188.121.119:39797",
	// us-node2.backbone2.energi.network
	"enode://1a6b27aa5b5becd6bb5998d1853b48b99ade2391f0340e00cd5623e67f1a7239a52571b79a60b7bde0d242a98471e3205958916bc2c4f38d858002bdecbe6052@85.239.240.237:39797",
	// us-node3.backbone2.energi.network
	"enode://a63c81108e1a5a45476f5fed6fadc0c6f8a22d51d50dccd0b7ae2218001b74de6336893d6a0fe9b6469a89601ecd44760e8a9f4b6e596229e2798b5c521e5784@18.223.42.176:39797",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	// us_node1n_test3_energi_network
	"enode://ccb54af42913ed8fa23b890c26f13e305057f8036aab94dd6e884da26c5853fe60079edf6a180cee72f395602c18bd115bbc06e41353d62589db52f26ac25964@144.126.155.146:49797",
	// de_node1_test3_energi_network
	"enode://439da3ce5873d7afb3f4516ecc8bf3fd3ff3772f00fdbe9248d13978ab9f2e5a95aeb20e808ef5248ca5e071853202c0bdaf9aae8c5bf907423de5d575f34665@167.86.90.170:49797",
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
