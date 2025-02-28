package queue

import (
	"gsadmin/core/config"
	"sync"
)

var (
	quClient QuClient
	once     sync.Once
)

type QuClient interface {
	Publish(topic string, message string) error
	Subscribe(topic string, f func(param string))
	RegisterTopic(topic string) error
}

func Instance() QuClient {
	if quClient == nil {
		once.Do(func() {
			switch config.Instance().App.QueueMode {
			case "redis":
				quClient = newRedisClient()
			case "mem":
				quClient = newMemClient()
			default:
				quClient = newMemClient()
			}
		})
	}
	return quClient
}
