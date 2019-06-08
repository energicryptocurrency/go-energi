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
const BlacklistRegistryV1 = artifacts.require('BlacklistRegistryV1');
const IBlacklistRegistry = artifacts.require('IBlacklistRegistry');

contract("BlacklistRegistryV1", async accounts => {
    let orig;
    let fake;
    let proxy;
    let proxy_abi;
    let token_abi;

    before(async () => {
        orig = await BlacklistRegistryV1.deployed();
        proxy = await MockProxy.at(await orig.proxy());
        fake = await MockContract.new(proxy.address);
        proxy_abi = await BlacklistRegistryV1.at(proxy.address);
        token_abi = await IBlacklistRegistry.at(proxy.address);
        await proxy.setImpl(orig.address);
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

    it('should destroy() after upgrade', async () => {
        const { logs } = await proxy.proposeUpgrade(
                fake.address, 0,
                { from: accounts[0], value: '1' });

        assert.equal(logs.length, 1);
        const proposal = await MockProposal.at(logs[0].args['1']);

        await proposal.setAccepted();
        await proxy.upgrade(proposal.address);

        try {
            await orig.proxy();
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /did it run Out of Gas/);
        }
    });
});
