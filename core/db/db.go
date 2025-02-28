package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gsadmin/core/config"
	"gsadmin/core/log"
	"gsadmin/core/utils/file"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func Instance() *gorm.DB {
	if db == nil {
		InitConn()
	}
	return db
}

func InitConn() {
	switch config.Instance().DB.DBType {
	case "mysql":
		db = GormMysql()
	case "sqlite":
		db = GormSqlite()
	default:
		log.Instance().Fatal("No DBType")
		os.Exit(0)
	}
}

func GormMysql() *gorm.DB {
	m := config.Instance().DB
	if m.DBName == "" {
		return nil
	}
	dsn := m.DBUser + ":" + m.DBPwd + "@tcp(" + m.DBHost + ")/" + m.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Instance().Fatal("MySQL启动异常: " + err.Error())
		os.Exit(0)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(300)
	db.SingularTable(true)
	db.LogMode(true)

	return db
}

func GormSqlite() *gorm.DB {
	dbFile := fmt.Sprintf("%s.db", config.Instance().DB.DBName)
	if file.CheckNotExist(dbFile) {
		if err = createDB(dbFile); err != nil {
			log.Instance().Fatal("创建数据库文件失败: " + err.Error())
		}
	}
	db, err = gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Instance().Fatal("连接数据库失败：" + err.Error())
		os.Exit(0)
	}
	db.SingularTable(true)
	db.LogMode(true)

	return db
}

func createDB(path string) error {
	fp, err := os.Create(path) // 如果文件已存在，会将文件清空。
	if err != nil {
		return err
	}
	defer fp.Close() //关闭文件，释放资源。
	return nil
}
