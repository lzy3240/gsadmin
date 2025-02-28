package service

import (
	"errors"
	"fmt"
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/str"
	"gsadmin/global/e"
	"strings"
)

type SysAuth struct {
	baseservice.Service
}

func (s *SysAuth) Insert(req model.SysAuth) (err error) {
	req.Icon = fmt.Sprintf("layui-icon %s", req.Icon) //格式化icon
	req.Status = 1                                    //默认启用
	err = db.Instance().Create(&req).Error
	if err != nil {
		log.Instance().Error("SysAuthService.Insert:" + err.Error())
		return err
	}
	return nil
}

func (s *SysAuth) Update(req model.SysAuth) (err error) {
	if strings.HasPrefix(req.Icon, "layui-icon ") == false {
		req.Icon = fmt.Sprintf("layui-icon %s", req.Icon)
	}

	err = db.Instance().Model(model.SysAuth{}).Where("id = ?", req.ID).Updates(&req).Error
	if err != nil {
		log.Instance().Error("SysAuthService.Update:" + err.Error())
		return err
	}
	return nil
}

func (s *SysAuth) Delete(req *dto.SysAuthIDForm) error {
	//查询是否有子节点
	var childCount int
	err := db.Instance().Model(model.SysAuth{}).Where("status = 1 AND pid = ?", req.ID).Count(&childCount).Error
	if err != nil {
		log.Instance().Error("SysAuthService.Delete:" + err.Error())
		return err
	}

	if childCount > 0 {
		return errors.New("请先删除子节点")
	}
	//查询是否系统默认菜单
	var auth model.SysAuth
	err = db.Instance().Model(model.SysAuth{}).Where("id = ?", req.ID).First(&auth).Error
	if err != nil {
		log.Instance().Error("SysAuthService.Delete:" + err.Error())
		return err
	}

	if auth.IsDefault == 1 {
		return errors.New("系统基础节点不能删除")
	}
	//删除角色菜单
	roleAuth := SysRoleAuth{}
	err = roleAuth.Delete("auth_id = ?", req.ID)
	if err != nil {
		log.Instance().Error("SysAuthService.Delete:" + err.Error())
		return err
	}

	//删除菜单
	err = db.Instance().Where("id = ?", req.ID).Delete(model.SysAuth{}).Error
	if err != nil {
		log.Instance().Error("SysAuthService.Delete:" + err.Error())
		return err
	}

	return nil
}

func (s *SysAuth) GetByID(req *dto.SysAuthIDForm) (*dto.NodeResp, error) {
	var auth model.SysAuth
	err := db.Instance().Scopes(
		s.SetCondition(*req),
	).Model(model.SysAuth{}).First(&auth).Error

	if err != nil {
		log.Instance().Error("SysAuthService.GetByID:" + err.Error())
		return nil, err
	}

	return &dto.NodeResp{
		ID:        assertion.AnyToInt(auth.ID),
		Pid:       auth.Pid,
		AuthName:  auth.AuthName,
		AuthUrl:   auth.AuthUrl,
		Sort:      auth.Sort,
		IsShow:    auth.IsShow,
		Icon:      auth.Icon,
		PowerType: auth.PowerType}, nil
}

func (s *SysAuth) ListAuth(req *dto.SysAuthListForm) (list []model.SysAuth, count int, err error) {
	err = db.Instance().Scopes(
		s.SetPaginate(e.DefaultLimit, e.DefaultPage),
		s.SetCondition(*req),
	).
		Model(model.SysAuth{}).Where("status = ?", 1).
		Order("power_type,sort,pid").Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		log.Instance().Error("SysAuthService.ListAuth:" + err.Error())
		return nil, 0, err
	}
	return list, count, nil
}

func (s *SysAuth) ListAuthByIDs(ids []string) (list []model.SysAuth, count int, err error) {
	err = db.Instance().Scopes(
		s.SetPaginate(e.DefaultLimit, e.DefaultPage),
	).
		Model(model.SysAuth{}).Where("status = ? and id in (?)", 1, ids).
		Order("power_type,sort,pid").Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		log.Instance().Error("SysAuthService.ListAuthByIDs:" + err.Error())
		return nil, 0, err
	}
	return list, count, nil
}

