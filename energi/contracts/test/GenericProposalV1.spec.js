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

const GenericProposalV1 = artifacts.require('GenericProposalV1');
const ITreasury = artifacts.require('ITreasury');
const MasternodeTokenV1 = artifacts.require('MasternodeTokenV1');
const MasternodeRegistryV1 = artifacts.require('MasternodeRegistryV1');

const common = require('./common');

contract("GenericProposalV1", async accounts => {
    let mntoken;
    let mnregistry;
    let treasury;

    before(async () => {
        const mntoken_orig = await MasternodeTokenV1.deployed();
        mntoken = await MasternodeTokenV1.at(await mntoken_orig.proxy());

        const mnregistry_orig = await MasternodeRegistryV1.deployed();
        mnregistry = await MasternodeRegistryV1.at(await mnregistry_orig.proxy());

        const treasury_proxy = await mnregistry.treasury_proxy();
        treasury = await ITreasury.at(treasury_proxy);
    });

    describe('Primary', () => {
        const { fromAscii, toBN, toWei } = web3.utils;

        const collateral1 = toBN(toWei('30000', 'ether'));
        const collateral2 = toBN(toWei('20000', 'ether'));
        const collateral3 = toBN(toWei('10000', 'ether'));

        const owner1 = accounts[0];
        const owner2 = accounts[1];
        const owner3 = accounts[2];
        const not_owner = accounts[3];

        const masternode1 = accounts[9];
        const masternode2 = accounts[8];
        const masternode3 = accounts[7];

        const ip1 = toBN(0x12345678);
        const ip2 = toBN(0x87654321);
        const ip3 = toBN(0x43218765);

        const enode_common = '123456789012345678901234567890'
        const enode1 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '11')];
        const enode2 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '22')];
        const enode3 = [fromAscii(enode_common + '11'), fromAscii(enode_common + '33')];

        before(async () => {
            await mntoken.depositCollateral({
                from: owner1,
                value: collateral1,
            });
            await mntoken.depositCollateral({
                from: owner2,
                value: collateral2,
            });
            await mntoken.depositCollateral({
                from: owner3,
                value: collateral3,
            });

            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
        });

        after(async () => {
            await mntoken.withdrawCollateral(collateral1, {
                from: owner1,
            });
            await mntoken.withdrawCollateral(collateral2, {
                from: owner2,
            });
            await mntoken.withdrawCollateral(collateral3, {
                from: owner3,
            });
        });

        it('should refuse wrong quorum percent', async () => {
            try {
                await GenericProposalV1.new(
                    mnregistry.address,
                    0,
                    60,
                    not_owner
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Quorum min/);
            }

            await GenericProposalV1.new(
                mnregistry.address,
                1,
                60,
                not_owner
            );

            await GenericProposalV1.new(
                mnregistry.address,
                100,
                60,
                not_owner
            );

            try {
                await GenericProposalV1.new(
                    mnregistry.address,
                    101,
                    60,
                    not_owner
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Quorum max/);
            }
        });

        it('should refuse on lack of quorum', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
            mnregistry.announce(masternode2, ip2, enode2, {from: owner2});
            mnregistry.announce(masternode3, ip3, enode3, {from: owner3});
            mnregistry.denounce(masternode1, {from: owner1});
            mnregistry.denounce(masternode2, {from: owner2});

            try {
                await GenericProposalV1.new(
                    mnregistry.address,
                    51,
                    60,
                    not_owner
                );
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Active weight < 1\/2 ever weight/);
            }
        });
        
        it('should accept half-way', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
            mnregistry.announce(masternode2, ip2, enode2, {from: owner2});
            mnregistry.announce(masternode3, ip3, enode3, {from: owner3});

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                51,
                60,
                not_owner
            );

            expect(await proposal.canVote(owner1)).true;
            await proposal.voteAccept({from: owner1});
            expect(await proposal.canVote(owner1)).false;
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.toString());

            expect(await proposal.canVote(owner3)).true;
            await proposal.voteAccept({from: owner3});
            expect(await proposal.canVote(owner3)).false;
            expect(await proposal.isAccepted()).true;
            expect(await proposal.isFinished()).true;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.add(collateral3).toString());

            expect(await proposal.canVote(owner2)).true;
            await proposal.voteAccept({from: owner2});
            expect(await proposal.canVote(owner2)).false;
            expect(await proposal.isAccepted()).true;
            expect(await proposal.isFinished()).true;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.add(collateral2).add(collateral3).toString());
        });

        it('should reject half-way', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
            mnregistry.announce(masternode2, ip2, enode2, {from: owner2});
            mnregistry.announce(masternode3, ip3, enode3, {from: owner3});

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                51,
                60,
                not_owner
            );

            await proposal.voteReject({from: owner1});
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral1.toString());
            
            await proposal.voteReject({from: owner3});
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).true;
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral1.add(collateral3).toString());

            await proposal.voteReject({from: owner2});
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).true;
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral1.add(collateral2).add(collateral3).toString());
        });
        
        it('should not accept half-way', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
            mnregistry.announce(masternode2, ip2, enode2, {from: owner2});
            mnregistry.announce(masternode3, ip3, enode3, {from: owner3});

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                51,
                60,
                not_owner
            );

            expect(await proposal.canVote(owner1)).true;
            await proposal.voteAccept({from: owner1});
            expect(await proposal.canVote(owner1)).false;
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.toString());

            expect(await proposal.canVote(owner3)).true;
            await proposal.voteReject({from: owner3});
            expect(await proposal.canVote(owner3)).false;
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.toString());
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral3.toString());

            await proposal.voteAccept({from: owner2});
            expect(await proposal.isAccepted()).true;
            expect(await proposal.isFinished()).true;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.add(collateral2).toString());
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral3.toString());

        });

        it('should not accept half-way and refuse vote after deadline', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
            mnregistry.announce(masternode2, ip2, enode2, {from: owner2});
            mnregistry.announce(masternode3, ip3, enode3, {from: owner3});

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                5,
                60,
                not_owner
            );

            expect(await proposal.canVote(owner1)).true;
            await proposal.voteAccept({from: owner1});
            expect(await proposal.canVote(owner1)).false;
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.toString());

            expect(await proposal.canVote(owner3)).true;
            await proposal.voteReject({from: owner3});
            expect(await proposal.canVote(owner3)).false;
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.toString());
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral3.toString());

            expect(await proposal.canVote(owner2)).true;
            await common.moveTime(web3, 70);
            expect(await proposal.canVote(owner2)).false;

            try {
                await proposal.voteAccept({from: owner2});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Finished/);
            }

            expect(await proposal.isAccepted()).true;
            expect(await proposal.isFinished()).true;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.toString());
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral3.toString());

        });

        it('should not accept half-way and refuse by quorum', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
            mnregistry.announce(masternode2, ip2, enode2, {from: owner2});
            mnregistry.announce(masternode3, ip3, enode3, {from: owner3});

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                90,
                60,
                not_owner
            );

            expect(await proposal.canVote(owner1)).true;
            await proposal.voteAccept({from: owner1});
            expect(await proposal.canVote(owner1)).false;
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.toString());

            expect(await proposal.canVote(owner3)).true;
            await proposal.voteReject({from: owner3});
            expect(await proposal.canVote(owner3)).false;
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.toString());
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral3.toString());

            expect(await proposal.canVote(owner2)).true;
            await proposal.voteAccept({from: owner2});
            expect(await proposal.canVote(owner2)).false;
            expect(await proposal.isAccepted()).false;
            expect(await proposal.isFinished()).false;
            await common.moveTime(web3, 70);
            expect(await proposal.isFinished()).true;
            expect((await proposal.accepted_weight()).toString())
                .eql(collateral1.add(collateral2).toString());
            expect((await proposal.rejected_weight()).toString())
                .eql(collateral3.toString());
        });

        it('should refuse vote of not eligible Masternode', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                51,
                60,
                not_owner
            );

            mnregistry.announce(masternode2, ip2, enode2, {from: owner2});
            expect(await proposal.canVote(owner2)).false;

            try {
                await proposal.voteAccept({from: owner2});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not eligible/);
            }

            expect(await proposal.canVote(owner1)).true;
            await proposal.voteAccept({from: owner1});
            expect(await proposal.canVote(owner1)).false;
        });

        it('should refuse already voted Masternode', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                51,
                60,
                not_owner
            );

            expect(await proposal.canVote(owner1)).true;
            await proposal.voteAccept({from: owner1});
            expect(await proposal.canVote(owner1)).false;

            try {
                await proposal.voteAccept({from: owner1});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Already voted/);
            }
        });

        it('should refuse withdraw() unless accepted', async () => {
            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                51,
                60,
                not_owner
            );

            try {
                await proposal.withdraw({from: not_owner});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not accepted/);
            }
        });

        it('should withdraw()', async () => {
            const bal_before = await web3.eth.getBalance(not_owner);
            
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                1,
                60,
                not_owner
            );
            await proposal.voteAccept({from: owner1});
            await common.moveTime(web3, 70);

            await proposal.setFee({ from: owner1, value: toWei('2', 'ether')});
            await proposal.setFee({ from: owner1, value: toWei('3', 'ether')});
            
            await proposal.withdraw({from: owner1});
            const bal_after = await web3.eth.getBalance(not_owner);
            expect(toBN(bal_after).sub(toBN(bal_before)).toString())
                .equal(toBN(toWei('5', 'ether')).toString());
        });

        it('should refuse destroy() unless parent', async () => {
            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                1,
                60,
                not_owner
            );

            try {
                await proposal.destroy({from: not_owner});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Only parent/);
            }
        });

        it('should destroy()', async () => {
            const bal_before = await web3.eth.getBalance(not_owner);
            
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});
            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                1,
                60,
                not_owner
            );
            await proposal.voteAccept({from: owner1});
            await common.moveTime(web3, 70);

            await proposal.setFee({ from: owner1, value: toWei('5', 'ether') });
            
            await proposal.destroy({from: owner1});
            const bal_after = await web3.eth.getBalance(not_owner);
            expect(toBN(bal_after).sub(toBN(bal_before)).toString())
                .equal(toBN(toWei('5', 'ether')).toString());
        });

        it('should refuse collect() unless rejected', async () => {
            mnregistry.announce(masternode1, ip1, enode1, {from: owner1});

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                1,
                60,
                not_owner
            );

            try {
                await proposal.collect({from: not_owner});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not collectable/);
            }

            await proposal.voteAccept({from: owner1});
            await common.moveTime(web3, 70);

            try {
                await proposal.collect({from: not_owner});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not collectable/);
            }
        });

        it('should collect()', async () => {
            expect(await mnregistry.treasury_proxy()).equal(treasury.address);

            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                1,
                60,
                not_owner
            );

            await proposal.setFee({ from: owner1, value: toWei('5', 'ether') });

            await common.moveTime(web3, 70);

            const bal_before = await treasury.balance();
            
            try {
                await proposal.collect({from: not_owner});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Only parent/);
            }

            await proposal.collect();

            const bal_after = await treasury.balance();
            expect(toBN(bal_after).sub(toBN(bal_before)).toString())
                .equal(toBN(toWei('5', 'ether')).toString());
        });
        
        it('should refuse payments', async () => {
            const proposal = await GenericProposalV1.new(
                mnregistry.address,
                1,
                60,
                not_owner
            );

            try {
                await proposal.send(toWei('1', 'ether'), {from: not_owner});
                assert.fail('It must fail');
            } catch (e) {
                assert.match(e.message, /Not allowed/);
            }
        });
    });
});
