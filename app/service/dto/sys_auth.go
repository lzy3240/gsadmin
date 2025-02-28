package dto

type SysAuthListForm struct {
	PowerType string `json:"power_type" form:"power_type" search:"type:exact;column:power_type;table:sys_auth"` //类型必须string,防止int被忽略
	Pid       string `json:"parentId" form:"parentId" search:"type:exact;column:pid;table:sys_auth"`            //类型必须string,防止int被忽略
}

type SysAuthIDForm struct {
	ID string `json:"id" form:"id" search:"type:exact;column:id;table:sys_auth"`
}

// 更新参数, 未使用, 直接使用model
//type SysAuthNodeForm struct {
//	ID        string `json:"id" form:"id" column:"id"`
//	Pid       string `json:"parentId" form:"parentId"  column:"pid"`
//	AuthName  string `json:"auth_name" form:"auth_name"  column:"auth_name"`
//	AuthUrl   string `json:"auth_url" form:"auth_url"  column:"auth_url"`
//	PowerType string `json:"power_type" form:"power_type"  column:"power_type"`
//	Sort      string `json:"sort" form:"sort"  column:"sort"`
//	IsShow    string `json:"is_show" form:"is_show"  column:"is_show"`
//	Icon      string `json:"icon" form:"icon"  column:"icon"`
//}

//type SysAuthListResp struct {
//	CateId   int    `json:"cate_id"`
//	CateName string `json:"cate_name"`
//	MenuId   int    `json:"menu_id"`
//	MenuName string `json:"menu_name"`
//	MenuUrl  string `json:"menu_url"`
//}

type NodeResp struct {
	ID        int    `json:"id"`
	Pid       int    `json:"pid"`
	AuthName  string `json:"auth_name"`
	AuthUrl   string `json:"auth_url"`
	Sort      int    `json:"sort"`
	IsShow    int    `json:"is_show"`
	Icon      string `json:"icon"`
	PowerType int    `json:"power_type"`
}

type NodesResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Open bool   `json:"open"`
	Pid  int    `json:"pId"`
}

type AuthResp struct {
	NodesResp
	PowerType int `json:"power_type"`
}

type FrontAuthResp struct {
	PowerID   string `json:"powerId"`
	PowerName string `json:"powerName"`
	PowerType string `json:"powerType"`
	PowerCode string `json:"powerCode"`
	PowerURL  string `json:"powerUrl"`
	OpenType  string `json:"openType"`
	ParentID  string `json:"parentId"`
	Icon      string `json:"icon"`
	Sort      int    `json:"sort"`
	CheckArr  string `json:"checkArr"`
}

type RolePower struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data []RolePowerData `json:"data"`
}

type RolePowerData struct {
	CheckArr   string      `json:"checkArr,omitempty"`
	CreateTime interface{} `json:"create_time,omitempty"`
	Enable     int         `json:"enable,omitempty"`
	Icon       string      `json:"icon,omitempty"`
	OpenType   interface{} `json:"openType,omitempty"`
	ParentID   string      `json:"parentId"`
	PowerID    string      `json:"powerId"`
	PowerName  string      `json:"powerName"`
	PowerType  string      `json:"powerType,omitempty"`
	PowerURL   interface{} `json:"powerUrl,omitempty"`
	Sort       int         `json:"sort,omitempty"`
	UpdateTime interface{} `json:"update_time,omitempty"`
}
