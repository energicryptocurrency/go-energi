// Copyright 2020 The Energi Core Authors
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
pragma solidity 0.5.16;
//pragma experimental SMTChecker;

import { IGovernedProxy } from "./IGovernedProxy.sol";
import { GovernedProxy } from "./GovernedProxy.sol";
import { IGovernedContract } from "./IGovernedContract.sol";
import { GlobalConstants } from "./constants.sol";

/**
 * GovernedContractV2 is a version of GovernedContract which deploys its own proxy.
 * This is useful to avoid a circular dependency between GovernedContract and GovernedProxy
 * wherein they need each other's address in the constructor.
 * This should only be used when deploying a contract which needs a new proxy, not a contract
 * for which a proxy already exists. If a proxy already exists, use GovernedContract
 */
contract GovernedContractAutoProxy is IGovernedContract, GlobalConstants {
    IGovernedProxy public proxy;

    constructor() public {
        proxy = new GovernedProxy(this, IGovernedProxy(SPORK_REGISTRY));
    }

    modifier requireProxy {
        require(msg.sender == address(proxy), "Not proxy");
        _;
    }

    function migrate(IGovernedContract _oldImpl) external requireProxy {
        _migrate(_oldImpl);
    }

    function destroy(IGovernedContract _newImpl) external requireProxy {
        _destroy(_newImpl);
        selfdestruct(address(_newImpl));
    }

    // solium-disable-next-line no-empty-blocks
    function _migrate(IGovernedContract) internal {}

    // solium-disable-next-line no-empty-blocks
    function _destroy(IGovernedContract) internal {}

    function _callerAddress()
        internal view
        returns (address payable)
    {
        if (msg.sender == address(proxy)) {
            // This is guarantee of the GovernedProxy
            // solium-disable-next-line security/no-tx-origin
            return tx.origin;
        } else {
            return msg.sender;
        }
    }
}
