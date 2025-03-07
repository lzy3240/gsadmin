package model

import "time"

type SysLoginLog struct {
	ID            int64     `json:"id" gorm:"primary_key"`
	LoginName     string    `json:"login_name"`
	IpAddr        string    `json:"ip_addr"`
	LoginLocation string    `json:"login_location"`
	Browser       string    `json:"browser"`
	Os            string    `json:"os"`
	Status        string    `json:"status"`
	Msg           string    `json:"msg"`
	LoginTime     time.Time `json:"login_time"`
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}
