package utils

import "math/rand"

const (
	alphaNum string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	size     int    = 16
)

// 生成32位的UUID
func NewUuid() string {
	return randStringN(size)
}

// 生成N为的UUID
func NewUuIDN(n int) string {
	if n <= 0 {
		return nil
	}
	return randStringN(n)
}

// randStringN 生成长度为N的随机字符串
func randStringN(n int) string {
	cBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		// rand.Intn(n)会生成[0, n)范围内的数字
		cBytes[i] = alphaNum[rand.Intn(size)]
	}
	return string(cBytes)
}
