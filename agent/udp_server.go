package agent

import (
    "net"
    "strings"
    "strconv"
    "encoding/json"
)

func StartUDPServer(conf *AgentConf) chan *MetricData {
    addr, _ := net.ResolveUDPAddr("udp", strings.Join([]string{conf.Address, strconv.Itoa(conf.Port)}, ":"))
    sock, err := net.ListenUDP("udp", addr)

    if err != nil {
		panic("Unable to start local agent at address" + conf.Address + ":" + strconv.Itoa(conf.Port))
    }

    L.Info("Started UDP monitor at " + conf.Address + ":" + strconv.Itoa(conf.Port));

    metricDataChannel := make(chan *MetricData)

    W.Add(1)
    go listenForMetrics(sock, metricDataChannel)

    return metricDataChannel
}

func listenForMetrics(sock *net.UDPConn, metricDataChannel chan *MetricData) {
    for {
        bytes := make([]byte, 1024)
        dataLen,_ := sock.Read(bytes)

        monitorData := new(MetricData)
        err := json.Unmarshal(bytes[:dataLen], monitorData)

        if err != nil {
            closeProcedure := string(bytes[:dataLen])
            if strings.Trim(closeProcedure, "!\n") == "close" {
                L.Info("Shutdown command received! Close data channel communication pipeline immediately!");
                close(metricDataChannel)
                break
            }

            L.Err("Unable to unpack monitor information");
        } else {
            metricDataChannel <- monitorData
        }
    }

    L.Info("I close the UDP monitor daemon")
    sock.Close()
    W.Done()
}
