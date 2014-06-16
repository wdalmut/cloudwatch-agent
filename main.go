package main

import (
    "sync"
    "github.com/wdalmut/cloudwatch-agent/agent"
)

var w sync.WaitGroup

func main() {
    w.Add(1)

    monitorChannel := agent.StartUDPServer("127.0.0.1", 1234)
    go agent.CollectData(monitorChannel)
    go agent.SendCollectedData()

    w.Wait()
}


