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
	a.Success(c, "操作成功").WriteHtmlExit("log_list.html", gin.H{})
}

func (a SysLog) LogLogin(c *gin.Context) {
	req := dto.SysLoginLogForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
		return
	}
	svice := service.SysLog{}

	list, count, err := svice.LoginLogList(&req)
	if err != nil {
		a.Error(c, "查询失败", err).WriteJsonExit()
		return
	}

	a.Custom(c, 0, "查询成功").SetPageData(count, list).WriteJsonExit()
}

func (a SysLog) LogOperate(c *gin.Context) {
	req := dto.SysOperLogForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
		return
	}
	svice := service.SysLog{}
	list, count, err := svice.OperLogList(&req)
	if err != nil {
		a.Error(c, "查询失败", err).WriteJsonExit()
		return
	}

	a.Custom(c, 0, "查询成功").SetPageData(count, list).WriteJsonExit()
}
