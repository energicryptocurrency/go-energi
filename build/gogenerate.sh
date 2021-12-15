
export GO111MODULE=on
export GOFLAGS=

go get -u github.com/fjl/gencodec
go get -u golang.org/x/tools/cmd/stringer
go get -u github.com/go-bindata/go-bindata/...

export GO111MODULE=on
export GOFLAGS=

go generate github.com/energicryptocurrency/energi/core/types
go generate github.com/energicryptocurrency/energi/core/vm
go generate github.com/energicryptocurrency/energi/core
go generate github.com/energicryptocurrency/energi/eth/tracers/internal/tracers/
go generate github.com/energicryptocurrency/energi/eth/
go generate github.com/energicryptocurrency/energi/internal/jsre/deps/
go generate github.com/energicryptocurrency/energi/p2p/discv5
go generate github.com/energicryptocurrency/energi/signer/rules/deps
go generate github.com/energicryptocurrency/energi/whisper/whisperv6/

