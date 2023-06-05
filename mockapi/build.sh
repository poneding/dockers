#!/bin/bash

export VERSION=v2

docker buildx build --push --platform=linux/amd64,linux/arm64 . \
 -t poneding/mockapi \
 -t poneding/mockapi:$VERSION \
 -t registry.cn-hangzhou.aliyuncs.com/pding/mockapi \
 -t registry.cn-hangzhou.aliyuncs.com/pding/mockapi:$VERSION \
 -f Dockerfile
