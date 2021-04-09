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

import { IGovernedContract } from "./IGovernedContract.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IUpgradeProposal } from "./IUpgradeProposal.sol";
import { SporkRegistryV1, UpgradeProposalV1 } from "./SporkRegistryV1.sol";

// solium-disable security/no-block-members
// solium-disable no-empty-blocks

/**
 * A special proposal for emergency consensus upgrades.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract EmergencyProposal is
    UpgradeProposalV1
{
    constructor(
        address _parent,
        IGovernedContract _impl,
        IGovernedProxy _mnregistry_proxy,
        uint _period,
        address payable _feePayer
    )
        public
        UpgradeProposalV1(
            _parent,
            _impl,
            _mnregistry_proxy,
            _period,
            _feePayer
        )
    {}

    function isAccepted() public view returns(bool) {
        return true;
    }

    function isFinished() public view returns(bool) {
        return true;
    }

    function voteAccept() external {
    }

    function voteReject() external {
    }
}


/**
 * Genesis hardcoded version of SporkRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract SporkRegistryV2 is
    SporkRegistryV1
{
    // Data for migration
    //---------------------------------
    address public Emergency_signer;
    //---------------------------------

    // IGovernedContract
    //---------------------------------
    constructor(address _proxy, IGovernedProxy _mnregistry_proxy,  address _emergency_signer)
        public SporkRegistryV1(_proxy, _mnregistry_proxy)
    {
        Emergency_signer = _emergency_signer;
    }

    // ISporkRegistry
    //---------------------------------
    function createUpgradeProposal(IGovernedContract _impl, uint _period, address payable _fee_payer)
        external payable
        returns (IUpgradeProposal proposal)
    {
        if (Emergency_signer == _fee_payer) {
            require(msg.value == 0, "Invalid fee");
            require(_period == 0, "Invalid period");

            return new EmergencyProposal(
                msg.sender,
                _impl,
                mnregistry_proxy,
                _period,
                _fee_payer
            );
        }

        // Parent has external method, unfortunately.
        //return super.createUpgradeProposal(_impl, _period, _fee_payer);

        require(msg.value == FEE_UPGRADE_V1, "Invalid fee");
        require(_period >= PERIOD_UPGRADE_MIN, "Period min");
        require(_period <= PERIOD_UPGRADE_MAX, "Period max");

        proposal = new UpgradeProposalV1(
            msg.sender,
            _impl,
            mnregistry_proxy,
            _period,
            _fee_payer
        );

        proposal.setFee.value(msg.value)();
    }
}
