# FROM golang:alpine AS builder

# WORKDIR /app

# COPY . .
# RUN go mod download 
# ARG TARGETPLATFORM TARGETOS TARGETARCH
# RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o hello-app main.go

FROM alpine
ARG TARGETARCH HELLO_APP_VERSION

WORKDIR /app
COPY ./mysettings /app/mysettings
COPY ./bin/${TARGETARCH}/hello-app /app/hello-app

ENV HELLO_APP_VERSION=$HELLO_APP_VERSION

CMD ["./hello-app"]
