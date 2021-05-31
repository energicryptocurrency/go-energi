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

// Energi Governance system is the fundamental part of Energi Core.

'use strict';

const MockProxy = artifacts.require('MockProxy');
const MockContract = artifacts.require('MockContract');
const MockProposal = artifacts.require('MockProposal');
const MockSporkRegistry = artifacts.require('MockSporkRegistry')
// const MockMasternodeTokenV2 = artifacts.require('MockMasternodeTokenV2')
const MasternodeRegistryV2_1 = artifacts.require('MasternodeRegistryV2_1');
const MasternodeRegistryV2 = artifacts.require('MasternodeRegistryV2');
const IBlockReward = artifacts.require('IBlockReward');
const IMasternodeRegistryV2 = artifacts.require('IMasternodeRegistryV2');
const IMasternodeToken = artifacts.require('IMasternodeToken');
const ITreasury = artifacts.require('ITreasury');
const StorageMasternodeRegistryV1 = artifacts.require('StorageMasternodeRegistryV1');

const common = require('./common');

contract("MasternodeRegistryV2_1", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
        storage: null,
    };

    const { toWei } = web3.utils;
    // const vperiod = common.mnregistry_config_v2[1];
    // const isTargetChanges = async (_token, _mn) => {
    //     return await _token.validationTarget(_mn, 'latest') !=  await _token.validationTarget(_mn, 'pending');
    // };
    // const sw_features = web3.utils.toBN((1 << 24) | (2 << 16) | (3 << 8));

    before(async () => {
        s.orig = await MasternodeRegistryV2_1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        await s.proxy.setImpl(s.orig.address);

        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await MasternodeRegistryV2_1.at(s.proxy.address);

        s.token_abi = await IMasternodeRegistryV2.at(s.proxy.address);
        s.reward_abi = await IBlockReward.at(s.proxy.address);
        s.storage = await StorageMasternodeRegistryV1.at(await s.proxy_abi.v1storage());

        s.mntoken_proxy_addr = await s.orig.token_proxy();
        //s.mntoken_proxy = await MockProxy.at(s.mntoken_proxy_addr);
        s.mntoken_abi = await IMasternodeToken.at(s.mntoken_proxy_addr);

        s.treasury_proxy_addr = await s.orig.treasury_proxy();
        s.treasury_proxy = await MockProxy.at(s.treasury_proxy_addr);
        s.treasury_abi = await ITreasury.at(s.treasury_proxy_addr);
        s.treasury_impl = await ITreasury.at(await s.treasury_proxy.impl());

        Object.freeze(s);
    });

    after(async () => {
        const impl = await MasternodeRegistryV2_1.new(
            s.proxy.address,
            s.mntoken_proxy_addr,
            s.treasury_proxy_addr,
            common.mnregistry_config_v2,
            common.mnreg_deploy_opts
        );
        await s.proxy.setImpl(impl.address);
    });

    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('Primary', () => {
        const { fromAscii, toBN } = web3.utils;

        const collateral1 = toWei('3000', 'ether');
        const collateral2 = toWei('2000', 'ether');
        const collateral3 = toWei('1000', 'ether');
        const reward = toBN(toWei('9.14', 'ether'));

        const owner1 = accounts[1];
        const owner2 = accounts[2];
        const owner3 = accounts[3];
        const owner4 = accounts[10];
        const not_owner = accounts[0];

        const masternode1 = accounts[9];
        const masternode2 = accounts[8];
        const masternode3 = accounts[7];

        const ip1 = toBN(0x12345678);
        const ip2 = toBN(0x87654321);
        const ip3 = toBN(0x43218765);

        const enode_common = '123456789012345678901234567890'
        const enode1 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '11')];
        const enode2 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '22')];
        const enode3 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '33')];

        before(async () => {
            await s.mntoken_abi.depositCollateral({
                from: owner1,
                value: collateral1,
            });
            await s.mntoken_abi.depositCollateral({
                from: owner2,
                value: collateral2,
            });
            await s.mntoken_abi.depositCollateral({
                from: owner3,
                value: collateral3,
            });
        });

        after(async () => {
            await s.mntoken_abi.withdrawCollateral(collateral1, {
                from: owner1,
            });
            await s.mntoken_abi.withdrawCollateral(collateral2, {
                from: owner2,
            });
            await s.mntoken_abi.withdrawCollateral(collateral3, {
                from: owner3,
            });
        });

        describe('Three MN', () => {
            const nodes = [
                {
                    masternode: masternode1,
                    ip: ip1,
                    enode: enode1,
                    owner: owner1,
                    collateral: collateral1,
                },
                {
                    masternode: masternode2,
                    ip: ip2,
                    enode: enode2,
                    owner: owner2,
                    collateral: collateral2,
                },
                {
                    masternode: masternode3,
                    ip: ip3,
                    enode: enode3,
                    owner: owner3,
                    collateral: collateral3,
                }
            ];

            it('should announce()', async () => {
                for (let mn of nodes) {
                    await s.token_abi.announce(
                        mn.masternode, mn.ip, mn.enode, { from: mn.owner });
                }

                const mn1_status = await s.orig.mn_status(masternode1);
                const mn2_status = await s.orig.mn_status(masternode2);
                const mn3_status = await s.orig.mn_status(masternode3);
                expect(mn1_status.seq_payouts.toString()).equal('3');
                expect(mn2_status.seq_payouts.toString()).equal('2');
                expect(mn3_status.seq_payouts.toString()).equal('1');
            });

            it('should re-announce MN', async () => {
                // back order to test current being left in place first
                for (let mn of Array.from(nodes).reverse()) {
                    await s.token_abi.announce(
                        mn.masternode, mn.ip, mn.enode, { from: mn.owner });

                    const ann_e = await s.orig.getPastEvents('Announced', common.evt_last_block);
                    expect(ann_e).lengthOf(1);
                    common.stringifyBN(web3, ann_e[0].args);
                    expect(ann_e[0].args).deep.include({
                        '0': mn.masternode,
                        '1': mn.owner,
                        '2': toBN(mn.ip).toString(),
                        '3': mn.enode,
                        '4': toBN(mn.collateral).toString(),
                        '__length__': 5,
                        'masternode': mn.masternode,
                        'owner': mn.owner,
                        'ipv4address': toBN(mn.ip).toString(),
                        'enode': mn.enode,
                        'collateral': toBN(mn.collateral).toString(),
                    });

                    const den_e = await s.orig.getPastEvents('Denounced', common.evt_last_block);
                    expect(den_e).lengthOf(1);
                    common.stringifyBN(web3, den_e[0].args);
                    expect(den_e[0].args).deep.include({
                        '0': mn.masternode,
                        '1': mn.owner,
                        '__length__': 2,
                        'masternode': mn.masternode,
                        'owner': mn.owner,
                    });
                }
            });

            it('should be isActive()', async () => {
                for (let mn of nodes) {
                    expect(await s.token_abi.isActive(mn.masternode)).true;
                    expect(await s.token_abi.isActive(mn.owner)).false;
                }
            });

            it('should correctly count', async () => {
                const res = await s.token_abi.count();
                common.stringifyBN(web3, res);
                expect(res).eql({
                    '0': '3',
                    '1': '3',
                    '2': toWei('6000', 'ether'),
                    '3': toWei('6000', 'ether'),
                    '4': toWei('10000', 'ether'),
                    'active': '3',
                    'total': '3',
                    'active_collateral': toWei('6000', 'ether'),
                    'total_collateral': toWei('6000', 'ether'),
                    'max_of_all_times': toWei('10000', 'ether'),
                });
            });

            it('should produce info()', async () => {
                for (let mn of nodes) {
                    const info = await s.token_abi.info(mn.masternode);
                    common.stringifyBN(web3, info);
                    expect(info).deep.include({
                        owner: mn.owner,
                        ipv4address: toBN(mn.ip).toString(),
                        enode: mn.enode,
                        collateral: toBN(mn.collateral).toString(),
                        sw_features: '0',
                    });
                }
            });

            it('should process reward() deactivate missing heartbeat', async () => {
                await common.moveTime(web3, 40*60);

                {
                    const b = await web3.eth.getBlock('latest');
                    await s.token_abi.heartbeat(b.number, b.hash, '12', {from:masternode1, ...common.zerofee_callopts});
                    await s.token_abi.heartbeat(b.number, b.hash, '23', {from:masternode2, ...common.zerofee_callopts});
                    await s.token_abi.heartbeat(b.number, b.hash, '34', {from:masternode3, ...common.zerofee_callopts});
                }

                await common.moveTime(web3, 70*60);

                const treasury_before = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                const owner1_before = toBN(await web3.eth.getBalance(owner1));
                const owner2_before = toBN(await web3.eth.getBalance(owner2));
                const owner3_before = toBN(await web3.eth.getBalance(owner3));
                const count = 18;
                let sb = false;

                for (let i = count; i > 0; --i) {
                    if (i == 12 || i == 8 || i == 4) {
                        const bn = await web3.eth.getBlockNumber();
                        const b = await web3.eth.getBlock(bn);
                        await s.token_abi.heartbeat(bn, b.hash, '12', {from:masternode1, ...common.zerofee_callopts});
                        await s.token_abi.heartbeat(bn, b.hash, '34', {from:masternode3, ...common.zerofee_callopts});
                        await common.moveTime(web3, 91*60);
                    }

                    let r = await s.reward_abi.getReward(i);
                    if (r.eq(toBN(0))) {
                        // superblock case
                        r = await s.reward_abi.getReward(i+1);
                        sb = true;
                    }

                    expect(r.toString()).eql(reward.toString());

                    await s.reward_abi.reward({
                        from: not_owner,
                        value: r,
                    });
                }

                expect(sb).true;

                const treasury_after = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                const owner1_after = toBN(await web3.eth.getBalance(owner1));
                const owner2_after = toBN(await web3.eth.getBalance(owner2));
                const owner3_after = toBN(await web3.eth.getBalance(owner3));

                // The treasury must get reward of nodes without votes by design
                expect(treasury_after.sub(treasury_before).toString())
                    .eql(reward.mul(toBN(0)).toString())
                expect(owner1_after.sub(owner1_before).toString())
                    .eql(reward.mul(toBN(3+3+3+3)).toString());
                expect(owner2_after.sub(owner2_before).toString())
                    .eql(reward.mul(toBN(2+0+0+0)).toString());
                expect(owner3_after.sub(owner3_before).toString())
                    .eql(reward.mul(toBN(1+1+1+1)).toString());

                expect(await s.token_abi.isActive(masternode1)).true;
                expect(await s.token_abi.isActive(masternode2)).false;
                expect(await s.token_abi.isActive(masternode3)).true;
            });

            it('should skip inactive node from validation', async () => {
                const target1 = await s.token_abi.validationTarget(masternode1);
                const target3 = await s.token_abi.validationTarget(masternode3);
                expect(target1).eql(masternode3);
                expect(target3).eql(masternode1);
            });

            it('should canInvalidate()', async () => {
                expect(await s.token_abi.canInvalidate(masternode1)).true;
                expect(await s.token_abi.canInvalidate(masternode2)).false;
                expect(await s.token_abi.canInvalidate(masternode3)).true;
            });

            it('should handle enumerate()', async () => {
                expect(await s.token_abi.enumerate()).members([
                    masternode1, masternode2, masternode3]);
            });

            it('should handle enumerateActive()', async () => {
                expect(await s.token_abi.enumerateActive()).members([
                    masternode1, masternode3]);
            });

            it('should denounce() on collateral withdrawal', async() => {
                await s.token_abi.announce(
                    masternode1, ip1, enode1, { from: owner1 });

                await s.mntoken_abi.withdrawCollateral(collateral1, {from: owner1});

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(1);
                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(0);
                expect(await s.token_abi.enumerate()).members([masternode2, masternode3]);
            });

            it('should re-announce() on collateral change', async() => {
                // Initial
                await s.mntoken_abi.depositCollateral({
                    from: owner1,
                    value: collateral1,
                });
                await s.token_abi.announce(
                    masternode1, ip1, enode1, { from: owner1 });

                // Change +
                await s.mntoken_abi.depositCollateral({
                    from: owner1,
                    value: collateral1,
                });

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(1);
                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(1);
                expect(await s.token_abi.enumerate()).members([masternode1, masternode2, masternode3]);

                // Change -
                await s.mntoken_abi.withdrawCollateral(collateral1, {from: owner1});

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(1);
                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(1);
                expect(await s.token_abi.enumerate()).members([masternode1, masternode2, masternode3]);
            });

            it('should cleanup inactive node', async () => {
                await common.moveTime(web3, 25*60*60);

                for (let i = 4; i > 0; --i) {
                    await s.reward_abi.reward({
                        from: not_owner,
                        value: reward
                    });
                }

                expect(await s.token_abi.enumerate()).members([masternode1, masternode3]);
            });

            it('should denounce()', async()=> {
                for (let mn of nodes) {
                    await s.token_abi.denounce(mn.masternode, { from: mn.owner });
                    const evt = await s.orig.getPastEvents('Denounced', common.evt_last_block);

                    if (mn.masternode == masternode2) {
                        expect(evt).lengthOf(0);
                        continue;
                    }

                    expect(evt).lengthOf(1);
                    expect(evt[0].args).deep.include({
                        '0': mn.masternode,
                        '1': mn.owner,
                        '__length__': 2,
                        'masternode': mn.masternode,
                        'owner': mn.owner,
                    });
                }
            });

            it('should correctly count() ever max', async () => {
                const res = await s.token_abi.count();
                common.stringifyBN(web3, res);
                expect(res).eql({
                    '0': '0',
                    '1': '0',
                    '2': toWei('0', 'ether'),
                    '3': toWei('0', 'ether'),
                    '4': toWei('10000', 'ether'),
                    'active': '0',
                    'total': '0',
                    'active_collateral': toWei('0', 'ether'),
                    'total_collateral': toWei('0', 'ether'),
                    'max_of_all_times': toWei('10000', 'ether'),
                });
            });

            it('should migrateStatusPartial() reject self migration', async () => {
                const mn_proxy = await MockProxy.new();
                const impl = await MasternodeRegistryV2_1.new(
                    mn_proxy.address,
                    s.mntoken_proxy_addr,
                    s.treasury_proxy_addr,
                    common.mnregistry_config_v2,
                    common.mnreg_deploy_opts,
                );
                await mn_proxy.setImpl(impl.address)

                try {
                    await impl.migrateStatusPartial({ from: owner4 });
                    assert.fail('It should fail');
                } catch(e) {
                    assert.match(e.message, /cannot migrate from self/);
                }
            });

            it('should migrate from V2 to V2_1', async () => {
                const collateral = toWei('10000', 'ether');

                // Spork registry
                const registry_proxy = await MockProxy.new();
                const registry = await MockSporkRegistry.new(registry_proxy.address);
                await registry_proxy.setImpl(registry.address);

                // MNReg proxy
                const mn_proxy = await MockProxy.new();
                const imn = await IMasternodeRegistryV2.at(mn_proxy.address);

                const impl1 = await MasternodeRegistryV2.new(
                    mn_proxy.address,
                    s.mntoken_proxy_addr,
                    s.treasury_proxy_addr,
                    common.mnregistry_config_v2,
                    common.mnreg_deploy_opts
                );
                const impl2 = await MasternodeRegistryV2_1.new(
                    mn_proxy.address,
                    s.mntoken_proxy_addr,
                    s.treasury_proxy_addr,
                    common.mnregistry_config_v2,
                    common.mnreg_deploy_opts,
                );
                await mn_proxy.setImpl(impl1.address);

                // Announce
                for (let mn of nodes) {
                    if (mn.masternode === masternode3) {
                        continue;
                    }

                    await s.mntoken_abi.depositCollateral({
                        from: mn.owner,
                        value: collateral,
                    });

                    await imn.announce(mn.masternode, mn.ip, mn.enode, { from: mn.owner });
                }

                expect(await imn.enumerate()).members([masternode1, masternode2]);

                expect(await impl2.enumerate()).members([]);

                // Expects no failure the first time its run.
                try {
                    await impl2.migrateStatusPartial({ from: owner4 });
                } catch(e) {
                    assert.fail('It should fail');
                }

                // Upgrade
                const { logs } = await mn_proxy.proposeUpgrade(impl2.address, 0);
                s.assert.equal(logs.length, 1);

                const proposal = await MockProposal.at(logs[0].args['1']);

                await proposal.setAccepted();
                await mn_proxy.upgrade(proposal.address);

                // Ensure MNs are still there
                expect(await imn.enumerate()).members([masternode1, masternode2]);

                expect(await impl2.enumerate()).members([masternode1, masternode2]);

                // Second migration should fail since it has already completed.
                try {
                    await impl2.migrateStatusPartial({ from: owner4 });
                    assert.fail('It should fail');
                } catch(e) {
                    assert.match(e.message, /migration already done/);
                }
            });
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});
