package agent

import(
    "testing"
    "fmt"
)

func TestMetricDataGenerateAKey(t *testing.T) {
	data := new(MetricData)

	data.Namespace = "Test"
	data.Metric = "Test"

	if data.Key() != "Test:Test" {
		t.Error(fmt.Sprintf("Metric namespace should be \"Test:Test\" given %s", data.Key()))
	}
}

func TestPutNotExistingPoint(t *testing.T) {
	Database.metrics["Test:Test"] = nil

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1235
    data.Metric = "Test"
    data.Namespace = "Test"

    Database.addPoint(data)

	if Database.metrics["Test:Test"] == nil {
        t.Error("Data is not saved into Database")
    }

    if Database.metrics["Test:Test"].Unit != data.Unit ||
        Database.metrics["Test:Test"].Value != data.Value ||
        Database.metrics["Test:Test"].Namespace != data.Namespace ||
        Database.metrics["Test:Test"].Metric != data.Metric {

        t.Error("You don't save the same metric structure")
    }

	Database.metrics["Test:Test"] = nil
}

func TestComputePointAverageOnAnExistingPoint(t *testing.T) {
    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1000
    data.Metric = "Test"
    data.Namespace = "Test"
    Database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
    Database.addPoint(data)

    if Database.metrics["Test:Test"].Unit != data.Unit ||
        Database.metrics["Test:Test"].Value != 600 ||
        Database.metrics["Test:Test"].Namespace != data.Namespace ||
        Database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the average -> %f should be 600", Database.metrics["Test:Test"].Value))
    }

    Database.metrics["Test:Test"] = nil
}

func TestComputePointMax(t *testing.T) {
    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "max"
    Database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "max"
    Database.addPoint(data)

    if Database.metrics["Test:Test"].Unit != data.Unit ||
        Database.metrics["Test:Test"].Value != 200 ||
        Database.metrics["Test:Test"].Namespace != data.Namespace ||
        Database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the maximum -> %f should be 200", Database.metrics["Test:Test"].Value))
    }

    Database.metrics["Test:Test"] = nil
}

func TestComputePointMin(t *testing.T) {
    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "min"
    Database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "min"
    Database.addPoint(data)

    if Database.metrics["Test:Test"].Unit != data.Unit ||
        Database.metrics["Test:Test"].Value != 1 ||
        Database.metrics["Test:Test"].Namespace != data.Namespace ||
        Database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the minimum -> %f should be 1", Database.metrics["Test:Test"].Value))
    }

    Database.metrics["Test:Test"] = nil
}


func TestComputePointSum(t *testing.T) {
    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "sum"
    Database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "sum"
    Database.addPoint(data)

    if Database.metrics["Test:Test"].Unit != data.Unit ||
        Database.metrics["Test:Test"].Value != 201 ||
        Database.metrics["Test:Test"].Namespace != data.Namespace ||
        Database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the minimum -> %f should be 201", Database.metrics["Test:Test"].Value))
    }

    Database.metrics["Test:Test"] = nil
}

func TestComputePointCount(t *testing.T) {
    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "sum"
    Database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "sum"
    Database.addPoint(data)

    if Database.metrics["Test:Test"].Unit != data.Unit ||
        Database.metrics["Test:Test"].Value != 2 ||
        Database.metrics["Test:Test"].Namespace != data.Namespace ||
        Database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the minimum -> %f should be 2", Database.metrics["Test:Test"].Value))
    }

    Database.metrics["Test:Test"] = nil
}


func TestMixedOperationConservesTheFirstSelection(t *testing.T) {
    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "sum"
    Database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 100
    data.Metric = "Test"
    data.Namespace = "Test"
	data.Op = "Max"
    Database.addPoint(data)

    if Database.metrics["Test:Test"].Unit != data.Unit ||
        Database.metrics["Test:Test"].Value != 101 ||
        Database.metrics["Test:Test"].Namespace != data.Namespace ||
        Database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not conserving first choice operation (sum -> max) -> %f should be 101", Database.metrics["Test:Test"].Value))
    }

    Database.metrics["Test:Test"] = nil
}
