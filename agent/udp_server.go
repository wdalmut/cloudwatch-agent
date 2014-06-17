package agent

import (
    "net"
    "strings"
    "strconv"
    "encoding/json"
)

func StartUDPServer(address string, port int) chan *MetricData {
    addr, _ := net.ResolveUDPAddr("udp", strings.Join([]string{address, strconv.Itoa(port)}, ":"))
    sock, err := net.ListenUDP("udp", addr)

    if err != nil {
        panic("Unable to start local agent at address" + address + ":" + strconv.Itoa(port))
    }

    L.Info("Started UDP monitor at " + address + ":" + strconv.Itoa(port));

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
                L.Info("Shutdown command received! Close metric data channel immediately");
                close(metricDataChannel)
                break
            }

            L.Err("Unable to unpack monitor information");
        } else {
            metricDataChannel <- monitorData
        }
    }

    L.Info("I'm ready to close the UDP monitor socket")
    sock.Close()
    W.Done()
}
