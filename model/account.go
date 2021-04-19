package model

// 用户结构体
type User struct {
	Name   string
	Avatar []byte // 图片等二进制文件一般存成字节数组
	Token  string
}

// 二维码信息
type QrCode struct {
	Name   string
	Data   []byte
	Scan   chan bool
	Auth   chan bool
	IsScan bool
	ScanBy string
	Token  string
}
