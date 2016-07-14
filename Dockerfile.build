FROM golang:1.6.2-alpine

RUN apk add -U make gcc git

ENV APP doenter
ENV REPO fntlnz/$(APP)

ADD . /go/src/github.com/fntlnz/doenter

WORKDIR /go/src/github.com/fntlnz/doenter
