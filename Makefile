ROOT := github.com/kyledinh/protoserver-go

export SHELL := /bin/bash

BUILD_DIR = ./build
OUTPUT_DIR = ./dist

# Current version of the project.
GITTAG ?= $(shell git describe --tags --always --dirty)
SEMVAR ?= $(shell head -n 1 semvar)

# Golang standard bin directory.
GOPATH ?= $(shell go env GOPATH)
GOROOT ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

# ACTIONS
.PHONY: build test

analyze:
	@./scripts/lint.sh
	go vet -v cmd/...
	staticcheck github.com/kyledinh/protoserver-go/cmd/...

check:
	@./scripts/check.sh

lint: 
	@./scripts/lint.sh

setup:
	@./scripts/setup.sh

test:
	go test ./...
