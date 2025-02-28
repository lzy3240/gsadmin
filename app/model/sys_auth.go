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

//type NodesResp struct {
//	ID   int    `json:"id"`
//	Name string `json:"name"`
//	Open bool   `json:"open"`
//	Pid  int    `json:"pId"`
//}

//type NodeResp struct {
//	ID        int    `json:"id"`
//	Pid       int    `json:"pid"`
//	AuthName  string `json:"auth_name"`
//	AuthUrl   string `json:"auth_url"`
//	Sort      int    `json:"sort"`
//	IsShow    int    `json:"is_show"`
//	Icon      string `json:"icon"`
//	PowerType int    `json:"power_type"`
//}

//type AuthResp struct {
//	NodesResp
//	PowerType int `json:"power_type"`
//}

//type FrontAuthResp struct {
//	PowerID   string `json:"powerId"`
//	PowerName string `json:"powerName"`
//	PowerType string `json:"powerType"`
//	PowerCode string `json:"powerCode"`
//	PowerURL  string `json:"powerUrl"`
//	OpenType  string `json:"openType"`
//	ParentID  string `json:"parentId"`
//	Icon      string `json:"icon"`
//	Sort      int    `json:"sort"`
//	CheckArr  string `json:"checkArr"`
//}

//type AuthListResp struct {
//	CateId   int    `json:"cate_id"`
//	CateName string `json:"cate_name"`
//	MenuId   int    `json:"menu_id"`
//	MenuName string `json:"menu_name"`
//	MenuUrl  string `json:"menu_url"`
//}

//type RolePower struct {
//	Status struct {
//		Code    int    `json:"code"`
//		Message string `json:"message"`
//	} `json:"status"`
//	Data []RolePowerData `json:"data"`
//}

//type RolePowerData struct {
//	CheckArr   string      `json:"checkArr,omitempty"`
//	CreateTime interface{} `json:"create_time,omitempty"`
//	Enable     int         `json:"enable,omitempty"`
//	Icon       string      `json:"icon,omitempty"`
//	OpenType   interface{} `json:"openType,omitempty"`
//	ParentID   string      `json:"parentId"`
//	PowerID    string      `json:"powerId"`
//	PowerName  string      `json:"powerName"`
//	PowerType  string      `json:"powerType,omitempty"`
//	PowerURL   interface{} `json:"powerUrl,omitempty"`
//	Sort       int         `json:"sort,omitempty"`
//	UpdateTime interface{} `json:"update_time,omitempty"`
//}
