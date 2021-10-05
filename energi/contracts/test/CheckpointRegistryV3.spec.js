// Copyright 2021 The Energi Core Authors
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
const CheckpointRegistryV3 = artifacts.require('CheckpointRegistryV3');
const ICheckpoint = artifacts.require('ICheckpoint');
const ICheckpointV2 = artifacts.require('ICheckpointV2');
const ICheckpointRegistryV2 = artifacts.require('ICheckpointRegistryV2');
const StorageCheckpointRegistryV1 = artifacts.require('StorageCheckpointRegistryV1');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');

const common = require('./common');
const ethjs = require('ethereumjs-util');

contract("CheckpointRegistryV3", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
    };

    before(async () => {
        s.registry_orig = await MasternodeRegistryV1.deployed();
        s.registry = await MasternodeRegistryV1.at(await s.registry_orig.proxy());
        s.mntoken_orig = await MasternodeTokenV1.deployed();
        s.mntoken = await MasternodeTokenV1.at(await s.mntoken_orig.proxy());
        s.orig = await CheckpointRegistryV3.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await CheckpointRegistryV3.at(s.proxy.address);
        s.token_abi = await ICheckpointRegistryV2.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        s.storage = await StorageCheckpointRegistryV1.at(await s.proxy_abi.v2storage());
        Object.freeze(s);
    });

    after(async () => {
        const impl = await CheckpointRegistryV3.new(s.proxy.address, s.registry.address, accounts[3]);
        await s.proxy.setImpl(impl.address);
    });

    describe('common pre', () => common.govPreTests(s));

    describe('Primary', () => {
        const { fromAscii, toBN, toWei } = web3.utils;
        const collateral1 = toBN(toWei('50000', 'ether'));
        const owner1 = accounts[0];
        const sigacc = web3.eth.accounts.privateKeyToAccount('0x4118811427785a33e8c61303e64b43d0d6b69db3caa4074f2ddbdec0b9d4c878');
        const mnacc1 = web3.eth.accounts.create();
        const nonmnacc1 = web3.eth.accounts.create();
        web3.eth.personal.importRawKey("0x4118811427785a33e8c61303e64b43d0d6b69db3caa4074f2ddbdec0b9d4c878","");
        web3.eth.personal.unlockAccount('0x2d0bc327d0843caf6fd9ae1efab0bf7196fc2fc8','');
        web3.eth.sendTransaction({to:'0x2d0bc327d0843caf6fd9ae1efab0bf7196fc2fc8', from: accounts[0], value: toWei('5000000', 'ether')});
        common.moveTime(web3, 10);
        
        const masternode1 = mnacc1.address;
        const ip1 = toBN(0x12345678);
        const enode_common = '123456789012345678901234567890';
        const enode1 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '11')];
        const ecsign = (acc, hash) => {
            const sig = ethjs.ecsign(
                toBN(hash).toArrayLike(Buffer),
                toBN(acc.privateKey).toArrayLike(Buffer)
            );
            return '0x' + [sig.r.toString('hex'), sig.s.toString('hex'), sig.v.toString(16)].join('');
        };
        const cp_count = 5;
        const cp_max_count = 10;
        const cp_sign = cp_max_count - 1;
        const block_hash = '0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef';

        let cp_list;
        let mn_sig;
        let cpp_sig;

        const mn_sig_reg = async (acc, num, block_hash) => {
            const sigbase = await s.token_abi.signatureBase(num, block_hash);
            return ecsign(acc, sigbase);
        };

        const mn_sig_cp = async (acc, cp_address) => {
            const cp = await ICheckpoint.at(cp_address);
            const sigbase = await cp.signatureBase();
            return ecsign(acc, sigbase);
        };

        before(async () => {
            await s.mntoken.depositCollateral({
                from: owner1,
                value: collateral1,
            });
            await s.registry.announce(masternode1, ip1, enode1, {from: owner1});
        });

        after(async () => {
            await s.mntoken.withdrawCollateral(collateral1, {
                from: owner1,
            });
        });

        it('should refuse propose() with invalid signature length', async () => {
            try {
                await s.token_abi.propose(1, block_hash, block_hash, {from : sigacc.address});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid signature length/);
            }
        });

        it('should refuse propose() from invalid signer', async () => {
            try {
                await s.token_abi.propose(1, block_hash, await mn_sig_reg(mnacc1, 1, block_hash), {from : sigacc.address});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid signer/);
            }
        });

        it('should refuse remove() from invalid signer', async () => {
            try {
                await s.token_abi.remove(1, block_hash, {from :accounts[0]});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not cpp signer!/);
            }
        });


        describe('Primary', () => {
            it('(QUEUE TEST) should propose() and add 5 checkpoints', async () => {
                for (let i = 0; i < cp_count; i++) {
                    cpp_sig = await mn_sig_reg(sigacc, i, block_hash);
                    await s.token_abi.propose(i, block_hash, cpp_sig, {from : sigacc.address});
                }
                assert.equal((await s.token_abi.checkpoints()).length, 5 , "didnt add checkpoints");
            });

            it('(QUEUE TEST) should allow remove() from valid signer but checkpoint will not be found (size remains same)', async () => {
                try {
                    const cps1 = await s.token_abi.checkpoints();
                    const num = 12831; //some random
                    cpp_sig = await mn_sig_reg(sigacc, num, block_hash);
                    await s.token_abi.remove(num, block_hash, {from : sigacc.address});
                    const cps2 = await s.token_abi.checkpoints();
                    assert.equal(cps1.length, cps2.length,"stored checkpoints should remain same");
                } catch (e) {
                    assert.fail("must fail");
                }
            });


            it('(QUEUE TEST) should allow remove() from valid signer and decrease checkpoint size', async () => {
                try {
                    const cps1 = await s.token_abi.checkpoints();
                    const num = 3; //some random
                    cpp_sig = await mn_sig_reg(sigacc, num, block_hash);
                    await s.token_abi.remove(num, block_hash, {from : sigacc.address});
                    const cps2 = await s.token_abi.checkpoints();
                    assert.equal(cps1.length - 1, cps2.length,"stored checkpoints should remain same");
                } catch (e) {
                    assert.fail("must fail");
                }
            });

            it('(QUEUE TEST) should allow remove() from valid signer but checkpoint will not be found (already deleted)', async () => {
                try {
                    const cps1 = await s.token_abi.checkpoints();
                    const num = 3; //some random
                    cpp_sig = await mn_sig_reg(sigacc, num, block_hash);
                    await s.token_abi.remove(num, block_hash, {from : sigacc.address});
                    const cps2 = await s.token_abi.checkpoints();
                    assert.equal(cps1.length, cps2.length,"stored checkpoints should remain same");
                } catch (e) {
                    assert.fail("must fail");
                }
            });




            it('(QUEUE TEST) should be 10 checkpoints in the end()', async () => {
                for (let i = 20; i < 100; i++) {
                    cpp_sig = await mn_sig_reg(sigacc, i, block_hash);
                    await s.token_abi.propose(i, block_hash, cpp_sig, {from : sigacc.address});
                }
                cp_list = await s.token_abi.checkpoints();
                assert.equal(cp_list.length, cp_max_count, "must me max 10 checkpoints");
            });

            it('(QUEUE TEST) should have removed the last element in queue()', async () => {
                cpp_sig = await mn_sig_reg(sigacc, 99, block_hash);
                await s.token_abi.remove(99, block_hash, {from : sigacc.address});
                cp_list = await s.token_abi.checkpoints();
                assert.equal(cp_list.length, cp_max_count - 1, "must me 9 checkpoints remaining");
            });

            it('(QUEUE TEST) should have removed the first element in queue()', async () => {
                cpp_sig = await mn_sig_reg(sigacc, 90, block_hash);
                await s.token_abi.remove(90, block_hash, {from : sigacc.address});
                cp_list = await s.token_abi.checkpoints();
                assert.equal(cp_list.length, cp_max_count - 2, "must me 8 checkpoints remaining");
            });

            it('(QUEUE TEST) should be max checkpoints ()', async () => {
                for (let i = 100; i < 103; i++) {
                    cpp_sig = await mn_sig_reg(sigacc, i, block_hash);
                    await s.token_abi.propose(i, block_hash, cpp_sig, {from : sigacc.address});
                }
                cp_list = await s.token_abi.checkpoints();
                assert.equal(cp_list.length, cp_max_count, "must me max 10 checkpoints");
            });

            it('(QUEUE TEST) should remove from the middle of the queue()', async () => {
                cpp_sig = await mn_sig_reg(sigacc, 100, block_hash);
                await s.token_abi.remove(100, block_hash, {from : sigacc.address});
                var cp_list2 = await s.token_abi.checkpoints();
                assert.equal(cp_list2[0], cp_list[0], "first elements must remain same");
                assert.equal(cp_list2[cp_list2.length-1], cp_list[cp_list.length-1], "last elements must remain same");
                assert.equal(cp_list2.length, cp_max_count-1, "1 minus element");
            });

            it('(QUEUE TEST) should be max checkpoints ()', async () => {
                for (let i = 1000; i < 1002; i++) {
                    cpp_sig = await mn_sig_reg(sigacc, i, block_hash);
                    await s.token_abi.propose(i, block_hash, cpp_sig, {from : sigacc.address});
                }
                cp_list = await s.token_abi.checkpoints();
                assert.equal(cp_list.length, cp_max_count, "must me max 10 checkpoints");
            });

            it('(QUEUE TEST) should refuse to sign() by non-MN', async() => {
                try {
                    await s.token_abi.sign(cp_list[cp_sign], ecsign(nonmnacc1, block_hash));
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Not active MN/);
                }
            });

            it('(QUEUE TEST) should refuse to sign() by invalid signature length', async() => {
                try {
                    await s.token_abi.sign(cp_list[cp_sign], block_hash);
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Invalid signature length/);
                }
            });

            it('should sign() by MN', async() => {
                mn_sig = await mn_sig_cp(mnacc1, cp_list[cp_sign]);
                await s.token_abi.sign(cp_list[cp_sign], mn_sig);
            });

            it('should refuse to sign() by already signed MN', async() => {
                try {
                    await s.token_abi.sign(cp_list[cp_sign], await mn_sig_cp(mnacc1, cp_list[cp_sign]));
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Already signed/);
                }
            });

            it('should refuse to sign() by CPP signer', async() => {
                try {
                    await s.token_abi.sign(cp_list[cp_sign], await mn_sig_cp(sigacc, cp_list[cp_sign]));
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Already signed/);
                }
            });

            it('should have correct signatureBase()', async () => {
                const hash = await s.token_abi.signatureBase(cp_sign + 1, block_hash);
                const reqhash = web3.utils.soliditySha3(
                    "||Energi Blockchain Checkpoint||",
                    toBN(cp_sign + 1),
                    toBN(block_hash),
                );
                expect(hash.toString()).equal(reqhash.toString());
            });
        });

        describe('CheckpointV2', async () => {
            it('should show info()', async () => {
                const cp = await ICheckpoint.at(cp_list[cp_sign]);
                const res = await cp.info();
                common.stringifyBN(web3, res);
                const num = 1001;
                expect(res).include({
                    number: toBN(parseInt(num)).toString(),
                    hash:   block_hash,
                });

                expect(res).include.keys('since');
            });

            it('should show signature()', async () => {
                const cp = await ICheckpoint.at(cp_list[cp_sign]);
                const res = await cp.signature(masternode1);

                expect(res.toString()).equal(mn_sig);
            });

            it('should show signature() of CPP signer', async () => {
                const cp = await ICheckpoint.at(cp_list[cp_sign]);
                const res = await cp.signature(sigacc.address);

                expect(res.toString()).equal(cpp_sig);
            });

            it('should fail signature() on not signed', async () => {
                const cp = await ICheckpoint.at(cp_list[cp_sign]);
                try {
                    await cp.signature(accounts[0]);
                    assert.fail('It must fail');
                } catch (e) {
                    assert.match(e.message, /Not signed yet/);
                }
            });

            it('should show signatures()', async () => {
                const cp = await ICheckpoint.at(cp_list[cp_sign]);
                const res = await cp.signatures();

                expect(res).include(cpp_sig);
                expect(res).include(mn_sig);
            });

            it('should have correct signatureBase()', async () => {
                const cp = await ICheckpoint.at(cp_list[cp_sign]);
                const hash = await cp.signatureBase();
                const num = 1001;
                const reqhash = web3.utils.soliditySha3(
                    "||Energi Blockchain Checkpoint||",
                    toBN(parseInt(num)),
                    toBN(block_hash),
                );
                expect(hash.toString()).equal(reqhash.toString());
            });

            it('should correctly handle canVote()', async () => {
                const num = 101;
                await s.token_abi.propose(num, block_hash, await mn_sig_reg(sigacc, num, block_hash), {from : sigacc.address});
                const cps = await s.token_abi.checkpoints();
                const cpa = cps[cps.length - 1];
                const cp = await ICheckpointV2.at(cpa);

                expect(await cp.canVote(sigacc.address)).false;
                expect(await cp.canVote(masternode1)).true;

                await s.token_abi.sign(cpa, await mn_sig_cp(mnacc1, cpa));

                expect(await cp.canVote(masternode1)).false;
                expect(await cp.canVote(nonmnacc1.address)).true;

                for (let i = 0; i < 24 * 60 - 1; ++i) {
                    try {
                        expect(await cp.canVote(nonmnacc1.address)).true;
                    } catch (e) {
                        console.log(`Iteration ${i}`);
                        throw e;
                    }
                    common.moveTime(web3, 1);
                }

                expect(await cp.canVote(nonmnacc1.address)).false;
            });
        });
    });

    describe('common post', () => common.govPostTests(s) );
});
