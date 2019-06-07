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

import { IProposal } from "./IProposal.sol";

/**
 * Genesis hardcoded version of GenericProposal
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract GenericProposalV1 is IProposal {
    uint public fee;
    uint public deadline;
    address payable public fee_payer;
    uint8 public quorum;

    constructor(
        uint8 _quorum,
        uint _period,
        address payable _fee_payer,
        uint _fee
    ) public {
        fee = _fee;
        // solium-disable-next-line security/no-block-members
        deadline = block.timestamp + _period;
        fee_payer = _fee_payer;
        quorum = _quorum;
    }

    function isAccepted() external view returns(bool) {
        return false;
    }

    function () external payable {}
}
