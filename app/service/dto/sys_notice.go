package dto

type SysNoticeListForm struct {
	PageForm    `search:"-"`
	NoticeType  string `json:"notice_type" form:"notice_type" search:"type:contains;column:notice_type;table:sys_notice"`
	NoticeTitle string `json:"notice_title" form:"notice_title" search:"type:contains;column:notice_title;table:sys_notice"`
	Status      string `json:"status" form:"status" search:"type:exact;column:status;table:sys_notice"`
}

type SysNoticeIDForm struct {
	Id string `json:"id" form:"id" search:"type:exact;column:id;table:sys_notice"`
}

type SysNoticeIDsForm struct {
	Ids string `json:"ids" form:"ids"`
}
