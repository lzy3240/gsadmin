package queue

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gsadmin/core/config"
)

//  实现消息异步处理
//  第三方消息队列: Redie

type redisClient struct {
	rmap map[string]*rdqueue
	db   *redis.Client
}

type rdqueue struct {
	name string
	pb   *redis.PubSub
}

func newRedisClient() *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Instance().Redis.RedisAddr,
		Password: config.Instance().Redis.RedisPWD,
		DB:       config.Instance().Redis.RedisDB,
	})

	return &redisClient{
		rmap: make(map[string]*rdqueue),
		db:   rdb,
	}
}

func (r *redisClient) RegisterTopic(topic string) error {
	r.rmap[topic] = &rdqueue{
		name: topic,
		pb:   r.db.Subscribe(context.Background(), topic),
	}
	return nil
}

func (r *redisClient) Publish(topic string, message string) error {
	_, err := r.db.Publish(context.Background(), topic, message).Result()
	return err
}

func (r *redisClient) Subscribe(topic string, f func(param string)) {
	for value := range r.rmap[topic].pb.Channel() {
		f(value.Payload)
	}
}
