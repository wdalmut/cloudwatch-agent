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

## Install the agent


### RPM Packages

*Install RPM repo*

```
wget https://bintray.com/wdalmut/rpm/rpm -O bintray-wdalmut-rpm.repo
sudo mv bintray-wdalmut-rpm.repo /etc/yum.repos.d/
```

*Install the agent*

```
sudo yum install cloudwatch-agent.x86_64
```

*Configure your CloudWatch keys*

Edit file at: `/etc/default/cloudwatch-agent` and put your `KEY` and `SECRET`

*Run the service*

```
sudo service cw-agent start
```

You can check: `start`, `stop`, `restart`, `status`

### DEB Packages

*Install DEB repo*

```
echo "deb http://dl.bintray.com/wdalmut/deb /" | sudo tee /etc/apt/sources.list.d/wdalmut.list
echo "#deb-src http://dl.bintray.com/wdalmut/deb /" | sudo tee -a /etc/apt/sources.list.d/wdalmut.list
```

*Update your packages*

```
sudo apt-get update
```

*Install the agent*

```
sudo apt-get install cloudwatch-agent
```

*Configure your CloudWatch keys*

Edit file at: `/etc/default/cloudwatch-agent` and put your `KEY` and `SECRET`

*Run the service*

```
sudo service cw-agent start
```

You can check: `start`, `stop`, `restart`, `status`

### Local configuration

You can tweak your configuration updating the `/etc/cloudwatch-agent/cloudwatch-agent.conf` file

A configuration example is:

```json
{
    "region": "eu-west-1",
    "address": "127.0.0.1",
    "port": 1234,
    "loop": 60
}
```

You can also configure AWS credentials using: `key` and `secret` but: put
your credentials into `/etc/default/cloudwatch-agent` is actually the preferrered configuration file

