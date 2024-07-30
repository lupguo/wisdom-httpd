package config

import (
	"strings"

	"github.com/labstack/gommon/log"
)

// GetLogLevel 日志等级
func GetLogLevel() log.Lvl {
	var level log.Lvl
	strLvl := strings.ToLower(appCfg.Log.LogLevel)
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
	return appCfg.Log.LogFormat + "\n"
}

// GetLogTimeFormat 日志时间格式
func GetLogTimeFormat() string {
	return appCfg.Log.LogTimeFormat
}
