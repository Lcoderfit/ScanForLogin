package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "47.101.48.37:6379",
		Password: "124541", // no password set
		DB:       0,        // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("err: ", err)
		return
	} else {
		fmt.Println("pong: ", pong)
	}
}

func main() {
	fmt.Println("begin")
}
