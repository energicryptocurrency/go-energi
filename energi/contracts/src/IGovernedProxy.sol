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
pragma solidity 0.5.10;
//pragma experimental SMTChecker;

import { IGovernedContract } from "./IGovernedContract.sol";
import { IUpgradeProposal } from "./IUpgradeProposal.sol";

/**
 * Genesis version of IGovernedProxy interface.
 *
 * Base Consensus interface for upgradable contracts proxy.
 * Unlike common approach, the implementation is NOT expected to be
 * called through delegatecall() to minimize risks of shared storage.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
interface IGovernedProxy {
    event UpgradeProposal(
        IGovernedContract indexed impl,
        IUpgradeProposal proposal
    );
    event Upgraded(
        IGovernedContract indexed impl,
        IUpgradeProposal proposal
    );

    function impl() external view returns(IGovernedContract);
    function proposeUpgrade(IGovernedContract _newImpl, uint _period)
        external payable returns(IUpgradeProposal);
    function upgrade(IUpgradeProposal _proposal) external;
    function upgradeProposalImpl(IUpgradeProposal _proposal) external view returns(IGovernedContract new_impl);
    function listUpgradeProposals() external view returns(IUpgradeProposal[] memory proposals);
    function collectUpgradeProposal(IUpgradeProposal _proposal) external;

    function () external payable;
}


