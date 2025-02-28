package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/str"
	"gsadmin/global/e"
	"strings"
)

type SysUser struct {
	baseapi.Api
}

//---------------用户信息管理------------------

func (a SysUser) UserAddPage(c *gin.Context) {
	a.MountCtx(c)
	roleSvice := service.SysRole{}
	var rolesShow []dto.RoleEditShow
	roles, _, err := roleSvice.ListByStatus(1)
	if err != nil {
		a.ErrorResp().SetLogTag(e.OperAdd, e.UserAdd).WriteStringExit("%s", err.Error())
		return
	}
	for _, v := range roles {
		rolesShow = append(rolesShow, dto.RoleEditShow{
			ID:       assertion.AnyToInt(v.ID),
			RoleName: v.RoleName,
			Status:   v.Status,
		})
	}
	a.SuccessResp().SetLogTag(e.OperAdd, e.UserAdd).WriteHtmlExit("user_add.html", gin.H{"role": rolesShow})
}

func (a SysUser) UserAdd(c *gin.Context) {
	a.MountCtx(c)
	userSvice := service.SysUser{}

	roles := c.PostFormArray("role_ids")
	roleIds := str.Array2Str(roles)
	status := c.PostForm("status")
	if status == "" {
		status = "0"
	} else if status == "on" {
		status = "1"
	}
	var req dto.SysUserAddForm
	if err := a.Bind(&req, binding.Form); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.UserAdd).WriteJsonExit()
		return
	}

	user := a.GetUserFromSession()
	if err := userSvice.UserAdd(roleIds, status, &req, assertion.AnyToInt(user.ID)); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.UserAdd).WriteJsonExit()
		return
	}

	a.SuccessResp().SetMsg("创建成功!").SetLogTag(e.OperAdd, e.UserAdd).WriteJsonExit()
}

func (a SysUser) UserListPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("user_list.html", gin.H{})
}

func (a SysUser) UserEditPage(c *gin.Context) {
	userSvice := service.SysUser{}
	req := dto.SysUserIDForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)

	if err != nil || req.ID == "" {
		a.ErrorResp().WriteStringExit("请检查参数")
		return
	}
	show, rolesShow, err := userSvice.UserEdit(&req)
	if err != nil {
		a.ErrorResp().SetLogTag(e.OperEdit, e.UserEdit).WriteStringExit("%s", a.TransErr(err))
		return
	}
	a.SuccessResp().SetLogTag(e.OperEdit, e.UserEdit).WriteHtmlExit("user_edit.html", gin.H{"show": show, "role": rolesShow})
}

func (a SysUser) UserEdit(c *gin.Context) {
	a.MountCtx(c)
	userSvice := service.SysUser{}

	roles := c.PostFormArray("role_ids") //待优化
	roleIds := str.Array2Str(roles)
	status := c.PostForm("status") //待优化
	if status == "" {
		status = "0"
	} else if status == "on" {
		status = "1"
	}
	var req dto.SysUserEditForm
	if err := a.Bind(&req, binding.Form); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.UserEdit).WriteJsonExit()
		return
	}

	req.RoleIds = roleIds
	req.Status = assertion.AnyToInt(status)

	if err := userSvice.UpdateUserAttr(&req); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.UserEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetMsg("更新成功").SetLogTag(e.OperEdit, e.UserEdit).WriteJsonExit()
}

func (a SysUser) UserStatus(c *gin.Context) {
	userSvice := service.SysUser{}
	operUser := a.MountCtx(c).GetUserFromSession()
	if userSvice.IsAdmin(operUser) == false {
		a.ErrorResp().SetMsg("权限不足，无法修改").SetLogTag(e.OperEdit, e.UserEdit).WriteJsonExit()
		return
	}

	req := dto.SysUserStatusForm{}
	err := a.Bind(&req, binding.Form)
	if err != nil || (req.ID == "" || req.Status == "") { //
		a.ErrorResp().SetMsg("参数不足").SetLogTag(e.OperEdit, e.UserEdit).WriteJsonExit()
		return
	}

	err = userSvice.UpdateUserStatus(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.UserEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetMsg("更新成功").SetLogTag(e.OperEdit, e.UserEdit).WriteJsonExit()
}

func (a SysUser) UserJson(c *gin.Context) {
	req := dto.SysUserListForm{}
	userSvice := service.SysUser{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}

	list, count, err := userSvice.List(&req)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}

	a.SuccessResp().SetCode(0).SetCount(count).SetData(list).WriteJsonExit()
}

