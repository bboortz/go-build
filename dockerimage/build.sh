#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

error() {
    error "$1" >&2
    exit 1
}

if [ -z "${APP}" ]; then
    error "Variable APP is not set" 
fi
if [ -z "${ARCH}" ]; then
    error "Variable ARCH is not set" 
fi

export CGO_ENABLED=0 
export GOARCH="${ARCH}" 

glide --home ./.glide install > /tmp/glide.out 2>&1 || cat /tmp/glide.out >&2
go fmt ./...
#go build -race -o $APP .
go build -o $APP .
