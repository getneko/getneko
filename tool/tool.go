package tool

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

// 生成随机字符串
func GetRandString(length int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result = append(result, letters[num.Int64()])
	}

	return string(result), nil
}

// 计算sha256
func GetSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 判断三个切片是否有重复元素(gpt生成，无法确保可用性)
func HasDuplicate(slice1, slice2, slice3 []string) bool {
	// 创建一个map来存储切片中的元素
	elements := make(map[string]bool)

	// 检查第一个切片
	for i := 0; i < len(slice1); i++ {
		if elements[slice1[i]] && slice1[i] != "" {
			return true
		}
		elements[slice1[i]] = true
	}

	// 检查第二个切片
	for i := 0; i < len(slice2); i++ {
		if elements[slice2[i]] && slice2[i] != "" {
			return true
		}
		elements[slice2[i]] = true
	}

	// 检查第三个切片
	for i := 0; i < len(slice3); i++ {
		if elements[slice3[i]] && slice3[i] != "" {
			return true
		}
		elements[slice3[i]] = true
	}

	// 如果没有重复元素，返回false
	return false
}
