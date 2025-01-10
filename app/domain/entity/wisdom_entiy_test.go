package entity

import (
	"strings"
	"testing"
)

func TestGenerateRandomHex(t *testing.T) {
	hex := GenerateRandomHex("")
	t.Logf("Generated hex: %s", hex)

	// 检查返回的字符串是否以 "0x" 开头
	if !strings.HasPrefix(hex, "0x") {
		t.Errorf("Expected hex string to start with '0x', got %s", hex)
	}

	// 检查剩余部分是否为有效的16进制字符
	hexPart := hex[2:] // 去掉 "0x"
	for _, char := range hexPart {
		if !isHexChar(char) {
			t.Errorf("Invalid hex character found: %c in %s", char, hex)
		}
	}
}

// 辅助函数，检查字符是否为有效的16进制字符
func isHexChar(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'F')
}

func TestGenerateRandomHex1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{"LongLong"}},
		{"t2", args{"Long.Long"}},
		{"t3", args{"Hello"}},
		{"t4", args{"时间精力、能力、见识有限，你关注圈大小也是有限的，聚焦在你能掌控影响的事情上（当下不杂），扩大和提升自己的影响力！"}},
		{"t5", args{"时间精力、能力、见识有限，你关注圈大小也是有限的，聚焦在你能掌控影响的事情上（当下不杂），扩大和提升自己的影响力."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomHex(tt.args.s); len(got) == 0 {
				t.Errorf("GenerateRandomHex() = %v", got)
			}
		})
	}
}
