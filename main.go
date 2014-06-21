package main

import (
    "os"
    "os/signal"
    "net"
    "strconv"
    "strings"
    "github.com/wdalmut/cloudwatch-agent/agent"
	"flag"
	"fmt"
)

const (
    endpoint = "127.0.0.1"
    port = 1234
)


func main() {
	var confPath string

	flag.StringVar(&confPath, "conf", "", "Local configuration path")
	flag.Parse()

	conf := agent.NewConf()
	mergeConfig(conf, confPath)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, os.Kill)
    agent.W.Add(1)
    go waitingForKillSignal(c)

    monitorChannel := agent.StartUDPServer(conf)
    agent.W.Add(1)
    go agent.CollectData(monitorChannel)
    go agent.SendCollectedData(conf)

    agent.W.Wait()
}

func mergeConfig(conf *agent.AgentConf, confPath string) {
	if confPath != "" {
		src, err := os.Open(confPath)
		if err == nil {
			agent.L.Info(fmt.Sprintf("Merge default configuration with file at path: %s", confPath))
			err = conf.MergeWithFile(src)
			if err != nil {
				agent.L.Warning(fmt.Sprintf("Unable to merge file %s, wrong JSON format probably", confPath))
			}
		} else {
			agent.L.Warning(fmt.Sprintf("Missing file at path %s", confPath))
		}
	}
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


