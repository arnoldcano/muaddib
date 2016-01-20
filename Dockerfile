FROM ubuntu:14.04

MAINTAINER Arnold Cano <arnoldcano@yahoo.com>

RUN apt-get update
RUN apt-get install -y golang

ENV GOPATH /go
WORKDIR /go/src/github.com/arnoldcano/muaddib
ADD . /go/src/github.com/arnoldcano/muaddib/
RUN go build

ENTRYPOINT ["go", "run", "main.go"]

EXPOSE 8081
