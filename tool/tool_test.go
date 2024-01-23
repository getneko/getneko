package tool

import (
	"testing"
)

// 测试生成随机数
func TestGetRandString(t *testing.T) {
	t.Log(GetRandString(10))
}

// 测试计算sha256
func TestGetSha256(t *testing.T) {
	t.Log(GetSha256("getneko"))
}
