# Makefile for Go project

# Variables
GO_CMD=go
MOCKERY_CMD=mockery
PKG=./...
BUILD_PKG=./...
CMD_PKG=./cmd/todo
COVERAGE_FILE=coverage.out

# Targets
.PHONY: all build run test test-coverage mocks clean

all: build

build:
	$(GO_CMD) build $(BUILD_PKG)

run:
	$(GO_CMD) run $(CMD_PKG)/main.go

test:
	$(GO_CMD) test -v $(shell go list $(PKG) | grep -v /test/)

test-coverage:
	$(GO_CMD) test -v -coverprofile=$(COVERAGE_FILE)  $(shell go list $(PKG) | grep -v /test/)
	$(GO_CMD) tool cover -html=$(COVERAGE_FILE)

mocks:
	$(MOCKERY_CMD) --all --output ./mocks

clean:
	$(GO_CMD) clean
	rm -f $(COVERAGE_FILE)
