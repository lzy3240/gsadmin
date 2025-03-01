package cache

import (
	"github.com/nutsdb/nutsdb"
	"strings"
)

type nutsClient struct {
	db     *nutsdb.DB
	bucket string
}

func newNutsClient() *nutsClient {
	ndb, _ := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir("./runtime/cache"), // 数据库会自动创建
	)

	return &nutsClient{db: ndb, bucket: "default"}
}

func (n *nutsClient) Set(name string, key string, value string, ttl int) error {
	_ = n.isExistBucket(n.bucket)
	err := n.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(n.bucket, []byte(name+":"+key), []byte(value), uint32(ttl)); err != nil {
				return err
			}
			return nil
		})
	return err
}

func (n *nutsClient) Get(name string, key string) (string, error) {
	var v string
	err := n.db.View(
		func(tx *nutsdb.Tx) error {
			if value, err := tx.Get(n.bucket, []byte(name+":"+key)); err != nil {
				return err
			} else {
				v = string(value)
			}
			return nil
		})
	return v, err
}

func (n *nutsClient) Put(name string, key string, value string, ttl int) error {
	err := n.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(n.bucket, []byte(name+":"+key), []byte(value), uint32(ttl)); err != nil {
				return err
			}
			return nil
		})
	return err
}

func (n *nutsClient) Del(name string, key string) error {
	err := n.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Delete(n.bucket, []byte(name+":"+key)); err != nil {
				return err
			}
			return nil
		})
	return err
}

func (n *nutsClient) GetKeys(name string) ([]string, error) {
	var keys []string
	err := n.db.View(
		func(tx *nutsdb.Tx) error {
			entries, err := tx.GetKeys(n.bucket)
			if err != nil {
				return err
			}

			for _, entry := range entries {
				tmp := strings.Split(string(entry), ":")
				if tmp[0] == name {
					keys = append(keys, tmp[1])
				}
			}
			return nil
		})
	return keys, err
}

func (n *nutsClient) Flush(name string) error {
	keys, err := n.GetKeys(name)
	if err != nil {
		return err
	}
	for _, key := range keys {
		_ = n.Del(name, key)
	}
	return nil
}

func (n *nutsClient) isExistBucket(bucket string) bool {
	err := n.db.Update(func(tx *nutsdb.Tx) error {
		if !tx.ExistBucket(nutsdb.DataStructureBTree, bucket) {
			//fmt.Println("create bucket:" + bucket)
			return tx.NewBucket(nutsdb.DataStructureBTree, bucket)
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}
