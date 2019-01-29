FROM ubuntu:latest AS release

WORKDIR /opt/

RUN apt update && apt upgrade -y

ADD bin/hookd ./

ENTRYPOINT ./hookd

