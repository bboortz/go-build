#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

GOBUILD=$( which go-build 2> /dev/null )
if [ -z "$GOBUILD" ]; then
	echo "please install github.com/bboortz/go-build first"
	exit 1
fi


${GOBUILD} build application
${GOBUILD} test
go install
