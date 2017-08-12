FROM golang:1.9-alpine
RUN apk update && \
	apk add \
		build-base \
		file \
		git && \
	rm -rf /var/cache/apk/*

RUN mkdir /build
RUN mkdir /.glide
WORKDIR /build
ADD build.sh .
ADD test.sh .

RUN go get "github.com/stretchr/testify/assert"
RUN go get "github.com/Masterminds/glide"
RUN go get "github.com/alecthomas/gometalinter"
RUN gometalinter --install

ENTRYPOINT [ "/bin/sh" ]
