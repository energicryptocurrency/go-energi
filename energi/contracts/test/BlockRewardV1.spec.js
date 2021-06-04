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
const BlockRewardV1 = artifacts.require('BlockRewardV1');
const IBlockReward = artifacts.require('IBlockReward');

const common = require('./common');

contract("BlockRewardV1", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
        backbone : accounts[5],
    };

    before(async () => {
        s.orig = await BlockRewardV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await BlockRewardV1.at(s.proxy.address);
        s.token_abi = await IBlockReward.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        Object.freeze(s);
    });

    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('Primary', () => {
        const { toBN, toWei } = web3.utils;
        const backbone_reward = toBN(toWei('2.28', 'ether'));
        const normal_reward = toBN(toWei('13.7', 'ether'));
        const superblock_reward = toBN(toWei('184004.56', 'ether'));

        it('should correctly getReward()', async () => {
            expect(toBN(await s.token_abi.getReward(0)).toString())
                .equal(toBN(0).toString());
            expect(toBN(await s.token_abi.getReward(1)).toString())
                .equal(toBN(normal_reward).toString());
            expect(toBN(await s.token_abi.getReward(common.superblock_cycles)).toString())
                .equal(toBN(superblock_reward).toString());
        });

        it('should reward()', async () => {
            const bal_before = toBN(await web3.eth.getBalance(s.backbone));
            await s.token_abi.reward({value: superblock_reward});
            const bal_after = toBN(await web3.eth.getBalance(s.backbone));
            expect(bal_after.sub(bal_before).toString()).equal(backbone_reward.toString());
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});
