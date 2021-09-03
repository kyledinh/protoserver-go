#!/bin/sh

REPO_BASE=kyledinh
IMG_NAME=$REPO_BASE/builder-protoserver-go

# Alpine linux 3.9 (use --no-cache once this is stable)
docker build -t $IMG_NAME -f dockerfile.builder-protoserver-go .
CODE=$?
if [ $CODE != 0 ]; then
  echo "Builder container build failed with code $CODE"; echo
  exit $CODE
fi

echo; echo "## LIST OF CONTAINER IMAGES:"
docker images | grep builder-protoserver-go