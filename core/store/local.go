package store

import (
	"context"
	"gsadmin/core/config"
	"gsadmin/core/utils/sysos"
	"gsadmin/global/e"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type localClient struct {
	ctx context.Context
}

func newLocalClient() *localClient {
	return &localClient{
		ctx: context.Background(),
	}
}

func (l *localClient) UploadFile(dstFileName string, localFilePath string) (string, error) {
	backFilePath := ""
	//本地存储时, 仅转换文件显示地址(local)
	_, err := os.Stat(localFilePath)
	if err == nil || os.IsExist(err) {
		day := time.Now().Format(e.TimeFormatDay)
		backFilePath = filepath.Join(filepath.Join(config.Instance().App.FileViewPath, day), dstFileName)
	} else {
		//考虑保存时间跨天时文件夹路径, 查看前一天的文件是否存在
		day := time.Now().Add(-24 * time.Hour).Format(e.TimeFormatDay)
		_, err = os.Stat(filepath.Join(filepath.Join(config.Instance().App.FileSavePath, day), dstFileName))
		if err == nil || os.IsExist(err) {
			backFilePath = filepath.Join(filepath.Join(config.Instance().App.FileViewPath, day), dstFileName)
		} else {
			return "", err
		}
	}

	if sysos.IsWindows() {
		backFilePath = strings.ReplaceAll(backFilePath, "\\", "/")
	}

	return backFilePath, nil
}

func (l *localClient) DeleteFile(dstFileName string) error {
	return nil
}
