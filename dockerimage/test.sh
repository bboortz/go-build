#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

error() {
    error "$1" >&2
    exit 1
}


#export CGO_ENABLED=0 
export GOOS=linux
export GOARCH="${ARCH}"


if [ -z "${ARCH}" ]; then
    error "Variable ARCH is not set"
fi
if [ ! ${IGNORE_PACKAGES+x} ]; then
    error "Variable IGNORE_PACKAGES is not set"
fi


PACKAGES=""
if [ -n "$IGNORE_PACKAGES" ]; then
	PACKAGES=$(go list ./... | grep -v $IGNORE_PACKAGES)
else
	PACKAGES=$(go list ./... )
fi

go fmt ./... >&2 || error "go fmt failed."
#go test -v -race -coverprofile=cover.out ./... >&2 || error "go test failed."
go test -v ./... >&2 || error "go test failed."
if [ -f cover.out ]; then
	go tool cover -func=cover.out
fi
for d in $PACKAGES; do
	echo -e "\nTEST DIR: /go/src/$d" >&2
	cd /go/src/$d
	go vet . >&2 || error "gometalinter failed for package: $d."
	gometalinter . >&2 || error "gometalinter failed for package: $d."
	cd -
done
