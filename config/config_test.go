package config

import (
	"testing"
)

func TestRootPath(t *testing.T) {
	appConfig = &AppConfig{
		RootPath: "/api",
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
