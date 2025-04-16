FROM golang:1.23.5-alpine3.20 as builder

RUN apk add --no-cache add make

WORKDIR /app
COPY . .

RUN make release

FROM alpine:3.21

RUN mkdir /etc/ezex-users

COPY --from=builder /app/build/ezex-users /usr/bin/ezex-users

EXPOSE 50051

ENTRYPOINT ["/usr/bin/ezex-users"] 