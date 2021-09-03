#!/bin/sh

## ENV and VARs
REPO_BASE=kyledinh
BUILDER_IMG_NAME=$REPO_BASE/builder-protoserver-go
APP_IMG_NAME=$REPO_BASE/protoserver

## DOWNLOAD GO MODULES AND DEPENDENCIES 
unset GOPATH
cd ..
go mod vendor
cd kube

## BUILD APPLICATION 
# docker pull $BUILDER_IMG_NAME

echo; echo "Building application in $(pwd) "
docker run --rm -v $(pwd)/../:/opt $BUILDER_IMG_NAME:latest
CODE=$?
if [ $CODE != 0 ]; then
  echo; echo "Application binary build failed with code $CODE"
  exit $CODE
fi

echo; echo $(cat ../sem-version)-$(git show -s --pretty="%h") >| buildArtifacts/version
BUILD_VERSION=$(cat buildArtifacts/version)

## BUILD CONTAINER IMAGE
echo; echo "Building $APP_IMG_NAME container image"
docker build -t $APP_IMG_NAME -f dockerfile.protoserver .
CODE=$?
if [ $CODE != 0 ]; then
  echo; echo "Application container build failed with code $CODE"
  exit $CODE
fi

## TAGGING CONTAINER IMAGE
echo; echo "Tagging image with version $CONNECTOR_CHASE_IMG_NAME:$BUILD_VERSION"
docker tag $APP_IMG_NAME:latest $APP_IMG_NAME:$BUILD_VERSION

echo; echo "## LIST OF CONTAINER IMAGES:"
docker images | grep protoserver
