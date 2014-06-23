# CloudWatch Agent

Just a simple agent that send your statistics to AWS CloudWatch using
an internal scheduler.

The monitor compute the average of all metrics and send them to
AWS CloudWatch every minute.

The agent listen for UDP packet on localhost on port 1234

 * Master: [![Build Status](https://travis-ci.org/wdalmut/cloudwatch-agent.svg?branch=master)](https://travis-ci.org/wdalmut/cloudwatch-agent)
 * Develop: [![Build Status](https://travis-ci.org/wdalmut/cloudwatch-agent.svg?branch=develop)](https://travis-ci.org/wdalmut/cloudwatch-agent)

## The data packet

The data packet is just a simple JSON message

```
{"namespace": "wdm", "metric": "indexer-timing", "value": 81.21}
```

### Complete data packet structure

The data packet could include other features, like: units, etc...

```
{
    "namespace": ...
    "metric": ...
    "unit": ...
    "value": ...
    "op": ...
}
```

Valid operations:
  * `avg` - Compute the average (default)
  * `sum` - Sum all samples
  * `max` - Only the max sample
  * `min` - Only the min sample

Valid unit:
  * `Seconds`
  * `Microseconds`
  * `Milliseconds`
  * `Bytes`
  * `Kilobytes`
  * `Megabytes`
  * `Gigabytes`
  * `Terabytes`
  * `Bits`
  * `Kilobits`
  * `Megabits`
  * `Gigabits`
  * `Terabits`
  * `Percent`
  * `Count`
  * `Bytes/Second`
  * `Kilobytes/Second`
  * `Megabytes/Second`
  * `Gigabytes/Second`
  * `Terabytes/Second`
  * `Bits/Second`
  * `Kilobits/Second`
  * `Megabits/Second`
  * `Gigabits/Second`
  * `Terabits/Second`
  * `Count/Second`

## Close the agent

You can kill it or send via UDP the message `close!`

