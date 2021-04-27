package model

// 用户结构体
type User struct {
	Name   string `json:"name"`
	Avatar []byte `json:"avatar"` // 图片等二进制文件一般存成字节数组
	Token  string `json:"token"`
}

// 二维码信息
type QrCode struct {
	Name   string    `json:"name"` // 二维码名称
	Data   []byte    `json:"data"` // 二维码转换而成的字节数组
	//Scan   chan bool `json:"scan"`
	//Auth   chan bool `json:"auth"`
	//IsScan bool      `json:"is_scan"`
	//ScanBy string    `json:"scan_by"`
	Token  string    `json:"token"`
}
