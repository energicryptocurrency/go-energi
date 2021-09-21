// Copyright 2021 The Energi Core Authors
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
import { BudgetProposalV2 } from "./BudgetProposalV2.sol";
import { IBudgetProposal } from "./ITreasury.sol";
import { TreasuryV1, StorageTreasuryV1 }  from "./TreasuryV1.sol";

/**
 * Genesis hardcoded version of Treasury
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract TreasuryV2 is TreasuryV1 {
    constructor(address _proxy, IGovernedProxy _mnregistry_proxy, uint _superblock_cycle) public TreasuryV1(_proxy, _mnregistry_proxy, _superblock_cycle) {}

    // create a new budget proposal
    function propose(uint _amount, uint _ref_uuid, uint _period)
        external payable
        noReentry
        returns(IBudgetProposal proposal)
    {
        require(msg.value == FEE_BUDGET_V2 + (_amount * 2) / 1000, "Invalid fee");
        require(_amount >= BUDGET_AMOUNT_MIN, "Too small amount");
        require(_amount <= BUDGET_AMOUNT_MAX, "Too large amount");
        require(_period >= PERIOD_BUDGET_MIN, "Too small period");
        require(_period <= PERIOD_BUDGET_MAX, "Too large period");

        StorageTreasuryV1 store = v1storage;
        address payable payout_address = _callerAddress();

        require(address(store.uuid_proposal(_ref_uuid)) == address(0), "UUID in use");

        // Find, if proposal slot is available.
        for (uint i = 0; i < BUDGET_PROPOSAL_MAX; ++i) {
            if (address(active_proposals[i]) == address(0)) {
                proposal = new BudgetProposalV2(
                    mnregistry_proxy,
                    payout_address,
                    _ref_uuid,
                    _amount,
                    _period
                );

                active_proposals[i] = proposal;
                break;
            }
        }

        require(address(proposal) != address(0), "Too many active proposals");

        // set/pay required fee for creating proposal
        proposal.setFee.value(msg.value)();

        // store proposal to be active
        store.setProposal(_ref_uuid, proposal);

        // NOTE: it's the only way to retrieve proposal on regular transaction
        emit BudgetProposal(
            _ref_uuid,
            proposal,
            payout_address,
            _amount,
            proposal.deadline()
        );

        return proposal;
    }

}