func (s *SysAuth) FindAuths() ([]dto.FrontAuthResp, int) {
	auths, count, _ := s.ListAuth(&dto.SysAuthListForm{})
	var resp []dto.FrontAuthResp
	for _, v := range auths {
		var powerCode string
		if v.IsShow == 0 && v.Pid != 0 {
			powerCode = v.AuthUrl
		}
		resp = append(resp, dto.FrontAuthResp{
			PowerID:   assertion.AnyToString(v.ID),
			PowerName: v.AuthName,
			PowerType: assertion.AnyToString(v.PowerType),
			PowerCode: powerCode,
			PowerURL:  v.AuthUrl,
			OpenType:  "",
			ParentID:  assertion.AnyToString(v.Pid),
			Icon:      v.Icon,
			Sort:      v.Sort,
			CheckArr:  "0",
		})
	}
	return resp, count
}

func (s *SysAuth) GetByType(powerType int) []dto.AuthResp {
	var authNames []dto.AuthResp

	pType := assertion.AnyToString(powerType)
	auths, _, _ := s.ListAuth(&dto.SysAuthListForm{PowerType: pType})
	if len(auths) < 1 {
		return authNames
	}
	for _, v := range auths {
		var resp dto.AuthResp
		resp.ID = int(v.ID)
		resp.Pid = v.Pid
		resp.Name = v.AuthName
		resp.PowerType = v.PowerType
		authNames = append(authNames, resp)
	}
	return authNames
}

func (s *SysAuth) FindAuthPower(roleId int) (power dto.RolePower) {
	power.Status.Code = 200
	power.Status.Message = "success"

	auths, _, _ := s.ListAuth(&dto.SysAuthListForm{})
	var (
		powerData []dto.RolePowerData
		roleAuth  []string
	)

	roleAuthService := SysRoleAuth{}
	res, err := roleAuthService.ListByRoleId(roleId)
	if err != nil {
		log.Instance().Error("SysAuthService.FindAuthPower:" + err.Error())
		return power
	}
	for _, ra := range res {
		roleAuth = append(roleAuth, assertion.AnyToString(ra.AuthID))
	}
	for _, auth := range auths {
		checkArr := "0"
		if str.IsContain(roleAuth, assertion.AnyToString(auth.ID)) {
			checkArr = "1"
		}
		openType := ""
		if auth.PowerType == 1 {
			openType = "_iframe"
		}
		powerData = append(powerData, dto.RolePowerData{
			CheckArr:   checkArr,
			Enable:     1,
			Icon:       auth.Icon,
			OpenType:   openType,
			ParentID:   assertion.AnyToString(auth.Pid),
			PowerID:    assertion.AnyToString(auth.ID),
			PowerName:  auth.AuthName,
			PowerType:  assertion.AnyToString(auth.PowerType),
			PowerURL:   auth.AuthUrl,
			Sort:       auth.Sort,
			UpdateTime: auth.UpdatedAt,
		})
	}
	power.Data = powerData
	return power
}

func (s *SysAuth) FindAllPower() (power dto.RolePower) {
	power.Status.Code = 200
	power.Status.Message = "success"

	auths, _, _ := s.ListAuth(&dto.SysAuthListForm{})
	var powerData []dto.RolePowerData
	powerData = append(powerData, dto.RolePowerData{
		ParentID:  "-1",
		PowerID:   "0",
		PowerName: "顶级权限",
	})
	for _, auth := range auths {
		openType := ""
		if auth.PowerType == 1 {
			openType = "_iframe"
		}
		powerData = append(powerData, dto.RolePowerData{
			Enable:     1,
			Icon:       auth.Icon,
			OpenType:   openType,
			ParentID:   assertion.AnyToString(auth.Pid),
			PowerID:    assertion.AnyToString(auth.ID),
			PowerName:  auth.AuthName,
			PowerType:  assertion.AnyToString(auth.PowerType),
			PowerURL:   auth.AuthUrl,
			Sort:       auth.Sort,
			CreateTime: auth.CreatedAt,
			UpdateTime: auth.UpdatedAt,
		})
	}
	power.Data = powerData
	return power
}
