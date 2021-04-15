package main

import (
	"bytes"
	"fmt"
	"github.com/shunde/avatar-go/avatar"
	"image/png"
	"os"
)

func init() {
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "47.101.48.37:6379",
	//	Password: "124541", // no password set
	//	DB:       0,        // use default DB
	//})
	//
	//pong, err := client.Ping().Result()
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//} else {
	//	fmt.Println("pong: ", pong)
	//}

	user := "http://blessing.lcoderfit.com"
	// 默认缓冲区大小为4096，当超出默认缓冲区大小会进行自动扩容
	// 返回一个bytes.Buffer类型的结构体
	buf := bytes.NewBuffer([]byte{})
	m := avatar.NewAvatar(user)
	png.Encode(buf, m)
	f, _ := os.OpenFile("a.png", os.O_CREATE|os.O_RDWR, 0755)
	f.Write(buf.Bytes())
}

func main() {
	fmt.Println("begin")
}
