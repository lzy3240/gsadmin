package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/global/e"
)

type SysNotice struct {
	baseapi.Api
}

func (a SysNotice) NoticeListPage(c *gin.Context) {
	a.Success(c, "操作成功").WriteHtmlExit("notice_list.html", gin.H{})
}

func (a SysNotice) NoticeStatus(c *gin.Context) {
	user := a.GetUserFromSession(c)
	req := dto.SysNoticeIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	svice := service.SysNotice{}
	err = svice.UpdateStatus(&req, user.ID)
	if err != nil {
		a.Error(c, "更新失败", err).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	a.Success(c, "更新成功").SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
}

func (a SysNotice) NoticeBatchStatus(c *gin.Context) {
	user := a.GetUserFromSession(c)
	req := dto.SysNoticeIDsForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	svice := service.SysNotice{}
	err = svice.UpdateBatchStatus(&req, user.ID)
	if err != nil {
		a.Error(c, "更新失败", err).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	a.Success(c, "更新成功").SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
}

func (a SysNotice) NoticeJson(c *gin.Context) {
	user := a.GetUserFromSession(c)
	req := dto.SysNoticeListForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
		return
	}
	svice := service.SysNotice{}
	list, count, err := svice.List(&req, user.ID)
	if err != nil {
		a.Error(c, "查询失败", err).WriteJsonExit()
		return
	}
	a.Custom(c, 0, "查询成功").SetPageData(count, list).WriteJsonExit()
}

func (a SysNotice) NoticeUnReadJson(c *gin.Context) {
	user := a.GetUserFromSession(c)
	svice := service.SysNotice{}
	list, count, err := svice.ListUnRead(user.ID)
	if err != nil {
		a.Error(c, "查询失败", err).WriteJsonExit()
		return
	}
	a.Custom(c, 0, "查询成功").SetPageData(count, list).WriteJsonExit()
}

func (a SysNotice) NoticeEditPage(c *gin.Context) {
	user := a.GetUserFromSession(c)
	req := dto.SysNoticeIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	svice := service.SysNotice{}
	data, err := svice.Get(&req, user.ID)
	if err != nil {
		a.Error(c, "查询失败", err).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	a.Success(c, "查询成功").SetLogTag(e.OperEdit, e.NoticeEdit).WriteHtmlExit("notice_edit.html", gin.H{"notice": data})
}
