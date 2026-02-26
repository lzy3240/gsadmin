package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/global/e"
	"gsadmin/utils/assertion"
)

type SysRole struct {
	baseapi.Api
}

func (a SysRole) RoleListPage(c *gin.Context) {
	a.Success(c, "操作成功").WriteHtmlExit("role_list.html", gin.H{})
}

func (a SysRole) RoleJson(c *gin.Context) {
	roleSvice := service.SysRole{}
	req := dto.SysRoleListForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
	}

	list, count, err := roleSvice.List(&req)
	if err != nil {
		a.Error(c, "查询失败", err).WriteJsonExit()
		return
	}
	a.Custom(c, 0, "查询成功").SetPageData(count, list).WriteJsonExit()
}

func (a SysRole) RoleAddPage(c *gin.Context) {
	a.Success(c, "操作成功").SetLogTag(e.OperAdd, e.RoleAdd).WriteHtmlExit("role_add.html", gin.H{})
}

func (a SysRole) RoleAdd(c *gin.Context) {
	user := a.GetUserFromSession(c)

	req := dto.SysRoleAddForm{}
	if err := a.Bind(c, &req, binding.Form); err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperAdd, e.RoleAdd).WriteJsonExit()
		return
	}

	if req.Status == "" {
		req.Status = "0"
	} else if req.Status == "on" {
		req.Status = "1"
	}

	req.SetCreate(user.ID)
	roleSvice := service.SysRole{}
	if err := roleSvice.Insert(&req); err != nil {
		a.Error(c, "新增失败", err).SetLogTag(e.OperAdd, e.RoleAdd).WriteJsonExit()
		return
	}

	a.Success(c, "新增成功").SetLogTag(e.OperAdd, e.RoleAdd).WriteJsonExit()
}

func (a SysRole) RolePowerPage(c *gin.Context) {
	req := dto.SysRolePowerIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		//a.Success().WriteCustomJsonExit(nil)
		a.Error(c, "参数校验失败", err).WriteCustomJsonExit(nil)
		return
	}

	a.Success(c, "操作成功").WriteHtmlExit("role_power.html", gin.H{"id": req.RoleId})
}

func (a SysRole) GetRolePower(c *gin.Context) {
	authService := service.SysAuth{}
	req := dto.SysRolePowerIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		//a.Success().WriteCustomJsonExit(nil)
		a.Error(c, "参数校验失败", err).WriteCustomJsonExit(nil)
		return
	}

	data := authService.FindAuthPower(assertion.AnyToInt(req.RoleId))
	a.Success(c, "操作成功").WriteCustomJsonExit(data)
}

func (a SysRole) SaveRolePower(c *gin.Context) {
	req := dto.SysRolePowerForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperAdd, e.RoleSave).WriteJsonExit()
		return
	}

	roleSvice := service.SysRole{}

	err = roleSvice.InsertRoleAuth(&req)
	if err != nil {
		a.Error(c, "保存失败", err).SetLogTag(e.OperAdd, e.RoleSave).WriteJsonExit()
		return
	}
	a.Success(c, "保存成功").SetLogTag(e.OperAdd, e.RoleSave).WriteJsonExit()
}

func (a SysRole) RoleEditPage(c *gin.Context) {
	roleSvice := service.SysRole{}
	req := dto.SysRoleIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.RoleEdit).WriteStringExit("%s", err.Error())
		return
	}

	role, err := roleSvice.Get(&req)
	if err != nil {
		a.Error(c, "查询失败", err).WriteStringExit("%s", err.Error())
		return
	}
	a.Success(c, "查询成功").SetLogTag(e.OperEdit, e.RoleEdit).WriteHtmlExit("role_edit.html", gin.H{"role": role})
}

func (a SysRole) RoleEdit(c *gin.Context) {
	user := a.GetUserFromSession(c)
	roleSvice := service.SysRole{}

	req := dto.SysRoleEditForm{}
	err := a.Bind(c, &req, binding.JSON)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
		return
	}

	req.SetUpdate(user.ID)

	if req.Status == "" {
		req.Status = "0"
	} else if req.Status == "on" {
		req.Status = "1"
	}

	err = roleSvice.Update(&req)
	if err != nil {
		a.Error(c, "更新失败", err).SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
		return
	}
	a.Success(c, "更新成功").SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
}

func (a SysRole) RoleDelete(c *gin.Context) {
	roleSvice := service.SysRole{}
	req := dto.SysAuthIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperDel, e.RoleDelete).WriteJsonExit()
		return
	}

	err = roleSvice.Delete(&req)
	if err != nil {
		a.Error(c, "删除失败", err).SetLogTag(e.OperDel, e.RoleDelete).WriteJsonExit()
		return
	}
	a.Success(c, "删除成功").SetLogTag(e.OperDel, e.RoleDelete).WriteJsonExit()
}

func (a SysRole) RoleStatus(c *gin.Context) {
	roleSvice := service.SysRole{}
	req := dto.SysRoleStatusForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
		return
	}

	err = roleSvice.UpdateRoleStatus(&req)
	if err != nil {
		a.Error(c, "更新失败", err).SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
		return
	}
	a.Success(c, "更新成功").SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
}
