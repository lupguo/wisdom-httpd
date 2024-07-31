package config

import (
	"os"
	"path"
	"time"

	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/go-shim/x/log"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// WisdomConfig wisdom.json文件配置
type WisdomConfig struct {
	FilePath        string        `json:"file_path" yaml:"file_path"`
	RefreshInterval time.Duration `json:"refresh_interval" yaml:"refresh_interval"`
}

// AssetConfig 视图路径配置
type AssetConfig struct {
	RootPath       string              `json:"root_path" yaml:"root_path"`
	DistPath       string              `json:"dist_path" yaml:"dist_path"`
	ViewPath       string              `json:"view_path" yaml:"view_path"`
	ViewParseFiles map[string][]string `json:"view_parse_files" yaml:"view_parse_files"`
}

// LogConfig 日志配置
type LogConfig struct {
	LogLevel      string `json:"log_level" yaml:"log_level"`
	LogFormat     string `json:"log_format" yaml:"log_format"`
	LogTimeFormat string `json:"log_time_format" yaml:"log_time_format"`
}

// AppConfig App配置
type AppConfig struct {
	Listen string        `json:"listen" yaml:"listen"`
	Log    *LogConfig    `json:"log" yaml:"log"`
	Assets *AssetConfig  `json:"assets" yaml:"assets"`
	Wisdom *WisdomConfig `json:"wisdom" yaml:"wisdom"`
}

// Config 应用配置
type Config struct {
	App *AppConfig `json:"app"`
}

// 系统默认配置
var appCfg *AppConfig

// ParseConfig 解析系统配置
func ParseConfig(filename string) (*AppConfig, error) {
	// 解析config.yaml文件
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "read filename fail: %v", filename)
	}
	var cfg *Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.Wrapf(err, "yaml unmarshal config fail")
	}

	// 设置系统默认值
	appCfg = cfg.App
	log.Infof("appConfig: %s", shim.ToJsonString(appCfg, true))

	// 基本检测
	if appCfg.Log == nil || appCfg.Assets == nil || appCfg.Wisdom == nil {
		return nil, errors.Errorf("ugly app config: %s", shim.ToJsonString(appCfg, false))
	}

	return cfg.App, nil
}

// AssetRoot 返回项目根目录
func AssetRoot() string {
	return appCfg.Assets.RootPath
}

// DistPath 静态资源路径
func DistPath() string {
	return path.Join(appCfg.Assets.RootPath, appCfg.Assets.DistPath)
}

// GetWisdomFilePath 获取wisdom文件
func GetWisdomFilePath() string {
	return appCfg.Wisdom.FilePath
}

// AssetViewPath  返回项目根目录 root_path/view_path/xx.tmpl
//
//	viewFilename示例：
//	- 支持表达式: main/*.tmpl
//	- 支持名称: index.tmpl、partial/error.tmpl
func AssetViewPath(view string) string {
	return path.Join(appCfg.Assets.RootPath, appCfg.Assets.ViewPath, view)
}

// GetViewPathList 获取视图文件的批path地址
func GetViewPathList(views ...string) []string {
	var paths []string
	for _, view := range views {
		paths = append(paths, AssetViewPath(view))
	}
	return paths
}

// GetViewParseFiles 视图解析的文件配置
func GetViewParseFiles() map[string][]string {
	return appCfg.Assets.ViewParseFiles
}
