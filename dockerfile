FROM golang:latest

MAINTAINER zs "810909753@qq.com"

WORKDIR /app
ENV GOPROXY   https://gocenter.io

ADD . /app
RUN go build  main.go

ENTRYPOINT ["./main"]