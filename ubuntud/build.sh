#!/bin/bash

docker buildx build . --push --platform linux/arm64,linux/amd64 \
    -t poneding/ubuntud:2204 \
    -t registry.cn-hangzhou.aliyuncs.com/pding/ubuntud:2204