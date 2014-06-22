VERSION:=$(shell cat ./VERSION)

all:
	go build -a -ldflags '-X github.com/wdalmut/cloudwatch-agent/agent.VERSION "$(VERSION)"'

