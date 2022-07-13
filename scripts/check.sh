#!/bin/bash

REPO="protoserver-go"
THISDIR=$(pwd)
REPO_DIR=${THISDIR/$REPO*/$REPO}
VERSION=$(head -n 1 $REPO_DIR/sem-version) 

echo "Checking development environment"
echo "REPO_DIR: $REPO_DIR"
echo "SEM_VERSION: $VERSION"
