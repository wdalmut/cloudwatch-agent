package main

import (
    "os"
    "os/signal"
    "net"
    "strconv"
    "strings"
    "github.com/wdalmut/cloudwatch-agent/agent"
)

const (
    endpoint = "127.0.0.1"
    port = 1234
)


func main() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, os.Kill)
    agent.W.Add(1)
    go waitingForKillSignal(c)

    monitorChannel := agent.StartUDPServer(endpoint, port)
    agent.W.Add(1)
    go agent.CollectData(monitorChannel)
    go agent.SendCollectedData()

    agent.W.Wait()
}

func waitingForKillSignal(c chan os.Signal) {
    _ = <-c

    agent.L.Info("Kill signal received!")

    addr, _ := net.ResolveUDPAddr("udp", strings.Join([]string{endpoint, strconv.Itoa(port)}, ":"))
    client,_ := net.DialUDP("udp", nil, addr)
    client.Write([]byte("close!\n"))
    close(c)

    agent.W.Done()
}


