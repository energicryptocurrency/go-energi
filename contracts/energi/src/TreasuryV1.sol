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

import { GlobalConstants } from "./constants.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IBlockReward } from "./IBlockReward.sol";
import { ITreasury, IProposal } from "./ITreasury.sol";
import { StorageBase }  from "./StorageBase.sol";

/**
 * Permanent storage of Treasury V1 data.
 */
contract StorageTreasuryV1 is
    StorageBase
{
    // NOTE: ABIEncoderV2 is not acceptable at the moment of development!
}

/**
 * Genesis hardcoded version of Treasury
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract TreasuryV1 is
    GlobalConstants,
    GovernedContract,
    IBlockReward,
    ITreasury
{
    // Data for migration
    //---------------------------------
    StorageTreasuryV1 public v1storage;
    uint public superblock_cycle;
    //---------------------------------

    constructor(address _proxy, uint _superblock_cycle)
        public
        GovernedContract(_proxy)
    {
        v1storage = new StorageTreasuryV1();
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
    function isSuperblock(uint _blockNumber)
        external view
        returns(bool)
    {
        return (_blockNumber % superblock_cycle) == 0;
    }

    function collect(IProposal proposal)
        external
    {
        // SECURITY: if this one is called as part of "zero-fee",
        //           then its parameter must be validated to avoid
        //           gas spending attacks.
        proposal.collect();
    }

    function contribute() external payable {
        if (msg.value > 0) {
            emit Contribution(msg.sender, msg.value);
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
    function reward() external payable {
    }

    function getReward(uint _blockNumber)
        external view
        returns(uint amount)
    {
        if (_blockNumber > 0) {
            amount = REWARD_TREASURY_V1;
        }
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}

