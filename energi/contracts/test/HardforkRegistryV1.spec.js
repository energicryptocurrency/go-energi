// Copyright 2020 The Energi Core Authors
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
const HardforkRegistryV1 = artifacts.require('HardforkRegistryV1');
const IHardforkRegistry = artifacts.require('IHardforkRegistry');
const StorageHardforkRegistryV1 = artifacts.require('StorageHardforkRegistryV1');

const common = require('./common');
const ethers = require('ethers');

contract("HardforkRegistryV1", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
    };

    before(async () => {
        s.orig = await HardforkRegistryV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.fake = await MockContract.new(s.proxy.address);

        s.proxy_abi = await HardforkRegistryV1.at(s.proxy.address);
        s.proxy_hf = await IHardforkRegistry.at(s.proxy.address);
        s.token_abi = s.proxy_hf;
        await s.proxy.setImpl(s.orig.address);

        s.storage = await StorageHardforkRegistryV1.at(await s.proxy_abi.v1storage());
        Object.freeze(s);
    });

    describe('common pre', () => common.govPreTests(s));

    describe("Primary", async () => {
        const b32 = (name) => ethers.utils.formatBytes32String(name);
        const emptyB32 = b32(""); // "0x0000000000000000000000000000000000000000000000000000000000000000"

        const hf_sw_feature = 3000600;
        const owner1 = "0x1be31a94361a391bbafb2a4ccd704f57dc04d4bb";
        let hf_names = [b32("Ba Sing Se"), b32("Hogwarts"), b32("Mars"), b32("Random")];

        let hf_blocks = [];
        let hf_hashes = [];

        before(async () => {
            let i = 0;
            for(; i < hf_names.length; i++) {
                common.moveTime(web3, 1);

                const b = await web3.eth.getBlock('latest');
                // Hardfork needs to be created way ahead of time.
                hf_blocks.push(b.number+2)
                if (i > 1) {
                    hf_hashes.push(emptyB32);
                } else {
                    hf_hashes.push(b.hash);
                }

                await s.proxy_hf.propose(hf_blocks[i], hf_names[i], hf_hashes[i], hf_sw_feature, {from: owner1});
            }
        })

        it("should fail to propose on invalid HF signer", async () => {
            common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');

            try {
                await s.proxy_hf.propose(b.number, b32("Invalid Signer"), emptyB32, hf_sw_feature, {from: s.fake.address});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid hardfork signer caller/);
            }
        });

        it("should fail to propose when empty hardfork name is used", async () => {
            common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');

            try {
                await s.proxy_hf.propose(b.number, emptyB32, emptyB32, hf_sw_feature);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Hardfork name cannot be empty/);
            }
        });

        it("should fail to propose on a past HF block", async () => {
            try {
                await s.proxy_hf.propose(5, b32("Adromeda"), emptyB32, hf_sw_feature);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Hardfork cannot be created in the past/);
            }
        });

        it("should fail to propose if duplicate name is used", async () => {
            common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');

            try {
                await s.proxy_hf.propose(b.number+5, b32("Ba Sing Se"), b.hash, hf_sw_feature);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Duplicate hardfork names are not allowed/);
            }
        });

        it("should propose on correct inputs", async ()=> {
            common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            let bn = b.number+3
        
            try {
                hf_blocks.push(bn);
                hf_names.push(b32("Krypton"));
                await s.proxy_hf.propose(bn, b32("Krypton"), b.hash, hf_sw_feature);
            } catch (e) {
                assert.match(e.message, /It must fail/);
            }

            const evt = await s.orig.getPastEvents('Hardfork', common.evt_last_block);
            expect(evt).lengthOf(1);
            common.stringifyBN(web3, evt[0].args);
            expect(evt[0].args).deep.include({
                block_no: bn.toString(10),
                block_hash: b.hash,
                name: b32("Krypton"),
            });
        });


        it("should fail to propose if the hardfork has already been finalized", async () => {
            const block_no = hf_blocks[hf_blocks.length -1];
            const b = await web3.eth.getBlock('latest');
            try {
                await s.proxy_hf.propose(block_no, b32("Ba Sing Se - Updated"), b.hash, hf_sw_feature);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /hardfork changes not editable/);
            }
        });

        it("should fail to propose during the hardfork finalization if block hash is empty", async () => {
            const b = await web3.eth.getBlock('latest');
            const bn = b.number+3;
            const hfn = b32("Barbage");

            try {
                await s.proxy_hf.propose(bn, hfn, emptyB32, hf_sw_feature);
                hf_blocks.push(bn);
            } catch (e) {
                assert.match(e.message, /It must fail/);
            }

            // Move time till the block finalization period.
            let i = 0;
            for(; i < 11; i++) {
                common.moveTime(web3, 1);
            }

            try {
                await s.proxy_hf.propose(bn, hfn, emptyB32, hf_sw_feature);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /HF finalization block hash cannot be empty/);
            }
        });

        it("should propose during the hardfork finalization if block hash is not empty", async () => {
            const b = await web3.eth.getBlock('latest');
            const bn = b.number+5;
            const hfn = b32("Karl Max");

            try {
                await s.proxy_hf.propose(bn, hfn, emptyB32, hf_sw_feature);
                hf_blocks.push(bn);
            } catch (e) {
                assert.match(e.message, /It must fail/);
            }

            // Move time till PAST the block finalization period.
            let i = 0;
            for(; i < 13; i++) {
                common.moveTime(web3, 1);
            }

            try {
                await s.proxy_hf.propose(bn, hfn, b.hash, hf_sw_feature);
            } catch (e) {
                assert.match(e.message, /It must fail/);
            }

            const evt = await s.orig.getPastEvents('Hardfork', common.evt_last_block);
            expect(evt).lengthOf(1);
            common.stringifyBN(web3, evt[0].args);
            expect(evt[0].args).deep.include({
                block_no: bn.toString(),
                block_hash: b.hash,
                name: hfn,
            });
        });

        it("should fail to propose when the hardfork finalization interval is exceeded", async () => {
            const b = await web3.eth.getBlock('latest');

            try {
                await s.proxy_hf.propose(hf_blocks[2], hf_names[2], b.hash, hf_sw_feature);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Hardfork finalization interval exceeded/);
            }
        });

        it("getByBlockNo with none existent block number should return empty values", async () => {
            common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');

            let info = await s.proxy_hf.getByBlockNo(b.number + 100);
            common.stringifyBN(web3, info);

            expect(info).deep.include({
                block_hash: emptyB32,
                name: emptyB32,
                sw_features: "0",
            });
        });

        it("getByBlockNo with existent block number should return correct values", async () => {
            let info = await s.proxy_hf.getByBlockNo(hf_blocks[2]);
            common.stringifyBN(web3, info);

            expect(info).deep.include({
                block_hash: hf_hashes[2],
                name: hf_names[2],
                sw_features: hf_sw_feature.toString(),
            });
        });

        it("getByName with none existent name should return empty values", async () => {
            let info = await s.proxy_hf.getByName(b32("knight"));
            common.stringifyBN(web3, info);

            expect(info).deep.include({
                block_no: "0",
                block_hash: emptyB32,
                sw_features: "0",
            });
        });

        it("getByName with existent name should return correct values", async () => {
            let info = await s.proxy_hf.getByName(hf_names[2]);
            common.stringifyBN(web3, info);

            expect(info).deep.include({
                block_no: hf_blocks[2].toString(),
                block_hash: hf_hashes[2],
                sw_features: hf_sw_feature.toString()
            });
        });

        it("should list all the hardforks blocks", async () => {
            let array  = await s.proxy_hf.enumerate();
            let i = 0;
            let returnArray = [];
            for(; i < array.length; i++) {
                returnArray.push(array[i].words[0]);
            }
            expect(returnArray).members(hf_blocks);
        });

        it("should fail to remove on invalid HF signer", async () => {
            try {
                await s.proxy_hf.remove(hf_blocks[2], {from: s.fake.address});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid hardfork signer caller/);
            }
        });

        it("should fail to remove a finalized hardfork information", async () => {
            try {
                await s.proxy_hf.remove(hf_blocks[1]);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Finalized hardfork cannot be deleted/);
            }

            // Confirm that the hardfork wasn't removed at all.
            let b = await s.proxy_hf.getByBlockNo(hf_blocks[1]);
            common.stringifyBN(web3, b);

            expect(b).deep.include({
                name: hf_names[1].toString(),
                block_hash: hf_hashes[1],
                sw_features: hf_sw_feature.toString()
            });

            let array  = await s.proxy_hf.enumerate();
            expect(array.length).to.equal(hf_blocks.length);

            let i = 0;
            let returnArray = [];
            for(; i < array.length; i++) {
                returnArray.push(array[i].words[0]);
            }
            expect(returnArray).to.contain(hf_blocks[1]);
        });

        it("should remove the unfinalized hardfork information completely", async () => {
            await s.proxy_hf.remove(hf_blocks[3]);

            // Confirm that the hardfork wasn't removed at all.
            let b = await s.proxy_hf.getByBlockNo(hf_blocks[3]);
            common.stringifyBN(web3, b);

            expect(b).deep.include({
                name: emptyB32,
                block_hash: emptyB32,
                sw_features: "0"
            });

            let array  = await s.proxy_hf.enumerate();
            expect(array.length).to.equal(hf_blocks.length - 1);

            let i = 0;
            let returnArray = [];
            for(; i < array.length; i++) {
                returnArray.push(array[i].words[0]);
            }
            expect(returnArray).not.contain(hf_blocks[3]);
        });

        describe("StorageHardforkRegistryV1", async () => {
            it("should refuse to update a hardfork directly", async () => {
                try {
                    await s.storage.setHardfork(hf_blocks[0], b32("Ba Sing Se- Updated"), hf_hashes[0], hf_sw_feature);
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Not owner/);
                }
            });

            it("should refuse to delete a hardfork directly", async () => {
                try {
                    await s.storage.deleteHardfork(hf_blocks[0]);
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Not owner/);
                }
            });
        });
    });

    describe('common post', () => common.govPostTests(s));
});