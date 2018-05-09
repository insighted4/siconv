PROJ=siconv
ORG_PATH=github.com/insighted4
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)

VERSION ?= $(shell ./scripts/git-version)
ENVIRONMENT ?= development
export ENVIRONMENT

$( shell mkdir -p bin )
$( shell mkdir -p release/bin )
$( shell mkdir -p release/images )
$( shell mkdir -p results )

user=$(shell id -u -n)
group=$(shell id -g -n)

export GOBIN=$(PWD)/bin
# Prefer ./bin instead of system packages for things like protoc, where we want
# to use the version orchestra uses, not whatever a developer has installed.
export PATH=$(GOBIN):$(shell printenv PATH)
export GO15VENDOREXPERIMENT=1

export GIN_MODE=release

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

build: clean bin/database bin/server bin/updater

bin/database: check-go-version
	@echo "Building Database Tool"
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/server

bin/server: check-go-version
	@echo "Building Server"
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/server

bin/updater: check-go-version
	@echo "Building Updater"
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/updater

.PHONY: release-binary
release-binary:
	@echo "Releasing binary files"
	@go build -race -o release/bin/database -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/database
	@go build -race -o release/bin/server -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/server
	@go build -race -o release/bin/updater -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/updater

.PHONY: start
start: build
	@echo "Staring Server"
	@$(GOBIN)/server start

.PHONY: revendor
revendor:
	@dep ensure

test:
	@echo "Testing"
	@go test -v -i $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

testcoverage:
	@echo "Testing with coverage"
	@mkdir -p results
	@go test -v ./... | go2xunit -output results/tests.xml
	@gocov test ./... | gocov-xml > results/cobertura-coverage.xml

testrace:
	@echo "Testing with race detection"
	@go test -v -i --race $(shell go list ./... | grep -v '/vendor/')
	@go test -v --race $(shell go list ./... | grep -v '/vendor/')

vet:
	@echo "Running go tool vet on packages"
	@go vet $(shell go list ./... | grep -v '/vendor/')

fmt:
	@echo "Running gofmt on package sources"
	@go fmt $(shell go list ./... | grep -v '/vendor/')

lint:
	@echo "lint"
	@for package in $(shell go list ./... | grep -v '/vendor/' | grep -v '/api' | grep -v '/server/internal'); do \
      golint -set_exit_status $$package $$i || exit 1; \
	done

.PHONY: check-go-version
check-go-version:
	@echo "Checking Golang version"
	@./scripts/check-go-version

.PHONY: clean
clean:
	@echo "Cleaning Binary Folders"
	@rm -rf bin/*
	@rm -rf release/*
	@rm -rf results/*

testall: testcoverage testrace vet fmt # lint

FORCE:

.PHONY: fmt lint test testall testcoverage testrace vet