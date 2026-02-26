package router

import (
	"github.com/gin-gonic/gin"
	"gsadmin/app/controller"
)

var (
	base    = controller.SysBase{}
	auth    = controller.SysAuth{}
	conf    = controller.SysConf{}
	dict    = controller.SysDict{}
	slog    = controller.SysLog{}
	notice  = controller.SysNotice{}
	role    = controller.SysRole{}
	user    = controller.SysUser{}
	monitor = controller.SysMonitor{}
)

// RegisterSystemRouter 注册路由
func RegisterSystemRouter(r *gin.Engine) {
	commonRouter(r)
	systemRouter(r)
	monitorRouter(r)
}
