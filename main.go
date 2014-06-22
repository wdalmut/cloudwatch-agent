package main

import (
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
				conf.MergeWithFileAtPath(c.String("conf"))

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
