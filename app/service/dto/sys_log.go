package dto

type SysLoginLogForm struct {
	PageForm  `search:"-"`
	LoginName string `json:"login_name" form:"login_name" search:"type:exact;column:login_name;table:sys_login_log"`
}

type SysOperLogForm struct {
	PageForm `search:"-"`
	OperName string `json:"oper_name" form:"oper_name" search:"type:exact;column:oper_name;table:sys_oper_log"`
}
