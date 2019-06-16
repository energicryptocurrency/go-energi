
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
            port: 8545,
            network_id: "*" // Match any network id
        }
    },
    compilers: {
        solc: {
            version: '0.5.9',
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
