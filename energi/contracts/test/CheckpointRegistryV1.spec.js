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
const CheckpointRegistryV1 = artifacts.require('CheckpointRegistryV1');
const ICheckpointRegistry = artifacts.require('ICheckpointRegistry');
const ICheckpoint = artifacts.require('ICheckpoint');
const StorageCheckpointRegistryV1 = artifacts.require('StorageCheckpointRegistryV1');

const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');

const common = require('./common');
const ethjs = require('ethereumjs-util');

contract("CheckpointRegistryV1", async accounts => {
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

        s.orig = await CheckpointRegistryV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await CheckpointRegistryV1.at(s.proxy.address);
        s.token_abi = await ICheckpointRegistry.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        s.storage = await StorageCheckpointRegistryV1.at(await s.proxy_abi.v1storage());
        Object.freeze(s);
    });

    after(async () => {
        const impl = await CheckpointRegistryV1.new(
            s.proxy.address, s.registry.address, accounts[3]);
        await s.proxy.setImpl(impl.address);
    });

    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('Primary', () => {
        const { fromAscii, toBN, toWei } = web3.utils;

        const collateral1 = toBN(toWei('50000', 'ether'));
        const owner1 = accounts[0];
        const mnacc1 = web3.eth.accounts.create();
        const nonmnacc1 = web3.eth.accounts.create();
        const masternode1 = mnacc1.address;
        const ip1 = toBN(0x12345678);
        const enode_common = '123456789012345678901234567890'
        const enode1 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '11')];

        const ecsign = (acc, hash) => {
            const sig = ethjs.ecsign(
                toBN(hash).toArrayLike(Buffer),
                toBN(acc.privateKey).toArrayLike(Buffer)
            );
            return '0x'+[sig.r.toString('hex'), sig.s.toString('hex'), sig.v.toString(16)].join('');
        };
    
        const cp_count = 100;
        const cp_sign = cp_count - 1;
        const block_hash = '0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef';
        let cp_list;
        const mn_sig = ecsign(mnacc1, block_hash);

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

        it('should refuse propose() from invalid signer', async () => {
            try {
                await s.token_abi.propose(1, block_hash);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid caller/);
            }
        });

        it('should propose() from valid signer', async () => {
            for (let i = 1; i <= cp_count; ++i) {
                await s.token_abi.propose(parseInt(i/2), block_hash, {from: accounts[3]});
            }
        });

        it('should checkpoints()', async () => {
            cp_list = await s.token_abi.checkpoints();
            expect(cp_list.length).equal(cp_count);
        });

        it('should refuse to sign() by non-MN', async() => {
            try {
                await s.token_abi.sign(cp_list[cp_sign], ecsign(nonmnacc1, block_hash));
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not active MN/);
            }
        });
        
        it('should refuse to sign() by invalid signature length', async() => {
            try {
                await s.token_abi.sign(cp_list[cp_sign], block_hash);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Invalid signature length/);
            }
        });

        it('should sign() by MN', async() => {
            await s.token_abi.sign(cp_list[cp_sign], mn_sig);
        });

        it('should refuse to sign() by already signed MN', async() => {
            try {
                await s.token_abi.sign(cp_list[cp_sign], mn_sig);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Already signed/);
            }
        });

        describe('CheckpointV1', async () => {
            it('should show info()', async () => {
                const cp = await ICheckpoint.at(cp_list[cp_sign]);
                const res = await cp.info();

                common.stringifyBN(web3, res);
                expect(res).include({
                    number: toBN(parseInt((cp_sign+1)/2)).toString(),
                    hash:   block_hash,
                });
            });

            it('should show signature()', async () => {
                const cp = await ICheckpoint.at(cp_list[cp_sign]);
                const res = await cp.signature(masternode1);

                expect(res.toString()).equal(mn_sig);
            });

            it('should fail signature() on not  signed', async () => {
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
                const res = await cp.signature(masternode1);

                expect(res).equal(mn_sig);
            });
        });
    });

    //---
    describe('StorageCheckpointRegistryV1', async () => {
        it ('should refuse add() from outside', async () => {
            try {
                await s.storage.add(s.fake.address);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });

        it ('should listCheckpoints() from outside', async () => {
            await s.storage.listCheckpoints();
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});
