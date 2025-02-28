package baserouter

import "github.com/gin-gonic/gin"

var appRouter = make([]func(r *gin.Engine), 0)

// 注册路由
func RegisterRouter(f func(r *gin.Engine)) {
	appRouter = append(appRouter, f)
}

// 初始化路由
func InitializeRouter(r *gin.Engine) {
	for _, f := range appRouter {
		f(r)
	}
}
