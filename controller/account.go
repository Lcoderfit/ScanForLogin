package controller

import (
	"ScanForLogin/model"
	"ScanForLogin/utils"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/shunde/avatar-go/avatar"
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
	c.HTML(http.StatusOK, "index.html", nil)
}

// QrCode 获取二维码
func QrCode(c *gin.Context) {

}
