package service

import (
	"errors"
	"go.uber.org/zap"
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/cache"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/global/e"
	"gsadmin/utils/assertion"
	"gsadmin/utils/str"
	"strings"
	"time"
)

type SysUser struct {
	baseservice.Service
}

func (s *SysUser) List(req *dto.SysUserListForm) (list []dto.SysUserShow, count int, err error) {
	var users []model.SysUser
	err = db.Instance().Scopes(
		s.SetPaginate(req.Limit, req.Page),
		s.SetCondition(*req),
	).Model(model.SysUser{}).Order("id desc").Find(&users).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		log.Instance().Error("SysUserService.List:" + err.Error())
		return
	}

	for _, v := range users {
		list = append(list, dto.SysUserShow{
			ID:        v.ID,
			LoginName: v.LoginName,
			RealName:  v.RealName,
			RoleIds:   v.RoleIds,
			Level:     v.Level,
			Phone:     v.Phone,
			Email:     v.Email,
			Avatar:    v.Avatar,
			Status:    v.Status,
			LastIp:    v.LastIp,
			LastLogin: v.LastLogin,
		})
	}
	return list, count, nil
}

func (s *SysUser) UpdateUserStatus(req *dto.SysUserStatusForm) error {
	user, err := getByID(assertion.AnyToUint(req.ID))

	if err != nil {
		log.Instance().Error("UpdateUserStatus.FindUser", zap.Error(err))
		return err
	}
	if user.ID < 1 {
		return errors.New("未找到用户")
	}

	err = db.Instance().Model(model.SysUser{}).Where("id = ?", req.ID).Update("status", req.Status).Error
	return err
}

func (s *SysUser) UpdateUserAttr(req *dto.SysUserEditForm) error {
	user, err := getByID(assertion.AnyToUint(req.ID))
	if err != nil {
		log.Instance().Error("UpdateUserAttr.FindUser:" + err.Error())
		return err
	}
	if user.ID < 1 {
		return errors.New("未找到用户")
	}
	attr := make(map[string]interface{})
	attr["status"] = req.Status
	attr["role_ids"] = req.RoleIds
	attr["login_name"] = req.LoginName
	attr["real_name"] = req.RealName
	attr["phone"] = req.Phone
	attr["email"] = req.Email

	err = db.Instance().Model(model.SysUser{}).Where("id = ?", req.ID).Updates(attr).Error
	if err != nil {
		log.Instance().Error("UpdateUserAttr.FindUser:" + err.Error())
		return err
	}
	return nil
}

func (s *SysUser) UserEdit(req *dto.SysUserIDForm) (show dto.SysUserShow, rolesShow []dto.RoleEditShow, err error) { //uid string
	roleSvice := SysRole{}
	user, err := getByID(assertion.AnyToUint(req.ID))
	if err != nil {
		log.Instance().Error("UserEdit.FindUser", zap.Error(err))
		return show, rolesShow, err
	}
	err = str.CopyFields(&show, user)
	if err != nil {
		log.Instance().Error("UserEdit.CopyFields", zap.Error(err))
		return show, rolesShow, err
	}
	roles, _, err := roleSvice.ListByStatus(1) // 查找全部的分组
	if err != nil {
		log.Instance().Error("UserEdit.FindRoles", zap.Error(err))
		return show, rolesShow, err
	}
	check := strings.Split(user.RoleIds, ",")
	for _, v := range roles {
		checked := 0
		if str.IsContain(check, assertion.AnyToString(v.ID)) {
			checked = 1
		}
		rolesShow = append(rolesShow, dto.RoleEditShow{
			ID:       assertion.AnyToInt(v.ID),
			RoleName: v.RoleName,
			Status:   v.Status,
			Checked:  checked,
		})
	}
	return show, rolesShow, nil
}

func (s *SysUser) UserAdd(roleIds, status string, req *dto.SysUserAddForm, userId int) error {
	if str.IsEmail([]byte(req.Email)) == false {
		return errors.New("邮箱格式不符合规范")
	}
	req.RoleIds = roleIds
	req.Status = assertion.AnyToInt(status)
	if err := createUser(req, userId); err != nil {
		log.Instance().Error("UserAdd.createUser:" + err.Error())
		return err
	}
	return nil
}

func (s *SysUser) Delete(req *dto.SysUserIDForm, operId uint) error {
	var user model.SysUser
	err := db.Instance().Model(model.SysUser{}).Where("id = ?", req.ID).First(&user).Error
	if err != nil {
		log.Instance().Error("SysUserService.Delete:" + err.Error())
		return err
	}
	if user.ID == 1 || user.Level == 99 {
		return errors.New("超级管理员用户不能删除")
	}

	if user.ID == operId {
		return errors.New("不能删除本人账号")
	}

	err = db.Instance().Where("id = ?", req.ID).Delete(model.SysUser{}).Error
	if err != nil {
		log.Instance().Error("SysUserService.Delete:" + err.Error())
		return err
	}
	return nil
}

// IsAdmin 判断是否是超级管理员
func (s *SysUser) IsAdmin(user *model.SysUser) bool {
	return user.ID == 1 && user.Level == 99 // 只有id = 1 && level = 99 的超级管理员
}

// SignIn 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *SysUser) SignIn(loginName, password string, clientIP string) (user model.SysUser, err error) {
	//查询用户信息
	user, err = getByName(loginName)
	if err != nil || (user.Password != str.Md5([]byte(password+user.Salt))) {
		return model.SysUser{}, errors.New("账号或密码错误")
	}
	//更新登录信息

	tmp := model.SysUser{}
	tmp.LastIp = clientIP
	tmp.LastLogin = time.Now().Format(e.TimeFormat)
	err = db.Instance().Model(model.SysUser{}).Where("id = ?", user.ID).Updates(&tmp).Error
	if err != nil {
		log.Instance().Error("SysUserService.SignIn:" + err.Error())
		return user, err
	}

	return user, nil
}

