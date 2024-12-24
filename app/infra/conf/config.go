package conf

import (
	"os"
	"path"
	"time"

	"github.com/lupguo/go-shim/shim"
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
	ViewPath       string              `json:"view_path" yaml:"view_path"`
	ViewParseFiles map[string][]string `json:"view_parse_files" yaml:"view_parse_files"`
}

// Config App配置
type Config struct {
	Root   string        `json:"root" yaml:"root"`
	Listen string        `json:"listen" yaml:"listen"`
	Log    *LogConfig    `json:"log" yaml:"log"`
	Assets *AssetConfig  `json:"assets" yaml:"assets"`
	Public string        `json:"public" yaml:"public"`
	Wisdom *WisdomConfig `json:"wisdom" yaml:"wisdom"`
}

// 系统默认配置
var cfg *Config

// ParseConfig 解析系统配置
func ParseConfig(filename string) (*Config, error) {
	// 解析config.yaml文件
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "os.ReadFile(%s) got err", filename)
	}
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.Wrapf(err, "yaml.Unmarshal got err")
	}

	// 基本检测
	if cfg.Log == nil || cfg.Assets == nil || cfg.Wisdom == nil {
		return nil, errors.Errorf("empty app config: %s", shim.ToJsonString(cfg, false))
	}

	return cfg, nil
}

// PublicPath 静态资源路径
func PublicPath() string {
	return path.Join(cfg.Root, cfg.Public)
}

// AppLogConfig 应用配置
func AppLogConfig() (*LogConfig, error) {
	if cfg.Log == nil {
		return nil, errors.New("empty log config")
	}
	return cfg.Log, nil
}

// GetWisdomFilePath 获取wisdom文件
func GetWisdomFilePath() string {
	return cfg.Wisdom.FilePath
}

// AssetViewPath  返回项目根目录 root_path/view_path/xx.tmpl
//
//	viewFilename示例：
//	- 支持表达式: main/*.tmpl
//	- 支持名称: index.tmpl、partial/error.tmpl
func AssetViewPath(view string) string {
	return path.Join(cfg.Root, cfg.Assets.ViewPath, view)
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
	return cfg.Assets.ViewParseFiles
}
