.PHONY: und und-cross evm all test clean
.PHONY: und-linux und-linux-386 und-linux-amd64 und-linux-mips64 und-linux-mips64le
.PHONY: und-darwin und-darwin-386 und-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= latest
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

und:
	build/env.sh go run build/ci.go install ./cmd/und
	@echo "Done building."
	@echo "Run \"$(GOBIN)/und\" to launch und."

bootnode:
	build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	build/env.sh go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

und-cross: und-linux und-darwin
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/und-*

und-linux: und-linux-386 und-linux-amd64 und-linux-mips64 und-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-*

und-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/und
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep 386

und-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/und
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep amd64

und-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/und
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep mips

und-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/und
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep mipsle

und-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/und
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep mips64

und-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/und
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep mips64le

und-darwin: und-darwin-386 und-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/und-darwin-*

und-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/und
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/und-darwin-* | grep 386

und-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/und
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/und-darwin-* | grep amd64

gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
