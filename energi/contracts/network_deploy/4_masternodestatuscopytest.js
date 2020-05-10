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

const MasternodeStatusCopyTest = artifacts.require('MasternodeStatusCopyTest');

module.exports = function(deployer, network) {
    try {
        var mn_registry = '0x0000000000000000000000000000000000000312';
        console.log("Deploying to " + network);

        if (network === "testnet") {
            mn_registry = '0x00f9ea265b3cd33fd693003fc6968ce8b5be5152';
        }

        deployer.deploy(MasternodeStatusCopyTest, mn_registry);
    } catch (e) {
        console.dir(e);
        throw e;
    }
};
