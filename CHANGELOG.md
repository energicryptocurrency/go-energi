# Core Node Changelog

## 3.1.0 (2021-08-12)

### Added
- masternode stats now available through RPC services
- hardfork registry
- checkpoint registry
- preimage hash correction upon detecting damage


### Changed
- improved logs
- disabled light client
- silenced zero fee tx errors
- disabled usb flag
- preimage insertion check
- stake cooldown reduced from 60 to 30 minutes

### Fixed
- data race concerning secKeyBuf buffer
- difficulty adjustment (based on PID controller)
