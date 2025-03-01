package model

import (
	"gsadmin/core/basemodel"
)

type SysAuth struct {
	ID        uint   `json:"id" form:"id" gorm:"primary_key"`
	AuthName  string `json:"auth_name" form:"auth_name"`
	AuthUrl   string `json:"auth_url" form:"auth_url"`
	UserId    int    `json:"user_id" form:"user_id"`
	Pid       int    `json:"pid" form:"parentId"`
	Sort      int    `json:"sort" form:"sort"`
	Icon      string `json:"icon" form:"icon"`
	IsShow    int    `json:"is_show" form:"is_show"`
	Status    int    `json:"status" form:"status"`
	PowerType int    `json:"power_type" form:"power_type"`
	IsDefault int    `json:"is_default" form:"is_default"`
	basemodel.Model
}

func (SysAuth) TableName() string {
	return "sys_auth"
}
