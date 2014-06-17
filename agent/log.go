package agent

import (
    "log/syslog"
)

var L, _ = syslog.New(syslog.LOG_INFO, "SYCWAGENT")
