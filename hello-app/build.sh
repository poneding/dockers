#!/bin/bash

VERSION=v1

docker buildx use gobuilder || docker buildx create --name gobuilder --use
docker buildx build --push --platform=linux/amd64,linux/arm64 . \
 -t poneding/hello-app \
 -t poneding/hello-app:$VERSION \
 -t registry.cn-hangzhou.aliyuncs.com/pding/hello-app \
 -t registry.cn-hangzhou.aliyuncs.com/pding/hello-app:$VERSION \
 -f Dockerfile
# docker buildx rm gobuilder
