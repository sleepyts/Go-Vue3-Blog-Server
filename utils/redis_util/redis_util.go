package redis_util

import (
	"Gin-Learn/globalVar"
	"encoding/json"
	"time"
)

func SetObject(key string, value interface{}, expiration time.Duration) error {
	valueJson,err:=json.Marshal(value)
    if err!= nil {
        return err
    }

	// 存储到 Redis
	if err := globalVar.RedisDb.Set(key, valueJson, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func GetObject(key string) (interface{}, error) {
	// 从 Redis 中获取数据
	valueJson,err := globalVar.RedisDb.Get(key).Result()
	if err == nil {
		// 反序列化
		var value interface{}
		if err := json.Unmarshal([]byte(valueJson), &value) ; err != nil {
			return nil, err
		}
		return value, nil	
	}
	return nil, err
}