package boot

import (
	"github.com/gin-gonic/gin"
	"gsadmin/app/router"
	"gsadmin/core/config"
	"gsadmin/core/utils/session"
	"gsadmin/global/f"
	"gsadmin/middleware"
	"html/template"
)

func initRouter() *gin.Engine { // staticFs, templateFs embed.FS
	gin.SetMode(config.Instance().App.RunMode)
	r := gin.New()

	//注册方法(html渲染使用)
	r.SetFuncMap(template.FuncMap{
		"DateFormat": f.DateFormat,
	})

	//模版文件加载, 两种模式选择一种:
	//方式一: 编译嵌入,修改静态文件后重新打包
	//tp, _ := template.ParseFS(templateFs, "template/**/**/*.html")
	//r.SetHTMLTemplate(tp)
	//fads, _ := fs.Sub(staticFs, "static")
	//r.StaticFS("/static", http.FS(fads))

	//方式二: 静态路径,修改静态文件后立即生效
	r.LoadHTMLGlob("template/**/**/*.html")
	r.Static("/static", "./static")

	//其他文件
	r.Static(config.Instance().App.FileViewPath, config.Instance().App.FileSavePath) //上传文件
	r.StaticFile("/favicon.ico", "favicon.ico")                                      //单独文件
	r.StaticFile("/robots.txt", "robots.txt")                                        //单独文件

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(session.EnableCookieSession(config.Instance().App.JwtSecret))

	//注册系统路由
	router.RegisterSystemRouter(r)

	//注册业务路由
	//TODO

	return r
}
