VERSION:=$(shell cat ./VERSION)

default: binary

all: test
	go build -a -ldflags '-X github.com/wdalmut/cloudwatch-agent/agent.VERSION "$(VERSION)"'

test:
	go test -v ./...

binary: all

