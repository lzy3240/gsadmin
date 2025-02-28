package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gsadmin/core/config"
	"gsadmin/core/log"
)

var redisCli *redis.Client

func Instance() *redis.Client {
	if redisCli == nil {
		InitRedis()
	}
	return redisCli
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Instance().Redis.RedisAddr,
		Password: config.Instance().Redis.RedisPWD, // no password set
		DB:       config.Instance().Redis.RedisDB,  // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Instance().Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		log.Instance().Info("redis connect ping response:", zap.String("pong", pong))
		redisCli = client
	}
}
