package model

type User struct {
	Name   string
	Avatar []byte // 图片等二进制文件一般存成字节数组
	Token  string
}
