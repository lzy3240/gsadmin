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

// 注册路由
func RegisterSystemRouter(r *gin.Engine) {
	CommonRouter(r)
	SystemRouter(r)
	MonitorRouter(r)
}
