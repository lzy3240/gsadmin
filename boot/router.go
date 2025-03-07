package boot

import (
	"github.com/gin-gonic/gin"
	"gsadmin/app/router"
	"gsadmin/core/baserouter"
	"gsadmin/core/config"
	"gsadmin/core/utils/session"
	"gsadmin/middleware"
)

func initRouter() *gin.Engine { // staticFs, templateFs embed.FS
	gin.SetMode(config.Instance().App.RunMode)
	r := gin.New()

	//模版文件加载, 两种模式选择一种: 方式一: 编译到程序; 方式二: 静态路径
	//t, _ := template.ParseFS(templateFs, "template/**/**/*.html")
	//r.SetHTMLTemplate(t)
	r.LoadHTMLGlob("template/**/**/*.html")
	//fads, _ := fs.Sub(staticFs, "static")
	//r.StaticFS("/static", http.FS(fads))
	r.Static("/static", "./static")

	//其他文件
	r.Static(config.Instance().App.FileViewPath, config.Instance().App.FileSavePath) //上传文件
	r.StaticFile("/favicon.ico", "./static/favicon.ico")                             //单独文件
	r.StaticFile("/robots.txt", "./static/robots.txt")                               //单独文件

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(session.EnableCookieSession(config.Instance().App.JwtSecret))

	//注册系统路由
	router.RegisterSystemRouter()

	//注册业务路由
	//TODO

	//初始化路由
	baserouter.InitializeRouter(r)

	return r
}
