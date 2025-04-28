package e

const (
	DefaultAvatar  = "/static/admin/images/avatar.jpg"
	DefUploadSize  = 2 * 1024 * 1024 // 默认最大上传
	DefaultSaltLen = 10

	DefaultPage  = 1
	DefaultLimit = 5000

	AllowAuth = "/system,/system/index,/system/main,/system/base,/system/menu,/system/notice,/system/dict/data/json" // 不需要验证的地址放在这里

	// 日期格式化
	TimeFormatDay = "20060102"
	TimeFormat    = "2006-01-02 15:04:05"

	// 操作类型
	OperOther = 0 //0其它
	OperAdd   = 1 //1新增
	OperEdit  = 2 //2修改
	OperDel   = 3 //3删除

	// 响应编码
	SUCCESS      = 200 // 成功
	ERROR        = 500 //错误
	UNAUTHORIZED = 401 //鉴权失败
	FORBIDDEN    = 403 //无操作权限
	FAIL         = -1  //失败

	// 消息主题
	TopicOperLog = "oper_log"

	// Session key
	UserInfo = "user_info"
	SysAuth  = "sys_auth"

	// 缓存分区
	SysBase   = "sys_base"   //系统基础缓存区
	SysDict   = "sys_dict"   //系统字典缓存区
	UserMenu  = "user_menu"  //用户菜单缓存区
	UserToken = "user_token" //用户token缓存区, 目前使用session, 不启用缓存
	SysLogin  = "sys_login"  //用户登录控制缓存区

	// 缓存key
	UserPwdErr    = "user_pwd_err_"
	UserLock      = "user_lock_"
	MenuCache     = "menu_cache_"
	AuthCache     = "auth_cache_"
	BaseConfig    = "base_conf"
	SiteConfig    = "site_conf"
	ServerMonitor = "server_monitor"
	CmsPvDate     = "cms_pv_date_" //cms pv date key

	// cache time etc
	MonCacheTime  = 60    // server monitor cache time
	MenuCacheTime = 3600  // user menu cache time
	UserErrTimes  = 5     // password err times
	UserLockTime  = 300   // user lock times
	ConfCacheTime = 86400 // config cache time

	// log title
	UserView   = "查看用户"
	UserAdd    = "新增用户"
	UserEdit   = "修改用户"
	UserLogin  = "用户登陆"
	UserDelete = "删除用户"

	AuthView     = "查看权限"
	AuthDelete   = "删除节点"
	AuthNodeEdit = "修改节点"
	AuthNodeAdd  = "新增节点"
	AuthNodeSave = "权限配置"

	RoleView   = "查看角色"
	RoleEdit   = "修改角色"
	RoleAdd    = "新增角色"
	RoleDelete = "删除角色"
	RoleSave   = "权限分配"

	DefaultUpload = "上传文件"

	SiteEdit = "更新站点设置"
	BaseEdit = "更新基础配置"

	AvatarEdit     = "更新头像"
	ProfileEdit    = "更新资料"
	PwdEditHandler = "修改密码"

	DictView       = "查看字典"
	DictTypeAdd    = "新增字典"
	DictTypeEdit   = "修改字典"
	DictTypeDelete = "删除字典"
	DictDataAdd    = "新增字典值"
	DictDataEdit   = "修改字典值"
	DictDataDelete = "删除字典值"

	NoticeView = "查看消息"
	NoticeEdit = "修改消息"

	// capt
	//ImgHeight    = 80
	//ImgWidth     = 240
	//ImgKeyLength = 4

	// task
	//Local    = 1
	//Remote   = 2
	//MaxPool  = 10
	//ToLocal  = 1
	//ToRemote = 2
)
