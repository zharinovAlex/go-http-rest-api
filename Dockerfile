FROM golang:1.14-alpine3.11 AS build

LABEL maintainer="Alex Zharinov"

RUN apk add --update \
  git

RUN mkdir -p go/src/app

ADD . /go/src/app

WORKDIR /go/src/app/cmd/apiserver

RUN go get -v
