# Core Node Changelog

## 3.2.0 (Banana Hard Fork)

### Changed
- Removed testing Interval Generator
- Updated staking simulator according to Banana hardfork changes

### Fixed
- `energi.compensationProcess()` now prints an error on wrong input arguments
- `energi` directory contains linting fixes

### New
- Upon activating `Banana-txfee` hardfork, 10% of block transaction fees will be sent to coinbase as a part of a block reward
- Upon activating `Banana-blocktime` hardfork, energi network block generation time will be reduced from 60 seconds to 15 seconds
- Upon activating `Banana-pos` hardfork, new pos algorithm will make block reward distribution proportional to the user's active stake
- `make lint` now runs golang linting for energi subdirectory

## 3.1.3

### Changed
- Updated Golang version to 1.17

## 3.1.2

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

## 3.1.1

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

## 3.1.0 (Apple Hard Fork)

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
