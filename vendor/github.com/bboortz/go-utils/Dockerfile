FROM golang:1.9-alpine
RUN apk update && \
	apk add \
		build-base \
		file \
		git && \
	rm -rf /var/cache/apk/*

RUN mkdir /app 
RUN mkdir /out
ADD . /app
WORKDIR /app 

ENV PROGRAM go-utils

RUN go get "github.com/davecgh/go-spew/spew"
RUN go get "github.com/stretchr/testify/assert"
RUN go get "github.com/bboortz/go-utils"

RUN go fmt ./...
RUN go test -v ./...

ENTRYPOINT [ "go" ]
CMD [ "test, "-v" ]
