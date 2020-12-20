// Copyright 2019 The Energi Core Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.


const contracts = [
    'BackboneRewardV1',
    'BlacklistRegistryV1',
    'CheckpointRegistryV1',
    'Gen2Migration',
    'GenericProposal',
    'GovernedProxy',
    'MasternodeTokenV1',
    'MasternodeRegistryV1',
    'SporkRegistryV1',
    'StakerRewardV1',
    'TreasuryV1',
];
const targets = [];

for (let c of contracts) {
    targets.push({
        properties: {
            contractName: c,
        },
        fileProperties: {
            abi: `./build/contracts/energi/${c}.abi`,
            bytecode: `./build/contracts/energi/${c}_evm.json`,
        }
    });
}

module.exports = {
    contracts_directory: './energi/contracts/src',
    contracts_build_directory: './build/contracts/truffle',
    migrations_directory: './energi/contracts/test_deploy',
    test_directory: './energi/contracts/test',
    test_file_extension_regexp: /.*\.spec\.js$/,
    verboseRpc: false,
    mocha: {
        spec: './energi/contracts/test/*.spec.js',
    },
    networks: {
        development: {
            host: "127.0.0.1",
            port: 7545,
            network_id: "*" // Match any network id
        }
    },
    compilers: {
        solc: {
            version: '0.5.16',
            evmVersion: 'petersburg',
            optimizer: {
                enabled: true,
                runs: 9999999999,
            }
        },
        /*external: {
            command: "make prebuild",
            targets,
        }*/
    }
}
