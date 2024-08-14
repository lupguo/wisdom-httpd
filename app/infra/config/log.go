package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
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
func (cfg *LogConfig) GetLogLevel() log.Lvl {
	var level log.Lvl
	strLvl := strings.ToLower(cfg.LogLevel)
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
func (cfg *LogConfig) GetLogFormat() string {
	return cfg.LogFormat + "\n"
}

// GetLogTimeFormat 日志时间格式
func (cfg *LogConfig) GetLogTimeFormat() string {
	return cfg.LogTimeFormat
}

// InitLogger 初始化日志输出
func (cfg *LogConfig) InitLogger() error {
	switch cfg.Type {
	case "console":
		// 输出到控制台
		log.SetOutput(os.Stdout)
	case "file":
		// 输出到文件
		file, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		log.SetOutput(file)
	default:
		return fmt.Errorf("unknown output type: %s", cfg.Type)
	}
	return nil
}
