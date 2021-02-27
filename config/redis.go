package config

import (
	"os"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func RedisInit() {
	host := os.Getenv("REDIS_HOST")
	pass := os.Getenv("REDIS_PASS")

	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisConnection() *redis.Client {
	return client
}

func SetCache(key string, value string, expirationTime time.Duration) error {
	setErr := GetRedisConnection().Set(key, value, expirationTime).Err()
	if setErr != nil {
		return setErr
	}

	return nil
}

func GetCache(key string) string {
	res := GetRedisConnection().Get(key)
	if res != nil {
		return res.Val()
	}
	return ""
}

func DeleteCache(key string) error {
	setErr := GetRedisConnection().Del(key).Err()
	if setErr != nil {
		return setErr
	}
	return nil
}
