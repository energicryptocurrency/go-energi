# Core Node Changelog

# 3.1.3

### Changed
- Updated Golang version to 1.17

# 3.1.2

### Added
- Governance submodule
- Hardfork check function
- Simple PoS simulator

### Removed
- Swarm functionality
- Symbols from releases

### Fixed
- miner.setAutoCompounding()
- Repeated USB enumeration failures
- Builds on Mac

# 3.1.1

### Added
- Preimage exception package which fixes full sync
- Checkpoint for Asgard hard fork block
- Checkpoints for testnet
- Icon on Windows

### Changed
- Various code optimizations
- Treasury proposal fee structure: from 100 NRG to 300 NRG + 0.2% of the amount
- Renamed `autocollateralize` to `autocompounding`

### Fixed
- Unit tests
- Wrong port log message on testnet
- Balance issue in EVM affecting `msg.sender`
- Further enhancements to checkpoint registry

## 3.1.0 (Asgard Hard Fork)

### Added
- Masternode stats now available through RPC services
- Hardfork registry to enable scheduling and alerting users of upcoming hardforks
- New optimized version of checkpoint registry
- Preimage hash correction upon detecting damage

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
