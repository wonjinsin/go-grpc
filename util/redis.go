package util

import (
	"context"
	"fmt"

	"phantom/config"

	"github.com/go-redis/redis/v8"
)

// RedisConnect ...
func RedisConnect(phantom *config.ViperConfig, zlog *Logger) (redisDB *redis.Client, err error) {
	host := fmt.Sprintf("%s:%d", phantom.GetString("redis.host"), phantom.GetInt("redis.port"))
	zlog.Infow("InitRedis", "redis_host", host)
	redisDB = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
	})
	if _, err := redisDB.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return redisDB, nil
}
