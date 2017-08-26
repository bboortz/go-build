#!/bin/sh

source /build/lib.sh


if [ -z "${APP}" ]; then
    error "Variable APP is not set" 
fi
if [ -z "${ARCH}" ]; then
    error "Variable ARCH is not set" 
fi


echo "Install dependencies ..."
/build/get_dependencies.sh
echo "Format sources ..."
go fmt ./...
echo "Compile sources ..."
#go build -race -o $APP .
#go build -o $APP .
go build --ldflags "-linkmode external -extldflags -static" -o $APP .
