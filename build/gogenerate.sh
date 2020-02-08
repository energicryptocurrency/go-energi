export GOPATH=$(pwd)/build/_workspace

go get -u github.com/fjl/gencodec
go get -u golang.org/x/tools/cmd/stringer
go get -u github.com/go-bindata/go-bindata/...

go generate github.com/ethereum/go-ethereum/core/types
go generate github.com/ethereum/go-ethereum/core/vm
go generate github.com/ethereum/go-ethereum/core
go generate github.com/ethereum/go-ethereum/eth/tracers/internal/tracers/
go generate github.com/ethereum/go-ethereum/eth/
go generate github.com/ethereum/go-ethereum/internal/jsre/deps/
go generate github.com/ethereum/go-ethereum/p2p/discv5
go generate github.com/ethereum/go-ethereum/signer/rules/deps
go generate github.com/ethereum/go-ethereum/whisper/whisperv6/

