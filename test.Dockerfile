FROM golang:1.19.2

ENV GOCACHE=/tmp/

ENV BASE_PATH=/go/src/github.com/james-cathcart/golog


WORKDIR $BASE_PATH

COPY . .
