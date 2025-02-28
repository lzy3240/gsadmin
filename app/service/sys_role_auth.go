package service

import (
	"gsadmin/app/model"
	"gsadmin/core/baseservice"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/core/utils/assertion"
	"strings"
)

type SysRoleAuth struct {
	baseservice.Service
}

func (s *SysRoleAuth) Delete(key string, value string) error {
	err := db.Instance().Where(key, value).Delete(model.SysRoleAuth{}).Error
	if err != nil {
		log.Instance().Error("SysRoleAuthService.Delete:" + err.Error())
		return err
	}
	return nil
}

func (s *SysRoleAuth) ListByRoleId(roleId int) (list []model.SysRoleAuth, err error) {
	err = db.Instance().Model(model.SysRoleAuth{}).Where("role_id = ?", roleId).Find(&list).Error
	if err != nil {
		log.Instance().Error("SysRoleAuthService.Delete:" + err.Error())
		return nil, err
	}
	return
}

func (s *SysRoleAuth) ListByRoleIDs(roleIds string) (authIds []string, err error) {
	var roleAuths []model.SysRoleAuth
	err = db.Instance().Model(model.SysRoleAuth{}).Where("role_id in (?)", strings.Split(roleIds, ",")).Find(&roleAuths).Error
	if err != nil {
		log.Instance().Error("SysRoleAuthService.Delete:" + err.Error())
		return nil, err
	}

	for _, ra := range roleAuths {
		if ra.AuthID != 0 && ra.AuthID != 1 {
			tmp := assertion.AnyToString(ra.AuthID)
			authIds = append(authIds, tmp)
		}
	}
	return
}

func (s *SysRoleAuth) Insert(auth model.SysRoleAuth) error {
	if auth.AuthID != 0 {
		err := db.Instance().
			Model(model.SysRoleAuth{}).Where("role_id = ? and auth_id = ?", auth.RoleID, auth.AuthID).FirstOrCreate(&auth).Error
		if err != nil {
			log.Instance().Error("SysRoleAuthService.Insert:" + err.Error())
			return err
		}
	}

	return nil
}
