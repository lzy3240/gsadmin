package model

import (
	"gsadmin/core/basemodel"
)

type SysNotice struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	NoticeType    string `json:"notice_type"`
	NoticeTitle   string `json:"notice_title"`
	NoticeContext string `json:"notice_context"`
	NoticeFrom    string `json:"notice_from"`
	NoticeTo      uint   `json:"notice_to"`
	Status        int    `json:"status"`
	basemodel.Model
}

func (SysNotice) TableName() string {
	return "sys_notice"
}
