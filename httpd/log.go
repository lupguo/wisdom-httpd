package httpd

import (
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/lupguo/wisdom-httpd/app/infra/config"
)

// GetLogLevel 日志等级
func GetLogLevel() log.Lvl {
	logCfg, _ := config.AppLogConfig()

	var level log.Lvl
	strLvl := strings.ToLower(logCfg.LogLevel)
	switch strLvl {
	case "debug":
		level = log.DEBUG
	case "info":
		level = log.INFO
	default:
		level = log.ERROR
	}
	return level
}

// GetLogFormat 获取应用的日志格式
func GetLogFormat() string {
	logCfg, _ := config.AppLogConfig()

	return logCfg.LogFormat + "\n"
}

// GetLogTimeFormat 日志时间格式
func GetLogTimeFormat() string {
	logCfg, _ := config.AppLogConfig()

	return logCfg.LogTimeFormat
}
