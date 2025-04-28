package dto

type SysRoleListForm struct {
	PageForm `search:"-"`
	ID       string `json:"id" form:"id" search:"type:exact;column:id;table:sys_role"`
	RoleName string `json:"role_name" form:"role_name" search:"type:contains;column:role_name;table:sys_role"`
	Detail   string `json:"detail" form:"detail" search:"type:contains;column:detail;table:sys_role"`
}

type SysRoleEditForm struct {
	ControlForm
	ID       string `json:"id" form:"id" binding:"required" zh:"权限ID" column:"id"`
	RoleName string `json:"role_name" form:"role_name" zh:"权限名称" column:"role_name"`
	Detail   string `json:"detail" form:"detail" column:"detail"`
	Status   string `json:"status" form:"status" column:"status"`
	//NodesData string `json:"nodes_data" form:"nodes_data"`
}

type SysRoleAddForm struct {
	ControlForm
	RoleName  string `json:"role_name" form:"role_name" binding:"required" zh:"权限名称"`
	Detail    string `json:"detail" form:"detail"`
	Status    string `json:"status" form:"status"`
	NodesData string `json:"nodes_data" form:"nodes_data"`
}

type SysRolePowerForm struct {
	RoleId   string `json:"roleId" form:"roleId"`
	PowerIds string `json:"powerIds" form:"powerIds"`
}

type SysRolePowerIDForm struct {
	RoleId string `json:"roleId" form:"roleId"`
}

type SysRoleIDForm struct {
	ID int `json:"id" form:"id" search:"type:exact;column:id;table:sys_role"`
}

type SysRoleStatusForm struct {
	ID     int `json:"id" form:"id"`
	Status int `json:"status" form:"status"`
}

type SysRoleResp struct {
	ID        int    `json:"id"`
	RoleName  string `json:"role_name"`
	Detail    string `json:"detail"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status    int    `json:"status"`
}

// 管理员信息修改页面展示
type RoleEditShow struct {
	ID       int
	RoleName string
	Status   int
	Checked  int
}
