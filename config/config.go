package config

import (
	"ScanForLogin/utils"
	"gopkg.in/ini.v1"
	"os"
	"path"
)

var (
	ServerCfg = new(serverConfig)
	RedisCfg  = new(redisConfig)
)

// 服务器配置
type serverConfig struct {
	HttpPort string
	AppMode  string
}

// redis配置
type redisConfig struct {
	Addr     string
	DB       int
	Password string
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		utils.Logger.Panic("获取config路径失败")
	}
	cfgPath := path.Join(pwd, "config.ini")

	// 加载配置文件
	cfg, err := ini.Load(cfgPath)
	if err != nil {
		utils.Logger.Panic("读取config.ini失败")
		return
	}
	// 读取服务器配置
	err = cfg.Section("server").MapTo(ServerCfg)
	if err != nil {
		utils.Logger.Panic("读取服务器配置失败")
	}
	// 读取redis配置
	err = cfg.Section("redis").MapTo(RedisCfg)
	if err != nil {
		utils.Logger.Panic("读取redis配置失败")
	}

}
