package model

type SysRoleAuth struct {
	ID     uint
	RoleID uint64
	AuthID uint
}

func (SysRoleAuth) TableName() string {
	return "sys_role_auth"
}