// ClearMenuCache 清空用户菜单缓存
//func (s *SysUser) ClearMenuCache(user *model.SysUser) {
//	if s.IsAdmin(user) {
//		cache.Instance().Delete(e.Menu)
//	} else {
//		cache.Instance().Delete(e.Menu + assertion.AnyToString(user.ID))
//	}
//}

// SignOut 用户注销
func (s *SysUser) SignOut(operUser *model.SysUser) error {
	if operUser != nil {
		_ = cache.Instance().Del(e.UserMenu, e.MenuCache+assertion.AnyToString(operUser.ID))
	}

	online := SysUserOnline{}
	err := online.Delete(assertion.AnyToString(operUser.ID))
	if err != nil {
		log.Instance().Error("SysUserService.SignOut:" + err.Error())
		return err
	}
	return nil
}

func (s *SysUser) GetAuth(user *model.SysUser) []model.SysAuth {
	authService := SysAuth{}
	roleAuthService := SysRoleAuth{}
	var auths []model.SysAuth
	if user.Level != 99 {
		authIds, _ := roleAuthService.ListByRoleIDs(user.RoleIds)
		auths, _, _ = authService.ListAuthByIDs(authIds)
		return auths
	}

	auths, _, _ = authService.ListAuth(&dto.SysAuthListForm{})
	return auths
}

// UpdateAvatar 更新用户头像
func (s *SysUser) UpdateAvatar(req *dto.AvatarForm, operUser *model.SysUser) error {
	if strings.HasPrefix(req.Avatar, "/") == false {
		req.Avatar = "/" + req.Avatar
	}
	err := db.Instance().Model(model.SysUser{}).Where("id = ?", operUser.ID).Updates(map[string]interface{}{"avatar": req.Avatar}).Error
	if err != nil {
		log.Instance().Error("SysUserService.UpdateAvatar:" + err.Error())
		return err
	}
	return nil
}

func (s *SysUser) UpdateProfile(req *dto.ProfileForm, operUser *model.SysUser) error {
	var user model.SysUser
	user.RealName = req.RealName
	user.Remark = req.Remark
	user.Phone = req.Phone
	user.Email = req.Email
	err := db.Instance().Model(model.SysUser{}).Where("id = ?", operUser.ID).Updates(&user).Error
	if err != nil {
		log.Instance().Error("SysUserService.UpdateProfile:" + err.Error())
		return err
	}
	return nil
}

func (s *SysUser) GetLoginInfo(operUser *model.SysUser) ([]model.SysLoginLog, error) {
	loginLog := SysLog{}
	req := dto.SysLoginLogForm{}
	req.LoginName = operUser.LoginName
	req.Page = 1
	req.Limit = 5
	list, _, err := loginLog.LoginLogList(&req)
	if err != nil {
		log.Instance().Error("SysUserservice.GetLoginInfo:" + err.Error())
		return nil, err
	}
	return list, nil
}

// UpdatePwd 修改用户密码
func (s *SysUser) UpdatePwd(profile *dto.PasswordForm, operUser *model.SysUser) error {
	if profile.NewPwd == profile.OldPwd {
		return errors.New("新旧密码不能相同")
	}
	if profile.ConfirmPwd != profile.NewPwd {
		return errors.New("确认密码不一致")
	}
	//校验密码
	user, _ := getByID(operUser.ID)
	if user.Password != str.Md5([]byte(profile.OldPwd+user.Salt)) {
		return errors.New("原密码不正确")
	}

	//新校验密码
	pwd, salt := str.SetPassword(e.DefaultSaltLen, profile.NewPwd)

	tmp := model.SysUser{}
	tmp.Password = pwd
	tmp.Salt = salt
	err := db.Instance().Model(model.SysUser{}).Where("id = ?", operUser.ID).Updates(&tmp).Error

	if err != nil {
		log.Instance().Error("SysUserService.UpdatePwd:" + err.Error())
		return errors.New("保存数据失败")
	}
	// 密码和盐不存session 此处不用更新session缓存
	return nil
}

func getByName(name string) (user model.SysUser, err error) {
	err = db.Instance().Model(model.SysUser{}).Where("login_name = ?", name).First(&user).Error
	return
}

func getByID(userId uint) (user model.SysUser, err error) {
	err = db.Instance().Model(model.SysUser{}).Where("id = ?", userId).First(&user).Error
	return
}

// checkLoginName 检查登陆名是否存在,存在返回true,否则false
func checkLoginName(loginName string) bool {
	if user, err := getByName(loginName); err != nil || user.ID < 1 { //dao.NewSysUserDaoImpl().FindByName(loginName)
		return false
	} else {
		return true
	}
}

func createUser(r *dto.SysUserAddForm, uid int) error {
	if checkLoginName(r.LoginName) {
		return errors.New("用户名已存在")
	}
	pwd, salt := str.SetPassword(e.DefaultSaltLen, r.Password)
	var sysUser model.SysUser
	sysUser.LoginName = r.LoginName
	sysUser.Password = pwd
	sysUser.Salt = salt
	sysUser.RealName = r.RealName
	sysUser.RoleIds = r.RoleIds
	sysUser.Level = 1
	sysUser.Phone = r.Phone
	sysUser.Email = r.Email
	sysUser.Status = r.Status
	sysUser.CreateId = uid
	sysUser.UpdateId = uid
	sysUser.LastLogin = time.Now().Format(e.TimeFormat)

	err := db.Instance().Create(&sysUser).Error
	if err != nil {
		log.Instance().Error("createUser.Insert", zap.Error(err))
		return err
	}
	return nil
}
