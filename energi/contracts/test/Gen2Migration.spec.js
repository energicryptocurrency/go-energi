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
const ITreasury = artifacts.require('ITreasury');
const MockGen2Migration = artifacts.require('MockGen2Migration');

const common = require('./common');
const ethjs = require('ethereumjs-util');

contract("Gen2Migration", async accounts => {
    let orig;
    let treasury_proxy;

    before(async () => {
        orig = await Gen2Migration.deployed();
        treasury_proxy = await ITreasury.at(await orig.treasury_proxy());
        orig = await MockGen2Migration.new(treasury_proxy.address, common.chain_id, accounts[0]); 
    });

    after(async () => {
        await Gen2Migration.new(treasury_proxy.address, common.chain_id, common.migration_signer);
    });
    
    // Primary stuff
    //---
    const { toBN, toWei } = web3.utils;

    const acc1 = web3.eth.accounts.create();
    const acc2 = web3.eth.accounts.create();
    const owner1 = acc1.address;
    const owner2 = acc2.address;
    const dst = accounts[1];

    const bal1 = toBN(toWei('100', 'ether'));
    const bal2 = toBN(toWei('200', 'ether'));
    
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
        ], [
            bal1,
            bal2,
        ], { value: bal1.add(bal2) });

        expect((await orig.itemCount()).toString()).equal(toBN(2).toString());

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

    it('should throw on invalid item ID', async () => {
        try {
            await orig.verifyClaim(3, dst, 0, '0x00', '0x00');
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Invalid ID/);
        }
    });

    it('should throw on invalid signature', async () => {
        try {
            await orig.verifyClaim(1, dst, 0, '0x00', '0x00');
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Invalid signature/);
        }
    });

    it('should verifyClaim()', async () => {
        const hash = await orig.hashToSign(dst);
        await orig.verifyClaim(1, dst, ...ecsign(acc2, hash));
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

    it('should claim() another account', async () => {
        const hash = await orig.hashToSign(dst);
        await orig.claim(0, dst, ...ecsign(acc1, hash));
    });

    it('should refuse to drain()', async() => {
        try {
            await orig.drain();
            assert.fail('It must fail');
        } catch (e) {
            assert.match(e.message, /Not treasury/);
        }
    });

    it.skip('should allow to drain() by Treasury', async() => {
        // TODO: only manually tested
    });

    it('should signerAddress()', async () => {
        const signer = await orig.signerAddress();
        expect(signer).equal(accounts[0]);
    });

    it('should totalAmount()', async () => {
        const total = await orig.totalAmount();
        expect(total.toString()).equal(bal1.add(bal2).toString());
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
