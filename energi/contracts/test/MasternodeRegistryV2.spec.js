// Copyright 2019-2020 The Energi Core Authors
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
const MockSporkRegistry = artifacts.require('MockSporkRegistry')
const MockMasternodeTokenV2 = artifacts.require('MockMasternodeTokenV2')
const MasternodeRegistryV2 = artifacts.require('MasternodeRegistryV2');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const IBlockReward = artifacts.require('IBlockReward');
const IMasternodeRegistryV2 = artifacts.require('IMasternodeRegistryV2');
const IMasternodeToken = artifacts.require('IMasternodeToken');
const ITreasury = artifacts.require('ITreasury');
const StorageMasternodeRegistryV1 = artifacts.require('StorageMasternodeRegistryV1');

const common = require('./common');

contract("MasternodeRegistryV2", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
        storage: null,
    };

    const { toWei } = web3.utils;
    const vperiod = common.mnregistry_config_v2[1];
    const isTargetChanges = async (_token, _mn) => {
        return await _token.validationTarget(_mn, 'latest') !=  await _token.validationTarget(_mn, 'pending');
    };
    
    before(async () => {
        s.orig = await MasternodeRegistryV2.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        await s.proxy.setImpl(s.orig.address);

        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await MasternodeRegistryV2.at(s.proxy.address);

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
        const impl = await MasternodeRegistryV2.new(
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

        const checkHeartbeat = async () => {
            try {
                const evt = await s.orig.getPastEvents('Heartbeat', common.evt_last_block);
                expect(evt).lengthOf(0);
            } catch(e) {
                expect(e.message).match(/doesn't exist in this contract/);
            }
        };

        describe('No MN', () => {
            it('should silently denounce()', async () => {
                await s.token_abi.denounce(masternode1);
                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(0);
            });

            it('should refuse to heartbeat() too old block', async () => {
                const bn = await web3.eth.getBlockNumber();
                const b = await web3.eth.getBlock(bn);

                try {
                    await s.token_abi.heartbeat(bn - 11, b.hash, '0', common.zerofee_callopts);
                    assert.fail('It should fail');
                } catch(e) {
                    assert.match(e.message, /Too old/);
                }
            });

            it('should refuse to heartbeat() wrong block', async () => {
                const bn = (await web3.eth.getBlockNumber());
                const b = await web3.eth.getBlock(bn);

                try {
                    await s.token_abi.heartbeat(bn - 10, b.hash, '0', common.zerofee_callopts);
                    assert.fail('It should fail');
                } catch(e) {
                    assert.match(e.message, /Block mismatch/);
                }
            });

            it('should refuse to heartbeat() not active', async () => {
                const bn = await web3.eth.getBlockNumber();
                const b = await web3.eth.getBlock(bn);

                try {
                    await s.token_abi.heartbeat(bn, b.hash, '0', common.zerofee_callopts);
                    assert.fail('It should fail');
                } catch(e) {
                    assert.match(e.message, /Not active/);
                }
            });

            it('should refuse to invalidate() vote for self', async () => {
                try {
                    await s.token_abi.invalidate(
                        owner1, {from: owner1, ...common.zerofee_callopts});
                    assert.fail('It should fail');
                } catch(e) {
                    assert.match(e.message, /Invalidation for self/);
                }
            });

            it('should refuse to invalidate() not active', async () => {
                try {
                    await s.token_abi.invalidate(masternode2, common.zerofee_callopts);
                    assert.fail('It should fail');
                } catch(e) {
                    assert.match(e.message, /Not active caller/);
                }
            });

            it('should not be isActive()', async () => {
                const res = await s.token_abi.isActive(masternode1);
                expect(res).false;
            });

            it('should correctly count()', async () => {
                const res = await s.token_abi.count();
                assert.equal(res[0], 0);
                assert.equal(res[1], 0);
                assert.equal(res[2], 0);
            });

            it('should handle info()', async () => {
                try {
                    await s.token_abi.info(masternode1);
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Unknown masternode/);
                }
            });

            it('should handle ownerInfo()', async () => {
                try {
                    await s.token_abi.ownerInfo(owner1);
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Unknown owner/);
                }
            });

            it('should process reward() to Treasury', async () => {
                const treasury_before = toBN(await web3.eth.getBalance(s.treasury_impl.address));

                let r = await s.reward_abi.getReward(0);
                assert.equal(r.valueOf(), 0);

                const count = 3;
                
                for (let i = count; i > 0; --i) {
                    r = await s.reward_abi.getReward(i);

                    if (r.eq(toBN(0))) {
                        // superblock case
                        r = await s.reward_abi.getReward(i+1);
                    }

                    expect(r.toString()).eql(reward.toString());
                    await s.reward_abi.reward({
                        from: not_owner,
                        value: r
                    });
                }

                // Kick the rest
                await s.reward_abi.reward();

                const treasury_after = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                expect(treasury_after.sub(treasury_before).toString())
                    .eql(reward.mul(toBN(count)).toString());
            });

            it('should handle onCollateralUpdate()', async () => {
                await s.token_abi.onCollateralUpdate(owner1);
            });

            it('collateralLimits() should return correct values', async () => {
                let limits = await s.token_abi.collateralLimits();
                common.stringifyBN(web3, limits);
                expect(limits).include({
                    min: '1000000000000000000000', // 1000 NRG
                    max: '100000000000000000000000', // 100,000 NRG
                });
            });

            it('should handle enumerate()', async () => {
                expect(await s.token_abi.enumerate()).lengthOf(0);
            });

            it.skip('must forbid more than one reward() per block', async () => {
                // Bug: https://github.com/trufflesuite/truffle/issues/1389
                const batch = web3.eth.BatchRequest();
                batch.add(s.reward_abi.reward.request({value: reward}));
                batch.add(s.reward_abi.reward.request({value: reward}));

                try {
                    await batch.execute();
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Call outside of governance/);
                }
            });
        });

        describe('Single MN', () => {
            let announced_block;

            it('should refuse announce() without collateral', async () => {
                try {
                    await s.token_abi.announce(
                        masternode1, ip1, enode1, { from: not_owner });
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Invalid collateral/);
                }

                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(0);
            });

            const non_routables = {
                '127.0.0.0/8' : [ 0x7F000001, 0x7FFFFFFF ],
                '10.0.0.0/8' : [ 0x0A000001, 0x0AFFFFFF ],
                '172.16.0.0/12' : [ 0xAC100001, 0xAC108001 ],
                '192.168.0.0/16' : [ 0xC0A80001, 0xC0A88001 ],
                '0.0.0.0/8' : [ 0x00123456 ],
                '100.64.0.0/10' : [ 0x64400001, 0x64480001 ],
                '169.254.0.0/16' : [ 0xA9FE0001, 0xA9FEFFFF ],
                '198.18.0.0/15' : [ 0xC6120001, 0xC613FFFF ],
                '198.51.100.0/24' : [ 0xC6336401, 0xC63364FF ],
                '203.0.113.0/24' : [ 0xCB007101, 0xCB0071FE ],
                '224.0.0.0/4' : [ 0xE0000001, 0xE80FF001 ],
                '240.0.0.0/4' : [ 0xF0000001, 0xF800FFFF ],
                '255.255.255.255/32' : [ 0xFFFFFFFF ],
            };

            for (let k in non_routables) {
                it(`should refuse announce() non-routable IPs: ${k}`, async () => {
                    for (let ip of non_routables[k]) {
                        try {
                            await s.token_abi.announce(
                                masternode1, ip, enode1, { from: owner1 });
                            assert.fail('It should fail');
                        } catch (e) {
                            assert.match(e.message, /Wrong IP/);
                        }
                    }
                });
            }

            it('should announce()', async () => {
                const res = await s.mntoken_abi.balanceInfo(owner1);
                assert.equal(res['0'].valueOf(), collateral1);

                await s.token_abi.announce(
                    masternode1, ip1, enode1, { from: owner1 });

                const ann_e = await s.orig.getPastEvents('Announced', common.evt_last_block);
                expect(ann_e).lengthOf(1);
                common.stringifyBN(web3, ann_e[0].args);
                expect(ann_e[0].args).deep.include({
                    '0': masternode1,
                    '1': owner1,
                    '2': ip1.toString(),
                    '3': enode1,
                    '4': toBN(collateral1).toString(),
                    '__length__': 5,
                    'masternode': masternode1,
                    'owner': owner1,
                    'ipv4address': ip1.toString(),
                    'enode': enode1,
                    'collateral': toBN(collateral1).toString(),
                });

                const den_e = await s.orig.getPastEvents('Denounced', common.evt_last_block);
                expect(den_e).lengthOf(0);
            });

            it('should re-announce MN', async () => {
                await s.token_abi.announce(
                    masternode1, ip1, enode1, { from: owner1 });

                const ann_e = await s.orig.getPastEvents('Announced', common.evt_last_block);
                expect(ann_e).lengthOf(1);
                expect(ann_e[0].args).deep.include({
                    '0': masternode1,
                    '1': owner1,
                    '2': ip1,
                    '3': enode1,
                    '4': toBN(collateral1),
                    '__length__': 5,
                    'masternode': masternode1,
                    'owner': owner1,
                    'ipv4address': ip1,
                    'enode': enode1,
                    'collateral': toBN(collateral1),
                });

                const den_e = await s.orig.getPastEvents('Denounced', common.evt_last_block);
                expect(den_e).lengthOf(1);
                expect(den_e[0].args).deep.include({
                    '0': masternode1,
                    '1': owner1,
                    '__length__': 2,
                    'masternode': masternode1,
                    'owner': owner1,
                });

                announced_block = await web3.eth.getBlockNumber();
            });

            it('should refuse announce() another owner\'s MN', async () => {
                try {
                    await s.token_abi.announce(
                        masternode1, ip2, enode2, { from: owner2 });
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Invalid owner/);
                }

                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(0);
            });

            it('should refuse denounce() another owner\'s MN', async () => {
                try {
                    await s.token_abi.denounce(masternode1, { from: owner2 });
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Invalid owner/);
                }

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(0);
            });

            it('should forbid heartbeat() too early', async () => {
                const bn = await web3.eth.getBlockNumber();
                const b = await web3.eth.getBlock(bn);

                try {
                    await s.token_abi.heartbeat(bn, b.hash, '0', {from: masternode1, ...common.zerofee_callopts});
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Too early/);
                }

                await common.moveTime(web3, 59*30);

                try {
                    await s.token_abi.heartbeat(bn, b.hash, '0', {from: masternode1, ...common.zerofee_callopts});
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Too early/);
                }
            });

            it('should be isActive()', async () => {
                expect(await s.token_abi.isActive(masternode1)).true;
                expect(await s.token_abi.isActive(masternode2)).false;
            });

            it('should canHeartbeat()', async () => {
                expect(await s.token_abi.canHeartbeat(masternode1)).false;
                expect(await s.token_abi.canHeartbeat(masternode2)).false;

                await common.moveTime(web3, 60*30+1);

                expect(await s.token_abi.canHeartbeat(masternode1)).true;
                expect(await s.token_abi.canHeartbeat(masternode2)).false;
            });

            it('should heartbeat()', async () => {
                const s1 = await s.orig.mn_status(masternode1);
                const bn = await web3.eth.getBlockNumber();
                const b = await web3.eth.getBlock(bn);

                await s.token_abi.heartbeat(bn, b.hash, '0', {from: masternode1, ...common.zerofee_callopts});

                const s2 = await s.orig.mn_status(masternode1);
                expect(s2.next_heartbeat.gt(s1.next_heartbeat)).true;
                expect(s2.next_heartbeat.gt(b.timestamp)).true;

                await checkHeartbeat();
            });

            it('should correctly count', async () => {
                const res = await s.token_abi.count();
                assert.equal(res[0], 1);
                assert.equal(res[1], 1);
                assert.equal(res[2].toString(), collateral1.toString());
            });

            it('should produce info()', async () => {
                const info = await s.token_abi.info(masternode1);
                common.stringifyBN(web3, info);
                expect(info).deep.include({
                    owner: owner1,
                    ipv4address: toBN(ip1).toString(),
                    enode: enode1,
                    collateral: toBN(collateral1).toString(),
                });
            });

            it('should produce ownerInfo()', async () => {
                const info = await s.token_abi.ownerInfo(owner1);
                common.stringifyBN(web3, info);
                expect(info).deep.include({
                    masternode: masternode1,
                    ipv4address: toBN(ip1).toString(),
                    enode: enode1,
                    collateral: toBN(collateral1).toString(),
                    announced_block: announced_block.toString(),
                });
            });

            it('should process reward()', async () => {
                const treasury_before = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                const owner_before = toBN(await web3.eth.getBalance(owner1));
                const count = 3;

                for (let i = count; i > 0; --i) {
                    const r = await s.reward_abi.getReward(i);
                    expect(r).eql(reward);
                    await s.reward_abi.reward({
                        from: owner2,
                        value: r
                    });
                }

                const treasury_after = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                expect(treasury_before.toString()).equal(treasury_after.toString());

                const owner_after = toBN(await web3.eth.getBalance(owner1));
                expect(owner_after.sub(owner_before).toString())
                    .eql(reward.mul(toBN(count)).toString());
            });

            it('should handle onCollateralUpdate()', async () => {
                await s.token_abi.onCollateralUpdate(owner1);
                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(0);
            });

            it('should handle enumerate()', async () => {
                expect(await s.token_abi.enumerate()).members([masternode1]);
            });

            it('should forbid heartbeat() too late', async () => {
                const bn = await web3.eth.getBlockNumber();
                const b = await web3.eth.getBlock(bn);

                expect(await s.token_abi.canHeartbeat(masternode1)).false;

                try {
                    await s.token_abi.heartbeat(bn, b.hash, '0', {from: masternode1, ...common.zerofee_callopts});
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Too early/);
                }

                const mn1_status = await s.orig.mn_status(masternode1);
                let to_move = (mn1_status.next_heartbeat.add(toBN(2*60*60+1)));
                to_move = to_move.sub(toBN((await web3.eth.getBlock('latest')).timestamp));

                await common.moveTime(web3, to_move.toNumber());

                expect(await s.token_abi.canHeartbeat(masternode1)).false;

                try {
                    await s.token_abi.heartbeat(bn, b.hash, '0', {from: masternode1, ...common.zerofee_callopts});
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Not active/);
                }

                // Denounce does not happen on read-only
                expect(await s.orig.getPastEvents(
                    'Denounced', common.evt_last_block)).lengthOf(0);
            });

            it('should denounce() on collateral withdrawal', async() => {
                await s.token_abi.announce(
                    masternode1, ip1, enode1, { from: owner1 });

                await s.mntoken_abi.withdrawCollateral(collateral1, {from: owner1});

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(1);
                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(0);
                expect(await s.token_abi.enumerate()).members([]);
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
                expect(await s.token_abi.enumerate()).members([masternode1]);

                // Change -
                await s.mntoken_abi.withdrawCollateral(collateral1, {from: owner1});

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(1);
                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(1);
                expect(await s.token_abi.enumerate()).members([masternode1]);
            });

            it('should denounce()', async()=> {
                await s.token_abi.denounce(masternode1, { from: owner1 });
                const evt = await s.orig.getPastEvents('Denounced', common.evt_last_block);
                expect(evt).lengthOf(1);
                expect(evt[0].args).deep.include({
                    '0': masternode1,
                    '1': owner1,
                    '__length__': 2,
                    'masternode': masternode1,
                    'owner': owner1,
                });
            });
        });

        describe('Two MN', () => {
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
            ];

            it('should announce()', async () => {
                for (let mn of nodes) {
                    await s.token_abi.announce(
                        mn.masternode, mn.ip, mn.enode, { from: mn.owner });
                }

                const mn1_status = await s.orig.mn_status(masternode1);
                const mn2_status = await s.orig.mn_status(masternode2);
                expect(mn1_status.seq_payouts.toString()).equal('3');
                expect(mn2_status.seq_payouts.toString()).equal('2');
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

            it('should refuse announce() another owner\'s MN', async () => {
                try {
                    await s.token_abi.announce(
                        masternode1, ip2, enode2, { from: owner2 });
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Invalid owner/);
                }

                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(0);
            });

            it('should refuse denounce() another owner\'s MN', async () => {
                try {
                    await s.token_abi.denounce(masternode1, { from: owner2 });
                    assert.fail('It should fail');
                } catch (e) {
                    assert.match(e.message, /Invalid owner/);
                }

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(0);
            });

            it('should be isActive()', async () => {
                for (let mn of nodes) {
                    expect(await s.token_abi.isActive(mn.masternode)).true;
                }
            });

            it('should heartbeat()', async () => {
                await common.moveTime(web3, 60*30+1);

                const s1 = await s.orig.mn_status(masternode1);
                const s1o = await s.orig.mn_status(masternode2);
                const bn = await web3.eth.getBlockNumber();
                const b = await web3.eth.getBlock(bn);

                await s.token_abi.heartbeat(bn, b.hash, '0', {from: masternode1, ...common.zerofee_callopts});
                
                const s2 = await s.orig.mn_status(masternode1);
                expect(s2.next_heartbeat.gt(s1.next_heartbeat)).true;
                expect(s2.next_heartbeat.gt(b.timestamp)).true;

                const s2o = await s.orig.mn_status(masternode2);
                expect(s2o.next_heartbeat.eq(s1o.next_heartbeat)).true;
                
                await checkHeartbeat();
            });

            it('should correctly count', async () => {
                const res = await s.token_abi.count();
                common.stringifyBN(web3, res);
                expect(res).eql({
                    '0': '2',
                    '1': '2',
                    '2': toWei('5000', 'ether'),
                    '3': toWei('5000', 'ether'),
                    '4': toWei('10000', 'ether'),
                    'active': '2',
                    'total': '2',
                    'active_collateral': toWei('5000', 'ether'),
                    'total_collateral': toWei('5000', 'ether'),
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
                        collateral: toBN(mn.collateral).toString()
                    });
                }
            });

            it('should process reward()', async () => {
                const treasury_before = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                const owner1_before = toBN(await web3.eth.getBalance(owner1));
                const owner2_before = toBN(await web3.eth.getBalance(owner2));
                const count = 10;
                let sb = false;

                for (let i = count; i > 0; --i) {
                    let r = await s.reward_abi.getReward(i);
                    if (r.eq(toBN(0))) {
                        // superblock case
                        r = await s.reward_abi.getReward(i+1);
                        sb = true;
                    }

                    expect(r.toString()).eql(reward.toString());

                    await s.reward_abi.reward({
                        from: owner3,
                        value: r
                    });
                }

                expect(sb).true;

                const treasury_after = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                expect(treasury_before.toString()).equal(treasury_after.toString());

                const owner1_after = toBN(await web3.eth.getBalance(owner1));
                const owner2_after = toBN(await web3.eth.getBalance(owner2));
                expect(owner1_after.sub(owner1_before).toString())
                    .eql(reward.mul(toBN(6)).toString());
                expect(owner2_after.sub(owner2_before).toString())
                    .eql(reward.mul(toBN(4)).toString());
            });

            it('should handle enumerate()', async () => {
                expect(await s.token_abi.enumerate()).members([masternode1, masternode2]);
            });

            it('should denounce() on collateral withdrawal', async() => {
                await s.token_abi.announce(
                    masternode1, ip1, enode1, { from: owner1 });

                await s.mntoken_abi.withdrawCollateral(collateral1, {from: owner1});

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(1);
                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(0);
                expect(await s.token_abi.enumerate()).members([masternode2]);
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
                expect(await s.token_abi.enumerate()).members([masternode1, masternode2]);

                // Change -
                await s.mntoken_abi.withdrawCollateral(collateral1, {from: owner1});

                expect(await s.orig.getPastEvents('Denounced', common.evt_last_block)).lengthOf(1);
                expect(await s.orig.getPastEvents('Announced', common.evt_last_block)).lengthOf(1);
                expect(await s.token_abi.enumerate()).members([masternode1, masternode2]);
            });

            it('should denounce()', async()=> {
                for (let mn of nodes) {
                    await s.token_abi.denounce(mn.masternode, { from: mn.owner });
                    const evt = await s.orig.getPastEvents('Denounced', common.evt_last_block);
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
                },
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
                        collateral: toBN(mn.collateral).toString()
                    });
                }
            });

            it('should process reward()', async () => {
                const treasury_before = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                const owner1_before = toBN(await web3.eth.getBalance(owner1));
                const owner2_before = toBN(await web3.eth.getBalance(owner2));
                const owner3_before = toBN(await web3.eth.getBalance(owner3));
                const count = 18;
                let sb = false;
                
                for (let i = count; i > 0; --i) {
                    let r = await s.reward_abi.getReward(i);
                    if (r.eq(toBN(0))) {
                        // superblock case
                        r = await s.reward_abi.getReward(i+1);
                        sb = true;
                    }

                    expect(r.toString()).eql(reward.toString());

                    await s.reward_abi.reward({
                        from: not_owner,
                        value: r
                    });

                    const target1 = await s.token_abi.validationTarget(masternode1);
                    const target2 = await s.token_abi.validationTarget(masternode2);
                    const target3 = await s.token_abi.validationTarget(masternode3);
                    expect(target1).not.equal(masternode1);
                    expect(target2).not.equal(masternode2);
                    expect(target3).not.equal(masternode3);
                    expect(target1).not.equal(target2);
                    expect(target1).not.equal(target3);

                    if (i > 6 && i < 12 || i == 13) {
                        let t = target1;

                        if (await isTargetChanges(s.token_abi, masternode1)) {
                            // still a chance of fail.
                            t = (t === masternode3) ? masternode2 : masternode3;
                        }

                        const invalidator = (t === masternode3) ? masternode1 : masternode2;

                        if (await s.token_abi.canInvalidate(invalidator)) {
                            await s.token_abi.invalidate(masternode3, {from: invalidator});
                        }
                    }
                }

                expect(sb).true;

                // One invalidation at round 2
                const treasury_after = toBN(await web3.eth.getBalance(s.treasury_impl.address));
                expect(treasury_after.sub(treasury_before).toString())
                    .eql(reward.mul(toBN(1)).toString())

                const owner1_after = toBN(await web3.eth.getBalance(owner1));
                const owner2_after = toBN(await web3.eth.getBalance(owner2));
                const owner3_after = toBN(await web3.eth.getBalance(owner3));
                expect(owner1_after.sub(owner1_before).toString())
                    .eql(reward.mul(toBN(3+3+3)).toString());
                expect(owner2_after.sub(owner2_before).toString())
                    .eql(reward.mul(toBN(2+2+2)).toString());
                expect(owner3_after.sub(owner3_before).toString())
                    .eql(reward.mul(toBN(1+0+1)).toString());
            });


            it('should calculate validation target by periods', async () => {
                let bn = await web3.eth.getBlockNumber();
                const valTarget = () => {
                    return s.token_abi.methods['validationTarget(address)'](masternode1, bn);
                }

                let tmp = await valTarget();
                let target = tmp;

                do {
                    --bn;
                    target = await valTarget();
                } while ( target === tmp );

                const obn = bn;

                for (let i = vperiod; i > 0; --i, --bn) {
                    tmp = await valTarget();
                    expect(tmp).equal(target);
                }

                bn = obn + 1;

                for (let i = vperiod; i > 0; --i, ++bn) {
                    tmp = await valTarget();
                    expect(tmp).not.equal(target);
                }
            });

            it('should refuse invalidate() wrong target', async () => {
                try {
                    let target = await s.token_abi.validationTarget(masternode1);

                    if ((target == masternode2) && !await isTargetChanges(s.token_abi, masternode1)) {
                        target = masternode3;
                    } else {
                        target = masternode2;
                    }
                    
                    await s.token_abi.invalidate(target, {from:masternode1, ...common.zerofee_callopts});
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Invalid target/);
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

            it('should refuse invalidate() by inactive node', async () => {
                try {
                    await s.token_abi.invalidate(masternode1, {from:masternode2, ...common.zerofee_callopts});
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Not active caller/);
                }
            });

            it('should refuse double invalidate()', async () => {
                expect(await s.token_abi.canInvalidate(masternode1)).true;
                await s.token_abi.invalidate(masternode3, {from:masternode1, ...common.zerofee_callopts});
                expect(await s.token_abi.canInvalidate(masternode1)).false;

                try {
                    for (let i = 0; i < vperiod; ++i) {
                        await s.token_abi.invalidate(masternode3, {from:masternode1, ...common.zerofee_callopts});
                    }
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Already invalidated/);
                }
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

            it('should migrate from V1', async () => {
                const collateral = toWei('10000', 'ether');

                // Spork registry
                const registry_proxy = await MockProxy.new();
                const registry = await MockSporkRegistry.new(registry_proxy.address);
                await registry_proxy.setImpl(registry.address);

                // MNReg proxy
                const mn_proxy = await MockProxy.new();
                const imn = await IMasternodeRegistryV2.at(mn_proxy.address);

                const impl1 = await MasternodeRegistryV1.new(
                    mn_proxy.address,
                    s.mntoken_proxy_addr,
                    s.treasury_proxy_addr,
                    common.mnregistry_config
                );
                const impl2 = await MasternodeRegistryV2.new(
                    mn_proxy.address,
                    s.mntoken_proxy_addr,
                    s.treasury_proxy_addr,
                    common.mnregistry_config_v2,
                    common.mnreg_deploy_opts
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

                // Upgrade
                const { logs } = await mn_proxy.proposeUpgrade(impl2.address, 0);
                s.assert.equal(logs.length, 1);

                const proposal = await MockProposal.at(logs[0].args['1']);

                await proposal.setAccepted();
                await mn_proxy.upgrade(proposal.address);

                // Ensure MNs are still there
                expect(await imn.enumerate()).members([masternode1, masternode2]);

                expect(await impl2.enumerate()).members([masternode1, masternode2]);
            });
        });

        const SPI_MN_COUNT = process.env.SPI_MN_COUNT || 10; // set to 10000 at runtime
        describe(`${SPI_MN_COUNT} MNs`, () => {
            const nodes = [];

            let mntoken_orig;
            let mntoken_proxy;
            let mnreg_orig;
            let mnreg_proxy;
            let mnreg_abi;

            before(async () => {
                mntoken_proxy = await MockProxy.new();
                mnreg_proxy = await MockProxy.new();

                mntoken_orig = await MockMasternodeTokenV2.new(
                    mntoken_proxy.address, mnreg_proxy.address,
                    common.mnreg_deploy_opts);
                await mntoken_proxy.setImpl(mntoken_orig.address);

                mnreg_orig = await MasternodeRegistryV2.new(
                    mnreg_proxy.address,
                    mntoken_proxy.address,
                    s.treasury_proxy_addr,
                    common.mnregistry_config_v2,
                    common.mnreg_deploy_opts
                );
                await mnreg_proxy.setImpl(mnreg_orig.address);
                mnreg_abi = await IMasternodeRegistryV2.at(mnreg_orig.address);
            });

            it('should generate metadata', async function() {
                this.timeout(7200e3);
                const c1k = toBN(toWei('1000', 'ether'));
                const for_fees = toBN(toWei('1', 'ether'));
                const e1 = fromAscii('12345678901234567890123456789012');

                for (let i = 1; i <= SPI_MN_COUNT; ++i) {
                    const e2 = (i.toString() + '000000').repeat(10).substr(0, 32);
                    const masternode = await web3.eth.personal.newAccount('');
                    const owner = await web3.eth.personal.newAccount('');
                    const collateral = c1k.mul(toBN(1 + i%100));

                    await web3.eth.personal.unlockAccount(masternode, '', 0);
                    await web3.eth.personal.unlockAccount(owner, '', 0);

                    nodes.push({
                        masternode,
                        owner,
                        ip: ip1,
                        enode: [e1, fromAscii(e2)],
                        collateral,
                    });

                    await web3.eth.sendTransaction({
                        to: owner,
                        from: owner1,
                        value: for_fees,
                    });
                    await web3.eth.sendTransaction({
                        to: masternode,
                        from: owner1,
                        value: for_fees,
                    });
                    await mntoken_orig.setBalance(owner, collateral);

                    if (i%100 === 0) {
                        // eslint-disable-next-line no-console
                        console.log(`Done ${i}`);
                    }
                }
            });

            it('should announce all', async function() {
                this.timeout(7200e3);
                let i = 0;

                for (let n of nodes) {
                    await mnreg_abi.announce(n.masternode, n.ip, n.enode, {
                        from: n.owner
                    });

                    if (++i%100 === 0) {
                        // eslint-disable-next-line no-console
                        console.log(`Done ${i}`);
                    }
                }
            });

            it('should heartbeat all', async function() {
                this.timeout(7200e3);
                let i = 0;

                await common.moveTime(web3, 60*30);

                for (let { masternode } of nodes) {
                    if (!await mnreg_abi.canHeartbeat(masternode)) {
                        const status = await mnreg_orig.mn_status(masternode);
                        let to_move = status.next_heartbeat.add(toBN(1));
                        to_move -= (await web3.eth.getBlock('latest')).timestamp;

                        await common.moveTime(web3, to_move.toNumber());
                    }

                    const bn = await web3.eth.getBlockNumber();
                    const b = await web3.eth.getBlock(bn);

                    await mnreg_abi.heartbeat(bn, b.hash, '0', {
                        from: masternode,
                        ...common.zerofee_callopts,
                    });

                    if (++i%100 === 0) {
                        // eslint-disable-next-line no-console
                        console.log(`Done ${i}`);
                    }
                }
            });
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});
