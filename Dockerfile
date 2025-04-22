FROM golang:1.24-alpine AS builder

RUN apk --no-cache add make

WORKDIR /app
COPY . .

RUN make release

FROM alpine:3.21

RUN mkdir /etc/ezex-users
COPY --from=builder /app/build/ezex-users /usr/bin/ezex-users

EXPOSE 50051

ENTRYPOINT ["/usr/bin/ezex-users"] 