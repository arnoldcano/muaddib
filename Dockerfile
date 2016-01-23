FROM ubuntu:14.04

MAINTAINER Arnold Cano <arnoldcano@yahoo.com>

RUN apt-get update
RUN apt-get install -y golang

RUN useradd -m muaddib

ENV GOPATH /home/muaddib/go
WORKDIR /home/muaddib/go/src/github.com/arnoldcano/muaddib
COPY . /home/muaddib/go/src/github.com/arnoldcano/muaddib
RUN go build

CMD ["./muaddib"]

EXPOSE 8081
