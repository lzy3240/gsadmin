package cache

import (
	"gsadmin/core/config"
	"sync"
)

var (
	caClient CaClient
	once     sync.Once
)

type CaClient interface {
	Set(string, string, string, int) error
	Get(string, string) (string, error)
	Put(string, string, string, int) error
	Del(string, string) error
	GetKeys(string) ([]string, error)
	Flush(string) error
}

func Instance() CaClient {
	if caClient == nil {
		once.Do(func() {
			switch config.Instance().App.CacheMode {
			case "nuts":
				caClient = newNutsClient()
			case "redis":
				caClient = newRedisClient()
			case "mem":
				caClient = newGoCacheClient()
			default:
				caClient = newGoCacheClient()
			}
		})
	}
	return caClient
}
