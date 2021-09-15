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
pragma solidity 0.5.16;
//pragma experimental SMTChecker;

import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IBudgetProposal } from "./ITreasury.sol";
import { GenericProposalV2 } from "./GenericProposalV2.sol";
import { GlobalConstants } from "./constants.sol";

/**
 * Budget Proposal V2 for Treasury distribution
 */
contract BudgetProposalV2 is
    GenericProposalV2,
    IBudgetProposal
{
    uint public paid_amount;
    uint public proposed_amount;
    uint public ref_uuid;

    constructor(
        IGovernedProxy _mnregistry_proxy,
        address payable _payout_address,
        uint _ref_uuid,
        uint _proposed_amount,
        uint _period
    )
        public
        GenericProposalV2(
            _mnregistry_proxy,
            GlobalConstants.QUORUM,
            _period,
            GlobalConstants.QUORUM_SUPERMAJORITY,
            _payout_address
        )
    {
        ref_uuid = _ref_uuid;
        proposed_amount = _proposed_amount;
    }

    // IBudgetProposal
    //---------------------------------

    // Just an alias
    function payout_address()
        external view
        returns(address payable)
    {
        return fee_payer;
    }

    // Called by Treasury on reward()
    function distributePayout() external payable {
        paid_amount += msg.value;
        assert(paid_amount <= proposed_amount);
    }

    // Optimized status retrieval in single call
    function budgetStatus()
        external view
        returns(
            uint uuid,
            bool is_accepted,
            bool is_finished,
            uint unpaid
        )
    {
        uuid = ref_uuid;
        is_accepted = isAccepted();
        is_finished = isFinished();
        assert(paid_amount <= proposed_amount);
        unpaid = proposed_amount - paid_amount;
    }
}
