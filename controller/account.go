package controller

import (
	"ScanForLogin/model"
	"ScanForLogin/utils"
	"bytes"
	"github.com/shunde/avatar-go/avatar"
	"image/png"
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
		m := avatar.NewAvatar(users[i].Name)
		// 将头像字节流存入buf中
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
