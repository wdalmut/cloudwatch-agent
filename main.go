package main

import (
    "os"
    "flag"
    "fmt"
    "github.com/wdalmut/cloudwatch-agent/agent"
)

func main() {
    var confPath string

    flag.StringVar(&confPath, "conf", "", "Local configuration path")
    flag.Parse()

    conf := agent.NewConf()
    completeConfig(conf, confPath)

    agent.Capture(conf)

    agent.L.Info("Waiting for the close signal")
    agent.W.Wait()
}

func completeConfig(conf *agent.AgentConf, confPath string) {
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



