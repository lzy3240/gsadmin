package model

import (
	"gsadmin/core/basemodel"
)

type SysUser struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	LoginName string `json:"login_name"`
	RealName  string `json:"real_name"`
	Password  string `json:"-"`
	Level     int    `json:"level"`
	RoleIds   string `json:"role_ids"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Remark    string `json:"remark"`
	Salt      string `json:"-"`
	LastIp    string `json:"-"`
	LastLogin string `json:"last_login"`
	Status    int    `json:"status"`
	basemodel.Model
}

func (SysUser) TableName() string {
	return "sys_user"
}
