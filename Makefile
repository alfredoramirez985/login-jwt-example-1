# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet

# Main package directory
PKG=./cmd/server

# Binary name
BINARY_NAME=main

# Port number for running the application
PORT=8080

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(PKG)

test:
	$(GOTEST) ./...

vet:
	$(GOVET) ./...

debug:
	$(GOBUILD) -o $(BINARY_NAME) -gcflags "all=-N -l" $(PKG)
	dlv --listen=:2345 --headless=true --api-version=2 exec ./$(BINARY_NAME) -- -port=$(PORT)

run: build
	./$(BINARY_NAME) -port=$(PORT)

.PHONY: all build test vet debug run
