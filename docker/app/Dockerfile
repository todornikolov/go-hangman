FROM golang:1.15.5-alpine
MAINTAINER Todor Nikolov <nikolov.todor@gmail.com>

RUN apk add --no-cache bash

COPY . /go/src/go-hangman
WORKDIR /go/src/go-hangman/cmd/hangman

RUN ["go", "get"]
RUN ["go", "build", "-o", "bin/hangman"]

RUN ["chmod", "+x", "/go/src/go-hangman/cmd/hangman/scripts/run-app.sh"]