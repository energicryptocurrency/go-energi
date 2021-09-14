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

const TreasuryV2 = artifacts.require('TreasuryV2');
const MockProxy = artifacts.require("MockProxy");
const common = require('../test/common');

module.exports = async function(deployer, network) {
    try {
        const mn_registry_proxy = '0x0000000000000000000000000000000000000302';
        const treasury_proxy = '0x0000000000000000000000000000000000000301';
        console.log("Deploying to " + network);

        if (network === "mainnet" || network === "testnet") {
            deployer.deploy(TreasuryV2,
                treasury_proxy,
                mn_registry_proxy,
                common.superblock_cycles,
                common.treasury_deploy_opts
            );
        } else {
            await deployer.deploy(MockProxy);
            const treasury_mock_proxy = MockProxy.address

            deployer.deploy(TreasuryV2,
                treasury_mock_proxy,
                mn_registry_proxy,
                common.superblock_cycles,
                common.treasury_deploy_opts
            );
        }
    } catch (e) {
        console.dir(e);
        throw e;
    }
};
