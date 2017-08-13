#!/bin/sh

source /build/lib.sh


glide --home ./.glide install > /tmp/glide.out 2>&1 || cat /tmp/glide.out >&2
