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
//pragma experimental SMTChecker;

import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { ISporkRegistry } from "./ISporkRegistry.sol";
import { GovernedProxy } from "./GovernedProxy.sol";

contract MockContract is GovernedContract
{
    constructor(address _proxy) public GovernedContract(_proxy) {}
    function migrate(IGovernedContract) external {}
    function destroy(IGovernedContract new_impl) external {
        selfdestruct(address(new_impl));
    }
    function getAddress() external view returns (address) {
        return address(this);
    }
    function () external payable {}
}

contract MockProxy is GovernedProxy
{
    constructor() public GovernedProxy(
        IGovernedContract(address(0)),
        ISporkRegistry(address(0))
    ) {}

    function setImpl(IGovernedContract _impl) external {
        current_impl = _impl;
    }

    function setSporkRegistry(ISporkRegistry _registry) external {
        spork_registry = _registry;
    }
}

