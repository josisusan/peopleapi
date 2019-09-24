FROM golang:1.13-alpine

ENV WORKSPACE /workspace

COPY ./ $WORKSPACE

WORKDIR $WORKSPACE

RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    go install github.com/graniticio/granitic-yaml/v2/cmd/grnc-yaml-bind

RUN go mod download

CMD grnc-yaml-bind && go build && ./peopleapi
