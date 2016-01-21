FROM ubuntu:14.04

MAINTAINER Arnold Cano <arnoldcano@yahoo.com>

RUN apt-get update
RUN apt-get install -y golang

ENV GOPATH /go
WORKDIR /go/src/github.com/arnoldcano/muaddib
COPY . /go/src/github.com/arnoldcano/muaddib/
RUN go build

CMD ["./muaddib"]

EXPOSE 8081
