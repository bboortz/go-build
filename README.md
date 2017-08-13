# go-build
go-build is tool to build golang application and libraries. With it you are able to simplify your build and test your application.
go-build is using docker to build and test your application. Therefore we are avoiding dependencies to your operating system.


## Goals

* standardizied and simplified build and test process
* reasonable build and test
* reasonable dependency management
* avoiding dependencies to your operating system


## Build

* ./build.sh


## Usage

* cd ~/go/src
* go-build create github.com/YOURNAME/YOURREPO
* cd github.com/YOURNAME/YOURREPO
* go-build build application
* go-build test
* go-build build container


## Features

* create a new application
* build golang application 
* build application container
* build build-container
* test golang application or library


## Build Dependencies

These tools must be installed first

* golang
* docker


## Runtime Dependencies

* github.com/urfave/cli
* github.com/BurntSushi/toml
* github.com/bboortz/go-utils


## Decisions

* My Development Environment: vim
* Dependency Management: with glide
* Repository Structure:
** dockerimage: contains some Dockerfiles
** testdata: test data
* Coding Style: 
** format: gofmt
** lint: gometalinter
* Configuration: toml files
* logging: own logging from github.com/bboortz/go-utils/logger
* testing: with go test and gometalinter

