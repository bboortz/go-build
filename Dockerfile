FROM alpine:3.6
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
ADD go-build .

ENTRYPOINT [ "/build/go-build" ]
