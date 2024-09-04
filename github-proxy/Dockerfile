FROM alpine
ARG TARGETARCH

WORKDIR /app
COPY ./bin/${TARGETARCH}/github-proxy /app/github-proxy

CMD ["./github-proxy"]
