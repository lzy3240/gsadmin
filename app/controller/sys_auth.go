package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/model"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/core/cache"
	"gsadmin/core/utils/assertion"
	"gsadmin/global/e"
)

type SysAuth struct {
	baseapi.Api
}

func (a SysAuth) AuthListPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("auth_list.html", gin.H{})
}

func (a SysAuth) AuthEditPage(c *gin.Context) {
	authService := service.SysAuth{}
	var req dto.SysAuthIDForm
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetLogTag(e.OperEdit, e.AuthNodeEdit).WriteStringExit("%s", a.TransErr(err))
		return
	}

	res, err := authService.GetByID(&req)
	if err != nil {
		a.ErrorResp().SetLogTag(e.OperEdit, e.AuthNodeEdit).WriteStringExit("%s", a.TransErr(err))
		return
	}
	firstAuths := authService.GetByType(0)
	secondAuths := authService.GetByType(1)
	a.SuccessResp().SetLogTag(e.OperEdit, e.AuthNodeEdit).WriteHtmlExit("auth_edit.html", gin.H{"parents": firstAuths, "seconds": secondAuths, "auth": res})
}

// AuthEdit 修改接口/新增接口
func (a SysAuth) AuthEdit(c *gin.Context) {
	authService := service.SysAuth{}
	user := a.MountCtx(c).GetUserFromSession()

	//新增/修改接口
	var req model.SysAuth
	//var req dto.SysAuthNodeForm
	err := a.Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperOther, e.AuthNodeEdit).WriteJsonExit()
		return
	}

	if req.ID == 0 { //兼容add
		req.SetCreate(user.ID)
		if err = authService.Insert(req); err != nil {
			a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.AuthNodeAdd).WriteJsonExit()
			return
		}
		cache.Instance().Delete(e.MenuCache + assertion.AnyToString(a.GetUidFromSession())) // 删除栏目列表缓存，重新进行设置service.GetUid(c)
		a.SuccessResp().SetLogTag(e.OperAdd, e.AuthNodeAdd).WriteJsonExit()
	} else { //执行edit
		req.SetUpdate(user.ID)
		if err = authService.Update(req); err != nil {
			a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperOther, e.AuthNodeEdit).WriteJsonExit()
			return
		}
		cache.Instance().Delete(e.MenuCache + assertion.AnyToString(a.GetUidFromSession())) // 删除栏目列表缓存，重新进行设置service.GetUid(c)
		a.SuccessResp().SetLogTag(e.OperEdit, e.AuthNodeEdit).WriteJsonExit()
	}
}

func (a SysAuth) AuthDelete(c *gin.Context) {
	authService := service.SysAuth{}
	req := dto.SysAuthIDForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperDel, e.AuthDelete).WriteJsonExit()
		return
	}

	if err = authService.Delete(&req); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperDel, e.AuthDelete).WriteJsonExit()
		return
	}
	cache.Instance().Delete(e.MenuCache + assertion.AnyToString(a.GetUidFromSession())) // 删除栏目列表缓存，重新进行设置service.GetUid(c)
	a.SuccessResp().SetLogTag(e.OperDel, e.AuthDelete).WriteJsonExit()
}

func (a SysAuth) AuthNode(c *gin.Context) {
	authService := service.SysAuth{}
	req := dto.SysAuthIDForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}

	data, err := authService.GetByID(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}
	a.SuccessResp().SetData(data).WriteJsonExit()
}

func (a SysAuth) AuthNodes(c *gin.Context) {
	authService := service.SysAuth{}
	data, count := authService.FindAuths()
	a.MountCtx(c).SuccessResp().SetCount(assertion.AnyToInt(count)).SetData(data).WriteJsonExit()
}

func (a SysAuth) AddNodePage(c *gin.Context) {
	authService := service.SysAuth{}
	firstAuths := authService.GetByType(0)
	secondAuths := authService.GetByType(1)
	a.MountCtx(c).SuccessResp().WriteHtmlExit("auth_add.html", gin.H{"parents": firstAuths, "seconds": secondAuths})
}

func (a SysAuth) Parent(c *gin.Context) {
	authService := service.SysAuth{}
	data := authService.FindAllPower()
	a.MountCtx(c).SuccessResp().WriteCustomJsonExit(data)
}
