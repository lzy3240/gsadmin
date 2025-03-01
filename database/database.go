package database

import (
	"github.com/jinzhu/gorm"
	"github.com/qustavo/dotsql"
	"gsadmin/app/model"
	"gsadmin/core/config"
	"gsadmin/core/db"
	"gsadmin/core/log"
)

func InitTables() {
	//init system tables
	checkTableData(&model.SysUser{})
	checkTableData(&model.SysUserOnline{})
	checkTableData(&model.SysAuth{})
	checkTableData(&model.SysLoginLog{})
	checkTableData(&model.SysOperLog{})
	checkTableData(&model.SysRole{})
	checkTableData(&model.SysRoleAuth{})
	checkTableData(&model.SysConf{})
	checkTableData(&model.SysDictType{})
	checkTableData(&model.SysDictData{})
	checkTableData(&model.SysNotice{})
	//init business tables
}

func checkTableData(tb interface{}) {
	DB := db.Instance()
	if DB.HasTable(tb) == false {
		if config.Instance().DB.DBType == "sqlite" {
			if err := DB.Debug().CreateTable(tb).Error; err != nil {
				log.Instance().Fatal("创建数据表失败: " + err.Error())
			}
		} else {
			if err := DB.Debug().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(tb).Error; err != nil {
				log.Instance().Fatal("创建数据表失败: " + err.Error())
			}
		}
		var sqlName string
		if _, ok := tb.(*model.SysUser); ok {
			sqlName = "create-sys-user"
		} else if _, ok := tb.(*model.SysAuth); ok {
			sqlName = "create-sys-auth"
		} else if _, ok := tb.(*model.SysRole); ok {
			sqlName = "create-sys-role"
		} else if _, ok := tb.(*model.SysRoleAuth); ok {
			sqlName = "create-sys-role-auth"
		} else if _, ok := tb.(*model.SysConf); ok {
			sqlName = "create-sys-conf"
		} else if _, ok := tb.(*model.SysDictType); ok {
			sqlName = "create-sys-dict-type"
		} else if _, ok := tb.(*model.SysDictData); ok {
			sqlName = "create-sys-dict-data"
		}

		if sqlName != "" {
			initData(DB, sqlName)
		}
	} else {
		// 已存在的表校验一下是否有新增字段
		if config.Instance().DB.DBType == "sqlite" {
			if err := DB.Debug().AutoMigrate(tb).Error; err != nil {
				log.Instance().Fatal("数据库初始化失败: " + err.Error())
			}
		} else {
			if err := DB.Debug().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(tb).Error; err != nil {
				log.Instance().Fatal("数据库初始化失败: " + err.Error())
			}
		}
	}
}

func initData(db *gorm.DB, sqlName string) {
	//dot, err := dotsql.LoadFromString(databaseInfo)
	dot, err := dotsql.LoadFromFile("database/db.sql")
	if err != nil {
		log.Instance().Fatal("无法加载初始数据")
		return
	}
	_, err = dot.Exec(db.DB(), sqlName)
	if err != nil {
		log.Instance().Fatal("执行 " + sqlName + " 失败，" + err.Error())
		return
	}
}
