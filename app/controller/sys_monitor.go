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
	a.MountCtx(c)
	svice := service.SysMonitor{}
	data, _ := svice.Server()

	a.SuccessResp().WriteHtmlExit("server_info.html", gin.H{"server_info": data})
}
