# Core Node Changelog

## 3.2.0 (Banana Hard Fork)

### Fixed
- `energi.compensationProcess()` now prints an error on wrong input arguments
- `energi` directory contains linting fixes

### New
- `make lint` now runs golang linting for energi subdirectory
- Upon activating `Banana` hardfork, 10% of block fees will be sent to coinbase as a reward

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
