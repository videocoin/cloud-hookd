FROM golang:1.10 as builder

RUN apt-get update
RUN apt-get install -y ca-certificates

COPY . /go/src/github.com/videocoin/cloud-hookd

WORKDIR /go/src/github.com/videocoin/cloud-hookd

RUN make test
RUN make build-bin

FROM bitnami/minideb:jessie

RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=builder /go/src/github.com/videocoin/cloud-hookd/bin/hookd /hookd

CMD ["/hookd"]
