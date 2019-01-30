FROM ubuntu:latest AS release

WORKDIR /opt/

RUN apt update && apt upgrade -y

ADD release/hookd-linux-amd64 ./

ENTRYPOINT ./hookd-linux-amd64

