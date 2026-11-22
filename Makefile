.PHONY: proto build test deps
export GOPRIVATE=github.com/altercpre
export PATH:=$(CURDIR)/deps:$(PATH)
BUILD_GOOS:=$(shell go env GOOS)
BUILD_GOARCH:=$(shell go env GOARCH)

build:
	@$(eval FLAGS := $$(shell PATH=$(PATH) govvv -flags -pkg github.com/altercpre/aparok-sync/app))
	GOOS=$(BUILD_GOOS) GOARCH=$(BUILD_GOARCH) go build -v -o bin/any-sync-consensusnode -ldflags "$(FLAGS) -X github.com/altercpre/aparok-sync/app.AppName=any-sync-consensusnode" github.com/altercpre/aparok-sync-consensusnode/cmd

test:
	go test ./... --cover

deps:
	go mod download
	go build -o deps github.com/ahmetb/govvv
	go build -o deps go.uber.org/mock/mockgen

mocks:
	echo 'Generating mocks...'
	go generate ./...
