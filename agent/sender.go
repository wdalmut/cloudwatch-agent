package agent

import (
    "time"
    "fmt"
    "github.com/crowdmob/goamz/aws"
    "github.com/crowdmob/goamz/cloudwatch"
)

var cw *cloudwatch.CloudWatch
var stopTicker chan struct{}

func initCloudWatchAgent(conf *AgentConf) {
    auth, err := aws.GetAuth(conf.Key, conf.Secret, "",  time.Time{})
    if err != nil {
        L.Err(fmt.Sprintf("Unable to create the auth structure for CloudWatch: %v", err))
        panic(fmt.Sprintf("Unable to create the auth structure for CloudWatch: %v", err))
    }

    region := aws.Regions[conf.Region]
    cw, err = cloudwatch.NewCloudWatch(auth, region.CloudWatchServicepoint)
    if err != nil {
        L.Err(fmt.Sprintf("Unable to login for send statistics: %v", err))
        panic(fmt.Sprintf("Unable to login for send statistics: %v", err))
    }
}

func sendCollectedData(conf *AgentConf) {
    initCloudWatchAgent(conf)

    stopTicker = make(chan struct{})
    doEvery(time.Duration(conf.Loop) * time.Second, func(time time.Time) {
        Database.Lock()
        for key, point := range Database.metrics {
            metric := new(cloudwatch.MetricDatum)

            metric.MetricName = point.Metric
            metric.Timestamp = time
            metric.Unit = point.Unit
            metric.Value = point.Value

            metrics := []cloudwatch.MetricDatum{*metric}

            if _, err := cw.PutMetricDataNamespace(metrics, point.Namespace); err != nil {
                L.Err(fmt.Sprintf("%v", err))
            } else {
                L.Info("Metric with key: \"" + key + "\" sent to cloud correcly")
            }


            delete(Database.metrics, key)
        }
        Database.Unlock()
    })
}

func doEvery(every time.Duration, f func(time.Time)) {
    ticker := time.NewTicker(every)

    W.Add(1)
    go func() {
        select {
        case <- ticker.C:
            f(time.Now())
        case <- stopTicker:
            ticker.Stop()

            L.Info("Received stop scheduler, force send operation to CloudWatch")
            f(time.Now())
            L.Info("All data sent to CloudWatch, closing data scheduler")

            W.Done()
            return
        }
    }()
}

