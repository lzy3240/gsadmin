package store

import (
	"gsadmin/core/config"
	"sync"
)

var (
	storeClient StoreClient
	once        sync.Once
)

type StoreClient interface {
	UploadFile(dstFileName string, localFilePath string) (string, error)
	DeleteFile(dstFileName string) error //暂无需求,不实现
}

func Instance() StoreClient {
	if storeClient == nil {
		once.Do(func() {
			switch config.Instance().Store.StoreType {
			case "minio":
				storeClient = newMinioClient()
			case "oss":
				storeClient = newOssClient()
			case "local":
				storeClient = newLocalClient()
			default:
				storeClient = newLocalClient()
			}
		})
	}
	return storeClient
}
