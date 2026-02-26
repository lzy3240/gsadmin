package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/model"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/core/cache"
	"gsadmin/global/e"
	"gsadmin/utils/assertion"
)

type SysAuth struct {
	baseapi.Api
}

func (a SysAuth) AuthListPage(c *gin.Context) {
	a.Success(c, "操作成功").WriteHtmlExit("auth_list.html", gin.H{})
}

func (a SysAuth) AuthEditPage(c *gin.Context) {
	authService := service.SysAuth{}
	var req dto.SysAuthIDForm
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数错误", err).SetLogTag(e.OperEdit, e.AuthNodeEdit).WriteStringExit("%s", err.Error())
		return
	}

	res, err := authService.GetByID(&req)
	if err != nil {
		a.Error(c, "查询失败", err).SetLogTag(e.OperEdit, e.AuthNodeEdit).WriteStringExit("%s", err.Error())
		return
	}
	firstAuths := authService.GetByType(0)
	secondAuths := authService.GetByType(1)
	a.Success(c, "操作成功").SetLogTag(e.OperEdit, e.AuthNodeEdit).WriteHtmlExit("auth_edit.html", gin.H{"parents": firstAuths, "seconds": secondAuths, "auth": res})
}

// AuthEdit 修改接口/新增接口
func (a SysAuth) AuthEdit(c *gin.Context) {
	authService := service.SysAuth{}
	user := a.GetUserFromSession(c)

	//新增/修改接口
	var req model.SysAuth
	//var req dto.SysAuthNodeForm
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数错误", err).SetLogTag(e.OperOther, e.AuthNodeEdit).WriteJsonExit()
		return
	}

	if req.ID == 0 { //兼容add
		req.SetCreate(user.ID)
		if err = authService.Insert(req); err != nil {
			a.Error(c, "新增失败", err).SetLogTag(e.OperAdd, e.AuthNodeAdd).WriteJsonExit()
			return
		}
		_ = cache.Instance().Del(e.UserMenu, e.MenuCache+assertion.AnyToString(a.GetUidFromSession(c))) // 删除栏目列表缓存，重新进行设置service.GetUid(c)
		a.Success(c, "操作成功").SetLogTag(e.OperAdd, e.AuthNodeAdd).WriteJsonExit()
	} else { //执行edit
		req.SetUpdate(user.ID)
		if err = authService.Update(req); err != nil {
			a.Error(c, "修改失败", err).SetLogTag(e.OperOther, e.AuthNodeEdit).WriteJsonExit()
			return
		}
		_ = cache.Instance().Del(e.UserMenu, e.MenuCache+assertion.AnyToString(a.GetUidFromSession(c))) // 删除栏目列表缓存，重新进行设置service.GetUid(c)
		a.Success(c, "操作成功").SetLogTag(e.OperEdit, e.AuthNodeEdit).WriteJsonExit()
	}
}

func (a SysAuth) AuthDelete(c *gin.Context) {
	authService := service.SysAuth{}
	req := dto.SysAuthIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数错误", err).SetLogTag(e.OperDel, e.AuthDelete).WriteJsonExit()
		return
	}

	if err = authService.Delete(&req); err != nil {
		a.Error(c, "删除失败", err).SetLogTag(e.OperDel, e.AuthDelete).WriteJsonExit()
		return
	}
	_ = cache.Instance().Del(e.UserMenu, e.MenuCache+assertion.AnyToString(a.GetUidFromSession(c))) // 删除栏目列表缓存，重新进行设置service.GetUid(c)
	a.Success(c, "操作成功").SetLogTag(e.OperDel, e.AuthDelete).WriteJsonExit()
}

func (a SysAuth) AuthNode(c *gin.Context) {
	authService := service.SysAuth{}
	req := dto.SysAuthIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数错误", err).WriteJsonExit()
		return
	}

	data, err := authService.GetByID(&req)
	if err != nil {
		a.Error(c, "查询失败", err).WriteJsonExit()
		return
	}
	a.Success(c, "操作成功").SetData(data).WriteJsonExit()
}

func (a SysAuth) AuthNodes(c *gin.Context) {
	authService := service.SysAuth{}
	data, count := authService.FindAuths()
	a.Success(c, "操作成功").SetPageData(assertion.AnyToInt(count), data).WriteJsonExit()
}

func (a SysAuth) AddNodePage(c *gin.Context) {
	authService := service.SysAuth{}
	firstAuths := authService.GetByType(0)
	secondAuths := authService.GetByType(1)
	a.Success(c, "操作成功").WriteHtmlExit("auth_add.html", gin.H{"parents": firstAuths, "seconds": secondAuths})
}

func (a SysAuth) Parent(c *gin.Context) {
	authService := service.SysAuth{}
	data := authService.FindAllPower()
	a.Success(c, "操作成功").WriteCustomJsonExit(data)
}
