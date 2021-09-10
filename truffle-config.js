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

const HDWalletProvider = require("@truffle/hdwallet-provider");
const mnemonicPhrase = process.env.TRUFFLE_MNEMONIC || "developers developers developers developers developers developers developers developers developers developers developers developers";
const mochaReporters = process.env.MOCHA_REPORTERS || "spec" // to get test output XML use MOCHA_REPORTERS="mocha-junit-reporter, spec"

const contracts = [
    'BackboneRewardV1',
    'BlacklistRegistryV1',
    'CheckpointRegistryV1',
    'Gen2Migration',
    'GenericProposal',
    'GovernedProxy',
    'HardforkRegistryV1',
    'MasternodeTokenV1',
    'MasternodeRegistryV1',
    'SporkRegistryV1',
    'StakerRewardV1',
    'TreasuryV2',
    'CheckpointRegistryV3',
    'StorageCheckpointRegistryV2',
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
    migrations_directory: './energi/contracts/migrations',
    test_directory: './energi/contracts/test',
    test_file_extension_regexp: /.*\.spec\.js$/,
    verboseRpc: false,
    mocha: {
        spec: './energi/contracts/test/*.spec.js',
        useColors: true,
        reporter: "mocha-multi-reporters",
        reporterOptions: {
            "reporterEnabled": mochaReporters,
            "mochaJunitReporterReporterOptions": {
                "mochaFile": ".test-sol-report.xml"
            }
        }
    },
    networks: {
        development: {
            host: "127.0.0.1",
            port: 8545,
            network_id: "*" // Match any network id
        },
        testnet: {
            provider: () =>
              new HDWalletProvider({
                mnemonic: {
                    phrase: mnemonicPhrase
                },
                providerOrUrl: "https://nodeapi.test.energi.network",
                derivationPath: "m/44'/49797'/0'/0/"
            }),
            network_id: "49797",
            gas: 30000000
        },
        mainnet: {
            provider: () =>
              new HDWalletProvider({
                mnemonic: {
                    phrase: mnemonicPhrase
                },
                providerOrUrl: "https://nodeapi.energi.network",
                derivationPath: "m/44'/39797'/0'/0/"
            }),
            network_id: "39797",
            gas: 30000000
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
