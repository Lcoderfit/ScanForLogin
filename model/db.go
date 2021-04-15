package model

import (
	"ScanForLogin/config"
	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

func init() {
	// 连接redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisCfg.Addr,
		DB:       config.RedisCfg.DB,
		Password: config.RedisCfg.Password,
	})
}
