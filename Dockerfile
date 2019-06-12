FROM golang:latest AS builder

LABEL maintainer="Videocoin" description="nginx hooks"

RUN apt update && apt upgrade -y
RUN apt install ca-certificates -y

WORKDIR /go/src/github.com/videocoin/hookd

ADD ./ ./

RUN make build

FROM ubuntu:latest AS release


COPY --from=builder /go/src/github.com/videocoin/hookd/bin/hookd ./
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT [ "./hookd" ]
