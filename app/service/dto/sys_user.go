package dto

type SysUserShow struct {
	ID        uint   `json:"id"`
	LoginName string `json:"login_name"`
	RealName  string `json:"real_name"`
	RoleIds   string `json:"role_ids"`
	Level     int    `json:"level"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Remark    string `json:"remark"`
	Status    int    `json:"status"`
	LastLogin string `json:"last_login"`
	LastIp    string `json:"last_ip"`
}

type SysUserIDForm struct {
	ID string `json:"id" form:"id" search:"type:exact;column:id;table:sys_user"`
}

type SysUserIDsForm struct {
	IDs string `json:"ids" form:"ids" search:"type:exact;column:id;table:sys_user"`
}

type SysUserListForm struct {
	PageForm  `search:"-"`
	ID        string `json:"id" form:"id" search:"type:exact;column:id;table:sys_user"`
	LoginName string `json:"login_name" form:"login_name" search:"type:contains;column:login_name;table:sys_user"`
	RealName  string `json:"real_name" form:"real_name" search:"type:contains;column:real_name;table:sys_user"`
	Phone     string `json:"phone" form:"phone" search:"type:contains;column:phone;table:sys_user"`
	Email     string `json:"email" form:"email" search:"type:contains;column:email;table:sys_user"`
}

type SysUserEditForm struct {
	ID        string `json:"id" form:"id" binding:"required" column:"id" zh:"用户ID"`
	LoginName string `json:"login_name" form:"login_name" column:"login_name"`
	RealName  string `json:"real_name" form:"real_name" column:"real_name"`
	Phone     string `json:"phone" form:"phone" column:"phone"`
	Email     string `json:"email" form:"email" column:"email"`
	Status    int    `column:"status"`
	RoleIds   string `column:"role_ids"`
}

type SysUserAddForm struct {
	LoginName string `json:"login_name" form:"login_name"`
	Password  string `json:"password" form:"password"`
	RealName  string `json:"real_name" form:"real_name"`
	Phone     string `json:"phone" form:"phone"`
	Email     string `json:"email" form:"email"`
	Status    int
	RoleIds   string
}

type SysUserStatusForm struct {
	ID     string `json:"id" form:"id" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}

type AvatarForm struct {
	Avatar string `json:"avatar" form:"avatar" binding:"required" zh:"头像"`
}

type ProfileForm struct {
	RealName string `json:"real_name" form:"real_name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Remark   string `json:"remark" form:"remark"`
}

type PasswordForm struct {
	OldPwd     string `json:"old_pwd" form:"old_pwd" binding:"required" zh:"原密码"`
	NewPwd     string `json:"new_pwd" form:"new_pwd" binding:"required" zh:"新密码"`
	ConfirmPwd string `json:"confirm_pwd" form:"confirm_pwd" binding:"required" zh:"重复密码"`
}
