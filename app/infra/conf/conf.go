package conf

import (
	"os"
	"path"

	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/go-shim/x/mysqlx"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// WisdomConfig wisdom.json文件配置
type WisdomConfig struct {
	FilePath string `json:"file_path" yaml:"file_path"`
	SKey     string `json:"skey" yaml:"skey"`
}

// AssetConfig 视图路径配置
type AssetConfig struct {
	ViewPath       string              `json:"view_path" yaml:"view_path"`
	ViewParseFiles map[string][]string `json:"view_parse_files" yaml:"view_parse_files"`
}

// Tool 工具配置
type Tool struct {
	RefreshToDB bool `json:"refresh_to_db" yaml:"refresh_to_db"`
}

// Config App配置
type Config struct {
	Root      string           `json:"root" yaml:"root"`
	Listen    string           `json:"listen" yaml:"listen"`
	LogConfig *log.Config      `json:"log" yaml:"log"`
	Assets    *AssetConfig     `json:"assets" yaml:"assets"`
	Public    string           `json:"public" yaml:"public"`
	Wisdom    *WisdomConfig    `json:"wisdom" yaml:"wisdom"`
	DBConfig  *mysqlx.DBConfig `yaml:"db"`
	Tool      *Tool            `json:"tool" yaml:"tool"`
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
	if cfg.LogConfig == nil || cfg.Assets == nil || cfg.Wisdom == nil {
		return nil, errors.Errorf("empty app config: %s", shim.ToJsonString(cfg, false))
	}

	return cfg, nil
}

// PublicPath 静态资源路径
func PublicPath() string {
	return path.Join(cfg.Root, cfg.Public)
}

// GetWisdomSentenceFilePath 获取wisdom文件
func GetWisdomSentenceFilePath() string {
	return cfg.Wisdom.FilePath
}

// GetDBConfig 从yaml获取DBConfig的配置信息
func GetDBConfig() (*mysqlx.DBConfig, error) {
	return cfg.DBConfig, nil
}

// GetRefreshToDBFlag true: 刷db， false: 不刷
func GetRefreshToDBFlag() bool {
	return cfg.Tool.RefreshToDB
}

// GetSKey 获取Yaml配置的密钥
func GetSKey() string {
	return cfg.Wisdom.SKey
}
