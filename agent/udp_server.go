package agent

import (
	"encoding/json"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

var closeAll chan struct{}

// Start the data capture
//
// The agent will listen for incoming messages on port
// 1234 (by default) via UDP/IP socket
func Capture(conf *AgentConf) {
	closeAll = make(chan struct{})

	database := startUDPServer(conf)
	sendCollectedData(conf, database)
	L.Info("CloudWatch agent started...")
}

// Start the UDP/IP server and prepare the empty databaes
//
// Just open a new UDP/IP socket server and return a new
// empty database (map key->value) in order to collect
// and compute statistics on incoming data
func startUDPServer(conf *AgentConf) *Samples {
	addr, _ := net.ResolveUDPAddr("udp", strings.Join([]string{conf.Address, strconv.Itoa(conf.Port)}, ":"))
	sock, err := net.ListenUDP("udp", addr)

	if err != nil {
		panic("Unable to start local agent at address" + conf.Address + ":" + strconv.Itoa(conf.Port))
	}

	L.Info("Started UDP monitor at " + conf.Address + ":" + strconv.Itoa(conf.Port))

	metricDataChannel := make(chan *MetricData)

	listenForMetrics(sock, metricDataChannel)
	database := collectData(metricDataChannel)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	waitingForKillSignal(c, sock)

	return database
}

// Stay up in order to collect incoming metrics
//
// Collect information and send them via metricDataChannel
// when a new valid packet is received via the opened socket
func listenForMetrics(sock *net.UDPConn, metricDataChannel chan *MetricData) {
	W.Add(1)
	go func() {
		for {
			bytes := make([]byte, 1024)
			if dataLen, err := sock.Read(bytes); err == nil {
				monitorData := new(MetricData)
				if err := json.Unmarshal(bytes[:dataLen], monitorData); err == nil {
					metricDataChannel <- monitorData
				} else {
					L.Err("Unable to unpack monitor information")
				}
			} else {
				break
			}
		}

		L.Info("I close the UDP monitor daemon")
		close(metricDataChannel)

		W.Done()
	}()
}

// Close the appllication on sigint or sigkill
//
// In case of interrupt or kill signal close the application
// closing all the opened sockets and channels
func waitingForKillSignal(c chan os.Signal, sock *net.UDPConn) {
	W.Add(1)
	go func() {
		_ = <-c
		L.Info("Kill signal received!")
		close(c)

		L.Info("Closing UDP/IP socket")
		sock.Close()

		L.Info("Closing sender data channel")
		close(closeAll)

		W.Done()
	}()
}
