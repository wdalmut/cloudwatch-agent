package agent

import (
    "sync"
)

type MetricData struct {
    Namespace string
    Metric string
    Unit string
    Value float64
}

func (h *MetricData)Update(point *MetricData) {
    h.Value += point.Value
    h.Value /= 2
}

type Samples struct {
    sync.Mutex
    metrics map[string]*MetricData
}

var Database = new(Samples)

func init() {
    Database.metrics = make(map[string]*MetricData)
}

func CollectData(metricPipe chan *MetricData) {
    for {
        data, ok := <-metricPipe
        if !ok {
            L.Info("The metric data pipeline is closed!")
            break
        }

        key := data.Namespace + ":" + data.Metric

        Database.Lock()
        addPoint(key, data)
        Database.Unlock()
    }

    L.Info("I close the metric data collection daemon")
    W.Done()
}

func addPoint(key string, data *MetricData) {
    actualPoint := Database.metrics[key]
    if (actualPoint == nil) {
        actualPoint = new(MetricData)

        actualPoint.Metric = data.Metric
        actualPoint.Namespace = data.Namespace

        actualPoint.Value = data.Value
        actualPoint.Unit = data.Unit

        Database.metrics[key] = actualPoint
    } else {
        actualPoint.Update(data)
    }
}

