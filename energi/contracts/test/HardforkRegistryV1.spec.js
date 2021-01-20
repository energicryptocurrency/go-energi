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

const MockAutoProxy = artifacts.require('MockAutoProxy');
const MockAutoContract = artifacts.require('MockAutoContract');
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
        s.proxy = await MockAutoProxy.at(await s.orig.proxy());
        s.fake = await MockAutoContract.new();

        s.proxy_abi = await HardforkRegistryV1.at(s.proxy.address);
        s.proxy_hf = await IHardforkRegistry.at(s.proxy.address);
        s.token_abi = s.proxy_hf;
        // await s.proxy.setImpl(s.orig.address);

        s.storage = await StorageHardforkRegistryV1.at(await s.proxy_abi.v1storage());
        Object.freeze(s);
    });

    describe('common pre', () => common.govPreTests(s));

    describe("Primary", async () => {
        const b32 = (name) => ethers.utils.formatBytes32String(name);
        const emptyB32 = b32(""); // "0x0000000000000000000000000000000000000000000000000000000000000000"

        const hf_sw_feature = 3000600;
        const hf_signer = common.hf_signer;
        let hf_names = [b32("Ba Sing Se"), b32("Hogwarts"), b32("Mars"), b32("Random")];
        let hf_active = [b32("Ba Sing Se"), b32("Hogwarts"), b32("Mars"), b32("Random")];
        let hf_pending = [];

        let hf_blocks = [];
        let hf_hashes = [];

        before(async () => {
            let i = 0;
            for(; i < hf_names.length; i++) {
                common.moveTime(web3, 1);

                const b = await web3.eth.getBlock('latest');
                // Hardfork needs to be created way ahead of time.
                hf_blocks.push(b.number+40)
                if (i > 1) {
                    hf_hashes.push(emptyB32);
                } else {
                    hf_hashes.push(b.hash);
                }
                await s.proxy_hf.add(hf_names[i], hf_blocks[i], hf_sw_feature, {from: hf_signer});
            }
        });

        it("should return existing hardfork names", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var existing_hf_names = await s.proxy_hf.enumerate({from: hf_signer});
                expect(existing_hf_names).members(hf_names);
            } catch (e) {
                assert.fail('cannot get Hardfork names');
            }
        });

        it("should refuse to set a hardfork directly", async () => {
          await common.moveTime(web3, 1);
            try {
                await s.storage.set(hf_names[0], 10, 10);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });

        it("should refuse to finalize a hardfork directly", async () => {
          await common.moveTime(web3, 1);
            try {
                await s.storage.finalize(hf_names[0], 10);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });

        it("should refuse to remove a hardfork directly", async () => {
          await common.moveTime(web3, 1);
            try {
                await s.storage.remove(hf_names[0]);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });

        it("should fail adding hardfork with empty name", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');

            try {
                await s.proxy_hf.add(b32(""),b.number+40, hf_sw_feature, {from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Hardfork name cannot be empty/);
            }
        });




        it("should show hardfork is active", async () => {
            for (var i=0;i<41;i++) {
              await common.moveTime(web3, 1);
            }
            const b = await web3.eth.getBlock('latest');
            try {
                var isactive = await s.proxy_hf.isActive(hf_names[0],{from: hf_signer});
                expect(isactive).to.equal(true);
            } catch (e) {
                assert.fail('It must fail');
            }
        });


        it("should fail adding incorrect block_number", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                await s.proxy_hf.add(hf_names[0], b.number, hf_sw_feature, {from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Hardfork is too soon/);
            }
        });


        it("should fail updating active hardfork", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                await s.proxy_hf.add(hf_names[0], b.number+100, hf_sw_feature, {from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Cannot modify active Hardfork/);
            }
        });


        it("should catch event for adding new hardfork", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.add(b32("Aristotle"), b.number+40, hf_sw_feature, {from: hf_signer});
                const evt = await s.orig.getPastEvents('HardforkCreated', common.evt_last_block);
                expect(res.logs).lengthOf(1);
                common.stringifyBN(web3, evt[0].args);
                expect(evt[0].args).deep.include({
                    block_number: (b.number+40).toString(),
                    name: b32("Aristotle"),
                    sw_features: hf_sw_feature.toString()
                });
            } catch (e) {
                assert.fail('It must fail');
            }
        });


        it("should return hardfork is not active (false)", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.isActive(b32("Aristotle"), {from: hf_signer});

                expect(res).to.equal(false);
            } catch (e) {
                assert.fail('It must fail');
            }
        });

        it("should have five active hardforks before removing", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.enumerate({from: hf_signer});
                expect(res.length).to.equal(5);
            } catch (e) {
                assert.fail('It must fail');
            }
        });

        it("should remove pending hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.remove(b32("Aristotle"), {from: hf_signer});
                const evt = await s.orig.getPastEvents('HardforkRemoved', common.evt_last_block);
                expect(res.logs).lengthOf(1);
                common.stringifyBN(web3, evt[0].args);
                expect(evt[0].args).deep.include({
                    name: b32("Aristotle")
                });
            } catch (e) {
                assert.fail('It must fail');
            }
        });


        it("should fail on removing non-existing hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.remove(b32("non existing hardfork"), {from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message,/Hardfork is already in effect or doesn't exist/);
            }
        });

        it("should fail on removing empty-named hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.remove(b32(""), {from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message,/Hardfork name cannot be empty/);
            }
        });


        it("should fail on removing active hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.remove(hf_names[0], {from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message,/Hardfork is already in effect or doesn't exist/);
            }
        });


        it("should have four active hardforks after removing", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.enumerate({from: hf_signer});
                expect(res).members(hf_names);
            } catch (e) {
                assert.fail('It must fail');
            }
        });













        //
        //
        // it("should return that hardfork is existing", async () => {
        //     common.moveTime(web3, 1);
        //     const b = await web3.eth.getBlock('latest');
        //     try {
        //         var ret  = await s.proxy_hf.add(hf_names[0], b.number + 40, hf_sw_feature,{from: hf_signer});
        //         common.stringifyBN(web3, ret);
        //         expect(info).deep.include({
        //             block_hash: hf_hashes[2],
        //             block_no: hf_blocks[2].toString(),
        //             sw_features: hf_sw_feature.toString(),
        //         });
        //     } catch (e) {
        //         assert.fail('It must fail');
        //     }
        // });
        //
        //
        //
        //
        //
        //
        //
        //
        //
        //
        // it("should fail to propose on a HF block within the block finalization count", async () => {
        //     try {
        //         await s.proxy_hf.propose(5, b32("Adromeda"), emptyB32, hf_sw_feature);
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /Hardfork cannot be scheduled immediately/);
        //     }
        // });
        //
        // it("should fail to propose if duplicate name is used", async () => {
        //     common.moveTime(web3, 1);
        //     const b = await web3.eth.getBlock('latest');
        //
        //     try {
        //         await s.proxy_hf.propose(b.number+40, b32("Ba Sing Se"), b.hash, hf_sw_feature);
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /Duplicate hardfork names are not allowed/);
        //     }
        // });
        //
        //
        //
        //
        //
        //
        //
        //
        //
        //
        //
        //
        //
        //
        //
        // it("should propose on correct inputs", async ()=> {
        //     common.moveTime(web3, 1);
        //     const b = await web3.eth.getBlock('latest');
        //     let bn = b.number+42
        //     let hfn = b32("Krypton");
        //
        //     try {
        //         hf_blocks.push(bn);
        //         hf_names.push(hfn);
        //         hf_active.push(hfn);
        //         await s.proxy_hf.propose(bn, hfn, b.hash, hf_sw_feature);
        //     } catch (e) {
        //         assert.match(e.message, /It must fail/);
        //     }
        //
        //     const evt = await s.orig.getPastEvents('Hardfork', common.evt_last_block);
        //     expect(evt).lengthOf(1);
        //     common.stringifyBN(web3, evt[0].args);
        //     expect(evt[0].args).deep.include({
        //         block_no: bn.toString(10),
        //         block_hash: b.hash,
        //         name: hfn,
        //         sw_features: hf_sw_feature.toString(),
        //     });
        // });
        //
        // it("should propose a hardfork on future on invalid HF signer", async () => {
        //     common.moveTime(web3, 1);
        //     const b = await web3.eth.getBlock('latest');
        //     const hfn = b32("lieze");
        //
        //     try {
        //         await s.proxy_hf.propose(b.number+40000, hfn, emptyB32, hf_sw_feature);
        //         hf_names.push(hfn);
        //         hf_pending.push(hfn);
        //     } catch (e) {
        //         assert.fail('It must fail');
        //     }
        // });
        //
        // it("should fail to propose if the hardfork has already been finalized", async () => {
        //     common.moveTime(web3, 1);
        //
        //     const block_no = hf_blocks[4];
        //     const hfns = hf_names[4];
        //     const b = await web3.eth.getBlock('latest');
        //     try {
        //         await s.proxy_hf.propose(block_no, hfns, b.hash, hf_sw_feature);
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /hardfork changes not editable/);
        //     }
        // });
        //
        // it("should fail to propose if duplicate block number is used", async () => {
        //     common.moveTime(web3, 1);
        //
        //     const block_no = hf_blocks[hf_blocks.length -1];
        //     const b = await web3.eth.getBlock('latest');
        //     try {
        //         await s.proxy_hf.propose(block_no, b32("Ba Sing Se - updated"), b.hash, hf_sw_feature);
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /Duplicate block numbers for unique hardforks are not allowed/);
        //     }
        // });
        //
        // it("should fail to propose during the hardfork finalization if block hash is empty", async () => {
        //     common.moveTime(web3, 1);
        //
        //     const b = await web3.eth.getBlock('latest');
        //     const bn = b.number+43;
        //     const hfn = b32("Barbage");
        //
        //     try {
        //         await s.proxy_hf.propose(bn, hfn, emptyB32, hf_sw_feature);
        //         hf_blocks.push(bn);
        //         hf_names.push(hfn);
        //         hf_active.push(hfn);
        //     } catch (e) {
        //         assert.match(e.message, /It must fail/);
        //     }
        //
        //     // Move time till the block finalization period.
        //     let i = 0;
        //     for(; i < 43; i++) {
        //         common.moveTime(web3, 1);
        //     }
        //
        //     try {
        //         await s.proxy_hf.propose(bn, hfn, emptyB32, hf_sw_feature);
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /HF finalization block hash cannot be empty/);
        //     }
        // });
        //
        // it("should propose during the hardfork finalization if block hash is not empty", async () => {
        //     const b = await web3.eth.getBlock('latest');
        //     const bn = b.number+50;
        //     const hfn = b32("Karl Max");
        //
        //     try {
        //         await s.proxy_hf.propose(bn, hfn, emptyB32, hf_sw_feature);
        //         hf_blocks.push(bn);
        //         hf_names.push(hfn);
        //         hf_active.push(hfn);
        //     } catch (e) {
        //         assert.match(e.message, /It must fail/);
        //     }
        //
        //     // Move time till PAST the block finalization period.
        //     let i = 0;
        //     for(; i < 43; i++) {
        //         common.moveTime(web3, 1);
        //     }
        //
        //     try {
        //         await s.proxy_hf.propose(bn, hfn, b.hash, hf_sw_feature);
        //     } catch (e) {
        //         assert.match(e.message, /It must fail/);
        //     }
        //
        //     const evt = await s.orig.getPastEvents('Hardfork', common.evt_last_block);
        //     expect(evt).lengthOf(1);
        //     common.stringifyBN(web3, evt[0].args);
        //     expect(evt[0].args).deep.include({
        //         block_no: bn.toString(),
        //         block_hash: b.hash,
        //         name: hfn,
        //         sw_features: hf_sw_feature.toString(),
        //     });
        // });
        //
        // it("should fail to propose when the hardfork finalization interval is exceeded", async () => {
        //     let i = 0;
        //     for(; i < 50; i++) {
        //         common.moveTime(web3, 1);
        //     }
        //     const b = await web3.eth.getBlock('latest');
        //
        //     try {
        //         await s.proxy_hf.propose(hf_blocks[2], hf_names[2], b.hash, hf_sw_feature);
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /Hardfork finalization interval exceeded/);
        //     }
        // });
        //
        // it("getHardfork with none existent block number should return empty values", async () => {
        //     common.moveTime(web3, 1);
        //     const b = await web3.eth.getBlock('latest');
        //
        //     let info = await s.proxy_hf.getHardfork(b32("none-existent"));
        //     common.stringifyBN(web3, info);
        //
        //     expect(info).deep.include({
        //         block_hash: emptyB32,
        //         block_no: "0",
        //         sw_features: "0",
        //     });
        // });
        //
        // it("getHardfork with existent block number should return correct values", async () => {
        //     let info = await s.proxy_hf.getHardfork(hf_names[2]);
        //     common.stringifyBN(web3, info);
        //
        //     expect(info).deep.include({
        //         block_hash: hf_hashes[2],
        //         block_no: hf_blocks[2].toString(),
        //         sw_features: hf_sw_feature.toString(),
        //     });
        // });
        //
        // it("should list all the hardforks names", async () => {
        //     let returnArray  = await s.proxy_hf.enumerateAll();
        //     expect(returnArray).members(hf_names);
        // });
        //
        // it("should list all the pending hardforks names", async () => {
        //     let returnArray  = await s.proxy_hf.enumeratePending();
        //     expect(returnArray).members(hf_pending);
        // });
        //
        // it("should list all the active hardforks names", async () => {
        //     let returnArray  = await s.proxy_hf.enumerateActive();
        //     expect(returnArray).members(hf_active);
        // });
        //
        // it("should fail to remove on invalid HF signer", async () => {
        //     try {
        //         await s.proxy_hf.remove(hf_names[2], {from: s.fake.address});
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /Invalid hardfork signer caller/);
        //     }
        // });
        //
        // it("should fail to remove on an invalid HF names", async () => {
        //     try {
        //         await s.proxy_hf.remove(b32("invalid-names"));
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /Hardfork name is unknown/);
        //     }
        // });
        //
        // it("should fail to remove a finalized hardfork information", async () => {
        //     try {
        //         await s.proxy_hf.remove(hf_names[1]);
        //         assert.fail('It must fail');
        //     } catch (e) {
        //         assert.match(e.message, /Finalized hardfork cannot be deleted/);
        //     }
        //
        //     // Confirm that the hardfork wasn't removed at all.
        //     let b = await s.proxy_hf.getHardfork(hf_names[1]);
        //     common.stringifyBN(web3, b);
        //
        //     expect(b).deep.include({
        //         block_no: hf_blocks[1].toString(),
        //         block_hash: hf_hashes[1],
        //         sw_features: hf_sw_feature.toString()
        //     });
        //
        //     let array  = await s.proxy_hf.enumerateAll();
        //     expect(array.length).to.equal(hf_names.length);
        //     expect(array).to.contain(hf_names[1]);
        // });
        //
        // it("should remove the unfinalized hardfork information completely", async () => {
        //     await s.proxy_hf.remove(hf_names[3]);
        //
        //     // Confirm that the hardfork wasn't removed at all.
        //     let b = await s.proxy_hf.getHardfork(hf_names[3]);
        //     common.stringifyBN(web3, b);
        //
        //     expect(b).deep.include({
        //         block_no: "0",
        //         block_hash: emptyB32,
        //         sw_features: "0"
        //     });
        //
        //     let array  = await s.proxy_hf.enumerateAll();
        //     expect(array.length).to.equal(hf_names.length - 1);
        //     expect(array).not.contain(hf_names[3]);
        // });
        //
        // it("should find hardfork with block that has passed as active", async() => {
        //     expect(await s.proxy_hf.isActive(b32("Krypton"))).to.be.true;
        // });
        //
        // it("should return false if an empty hardfork name is used", async() => {
        //     expect(await s.proxy_hf.isActive(emptyB32)).to.be.false;
        // });
        //
        // it("should find hardfork with block not yet passed inactive active", async() => {
        //     const b = await web3.eth.getBlock('latest');
        //     const bn = b.number+45
        //     const name = b32("Atlanta-")
        //     try {
        //         await s.proxy_hf.propose(bn, name, emptyB32, hf_sw_feature);
        //     } catch (e) {
        //         assert.match(e.message, /It must fail/);
        //     }
        //
        //     // The hardfork should be found in the system.
        //     let info = await s.proxy_hf.getHardfork(name);
        //     common.stringifyBN(web3, info);
        //     expect(info).deep.include({
        //         block_no: bn.toString(),
        //         block_hash: emptyB32,
        //         sw_features: hf_sw_feature.toString(),
        //     });
        //
        //     // The hardfork should be inactive because the block no is not yet mined.
        //     expect(await s.proxy_hf.isActive(name)).to.be.false;
        // });

        // describe("StorageHardforkRegistryV1", async () => {
        //     it("should refuse to update a hardfork directly", async () => {
        //         try {
        //             await s.storage.setHardfork(hf_blocks[0], b32("Ba Sing Se- Updated"), hf_hashes[0], hf_sw_feature);
        //             assert.fail('It must fail');
        //         } catch (e) {
        //             assert.match(e.message, /Not owner/);
        //         }
        //     });
        //
        //     it("should refuse to delete a hardfork directly", async () => {
        //         try {
        //             await s.storage.deleteHardfork(hf_names[0]);
        //             assert.fail('It must fail');
        //         } catch (e) {
        //             assert.match(e.message, /Not owner/);
        //         }
        //     });
        // });
    });

    // describe('common post', () => common.govPostTests(s));
});
