PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NameService \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=nsd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=nscli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

include Makefile.ledger
all: install

install: go.sum
		@echo "--> Installing nsd & nscli & ebrelayer"
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/nsd
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/nscli
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/ebrelayer

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify
		@go mod tidy

test:
	@go test -mod=readonly $(PACKAGES)
