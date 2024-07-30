package config

import (
	"testing"
)

func TestRootPath(t *testing.T) {
	appCfg = &Config{
		Path: &Path{
			RootPath: "/api",
			Assets: struct {
				AssetPath string `yaml:"asset_path"`
				ViewPath  string `yaml:"view_path"`
			}(struct {
				AssetPath string
				ViewPath  string
			}{}),
		},
	}

	tests := []struct {
		name   string
		paths  []any
		output string
	}{
		{"t1", []any{"users"}, "/api/users"},
		{"t2", []any{"users", "profile"}, "/api/users/profile"},
		{"t3", []any{"products", "123", "reviews"}, "/api/products/123/reviews"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RootPath(tt.paths...)

			if got != tt.output {
				t.Errorf("got %s, but want %s", got, tt.output)
			}
		})
	}
}
