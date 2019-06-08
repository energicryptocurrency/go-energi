'use strict';

const BackboneRewardV1 = artifacts.require('BackboneRewardV1');
const BlacklistRegistryV1 = artifacts.require('BlacklistRegistryV1');
const CheckpointRegistryV1 = artifacts.require('CheckpointRegistryV1');
const Gen2Migration = artifacts.require('Gen2Migration');
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

module.exports = async (deployer) => {
    // V1 instances
    await deployer.deploy(MockProxy);
    await deployer.deploy(BackboneRewardV1, MockProxy.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(BlacklistRegistryV1, MockProxy.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(CheckpointRegistryV1, MockProxy.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(Gen2Migration, MockProxy.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(MasternodeTokenV1, MockProxy.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(MasternodeRegistryV1, MockProxy.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(SporkRegistryV1, MockProxy.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(StakerRewardV1, MockProxy.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(TreasuryV1, MockProxy.address);

    // For unit test
    await deployer.deploy(GovernedProxy, BackboneRewardV1.address, SporkRegistryV1.address);
    await deployer.deploy(MockProxy);
    await deployer.deploy(MockContract, MockProxy.address);
    await deployer.deploy(MockSporkRegistry, MockProxy.address);
    await deployer.deploy(MockProposal);
};
