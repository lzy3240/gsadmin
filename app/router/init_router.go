package router

import (
	"gsadmin/app/controller"
	"gsadmin/core/baserouter"
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
func RegisterSystemRouter() {
	baserouter.RegisterRouter(CommonRouter)
	baserouter.RegisterRouter(SystemRouter)
	baserouter.RegisterRouter(MonitorRouter)
}
