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

const GovernedProxy = artifacts.require('GovernedProxy');
const MockContract = artifacts.require('MockContract');
const MockSporkRegistry = artifacts.require('MockSporkRegistry');
const MockProposal = artifacts.require('MockProposal');

contract("GovernedProxy", async accounts => {
    let first;
    let second;
    let third;
    let fourth;
    let proxy;
    let proxy_abi;
    let registry;
    const weeks = 60*60*24*7;

    before(async () => {
        registry = await MockSporkRegistry.deployed();
        first = await MockContract.new(registry.address);
        proxy = await GovernedProxy.new(first.address, registry.address);
        second = await MockContract.new(proxy.address);
        third = await MockContract.new(proxy.address);
        fourth = await MockContract.new(proxy.address);
        proxy_abi = await MockContract.at(proxy.address);
    });
    it('should refuse migrate()', async () => {
        try {
            await proxy.migrate(second.address, { from: accounts[0] });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Good try/);
        }
    });

    it('should refuse destroy()', async () => {
        try {
            await proxy.destroy(second.address, { from: accounts[0] });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Good try/);
        }
    });


    it('should proxy calls', async () => {
        const res = await proxy_abi.getAddress({ from: accounts[0] });
        assert.equal(first.address.valueOf(), res.valueOf());
    });

    it('should refuse proposal - same impl', async () => {
        try {
            await proxy.proposeUpgrade(
                    first.address, 2 * weeks,
                    { from: accounts[0], value: web3.utils.toWei('1', 'ether') });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Already active!/);
        }
    });

    it('should refuse proposal - wrong proxy', async () => {
        try {
            await proxy.proposeUpgrade(
                    registry.address, 2 * weeks,
                    { from: accounts[0], value: web3.utils.toWei('1', 'ether') });
            assert.fail("It must fail");
        } catch (e) {
            //assert.match(e.message, /Wrong proxy!/);
            assert.match(e.message, /revert/);
        }
    });

    it('should accept proposal', async () => {
        await proxy.proposeUpgrade(
                second.address, 2 * weeks,
                // NOTE: it's mock registry - no fee check
                { from: accounts[0], value: '1' });
    });

    it('should refuse upgrade - Not accepted!', async () => {
        const { logs } = await proxy.proposeUpgrade(
                second.address, 2 * weeks,
                { from: accounts[0], value: '1' });

        assert.equal(logs.length, 1);
        const proposal = logs[0].args['1'];

        try {
            await proxy.upgrade(proposal);
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not accepted!/);
        }
    });

    it('should refuse upgrade - Not registered!', async () => {
        let proposal = await MockProposal.new();

        try {
            await proxy.upgrade(proposal.address);
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not registered!/);
        }
    });

    it('should accept upgrade', async () => {
        const { logs } = await proxy.proposeUpgrade(
                second.address, 2 * weeks,
                { from: accounts[0], value: '1' });
        assert.equal(logs.length, 1);
        const proposal = await MockProposal.at(logs[0].args['1']);

        await proposal.setAccepted();

        const res = await proxy.upgrade(proposal.address);
        assert.equal(res.logs.length, 1);
    });

    it('should refuse upgrade AFTER upgrade - Not registered!', async () => {
        const { logs } = await proxy.proposeUpgrade(
                third.address, 2 * weeks,
                { from: accounts[0], value: '1' });
        const proposal = await MockProposal.at(logs[0].args['1']);

        await proposal.setAccepted();
        await proxy.upgrade(proposal.address);

        try {
            await proxy.upgrade(proposal.address);
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not registered!/);
        }
    });

    it('should refuse upgrade - Already active!', async () => {
        let proposal1 = await proxy.proposeUpgrade(
                fourth.address, 2 * weeks,
                { from: accounts[0], value: '1' });
        let proposal2 = await proxy.proposeUpgrade(
                fourth.address, 2 * weeks,
                { from: accounts[0], value: '1' });
        proposal1 = await MockProposal.at(proposal1.logs[0].args['1']);
        proposal2 = await MockProposal.at(proposal2.logs[0].args['1']);
        await proposal1.setAccepted();
        await proposal2.setAccepted();
        await proxy.upgrade(proposal1.address);

        try {
            await proxy.upgrade(proposal2.address);
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Already active!/);
        }
    });
});
