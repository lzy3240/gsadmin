package model

import (
	"gsadmin/core/basemodel"
)

type SysDictType struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	DictType  string `json:"dict_type"`
	DictName  string `json:"dict_name"`
	Status    int    `json:"status"`
	Remark    string `json:"remark"`
	IsDefault int    `json:"is_default"`
	basemodel.Model
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}

type SysDictData struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	DictType  string `json:"dict_type"`
	DictLabel string `json:"dict_label"`
	DictValue string `json:"dict_value"`
	DictSort  int    `json:"dict_sort"`
	Status    int    `json:"status"`
	Remark    string `json:"remark"`
	IsDefault int    `json:"is_default"`
	basemodel.Model
}

func (SysDictData) TableName() string {
	return "sys_dict_data"
}
