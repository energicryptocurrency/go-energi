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

const GovernedProxy = artifacts.require('GovernedProxy');
const HardforkRegistryV1 = artifacts.require('HardforkRegistryV1');

module.exports = function(deployer, network) {
    try {
        var hf_signer;
        const hf_finalized_period = 30;
        const sporky_address =  '0x0000000000000000000000000000000000000305';
        const hf_proxy_address = '0x000000000000000000000000000000000000030D';

        if (network === "mainnet") {
            hf_signer = '0x44D16E845ec2d2D6A99a10fe44EE99DA0541CF31';
        } else if (network === "testnet") {
            hf_signer = '0x5b00118464fa6e73f9c2a4ea44e1cbfa9f5b83c6';
        } else {
            // Not supported for any other network
            return;
        }

        console.log("Deploying to " + network);

        deployer.deploy(GovernedProxy,
            hf_proxy_address,
            sporky_address
        )

        deployer.deploy(HardforkRegistryV1,
            hf_proxy_address, 
            hf_signer,
            hf_finalized_period,
            {overwrite: false}
        );
    } catch (e) {
        console.dir(e);
        throw e;
    }
};
