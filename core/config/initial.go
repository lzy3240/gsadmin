package config

const configToml = `
# App运行配置
[app]
Name = "gsadmin"
Version = "V1.0.0"
HttpPort = 8010
BaseURL = ""
PageSize = 20
RunMode = "debug" #运行模式: debug; release
CacheMode = "mem" #缓存模式: mem; redis; nuts; 默认mem
QueueMode = "mem" #队列模式: mem; redis; 默认mem
JwtSecret = "0102$%#&*^*&150405"
FileSavePath = "runtime/file"
FileViewPath = "runtime/upload"

# db配置 以下 两者任选其一 不要同时使用
[db]
DBType = "mysql" # "sqlite" 或 "mysql"
DBName = "gsadmin"
DBUser = "root"
DBPwd = "Aa111111"
DBHost = "127.0.0.1:3306"

# redis配置
[redis]
RedisAddr = "127.0.0.1:6379"
RedisPWD = ""
RedisDB = 9

# store配置
[store]
StoreType = "" #文件存储模式, minio:MinIO; oss:阿里云OSS; local:本地存储; 默认为空时本地存储; minio和oss需要打开公开只读授权
EndPoint = "" #存储服务端URL
AccessKey = ""
AccessSecret = ""
BucketName = "" #存储桶名称
ShowPrefix = "" #链接地址前缀, minio: http://minio_public_url/bucketname/file/to/name; oss: https://bucketname.endpoint/file/to/name

# log配置
[zaplog]
director = 'runtime/log'
SaveMode = "db" #系统日志记录模式, file:文件; db:数据库; 默认为file

# chptchar配置
[captchar]
ImgHeight = 80
ImgWidth = 240
ImgKeyLength = 4
`
