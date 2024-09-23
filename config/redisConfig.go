package config

import (
	"Go-Vue3-Blog-Server/globalVar"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", AppConfig.Redis.Host, AppConfig.Redis.Port),
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Fatalf("Redis connect error:%v", err)
	}
	globalVar.RedisDb = client
}
