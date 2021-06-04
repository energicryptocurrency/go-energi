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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// gb-node1.backbone.energi.network
	"enode://4a9f03b85219dc587a763ac0eb282f13bcfafed821703fb62d57854b4cd153866fb872f7461eb9db5e0cd1147e23aa37212db567edd076ef3d959fa96bee8a08@18.130.217.124:39797",
	// gb-node2.backbone.energi.network
	"enode://19d706e65a2e87d636c8f21788f7d7342aacfcbe3ee1e10b4f5681c3ba15c879d7bac4e24586f00d85fba655a33922b814f95006d648bba46af2bea485e6a1ef@35.177.206.152:39797",
	// gb-node3.backbone2.energi.network
	"enode://b220d4308c138004528cb5ff1686c94322db6ca0e720845e8ae49678d30c26fb4f230a79c76d8ac1c6d388b10806f7d9b93978669a13592ac40e5b8b40002ebc@35.178.54.163:39797",
	// gb-node4.backbone2.energi.network
	"enode://3595271c5747f2bac2089fcfba2ca92ed2798141235d9fab3276ca877d089a3e8baa64b66f0537eafe4de4ab3a6d9ff415d5a722f76d6d1b0679a82532a4db3d@52.56.172.161:39797",
	// gb-node5.backbone2.energi.network
	"enode://8e62dc0e08161efceaae4d565d903c8107a103a42f5cd10daf04c12606abdeca94d230e687ced6be26ff94614a5d6d3b1539dcf354635a4d76af84dd62a17217@52.56.87.53:39797",
	// kr-node1.backbone2.energi.network
	"enode://1d9219f99fdfec678e4ab8dc9846cd4838a944ddfefbaabe10038a4e2e45ec28f43fd11ad8b60b4026165b9a4020f685c722ea86e3b9d9808bf92c3651c4f685@13.209.138.91:39797",
	// kr-node2.backbone2.energi.network
	"enode://98d35115395531b6079096ccfc218aa206453ec3d60174670e82552d00b3ddc61bc0e0b3f254e81e8d0a06386406ff33948ebdda3c3e9ff3a2f43c2397726644@13.209.141.8:39797",
	// kr-node3.backbone2.energi.network
	"enode://eead2c8dfc7eb7fe58bd866f9e6a48802f889e4fb5e7c68210375a4c4ff71541153b7ef86f51dc28909de375ee57387a0aeba1072ea7aaf8513f17922df788cb@13.209.220.197:39797",
	// kr-node4.backbone2.energi.network
	"enode://acee420c23259e4af58184ef7f2b7a7315bc7fe414495133167d39a5d70886bc77ea1edb4e9240edbd621217464a10e17dac31bf03d74fad1e4f168c86af66ec@52.78.34.224:39797",
	// kr-node5.backbone2.energi.network
	"enode://785e80f15da881d9d58b4e0cb95eb5129578e614d03d91ee716fdceffba2ea0d57f2300ac715d170bbe447fd6b20ea2e55a9361fcbe930349ed22d0dfe12fd1b@52.79.246.69:39797",
	// us-node1.backbone2.energi.network
	"enode://6f7d6401de7be03fe172777590e97a2e5198b6a9ad4986ae3b7fd53a03961c640d4e88e1816b9be91671e07ba63c2e64d712425e4a7eec510e70f653efd17f9d@18.188.121.119:39797",
	// us-node2.backbone2.energi.network
	"enode://20b5679ca81eb3ac65369c639246c6029d0cfde9e88471349a5250067eae6617183b86890237271150688d725475c4cc17eb594bfae89da8fcc51d3e9eb7189a@18.216.223.87:39797",
	// us-node3.backbone2.energi.network
	"enode://a63c81108e1a5a45476f5fed6fadc0c6f8a22d51d50dccd0b7ae2218001b74de6336893d6a0fe9b6469a89601ecd44760e8a9f4b6e596229e2798b5c521e5784@18.223.42.176:39797",
	// us-node4.backbone2.energi.network
	"enode://5b1928c026ed84f5b075a3387d4c1dfd0cab56608223d1b6cdb56831c90a28a1c8374b1ffbcd2b2859ae740e342426077fa0c33ef8785674e10bc1e2829985fe@18.223.72.151:39797",
	// us-node5.backbone2.energi.network
	"enode://527524b27571d9deedd859c3355ad292e9744481ac7499ffa24e0e853966ea5f727199e84b7371f99d1949a21486d6629ef7cd406ec3f924eba3bbf84bd4c423@18.223.88.189:39797",
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
