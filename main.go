package main

import (
	"fmt"
	"net/url"
	"path"
)

//func init() {
//	//client := redis.NewClient(&redis.Options{
//	//	Addr:     "47.101.48.37:6379",
//	//	Password: "124541", // no password set
//	//	DB:       0,        // use default DB
//	//})
//	//
//	//pong, err := client.Ping().Result()
//	//if err != nil {
//	//	fmt.Println("err: ", err)
//	//	return
//	//} else {
//	//	fmt.Println("pong: ", pong)
//	//}
//
//	user := "http://blessing.lcoderfit.com"
//	// 默认缓冲区大小为4096，当超出默认缓冲区大小会进行自动扩容
//	// 返回一个bytes.Buffer类型的结构体
//	buf := bytes.NewBuffer([]byte{})
//	m := avatar.NewAvatar(user)
//	png.Encode(buf, m)
//	f, _ := os.OpenFile("a.png", os.O_CREATE|os.O_RDWR, 0755)
//	f.Write(buf.Bytes())
//}

func main() {
	//fmt.Println("begin")
	//var u uuid.UUID
	//u = uuid.NewV4()
	//fmt.Println(len(u), u)
	//fmt.Println()

	//content := "http://blessing.lcoderfit.com"
	//c, _ := qr.Encode(content, qr.H)
	//f, _ := os.OpenFile("a.png", os.O_CREATE|os.O_RDWR, 0755)
	//f.Write(c.PNG())

	s, err := UrlJoin("http://localhost:" + "8060", "/b", "/c/", "/d")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

// UrlJoin 对URL进行拼接
func UrlJoin(paths ...string) (string, error) {
	baseUrl, err := url.Parse(paths[0])
	if err != nil {
		return "", err
	}
	baseUrl.Path = path.Join(paths[1:]...)
	return baseUrl.String(), nil
}
