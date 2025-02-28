package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
)

type SysLog struct {
	baseapi.Api
}

func (a SysLog) LogListPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("log_list.html", gin.H{})
}

func (a SysLog) LogLogin(c *gin.Context) {
	req := dto.SysLoginLogForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}
	svice := service.SysLog{}

	list, count, err := svice.LoginLogList(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}

	a.SuccessResp().SetCode(0).SetCount(count).SetData(list).WriteJsonExit()
}

func (a SysLog) LogOperate(c *gin.Context) {
	req := dto.SysOperLogForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}
	svice := service.SysLog{}
	list, count, err := svice.OperLogList(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}

	a.SuccessResp().SetCode(0).SetCount(count).SetData(list).WriteJsonExit()
}
