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

        const hf_sw_feature = 3000700;
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

          it("should show hardfork is not active (as it doesn't exist)", async () => {
              for (var i=0;i<41;i++) {
                await common.moveTime(web3, 1);
              }
              const b = await web3.eth.getBlock('latest');
              try {
                  var isactive = await s.proxy_hf.isActive(b32("some name"),{from: hf_signer});
                  expect(isactive).to.equal(false);
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
                  assert.match(e.message, /Hardfork is already in effect or doesn't exist/);
              }
          });


          it("should fail updating active hardfork", async () => {
            await common.moveTime(web3, 1);
              const b = await web3.eth.getBlock('latest');
              try {
                  await s.proxy_hf.add(hf_names[0], b.number+100, hf_sw_feature, {from: hf_signer});
                  assert.fail('It must fail');
              } catch (e) {
                  assert.match(e.message, /Hardfork is already in effect or doesn't exist/);
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
            } catch (e) {
                assert.match(e.message,/Hardfork is already in effect or doesn't exist/);
            }
        });

        it("should fail on removing hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.remove(hf_active[0], {from: hf_signer});
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


        it("should not finalize empty name hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                await s.proxy_hf.finalize(b32(""),{from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message,/Hardfork doesn't exist/);
            }
        });


        it("should not finalize non existing hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                await s.proxy_hf.finalize(b32("some non existing hardfork"),{from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message,/Hardfork doesn't exist/);
            }
        });


        it("should finalize hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.finalize(hf_active[0],{from: hf_signer});
                const evt = await s.orig.getPastEvents('HardforkFinalized', common.evt_last_block);
                expect(res.logs).lengthOf(1);
                common.stringifyBN(web3, evt[0].args);
                expect(evt[0].args).deep.include({
                    name: hf_active[0],
                    sw_features: hf_sw_feature.toString()
                });
            } catch (e) {
                assert.fail('It must fail')
            }
        });


        it("should not finalize already finalized hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                await s.proxy_hf.finalize(hf_active[0],{from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Hardfork already finalized/);
            }
        });




        it("should catch event for adding new hardfork", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.add(b32("Fyodor Dostoevsky"), b.number+40, hf_sw_feature, {from: hf_signer});
                const evt = await s.orig.getPastEvents('HardforkCreated', common.evt_last_block);
                expect(res.logs).lengthOf(1);
                common.stringifyBN(web3, evt[0].args);
                hf_pending.push(b32("Fyodor Dostoevsky"));
                expect(evt[0].args).deep.include({
                    block_number: (b.number+40).toString(),
                    name: b32("Fyodor Dostoevsky"),
                    sw_features: hf_sw_feature.toString()
                });
            } catch (e) {
                assert.fail('It must fail');
            }
        });

        it("should catch event for adding new hardfork", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.add(b32("Leo Tolstoy"), b.number+40, hf_sw_feature, {from: hf_signer});
                const evt = await s.orig.getPastEvents('HardforkCreated', common.evt_last_block);
                expect(res.logs).lengthOf(1);
                common.stringifyBN(web3, evt[0].args);
                hf_pending.push(b32("Leo Tolstoy"));
                expect(evt[0].args).deep.include({
                    block_number: (b.number+40).toString(),
                    name: b32("Leo Tolstoy"),
                    sw_features: hf_sw_feature.toString()
                });
            } catch (e) {
                assert.fail('It must fail');
            }
        });

        it("should not finalize pending hardfork", async () => {
            await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                await s.proxy_hf.finalize(hf_pending[0],{from: hf_signer});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message,/Hardfork not eligible for finalizing/);
            }
        });


        it("should enumerate pending hardforks", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.enumeratePending({from: hf_signer});
                expect(res).members(hf_pending);
            } catch (e) {
                assert.fail('It must fail');
            }
        });

        it("should enumerate active hardforks", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.enumerateActive({from: hf_signer});
                expect(res).members(hf_active);
            } catch (e) {
                assert.fail('It must fail');
            }
        });



        it("should enumerate hardforks", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.enumerate({from: hf_signer});
                var enumerates = hf_active.concat(hf_pending);
                expect(res).members(enumerates);
            } catch (e) {
                assert.fail('It must fail');
            }
        });

        it("should return hf to be active", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.isActive(hf_active[0] ,{from: hf_signer});
                expect(res).to.equal(true);
            } catch (e) {
                assert.fail('It must fail');
            }
        });

        it("should return hf not to be active", async () => {
          await common.moveTime(web3, 1);
            const b = await web3.eth.getBlock('latest');
            try {
                var res = await s.proxy_hf.isActive(hf_pending[0] ,{from: hf_signer});
                expect(res).to.equal(false);
            } catch (e) {
                assert.fail('It must fail');
            }
        });






    });

});
