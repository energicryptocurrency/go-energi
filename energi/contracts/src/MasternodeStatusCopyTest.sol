// Copyright 2019-2020 The Energi Core Authors
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

import { NonReentrant } from "./NonReentrant.sol";
import { MasternodeRegistryV2 } from "./MasternodeRegistryV2.sol";

contract owned {
    constructor() public { owner = msg.sender; }
    address payable owner;

    // This contract only defines a modifier but does not use
    // it: it will be used in derived contracts.
    // The function body is inserted where the special symbol
    // `_;` in the definition of a modifier appears.
    // This means that if the owner calls this function, the
    // function is executed and otherwise, an exception is
    // thrown.
    modifier onlyOwner {
        require(
            msg.sender == owner,
            "Only owner can call this function."
        );
        _;
    }
}

contract destructible is owned {
    // This contract inherits the `onlyOwner` modifier from
    // `owned` and applies it to the `destroy` function, which
    // causes that calls to `destroy` only have an effect if
    // they are made by the stored owner.
    function destroy() public onlyOwner {
        selfdestruct(owner);
    }
}

contract MasternodeStatusCopyTest is destructible, NonReentrant {
    struct Status {
        uint256 sw_features;
        uint next_heartbeat;
        uint inactive_since;
        uint validator_index;
        uint invalidations;
        uint seq_payouts;
        uint last_vote_epoch;
    }

    mapping(address => Status) public mn_status;
    address[] public validator_list;
    MasternodeRegistryV2 public mn_registry;
    uint public max_mns;

    constructor(MasternodeRegistryV2 _mn_registry) public {
        mn_registry = _mn_registry;
        max_mns = 320; // 320 masternode statuses seems to hit the block gas limit of 40M
    }

    function setMax(uint max) public onlyOwner {
        max_mns = max;
    }

    function setRegistry(address payable _mn_registry) public onlyOwner {
        mn_registry = MasternodeRegistryV2(_mn_registry);
    }

    function copyStatus() public noReentry {
        uint mn_active = mn_registry.mn_active();

        // Restore the mn status information.
        // NOTE: Only active masternodes (validator_list()) are considered for migration
        // NOTE: this may be a serious gas consumption problem due to open limit.
        for (uint i = 0; i < mn_active; ++i) {
            if (i > max_mns) break;

            address mn = mn_registry.validator_list(i);
            Status memory status;
            (
                status.sw_features,
                status.next_heartbeat,
                status.inactive_since,
                status.validator_index,
                status.invalidations,
                status.seq_payouts,
                status.last_vote_epoch
            ) = mn_registry.mn_status(mn);

            validator_list.push(mn);
            mn_status[mn] = status;
        }
    }

    function copyStatusWithoutInvalidations() public noReentry {
        uint mn_active = mn_registry.mn_active();

        // Restore the mn status information.
        // NOTE: Only active masternodes (validator_list()) are considered for migration
        // NOTE: this may be a serious gas consumption problem due to open limit.
        for (uint i = 0; i < mn_active; ++i) {
            if (i > max_mns) break;

            address mn = mn_registry.validator_list(i);
            Status memory status;
            (
                ,
                status.next_heartbeat,
                status.inactive_since,
                status.validator_index,
                ,
                status.seq_payouts,
            ) = mn_registry.mn_status(mn);

            validator_list.push(mn);
            mn_status[mn] = status;
        }
    }

    function reset() public noReentry {
        // delete all of the mapping keys
        for (uint i = 0; i < validator_list.length; ++i) {
            delete mn_status[validator_list[i]];
        }
        // delete the validator list
        delete validator_list;
    }

    function statusCount() public view returns (uint) {
        return validator_list.length;
    }
}
