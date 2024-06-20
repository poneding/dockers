#!/bin/bash

docker buildx use mybuilder || docker buildx create --use --name mybuilder
docker buildx build --platform linux/amd64,linux/arm64 --push -t poneding/infinity -t registry.cn-hangzhou.aliyuncs.com/pding/infinity .
# docker buildx rm mybuilder