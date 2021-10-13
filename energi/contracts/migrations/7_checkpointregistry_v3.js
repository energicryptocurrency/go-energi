// Copyright 2020 The Energi Core Authors
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

'use strict';

const CheckpointRegistryV3 = artifacts.require('CheckpointRegistryV3');
const common = require('../test/common');

module.exports = async function(deployer, network) {
    try {
        var checkpointSigner = common.cpp_signer;
        var checkpointProxyAddress = '0x0000000000000000000000000000000000000306';
        var masternodeRegistryProxyAddress = '0x0000000000000000000000000000000000000302';

        console.log("Deploying to " + network);

        if (network === "mainnet") {
            checkpointSigner = '0xBD1C57eACcfD1519E342F870C1c551983F839479';
        } else if (network === "testnet") {
            checkpointSigner = '0xb1372ea07f6a92bc86fd5f8cdf468528f79f87ca';
        } else {
            // CheckpointRegistryV3 doesn't need to be deployed here for tests
            // it will be deployed by the CheckpointRegistryV3.spec.js test file
            return
        }

        // since this uses GovernedContractAutoProxy, make sure we capture the new proxy address
        await deployer.deploy(CheckpointRegistryV3, checkpointProxyAddress, masternodeRegistryProxyAddress, checkpointSigner);
        var instance = await CheckpointRegistryV3.deployed();
        var proxyAddress = await instance.proxy();
        console.log("   > proxy address:       " + proxyAddress);
    } catch (e) {
        console.dir(e);
        throw e;
    }
};
