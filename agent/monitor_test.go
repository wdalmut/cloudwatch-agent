package agent

import(
    "testing"
    "fmt"
)

func TestPutNotExistingPoint(t *testing.T) {
    Database.metrics["key"] = nil

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1235
    data.Metric = "Test"
    data.Namespace = "Test"

    addPoint("key", data)

    if Database.metrics["key"] == nil {
        t.Error("Data is not saved into Database")
    }

    if Database.metrics["key"].Unit != data.Unit ||
        Database.metrics["key"].Value != data.Value ||
        Database.metrics["key"].Namespace != data.Namespace ||
        Database.metrics["key"].Metric != data.Metric {

        t.Error("You don't save the same metric structure")
    }

    Database.metrics["key"] = nil
}

func TestComputePointAverageOnAnExistingPoint(t *testing.T) {
    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1000
    data.Metric = "Test"
    data.Namespace = "Test"
    addPoint("key", data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
    addPoint("key", data)

    if Database.metrics["key"].Unit != data.Unit ||
        Database.metrics["key"].Value != 600 ||
        Database.metrics["key"].Namespace != data.Namespace ||
        Database.metrics["key"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the average -> %f should be 600", Database.metrics["key"].Value))
    }

    Database.metrics["key"] = nil
}
