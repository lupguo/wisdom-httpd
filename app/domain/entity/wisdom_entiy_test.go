package entity

import (
	"strings"
	"testing"
)

func TestGenerateRandomHex(t *testing.T) {
	hex := GenerateRandomHex()
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
