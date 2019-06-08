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
const MockProposal = artifacts.require('MockProposal');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const IMasternodeRegistry = artifacts.require('IMasternodeRegistry');
const StorageMasternodeRegistryV1 = artifacts.require('StorageMasternodeRegistryV1');

contract("MasternodeRegistryV1", async accounts => {
    let orig;
    let fake;
    let proxy;
    let proxy_abi;
    let token_abi;
    let storage;

    before(async () => {
        orig = await MasternodeRegistryV1.deployed();
        proxy = await MockProxy.at(await orig.proxy());
        fake = await MockContract.new(proxy.address);
        proxy_abi = await MasternodeRegistryV1.at(proxy.address);
        token_abi = await IMasternodeRegistry.at(proxy.address);
        await proxy.setImpl(orig.address);
        storage = await StorageMasternodeRegistryV1.at(await proxy_abi.v1storage());
    });

    it('should refuse migrate() through proxy', async () => {
        try {
            await proxy_abi.migrate(fake.address, { from: accounts[0] });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Good try/);
        }
    });

    it('should refuse destroy() through proxy', async () => {
        try {
            await proxy_abi.destroy(fake.address, { from: accounts[0] });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Good try/);
        }
    });

    it('should refuse migrate() directly', async () => {
        try {
            await orig.migrate(fake.address, { from: accounts[0] });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not proxy/);
        }
    });

    it('should refuse destroy() directly', async () => {
        try {
            await orig.destroy(fake.address, { from: accounts[0] });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not proxy/);
        }
    });


    // Primary stuff
    //---


    // Safety & Cleanup
    //---
    it('should refuse to accept funds', async () => {
        try {
            await token_abi.send(web3.utils.toWei('1', "ether"));
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not supported/);
        }
    });
    
    it('should refuse to accept funds to storage', async () => {
        try {
            await storage.send(web3.utils.toWei('1', "ether"));
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /revert/);
        }
    });

    it('should refuse kill() storage', async () => {
        try {
            await storage.kill();
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not owner/);
        }
    });

    it('should refuse setOwner() on storage', async () => {
        try {
            await storage.setOwner(proxy.address);
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not owner/);
        }
    });

    it('should destroy() after upgrade', async () => {
        const orig_balance = await web3.eth.getBalance(orig.address)
        const { logs } = await proxy.proposeUpgrade(
                fake.address, 0,
                { from: accounts[0], value: '1' });

        assert.equal(logs.length, 1);
        const proposal = await MockProposal.at(logs[0].args['1']);
        
        await proposal.setAccepted();
        await proxy.upgrade(proposal.address);

        const fake_balance = await web3.eth.getBalance(fake.address)
        assert.equal(orig_balance.valueOf(), fake_balance.valueOf());

        try {
            await orig.proxy();
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /did it run Out of Gas/);
        }
    });

    it('should transfer storage & allow to kill() it', async () => {
        await fake.killStorage(storage.address);
    });
});
