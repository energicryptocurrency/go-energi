# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: energi android ios energi3-cross swarm evm all test clean
.PHONY: energi3-linux energi3-linux-386 energi3-linux-amd64 energi3-linux-mips64 energi3-linux-mips64le
.PHONY: energi3-linux-arm energi3-linux-arm-5 energi3-linux-arm-6 energi3-linux-arm-7 energi3-linux-arm64
.PHONY: energi3-darwin energi3-darwin-386 energi3-darwin-amd64
.PHONY: energi3-windows energi3-windows-386 energi3-windows-amd64
.PHONY: prebuild

include energi/contracts/Makefile.include

GOBIN = $(shell pwd)/build/bin
GO ?= $(goVer)
GO ?= latest

prebuild:

energi:
	build/env.sh go run build/ci.go install ./cmd/energi
	@echo "Done building."
	@echo "Run \"$(GOBIN)/energi\" to launch energi."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all: prebuild
	build/env.sh go run build/ci.go install

package: all
	build/env.sh go run build/ci.go archive
	# TODO: architecture-specific packaging

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/energi3.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/EnergiCore.framework\" to use the library."

check: lint test

test: all test-go test-sol

test-go:
	git submodule update --init --recursive
	build/env.sh go run build/ci.go test

test-sol: lint-sol-tests lint-sol test-sol-contracts

lint: lint-go lint-sol-tests lint-sol

lint-go:
	build/env.sh go run build/ci.go lint

lint-sol-tests:
	yarn run eslint energi/contracts/

lint-sol:
	yarn run solium -d energi/contracts/

clean:
	./build/clean_go_build_cache.sh
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

clean-vcs:
	git clean -fdx . || true

update-license:
	go run ./build/update-license.go

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

swarm-devtools:
	env GOBIN= go install ./cmd/swarm/mimegen

# Cross Compilation Targets (xgo)

energi3-cross: energi3-linux energi3-darwin energi3-windows energi3-android energi3-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/energi3-*

energi3-linux: energi3-linux-386 energi3-linux-amd64 energi3-linux-arm energi3-linux-mips64 energi3-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-*

energi3-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/energi
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep 386

energi3-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/energi
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep amd64

energi3-linux-arm: energi3-linux-arm-5 energi3-linux-arm-6 energi3-linux-arm-7 energi3-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep arm

energi3-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/energi
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep arm-5

energi3-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/energi
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep arm-6

energi3-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/energi
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep arm-7

energi3-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/energi
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep arm64

energi3-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/energi
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep mips

energi3-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/energi
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep mipsle

energi3-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/energi
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep mips64

energi3-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/energi
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/energi3-linux-* | grep mips64le

energi3-darwin: energi3-darwin-386 energi3-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/energi3-darwin-*

energi3-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/energi
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-darwin-* | grep 386

energi3-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/energi
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-darwin-* | grep amd64

energi3-windows: energi3-windows-386 energi3-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/energi3-windows-*

energi3-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/energi
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-windows-* | grep 386

energi3-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/energi
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/energi3-windows-* | grep amd64
