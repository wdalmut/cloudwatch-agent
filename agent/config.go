package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type AgentConf struct {
	Key     string
	Secret  string
	Region  string
	Address string
	Port    int
	Loop    int
}

var (
	VERSION string
)

// Prepare a new agent configuration
//
// The default parameters are
//  * REGION: eu-west-1
// 	* ADDRESS: 127.0.0.1
//  * PORT: 1234
//  * LOOP: 60
func NewConf() *AgentConf {
	daemonConf := new(AgentConf)
	daemonConf.Key = os.Getenv("AWS_ACCESS_KEY_ID")

	daemonConf.Secret = os.Getenv("AWS_SECRET_ACCESS_KEY")
	if daemonConf.Secret == "" {
		daemonConf.Secret = os.Getenv("AWS_SECRET_KEY")
	}

	daemonConf.Region = "eu-west-1"
	daemonConf.Address = "127.0.0.1"
	daemonConf.Port = 1234
	daemonConf.Loop = 60

	return daemonConf
}

// Merge default configuration with a JSON configuration
func (d *AgentConf) MergeWithFileAtPath(path string) error {
	src, err := os.Open(path)
	if err == nil {
		L.Info(fmt.Sprintf("Merge default configuration with file at path: %s", path))
		err = d.MergeWithReader(src)
		if err != nil {
			L.Warning(fmt.Sprintf("Unable to merge file %s, wrong JSON format probably", path))
		}
	} else {
		L.Warning(fmt.Sprintf("Missing file at path %s", path))
	}

	return err
}

// Merge default configuration with a JSON configuration
func (d *AgentConf) MergeWithReader(src io.Reader) error {
	file, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, d)

	return err
}
