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
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IBlockReward } from "./IBlockReward.sol";
import { ITreasury, IBudgetProposal } from "./ITreasury.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { BudgetProposalV1 } from "./BudgetProposalV1.sol";
import { NonReentrant } from "./NonReentrant.sol";
import { StorageBase }  from "./StorageBase.sol";

/**
 * Permanent storage of Treasury V1 data.
 */
contract StorageTreasuryV1 is
    StorageBase
{
    mapping(uint => IBudgetProposal) public uuid_proposal;
    mapping(address => uint) public proposal_uuid;

    function setProposal(uint _uuid, IBudgetProposal _proposal)
        external
        requireOwner
    {
        uuid_proposal[_uuid] = _proposal;
        proposal_uuid[address(_proposal)] = _uuid;
    }

    function deleteProposal(IBudgetProposal _proposal)
        external
        requireOwner
    {
        uint uuid = proposal_uuid[address(_proposal)];
        delete proposal_uuid[address(_proposal)];
        delete uuid_proposal[uuid];
    }
}

/**
 * Genesis hardcoded version of Treasury
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract TreasuryV2 is
    GlobalConstants,
    GovernedContract,
    NonReentrant,
    IBlockReward,
    ITreasury
{
    // Data for migration
    //---------------------------------
    StorageTreasuryV1 public v1storage;
    IGovernedProxy public mnregistry_proxy;
    uint public superblock_cycle;
    IBudgetProposal[BUDGET_PROPOSAL_MAX] public active_proposals;
    //---------------------------------

    constructor(address _proxy, IGovernedProxy _mnregistry_proxy, uint _superblock_cycle)
        public
        GovernedContract(_proxy)
    {
        v1storage = new StorageTreasuryV1();
        mnregistry_proxy = _mnregistry_proxy;
        superblock_cycle = _superblock_cycle;
        assert(superblock_cycle > 0);
    }

    // IGovernedContract
    //---------------------------------
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // ITreasury
    //---------------------------------
    function uuid_proposal(uint _ref_uuid) external view returns(IBudgetProposal) {
        return v1storage.uuid_proposal(_ref_uuid);
    }

    function proposal_uuid(IBudgetProposal proposal) external view returns(uint) {
        return v1storage.proposal_uuid(address(proposal));
    }

    function propose(uint _amount, uint _ref_uuid, uint _period)
        external payable
        noReentry
        returns(IBudgetProposal proposal)
    {
        require(msg.value == FEE_BUDGET_V2 * 2 / 1000, "Invalid fee");
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
                proposal = new BudgetProposalV1(
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
        //---

        proposal.setFee.value(msg.value)();
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

    function listProposals() external view returns(IBudgetProposal[] memory proposals) {
        IBudgetProposal[] memory tmp = new IBudgetProposal[](BUDGET_PROPOSAL_MAX);
        uint tmp_len = 0;

        for (uint i = 0; i < BUDGET_PROPOSAL_MAX; ++i) {
            IBudgetProposal p = active_proposals[i];

            if (address(p) != address(0)) {
                tmp[tmp_len++] = p;
            }
        }

        proposals = new IBudgetProposal[](tmp_len);

        for (uint i = 0; i < tmp_len; ++i) {
            proposals[i] = tmp[i];
        }

        return proposals;
    }

    function isSuperblock(uint _blockNumber)
        public view
        returns(bool)
    {
        return (_blockNumber % superblock_cycle) == 0 && (_blockNumber > 0);
    }

    function contribute() external payable {
        if (msg.value > 0) {
            emit Contribution(_callerAddress(), msg.value);
        }
    }

    // NOTE: usually Treasury is behind proxy and this one
    //       minimizes possible errors.
    function balance()
        external view
        returns(uint amount)
    {
        return address(this).balance;
    }

    // IBlockReward
    //---------------------------------
    struct AcceptedProposal {
        IBudgetProposal proposal;
        uint ref_uuid;
        uint unpaid;
    }

    function reward()
        external payable
        noReentry
    {
        AcceptedProposal[BUDGET_PROPOSAL_MAX] memory accepted;

        uint unpaid_total = _reward_status(accepted);
        uint curr_balance = address(this).balance;

        if ((curr_balance > 0) && (unpaid_total > 0)) {
            uint permille = 1000;

            if (unpaid_total > curr_balance) {
                // Due to lack of floating-point precision,
                // it may require a few blocks to process
                // full payouts.
                permille = curr_balance * 1000 / unpaid_total;
            }

            _reward_distribute(permille, accepted);
        }
    }

    function _reward_status(AcceptedProposal[BUDGET_PROPOSAL_MAX] memory accepted)
        internal
        returns(uint unpaid_total)
    {
        IBudgetProposal proposal;
        uint ref_uuid;
        bool is_accepted;
        bool is_finished;
        uint unpaid = 0;

        for (uint i = 0; i < BUDGET_PROPOSAL_MAX; ++i) {
            proposal = active_proposals[i];

            if (address(proposal) != address(0)) {
                (ref_uuid, is_accepted, is_finished, unpaid) = proposal.budgetStatus();

                if (is_accepted) {
                    if (unpaid > 0) {
                        unpaid_total += unpaid;
                        accepted[i].proposal = proposal;
                        accepted[i].ref_uuid = ref_uuid;
                        accepted[i].unpaid = unpaid;
                    } else {
                        // Fulfilled
                        proposal.destroy();
                        delete active_proposals[i];
                    }
                } else if (is_finished) {
                    // Rejected
                    proposal.collect();
                    delete active_proposals[i];
                }
            }
        }
    }

    function _reward_distribute(
        uint permille,
        AcceptedProposal[BUDGET_PROPOSAL_MAX] memory accepted
    )
        internal
    {
        IBudgetProposal proposal;

        for (uint i = 0; i < BUDGET_PROPOSAL_MAX; ++i) {
            proposal = accepted[i].proposal;

            if (address(proposal) != address(0)) {
                uint amount = accepted[i].unpaid * permille / 1000;
                assert(amount <= accepted[i].unpaid);
                proposal.distributePayout.value(amount)();
                emit Payout(
                    accepted[i].ref_uuid,
                    proposal,
                    amount
                );
            }
        }
    }

    function getReward(uint _blockNumber)
        external view
        returns(uint amount)
    {
        if (isSuperblock(_blockNumber)) {
            amount = REWARD_TREASURY_V1;
        }
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
