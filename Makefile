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
.PHONY: build kube test 

analyze:
	@./scripts/lint.sh
	go vet -v cmd/...
	staticcheck github.com/kyledinh/protoserver-go/cmd/...

check:
	@./scripts/check.sh

db-init:
	@go run ./cmd/cli -migrate initialize

db-ping:
	@go run ./cmd/cli -migrate ping

db-seed:
	@go run ./cmd/cli -migrate seed 


kube:
	@./kube/make-builder-container.sh
	@./kube/make-protoserver-container.sh

docker-up:
# currently postgres db only for local development
	@docker-compose -f docker/local-compose.yaml up -d

dockder-down:
	@docker-compose -f docker/local-compose.yaml down -d

gen-models:
# requires https://github.com/kyledinh/btk-go
	@btk -gen=models -i ./spec/jwt-tokens.latest.yaml 

lint: 
	@./scripts/lint.sh

setup:
	@./scripts/setup.sh

test:
	go test ./...
