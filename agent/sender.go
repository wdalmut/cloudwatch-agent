package agent

import (
    "time"
    "fmt"
    "github.com/crowdmob/goamz/aws"
    "github.com/crowdmob/goamz/cloudwatch"
)

const (
    SCHEDULED_LOOP = 6
)

var cw *cloudwatch.CloudWatch

func init() {
    region := aws.Regions["eu-west-1"]
    auth, err := aws.EnvAuth()

    if err != nil {
        L.Err("Unable to send data to CloudWatch")

        panic("Unable to send data to CloudWatch")
    }

    cw,_ = cloudwatch.NewCloudWatch(auth, region.CloudWatchServicepoint)
}

func SendCollectedData() {
    doEvery(SCHEDULED_LOOP * time.Second, func(time time.Time) {
        Database.Lock()

        for key, point := range Database.metrics {
            metric := new(cloudwatch.MetricDatum)

            L.Info(fmt.Sprintf("letto %v", point))

            metric.MetricName = point.Metric
            metric.Timestamp = time
            metric.Unit = ""
            metric.Value = point.Value

            metrics := []cloudwatch.MetricDatum{*metric}

            L.Info(fmt.Sprintf("%v", metric))
            if _, err := cw.PutMetricDataNamespace(metrics, point.Namespace); err != nil {
                L.Err(fmt.Sprintf("%", err))
            } else {
                L.Info("Data sent to cloud")
            }


            delete(Database.metrics, key)
        }
        Database.Unlock()
    })
}

func doEvery(d time.Duration, f func(time.Time)) {
    for {
        L.Info("issue")
        time.Sleep(d)
        f(time.Now())
    }
}
