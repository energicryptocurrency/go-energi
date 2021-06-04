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
const BlacklistRegistryV1 = artifacts.require('BlacklistRegistryV1');
const IBlacklistRegistry = artifacts.require('IBlacklistRegistry');
const IProposal = artifacts.require('IProposal');
const Gen2Migration = artifacts.require('Gen2Migration');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');
const TreasuryV1 = artifacts.require('TreasuryV1');
const StorageBlacklistRegistryV1 = artifacts.require('StorageBlacklistRegistryV1');
const MockBlacklistProposalV1 = artifacts.require('MockBlacklistProposalV1');

const common = require('./common');

contract("BlacklistRegistryV1", async accounts => {
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

        s.treasury_orig = await TreasuryV1.deployed();
        s.treasury = await TreasuryV1.at(await s.treasury_orig.proxy());
        
        s.orig = await BlacklistRegistryV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.mnregistry_proxy = await MockProxy.at(await s.orig.mnregistry_proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await BlacklistRegistryV1.at(s.proxy.address);
        s.token_abi = await IBlacklistRegistry.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        s.storage = await StorageBlacklistRegistryV1.at(await s.proxy_abi.v1storage());
        s.compensation_fund = await s.orig.compensation_fund();
        Object.freeze(s);
    });

    after(async () => {
        const impl = await BlacklistRegistryV1.new(
            s.proxy.address, s.mnregistry_proxy.address,
            Gen2Migration.address, s.compensation_fund,
            accounts[9],
            { gas: "10000000" });
        await s.proxy.setImpl(impl.address);
    });

    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('Primary', () => {
        const { fromAscii, toBN, toWei } = web3.utils;
        const enforce_fee = toBN(toWei('1000', 'ether'));
        const revoke_fee = toBN(toWei('100', 'ether'));
        const drain_fee = toBN(toWei('100', 'ether'));

        const collateral1 = toBN(toWei('50000', 'ether'));
        const owner1 = accounts[0];
        const masternode1 = accounts[9];
        const ip1 = toBN(0x12345678);
        const enode_common = '123456789012345678901234567890'
        const enode1 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '11')];

        const target1 = accounts[1];
        const target2 = accounts[2];

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

        it('should not be enforcelisted by default', async () => {
            expect(await s.token_abi.isBlacklisted(target1)).false;
        });

        it('should return EBI signer', async () => {
            expect(await s.token_abi.EBI_signer()).equal(accounts[3]);
        });

        it('should refuse propose() without proper fee', async () => {
            try {
                await s.token_abi.propose(target1, { value: enforce_fee.sub(toBN(1)) });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            try {
                await s.token_abi.propose(target1, { value: enforce_fee.add(toBN(1)) });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            try {
                await s.token_abi.propose(target1, {
                    value: enforce_fee, from: await s.token_abi.EBI_signer()});
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }
        });

        it('should refuse proposeRevoke() without proper fee', async () => {
            try {
                await s.token_abi.propose(target1, { value: revoke_fee.sub(toBN(1)) });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            try {
                await s.token_abi.propose(target1, { value: revoke_fee.add(toBN(1)) });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            try {
                await s.token_abi.propose(target1, {
                    value: revoke_fee, from: await s.token_abi.EBI_signer()});
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }
        });


        it('should refuse proposeDrain() without proper fee', async () => {
            try {
                await s.token_abi.proposeDrain(target1, { value: drain_fee.sub(toBN(1)) });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            try {
                await s.token_abi.proposeDrain(target1, { value: drain_fee.add(toBN(1)) });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }

            try {
                await s.token_abi.proposeDrain(target1, {
                    value: drain_fee, from: await s.token_abi.EBI_signer()});
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid fee/);
            }
        });

        it('should refuse proposeRevoke() on no enforce voting', async () => {
            try {
                await s.token_abi.proposeRevoke(target1, { value: revoke_fee });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /No need \(1\)/);
            }
        });

        it('should enumerate*() empty', async () => {
            expect(await s.token_abi.enumerateAll()).length(0);
            expect(await s.token_abi.enumerateBlocked()).length(0);
        });

        it('should propose()', async () => {
            await s.token_abi.propose(target1, { value: enforce_fee });

            const evt = await s.orig.getPastEvents('BlacklistProposal', common.evt_last_block);
            expect(evt).lengthOf(1);
            common.stringifyBN(web3, evt[0].args);
            expect(evt[0].args).deep.include({
                '0': target1,
                '__length__': 2,
                'target': target1,
            });
            expect(evt[0].args).include.keys('proposal');

            expect(await s.token_abi.isBlacklisted(target1)).false;
        });

        it('should enumerate*() one', async () => {
            const all = await s.token_abi.enumerateAll();
            const blocked = await s.token_abi.enumerateBlocked();

            expect(all).eql([target1]);
            expect(blocked).eql(['0x0000000000000000000000000000000000000000']);
        });

        it('should refuse propose() on active enforce voting', async () => {
            try {
                await s.token_abi.propose(target1, { value: enforce_fee });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Already active \(2\)/);
            }
        });

        it('should refuse proposeRevoke() on active enforce voting', async () => {
            try {
                await s.token_abi.proposeRevoke(target1, { value: revoke_fee });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Not applicable/);
            }
        });

        it('should refuse proposeRevoke() on rejected enforcement', async () => {
            const old_proposal = await IProposal.at((await s.token_abi.proposals(target1))['0']);
            await old_proposal.voteReject({from: owner1});

            try {
                await s.token_abi.proposeRevoke(target1, { value: revoke_fee });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /No need \(2\)/);
            }
        });

        it('should propose() and collect rejected', async () => {
            await s.token_abi.propose(target1, { value: enforce_fee });
            expect(await s.treasury_orig.getPastEvents('Contribution', common.evt_last_block)).lengthOf(1);
            expect(await s.token_abi.isBlacklisted(target1)).false;
            expect(await s.token_abi.isBlacklisted(target2)).false;
        });

        it('should refuse collect() on lack of enforcement', async () => {
            try {
                await s.token_abi.collect(target2);
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Nothing to collect/);
            }
        });

        it('should refuse collect() on enforce voting', async () => {
            try {
                await s.token_abi.collect(target1);
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Enforce voting in progress/);
            }
        });

        it('should refuse propose() on enabled', async () => {
            const old_proposal = await IProposal.at((await s.token_abi.proposals(target1))['0']);
            await old_proposal.voteAccept({from: owner1});
            expect(await s.token_abi.isBlacklisted(target1)).true;
            expect(await s.token_abi.isBlacklisted(target2)).false;

            try {
                await s.token_abi.propose(target1, { value: enforce_fee });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Already active \(2\)/);
            }            
        });

        it('should refuse collect() on enforced', async () => {
            try {
                await s.token_abi.collect(target1);
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /No proposals ready to collect/);
            }
        });

        it('should proposeRevoke()', async () => {
            await s.token_abi.proposeRevoke(target1, { value: revoke_fee });
            expect(await s.token_abi.isBlacklisted(target1)).true;
            expect(await s.token_abi.isBlacklisted(target2)).false;
        });

        it('should refuse proposeRevoke() on active revoke voting', async () => {
            try {
                await s.token_abi.proposeRevoke(target1, { value: revoke_fee });
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Already active/);
            }
        });

        it('should proposeRevoke() and collect rejected', async () => {
            const old_proposal = await IProposal.at((await s.token_abi.proposals(target1))['1']);
            await old_proposal.voteReject({from: owner1});
            await s.token_abi.proposeRevoke(target1, { value: revoke_fee });
            expect(await s.treasury_orig.getPastEvents('Contribution', common.evt_last_block)).lengthOf(1);
            expect(await s.token_abi.isBlacklisted(target1)).true;
            expect(await s.token_abi.isBlacklisted(target2)).false;
        });

        it('should refuse collect() on revoke voting', async () => {
            try {
                await s.token_abi.collect(target1);
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Revoke voting in progress/);
            }
        });

        it('should collect() after rejected revocation', async () => {
            const old_proposal = await IProposal.at((await s.token_abi.proposals(target1))['1']);
            await old_proposal.voteReject({from: owner1});
            await s.token_abi.collect(target1);
            expect(await s.treasury_orig.getPastEvents('Contribution', common.evt_last_block)).lengthOf(1);
            expect(await s.token_abi.isBlacklisted(target1)).true;
            expect(await s.token_abi.isBlacklisted(target2)).false;
        });

        it('should collect() after approved revocation', async () => {
            await s.token_abi.proposeRevoke(target1, { value: revoke_fee });
            const old_proposal = await IProposal.at((await s.token_abi.proposals(target1))['1']);
            await old_proposal.voteAccept({from: owner1});
            await s.token_abi.collect(target1);
            expect(await s.treasury_orig.getPastEvents('Contribution', common.evt_last_block)).lengthOf(0);
            expect(await s.token_abi.isBlacklisted(target1)).false;
            expect(await s.token_abi.isBlacklisted(target2)).false;

        });

        it('should isObey() on accept', async () => {
            const proposal = await MockBlacklistProposalV1.new(s.mnregistry_proxy.address, accounts[0]);
            expect(await proposal.isObeyed()).false;
            await proposal.voteAccept();
            expect(await proposal.isObeyed()).true;
        });

        it('should isObeyed() on 2:1 weight over 10x votes', async () => {
            const proposal = await MockBlacklistProposalV1.new(s.mnregistry_proxy.address, accounts[0]);
            expect(await proposal.isObeyed()).false;

            const b = toBN(toWei('10000', 'ether'));

            // Must obey 10x
            await proposal.setWeights(toBN(20), toBN(9), toBN(10), toBN(30));
            expect(await proposal.isObeyed()).false;

            // 2>1
            await proposal.setWeights(toBN(20).mul(b), toBN(9).mul(b), toBN(1).mul(b), toBN(50).mul(b));
            expect(await proposal.isObeyed()).true;
            expect(await proposal.isAccepted()).false;

            // 2:1
            await proposal.setWeights(toBN(20).mul(b), toBN(10).mul(b), toBN(1).mul(b), toBN(50).mul(b));
            expect(await proposal.isObeyed()).false;
            expect(await proposal.isAccepted()).false;

            // < 10x
            await proposal.setWeights(toBN(7).mul(b), toBN(3).mul(b), toBN(1).mul(b), toBN(50).mul(b));
            expect(await proposal.isObeyed()).false;
            expect(await proposal.isAccepted()).false;
        });

        it('should proposeDrain() collect rejected', async () => {
            await s.token_abi.propose(target1, { value: enforce_fee });
            const enforce = await IProposal.at((await s.token_abi.proposals(target1))['0']);
            await enforce.voteAccept();
            await s.token_abi.proposeDrain(target1, { value: drain_fee });
            const drain = await IProposal.at((await s.token_abi.proposals(target1))['2']);

            expect(await s.token_abi.isDrainable(target1)).false;
            await drain.voteReject();
            expect(await s.token_abi.isDrainable(target1)).false;

            s.token_abi.collect(target1);
        });

        it('should proposeDrain()', async () => {
            await s.token_abi.proposeDrain(target1, { value: drain_fee });
            const drain = await IProposal.at((await s.token_abi.proposals(target1))['2']);

            expect(await s.token_abi.isDrainable(target1)).false;
            await drain.voteAccept();
            expect(await s.token_abi.isDrainable(target1)).true;
        });

        it('should refuse onDrain()', async () => {
            try {
                await s.token_abi.onDrain(target1);
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Not consensus/);
            }
        });

        it('should accept drainMigration()', async () => {
            try {
                await s.token_abi.drainMigration(0, target1);
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Invalid ID/);
            }
        });

        it('should refuse drainMigration()', async () => {
            try {
                await s.token_abi.drainMigration(0, target2);
                assert.fail('It should fail');
            } catch (e) {
                assert.match(e.message, /Not drainable/);
            }
        });
    });


    //---
    describe('StorageBlacklistRegistryV1', async () => {
        it ('should refuse setRevoke() from outside', async () => {
            try {
                await s.storage.setRevoke(s.fake.address, s.fake.address);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });

        it ('should refuse setEnforce() from outside', async () => {
            try {
                await s.storage.setEnforce(s.fake.address, s.fake.address);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });

        it ('should refuse remove() from outside', async () => {
            try {
                await s.storage.remove(s.fake.address);
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not owner/);
            }
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});

