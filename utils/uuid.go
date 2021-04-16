package utils

import "math/rand"

const (
	alphaNum string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	size     int    = 32
)

// 生成32位的UUID
func NewUuid() []byte {
	return randStringN(size)
}

// 生成N为的UUID
func NewUuIDN(n int) []byte {
	if n <= 0 {
		return nil
	}
	return randStringN(n)
}

// randStringN 生成长度为N的随机字符串
func randStringN(n int) []byte {
	uid := make([]byte, n)
	for i := 0; i < n; i++ {
		// rand.Intn(n)会生成[0, n)范围内的数字
		uid[i] = alphaNum[rand.Intn(size)]
	}
	return uid
}
