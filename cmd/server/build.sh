#!/bin/bash
APP_VERSION=$(head -1 ../../sem-version)
echo "APP_VERSION: $APP_VERSION"

# clean up
rm protoserverLinux
rm protoserverMac

# Linux
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -X 'main.Version=$APP_VERSION'" -o protoserverLinux

# MacOS
env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -X 'main.Version=$APP_VERSION'" -o protoserverMac
