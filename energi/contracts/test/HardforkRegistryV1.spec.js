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
        s.proxy = await MockProxy.at(s.orig.proxy());
        await s.proxy.setImpl(s.orig.address);

        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_hf = await IHardforkRegistry.at(s.proxy.address);
        s.storage = await StorageHardforkRegistryV1.at(s.proxy.address);
        Object.freeze(s);
    });

    describe('common pre', () => common.govPreTests(s));

    describe("Primary", () => {
        const hf_blocks = [10, 20]
        const hf_names = ["Ba Sing Se", "Hogwarts"]
        const hf_sw_features = [3000600, 3000600]

        let pos = 0;
        for(i = 0; i <= Math.max(...hf_blocks)+1; i++) {
            common.moveTime(web3, 1);

            if (pos < hf_blocks.length && i == hf_blocks[pos]) {
                let hash = "";

                if (pos == 0) {
                    // Finalize only one of the hardforks.
                    const bn = await web3.eth.getBlockNumber();
                    const b = await web3.eth.getBlock(bn);

                    hash = b.hash
                }

                await s.proxy_hf.propose(hf_blocks[pos], hash, hf_names[pos], hf_sw_features[pos]);
                pos++;
            }
        }

        it("should fail to propose on invalid HF signer", async () => {
            try {
                await s.fake.propose(50, "", "Invalid Signer", 3000600);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid hardfork signer caller/);
            }
        });

        it("should fail to propose on past HF", async () => {
            try {
                await s.proxy_hf.propose(5, "", "Adromeda", 3000600);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Hardfork cannot be created in the past/);
            }
        });

        it("should fail to propose if duplicate name is used", async () => {
            let blockHash = "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421";
            try {
                await s.proxy_hf.propose(25, blockHash, "Ba Sing Se", 3000600);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Duplicate hardfork names are not allowed/);
            }
        });

        it("should propose on correct inputs", async ()=> {
            common.moveTime(web3, 1);

            const bn = web3.eth.getBlockNumber();
            const b = web3.eth.getBlock(bn);
        
            try {
                await s.proxy_hf.propose(bn, b.hash, "Krypton", 3000600);
            } catch (e) {
                assert.match(e.message, /It must fail/);
            }

            const evt = await s.orig.getPastEvents('hardfork', common.evt_last_block);
            expect(evt).lengthOf(1);
        });

        it("getByBlockNo with none existent block number should return empty values", async () => {
            let info = await s.proxy_hf.getByBlockNo(100);
            common.stringifyBN(web3, info);

            expect(info).deep.include({
                block_hash: "0x000000000000000000000000000000000000",
                name: "0x00000000",
                sw_fetaures: 0
            });
        });

        it("getByBlockNo with existent block number should return correct values", async () => {
            let info = await s.proxy_hf.getByBlockNo(10);
            common.stringifyBN(web3, info);

            expect(info).deep.include({
                block_hash: "0x000000000000000000000000000000000000",
                name: "0x00000000",
                sw_fetaures: 0
            });
        });

        it("getByName with none existent name should return empty values", async () => {
            let info = await s.proxy_hf.getByName("knight");
            common.stringifyBN(web3, info);

            expect(info).deep.include({
                block_hash: "0x000000000000000000000000000000000000",
                name: "0x00000000",
                sw_fetaures: 0
            });
        });

        it("getByName with existent name should return correct values", async () => {
            let info = await s.proxy_hf.getByName("Ba Sing Se");
            common.stringifyBN(web3, info);

            expect(info).deep.include({
                block_hash: "0x000000000000000000000000000000000000",
                name: "0x00000000",
                sw_fetaures: 0
            });
        });

        it("should list all the hardforks blocks", async () => {
            let array  = await s.proxy_hf.enumerate();
            expect(array).to.eql([10, 20, 50]);
        });

        it("should fail to remove a finalized hardfork information", async () => {
            try {
                await s.proxy_hf.remove(10);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Finalized hardfork cannot be deleted/);
            }

            // Confirm that the hardfork wasn't removed at all.
            let b_hash, name, sw_features = await s.proxy_hf.getByBlockNo(10);

            expect(b_hash).to.equal("0x000000000000000000000000000000000000");
            expect(name).to.equal("0x00000000");
            expect(sw_features).to.equal(0);

            let array  = await s.proxy_hf.enumerate();
            expect(array).to.contain(10);
        });

        it("should remove the unfinalized hardfork information completely", async () => {
            try {
                await s.proxy_hf.remove(20);
            } catch (e) {
                assert.match(e.message, /It must fail/);
            }

            // Confirm that the hardfork wasn't removed at all.
            let b_hash, name, sw_features = await s.proxy_hf.getByBlockNo(20);

            expect(b_hash).to.equal("0x000000000000000000000000000000000000");
            expect(name).to.equal("0x00000000");
            expect(sw_features).to.equal(0);

            let array  = await s.proxy_hf.enumerate();
            expect(array).to.contain(20);
        });

        describe("StorageHardforkRegistryV1", async () => {
            it("should refuse to update a finalised hardfork", async () => {
                const b = web3.eth.getBlock(11);
                try {
                    await s.storage.setHardfork(10, b.hash, "Ba Sing Se- Updated", 3000600);
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /hardfork changes not editable/);
                }
            });

            it("should refuse to delete a finalized hardfork", async () => {
                try {
                    await s.storage.deleteHardfork(10);
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /hardfork changes not editable/);
                }
            });
        });
    });

    describe('common post', () => common.govPostTests(s));
});