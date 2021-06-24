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

const MasternodeRegistryV2_1 = artifacts.require('MasternodeRegistryV2_1');
const MockProxy = artifacts.require("MockProxy");
const common = require('../test/common');

module.exports = async function(deployer, network) {
    try {
        const mn_registry_proxy = '0x0000000000000000000000000000000000000302';
        const mn_token_proxy = '0x0000000000000000000000000000000000000309';
        const treasury_proxy = '0x0000000000000000000000000000000000000301';
        let mnregistry_config_v2_1 = [];

        console.log("Deploying to " + network);

        if (network === "mainnet") {
            mnregistry_config_v2_1 = [
                10,                         // RequireValidation
                5,                          // ValidationPeriods
                24*60*60,                   // CleanupPeriod
                '1000000000000000000000',   // minimum collateral 1000 NRG
                10                          // MNRewardsPerBlock
            ];


            deployer.deploy(MasternodeRegistryV2_1,
                mn_registry_proxy,
                mn_token_proxy,
                treasury_proxy,
                mnregistry_config_v2_1,
                common.mnreg_deploy_opts
            );
        } else if (network === "testnet") {
            mnregistry_config_v2_1 = [
                5,                          // RequireValidation
                5,                          // ValidationPeriods
                3*60*60,                    // CleanupPeriod
                '1000000000000000000000',   // minimum collateral 1000 NRG
                10                          // MNRewardsPerBlock
            ];


            deployer.deploy(MasternodeRegistryV2_1,
                mn_registry_proxy,
                mn_token_proxy,
                treasury_proxy,
                mnregistry_config_v2_1,
                common.mnreg_deploy_opts
            );
        } else {
            mnregistry_config_v2_1 = common.mnregistry_config_v2;
            await deployer.deploy(MockProxy);
            const mn_registry_mock_proxy = MockProxy.address;
            await deployer.deploy(MockProxy);
            const mn_token_mock_proxy = MockProxy.address;
            await deployer.deploy(MockProxy);
            const treasury_mock_proxy = MockProxy.address


            deployer.deploy(MasternodeRegistryV2_1,
                mn_registry_mock_proxy,
                mn_token_mock_proxy,
                treasury_mock_proxy,
                mnregistry_config_v2_1,
                common.mnreg_deploy_opts
            );
        }
    } catch (e) {
        console.dir(e);
        throw e;
    }
};
