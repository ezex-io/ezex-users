FROM golang:1.23.5-alpine3.20 as builder
RUN apk add --no-cache git gmp-dev build-base g++ openssl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ./build/ezex-users ./cmd/server.go

FROM alpine:3.19
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/build/ezex-users /usr/bin/ezex-users

ENV WORKING_DIR="/app" \
    EZEX_USERS_HTTP_SERVER_ADDRESS=":8080" \
    EZEX_USERS_GRPC_SERVER_ADDRESS="0.0.0.0:50051"

WORKDIR $WORKING_DIR

EXPOSE 8080 50051

VOLUME $WORKING_DIR
ENTRYPOINT ["/usr/bin/ezex-users"] 