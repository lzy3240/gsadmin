package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gsadmin/app/model"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/cache"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/session"
	"gsadmin/core/utils/str"
	"gsadmin/global/e"
	"net/http"
	"strings"
)

// Auth 权限中间件
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		if isSignedIn(c) == false {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			if authByPage(c) == false {
				c.Redirect(http.StatusFound, "/not_found")
				c.Abort()
			} else {
				if c.Request.URL.Path == "/system" {
					c.Redirect(http.StatusFound, "/system/")
				}
				c.Next()
			}
		}
	}
}

func CheckLoginPage(c *gin.Context) {
	if isSignedIn(c) == true {
		c.Redirect(http.StatusFound, "/system/index")
	}
}

func CheckDefaultPage(c *gin.Context) {
	if isSignedIn(c) == false {
		c.Redirect(http.StatusFound, "/login")
	} else {
		c.Redirect(http.StatusFound, "/system/index")
	}
}

func isSignedIn(c *gin.Context) bool {
	uid := session.Get(c, e.SysAuth)
	return uid != nil
}

func authByPage(c *gin.Context) bool {
	userSvice := service.SysUser{}
	//获取用户信息
	tmp := session.Get(c, e.UserInfo)

	//类型转换
	user := new(model.SysUser)
	if s, ok := tmp.(string); ok {
		_ = json.Unmarshal([]byte(s), &user)
	}

	if userSvice.IsAdmin(user) { // 管理员不限制
		return true
	}
	url := c.Request.URL.Path
	if strings.HasSuffix(url, "/") {
		url = strings.TrimRight(url, "/")
	}

	allowAuthArr := strings.Split(e.AllowAuth, ",") // 校验公共路径
	if str.IsContain(allowAuthArr, url) {
		return true
	}

	var allowUrlArr []string
	menuCache, err := cache.Instance().Get(e.UserMenu, e.MenuCache+assertion.AnyToString(user.ID))
	if err == nil { //从缓存取菜单
		var mc dto.CacheMenuV2
		_ = json.Unmarshal([]byte(menuCache), &mc)
		allowUrlArr = strings.Split(mc.AllowUrl, ",")
	} else {
		result := userSvice.GetAuth(user)
		for _, v := range result {
			if str.IsContain([]string{"", "/"}, v.AuthUrl) == false {
				allowUrlArr = append(allowUrlArr, v.AuthUrl)
			}
		}
	}
	if str.IsContain(allowUrlArr, url) == false { // 校验用户路径
		return false
	}
	return true
}
