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

const Gen2Migration = artifacts.require('Gen2Migration');
const MockGen2Migration = artifacts.require('MockGen2Migration');
const MockGen2MigrationBlacklist = artifacts.require('MockGen2MigrationBlacklist');
const MockProxy = artifacts.require('MockProxy');
const BlacklistRegistryV1 = artifacts.require('BlacklistRegistryV1');

const common = require('./common');
const ethjs = require('ethereumjs-util');

contract("Gen2Migration", async accounts => {
    let orig;
    let blacklist_registry;

    before(async () => {
        const blacklist_proxy = await MockProxy.new();

        orig = await MockGen2Migration.new(
            blacklist_proxy.address, common.chain_id, accounts[0]);

        blacklist_registry = await BlacklistRegistryV1.deployed();
        blacklist_registry = await MockGen2MigrationBlacklist.new(
            blacklist_proxy.address, await blacklist_registry.mnregistry_proxy(),
            orig.address, await blacklist_registry.compensation_fund(),
            accounts[9],
            { gas: "30000000" });
        await blacklist_proxy.setImpl(blacklist_registry.address);
    });

    after(async () => {
        await Gen2Migration.new(
            blacklist_registry.address, common.chain_id, common.migration_signer);
    });
    
    // Primary stuff
    //---
    const { toBN, toWei } = web3.utils;

    const acc1 = web3.eth.accounts.create();
    const acc2 = web3.eth.accounts.create();
    const acc3 = web3.eth.accounts.create();
    const owner1 = acc1.address;
    const owner2 = acc2.address;
    const owner3 = acc3.address;
    const dst = accounts[1];

    const bal1 = toBN(toWei('100', 'ether'));
    const bal2 = toBN(toWei('200', 'ether'));
    const bal3 = toBN(toWei('300', 'ether'));
    
    const ecsign = (acc, hash) => {
        const sig = ethjs.ecsign(
            toBN(hash).toArrayLike(Buffer),
            toBN(acc.privateKey).toArrayLike(Buffer)
        );
        return [sig.v, '0x'+sig.r.toString('hex'), '0x'+sig.s.toString('hex')];
    };
    
    it('should correctly reflect itemCount()', async () => {
        await orig.setCoins([
            owner1,
            owner2,
            owner3,
        ], [
            bal1,
            bal2,
            bal3,
        ], [], { value: bal1.add(bal2).add(bal3) });

        expect((await orig.itemCount()).toString()).equal(toBN(3).toString());

        const c = await orig.coins(1);
        expect(c.amount.toString()).equal(bal2.toString());
        expect(toBN(c.owner).toString()).equal(toBN(owner2).toString());
    });

    it('should correctly create hashToSign()', async () => {
        const hash = await orig.hashToSign(dst);
        const reqhash = web3.utils.soliditySha3(
            dst,
            "||Energi Gen 2 migration claim||",
            toBN(common.chain_id)
        );
        expect(hash.toString()).equal(reqhash.toString());
    });

    it('should refuse to claim() on invalid item ID', async () => {
        try {
            await orig.verifyClaim(4, dst, 0, '0x00', '0x00');
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Invalid ID/);
        }
    });

    it('should refuse to claim() on invalid signature', async () => {
        try {
            await orig.verifyClaim(1, dst, 0, '0x00', '0x00');
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Invalid signature/);
        }
    });

    it('should refuse to blacklistClaim() on invalid ID', async() => {
        try {
            await orig.blacklistClaim(3, dst);
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Invalid ID/);
        }
    });

    it('should refuse to blacklistClaim() on invalid owner', async() => {
        try {
            await orig.blacklistClaim(1, owner3);
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Invalid Owner/);
        }
    });

    it('should verifyClaim()', async () => {
        const hash = await orig.hashToSign(dst);
        await orig.verifyClaim(1, dst, ...ecsign(acc2, hash));
    });

    it('should refuse claim() to claim blacklisted', async () => {
        await blacklist_registry.setBlacklisted(owner2, true);

        try {
            const hash = await orig.hashToSign(dst);
            await orig.claim(1, dst, ...ecsign(acc2, hash), common.zerofee_callopts);
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Owner is blacklisted/);
        } finally {
            await blacklist_registry.setBlacklisted(owner2, false);
        }
    });

    it('should claim()', async () => {
        const bal_before = await web3.eth.getBalance(dst);

        const hash = await orig.hashToSign(dst);
        await orig.claim(1, dst, ...ecsign(acc2, hash), common.zerofee_callopts);

        const bal_after = await web3.eth.getBalance(dst);
        expect(toBN(bal_after).sub(toBN(bal_before)).toString()).equal(bal2.toString());
        
        const evt = await orig.getPastEvents('Migrated', common.evt_last_block);
        expect(evt).lengthOf(1);
        common.stringifyBN(web3, evt[0].args);
        expect(evt[0].args).deep.include({
            '0': '1',
            '1': dst,
            '2': bal2.toString(),
            '__length__': 3,
            'item_id': '1',
            'destination': dst,
            'amount': bal2.toString()
        });
    });

    it('should refuse to claim() again', async () => {
        const hash = await orig.hashToSign(dst);

        try {
            await orig.claim(1, dst, ...ecsign(acc2, hash), common.zerofee_callopts);
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Already spent/);
        }

        const evt = await orig.getPastEvents('Migrated', common.evt_last_block);
        expect(evt).lengthOf(0);
    });

    it('should refuse to blacklistClaim() on spent', async() => {
        try {
            await orig.blacklistClaim(1, owner2);
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Already spent/);
        }
    });

    it('should claim() another account', async () => {
        const hash = await orig.hashToSign(dst);
        await orig.claim(0, dst, ...ecsign(acc1, hash), common.zerofee_callopts);
    });

    it('should refuse to blacklistClaim() not by blacklist', async() => {
        try {
            await orig.blacklistClaim(2, owner3);
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Not blacklist registry/);
        }
    });

    it('should refuse claim() on hard blacklist', async () => {
        try {
            await orig.setCoins([
                owner1,
                owner2,
                owner3,
            ], [
                bal1,
                bal2,
                bal3,
            ], [ owner3 ]);

            const hash = await orig.hashToSign(dst);
            await orig.claim(2, dst, ...ecsign(acc3, hash), common.zerofee_callopts);
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Owner is hard blacklisted/);
        }
    });

    it('should allow to blacklistClaim() by Blacklist', async() => {
        const fund = await blacklist_registry.compensation_fund();
        const bal_before = await web3.eth.getBalance(fund);
        await blacklist_registry.drainMigration(2, owner3);
        const bal_after = await web3.eth.getBalance(fund);
        expect(toBN(bal_after).sub(toBN(bal_before)).toString()).equal(bal3.toString());
    });

    it('should signerAddress()', async () => {
        const signer = await orig.signerAddress();
        expect(signer).equal(accounts[0]);
    });

    it('should totalAmount()', async () => {
        const total = await orig.totalAmount();
        expect(total.toString()).equal(bal1.add(bal2).add(bal3).toString());
    });

    // Safety & Cleanup
    //---
    it('should refuse to accept funds', async () => {
        try {
            await orig.send(web3.utils.toWei('1', "ether"));
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Not supported/);
        }
    });
});
