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

	// 缓存分区, 暂不分区
	SysConf   = "sys_config"
	UserMenu  = "user_menu"
	UserToken = "user_token"

	// 缓存KEY
	Menu          = "menu_list"
	UserInfo      = "user_info"
	SysAuth       = "sys_auth"
	UserLoginErr  = "user_pwd_err_"
	UserLock      = "user_lock_"
	MaxErrNum     = 5
	MenuCache     = "menu_cache"
	AuthList      = "auth_list"
	BaseConfig    = "base_conf"
	SiteConfig    = "site_conf"
	ServerMonitor = "server_monitor"
	LongMonTime   = 1 //server info缓存时间, 分钟

	// conf
	BaseConfigType = "base_conf"
	SiteConfigType = "site_conf"

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

	DefaultUpload = "上传图片"

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
