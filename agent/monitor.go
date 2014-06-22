package agent

import (
    "sync"
    "strings"
)

const (
    OP_AVG = "avg"
    OP_SUM = "sum"
    OP_MAX = "max"
    OP_MIN = "min"
)

type MetricData struct {
    Namespace string
    Metric string
    Unit string
    Value float64
    Op string
}

type Samples struct {
    sync.Mutex
    metrics map[string]*MetricData
}

func (h *MetricData)Key() string {
    return h.Namespace + ":" + h.Metric
}

func (h *MetricData)Update(point *MetricData) {
    switch h.Op {
    default:
        h.avg(point)
    case OP_AVG:
        h.avg(point)
    case OP_MAX:
        h.max(point)
    case OP_MIN:
        h.min(point)
    case OP_SUM:
        h.sum(point)
    }
}

func (h *MetricData)max(point *MetricData) {
    if point.Value > h.Value {
        h.Value = point.Value
    }
}

func (h *MetricData)min(point *MetricData) {
    if point.Value < h.Value {
        h.Value = point.Value
    }
}

func (h *MetricData)sum(point *MetricData) {
    h.Value += point.Value
}

func (h *MetricData)avg(point *MetricData) {
    h.Value += point.Value
    h.Value /= 2
}

func (h *Samples)addPoint(data *MetricData) {
    actualPoint := h.metrics[data.Key()]

    if (actualPoint == nil) {
        actualPoint = new(MetricData)

        actualPoint.Metric = data.Metric
        actualPoint.Namespace = data.Namespace

        actualPoint.Value = data.Value
        actualPoint.Unit = data.Unit

        actualPoint.Op = strings.ToLower(data.Op)

        h.metrics[data.Key()] = actualPoint
    } else {
        actualPoint.Update(data)
    }
}

func collectData(metricPipe chan *MetricData) *Samples {

	database := new(Samples)
	database.metrics = make(map[string]*MetricData)

    W.Add(1)
    go func() {
        for {
            data, ok := <-metricPipe
            if !ok {
                L.Info("The metric data pipeline is closed!")
                break
            }

            database.Lock()
            database.addPoint(data)
            database.Unlock()
        }

        L.Info("I close the metric data collection daemon")
        W.Done()
    }()

	return database
}


