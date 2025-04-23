package config

import (
	"github.com/spf13/viper"
	"gsadmin/core/utils/file"
	"io/ioutil"
	"log"
)

var c *Conf

func Instance() *Conf {
	if c == nil {
		InitConfig("./config.toml")
	}
	return c
}

type Conf struct {
	App      AppConf
	DB       DBConf
	Redis    RedisConf
	Store    StoreConf
	ZapLog   ZapLogConf
	CaptChar CaptCharConf
}

type AppConf struct {
	Name         string
	Version      string
	HttpPort     int
	BaseURL      string
	RunMode      string
	CacheMode    string
	QueueMode    string
	PageSize     int
	JwtSecret    string
	FileSavePath string
	FileViewPath string
}

type DBConf struct {
	DBType string
	DBUser string
	DBPwd  string
	DBHost string
	DBName string
}

type RedisConf struct {
	RedisAddr string
	RedisPWD  string
	RedisDB   int
}

type StoreConf struct {
	StoreType    string
	EndPoint     string
	AccessKey    string
	AccessSecret string
	BucketName   string
	ShowPrefix   string
}

type ZapLogConf struct {
	Director string
	SaveMode string
}

type CaptCharConf struct {
	ImgHeight    int
	ImgWidth     int
	ImgKeyLength int
}

func InitConfig(tomlPath ...string) {
	if len(tomlPath) > 1 {
		log.Fatal("配置路径数量不正确")
	}
	if file.CheckNotExist(tomlPath[0]) {
		err := ioutil.WriteFile(tomlPath[0], []byte(configToml), 0777)
		if err != nil {
			log.Fatal("无法写入配置模板", err.Error())
		}
		log.Fatal("配置文件不存在，已在根目录下生成示例模板文件【config.toml】，请修改后重新启动程序！")
	}
	v := viper.New()
	v.SetConfigFile(tomlPath[0])
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("配置文件读取失败: ", err.Error())
	}
	err = v.Unmarshal(&c)
	if err != nil {
		log.Fatal("配置解析失败:", err.Error())
	}
}
