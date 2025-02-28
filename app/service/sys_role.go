package service

import (
	"errors"
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/cache"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/core/utils/assertion"
	"gsadmin/global/e"
	"strings"
)

type SysRole struct {
	baseservice.Service
}

func (s *SysRole) ListByStatus(status int) (list []model.SysRole, count int, err error) {
	err = db.Instance().Model(model.SysRole{}).Where("status = ?", status).Find(&list).Count(&count).Error
	if err != nil {
		log.Instance().Error("SysRoleService.ListByStatus:" + err.Error())
		return nil, 0, err
	}
	return
}

func (s *SysRole) List(req *dto.SysRoleListForm) (list []dto.SysRoleResp, count int, err error) {
	var roles []model.SysRole
	err = db.Instance().Scopes(
		s.SetPaginate(req.Limit, req.Page),
		s.SetCondition(*req),
	).Model(model.SysRole{}).
		Order("id desc").Count(&count).Limit(-1).Offset(-1).Find(&roles).Error
	if err != nil {
		log.Instance().Error("SysRoleService.List:" + err.Error())
		return nil, 0, err
	}

	//过滤
	for _, v := range roles {
		list = append(list, dto.SysRoleResp{
			ID:        assertion.AnyToInt(v.ID),
			RoleName:  v.RoleName,
			Detail:    v.Detail,
			CreatedAt: v.CreatedAt.Format(e.TimeFormat),
			UpdatedAt: v.UpdatedAt.Format(e.TimeFormat),
			Status:    v.Status,
		})
	}
	return list, count, nil
}

func (s *SysRole) Get(req *dto.SysRoleIDForm) (model.SysRole, error) {
	var role model.SysRole
	err := db.Instance().
		Scopes(s.SetCondition(*req)).Preload("RoleAuths").
		Model(model.SysRole{}).First(&role).Error
	if err != nil {
		log.Instance().Error("SysRoleService.Get:" + err.Error())
		return model.SysRole{}, err
	}
	return role, nil
}

func (s *SysRole) Insert(req *dto.SysRoleAddForm) error {
	var role model.SysRole
	err := db.Instance().Model(model.SysRole{}).Where("role_name = ?", req.RoleName).First(&role).Error
	if err == nil || role.ID > 0 {
		return errors.New("角色名已存在")
	}

	role.RoleName = req.RoleName
	role.Detail = req.Detail
	role.CreateId = req.CreateId
	role.UpdateId = req.UpdateId
	role.Status = req.Status

	err = db.Instance().Create(&role).Error
	if err != nil {
		log.Instance().Error("SysRoleService.Insert:" + err.Error())
		return err
	}

	roleAuthSvice := SysRoleAuth{}
	if strings.HasSuffix(req.NodesData, ",") {
		req.NodesData = string([]rune(req.NodesData)[:len(req.NodesData)-1])
	}
	nodesArr := strings.Split(req.NodesData, ",")
	for _, v := range nodesArr {
		err = roleAuthSvice.Insert(model.SysRoleAuth{
			RoleID: assertion.AnyToUint64(role.ID),
			AuthID: assertion.AnyToUint(v),
		})

		if err != nil {
			log.Instance().Error("SysRoleService.Insert:" + err.Error())
			return err
		}
	}
	return nil
}

func (s *SysRole) Update(req *dto.SysRoleEditForm) error {
	attr := s.StructToMapByTag(*req, "column")
	attr["update_id"] = req.UpdateId
	err := db.Instance().Model(model.SysRole{}).Where("id = ?", req.ID).Updates(attr).Error
	if err != nil {
		log.Instance().Error("SysRoleService.Update:" + err.Error())
		return err
	}
	return nil
}

func (s *SysRole) Delete(req *dto.SysAuthIDForm) error {
	if req.ID == "1" {
		return errors.New("超级管理员角色不能删除")
	}
	err := db.Instance().Where("id = ?", req.ID).Delete(model.SysRole{}).Error
	if err != nil {
		log.Instance().Error("SysRoleService.Delete:" + err.Error())
		return err
	}
	//删除权限关联表
	roleAuthSvice := SysRoleAuth{}
	err = roleAuthSvice.Delete("role_id = ?", assertion.AnyToString(req.ID))
	return err
}

func (s *SysRole) InsertRoleAuth(req *dto.SysRolePowerForm) (err error) { //SaveRoleAuth roleId, authIds string
	authIdMap := strings.Split(req.PowerIds, ",")
	if len(authIdMap) < 1 {
		return errors.New("权限分配出错")
	}
	roleAuthSvice := SysRoleAuth{}
	if err = roleAuthSvice.Delete("role_id = ?", req.RoleId); err != nil {
		return err
	}

	//if err = dao.NewSysRoleAuthDaoImpl().Delete("role_id = ?", assertion.AnyToString(roleId)); err != nil {
	//	return err
	//}
	//var roleAuth model.SysRoleAuth

	tx := db.Instance().Begin()
	for _, v := range authIdMap {
		err = roleAuthSvice.Insert(model.SysRoleAuth{
			AuthID: assertion.AnyToUint(v),
			RoleID: assertion.AnyToUint64(req.RoleId),
		})
		if err != nil {
			log.Instance().Warn("InsertRoleAuth.Insert error:" + err.Error())
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	deleteMenuCache()
	return nil
}

func (s *SysRole) UpdateRoleStatus(req *dto.SysRoleStatusForm) error {
	err := db.Instance().Model(model.SysRole{}).Where("id = ?", req.ID).Update("status", req.Status).Error
	if err != nil {
		log.Instance().Error("SysRoleService.UpdateRoleStatus:" + err.Error())
		return err
	}
	return nil
}

func deleteMenuCache() {
	items := cache.Instance().Items()
	for k, _ := range items {
		if strings.HasPrefix(k, e.MenuCache) {
			cache.Instance().Delete(k)
		}
	}
}
