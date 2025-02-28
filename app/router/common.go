package router

import (
	"github.com/gin-gonic/gin"
	"gsadmin/middleware"
)

func CommonRouter(r *gin.Engine) {
	r.GET("/", middleware.CheckDefaultPage)
	r.GET("login", middleware.CheckLoginPage, base.LoginPage)
	r.POST("login", base.LoginHandler)
	r.GET("logout", base.Logout)
	r.GET("not_found", base.NotFoundPage)
	r.GET("captcha", base.GetCaptcha)
	r.POST("captcha_verify", base.CaptchaVerify)
}
