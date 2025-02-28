package router

import (
	"github.com/gin-gonic/gin"
	"gsadmin/middleware"
)

func MonitorRouter(r *gin.Engine) {
	tr := r.Group("monitor")
	tr.Use(middleware.Auth())

	tr.GET("server", monitor.Server)
}
