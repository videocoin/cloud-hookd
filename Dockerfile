FROM ubuntu:latest AS release

LABEL maintainer="Videocoin" description="nginx hooks"

RUN apt update  && apt upgrade -y

WORKDIR /go/src/github.com/VideoCoin/hookd

ADD release/hookd-linux-amd64 ./

EXPOSE 50051 50052 50053 50054 50055

ENTRYPOINT [ "./hookd-linux-amd64" ]
