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
	database := new(Samples)
	database.metrics = make(map[string]*MetricData)

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1235
    data.Metric = "Test"
    data.Namespace = "Test"

    database.addPoint(data)

    if database.metrics["Test:Test"] == nil {
        t.Error("Data is not saved into database")
    }

    if database.metrics["Test:Test"].Unit != data.Unit ||
        database.metrics["Test:Test"].Value != data.Value ||
        database.metrics["Test:Test"].Namespace != data.Namespace ||
        database.metrics["Test:Test"].Metric != data.Metric {

        t.Error("You don't save the same metric structure")
    }

    database.metrics["Test:Test"] = nil
}

func TestComputePointAverageOnAnExistingPoint(t *testing.T) {
	database := new(Samples)
	database.metrics = make(map[string]*MetricData)

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1000
    data.Metric = "Test"
    data.Namespace = "Test"
    database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
    database.addPoint(data)

    if database.metrics["Test:Test"].Unit != data.Unit ||
        database.metrics["Test:Test"].Value != 600 ||
        database.metrics["Test:Test"].Namespace != data.Namespace ||
        database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the average -> %f should be 600", database.metrics["Test:Test"].Value))
    }

    database.metrics["Test:Test"] = nil
}

func TestComputePointMax(t *testing.T) {
	database := new(Samples)
	database.metrics = make(map[string]*MetricData)

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "max"
    database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "max"
    database.addPoint(data)

    if database.metrics["Test:Test"].Unit != data.Unit ||
        database.metrics["Test:Test"].Value != 200 ||
        database.metrics["Test:Test"].Namespace != data.Namespace ||
        database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the maximum -> %f should be 200", database.metrics["Test:Test"].Value))
    }

    database.metrics["Test:Test"] = nil
}

func TestComputePointMin(t *testing.T) {
	database := new(Samples)
	database.metrics = make(map[string]*MetricData)

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "min"
    database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "min"
    database.addPoint(data)

    if database.metrics["Test:Test"].Unit != data.Unit ||
        database.metrics["Test:Test"].Value != 1 ||
        database.metrics["Test:Test"].Namespace != data.Namespace ||
        database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the minimum -> %f should be 1", database.metrics["Test:Test"].Value))
    }

    database.metrics["Test:Test"] = nil
}


func TestComputePointSum(t *testing.T) {
	database := new(Samples)
	database.metrics = make(map[string]*MetricData)

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 200
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "sum"
    database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "sum"
    database.addPoint(data)

    if database.metrics["Test:Test"].Unit != data.Unit ||
        database.metrics["Test:Test"].Value != 201 ||
        database.metrics["Test:Test"].Namespace != data.Namespace ||
        database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the minimum -> %f should be 201", database.metrics["Test:Test"].Value))
    }

    database.metrics["Test:Test"] = nil
}

func TestComputePointCount(t *testing.T) {
	database := new(Samples)
	database.metrics = make(map[string]*MetricData)

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "sum"
    database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "sum"
    database.addPoint(data)

    if database.metrics["Test:Test"].Unit != data.Unit ||
        database.metrics["Test:Test"].Value != 2 ||
        database.metrics["Test:Test"].Namespace != data.Namespace ||
        database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not computing the minimum -> %f should be 2", database.metrics["Test:Test"].Value))
    }

    database.metrics["Test:Test"] = nil
}


func TestMixedOperationConservesTheFirstSelection(t *testing.T) {
	database := new(Samples)
	database.metrics = make(map[string]*MetricData)

    data := new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 1
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "sum"
    database.addPoint(data)

    data = new(MetricData)
    data.Unit = "Milliseconds"
    data.Value = 100
    data.Metric = "Test"
    data.Namespace = "Test"
    data.Op = "Max"
    database.addPoint(data)

    if database.metrics["Test:Test"].Unit != data.Unit ||
        database.metrics["Test:Test"].Value != 101 ||
        database.metrics["Test:Test"].Namespace != data.Namespace ||
        database.metrics["Test:Test"].Metric != data.Metric {

        t.Error(fmt.Sprintf("You are not conserving first choice operation (sum -> max) -> %f should be 101", database.metrics["Test:Test"].Value))
    }

    database.metrics["Test:Test"] = nil
}
