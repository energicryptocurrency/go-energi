// Copyright 2019 The Energi Core Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

// Energi Governance system is the fundamental part of Energi Core.

'use strict';

const MockProxy = artifacts.require('MockProxy');
const MockContract = artifacts.require('MockContract');
const SporkRegistryV1 = artifacts.require('SporkRegistryV1');
const ISporkRegistry = artifacts.require('ISporkRegistry');

const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');

const common = require('./common');

contract("SporkRegistryV1", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
    };

    before(async () => {
        s.registry_orig = await MasternodeRegistryV1.deployed();
        s.registry = await MasternodeRegistryV1.at(await s.registry_orig.proxy());

        s.mntoken_orig = await MasternodeTokenV1.deployed();
        s.mntoken = await MasternodeTokenV1.at(await s.mntoken_orig.proxy());

        s.orig = await SporkRegistryV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.mnregistry_proxy = await MockProxy.at(await s.orig.mnregistry_proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await SporkRegistryV1.at(s.proxy.address);
        s.token_abi = await ISporkRegistry.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        Object.freeze(s);
    });

    after(async () => {
        const impl = await SporkRegistryV1.new(s.proxy.address, s.mnregistry_proxy.address);
        await s.proxy.setImpl(impl.address);
    });

    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('Primary', () => {
        const { fromAscii, toBN, toWei } = web3.utils;

        const fee_amount = toBN(toWei('10000', 'ether'));

        const collateral1 = toBN(toWei('50000', 'ether'));
        const owner1 = accounts[0];
        const masternode1 = accounts[9];
        const ip1 = toBN(0x12345678);
        const enode_common = '123456789012345678901234567890'
        const enode1 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '11')];

        before(async () => {
            await s.mntoken.depositCollateral({
                from: owner1,
                value: collateral1,
            });
            await s.registry.announce(masternode1, ip1, enode1, {from: owner1});
        });

        after(async () => {
            await s.mntoken.withdrawCollateral(collateral1, {
                from: owner1,
            });
        });
        
        it ('should refuse to createUpgradeProposal() with invalid fee', async () => {
            try {
                await s.token_abi.createUpgradeProposal(
                    s.fake.address, 14*24*60*60, accounts[0],
                    { value: fee_amount.add(toBN(1)) });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            await s.token_abi.createUpgradeProposal(
                s.fake.address, 14*24*60*60, accounts[0],
                { value: fee_amount });

            try {
                await s.token_abi.createUpgradeProposal(
                    s.fake.address, 14*24*60*60, accounts[0],
                    { value: fee_amount.sub(toBN(1)) });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }
        });

        it ('should refuse to createUpgradeProposal() with invalid period', async () => {
            try {
                await s.token_abi.createUpgradeProposal(
                    s.fake.address, 14*24*60*60-1, accounts[0],
                    { value: fee_amount });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Period min/);
            }

            await s.token_abi.createUpgradeProposal(
                s.fake.address, 14*24*60*60, accounts[0],
                { value: fee_amount });

            await s.token_abi.createUpgradeProposal(
                s.fake.address, 365*24*60*60, accounts[0],
                { value: fee_amount });

            try {
                await s.token_abi.createUpgradeProposal(
                    s.fake.address, 365*24*60*60+1, accounts[0],
                    { value: fee_amount });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Period max/);
            }
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});

