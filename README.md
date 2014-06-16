# CloudWatch Agent

Just a simple agent that send your statistics to AWS CloudWatch using
an internal scheduler.

The monitor compute the average of all metrics and send them to
AWS CloudWatch every minute.

The agent listen for UDP packet on localhost on port 1234

## The data packet

The data packet is just a simple JSON message

```
{"namespace": "wdm", "metric": "indexer-timing", "value": 81.21}
```

## Close the agent

You can kill it or send via UDP the message `close!`

