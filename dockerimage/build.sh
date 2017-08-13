#!/bin/sh

source /build/lib.sh


if [ -z "${APP}" ]; then
    error "Variable APP is not set" 
fi
if [ -z "${ARCH}" ]; then
    error "Variable ARCH is not set" 
fi


/build/get_dependencies.sh
go fmt ./...
#go build -race -o $APP .
#go build -o $APP .
go build --ldflags "-linkmode external -extldflags -static" -o $APP .
