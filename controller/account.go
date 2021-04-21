package controller

import (
	"ScanForLogin/config"
	"ScanForLogin/constant"
	"ScanForLogin/model"
	"ScanForLogin/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/shunde/avatar-go/avatar"
	"github.com/shunde/rsc/qr"
	"image/png"
	"net/http"
	"strconv"
	"strings"
	"time"
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

// Login 登录
func Login(c *gin.Context) {
	var uid string
	// 如果uuid已存在于缓存中，则重新创建，直到创建一个唯一的uuid为止
	for {
		// 生成32位的uuid字符串
		uid = strings.Replace(fmt.Sprint(uuid.NewV4()), "-", "", -1)
		if ok, _ := model.RedisClient.HExists("user", uid).Result(); !ok {
			break
		}
	}
	// 将字符串进行编码
	url, err := utils.UrlJoin("http://localhost:"+config.ServerCfg.HttpPort, "/code")
	if err != nil {
		utils.Logger.Error("url拼接失败")
		fail(c, constant.UrlJoinError)
		return
	}
	code, err := qr.Encode(url, qr.H)
	if err != nil {
		utils.Logger.Error("二维码生成失败")
		fail(c, constant.QrCodeEncodeError)
		return
	}
	qrCode := model.QrCode{
		Name: uid + ".png",
		Data: code.PNG(),
		Scan: make(chan bool),
	}
	// 将二维码信息缓存5分钟
	_, err = model.RedisClient.Set(uid, &qrCode, 5*time.Minute).Result()
	if err != nil {
		utils.Logger.Error("二维码信息缓存失败")
		fail(c, constant.QrCodeCacheError)
		return
	}
	successWithData(c, gin.H{
		"uid": uid,
	})
}

// Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// QrCode 获取二维码
func QrCode(c *gin.Context) {
	uid := c.Param("uid")
	val, err := model.RedisClient.HGet("user", uid).Result()
	if err != nil {
		utils.Logger.Error("uid不存在")
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": val,
	})
}
