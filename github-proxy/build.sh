#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/github-proxy main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/arm64/github-proxy main.go

docker buildx use gobuilder || docker buildx create --name gobuilder --use

docker buildx build --build-arg HELLO_APP_VERSION=$VERSION --push --platform=linux/amd64,linux/arm64 . \
 -t poneding/github-proxy \
 -t registry.cn-hangzhou.aliyuncs.com/pding/github-proxy \
 -f Dockerfile

 rm -rf bin/

 # docker buildx rm gobuilder