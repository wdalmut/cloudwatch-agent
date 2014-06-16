package main

import (
    "github.com/wdalmut/cloudwatch-agent/agent"
)

func main() {
    monitorChannel := agent.StartUDPServer("127.0.0.1", 1234)

    agent.W.Add(1)
    go agent.CollectData(monitorChannel)
    //Skip wait group
    go agent.SendCollectedData()

    agent.W.Wait()
}


