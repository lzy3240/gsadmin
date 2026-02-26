package router

import (
	"github.com/gin-gonic/gin"
	"gsadmin/middleware"
)

func systemRouter(r *gin.Engine) {
	sr := r.Group("system")

	// default
	sr.GET("base", base.Base)
	sr.GET("menu", base.Menu)
	sr.GET("server_err", base.ServerErrPage) //错误页面
	sr.GET("ui/icon", base.IconShow)
	sr.GET("file", base.ShowFile)

	sr.Use(middleware.Auth()) //下面的路由验证中间件

	// index
	sr.GET("/", base.IndexPage)     //Index页面
	sr.GET("index", base.IndexPage) //Index页面
	sr.GET("main", base.MainPage)   //Main页面

	// upload
	sr.POST("upload", base.Upload)
	sr.GET("uploadFile", base.UploadFile)

	// log
	sr.GET("log/list", slog.LogListPage)
	sr.GET("log/operate", slog.LogOperate)
	sr.GET("log/login", slog.LogLogin)

	sr.Use(middleware.LogTo()) //下面的路由记录日志

	// notice
	sr.GET("notice", notice.NoticeUnReadJson)
	sr.GET("notice/list", notice.NoticeListPage)
	sr.GET("notice/json", notice.NoticeJson)
	sr.GET("notice/edit", notice.NoticeEditPage)
	sr.POST("notice/status", notice.NoticeStatus)
	sr.POST("notice/batchRead", notice.NoticeBatchStatus)

	// user 用户列表页
	sr.GET("user/list", user.UserListPage)
	sr.GET("user/json", user.UserJson)
	sr.GET("user/add", user.UserAddPage)
	sr.POST("user/add", user.UserAdd)
	sr.GET("user/edit", user.UserEditPage)
	sr.POST("user/edit", user.UserEdit)
	sr.POST("user/status", user.UserStatus)
	sr.POST("user/delete", user.UserDelete)
	sr.POST("user/batchRemove", user.UserBatchDelete)

	// user 个人信息编辑页
	sr.POST("user/avatar", user.AvatarEdit)
	sr.GET("user/uploadPage", user.UploadPage)
	sr.GET("user/profile", user.ProfilePage)
	sr.POST("user/profile", user.ProfileEdit)
	sr.GET("user/pwd", user.PwdEditPage)
	sr.POST("user/pwd", user.PwdEdit)

	// role 角色管理列表页
	sr.GET("role/list", role.RoleListPage)
	sr.GET("role/json", role.RoleJson)
	sr.GET("role/add", role.RoleAddPage)
	sr.POST("role/add", role.RoleAdd)
	sr.GET("role/power", role.RolePowerPage)
	sr.GET("role/getRolePower", role.GetRolePower)
	sr.POST("role/saveRolePower", role.SaveRolePower)
	sr.GET("role/edit", role.RoleEditPage)
	sr.POST("role/edit", role.RoleEdit)
	sr.POST("role/delete", role.RoleDelete)
	sr.POST("role/status", role.RoleStatus)

	// auth 功能权限列表页
	sr.GET("auth/list", auth.AuthListPage)  // 权限列表页面
	sr.GET("auth/nodes", auth.AuthNodes)    // 权限配置
	sr.GET("auth/add", auth.AddNodePage)    // 新增权限页面
	sr.GET("auth/edit", auth.AuthEditPage)  // 修改权限页面
	sr.POST("auth/edit", auth.AuthSave)     // 保存权限, 兼容新增、修改
	sr.POST("auth/node", auth.AuthNode)     // 权限配置列表页
	sr.POST("auth/delete", auth.AuthDelete) // 权限配置列表页
	sr.GET("auth/parent", auth.Parent)      // 权限列表

	// site 站点设置
	sr.GET("site/edit", conf.SiteEditPage) // 站点配置页面
	sr.POST("site/edit", conf.SiteEdit)

	// base 基础设置
	//sr.GET("base/edit", conf.BaseEditPage) // 基础配置页面
	//sr.POST("base/edit", conf.BaseEdit)

	// dict type字典设置
	sr.GET("dict/type/list", dict.DictTypeListPage)
	sr.GET("dict/type/json", dict.DictTypeJson)
	sr.POST("dict/type/status", dict.DictTypeStatus)
	sr.GET("dict/type/add", dict.DictTypeAddPage)
	//sr.POST("dict/type/add", dict.DictTypeSave)
	sr.GET("dict/type/edit", dict.DictTypeEditPage)
	sr.POST("dict/type/edit", dict.DictTypeSave) // 保存字典类型, 兼容新增、修改
	sr.POST("dict/type/delete", dict.DictTypeDelete)

	// dict data字典设置
	sr.GET("dict/data/list", dict.DictDataListPage)
	sr.GET("dict/data/json", dict.DictDataJson)
	sr.POST("dict/data/status", dict.DictDataStatus)
	sr.GET("dict/data/add", dict.DictDataAddPage)
	//sr.POST("dict/data/add", dict.DictDataAdd)
	sr.GET("dict/data/edit", dict.DictDataEditPage)
	sr.POST("dict/data/edit", dict.DictDataSave) // 保存字典数据, 兼容新增、修改
	sr.POST("dict/data/delete", dict.DictDataDelete)
}
