#!/bin/bash

## SETUP ##############
APP_VERSION=$(git rev-parse HEAD)
ARTIFACTS_DIR=/opt/kube/buildArtifacts
CONTAINER_FILES_DIR=/opt/kube/containerFiles
DEBUG=true

# Create an empty buildArtifacts directory
mkdir -m 777 -p $ARTIFACTS_DIR
rm -rf  $ARTIFACTS_DIR/*

## FUNCTIONS ##########

# Debugger 
fn_debug() {
  if [ $DEBUG = true ]; then 
    echo "======== DEBUG =============="
    echo $(pwd)
    ls -la 
    echo "============================="
    echo
  fi
}

# Helper function ensures global $TEST_STATUS is 0. Exits if not.
fn_test_for_failure() {
  if [[ $TEST_STATUS -ne 0 ]]; then
     echo; echo "DETECTED A TEST FAILURE!! Exiting build..."
     exit $TEST_STATUS
  fi
}

## BUILD ##############

## Protoserver 
echo; echo "Compile and Build : protoserver"
cd cmd/server # Change directory to run build the protoserver binary in the server/ directory.
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -X 'main.Version=$APP_VERSION'" -o protoserver > $ARTIFACTS_DIR/compile-results.txt
TEST_STATUS=$?
cat $ARTIFACTS_DIR/compile-results.txt
fn_test_for_failure

fn_debug
mv protoserver $ARTIFACTS_DIR/.
cd ../.. # back to /opt

# For unit testing purposes
export PROTOSERVER_LOG_TYPE=stderr
export PROTOSERVER_LOG_LEVEL=debug

if [[ ! -f /etc/protoserver.json ]]; then
  echo '{"log": {"type": "stderr","level": "info"}}' > /etc/protoserver.json;
fi

echo; echo "Testing : protoserver"
go test -failfast -mod vendor protoserver-go... > $ARTIFACTS_DIR/test-results.txt
TEST_STATUS=$?
cat $ARTIFACTS_DIR/test-results.txt
fn_test_for_failure

echo; echo "Benchmarking : protoserver"
go test -mod vendor -bench=. -benchmem protoserver-go... > $ARTIFACTS_DIR/benchmarks-results.txt
TEST_STATUS=$?
cat $ARTIFACTS_DIR/benchmarks-results.txt
fn_test_for_failure

echo; echo "Done!"
echo "╔══ Should list the "protoserver" artifact built at $(date) ══╗"
ls -la $ARTIFACTS_DIR/protoserver
echo "╚══ Should list the "protoserver" artifact built at $(date) ══╝"