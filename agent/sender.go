package agent

import (
    "time"
    "fmt"
    "github.com/crowdmob/goamz/aws"
    "github.com/crowdmob/goamz/cloudwatch"
)

const (
    SCHEDULED_LOOP = 60
)

var cw *cloudwatch.CloudWatch

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

func SendCollectedData(conf *AgentConf) {
	initCloudWatchAgent(conf)

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

func doEvery(d time.Duration, f func(time.Time)) {
    for {
        time.Sleep(d)
        f(time.Now())
    }
}

