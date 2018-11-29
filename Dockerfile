FROM golang:1.11-rc-alpine AS builder

RUN apk update && apk add --update build-base alpine-sdk musl-dev musl

WORKDIR /go/src/github.com/videocoin/hookd

ADD . ./

ENV GO111MODULE off

RUN make build-alpine

FROM alpine:latest AS release

COPY --from=builder /go/src/github.com/videocoin/hookd/bin/hookd ./

ENTRYPOINT ./hookd

