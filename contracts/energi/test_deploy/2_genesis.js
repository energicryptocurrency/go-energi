'use strict';

const BackboneRewardV1 = artifacts.require('BackboneRewardV1');
const BlacklistRegistryV1 = artifacts.require('BlacklistRegistryV1');
const CheckpointRegistryV1 = artifacts.require('CheckpointRegistryV1');
const Gen2Migration = artifacts.require('Gen2Migration');
const GenericProposal = artifacts.require('GenericProposal');
const GovernedProxy = artifacts.require('GovernedProxy');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const SporkRegistryV1 = artifacts.require('SporkRegistryV1');
const StakerRewardV1 = artifacts.require('StakerRewardV1');
const TreasuryV1 = artifacts.require('TreasuryV1');

module.exports = function(deployer) {
    // V1 instances
    deployer.deploy(BackboneRewardV1);
    deployer.deploy(BlacklistRegistryV1);
    deployer.deploy(CheckpointRegistryV1);
    deployer.deploy(Gen2Migration);
    deployer.deploy(MasternodeTokenV1);
    deployer.deploy(MasternodeRegistryV1);
    deployer.deploy(SporkRegistryV1);
    deployer.deploy(StakerRewardV1);
    deployer.deploy(TreasuryV1).then(() => {
        // This instances a put here just for clarity, but
        // tests will use a new instance due to Truffle limitations.
        GovernedProxy.new(BackboneRewardV1.address);
        GovernedProxy.new(BlacklistRegistryV1.address);
        GovernedProxy.new(CheckpointRegistryV1.address);
        GovernedProxy.new(MasternodeTokenV1.address);
        GovernedProxy.new(MasternodeRegistryV1.address);
        GovernedProxy.new(SporkRegistryV1.address);
        GovernedProxy.new(StakerRewardV1.address);
        GovernedProxy.new(TreasuryV1.address);
    });
};
