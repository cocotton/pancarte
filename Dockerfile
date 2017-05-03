FROM golang:1.7.5-alpine3.5

RUN apk update && apk add --no-cache git

COPY  . $GOPATH/src/github.com/cocotton/pancarte
RUN go get github.com/cocotton/pancarte
RUN go install github.com/cocotton/pancarte

ENV PANCARTE_JWT_SECRET someSecret

RUN $GOPATH/bin/pancarte