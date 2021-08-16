# Core Node Changelog

## 3.1.0 (Asgard Hard Fork)

### Added
- masternode stats now available through RPC services
- hardfork registry to enable scheduling and alerting users of upcoming hardforks
- new optimized version of checkpoint registry
- preimage hash correction upon detecting damage

### Changed
- improved log messages
- disabled light client
- disabled usb flag
- stake cooldown reduced from 60 to 30 minutes

### Fixed
- data race concerning secKeyBuf buffer
- difficulty adjustment (based on PID controller)
- issue syncing on mainnet
- issue where preimages may become corrupted
