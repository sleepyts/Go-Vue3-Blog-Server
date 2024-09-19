package globalVar

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
	RedisDb *redis.Client
)