func (a SysUser) UserDelete(c *gin.Context) {
	req := dto.SysUserIDForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg("id不能为空").SetLogTag(e.OperDel, e.UserDelete).WriteJsonExit()
		return
	}

	operUser := a.GetUserFromSession()
	userSvice := service.SysUser{}
	err = userSvice.Delete(&req, operUser.ID)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperDel, e.UserDelete).WriteJsonExit()
		return
	}

	a.SuccessResp().SetLogTag(e.OperDel, e.UserDelete).WriteJsonExit()
}

func (a SysUser) UserBatchDelete(c *gin.Context) {
	req := dto.SysUserIDsForm{}
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperDel, e.UserDelete).WriteJsonExit()
		return
	}

	tmpIds := str.SplitStr(req.IDs)
	var failedIds []string
	userSvice := service.SysUser{}

	operUser := a.GetUserFromSession()
	for _, id := range tmpIds {
		if err = userSvice.Delete(&dto.SysUserIDForm{ID: id}, operUser.ID); err != nil {
			failedIds = append(failedIds, id)
		}
	}

	if err != nil {
		a.ErrorResp().SetMsg("部分删除失败, id: "+strings.Join(failedIds, ",")).SetLogTag(e.OperDel, e.UserDelete).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperDel, e.UserDelete).WriteJsonExit()
}

//--------------个人资料修改----------------

func (a SysUser) ProfilePage(c *gin.Context) {
	a.MountCtx(c)
	operUser := a.GetUserFromSession()
	userSvice := service.SysUser{}

	if operUser.Avatar == "" {
		operUser.Avatar = e.DefaultAvatar
	}
	info, _ := userSvice.GetLoginInfo(operUser)
	a.SuccessResp().SetLogTag(e.OperEdit, e.ProfileEdit).WriteHtmlExit("user_show.html", gin.H{"user": operUser, "info": info})
}

func (a SysUser) ProfileEdit(c *gin.Context) {
	a.MountCtx(c)
	operUser := a.GetUserFromSession()
	userSvice := service.SysUser{}

	var pro dto.ProfileForm
	err := a.Bind(&pro, binding.JSON)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.ProfileEdit).WriteJsonExit()
		return
	}

	err = userSvice.UpdateProfile(&pro, operUser)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.ProfileEdit).WriteJsonExit()
		return
	}

	//更新session
	operUser.RealName = pro.RealName
	operUser.Email = pro.Email
	operUser.Remark = pro.Remark
	operUser.Phone = pro.Phone
	_, _ = a.SetUserToSession(*operUser)
	a.SuccessResp().SetLogTag(e.OperEdit, e.ProfileEdit).WriteJsonExit()
}

func (a SysUser) UploadPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("upload_profile.html", nil)
}

func (a SysUser) AvatarEdit(c *gin.Context) {
	userSvice := service.SysUser{}
	req := dto.AvatarForm{}
	err := a.MountCtx(c).Bind(&req, binding.JSON)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.AvatarEdit).WriteJsonExit()
		return
	}

	operUser := a.GetUserFromSession()

	err = userSvice.UpdateAvatar(&req, operUser)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.AvatarEdit).WriteJsonExit()
		return
	}
	//更新session
	operUser.Avatar = req.Avatar
	_, _ = a.SetUserToSession(*operUser)
	a.SuccessResp().SetMsg(req.Avatar).SetLogTag(e.OperEdit, e.AvatarEdit).WriteJsonExit()
}

func (a SysUser) PwdEditPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().SetLogTag(e.OperEdit, e.UserEdit).WriteHtmlExit("user_pwd.html", gin.H{})
}

func (a SysUser) PwdEdit(c *gin.Context) {
	a.MountCtx(c)

	var pwd dto.PasswordForm
	err := a.Bind(&pwd, binding.JSON)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.PwdEditHandler).WriteJsonExit()
		return
	}
	userSvice := service.SysUser{}

	operUser := a.GetUserFromSession()

	if err = userSvice.UpdatePwd(&pwd, operUser); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.PwdEditHandler).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperEdit, e.PwdEditHandler).WriteJsonExit()
}
