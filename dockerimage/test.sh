#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

error() {
    error "$1" >&2
    exit 1
}

export CGO_ENABLED=0 
PACKAGES=""
if [ -n "$IGNORE_PACKAGES" ]; then
	PACKAGES=$(go list ./... | grep -v $IGNORE_PACKAGES)
else
	PACKAGES=$(go list ./... )
fi
#echo "PACKAGES to test: $PACKAGES" >&2

glide --home ./.glide install > /tmp/glide.out 2>&1 || cat /tmp/glide.out >&2
go fmt ./... >&2 || error "go fmt failed."
go test -v ./... >&2 || error "go test failed."
for d in $PACKAGES; do
	echo -e "\nTEST DIR: /go/src/$d" >&2
	cd /go/src/$d
	gometalinter . >&2 || error "gometalinter failed for package: $d."
	cd -
done
