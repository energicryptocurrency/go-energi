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

import { GlobalConstants } from "./constants.sol";
import { IProposal } from "./IProposal.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IMasternodeRegistry } from "./IMasternodeRegistry.sol";
import { ITreasury } from "./ITreasury.sol";
import { StorageBase }  from "./StorageBase.sol";

// solium-disable security/no-block-members

/**
 * Genesis hardcoded version of GenericProposal V1
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract GenericProposalV1 is
    GlobalConstants,
    IProposal
{
    IGovernedProxy public mnregistry_proxy;

    address public parent;

    //! Block of proposal creation to check Masternode eligibility
    uint public created_block;

    //! Deadline for voting
    uint public deadline;

    //! The one who initiated the voting
    address payable public fee_payer;

    //! Fee amount
    uint public fee_amount;

    //! Weight of approval votes
    uint public accepted_weight;

    //! Weight of rejection votes
    uint public rejected_weight;

    //! Total masternode weight at the moment of proposal creation
    uint public total_weight;

    //! Weight of votes when the result is consider eligible
    uint public quorum_weight;

    //! Weight of votes when the voting can finalize before deadline
    uint public finish_weight;

    //! Registry of votes masternodes (vote once only)
    mapping(address => bool) public voted;

    /**
     * C-tor
     *
     * @param _mnregistry_proxy - IMasternodeRegistry proxy
     * @param _quorum - 1..100
     * @param _period - in seconds until deadline
     * @param _feePayer - the proposal initiator
     */
    constructor(
        IGovernedProxy _mnregistry_proxy,
        uint8 _quorum,
        uint _period,
        address payable _feePayer
    ) public {
        parent = msg.sender;
        created_block = block.number;

        mnregistry_proxy = _mnregistry_proxy;
        deadline = block.timestamp + _period;
        fee_payer = _feePayer;

        (
            ,
            ,
            uint _total_weight, // active_collaterel
            , // total_collateral
            uint _ever_weight
        ) = IMasternodeRegistry(address(_mnregistry_proxy.impl())).count();

        require(_ever_weight > 0, "Not ready for proposals");
        require(_total_weight >= (_ever_weight/2), "Active weight < 1/2 ever weight");
        require(_quorum >= QUORUM_MIN, "Quorum min");
        require(_quorum <= QUORUM_MAX, "Quorum max");

        total_weight = _total_weight;
        quorum_weight = _total_weight * _quorum / QUORUM_MAX;

        if (_quorum >= QUORUM_MAJORITY) {
            finish_weight = quorum_weight;
        } else {
            finish_weight = _total_weight * QUORUM_MAJORITY / QUORUM_MAX;
        }

        require(quorum_weight > 0, "Quorum weight");
        require(finish_weight > 0, "Finish weight");
    }

    /**
     * Check if the proposal is considered accepted.
     * NOTE: It can happen before the deadline.
     */
    function isAccepted() public view returns(bool) {
        // Before the deadline condition
        if (accepted_weight >= finish_weight) {
            return true;
        }

        // Ensure finish condition is reaches otherwise
        if (!isFinished()) {
            return false;
        }

        // Check quorum
        if ((accepted_weight + rejected_weight) < quorum_weight) {
            return false;
        }

        // Simply majority
        return accepted_weight > rejected_weight;
    }

    /**
     * Check finish condition
     */
    function isFinished() public view returns(bool) {
        return (
            (deadline <= block.timestamp) ||
            (accepted_weight >= finish_weight) ||
            (rejected_weight > finish_weight)
        );
    }

    function _voteCommon() internal returns(uint collateral) {
        // NOTE: do not use isFinished() to allow to accept votes before the deadline
        require(deadline > block.timestamp, "Finished");

        IMasternodeRegistry registry = IMasternodeRegistry(address(mnregistry_proxy.impl()));
        address owner = msg.sender;

        uint announced_block;
        (,,, collateral, announced_block,) = registry.ownerInfo(owner);
        require(announced_block < created_block, "Not eligible");
        require(!voted[owner], "Already voted");
        voted[owner] = true;
    }

    /**
     * Check if particular MN owner can vote
     */
    function canVote(address owner) external view returns(bool) {
        IMasternodeRegistry registry = IMasternodeRegistry(address(mnregistry_proxy.impl()));

        uint announced_block;
        (,,,, announced_block,) = registry.ownerInfo(owner);

        return (
            (deadline > block.timestamp) &&
            (announced_block < created_block) &&
            !voted[owner]
        );
    }

    /**
     * Masternode Owner approval vote
     */
    function voteAccept() external {
        accepted_weight += _voteCommon();
    }

    /**
     * Masternode Owner rejection vote
     */
    function voteReject() external {
        rejected_weight += _voteCommon();
    }

    /**
     * Withdrawal from accepted proposal.
     * NOTE: Usually for fee, but can be for budget as well.
     */
    function withdraw() external {
        // NOTE: anyone should be able to do that for cases when payer is a contract
        require(isAccepted(), "Not accepted");
        fee_payer.transfer(address(this).balance);
    }

    /**
     * Destruction via Governance logic.
     */
    function destroy() external {
        // NOTE: unfinished voting must get canceled
        require(msg.sender == parent, "Only parent");
        selfdestruct(fee_payer);
    }

    /**
     * Allow Treasury to collect the fee of rejected proposals.
     */
    function collect() external {
        require(isFinished() && !isAccepted(), "Not collectable");
        require(msg.sender == parent, "Only parent");

        IMasternodeRegistry registry = IMasternodeRegistry(address(mnregistry_proxy.impl()));
        ITreasury treasury = ITreasury(address(registry.treasury_proxy().impl()));

        treasury.contribute.value(address(this).balance)();
    }

    /**
     * Set fee amount by parent
     */
    function setFee() external payable {
        require(msg.sender == parent, "Only parent");
        // NOTE: make sure it correctly handles multiple calls
        fee_amount += msg.value;
    }

    /**
     * Only accept fee from the parent creating contract
     */
    function () external payable {
        revert("Not allowed");
    }
}
