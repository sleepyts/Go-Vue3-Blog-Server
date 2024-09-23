package redis_util

import (
	"Go-Vue3-Blog-Server/globalVar"
	"encoding/json"
	"time"
)

func SetObject(key string, value interface{}, expiration time.Duration) error {
	valueJson, err := json.Marshal(value)
	if err != nil {
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
	valueJson, err := globalVar.RedisDb.Get(key).Result()
	if err == nil {
		// 反序列化
		var value interface{}
		if err := json.Unmarshal([]byte(valueJson), &value); err != nil {
			return nil, err
		}
		return value, nil
	}
	return nil, err
}

func DeleteKeysWithPrefix(prefix string) error {
	// 使用 SCAN 遍历所有键
	var cursor uint64
	var keys []string
	var err error
	for {
		// SCAN 命令查找带有特定前缀的键
		keys, cursor, err = globalVar.RedisDb.Scan(cursor, prefix+"*", 10).Result()
		if err != nil {
			println(err.Error())
			return err
		}

		if len(keys) > 0 {
			// 删除匹配的键
			err = globalVar.RedisDb.Del(keys...).Err()
			if err != nil {
				return err
			}
		}

		// 如果 cursor 为 0，说明遍历结束
		if cursor == 0 {
			break
		}
	}
	return nil
}
