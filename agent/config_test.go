package agent

import(
	"testing"
	"strings";
)

func TestConfAreNotOverdriven(t *testing.T) {
	conf := &AgentConf{Port:1234}
	conf.MergeWithFile(strings.NewReader("{}"));

	if conf.Port != 1234 {
		t.Error("Conf are not merged")
	}
}

func TestConfAreOverdriven(t *testing.T) {
	conf := &AgentConf{Port:1234}
	conf.MergeWithFile(strings.NewReader("{\"port\": 1235}"))

	if conf.Port != 1235 {
		t.Error("Conf are not overwritten")
	}
}

func TestConfArePartiallyOverwritten(t *testing.T) {
	conf := &AgentConf{Port:1234}
	conf.MergeWithFile(strings.NewReader("{\"Key\": \"The key\"}"))

	if conf.Port != 1234 || conf.Key != "The key" {
		t.Error("Conf are not merged")
	}
}

