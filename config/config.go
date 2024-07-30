package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	Listen string `yaml:"listen"` // 监听
	Log    struct {
		LogLevel      string `yaml:"log_level"`  // 日志等级(debug, info, error)
		LogFormat     string `yaml:"log_format"` // 日志格式
		LogTimeFormat string `yaml:"log_time_format"`
	} `yaml:"log"`
	Path   *Path `yaml:"view_path"`
	Wisdom struct {
		FileType string `yaml:"file_type"` // 后期考虑DB、Http文件
		FileName string `yaml:"file_name"` // 文件名称（目前默认为本地文件)
	}
}

// Path 视频配置
type Path struct {
	RootPath string `yaml:"root_path"` // 根路径
	Assets   struct {
		AssetPath string `yaml:"asset_path"` // 静态资源path
		ViewPath  string `yaml:"view_path"`  // 视图资源path
	} `yaml:"assets"`
}

// 系统默认配置
var appCfg *Config

// ParseConfig 解析系统配置
func ParseConfig(filename string) (*Config, error) {
	// parse
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "read filename fail: %v", filename)
	}
	var cfg *Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "yaml unmarshal config fail")
	}

	// 设置系统默认值
	appCfg = cfg

	return cfg, nil
}

// GetRootPath 返回项目根目录
func GetRootPath() string {
	return appCfg.Path.RootPath
}

// RootPath 返回path=routePath/subPath
func RootPath(paths ...any) string {
	format := strings.TrimRight(strings.Repeat("%s/", len(paths)), "/")
	subPath := fmt.Sprintf(format, paths...)
	return fmt.Sprintf("%s/%s", GetRootPath(), subPath)
}

// GetAssetPath 静态资源路径
func GetAssetPath() string {
	return RootPath(appCfg.Path.Assets.AssetPath)
}

// GetViewRealPath 视图地址, assets/views/path
func GetViewRealPath(path string) string {
	return RootPath(appCfg.Path.Assets.ViewPath, path)
}

// GetViewRealPathList 获取视图文件的批path地址
func GetViewRealPathList(paths ...string) (fullPaths []string) {
	var ret []string
	for _, path := range paths {
		ret = append(ret, GetViewRealPath(path))
	}
	return ret
}

// GetWisdomFile 获取wisdom文件
func GetWisdomFile() string {
	return RootPath(appCfg.Wisdom.FileName)
}
