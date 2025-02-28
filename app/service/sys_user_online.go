package service

import (
	"gsadmin/app/model"
	"gsadmin/core/baseservice"
	"gsadmin/core/db"
	"gsadmin/core/log"
)

type SysUserOnline struct {
	baseservice.Service
}

func (s *SysUserOnline) Insert(online model.SysUserOnline) error {
	err := db.Instance().Create(&online).Error
	if err != nil {
		log.Instance().Error("SysUserOnlineService.Insert:" + err.Error())
	}
	return nil
}

func (s *SysUserOnline) Delete(sessionID string) error {
	err := db.Instance().Where("session_id = ?", sessionID).Delete(model.SysUserOnline{}).Error
	if err != nil {
		log.Instance().Error("SysUserOnlineService.Delete:" + err.Error())
	}
	return nil
}
