package store

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gsadmin/core/config"
	"gsadmin/core/log"
	"mime"
	"path/filepath"
)

type minioClient struct {
	client *minio.Client
	ctx    context.Context
}

func newMinioClient() *minioClient {
	mc, err := minio.New(config.Instance().Store.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Instance().Store.AccessKey, config.Instance().Store.AccessSecret, ""),
		Secure: false,
	})
	if err != nil {
		log.Instance().Error("Conn MinIO Failed..." + err.Error())
	}
	return &minioClient{client: mc, ctx: context.Background()}
}

func (m *minioClient) UploadFile(dstFileName string, localFilePath string) (string, error) {
	//检验桶是否存在
	if exists, err := m.client.BucketExists(m.ctx, config.Instance().Store.BucketName); !exists {
		err = m.client.MakeBucket(m.ctx, config.Instance().Store.BucketName, minio.MakeBucketOptions{Region: "us-east-1"}) //时区
		if err != nil {
			return "", err
		}
	}
	//上传文件
	fileType := mime.TypeByExtension(filepath.Ext(localFilePath))
	_, err := m.client.FPutObject(m.ctx, config.Instance().Store.BucketName, dstFileName, localFilePath, minio.PutObjectOptions{
		ContentType: fileType,
	})
	if err != nil {
		return "", err
	}
	//返回文件链接
	return config.Instance().Store.ShowPrefix + "/" + config.Instance().Store.BucketName + "/" + dstFileName, nil
}

func (m *minioClient) DeleteFile(dstFileName string) error {
	return nil
}
