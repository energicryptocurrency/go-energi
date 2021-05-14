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
const StakerRewardV1 = artifacts.require('StakerRewardV1');
const IBlockReward = artifacts.require('IBlockReward');

const common = require('./common');

contract("StakerRewardV1", async accounts => {
    const s = {
        artifacts,
        accounts,
        assert,
        it,
        web3,
    };

    before(async () => {
        s.orig = await StakerRewardV1.deployed();
        s.proxy = await MockProxy.at(await s.orig.proxy());
        s.fake = await MockContract.new(s.proxy.address);
        s.proxy_abi = await StakerRewardV1.at(s.proxy.address);
        s.token_abi = await IBlockReward.at(s.proxy.address);
        await s.proxy.setImpl(s.orig.address);
        Object.freeze(s);
    });

    after(async () => {
        const impl = await StakerRewardV1.new(s.proxy.address);
        await s.proxy.setImpl(impl.address);
    });

    describe('common pre', () => common.govPreTests(s) );

    //---
    describe('Primary', () => {
        const { toBN, toWei } = web3.utils;
        const reward = toBN(toWei('2.28', 'ether'));

        it('should correctly getReward()', async () => {
            expect(toBN(await s.token_abi.getReward(0)).toString())
                .equal(toBN(0).toString());
            expect(toBN(await s.token_abi.getReward(1)).toString())
                .equal(toBN(reward).toString());
            expect(toBN(await s.token_abi.getReward(common.superblock_cycles)).toString())
                .equal(toBN(reward).toString());
        });

        it('should reward()', async () => {
            const coinbase = (await web3.eth.getBlock('latest')).miner;
            const trunc = toBN(toWei('0.01', 'ether')); // truncate gas
            const eth_reward = toBN(toWei('2', 'ether')); // standard reward

            const bal_before = toBN(await web3.eth.getBalance(coinbase));
            await s.token_abi.reward({value: reward, from: accounts[1]});
            const bal_after = toBN(await web3.eth.getBalance(coinbase));

            expect(bal_after.sub(bal_before).div(trunc).mul(trunc).toString())
                .equal(reward.add(eth_reward).toString());
        });
    });

    //---
    describe('common post', () => common.govPostTests(s) );
});
