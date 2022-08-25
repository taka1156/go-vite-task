FROM golang:1.18-alpine

RUN apk update && apk add git

RUN apk add alpine-sdk

ENV GO111MODULE=on

RUN go install github.com/99designs/gqlgen@latest
RUN go install github.com/rubenv/sql-migrate/...@latest

WORKDIR /go/src/app
