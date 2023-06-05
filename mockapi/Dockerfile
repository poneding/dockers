FROM --platform=$TARGETPLATFORM golang:alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download 
ARG TARGETPLATFORM TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o mockapi main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/mysettings /app/mysettings
COPY --from=builder /app/mockapi /app/mockapi

CMD ["./mockapi"]
