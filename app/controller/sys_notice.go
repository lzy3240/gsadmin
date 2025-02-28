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
	a.MountCtx(c).SuccessResp().WriteHtmlExit("notice_list.html", gin.H{})
}

func (a SysNotice) NoticeStatus(c *gin.Context) {
	user := a.MountCtx(c).GetUserFromSession()
	req := dto.SysNoticeIDForm{}
	err := a.Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	svice := service.SysNotice{}
	err = svice.UpdateStatus(&req, user.ID)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetMsg("更新成功").SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
}

func (a SysNotice) NoticeBatchStatus(c *gin.Context) {
	user := a.MountCtx(c).GetUserFromSession()
	req := dto.SysNoticeIDsForm{}
	err := a.Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	svice := service.SysNotice{}
	err = svice.UpdateBatchStatus(&req, user.ID)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetMsg("更新成功").SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
}

func (a SysNotice) NoticeJson(c *gin.Context) {
	user := a.MountCtx(c).GetUserFromSession()
	req := dto.SysNoticeListForm{}
	err := a.Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}
	svice := service.SysNotice{}
	list, count, err := svice.List(&req, user.ID)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}
	a.SuccessResp().SetCode(0).SetCount(count).SetData(list).WriteJsonExit()
}

func (a SysNotice) NoticeUnReadJson(c *gin.Context) {
	user := a.MountCtx(c).GetUserFromSession()
	svice := service.SysNotice{}
	list, count, err := svice.ListUnRead(user.ID)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}
	a.SuccessResp().SetCode(0).SetCount(count).SetData(list).WriteJsonExit()
}

func (a SysNotice) NoticeEditPage(c *gin.Context) {
	user := a.MountCtx(c).GetUserFromSession()
	req := dto.SysNoticeIDForm{}
	err := a.Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	svice := service.SysNotice{}
	data, err := svice.Get(&req, user.ID)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.NoticeEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperEdit, e.NoticeEdit).WriteHtmlExit("notice_edit.html", gin.H{"notice": data})
}
