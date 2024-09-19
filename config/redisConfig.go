package config

import (
	"Gin-Learn/globalVar"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",AppConfig.Redis.Host,AppConfig.Redis.Port),
		Password: AppConfig.Redis.Password,
	})

	if _,err:=client.Ping().Result(); err!=nil{
		log.Fatalf("Redis connect error:%v",err)
	}
	globalVar.RedisDb = client
}