FROM golang:1.15.5-alpine
MAINTAINER Todor Nikolov <nikolov.todor@gmail.com>

RUN apk add --no-cache bash

COPY . /go/src/go-game
WORKDIR /go/src/go-game/cmd/gogame

RUN ["go", "get"]
RUN ["go", "build", "-o", "bin/go-game"]

RUN ["chmod", "+x", "/go/src/go-game/cmd/gogame/scripts/run-app.sh"]