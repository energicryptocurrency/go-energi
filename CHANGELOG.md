# Core Node Changelog

## 1.2.0 (Banana Hard Fork)

### Changed
- Updated staking simulator according to Banana hardfork changes

### New
- Upon activating `Banana-txfee` hardfork, 10% of block transaction fees will be sent to coinbase as a part of a block reward
- Upon activating `Banana-blocktime` hardfork, energi network block generation time will be reduced from 60 seconds to 15 seconds
- Upon activating `Banana-pos` hardfork, new pos algorithm will make block reward distribution proportional to the user's active stake
- Upon activating `Banana-difficulty-adjustment` hardfork, new values will be used for difficulty adjustment algorithm

## 1.1.8

### Fixed
- Using top 100 oldest masternodes as bootnodes

## 1.1.7

### Changed
- bumped golang version to v1.21.4

### Fixed
- fixes debug_traceBlock methods that were missing governance transaction info

## 1.1.6

### Changed
- Removed armv5 support
- Removed mips support
- Supported Linux upgraded to Ubuntu 22.04
- Bootnodes settings
- Removed testing Interval Generator

### Fixed
- `energi` directory contains linting fixes
- unit test fixes
- build pipeline fixes

## 1.1.5

### Fixed
- Go module support
- Upgraded Golang version to 1.21.0

## 1.1.4

### Fixed
- preimages hotfix

## 1.1.3

### Changed
- Updated Golang version to 1.17

## 1.1.2

### Changed
- Removed swarm functionality
- Removed symbols from releases

### Fixed
- miner.setAutoCompounding()
- Repeated USB enumeration failures
- Builds on Mac

### New
- Governance submodule
- Hardfork check function
- Simple PoS simulator

## 1.1.1

### Changed
- Various code optimizations
- Treasury proposal fee structure: from 100 NRG to 300 NRG + 0.2% of the amount
- Renamed `autocollateralize` to `autocompounding`

### Fixed
- Unit tests
- Wrong port log message on testnet
- Balance issue in EVM affecting `msg.sender`
- Further enhancements to checkpoint registry

### New
- Preimage exception package which fixes full sync
- Checkpoint for Asgard hard fork block
- Checkpoints for testnet
- Icon on Windows

## 1.1.0 (Apple Hard Fork)

### Changed
- Improved log messages
- Disabled light client
- Disabled usb flag
- Stake cooldown reduced from 60 to 30 minutes

### Fixed
- Data race concerning secKeyBuf buffer
- Difficulty adjustment (based on PID controller)
- Issue syncing on mainnet
- Issue where preimages may become corrupted

### New
- Masternode stats now available through RPC services
- Hardfork registry to enable scheduling and alerting users of upcoming hardforks
- New optimized version of checkpoint registry
- Preimage hash correction upon detecting damage
