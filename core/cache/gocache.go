package cache

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

type goCacheClient struct {
	db *cache.Cache
}

func newGoCacheClient() *goCacheClient {
	gdb := cache.New(1*time.Minute, 3*time.Minute)
	return &goCacheClient{db: gdb}
}

func (g *goCacheClient) Set(name string, key string, value string, ttl int) error {
	g.db.Set(name+":"+key, value, time.Second*time.Duration(ttl))
	return nil
}

func (g *goCacheClient) Get(name string, key string) (string, error) {
	res, found := g.db.Get(name + ":" + key)
	if found {
		return res.(string), nil
	}
	return "", errors.New("key not found")
}

func (g *goCacheClient) Put(name string, key string, value string, ttl int) error {
	if err := g.db.Replace(name+":"+key, value, time.Second*time.Duration(ttl)); err != nil {
		return err
	}
	return nil
}

func (g *goCacheClient) Del(name string, key string) error {
	g.db.Delete(name + ":" + key)
	return nil
}

func (g *goCacheClient) GetKeys(name string) ([]string, error) {
	var keys []string
	for key := range g.db.Items() {
		tmp := strings.Split(key, ":")
		if tmp[0] == name {
			keys = append(keys, tmp[1])
		}
	}
	return keys, nil
}

func (g *goCacheClient) Flush(name string) error {
	for key := range g.db.Items() {
		if strings.Split(key, ":")[0] == name {
			g.db.Delete(key)
		}
	}
	return nil
}
