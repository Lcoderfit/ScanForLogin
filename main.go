package main

import (
	"ScanForLogin/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
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

//func index(w http.ResponseWriter, r *http.Request) {
//	// 定义模板
//	// 解析模板
//	// 父模板和子模板的顺序不能乱，父在前，子在后
//	t, err := template.ParseFiles("./static/base.html","./static/helllo.html")
//	if err != nil{
//		fmt.Printf("parse files failed, err : %v\n", err)
//		return
//	}
//	// 渲染模板
//	// 渲染模板时使用ExecuteTemplate函数，需要制定要被渲染的模板名称
//	err = t.ExecuteTemplate(w, "helllo.html","index")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}

func index(c *gin.Context) {
	// 定义模板
	// 解析模板
	// 父模板和子模板的顺序不能乱，父在前，子在后
	t, err := template.ParseFiles("./static/base.html", "./static/helllo.html")
	if err != nil {
		fmt.Printf("parse files failed, err : %v\n", err)
		return
	}
	// 渲染模板
	// 渲染模板时使用ExecuteTemplate函数，需要制定要被渲染的模板名称
	err = t.ExecuteTemplate(c.Writer, "helllo.html", "index")
	if err != nil {
		fmt.Println(err)
		return
	}
}

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

	//r := gin.New()
	//r.GET("/index", index)
	//r.Run(":8063")

	//content := "http://blessing.lcoderfit.com"
	//c, _ := qr.Encode(content, qr.H)
	//f, _ := os.OpenFile("a.png", os.O_CREATE|os.O_RDWR, 0755)
	//f.Write(c.PNG())

	//s, err := UrlJoin("http://localhost:" + "8060", "/b", "/c/", "/d")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(s)

	//client := redis.NewClient(&redis.Options{
	//	DB:       0,
	//	Password: "124541",
	//	Addr:     "47.101.48.37:6379",
	//})
	//
	//a := struct {
	//	C string `json:"c"`
	//	B int    `json:"b"`
	//}{
	//	C: "robert lu",
	//	B: 124,
	//}
	//aj, err := json.Marshal(&a)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//} else {
	//	fmt.Println("json marshal success")
	//}
	//_, err = client.Set("t", aj, 5*time.Minute).Result()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("set ok")
	//}
	//res, err := client.Get("gg").Result()
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//} else {
	//	fmt.Println("res: ", res)
	//}
	routes.InitRouter()
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
