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

// NOTE: It's not allowed to change the compiler due to byte-to-byte
//       match requirement.
pragma solidity 0.5.9;

import { Gen2Migration } from "./Gen2Migration.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";

contract MockGen2Migration is Gen2Migration {
    constructor(IGovernedProxy _treasury_proxy, uint _chain_id) public
        Gen2Migration(_treasury_proxy, _chain_id)
    // solium-disable-next-line no-empty-blocks
    {
    }

    function setCoins(bytes20[] calldata _owners, uint[] calldata _amounts) external payable {
        require(_owners.length == _amounts.length, "match length");
        require(_owners.length > 0, "has data");

        coins.length = _owners.length;

        for (uint i = _owners.length; i-- > 0;) {
            coins[i].owner = _owners[i];
            coins[i].amount = _amounts[i];
        }
    }
}
