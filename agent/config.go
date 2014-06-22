package agent

import(
    "encoding/json"
    "io/ioutil"
    "os"
    "io"
)

type AgentConf struct {
    Key string
    Secret string
    Region string
    Address string
    Port int
    Loop int
}

func NewConf() (*AgentConf){
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

func (d *AgentConf)MergeWithFile(src io.Reader) (error) {
    file, err := ioutil.ReadAll(src)
    if err != nil {
        return err
    }

    err = json.Unmarshal(file, d)

    return err
}

