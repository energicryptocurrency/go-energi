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

'use strict';

const BackboneRewardV1 = artifacts.require('BackboneRewardV1');
const BlacklistRegistryV1 = artifacts.require('BlacklistRegistryV1');
const BlockRewardV1 = artifacts.require('BlockRewardV1');
const CheckpointRegistryV1 = artifacts.require('CheckpointRegistryV1');
const CheckpointRegistryV2 = artifacts.require('CheckpointRegistryV2');
const CheckpointRegistryV3 = artifacts.require('CheckpointRegistryV3');
const StorageCheckpointRegistryV2 = artifacts.require('StorageCheckpointRegistryV2');

const Gen2Migration = artifacts.require('Gen2Migration');
//const GenericProposalV1 = artifacts.require('GenericProposalV1');
const GovernedProxy = artifacts.require('GovernedProxy');
const HardforkRegistryV1 = artifacts.require('HardforkRegistryV1');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');
const MasternodeTokenV2 = artifacts.require('MasternodeTokenV2');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const MasternodeRegistryV2 = artifacts.require('MasternodeRegistryV2');
const MasternodeRegistryV2_1 = artifacts.require('MasternodeRegistryV2_1');
const SporkRegistryV1 = artifacts.require('SporkRegistryV1');
const SporkRegistryV2 = artifacts.require('SporkRegistryV2');
const StakerRewardV1 = artifacts.require('StakerRewardV1');
const TreasuryV1 = artifacts.require('TreasuryV1');
const TreasuryV2 = artifacts.require('TreasuryV2');


const MockProxy = artifacts.require("MockProxy");
//const MockAutoProxy = artifacts.require("MockAutoProxy");
const MockContract = artifacts.require("MockContract");
const MockSporkRegistry = artifacts.require("MockSporkRegistry");
const MockProposal = artifacts.require("MockProposal");

const common = require('../test/common');

module.exports = async (deployer, network, accounts) => {
    // mainnet and testnet don't do genesis deployment, they already have a genesis block
    if ((network === "mainnet") || (network === "testnet")) {
        console.log("Skipping genesis migration on " + network);
        return;
    }

    try {
        // V1 instances
        await deployer.deploy(MockProxy);
        const mn_token_proxy = MockProxy.address;
        await deployer.deploy(MockProxy);
        const treasury_proxy = MockProxy.address;
        await deployer.deploy(MockProxy);
        const blacklist_registry = MockProxy.address;
        await deployer.deploy(MockProxy);
        const mn_registry_proxy = MockProxy.address;
        await deployer.deploy(MockProxy);
        const staker_proxy = MockProxy.address;
        await deployer.deploy(MockProxy);
        const backbone_proxy = MockProxy.address;

        const deploy_common = async (type, proxy, ...args) => {
            if (!proxy) {
                await deployer.deploy(MockProxy);
                proxy = MockProxy.address;
            }

            const instance = await deployer.deploy(type, proxy, ...args);
            if (proxy !== common.default_address) {
                await (await MockProxy.at(proxy)).setImpl(instance.address);
            }
        };

        await deployer.deploy(Gen2Migration, blacklist_registry, common.chain_id, common.migration_signer);

        const compensation_fund = await TreasuryV1.new(treasury_proxy, mn_registry_proxy, 1);
        await deploy_common(BlacklistRegistryV1,
            blacklist_registry, mn_registry_proxy,
            Gen2Migration.address, compensation_fund.address,
            accounts[3],
            { gas: "10000000" });
        await deploy_common(BackboneRewardV1, backbone_proxy, accounts[5]);
        await deploy_common(CheckpointRegistryV1, undefined, mn_registry_proxy, common.cpp_signer);
        await deploy_common(CheckpointRegistryV2, undefined, mn_registry_proxy, common.cpp_signer);
        await deploy_common(CheckpointRegistryV3, undefined, mn_registry_proxy, common.cpp_signer);
        await deploy_common(StorageCheckpointRegistryV2, undefined, mn_registry_proxy, common.cpp_signer);
        await deploy_common(HardforkRegistryV1, common.default_address, common.hf_signer, common.hf_finalization_period);
        await deploy_common(MasternodeTokenV1, mn_token_proxy, mn_registry_proxy);
        await deploy_common(MasternodeTokenV2, mn_token_proxy, mn_registry_proxy);
        await deploy_common(MasternodeRegistryV1,
            mn_registry_proxy, mn_token_proxy, treasury_proxy,
            common.mnregistry_config);
        await deploy_common(MasternodeRegistryV2,
            mn_registry_proxy, mn_token_proxy, treasury_proxy,
            common.mnregistry_config_v2, common.mnreg_deploy_opts);
        await deploy_common(MasternodeRegistryV2_1,
            mn_registry_proxy, mn_token_proxy, treasury_proxy,
            common.mnregistry_config_v2, common.mnreg_deploy_opts);
        await deploy_common(SporkRegistryV1, undefined, mn_registry_proxy);
        await deploy_common(SporkRegistryV2, undefined, mn_registry_proxy, common.emergency_signer);
        await deploy_common(StakerRewardV1, staker_proxy);
        await deploy_common(TreasuryV1, treasury_proxy, mn_registry_proxy, common.superblock_cycles);
        await deploy_common(TreasuryV2, treasury_proxy, mn_registry_proxy, common.superblock_cycles);

        await deploy_common(BlockRewardV1, undefined, [
            staker_proxy,
            backbone_proxy,
            treasury_proxy,
            mn_registry_proxy,
        ]);

        // For unit test
        await deployer.deploy(GovernedProxy, BackboneRewardV1.address, SporkRegistryV1.address);
        await deployer.deploy(MockProxy);
        await deployer.deploy(MockContract, MockProxy.address);
        await deployer.deploy(MockSporkRegistry, MockProxy.address);
        await deployer.deploy(MockProposal, MockProxy.address, MockContract.address);
        //await deployer.deploy(GenericProposalV1, mn_registry_proxy, 1, 1, mn_registry_proxy)
    } catch (e) {
        // eslint-disable-next-line no-console
        console.dir(e);
        throw e;
    }
};
