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

const MasternodeRegistryV2_2 = artifacts.require('MasternodeRegistryV2_2');

const common = require('../test/common');

module.exports = function(deployer, network) {
    try {
        const mn_registry_proxy = '0x0000000000000000000000000000000000000302';
        const mn_token_proxy = '0x0000000000000000000000000000000000000309';
        const treasury_proxy = '0x0000000000000000000000000000000000000301';
        var mnregistry_config_v2_2 = [];

        console.log("Deploying to " + network);

        if (network === "mainnet") {
            mnregistry_config_v2_2 = [
                10,                         // RequireValidation
                5,                          // ValidationPeriods
                24*60*60,                   // CleanupPeriod
                '1000000000000000000000',   // minimum collateral 1000 NRG
                10                          // MNRewardsPerBlock
            ];
        } else if (network === "testnet") {
            mnregistry_config_v2_2 = [
                5,                          // RequireValidation
                5,                          // ValidationPeriods
                3*60*60,                    // CleanupPeriod
                '1000000000000000000000',   // minimum collateral 1000 NRG
                10                          // MNRewardsPerBlock
            ];
        } else {
            mnregistry_config_v2_2 = common.mnregistry_config_v2;
        }

        deployer.deploy(MasternodeRegistryV2_2,
            mn_registry_proxy,
            mn_token_proxy,
            treasury_proxy,
            mnregistry_config_v2_2,
            common.mnreg_deploy_opts
        );
    } catch (e) {
        console.dir(e);
        throw e;
    }
};
