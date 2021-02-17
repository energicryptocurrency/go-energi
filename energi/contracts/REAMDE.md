# Energi Governance Contracts
This directory contains the Energi Governance smart contracts used by the Energi Core node as a key part of consensus. Most of these contracts were initially deployed as a part of the genesis block on the Energi Testnet and Energi Mainnet. They can be tested against a local ganache instance. All commands given here unless otherwise specified are meant to be run from the root of the repository, where [truffle-config.js](truffle-config.js) resides.

## Prerequisites
At a minimum to build and test these repositories you need `nodejs-v12`, `truffle`, and `ganache-cli`. Some dependencies here have shown some issues with later versions of nodejs. If you have a suitable version of nodejs, running `npm install` should get you to a workable state.

## Building
The contracts can be built using `truffle compile`. In order to build the golang bridge to the core node, `abigen` must be used. This can be invoked as part of the normal build process using `make all`.

## Testing
The quickest way to run all the tests are to run `make test-sol`. If you need some more control over the testing process, they can be run manually with `truffle` and `ganache-cli`. The tests for these contracts require some specific accounts set up on `ganache-cli` so run it as follows:
```
ganache-cli \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef,10000000000000000000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdff,200000000000000000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcfff,100000000000000000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abffff,11000000000000000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890afffff,1000000000000000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890ffffff,1000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef123456789fffffff,0' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef12345678ffffffff,100000000000000000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef1234567fffffffff,100000000000000000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef123456ffffffffff,100000000000000000000000' \
      --account='0x1234567890abcdef1234567890abcdef1234567890abcdef12345fffffffffff,10000000000000000000000000' \
      -l 10000000
```

Once `ganache-cli` is running, you can simply run `truffle test` in order to run all the tests, or to run a specific test you should run `truffle test energi/contracts/test/MyTestFile.spec.js`.

## Deploying
There are two types of deployment for these contracts. First is genesis block deployment, and second is deployment with truffle. The difference is whether contracts were deployed with the network as a part of the genesis block, or deployed after for the purpose of a governance upgrade or additional governance contract.

### Genesis Block Deployment
Genesis block deployment was used for the initial version of Energi Governance contracts that were deployed with the Energi Mainnet and Energi Testnet. The Simnet and Devnet still use genesis block deployment so local development networks can have the latest version of governance contracts. Generally, contracts deployed as a part of the genesis block will have a special address that is easily recognizable. These special addresses are defined in [energi/params/accounts.go](energi/params/accounts.go). You can see how the genesis block deployment works by looking at [core/genesis.go](core/genesis.go).

### Deployment with Truffle
For the purposes of running tests, or deploying upgrades to live networks, these should be deployed with truffle. In order to deploy contracts you will need to set up an account which has funds to pay for gas for the deployment cost. The best way to do this is to use a mnemonic phrase kept in the environment variable `TRUFFLE_MNEMONIC`. Be careful not to set this environment variable system wide, but only in a local shell as needed. [direnv](https://direnv.net/) is a recommended tool to manage environment variables conveniently. If you are ready to deploy you can easily migrate with `truffle migrate` which will deploy to the local `ganache-cli` instance by default. To deploy to testnet or mainnet simply use `truffle migrate --network testnet` or `truffle migrate --network mainnet`. Truffle migration files, which are kept at [energi/contracts/migrations/](energi/contracts/migrations/) begin with a number. You can do a single deployment using `truffle migrate -f N --to N` where `N` is the prefix of the migration you wish to run.
