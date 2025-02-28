package dto

import "gsadmin/core/utils/assertion"

//Form: 提交表单
//Resp: 响应表单
//Show: 展示表单

type PageForm struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type ControlForm struct {
	CreateId int `json:"create_id" form:"create_id" column:"create_id"`
	UpdateId int `json:"update_id" form:"update_id" column:"update_id"`
}

func (c *ControlForm) SetCreate(userId uint) {
	c.CreateId = assertion.AnyToInt(userId)
	c.UpdateId = assertion.AnyToInt(userId)
}

func (c *ControlForm) SetUpdate(userId uint) {
	c.UpdateId = assertion.AnyToInt(userId)
}

type TimeForm struct {
	BeginTime string `json:"begin_time" form:"begin_time" search:"type:gte;column:create_at"`
	EndTime   string `json:"end_time" form:"end_time" search:"type:lte;column:create_at"`
}

type CaptchaVerifyForm struct {
	ID   string `json:"id" form:"id" binding:"required"`
	Capt string `json:"capt" form:"capt" binding:"required"`
}

type LoginForm struct {
	UserName string `json:"username" form:"username" binding:"required,min=3,max=30" zh:"用户名"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=30" zh:"密码"`
}

// --------------------------------------

type CacheMenuV2 struct {
	MenuResp []MenuResp `json:"menu_resp"`
	AllowUrl string     `json:"allow_url"`
}

type MenuResp struct {
	ID       int                `json:"id"`
	Title    string             `json:"title"`
	Type     int                `json:"type"`
	Icon     string             `json:"icon"`
	Href     string             `json:"href"`
	Children []MenuChildrenResp `json:"children"`
}

type MenuChildrenResp struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Type     int    `json:"type"`
	OpenType string `json:"openType"`
	Href     string `json:"href"`
}
