docker buildx create --use --name sparrowbuilder
docker buildx use sparrowbuilder

docker buildx build --platform linux/amd64,linux/arm64 --push -t poneding/sparrow -t registry.cn-hangzhou.aliyuncs.com/pding/sparrow .