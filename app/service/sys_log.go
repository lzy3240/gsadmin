package service

import (
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/db"
	"gsadmin/core/log"
)

type SysLog struct {
	baseservice.Service
}

func (s *SysLog) OperLogList(req *dto.SysOperLogForm) (list []model.SysOperLog, count int, err error) {
	err = db.Instance().Scopes(
		s.SetPaginate(req.Limit, req.Page),
		s.SetCondition(*req),
	).
		Model(model.SysOperLog{}).
		Order("id desc").Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		log.Instance().Error("SysLogService.OperLogList:" + err.Error())
		return nil, 0, err
	}

	return list, count, nil
}

func (s *SysLog) LoginLogList(req *dto.SysLoginLogForm) (list []model.SysLoginLog, count int, err error) {
	err = db.Instance().Scopes(
		s.SetPaginate(req.Limit, req.Page),
		s.SetCondition(*req),
	).
		Model(model.SysLoginLog{}).
		Order("id desc").Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		log.Instance().Error("SysLogService.LoginLogList:" + err.Error())
		return nil, 0, err
	}

	return list, count, nil
}

func (s *SysLog) InsertSysLoginLog(logInfo model.SysLoginLog) error {
	err := db.Instance().Create(&logInfo).Error
	if err != nil {
		log.Instance().Error("SysLogService.InsertSysLoginLog:" + err.Error())
		return err
	}
	return nil
}
