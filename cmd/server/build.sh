#!/bin/bash
KUBE_DIR=$(pwd)/../../kube
APP_VERSION=$(cat $KUBE_DIR/sem-version)

rm protoserverMac

# Linux
# env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -X 'main.Version=$APP_VERSION'" -o protoserver

# MacOS
env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -X 'main.Version=$APP_VERSION'" -o protoserverMac
