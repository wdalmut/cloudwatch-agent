package agent

import (
	"fmt"
	"time"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/cloudwatch"
)

// Init the CloudWatch client
//
// The configuration is used in order to prepare the client with
// key, secret and region.
func initCloudWatchAgent(conf *AgentConf) *cloudwatch.CloudWatch {
	auth, err := aws.GetAuth(conf.Key, conf.Secret, "", time.Time{})
	if err != nil {
		L.Err(fmt.Sprintf("Unable to create the auth structure for CloudWatch: %v", err))
		panic(fmt.Sprintf("Unable to create the auth structure for CloudWatch: %v", err))
	}

	region := aws.Regions[conf.Region]
	cw, err := cloudwatch.NewCloudWatch(auth, region.CloudWatchServicepoint)
	if err != nil {
		L.Err(fmt.Sprintf("Unable to login for send statistics: %v", err))
		panic(fmt.Sprintf("Unable to login for send statistics: %v", err))
	}

	return cw
}

// Send collected data to CloudWatch
//
// Uses doEvery method in order to send collected data to
// AWS Cloudwatch. Just one point is actually sent to CloudWatch
func sendCollectedData(conf *AgentConf, database *Samples) {
	cw := initCloudWatchAgent(conf)

	doEvery(time.Duration(conf.Loop)*time.Second, func(time time.Time) {
		database.Lock()
		for key, point := range database.metrics {
			metric := point.Datum()
			metric.Timestamp = time

			if _, err := cw.PutMetricDataNamespace([]cloudwatch.MetricDatum{metric}, point.Namespace); err != nil {
				L.Err(fmt.Sprintf("%v", err))
			} else {
				L.Info("Metric with key: \"" + key + "\" sent to cloud correcly")
			}

			delete(database.metrics, key)
		}
		database.Unlock()
	})
}

// Execute the anonymous function every given amount of time
//
// The anon func is executed tipically every 60 seconds and
// the timer is closed when a message appears on closeAll
// channel (just close it)
func doEvery(every time.Duration, f func(time.Time)) {
	ticker := time.NewTicker(every)

	W.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				f(time.Now())
			case <-closeAll:
				ticker.Stop()

				L.Info("Received stop scheduler, force send operation to CloudWatch")
				f(time.Now())
				L.Info("All data sent to CloudWatch, closing data scheduler")

				W.Done()
				return
			}
		}
	}()
}
