package router

import (
	"github.com/gin-gonic/gin"
	"gsadmin/middleware"
)

func monitorRouter(r *gin.Engine) {
	tr := r.Group("monitor")
	tr.Use(middleware.Auth())

	tr.GET("server", monitor.Server)
}
