package conf

import (
	"fmt"
	"io"
	"os"
	"strings"

	elog "github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// LogConfig 日志配置
type LogConfig struct {
	Type          string `yaml:"type"`
	File          string `yaml:"file,omitempty"` // omitempty表示如果没有设置该字段则不输出
	LogLevel      string `json:"log_level" yaml:"log_level"`
	LogFormat     string `json:"log_format" yaml:"log_format"`
	LogTimeFormat string `json:"log_time_format" yaml:"log_time_format"`
}

// GetLogLevel 日志等级
func (cfg *LogConfig) GetLogLevel() elog.Lvl {
	var level elog.Lvl
	strLvl := strings.ToLower(cfg.LogLevel)
	switch strLvl {
	case "debug":
		level = elog.DEBUG
	case "info":
		level = elog.INFO
	default:
		level = elog.ERROR
	}
	return level
}

// GetLogFormat 获取应用的日志格式
func (cfg *LogConfig) GetLogFormat() string {
	return cfg.LogFormat + "\n"
}

// GetLogTimeFormat 日志时间格式
func (cfg *LogConfig) GetLogTimeFormat() string {
	return cfg.LogTimeFormat
}

// InitLogger 初始化日志输出
func InitLogger(cfg *LogConfig) (*logrus.Logger, error) {
	// 注入logrus
	stdLog := logrus.StandardLogger()

	// 设置log等级
	level, err := logrus.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		return nil, errors.Wrap(err, "parse logrus level")
	}
	stdLog.SetLevel(level)

	// 设置log输出位置
	var output io.Writer
	switch cfg.Type {
	case "console":
		// 输出到控制台
		output = os.Stdout
	case "file":
		// 输出到文件
		file, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		output = file
	default:
		return nil, fmt.Errorf("unknown output type: %s", cfg.Type)
	}
	stdLog.SetOutput(output)

	return stdLog, nil
}
