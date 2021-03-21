package config

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05", //时间格式
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
	//log.SetReportCaller(true)
}
