package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/wdalmut/cloudwatch-agent/agent"
)

const (
	VERSION = "0.0.9"
)

func main() {
	app := cli.NewApp()
	app.Name = "cloudwatch-agent"
	app.Usage = "Monitor your application via UDP/IP messages"
	app.Version = VERSION
	app.Commands = []cli.Command{
		{
			Name:  "capture",
			Usage: "Start capture messages on UDP/IP socket",
			Action: func(c *cli.Context) {
				conf := agent.NewConf()
				completeConfig(conf, c.String("conf"))

				agent.Capture(conf)
			},
			Flags: []cli.Flag{
				cli.StringFlag{"conf, c", "", "Your local configuration path"},
			},
		},
	}

	app.Run(os.Args)
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
