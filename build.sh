#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

GOBUILD=$( which go-build 2> /dev/null || echo -n )
if [ -z "$GOBUILD" ]; then
	if [ -f ./go-build ]; then
		echo
		echo "INFO: build with local application: ./go-build"
		GOBUILD=./go-build
	else
		echo
		echo "INFO: build build-container: go-build-base"
		cd dockerimage
		docker build -t go-build-base -f Dockerfile .
		cd ..
		echo
		echo "INFO: build application: go-build"
		docker run -it -u $(id -u):$(id -g) -v $(pwd):/go/src/github.com/bboortz/go-build -w /go/src/github.com/bboortz/go-build -e APP=go-build -e ARCH=amd64 go-build-base /build/build.sh
		GOBUILD=./go-build
		echo
		echo "INFO: build application again with go-build"
	fi
fi


${GOBUILD} build build-container
${GOBUILD} build application
GOBUILD=./go-build
${GOBUILD} test application
${GOBUILD} build container
#${GOBUILD} test container
#go install -race
go install
