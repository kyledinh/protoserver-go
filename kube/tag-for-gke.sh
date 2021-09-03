#!/bin/bash

## ENV and VARs
REPO_BASE=kyledinh
APP_IMG_NAME=$REPO_BASE/protoserver

GKE_IMG_NAME=gcr.io/protoserve/protoserver
GKE_VERSION=$(cat sem-version)-$(git show -s --pretty="%h")

docker tag $APP_IMG_NAME:latest $GKE_IMG_NAME:$GKE_VERSION
docker tag $APP_IMG_NAME:latest $GKE_IMG_NAME:latest

docker images | grep protoserver

echo; echo "Push to GCR (Google Container Registry)? y/n"
read CHOICE 

case "$CHOICE" in 
    y|Y|yes)
        docker push $GKE_IMG_NAME:$GKE_VERSION
        docker push $GKE_IMG_NAME:latest
        exit 0;;
    *)
        echo "Skipping push."
esac

