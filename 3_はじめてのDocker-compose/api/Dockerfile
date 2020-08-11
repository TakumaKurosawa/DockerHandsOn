FROM golang:alpine3.11 as builder
ADD . /go/src/github.com/volare
WORKDIR /go/src/github.com/volare

ENV GO111MODULE=on

RUN apk update && \
    go build main.go

ENTRYPOINT ["./main"]