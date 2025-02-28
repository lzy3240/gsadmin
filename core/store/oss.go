package store

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gsadmin/core/config"
	"gsadmin/core/log"
)

type ossClient struct {
	client *oss.Client
	ctx    context.Context
}

func newOssClient() *ossClient {
	oc, err := oss.New(config.Instance().Store.EndPoint, config.Instance().Store.AccessKey, config.Instance().Store.AccessSecret)
	if err != nil {
		log.Instance().Error("Conn OSS Failed..." + err.Error())
	}
	return &ossClient{client: oc, ctx: context.Background()}
}

func (o *ossClient) UploadFile(dstFileName string, localFilePath string) (string, error) {
	//检验桶是否存在
	if exists, err := o.client.IsBucketExist(config.Instance().Store.BucketName); !exists {
		err = o.client.CreateBucket(config.Instance().Store.BucketName)
		if err != nil {
			return "", err
		}
	}
	//获取bucket对象
	bucket, err := o.client.Bucket(config.Instance().Store.BucketName)
	if err != nil {
		return "", err
	}
	//上传文件
	err = bucket.PutObjectFromFile(dstFileName, localFilePath)
	if err != nil {
		return "", err
	}
	//返回文件链接地址
	return config.Instance().Store.ShowPrefix + "/" + dstFileName, nil
}

func (o *ossClient) DeleteFile(dstFileName string) error {
	return nil
}
