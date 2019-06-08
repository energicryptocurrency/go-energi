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
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');
const IMasternodeToken = artifacts.require('IMasternodeToken');
const StorageMasternodeTokenV1 = artifacts.require('StorageMasternodeTokenV1');

contract("MasternodeTokenV1", async accounts => {
    let orig;
    let fake;
    let proxy;
    let proxy_abi;
    let token_abi;
    let storage;

    // NOTE: some BigNumber issues with Truffle exposed web3...
    const COLLATERAL_1 = web3.utils.toWei('10000', 'ether');
    const COLLATERAL_2 = web3.utils.toWei('20000', 'ether');
    const COLLATERAL_3 = web3.utils.toWei('30000', 'ether');
    const COLLATERAL_4 = web3.utils.toWei('40000', 'ether');
    const COLLATERAL_7 = web3.utils.toWei('70000', 'ether');
    const COLLATERAL_9 = web3.utils.toWei('90000', 'ether');
    const COLLATERAL_10 = web3.utils.toWei('100000', 'ether');
    const COLLATERAL_13 = web3.utils.toWei('130000', 'ether');
    const check_age = (age) => {
        assert.isBelow(parseInt(age.valueOf()), 10);
    };

    before(async () => {
        orig = await MasternodeTokenV1.deployed();
        proxy = await MockProxy.at(await orig.proxy());
        fake = await MockContract.new(proxy.address);
        proxy_abi = await MasternodeTokenV1.at(proxy.address);
        token_abi = await IMasternodeToken.at(proxy.address);
        await proxy.setImpl(orig.address);
        storage = await StorageMasternodeTokenV1.at(await proxy_abi.v1storage());
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

    // ERC20 stuff
    //---
    it.skip('should emit Transfer in c-tor', async () => {});

    it('should support totalSupply()', async () => {
        const res = await token_abi.totalSupply();
        assert.equal(res.valueOf(), 0);
    });

    it('should support name()', async () => {
        const res = await token_abi.name();
        assert.equal(res, "Masternode Collateral");
    });

    it('should support symbol()', async () => {
        const res = await token_abi.symbol();
        assert.equal(res, "MNGR");
    });

    it('should support decimals()', async () => {
        const res = await token_abi.decimals();
        assert.equal(res.valueOf(), 22);
    });

    it('should support balanceOf()', async () => {
        const res = await token_abi.balanceOf(fake.address);
        assert.equal(res.valueOf(), 0);
    });

    it('should support allowance()', async () => {
        const res = await token_abi.allowance(fake.address, fake.address);
        assert.equal(res.valueOf(), 0);
    });

    it('should refuse transfer()', async () => {
        try {
            await token_abi.transfer(fake.address, '0');
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not allowed/);
        }
    });

    it('should refuse transferFrom()', async () => {
        try {
            await token_abi.transferFrom(fake.address, fake.address, '0');
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not allowed/);
        }
    });

    it('should refuse approve()', async () => {
        try {
            await token_abi.approve(fake.address, '0');
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not allowed/);
        }
    });

    // Energi stuff
    //---
    it('should support balanceInfo()', async () => {
        const res = await token_abi.balanceInfo(fake.address);
        assert.equal(res['0'].valueOf(), 0);
    });

    it('should allow depositCollateral()', async () => {
        const { logs } = await token_abi.depositCollateral({
            from: accounts[0],
            value: COLLATERAL_1,
        });
        assert.equal(logs.length, 1);
        const res = await token_abi.balanceInfo(accounts[0]);
        assert.equal(res['0'].valueOf(), COLLATERAL_1);
        check_age(res['1']);
    });

    it('should correctly reflect age', async () => {
        await new Promise((resolve, reject) => {
            web3.currentProvider.send({
                jsonrpc: "2.0",
                method: "evm_increaseTime",
                params: [3600],
                id: new Date().getSeconds()
            }, resolve);
        });
        await new Promise((resolve, reject) => {
            web3.currentProvider.send({
                jsonrpc: "2.0",
                method: "evm_mine",
                params: [],
                id: new Date().getSeconds() + 1
            }, resolve);
        });
        const res = await token_abi.balanceInfo(accounts[0]);
        assert.equal(res['0'].valueOf(), COLLATERAL_1);
        assert.isAtLeast(parseInt(res['1'].valueOf()), 3600);
    });
    
    it('should allow depositCollateral() direct', async () => {
        const { logs } = await orig.depositCollateral({
            from: accounts[0],
            value: COLLATERAL_2,
        });
        assert.equal(logs.length, 1);
        const res = await token_abi.balanceInfo(accounts[0]);
        assert.equal(res['0'].valueOf(), COLLATERAL_3);
        check_age(res['1']);
    });

    it('should refuse depositCollateral() not a multiple of', async () => {
        try {
            await token_abi.depositCollateral({
                from: accounts[0],
                value: web3.utils.toWei('10001', 'ether'),
            });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not a multiple/);
        }
    });

    it('should allow depositCollateral() - max', async () => {
        const { logs } = await token_abi.depositCollateral({
            from: accounts[0],
            value: COLLATERAL_7,
        });
        assert.equal(logs.length, 1);
        const res = await token_abi.balanceInfo(accounts[0]);
        assert.equal(res['0'].valueOf(), COLLATERAL_10);
        check_age(res['1']);
    });

    it('should refuse to depositCollateral() over max', async () => {
        try {
            await token_abi.depositCollateral({
                from: accounts[0],
                value: web3.utils.toWei(COLLATERAL_1, 'ether'),
            });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Too much/);
        }
    });

    it('should allow depositCollateral() another account', async () => {
        const { logs } = await orig.depositCollateral({
            from: accounts[1],
            value: COLLATERAL_3,
        });
        assert.equal(logs.length, 1);

        const res = await token_abi.balanceInfo(accounts[1]);
        assert.equal(res['0'].valueOf(), COLLATERAL_3);
        check_age(res['1']);

        const total = await token_abi.totalSupply();
        assert.equal(total.valueOf(), COLLATERAL_13);

    });

    it('should allow withdrawCollateral()', async () => {
        const { logs } = await token_abi.withdrawCollateral(COLLATERAL_9, {
            from: accounts[0],
        });
        assert.equal(logs.length, 1);
        const res = await token_abi.balanceInfo(accounts[0]);
        assert.equal(res['0'].valueOf(), COLLATERAL_1);
        check_age(res['1']);

        const total = await token_abi.totalSupply();
        assert.equal(total.valueOf(), COLLATERAL_4);
    });

    it('should refuse withdrawCollateral() over balance', async () => {
        try {
            await token_abi.withdrawCollateral(COLLATERAL_2, {
                from: accounts[0],
            });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not enough/);
        }
    });

    it('should refuse setBalance() on storage', async () => {
        try {
            await storage.setBalance(fake.address, COLLATERAL_1, COLLATERAL_1);
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not owner/);
        }
    });


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

    it('should transfer storage & allow to kill() it', async () => {
        await fake.killStorage(storage.address);
    });
});
