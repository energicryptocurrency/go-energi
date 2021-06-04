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
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');
const IMasternodeToken = artifacts.require('IMasternodeToken');
const StorageMasternodeTokenV1 = artifacts.require('StorageMasternodeTokenV1');

const common = require('./common');

contract("MasternodeTokenV1", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
        storage: null,
    };

    const COLLATERAL_1 = web3.utils.toWei('10000', 'ether');
    const COLLATERAL_2 = web3.utils.toWei('20000', 'ether');
    const COLLATERAL_3 = web3.utils.toWei('30000', 'ether');
    const COLLATERAL_4 = web3.utils.toWei('40000', 'ether');
    const COLLATERAL_7 = web3.utils.toWei('70000', 'ether');
    const COLLATERAL_9 = web3.utils.toWei('90000', 'ether');
    const COLLATERAL_10 = web3.utils.toWei('100000', 'ether');
    const COLLATERAL_13 = web3.utils.toWei('130000', 'ether');
    const check_age = async (age) => {
        const bn = await web3.eth.getBlockNumber();
        expect(age.toString()).equal(bn.toString());
    };

    before(async () => {
        s.orig = await MasternodeTokenV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.registry_proxy = await MockProxy.at(await s.orig.registry_proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await MasternodeTokenV1.at(s.proxy.address);
        s.token_abi = await IMasternodeToken.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        s.storage = await StorageMasternodeTokenV1.at(await s.proxy_abi.v1storage());
        Object.freeze(s);
    });

    after(async () => {
        const impl = await MasternodeTokenV1.new(s.proxy.address, s.registry_proxy.address);
        await s.proxy.setImpl(impl.address);

        await s.fake.testDrain(COLLATERAL_3, {from: accounts[1]});
        await s.fake.testDrain(COLLATERAL_1, {from: accounts[0]});
    });
    
    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('ERC20', () => {
        it('should emit Transfer in c-tor', async () => {
            const tmp = await MasternodeTokenV1.new(s.proxy.address, s.registry_proxy.address);

            const evt = await tmp.getPastEvents('Transfer', common.evt_last_block);
            expect(evt).lengthOf(1);
            expect(evt[0].args).deep.include({
                '__length__': 3,
                'from': '0x0000000000000000000000000000000000000000',
                'to': '0x0000000000000000000000000000000000000000',
                'value': web3.utils.toBN('0'),
            });
        });

        it('should support totalSupply()', async () => {
            const res = await s.token_abi.totalSupply();
            assert.equal(res.valueOf(), 0);
        });

        it('should support name()', async () => {
            const res = await s.token_abi.name();
            assert.equal(res, "Masternode Collateral");
        });

        it('should support symbol()', async () => {
            const res = await s.token_abi.symbol();
            assert.equal(res, "MNGR");
        });

        it('should support decimals()', async () => {
            const res = await s.token_abi.decimals();
            assert.equal(res.valueOf(), 22);
        });

        it('should support balanceOf()', async () => {
            const res = await s.token_abi.balanceOf(s.fake.address);
            assert.equal(res.valueOf(), 0);
        });

        it('should support allowance()', async () => {
            const res = await s.token_abi.allowance(s.fake.address, s.fake.address);
            assert.equal(res.valueOf(), 0);
        });

        it('should refuse transfer()', async () => {
            try {
                await s.token_abi.transfer(s.fake.address, '0');
                assert.fail("It must fail");
            } catch (e) {
                assert.match(e.message, /Not allowed/);
            }
        });

        it('should refuse transferFrom()', async () => {
            try {
                await s.token_abi.transferFrom(s.fake.address, s.fake.address, '0');
                assert.fail("It must fail");
            } catch (e) {
                assert.match(e.message, /Not allowed/);
            }
        });

        it('should refuse approve()', async () => {
            try {
                await s.token_abi.approve(s.fake.address, '0');
                assert.fail("It must fail");
            } catch (e) {
                assert.match(e.message, /Not allowed/);
            }
        });
    });

    //---
    describe('Primary', () => {
        it('should support balanceInfo()', async () => {
            const res = await s.token_abi.balanceInfo(s.fake.address);
            assert.equal(res['0'].valueOf(), 0);
        });

        it('should allow depositCollateral()', async () => {
            const { logs } = await s.token_abi.depositCollateral({
                from: accounts[0],
                value: COLLATERAL_1,
            });
            assert.equal(logs.length, 1);
            const res = await s.token_abi.balanceInfo(accounts[0]);
            assert.equal(res['0'].valueOf(), COLLATERAL_1);
            await check_age(res['1']);

            const res2 = await s.token_abi.balanceOf(accounts[0]);
            assert.equal(res2.valueOf(), COLLATERAL_1);

            const res3 = await s.token_abi.totalSupply();
            assert.equal(res3.valueOf(), COLLATERAL_1);

            const evt = await s.orig.getPastEvents('Transfer', common.evt_last_block);
            expect(evt).lengthOf(1);
            expect(evt[0].args).deep.include({
                '__length__': 3,
                'from': '0x0000000000000000000000000000000000000000',
                'to': accounts[0],
                'value': web3.utils.toBN(COLLATERAL_1),
            });
        });

        it('should correctly reflect last block', async () => {
            const res1 = await s.token_abi.balanceInfo(accounts[0]);
            await common.moveTime(web3, 3600);

            const res2 = await s.token_abi.balanceInfo(accounts[0]);
            assert.equal(res2['0'].valueOf(), COLLATERAL_1);
            assert.equal(res1['1'].toString(), res2['1'].toString());
        });
        
        it('should allow depositCollateral() direct', async () => {
            const { logs } = await s.orig.depositCollateral({
                from: accounts[0],
                value: COLLATERAL_2,
            });
            assert.equal(logs.length, 1);
            const res = await s.token_abi.balanceInfo(accounts[0]);
            assert.equal(res['0'].valueOf(), COLLATERAL_3);
            await check_age(res['1']);

            const res2 = await s.token_abi.balanceOf(accounts[0]);
            assert.equal(res2.valueOf(), COLLATERAL_3);

            const total = await s.token_abi.totalSupply();

            assert.equal(total.valueOf(), COLLATERAL_3);

            const evt = await s.orig.getPastEvents('Transfer', common.evt_last_block);
            expect(evt).lengthOf(1);
            expect(evt[0].args).deep.include({
                '__length__': 3,
                'from': '0x0000000000000000000000000000000000000000',
                'to': accounts[0],
                'value': web3.utils.toBN(COLLATERAL_2),
            });
        });

        it('should refuse depositCollateral() not a multiple of', async () => {
            try {
                await s.token_abi.depositCollateral({
                    from: accounts[0],
                    value: web3.utils.toWei('10001', 'ether'),
                });
                assert.fail("It must fail");
            } catch (e) {
                assert.match(e.message, /Not a multiple/);
            }

            const evt = await s.orig.getPastEvents('Transfer', common.evt_last_block);
            expect(evt).lengthOf(0);
        });

        it('should allow depositCollateral() - max', async () => {
            const { logs } = await s.token_abi.depositCollateral({
                from: accounts[0],
                value: COLLATERAL_7,
            });
            assert.equal(logs.length, 1);
            const res = await s.token_abi.balanceInfo(accounts[0]);
            assert.equal(res['0'].valueOf(), COLLATERAL_10);
            await check_age(res['1']);

            const res2 = await s.token_abi.balanceOf(accounts[0]);
            assert.equal(res2.valueOf(), COLLATERAL_10);

            const total = await s.token_abi.totalSupply();
            assert.equal(total.valueOf(), COLLATERAL_10);
        });

        it('should refuse to depositCollateral() over max', async () => {
            try {
                await s.token_abi.depositCollateral({
                    from: accounts[0],
                    value: web3.utils.toWei(COLLATERAL_1, 'ether'),
                });
                assert.fail("It must fail");
            } catch (e) {
                assert.match(e.message, /Too much/);
            }
        });

        it('should allow depositCollateral() another account', async () => {
            const { logs } = await s.orig.depositCollateral({
                from: accounts[1],
                value: COLLATERAL_3,
            });
            assert.equal(logs.length, 1);

            const res = await s.token_abi.balanceInfo(accounts[1]);
            assert.equal(res['0'].valueOf(), COLLATERAL_3);
            await check_age(res['1']);

            const res2 = await s.token_abi.balanceOf(accounts[1]);
            assert.equal(res2.valueOf(), COLLATERAL_3);

            const total = await s.token_abi.totalSupply();
            assert.equal(total.valueOf(), COLLATERAL_13);
        });

        it('should allow withdrawCollateral()', async () => {
            const { logs } = await s.token_abi.withdrawCollateral(COLLATERAL_9, {
                from: accounts[0],
            });
            assert.equal(logs.length, 1);
            const res = await s.token_abi.balanceInfo(accounts[0]);
            assert.equal(res['0'].valueOf(), COLLATERAL_1);
            await check_age(res['1']);

            const total = await s.token_abi.totalSupply();
            assert.equal(total.valueOf(), COLLATERAL_4);

            const evt = await s.orig.getPastEvents('Transfer', common.evt_last_block);
            expect(evt).lengthOf(1);
            expect(evt[0].args).deep.include({
                '__length__': 3,
                'from': accounts[0],
                'to': '0x0000000000000000000000000000000000000000',
                'value': web3.utils.toBN(COLLATERAL_9),
            });
        });

        it('should refuse withdrawCollateral() over balance', async () => {
            try {
                await s.token_abi.withdrawCollateral(COLLATERAL_2, {
                    from: accounts[0],
                });
                assert.fail("It must fail");
            } catch (e) {
                assert.match(e.message, /Not enough/);
            }

            const evt = await s.orig.getPastEvents('Transfer', common.evt_last_block);
            expect(evt).lengthOf(0);
        });

        it('should refuse setBalance() on s.storage', async () => {
            try {
                await s.storage.setBalance(s.fake.address, COLLATERAL_1, COLLATERAL_1);
                assert.fail("It must fail");
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});
