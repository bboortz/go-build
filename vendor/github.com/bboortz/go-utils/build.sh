#!/bin/bash

set -e
set -u

. ./build.settings

mkdir -p testdata
gofmt -w .
docker build -t ${program_name}-build .

# docker run -it -v $PWD/out:/out scm-build
