// Copyright 2021 The Energi Core Authors
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

const MasternodeRegistryV2_3 = artifacts.require('MasternodeRegistryV2_3');
const common = require('../test/common');

module.exports = async function(deployer, network, accounts) {
    try {
        const mn_registry_proxy = '0x0000000000000000000000000000000000000302';
        const mn_token_proxy = '0x0000000000000000000000000000000000000309';
        const treasury_proxy = '0x0000000000000000000000000000000000000301';
        const migrate_gas_limit = 30000000;         // gas limit for each call to migrateStatusPartial()
        const migrate_gas_price = 100000000000;     // gas price for each call to migrateStatusPartial()
        const migrate_max_calls = 10;               // how many times we are allowed to call migrateStatusPartial()
        var mnregistry_config_v2_3 = [];

        console.log("Deploying to " + network);

        if (network === "mainnet") {
            mnregistry_config_v2_3 = [
                10,                         // RequireValidation
                5,                          // ValidationPeriods
                24*60*60,                   // CleanupPeriod
                '1000000000000000000000',   // minimum collateral 1000 NRG
                10                          // MNRewardsPerBlock
            ];
        } else if (network === "testnet") {
            mnregistry_config_v2_3 = [
                5,                          // RequireValidation
                5,                          // ValidationPeriods
                3*60*60,                    // CleanupPeriod
                '1000000000000000000000',   // minimum collateral 1000 NRG
                10                          // MNRewardsPerBlock
            ];
        } else {
            mnregistry_config_v2_3 = common.mnregistry_config_v2;
        }

        await deployer.deploy(MasternodeRegistryV2_3,
            mn_registry_proxy,
            mn_token_proxy,
            treasury_proxy,
            mnregistry_config_v2_3,
            common.mnreg_deploy_opts
        );

        // perform data migration from previous masternode registry
        if ((network === "mainnet") || (network === "testnet")) {
            var instance = await MasternodeRegistryV2_3.deployed();
            for (var i = 0; i < migrate_max_calls; i++) {
                console.log('   Calling function \'migrateStatusPartial()\'');
                console.log('   ---------------------------------------------');
                try {
                    const ret = await instance.migrateStatusPartial({from: accounts[0], gas: migrate_gas_limit, gasPrice: migrate_gas_price});
                    // TODO: get BigNumber working correctly to be able to call fromWei
                    //const bal = BigNumber(web3.eth.getBalance(accounts[0]));
                    //const total_cost = BigNumber(ret.receipt.gasUsed * migrate_gas_price);
                    console.log("   > transaction hash:\t" + ret.tx);
                    console.log("   > block number:\t\t" + ret.receipt.blockNumber);
                    console.log("   > account:\t\t\t" + accounts[0]);
                    //console.log("   > balance:\t\t\t" + web3.utils.fromWei(bal));
                    console.log("   > gas used:\t\t\t" + ret.receipt.gasUsed);
                    console.log("   > gas price:\t\t\t" + (migrate_gas_price / 10e8) + " gwei");
                    //console.log("   > total cost:\t\t\t" + web3.utils.fromWei(total_cost) + " NRG");
                    console.log();
                } catch (e) {
                    // expected revert
                    if ((e.name === 'StatusError') && (e.reason === 'migration already done')) {
                        console.log("   > migrateStatusPartial(): complete");
                        console.log();
                        break;
                    } else {
                        console.dir(e);
                        throw e;
                    }
                }
            }
        }
    } catch (e) {
        console.dir(e);
        throw e;
    }
};
