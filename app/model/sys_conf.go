package model

import (
	"gsadmin/core/basemodel"
)

type SysConf struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Type   string `json:"type" gorm:"unique"`
	Info   string `json:"info" gorm:"type:varchar(3000);"`
	Status int    `json:"status"`
	basemodel.Model
}

func (SysConf) TableName() string {
	return "sys_conf"
}

// 基础样式
type Base struct {
	Colors []Colors `json:"colors"`
	Header Header   `json:"header"`
	Links  []Links  `json:"links"`
	Logo   Logo     `json:"logo"`
	Menu   Menu     `json:"menu"`
	Other  Other    `json:"other"`
	Tab    Tab      `json:"tab"`
	Theme  Theme    `json:"theme"`
}

type Menu struct {
	Accordion bool   `json:"accordion"`
	Control   bool   `json:"control"`
	Data      string `json:"data"`
	Method    string `json:"method"`
	Select    string `json:"select"`
}

type Theme struct {
	AllowCustom  bool   `json:"allowCustom"`
	DefaultColor string `json:"defaultColor"`
	DefaultMenu  string `json:"defaultMenu"`
}

type Other struct {
	AutoHead bool  `json:"autoHead"`
	KeepLoad int64 `json:"keepLoad"`
}

type Colors struct {
	Color string `json:"color"`
	ID    string `json:"id"`
}

type Links struct {
	Href  string `json:"href"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
}

type Logo struct {
	Image string `json:"image"`
	Title string `json:"title"`
}

type Tab struct {
	Index     TabIndex `json:"index"`
	KeepState bool     `json:"keepState"`
	MuiltTab  bool     `json:"muiltTab"`
	TabMax    int64    `json:"tabMax"`
}

type TabIndex struct {
	Href  string `json:"href"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Header struct {
	Message string `json:"message"`
}

// Site 网站信息
type Site struct {
	WebName     string `json:"web_name"`
	WebUrl      string `json:"web_url"`
	LogoUrl     string `json:"logo_url"`
	KeyWords    string `json:"key_words"`
	Description string `json:"description"`
	Copyright   string `json:"copyright"`
	Icp         string `json:"icp"`
	SiteStatus  uint8  `json:"site_status"`
}
