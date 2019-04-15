.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/unification-com/mainchain/cmd/swarm
	go build -o bin/tomo     github.com/unification-com/mainchain/cmd/tomo
	go build -o bin/bootnode github.com/unification-com/mainchain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/unification-com/mainchain/swarm/...
