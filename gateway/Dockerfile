FROM golang:1.26.1
WORKDIR /app

COPY proto/ proto/
COPY gateway/go.mod gateway/go.sum ./gateway/

WORKDIR /app/gateway

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY gateway/ .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o gateway ./cmd

CMD ["./gateway"]
