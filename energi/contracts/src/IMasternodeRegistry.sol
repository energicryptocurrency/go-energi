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

/**
 * Genesis version of MasternodeRegistry interface.
 *
 * Base Consensus interface for masternodes.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
interface IMasternodeRegistry {
    event Announced(
        address indexed masternode,
        address indexed owner,
        uint32 ipv4address,
        bytes32[2] enode,
        uint collateral
    );
    event Denounced(
        address indexed masternode,
        address indexed owner
    );
    event Invalidation(
        address indexed masternode,
        address indexed validator
    );
    event Deactivated(
        address indexed masternode
    );

    function token_proxy() external view returns(IGovernedProxy);
    function treasury_proxy() external view returns(IGovernedProxy);

    function announce(address masternode, uint32 ipv4address, bytes32[2] calldata enode) external;
    function denounce(address masternode) external;
    function heartbeat(uint block_number, bytes32 block_hash, uint sw_features) external;
    function invalidate(address masternode) external;
    function validationTarget(address masternode) external view returns(address target);
    function isActive(address masternode) external view
        returns(bool);
    function count() external view
        returns(
            uint active, uint total,
            uint active_collateral, uint total_collateral,
            uint max_of_all_times);
    function info(address masternode) external view
        returns(address owner, uint32 ipv4address, bytes32[2] memory enode,
                uint collateral, uint announced_block);
    function ownerInfo(address owner) external view
        returns(address masternode, uint32 ipv4address, bytes32[2] memory enode,
                uint collateral, uint announced_block);
    function onCollateralUpdate(address owner) external;
    function enumerate() external view returns(address[] memory masternodes);
    function enumerateActive() external view returns(address[] memory masternodes);
}
