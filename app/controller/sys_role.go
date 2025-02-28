package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/core/utils/assertion"
	"gsadmin/global/e"
)

type SysRole struct {
	baseapi.Api
}

func (a SysRole) RoleListPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("role_list.html", gin.H{})
}

func (a SysRole) RoleJson(c *gin.Context) {
	roleSvice := service.SysRole{}
	req := dto.SysRoleListForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
	}

	list, count, err := roleSvice.List(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}
	a.SuccessResp().SetCode(0).SetCount(count).SetData(list).WriteJsonExit()
}

func (a SysRole) RoleAddPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().SetLogTag(e.OperAdd, e.RoleAdd).WriteHtmlExit("role_add.html", gin.H{})
}

func (a SysRole) RoleAdd(c *gin.Context) {
	user := a.MountCtx(c).GetUserFromSession()

	req := dto.SysRoleAddForm{}
	if err := a.Bind(&req, binding.Form); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.RoleAdd).WriteJsonExit()
		return
	}

	req.SetCreate(user.ID)
	roleSvice := service.SysRole{}
	if err := roleSvice.Insert(&req); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.RoleAdd).WriteJsonExit()
		return
	}

	a.SuccessResp().SetLogTag(e.OperAdd, e.RoleAdd).WriteJsonExit()
}

func (a SysRole) RolePowerPage(c *gin.Context) {
	req := dto.SysRolePowerIDForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.SuccessResp().WriteCustomJsonExit(nil)
		return
	}

	a.SuccessResp().WriteHtmlExit("role_power.html", gin.H{"id": req.RoleId})
}

func (a SysRole) GetRolePower(c *gin.Context) {
	authService := service.SysAuth{}
	req := dto.SysRolePowerIDForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.SuccessResp().WriteCustomJsonExit(nil)
		return
	}

	data := authService.FindAuthPower(assertion.AnyToInt(req.RoleId))
	a.SuccessResp().WriteCustomJsonExit(data)
}

func (a SysRole) SaveRolePower(c *gin.Context) {
	req := dto.SysRolePowerForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.RoleSave).WriteJsonExit()
		return
	}

	roleSvice := service.SysRole{}

	err = roleSvice.InsertRoleAuth(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.RoleSave).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperAdd, e.RoleSave).WriteJsonExit()
}

func (a SysRole) RoleEditPage(c *gin.Context) {
	roleSvice := service.SysRole{}
	req := dto.SysRoleIDForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetLogTag(e.OperEdit, e.RoleEdit).WriteStringExit("%s", err.Error())
		return
	}

	role, err := roleSvice.Get(&req)
	if err != nil {
		a.ErrorResp().WriteStringExit("%s", err.Error())
		return
	}
	a.SuccessResp().SetLogTag(e.OperEdit, e.RoleEdit).WriteHtmlExit("role_edit.html", gin.H{"role": role})
}

func (a SysRole) RoleEdit(c *gin.Context) {
	user := a.MountCtx(c).GetUserFromSession()
	roleSvice := service.SysRole{}

	req := dto.SysRoleEditForm{}
	err := a.Bind(&req, binding.JSON)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
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
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
}

func (a SysRole) RoleDelete(c *gin.Context) {
	roleSvice := service.SysRole{}
	req := dto.SysAuthIDForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperDel, e.RoleDelete).WriteJsonExit()
		return
	}

	err = roleSvice.Delete(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperDel, e.RoleDelete).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperDel, e.RoleDelete).WriteJsonExit()
}

func (a SysRole) RoleStatus(c *gin.Context) {
	roleSvice := service.SysRole{}
	req := dto.SysRoleStatusForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
		return
	}

	err = roleSvice.UpdateRoleStatus(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperEdit, e.RoleEdit).WriteJsonExit()
}
