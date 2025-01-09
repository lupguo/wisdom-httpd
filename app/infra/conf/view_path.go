package conf

import (
	"path"
)

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
