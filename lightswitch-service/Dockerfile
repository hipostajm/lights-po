FROM golang:1.26.1
WORKDIR /app

COPY proto/ proto/
COPY lightswitch-service/go.mod lightswitch-service/go.sum ./lightswitch-service/

WORKDIR /app/lightswitch-service

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY lightswitch-service/ .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o lightswitch-service ./cmd

CMD ["./lightswitch-service"]
