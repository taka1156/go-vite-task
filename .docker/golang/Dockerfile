FROM golang:1.18-alpine

RUN apk update && apk add git

RUN apk add alpine-sdk

ENV GO111MODULE=on

WORKDIR /go/src/app

RUN go get github.com/99designs/gqlgen
RUN go install github.com/rubenv/sql-migrate/...@latest
