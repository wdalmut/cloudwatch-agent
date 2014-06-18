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

### Complete data packet structure

The data packet could include other features, like: units, etc...

```
{
    "namespace": ...
    "metric": ...
    "unit": ...
    "value": ...
}
```

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
echo "deb-src http://dl.bintray.com/wdalmut/deb /" | sudo tee -a /etc/apt/sources.list.d/wdalmut.list
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

