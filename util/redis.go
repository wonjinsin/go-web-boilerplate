package util

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// RedisConnect ...
func RedisConnect(host string, port int, zlog *Logger) (redisDB *redis.Client, err error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	zlog.Infow("InitRedis", "addr", addr)
	redisDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
	})
	if _, err := redisDB.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return redisDB, nil
}
