package model

type SysOperLog struct {
	ID            uint64 `json:"id" gorm:"primary_key"`
	Title         string `json:"title"`
	BusinessType  int    `json:"business_type"`
	Method        string `json:"method"`
	RequestMethod string `json:"request_method"`
	OperatorType  int    `json:"operator_type"`
	UserId        int    `json:"user_id"`
	OperName      string `json:"oper_name"`
	DeptName      string `json:"dept_name"`
	OperUrl       string `json:"oper_url"`
	OperIp        string `json:"oper_ip"`
	OperLocation  string `json:"oper_location"`
	OperParam     string `json:"oper_param" gorm:"type:longtext;"`
	Status        int    `json:"status"`
	JsonResult    string `json:"json_result"`
	ErrorMsg      string `json:"error_msg"`
	LatencyTime   int    `json:"latency_time"`
	UserAgent     string `json:"user_agent"`
	OperTime      string `json:"oper_time"`
}

func (SysOperLog) TableName() string {
	return "sys_oper_log"
}
