package model

import (
	"gsadmin/core/basemodel"
)

type SysRole struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	RoleName string `json:"role_name"`
	Detail   string `json:"detail"`
	Status   int    `json:"status"`
	basemodel.Model
	RoleAuths []SysRoleAuth `json:"-" gorm:"foreignkey:RoleID"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

//// 管理员信息修改页面展示
//type RoleEditShow struct {
//	ID       int
//	RoleName string
//	Status   int
//	Checked  int
//}
