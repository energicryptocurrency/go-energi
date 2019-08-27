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

import { BlacklistProposalV1 } from "./BlacklistRegistryV1.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";

contract MockBlacklistProposalV1
    is BlacklistProposalV1
{
    constructor(IGovernedProxy _mnregistry_proxy, address payable fee_payer)
        public
        BlacklistProposalV1(_mnregistry_proxy, fee_payer)
    // solium-disable-next-line no-empty-blocks
    {}

    function setWeights(uint accepted, uint rejected, uint quorum, uint finish) external {
        accepted_weight = accepted;
        rejected_weight = rejected;
        quorum_weight = quorum;
        finish_weight = finish;
    }
}
