package controller

import (
	"github.com/gin-gonic/gin"
	"gsadmin/app/service"
	"gsadmin/core/baseapi"
)

type SysMonitor struct {
	baseapi.Api
}

func (a SysMonitor) Server(c *gin.Context) {
	svice := service.SysMonitor{}
	data, _ := svice.GetServerInfo()

	a.Success(c, "操作成功").WriteHtmlExit("server_info.html", gin.H{"server_info": data})
}
