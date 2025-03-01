package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gsadmin/core/config"
	"strings"
	"time"
)

type redisClient struct {
	db *redis.Client
}

func newRedisClient() *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Instance().Redis.RedisAddr,
		Password: config.Instance().Redis.RedisPWD, // no password set
		DB:       config.Instance().Redis.RedisDB,  // use default DB
	})

	return &redisClient{db: rdb}
}

func (r *redisClient) Set(name string, key string, value string, ttl int) error {
	err := r.db.Set(context.Background(), name+":"+key, value, time.Second*time.Duration(ttl)).Err()
	return err
}

func (r *redisClient) Get(name string, key string) (string, error) {
	res, err := r.db.Get(context.Background(), name+":"+key).Result()
	return res, err
}

func (r *redisClient) Put(name string, key string, value string, ttl int) error {
	err := r.db.Set(context.Background(), name+":"+key, value, time.Second*time.Duration(ttl)).Err()
	return err
}

func (r *redisClient) Del(name string, key string) error {
	err := r.db.Del(context.Background(), name+":"+key).Err()
	return err
}

func (r *redisClient) GetKeys(name string) ([]string, error) {
	var keys []string
	res, err := r.db.Keys(context.Background(), name+":*").Result()
	if err != nil {
		return nil, err
	}
	for _, key := range res {
		tmp := strings.Split(key, ":")
		if tmp[0] == name {
			keys = append(keys, tmp[1])
		}
	}
	return keys, nil
}

func (r *redisClient) Flush(name string) error {
	keys, err := r.db.Keys(context.Background(), name+":*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		if err1 := r.db.Del(context.Background(), key).Err(); err1 != nil {
			return err1
		}
	}
	return nil
}
