#!/usr/bin/env bash

cd $(dirname $0)

docker run \
    -v $PWD/..:/go/src/github.com/BeardyC/draft \
    --workdir /go/src/github.com/BeardyC/draft \
    deis/go-dev:v0.22.0 \
    make $@
