#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/hello-app main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/arm64/hello-app main.go

docker buildx use gobuilder || docker buildx create --name gobuilder --use

# VERSION=v1
# docker buildx build --build-arg HELLO_APP_VERSION=$VERSION --push --platform=linux/amd64,linux/arm64 . \
#  -t poneding/hello-app:$VERSION \
#  -t registry.cn-hangzhou.aliyuncs.com/pding/hello-app:$VERSION \
#  -f Dockerfile

# VERSION=v2
# docker buildx build --build-arg HELLO_APP_VERSION=$VERSION --push --platform=linux/amd64,linux/arm64 . \
#  -t poneding/hello-app:$VERSION \
#  -t registry.cn-hangzhou.aliyuncs.com/pding/hello-app:$VERSION \
#  -f Dockerfile

VERSION=v3
docker buildx build --build-arg HELLO_APP_VERSION=$VERSION --push --platform=linux/amd64,linux/arm64 . \
 -t poneding/hello-app \
 -t poneding/hello-app:$VERSION \
 -t registry.cn-hangzhou.aliyuncs.com/pding/hello-app \
 -t registry.cn-hangzhou.aliyuncs.com/pding/hello-app:$VERSION \
 -f Dockerfile

 rm -rf bin/

 # docker buildx rm gobuilder