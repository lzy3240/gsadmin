package service

import (
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/utils/str"
)

type SysNotice struct {
	baseservice.Service
}

// UpdateStatus 更新消息状态
func (s *SysNotice) UpdateStatus(req *dto.SysNoticeIDForm, userId uint) error {
	err := db.Instance().Scopes(s.SetCondition(*req)).
		Model(model.SysNotice{}).
		Where("notice_to = ?", userId).
		Update("status", 1).Error

	if err != nil {
		log.Instance().Error("SysNoticeService.UpdateStatus:" + err.Error())
		return err
	}
	return nil
}

// UpdateBatchStatus 批量更新消息状态
func (s *SysNotice) UpdateBatchStatus(req *dto.SysNoticeIDsForm, userId uint) error {
	err := db.Instance().
		Model(model.SysNotice{}).
		Where("notice_to = ? and id in (?)", userId, str.SplitStr(req.Ids)).
		Update("status", 1).
		Error
	if err != nil {
		log.Instance().Error("SysNoticeService.UpdateBatchStatus:" + err.Error())
		return err
	}
	return nil
}

// List 查询通知列表
func (s *SysNotice) List(req *dto.SysNoticeListForm, userId uint) (list []model.SysNotice, count int, err error) {
	err = db.Instance().Scopes(
		s.SetPaginate(req.Limit, req.Page), //默认分页
		s.SetCondition(*req),
	).
		Model(model.SysNotice{}).
		Where("notice_to = ?", userId).
		Order("id desc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error

	if err != nil {
		log.Instance().Error("SysNoticeService.List:" + err.Error())
		return nil, 0, err
	}
	return
}

// ListUnRead 未读消息列表
func (s *SysNotice) ListUnRead(userId uint) (list []model.SysNotice, count int, err error) {
	err = db.Instance().
		Model(model.SysNotice{}).
		Where("notice_to = ? and status = ?", userId, 0).
		Order("created_at desc").Count(&count).Find(&list).Error
	if err != nil {
		log.Instance().Error("SysNoticeService.ListUnRead:" + err.Error())
		return nil, 0, err
	}
	return
}

// Get 获取消息
func (s *SysNotice) Get(req *dto.SysNoticeIDForm, userId uint) (data model.SysNotice, err error) {
	err = db.Instance().Scopes(
		s.SetCondition(*req),
	).
		Model(model.SysNotice{}).Where("notice_to = ?", userId).First(&data).Error
	if err != nil {
		log.Instance().Error("SysNoticeService.Get:" + err.Error())
		return model.SysNotice{}, err
	}
	return
}

// Insert 新增消息
func (s *SysNotice) Insert(notice model.SysNotice) error {
	err := db.Instance().Create(&notice).Error
	if err != nil {
		log.Instance().Error("SysNoticeService.Insert:" + err.Error())
		return err
	}
	return nil
}
