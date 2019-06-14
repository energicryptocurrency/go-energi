'use strict';

const BackboneRewardV1 = artifacts.require('BackboneRewardV1');
const BlacklistRegistryV1 = artifacts.require('BlacklistRegistryV1');
const CheckpointRegistryV1 = artifacts.require('CheckpointRegistryV1');
const Gen2Migration = artifacts.require('Gen2Migration');
//const GenericProposalV1 = artifacts.require('GenericProposalV1');
const GovernedProxy = artifacts.require('GovernedProxy');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const SporkRegistryV1 = artifacts.require('SporkRegistryV1');
const StakerRewardV1 = artifacts.require('StakerRewardV1');
const TreasuryV1 = artifacts.require('TreasuryV1');

const MockProxy = artifacts.require("MockProxy");
const MockContract = artifacts.require("MockContract");
const MockSporkRegistry = artifacts.require("MockSporkRegistry");
const MockProposal = artifacts.require("MockProposal");

const common = require('../test/common');

module.exports = async (deployer, _, accounts) => {
    try {
        // V1 instances
        await deployer.deploy(MockProxy);
        const mn_token_proxy = MockProxy.address;
        await deployer.deploy(MockProxy);
        const treasury_proxy = MockProxy.address;
        await deployer.deploy(MockProxy);
        const mn_registry_proxy = MockProxy.address;

        const deploy_common = async (type, proxy, ...args) => {
            if (!proxy) {
                await deployer.deploy(MockProxy);
                proxy = MockProxy.address;
            }

            const instance = await deployer.deploy(type, proxy, ...args);
            await (await MockProxy.at(proxy)).setImpl(instance.address);
        };

        await deployer.deploy(Gen2Migration, treasury_proxy, common.chain_id);

        await deploy_common(BlacklistRegistryV1, undefined, mn_registry_proxy);
        await deploy_common(BackboneRewardV1, undefined, accounts[5]);
        await deploy_common(CheckpointRegistryV1);
        await deploy_common(MasternodeTokenV1, mn_token_proxy, mn_registry_proxy);
        await deploy_common(MasternodeRegistryV1,
            mn_registry_proxy, mn_token_proxy, treasury_proxy,
            common.mnregistry_config);
        await deploy_common(SporkRegistryV1, undefined, mn_registry_proxy);
        await deploy_common(StakerRewardV1);
        await deploy_common(TreasuryV1, treasury_proxy, mn_registry_proxy, common.superblock_cycles);

        // For unit test
        await deployer.deploy(GovernedProxy, BackboneRewardV1.address, SporkRegistryV1.address);
        await deployer.deploy(MockProxy);
        await deployer.deploy(MockContract, MockProxy.address);
        await deployer.deploy(MockSporkRegistry, MockProxy.address);
        await deployer.deploy(MockProposal);
        //await deployer.deploy(GenericProposalV1, mn_registry_proxy, 1, 1, mn_registry_proxy)
    } catch (e) {
        // eslint-disable-next-line no-console
        console.dir(e);
        throw e;
    }
};
