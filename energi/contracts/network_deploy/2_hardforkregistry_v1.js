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

'use strict';

const HardforkRegistryV1 = artifacts.require('HardforkRegistryV1');
const common = require('../test/common');

module.exports = function(deployer, network) {
  try {
    var hf_signer = common.hf_signer;
    var hf_finalization_period = common.hf_finalization_period;

    console.log("Deploying to " + network);

    if (network === "mainnet") {
      hf_signer = '0x44D16E845ec2d2D6A99a10fe44EE99DA0541CF31';
      hf_finalization_period = 30;
    } else if (network === "testnet") {
      hf_signer = '0x5b00118464fa6e73f9c2a4ea44e1cbfa9f5b83c6';
      hf_finalization_period = 10;
    }

    deployer.deploy(HardforkRegistryV1, hf_signer, hf_finalization_period);
  } catch (e) {
    console.dir(e);
    throw e;
  }
};
