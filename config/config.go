package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// ServerConfig wisdom-httpd服务配置
type ServerConfig struct {
	App    *AppConfig `yaml:"app"`
	Listen string     `yaml:"listen"` // 监听
}

// AppConfig 应用配置
type AppConfig struct {
	RootPath string `yaml:"root_path"` // 配置文件
	Log      struct {
		LogLevel      string `yaml:"log_level"`  // 日志等级(debug, info, error)
		LogFormat     string `yaml:"log_format"` // 日志格式
		LogTimeFormat string `yaml:"log_time_format"`
	} `yaml:"log"`
	Assets struct {
		AssetPath string `yaml:"asset_path"` // 静态资源path
		ViewPath  string `yaml:"view_path"`  // 视图资源path
	} `yaml:"assets"`
	Wisdom struct {
		FileType string `yaml:"file_type"` // 后期考虑DB、Http文件
		FileName string `yaml:"file_name"` // 文件名称（目前默认为本地文件)
	}
}

// 服务配置
var srvConfig *ServerConfig

// 应用配置
var appConfig *AppConfig

// ParseConfig 解析系统配置
func ParseConfig(filename string) (cfg *ServerConfig, err error) {
	// parse
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "read filename fail: %v", filename)
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "yaml unmarshal config fail")
	}

	// default setting
	if cfg.App == nil {
		return nil, errors.New("yaml app config error")
	}
	appConfig = cfg.App

	return cfg, nil
}

// GetRootPath 返回项目根目录
func GetRootPath() string {
	return appConfig.RootPath
}

// GetLogLevel 日志等级
func GetLogLevel() log.Lvl {
	var level log.Lvl
	strLvl := strings.ToLower(appConfig.Log.LogLevel)
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
	return appConfig.Log.LogFormat + "\n"
}

// GetLogTimeFormat 日志时间格式
func GetLogTimeFormat() string {
	return appConfig.Log.LogTimeFormat
}

// RootPath 返回path=routePath/subPath
func RootPath(paths ...any) string {
	format := strings.TrimRight(strings.Repeat("%s/", len(paths)), "/")
	subPath := fmt.Sprintf(format, paths...)
	return fmt.Sprintf("%s/%s", GetRootPath(), subPath)
}

// GetAssetPath 静态资源路径
func GetAssetPath() string {
	return RootPath(appConfig.Assets.AssetPath)
}

// GetSpecialViewPath 视图地址, assets/views/path
func GetSpecialViewPath(path string) string {
	return RootPath(appConfig.Assets.AssetPath, path)
}

// GetSpecialViewPathList 获取一批path地址
func GetSpecialViewPathList(paths ...string) []string {
	var ret []string
	for _, path := range paths {
		ret = append(ret, GetSpecialViewPath(path))
	}
	return ret
}

// GetWisdomFilename 获取wisdom文件
func GetWisdomFilename() string {
	return RootPath(appConfig.Wisdom.FileName)
}
