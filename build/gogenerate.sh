
export GO111MODULE=on
export GOFLAGS=

go get -u github.com/fjl/gencodec
go get -u golang.org/x/tools/cmd/stringer
go get -u github.com/go-bindata/go-bindata/...

export GO111MODULE=on
export GOFLAGS=

go generate github.com/energicryptocurrency/go-energi/core/types
go generate github.com/energicryptocurrency/go-energi/core/vm
go generate github.com/energicryptocurrency/go-energi/core
go generate github.com/energicryptocurrency/go-energi/eth/tracers/internal/tracers/
go generate github.com/energicryptocurrency/go-energi/eth/
go generate github.com/energicryptocurrency/go-energi/internal/jsre/deps/
go generate github.com/energicryptocurrency/go-energi/p2p/discv5
go generate github.com/energicryptocurrency/go-energi/signer/rules/deps
go generate github.com/energicryptocurrency/go-energi/whisper/whisperv6/

