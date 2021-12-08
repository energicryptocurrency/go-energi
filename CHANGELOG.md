# Core Node Changelog

# 3.1.1 (Asgard Hard Fork)

### Added
- Preimage exception package
- Checkpoint for Asgard hard fork block
- Checkpoints for testnet
- New chekpoint registry (allows sync from the genesis block)
- Icon on Windows

### Changed
- Various code optimizations
- Proposal fee structure: from 100 NRG to 300 NRG + 0.2% of the amount
- Renamed `autocollateralize` to `autocompounding`

### Fixed
- Errors on tests
- Wrong port on testnet
- Balance issue in EVM affecting `msg.sender`

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
