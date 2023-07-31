#!/bin/bash

docker buildx use pybuilder || docker buildx create --name pybuilder --use
docker buildx build --push --platform=linux/amd64,linux/arm64 . \
 -t poneding/counter \
 -t registry.cn-hangzhou.aliyuncs.com/pding/counter \
 -f Dockerfile
docker buildx rm pybuilder
