package agent

import (
	"encoding/json"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

func Capture(conf *AgentConf) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	waitingForKillSignal(c, conf)
	database := startUDPServer(conf)
	sendCollectedData(conf, database)
}

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

	return database
}

func listenForMetrics(sock *net.UDPConn, metricDataChannel chan *MetricData) {
	W.Add(1)
	go func() {
		for {
			bytes := make([]byte, 1024)
			dataLen, _ := sock.Read(bytes)

			monitorData := new(MetricData)
			err := json.Unmarshal(bytes[:dataLen], monitorData)

			if err != nil {
				closeProcedure := string(bytes[:dataLen])
				if strings.Trim(closeProcedure, "!\n") == "close" {
					L.Info("Shutdown command received! Close data channel communication pipeline immediately!")
					close(metricDataChannel)
					close(stopTicker)
					break
				}

				L.Err("Unable to unpack monitor information")
			} else {
				metricDataChannel <- monitorData
			}
		}

		L.Info("I close the UDP monitor daemon")
		sock.Close()
		W.Done()
	}()
}

func waitingForKillSignal(c chan os.Signal, conf *AgentConf) {
	W.Add(1)
	go func() {
		_ = <-c

		L.Info("Kill signal received!")

		addr, _ := net.ResolveUDPAddr("udp", strings.Join([]string{conf.Address, strconv.Itoa(conf.Port)}, ":"))
		client, _ := net.DialUDP("udp", nil, addr)
		client.Write([]byte("close!\n"))
		close(c)

		W.Done()
	}()
}
