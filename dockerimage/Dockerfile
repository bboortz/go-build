FROM golang:1.9-alpine
RUN apk update && \
	apk add \
		build-base \
		file \
		git && \
	rm -rf /var/cache/apk/* && \
	addgroup -g 500 testgroup && \
	adduser -u 500 -G testgroup -h /home/testuser -s /bin/sh -D testuser testgroup 


RUN mkdir /build
RUN mkdir /.glide
WORKDIR /build
ADD lib.sh .
ADD get_dependencies.sh .
ADD build.sh .
ADD test.sh .


RUN go get "github.com/stretchr/testify/assert"
RUN go get "github.com/Masterminds/glide"
RUN go get "github.com/alecthomas/gometalinter"
RUN gometalinter --install


ENTRYPOINT [ "/bin/sh" ]
