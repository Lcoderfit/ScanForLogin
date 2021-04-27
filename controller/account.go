package controller

import (
	"ScanForLogin/config"
	"ScanForLogin/constant"
	"ScanForLogin/model"
	"ScanForLogin/utils"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/shunde/avatar-go/avatar"
	"github.com/shunde/rsc/qr"
	"html/template"
	"image/png"
	"net/http"
	"strconv"
)

func init() {
	users := []*model.User{
		{Name: "张三"},
		{Name: "李四"},
		{Name: "王二"},
	}

	for i := 0; i < len(users); i++ {
		buf := bytes.NewBuffer([]byte{})
		// 将字符串转换成图片类型
		m := avatar.NewAvatar(users[i].Name)
		// 将头像图片编码成字节流存入buf
		err := png.Encode(buf, m)
		if err != nil {
			utils.Logger.Error("头像编码失败")
		}
		users[i].Avatar = buf.Bytes()
		// field必须是字符串，value是一个接口类型
		_, err = model.RedisClient.HSet("user", strconv.Itoa(i), users[i]).Result()
		if err != nil {
			utils.Logger.Errorf("user cache设置失败, %s, %v\n", i, users[i])
		}
	}
}

// Index 首页
func Index(c *gin.Context) {
	//c.HTML(http.StatusOK, "index.html", nil)
	c.Redirect(http.StatusFound, "/pc-login")
}

// PcLogin pc端登录
func PcLogin(c *gin.Context) {
	var uid string
	//for {
	//	// 生成32位uuid
	//	uid = strings.Replace(fmt.Sprint(uuid.NewV4()), "-", "", -1)
	//	if ok, _ := model.RedisClient.HExists("user", uid).Result(); !ok {
	//		break
	//	}
	//}
	// 设置cookie
	c.SetCookie("uid", uid, 3600, "/", "localhost", false, true)
	url, err := utils.UrlJoin("http://localhost"+config.ServerCfg.HttpPort, "/qr-code", uid)
	if err != nil {
		utils.Logger.Error("url拼接失败，error：", err)
		fail(c, constant.UrlJoinError)
		return
	}
	c.Redirect(http.StatusFound, url)
}

// QrCode 获取二维码
func QrCode(c *gin.Context) {
	uid := c.Param("uid")
	if uc, err := c.Cookie("uid"); err != nil || uc == "" {
		utils.Logger.Error("uid在cookie中不存在")
		c.Redirect(http.StatusFound, "/pc-login")
	}
	if c.Request.Method == "GET" {
		url, err := utils.UrlJoin("http://localhost"+config.ServerCfg.HttpPort, "/cellphone", uid)
		if err != nil {
			utils.Logger.Error("url拼接失败")
			fail(c, constant.UrlJoinError)
			return
		}
		code, err := qr.Encode(url, qr.H)
		if err != nil {
			utils.Logger.Error("url编码失败")
			fail(c, constant.QrCodeEncodeError)
			return
		}
		t, err := template.ParseFiles("../static/base.html", "../static/pc.html")
		if err != nil {
			utils.Logger.Error("模板解析失败")
			fail(c, constant.TemplateParseError)
			return
		}
		err = t.ExecuteTemplate(c.Writer, "content", code.PNG())
		if err != nil {
			utils.Logger.Error("模板执行失败")
			fail(c, constant.TemplateExecuteError)
			return
		}
	}
}

// ConfirmScanStatus 确认扫描状态
func ConfirmScanStatus(c *gin.Context) {
	uid := c.Param("uid")
	uc, err := c.Cookie("uid")
	if err != nil || uc != uid {
		utils.Logger.Error("uid在cookie中不存在")
		fail(c, constant.UidNotExistError)
		return
	}
	var content string
	if c.Request.Method == "GET" {
		
	} else if c.Request.Method == "POST" {
		content = "PC端登录成功"
	}
	t, err := template.ParseFiles("../static/base.html", "../static/cellphone.html")
	if err != nil {
		utils.Logger.Error("模板解析失败")
		fail(c, constant.TemplateParseError)
		return
	}
	err = t.ExecuteTemplate(c.Writer, "content", content)
	if err != nil {
		utils.Logger.Error("模板执行失败")
		fail(c, constant.TemplateExecuteError)
		return
	}
}

//
//// Login 登录
//// 会生成一个32位的uuid字符串，然后根据该uuid和base url生成二维码
//func Login(c *gin.Context) {
//	var uid string
//	// 如果uuid已存在于缓存中，则重新创建，直到创建一个唯一的uuid为止
//	for {
//		// 生成32位的uuid字符串
//		uid = strings.Replace(fmt.Sprint(uuid.NewV4()), "-", "", -1)
//		if ok, _ := model.RedisClient.HExists("user", uid).Result(); !ok {
//			break
//		}
//	}
//	// 将字符串进行编码
//	url, err := utils.UrlJoin("http://localhost:"+config.ServerCfg.HttpPort, "/scan-code")
//	if err != nil {
//		utils.Logger.Error("url拼接失败")
//		fail(c, constant.UrlJoinError)
//		return
//	}
//	code, err := qr.Encode(url, qr.H)
//	if err != nil {
//		utils.Logger.Error("二维码生成失败")
//		fail(c, constant.QrCodeEncodeError)
//		return
//	}
//	qrCode := model.QrCode{
//		Name: uid + ".png",
//		Data: code.PNG(),
//		//Scan: make(chan bool),
//	}
//	// 结构体无法直接存入redis，需先转为json格式
//	val, err := json.Marshal(&qrCode)
//	if err != nil {
//		utils.Logger.Error("二维码信息转换为json格式失败")
//		fail(c, constant.QrCodeConvertJsonError)
//		return
//	}
//
//	// 将二维码信息缓存5分钟
//	_, err = model.RedisClient.Set(uid, &val, 5*time.Minute).Result()
//	if err != nil {
//		utils.Logger.Error("二维码信息缓存失败")
//		fail(c, constant.QrCodeCacheError)
//		return
//	}
//	successWithData(c, gin.H{
//		"uid": uid,
//	})
//}
//
//// QrCode 获取二维码
//func QrCode(c *gin.Context) {
//	uid := c.Param("uid")
//	val, err := model.RedisClient.Get(uid).Result()
//	if err != nil {
//		utils.Logger.Error("uid不存在")
//		fail(c, constant.UidNotExistError)
//		return
//	}
//	successWithData(c, val)
//}
