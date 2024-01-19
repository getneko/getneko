package tool

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
)

// 生成随机字符串
func GetRandString(stringLen int) string {
	var stringArray = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	str := strings.Builder{}
	for i := 0; i < stringLen; i++ {
		str.WriteString(stringArray[rand.Intn(len(stringArray))])
	}
	return str.String()
}

// 计算sha256
func GetSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
