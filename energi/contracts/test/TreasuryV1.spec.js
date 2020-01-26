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

// Energi Governance system is the fundamental part of Energi Core.

'use strict';

const MockProxy = artifacts.require('MockProxy');
const MockContract = artifacts.require('MockContract');
const TreasuryV1 = artifacts.require('TreasuryV1');
const ITreasury = artifacts.require('ITreasury');
const IBlockReward = artifacts.require('IBlockReward');
const IProposal = artifacts.require('IProposal');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');
const StorageTreasuryV1 = artifacts.require('StorageTreasuryV1');

const common = require('./common');

contract("TreasuryV1", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
        storage: null,
    };

    before(async () => {
        s.registry_orig = await MasternodeRegistryV1.deployed();
        s.registry = await MasternodeRegistryV1.at(await s.registry_orig.proxy());

        s.mntoken_orig = await MasternodeTokenV1.deployed();
        s.mntoken = await MasternodeTokenV1.at(await s.mntoken_orig.proxy());

        s.orig = await TreasuryV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await TreasuryV1.at(s.proxy.address);
        s.token_abi = await ITreasury.at(s.proxy.address);
        s.treasury_abi = s.token_abi;
        s.reward_abi = await IBlockReward.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        s.storage = await StorageTreasuryV1.at(await s.proxy_abi.v1storage());
        Object.freeze(s);
    });

    after(async () => {
        const impl = await TreasuryV1.new(s.proxy.address, s.registry.address, common.superblock_cycles);
        await s.proxy.setImpl(impl.address);
    });

    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('ITreasury', () => {
        const { fromAscii, toBN, toWei } = web3.utils;
        const superblock_reward = toBN(toWei('184000', 'ether'));
        const payer1 = accounts[0];
        const def_period = 14*24*60*60; // 2 weeks
        const def_amount = toBN(toWei('1000', 'ether'));
        const fee_amount = toBN(toWei('100', 'ether'));

        const collateral1 = toBN(toWei('30000', 'ether'));
        const collateral2 = toBN(toWei('20000', 'ether'));

        const owner1 = accounts[0];
        const owner2 = accounts[1];
        const owner3 = accounts[3];
        const owner4 = accounts[4];

        const masternode1 = accounts[9];
        const masternode2 = accounts[8];

        const ip1 = toBN(0x12345678);
        const ip2 = toBN(0x87654321);

        const enode_common = '123456789012345678901234567890'
        const enode1 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '11')];
        const enode2 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '22')];

        before(async () => {
            await s.mntoken.depositCollateral({
                from: owner1,
                value: collateral1,
            });
            await s.mntoken.depositCollateral({
                from: owner2,
                value: collateral2,
            });
            await s.registry.announce(masternode1, ip1, enode1, {from: owner1});
            await s.registry.announce(masternode2, ip2, enode2, {from: owner2});
        });

        after(async () => {
            await s.mntoken.withdrawCollateral(collateral1, {
                from: owner1,
            });
            await s.mntoken.withdrawCollateral(collateral2, {
                from: owner2,
            });
        });
        
        it ('should correctly identify superblocks', async () => {
            const period = 8;
            expect(await s.treasury_abi.isSuperblock(0)).false;

            for (let i = 1; i < period; ++i) {
                expect(await s.treasury_abi.isSuperblock(i)).false;
            }

            expect(await s.treasury_abi.isSuperblock(period)).true;

            for (let i = period + 1; i < 2*period; ++i) {
                expect(await s.treasury_abi.isSuperblock(i)).false;
            }

            expect(await s.treasury_abi.isSuperblock(2*period)).true;
        });

        it ('should correctly identify reward', async () => {
            const period = 8;
            expect(await s.treasury_abi.isSuperblock(0)).false;

            for (let i = 1; i < period; ++i) {
                expect(await s.treasury_abi.isSuperblock(i)).false;
            }

            expect(await s.treasury_abi.isSuperblock(period)).true;

            for (let i = period + 1; i < 2*period; ++i) {
                expect(await s.treasury_abi.isSuperblock(i)).false;
            }

            expect(await s.treasury_abi.isSuperblock(2*period)).true;
        });
        
        it ('should correctly reflect balance()', async () => {
            await s.treasury_abi.contribute({value: '234'});

            const orig_bal = toBN(await web3.eth.getBalance(s.orig.address));
            expect(toBN(await s.treasury_abi.balance()).toString()).equal(orig_bal.toString());
        });

        it ('should handle contribute()', async () => {
            const orig_bal = toBN(await web3.eth.getBalance(s.orig.address));

            await s.treasury_abi.contribute({value: '123'});

            const bal_after = toBN(await s.treasury_abi.balance());
            expect(bal_after.sub(orig_bal).toString()).equal('123');

            const evt = await s.orig.getPastEvents('Contribution', common.evt_last_block);
            expect(evt).lengthOf(1);
            common.stringifyBN(web3, evt[0].args);
            expect(evt[0].args).deep.include({
                '0': payer1,
                '1': '123',
                '__length__': 2,
                'from': payer1,
                'amount': '123',
            });
        });

        it ('should refuse propose() without proper fee', async () => {
            try {
                await s.treasury_abi.propose(def_amount, '1', def_period);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            try {
                await s.treasury_abi.propose(
                    def_amount, '1', def_period,
                    { value: toWei('99.99', 'ether') }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            try {
                await s.treasury_abi.propose(
                    def_amount, '1', def_period,
                    { value: toWei('100.1', 'ether') }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }
        });

        it ('should refuse propose() without proper amount', async () => {
            try {
                await s.treasury_abi.propose(
                    toWei('99', 'ether'), '1', def_period,
                    { value: fee_amount }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Too small amount/);
            }

            try {
                await s.treasury_abi.propose(
                    toWei('184000.1', 'ether'), '1', def_period,
                    { value: fee_amount }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Too large amount/);
            }
        });

        it ('should refuse propose() without proper period', async () => {
            try {
                await s.treasury_abi.propose(
                    def_amount, '1', 7*24*60*60-1,
                    { value: fee_amount }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Too small period/);
            }

            try {
                await s.treasury_abi.propose(
                    def_amount, '1', 30*24*60*60+1,
                    { value: fee_amount }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Too large period/);
            }
        });

        it ('should propose()', async () => {
            const min_amount = toWei('100', 'ether');
            await s.treasury_abi.propose(
                min_amount, '11111111', def_period,
                { value: fee_amount }
            );

            const bn = await web3.eth.getBlockNumber();
            const b = await web3.eth.getBlock(bn);
            
            const evt = await s.orig.getPastEvents('BudgetProposal', common.evt_last_block);
            expect(evt).lengthOf(1);
            common.stringifyBN(web3, evt[0].args);
            expect(evt[0].args).deep.include({
                '0': '11111111',
                '2': payer1,
                '3': min_amount.toString(),
                '4': toBN(b.timestamp + def_period).toString(),
                '__length__': 5,
                'ref_uuid': '11111111',
                'payout_address': payer1,
                'amount': min_amount.toString(),
                'deadline': toBN(b.timestamp + def_period).toString(),
            });
            expect(evt[0].args).include.keys('proposal');
        });

        it ('should refuse propose() on UUID duplicate', async () => {
            try {
                await s.treasury_abi.propose(
                    def_amount, '11111111', def_period,
                    { value: fee_amount }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /UUID in use/);
            }
        });

        it ('should record & map proposals()', async () => {
            const max_amount = superblock_reward;
            await s.treasury_abi.propose(
                max_amount, '22222222', def_period,
                { value: fee_amount }
            );
            const proposal2 = await s.treasury_abi.uuid_proposal('22222222');
            expect(proposal2.toString()).not.equal(toBN('0').toString());
            
            const proposal1 = await s.treasury_abi.uuid_proposal('11111111');
            expect(proposal1.toString()).not.equal(toBN('0').toString());

            expect((await s.treasury_abi.proposal_uuid(proposal1)).toString()).equal('11111111');
            expect((await s.treasury_abi.proposal_uuid(proposal2)).toString()).equal('22222222');
        });

        it ('should refuse propose() over total limit', async () => {
            for (let i = 3; i <= 8; ++i) {
                await s.treasury_abi.propose(
                    def_amount, String(i).repeat(8), def_period + (i * def_period / 10),
                    { value: fee_amount }
                );
            }

            try {
                await s.treasury_abi.propose(
                    def_amount, '9'.repeat(8), def_period,
                    { value: fee_amount }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Too many active proposals/);
            }
        });

        it ('should collect rejected proposals', async () => {
            const orig_bal = toBN(await s.treasury_abi.balance());

            // No deadline
            await s.reward_abi.reward();
            expect(toBN(await s.treasury_abi.balance()).sub(orig_bal).toString())
                .equal(toBN(toWei('0', 'ether')).toString());

            // 3 deadlined
            await common.moveTime(web3, def_period+1);
            await s.reward_abi.reward();
            expect(toBN(await s.treasury_abi.balance()).sub(orig_bal).toString())
                .equal(toBN(toWei('200', 'ether')).toString());

            // another 2
            await common.moveTime(web3, (4*def_period/10)+1);
            await s.reward_abi.reward();
            expect(toBN(await s.treasury_abi.balance()).sub(orig_bal).toString())
                .equal(toBN(toWei('400', 'ether')).toString());

            // remaining 4
            await common.moveTime(web3, (5*def_period/10)+1);
            await s.reward_abi.reward();
            expect(toBN(await s.treasury_abi.balance()).sub(orig_bal).toString())
                .equal(toBN(toWei('800', 'ether')).toString());
        });

        it ('should correctly distribute full payouts', async() => {
            const orig_bal = toBN(await s.treasury_abi.balance());

            const get_proposal = async () => {
                const evt = await s.orig.getPastEvents('BudgetProposal', common.evt_last_block);
                expect(evt).lengthOf(1);
                return await IProposal.at(evt[0].args.proposal);
            };
            
            await s.treasury_abi.propose(
                toWei('300', 'ether'), '201', def_period,
                { value: fee_amount }
            );
            const proposal1 = await get_proposal();
            await s.treasury_abi.propose(
                toWei('300', 'ether'), '202', def_period,
                { value: fee_amount }
            );
            const proposal2 = await get_proposal();
            await s.treasury_abi.propose(
                toWei('300', 'ether'), '203', def_period,
                { value: fee_amount }
            );

            await proposal1.voteAccept({from: owner1});
            await proposal1.voteAccept({from: owner2});
            await proposal2.voteReject({from: owner1});
            await proposal2.voteReject({from: owner2});

            await s.reward_abi.reward();
            expect(toBN(await s.treasury_abi.balance()).sub(orig_bal).toString())
                .equal(toBN(toWei('-200', 'ether')).toString());

            await common.moveTime(web3, def_period+1);
            await s.reward_abi.reward();
            expect(toBN(await s.treasury_abi.balance()).sub(orig_bal).toString())
                .equal(toBN(toWei('-100', 'ether')).toString());
        });

        it ('should correctly distribute partial payouts', async() => {
            const orig_bal = toBN(await s.treasury_abi.balance());
            const start_bal3 = toBN(await web3.eth.getBalance(owner3));
            const start_bal4 = toBN(await web3.eth.getBalance(owner4));

            const get_proposal = async () => {
                const evt = await s.orig.getPastEvents('BudgetProposal', common.evt_last_block);
                expect(evt).lengthOf(1);
                return await IProposal.at(evt[0].args.proposal);
            };
            const mul1 = toBN(2);
            const mul2 = toBN(3);
            const tot = mul1.add(mul2);
            const addbal = orig_bal.mul(tot.sub(toBN(1)));
            const trunc = toBN(10);
            const trunc_bal = toBN('10000000000000000');
            const extra = toBN(12345678);
            
            await s.treasury_abi.propose(
                orig_bal.mul(mul1), '301', def_period,
                { value: fee_amount, from: owner3 }
            );
            const proposal1 = await get_proposal();
            await s.treasury_abi.propose(
                orig_bal.mul(mul2), '302', def_period,
                { value: fee_amount, from: owner4 }
            );
            const proposal2 = await get_proposal();

            await proposal1.voteAccept({from: owner1});
            await proposal1.voteAccept({from: owner2});
            await proposal2.voteAccept({from: owner1});
            await proposal2.voteAccept({from: owner2});

            await s.reward_abi.reward();
            expect(toBN(await s.treasury_abi.balance()).div(trunc).toString())
                .equal(toBN(0).toString());

            const evt1 = await s.orig.getPastEvents('Payout', common.evt_last_block);
            expect(evt1.length).equal(2);
            expect(evt1[0].args.amount.div(trunc).toString())
                .equal(orig_bal.mul(mul1).div(tot).div(trunc).toString());
            expect(evt1[1].args.amount.div(trunc).toString())
                .equal(orig_bal.mul(mul2).div(tot).div(trunc).toString());
            await proposal1.withdraw({from: owner3});
            await proposal2.withdraw({from: owner4});

            await s.reward_abi.reward({from: owner1, value: addbal.add(extra)});
            expect(toBN(await s.treasury_abi.balance()).div(trunc).toString())
                .equal(extra.div(trunc).toString());
            const evt2 = await s.orig.getPastEvents('Payout', common.evt_last_block);
            expect(evt2.length).equal(2);
            expect(evt2[0].args.amount.div(trunc).toString())
                .equal(addbal.mul(mul1).div(tot).div(trunc).toString());
            expect(evt2[1].args.amount.div(trunc).toString())
                .equal(addbal.mul(mul2).div(tot).div(trunc).toString());

            await s.reward_abi.reward();
            const evt3 = await s.orig.getPastEvents('Payout', common.evt_last_block);
            expect(evt3.length).equal(0);

            const after_bal1 = orig_bal.mul(mul1).add(start_bal3)
                .sub(toBN(await web3.eth.getBalance(owner3)))
                .div(trunc_bal);
            const after_bal2 = orig_bal.mul(mul2).add(start_bal4)
                .sub(toBN(await web3.eth.getBalance(owner4)))
                .div(trunc_bal);
            expect(after_bal1.toString()).equal(toBN(2).toString());
            expect(after_bal2.toString()).equal(toBN(2).toString());
        });

        it ('should refuse propose() on UUID duplicate of past proposal', async () => {
            try {
                await s.treasury_abi.propose(
                    def_amount, '201', def_period,
                    { value: fee_amount }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /UUID in use/);
            }
            try {
                await s.treasury_abi.propose(
                    def_amount, '202', def_period,
                    { value: fee_amount }
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /UUID in use/);
            }

        });

    });

    //---
    describe('StorageTreasuryV1', async () => {
        it ('should refuse setProposal() from outside', async () => {
            try {
                await s.storage.setProposal(0, s.fake.address);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });

        it ('should refuse deleteProposal() from outside', async () => {
            try {
                await s.storage.deleteProposal(s.fake.address);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});

