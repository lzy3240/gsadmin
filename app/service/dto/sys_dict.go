package dto

type DictTypeListForm struct {
	PageForm `search:"-"`
	ID       string `json:"id" form:"id" search:"type:exact;column:id;table:sys_dict_type"`
	DictType string `json:"dict_type" form:"dict_type" search:"type:contains;column:dict_type;table:sys_dict_type"`
	DictName string `json:"dict_name" form:"dict_name" search:"type:contains;column:dict_name;table:sys_dict_type"`
	Status   string `json:"status" form:"status" search:"type:exact;column:status;table:sys_dict_type"`
}

type DictTypeIDForm struct {
	ID string `json:"id" form:"id" binding:"required" search:"type:exact;column:id;table:sys_dict_type"`
}

type DictTypeStatusForm struct {
	ID     string `json:"id" form:"id" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}

type DictTypeEditForm struct {
	ControlForm
	ID       string `json:"id" form:"id" column:"id"`
	DictType string `json:"dict_type" form:"dict_type" column:"dict_type"`
	DictName string `json:"dict_name" form:"dict_name" column:"dict_name"`
	Remark   string `json:"remark" form:"remark" column:"remark"`
	Status   string `json:"status" form:"status" column:"status"`
}

// ----------dict data----------------------------

type DictDataJsonForm struct {
	DictType string `json:"dict_type" form:"dict_type" binding:"required" search:"type:exact;column:dict_type;table:sys_dict_data"`
}

type DictDataIDForm struct {
	ID string `json:"id" form:"id" binding:"required" search:"type:exact;column:id;table:sys_data_type"`
}

type DictDataStatusForm struct {
	ID     string `json:"id" form:"id" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}

type DictDataEditForm struct {
	ControlForm
	ID        string `json:"id" form:"id" column:"id"`
	DictType  string `json:"dict_type" form:"dict_type" column:"dict_type"`
	DictLabel string `json:"dict_label" form:"dict_label" column:"dict_label"`
	DictValue string `json:"dict_value" form:"dict_value" column:"dict_value"`
	DictSort  int    `json:"dict_sort" form:"dict_sort" column:"dict_sort"`
	Remark    string `json:"remark" form:"remark" column:"remark"`
	Status    string `json:"status" form:"status" column:"status"`
}
