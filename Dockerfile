FROM golang:1.13-alpine

ENV WORKSPACE /workspace

WORKDIR $WORKSPACE

COPY go.mod /workspace
COPY go.sum /workspace

RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    go mod download && \
    go install github.com/graniticio/granitic-yaml/v2/cmd/grnc-yaml-bind

COPY ./ $WORKSPACE

CMD grnc-yaml-bind && go build && ./peopleapi
