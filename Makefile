VERSION:=$(shell cat ./VERSION)

default: binary

prepare:
	go get github.com/wdalmut/cloudwatch-agent

all: prepare test
	go build -a -ldflags '-X github.com/wdalmut/cloudwatch-agent/agent.VERSION "$(VERSION)"'

test:
	go test -v ./...

binary: all

