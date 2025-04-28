package service

import (
	"errors"
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/core/utils/assertion"
)

type SysDict struct {
	baseservice.Service
}

// 获取字典类型列表
func (s *SysDict) DictTypeListJson(req *dto.DictTypeListForm) (list []model.SysDictType, count int, err error) {
	err = db.Instance().Scopes(
		s.SetPaginate(req.Limit, req.Page),
		s.SetCondition(*req),
	).
		Model(model.SysDictType{}).Order("id asc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		log.Instance().Error("SysDictService.DictTypeListJson:" + err.Error())
		return nil, 0, err
	}
	return list, count, nil
}

// 获取字典类型信息
func (s *SysDict) GetDictTypeByID(req *dto.DictTypeIDForm) (data model.SysDictType, err error) {
	err = db.Instance().Model(model.SysDictType{}).Where("id = ?", req.ID).First(&data).Error
	if err != nil {
		log.Instance().Error("SysDictService.GetDictTypeByID:" + err.Error())
		return model.SysDictType{}, err
	}
	return data, nil
}

// 切换可用状态
func (s *SysDict) UpdateDictTypeStatus(req *dto.DictTypeStatusForm) error {
	//直接更新,不再先查后改
	err := db.Instance().Model(model.SysDictType{}).Where("id = ?", req.ID).Update("status", req.Status).Error
	if err != nil {
		log.Instance().Error("SysDictService.UpdateDictTypeStatus:" + err.Error())
		return err
	}
	return nil
}

// 修改字典属性
func (s *SysDict) UpdateDictTypeAttr(req *dto.DictTypeEditForm) error {
	attr := s.StructToMapByTag(*req, "column")
	attr["update_id"] = req.UpdateId

	//直接更新,不再先查后改
	err := db.Instance().Model(model.SysDictType{}).Where("id = ?", req.ID).Updates(attr).Error
	if err != nil {
		log.Instance().Error("UpdateDictTypeAttr.UpdateDictType:" + err.Error())
		return err
	}

	return nil
}

// 新增字典
func (s *SysDict) InsertDictType(req *dto.DictTypeEditForm) error {
	var dictType model.SysDictType
	dictType.DictType = req.DictType
	dictType.DictName = req.DictName
	dictType.Remark = req.Remark
	dictType.Status = assertion.AnyToInt(req.Status)
	dictType.CreateId = req.CreateId
	dictType.UpdateId = req.UpdateId

	err := db.Instance().Create(&dictType).Error
	if err != nil {
		log.Instance().Error("UpdateDictTypeAttr.InsertDictType:" + err.Error())
		return err
	}
	return nil
}

// 删除字典
func (s *SysDict) DeleteDictType(req *dto.DictTypeIDForm) error {
	var dictType model.SysDictType
	err := db.Instance().Model(model.SysDictType{}).Where("id = ?", req.ID).First(&dictType).Error
	if err != nil {
		log.Instance().Error("SysDictService.DictTypeDeleteService:" + err.Error())
		return err
	}
	if dictType.IsDefault == 1 {
		return errors.New("系统字段不能删除")
	}

	err = db.Instance().Where("id = ?", req.ID).Delete(model.SysDictType{}).Error
	if err != nil {
		log.Instance().Error("SysDictService.DictTypeDeleteService:" + err.Error())
		return err
	}
	return nil
}

// -------dict data operation---------

// 获取字典值列表
func (s *SysDict) DictDataListJson(req *dto.DictDataJsonForm) (list []model.SysDictData, count int, err error) {
	err = db.Instance().Scopes(s.SetCondition(*req)).Model(model.SysDictData{}).Count(&count).Find(&list).Error

	if err != nil {
		log.Instance().Error("SysDictService.DictDataListJson:" + err.Error())
		return nil, 0, err
	}
	return list, count, nil
}

func (s *SysDict) GetDictDataByID(req *dto.DictDataIDForm) (data model.SysDictData, err error) {
	err = db.Instance().Model(model.SysDictData{}).Where("id = ?", req.ID).First(&data).Error
	if err != nil {
		log.Instance().Error("SysDictService.GetDictDataByID:" + err.Error())
		return model.SysDictData{}, err
	}
	return data, nil
}

func (s *SysDict) InsertDictData(req *dto.DictDataEditForm) error {
	var dictData model.SysDictData
	dictData.DictType = req.DictType
	dictData.DictLabel = req.DictLabel
	dictData.DictValue = req.DictValue
	dictData.DictSort = req.DictSort
	dictData.Remark = req.Remark
	dictData.Status = assertion.AnyToInt(req.Status)
	dictData.CreateId = req.CreateId
	dictData.UpdateId = req.UpdateId

	err := db.Instance().Create(&dictData).Error
	if err != nil {
		log.Instance().Error("SysDictService.InsertDictData:" + err.Error())
		return err
	}

	return nil
}

func (s *SysDict) UpdateDictDataAttr(req *dto.DictDataEditForm) error {
	attr := s.StructToMapByTag(*req, "column")
	attr["update_id"] = req.UpdateId

	err := db.Instance().Model(model.SysDictData{}).Where("id = ?", req.ID).Updates(attr).Error
	if err != nil {
		log.Instance().Error("SysDictService.UpdateDictDataAttr:" + err.Error())
		return err
	}
	return nil
}

func (s *SysDict) UpdateDictDataStatus(req *dto.DictDataStatusForm) error {
	err := db.Instance().Model(model.SysDictData{}).Where("id = ?", req.ID).Update("status", req.Status).Error
	if err != nil {
		log.Instance().Error("SysDictService.UpdateDictDataStatus:" + err.Error())
		return err
	}
	return nil
}

func (s *SysDict) DeleteDictData(req *dto.DictDataIDForm) error {
	err := db.Instance().Where("id = ?", req.ID).Delete(model.SysDictData{}).Error
	if err != nil {
		log.Instance().Error("SysDictService.DeleteDictData:" + err.Error())
		return err
	}
	return nil
}
