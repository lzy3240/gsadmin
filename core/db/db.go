package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gsadmin/core/config"
	"gsadmin/core/log"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func Instance() *gorm.DB {
	if db == nil {
		db = InitConn()
	}
	return db
}

func InitConn() *gorm.DB {
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
