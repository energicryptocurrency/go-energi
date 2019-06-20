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
 * Interface of BudgetProposal
 */
interface IBudgetProposal {
    function payout_address() external view returns(address payable);

    function distributePayout() external payable;

    function budgetStatus() external view returns(
        uint ref_uuid,
        bool is_accepted,
        bool is_finished,
        uint unpaid
    );
}

/**
 * Interface for the Treasury
 */
interface ITreasury {
    event BudgetProposal(
        uint indexed ref_uuid,
        IProposal proposal,
        address payout_address,
        uint amount,
        uint deadline
    );
    event Contribution(
        address from,
        uint amount
    );
    event Payout(
        uint indexed ref_uuid,
        IProposal proposal,
        uint amount
    );

    function uuid_proposal(uint _ref_uuid) external view returns(IProposal);
    function proposal_uuid(IProposal proposal) external view returns(uint);
    function propose(uint _amount, uint _ref_uuid, uint _period)
        external payable returns(IProposal proposal);
    function isSuperblock(uint _blockNumber) external view returns(bool);
    function contribute() external payable;
    function balance() external view returns(uint amount);
}
