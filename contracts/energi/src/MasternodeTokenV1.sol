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
import { IMasternodeToken } from "./IMasternodeToken.sol";

/**
 * Genesis hardcoded version of MasternodeToken
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract MasternodeTokenV1 is
    GlobalConstants,
    GovernedContract,
    IMasternodeToken
{
    constructor(address _proxy) public GovernedContract(_proxy) {
        // ERC20
        emit Transfer(address(0), address(0), 0);
    }

    // IGovernedContract
    //---------------------------------
    // solium-disable-next-line no-empty-blocks
    function migrate(IGovernedContract) external requireProxy {
        // pass
    }

    function destroy(IGovernedContract _newImpl) external requireProxy {
        selfdestruct(address(_newImpl));
    }

    // MasternodeTokenV1
    //---------------------------------
    uint256 public totalSupply;
    mapping(address => uint256) public owners;

    function name() external view returns (string memory) {
        return "Masternode Collateral";
    }

    function symbol() external view returns (string memory) {
        return "MNGR";
    }

    function decimals() external view returns (uint8) {
        return 18 + 4;
    }

    function balanceOf(address _owner) external view returns (uint256) {
        return owners[_owner];
    }

    function transfer(address, uint256) external returns (bool) {
        revert("Not allowed");
    }

    function transferFrom(address, address, uint256) external returns (bool) {
        revert("Not allowed");
    }

    function approve(address, uint256) external returns (bool) {
        revert("Not allowed");
    }

    function allowance(address, address) external view returns (uint256) {
        return 0;
    }

    // Energi
    //---------------------------------

    // solium-disable security/no-block-members

    struct Balance {
        uint256 amount;
        uint256 last_change;
    }
    mapping(address => Balance) public balances;

    function balanceInfo(address _tokenOwner)
        external view
        returns (uint256 balance, uint256 age)
    {
        Balance storage item = balances[_tokenOwner];

        assert(block.timestamp >= item.last_change);

        balance = item.amount;
        age = block.timestamp - item.last_change;
    }

    function withdrawCollateral(uint256 _amount) external {
        address payable owner = _ownerAddress();
        Balance storage item = balances[owner];
        uint256 balance = item.amount;

        if (balance < _amount) {
            revert("Not enough");
        }

        balance -= _amount;
        totalSupply -= _amount;

        _validateBalance(balance);

        item.amount = balance;
        item.last_change = block.timestamp;

        emit Transfer(owner, address(0), _amount);

        owner.transfer(_amount);
    }

    function depositCollateral() external payable {
        address payable owner = _ownerAddress();
        Balance storage item = balances[owner];
        uint256 balance = item.amount + msg.value;

        _validateBalance(balance);

        item.amount = balance;
        item.last_change = block.timestamp;
        totalSupply += msg.value;

        emit Transfer(address(0), owner, msg.value);
    }

    function _validateBalance(uint256 _amount) internal pure {
        // NOTE: "Too small" check makes no sense as it would be just zero.

        if (_amount > MN_COLLATERAL_MAX) {
            revert("Too much");
        }

        if ((_amount % MN_COLLATERAL_MIN) != 0) {
            revert("Not a multiple");
        }
    }

    function _ownerAddress()
        internal view
        returns (address payable)
    {
        if (msg.sender == proxy) {
            // This is guarantee of the GovernedProxy
            // solium-disable-next-line security/no-tx-origin
            return tx.origin;
        } else {
            return msg.sender;
        }
    }

    function () external payable {
        revert("Not supported");
    }

}
