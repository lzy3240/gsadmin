package service

import (
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/cache"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/str"
	"gsadmin/global/e"
	"strconv"
	"time"
)

type SysBase struct {
	baseservice.Service
}

func (s *SysBase) CheckLock(loginName string) bool {
	res, ok := cache.Instance().Get(e.UserLock + loginName)
	if ok && res == true {
		return true
	}
	return false
}

func (s *SysBase) SetPwdErrNum(loginName string) int {
	countNum := 0
	errNum, _ := cache.Instance().Get(e.UserLoginErr + loginName)
	if errNum != nil {
		countNum = assertion.AnyToInt(errNum)
	}
	countNum = countNum + 1
	cache.Instance().Set(e.UserLoginErr+loginName, countNum, time.Minute*1)
	if countNum >= 5 {
		lock(loginName)
	}
	return countNum
}

func (s *SysBase) RemovePwdErrNum(loginName string) {
	cache.Instance().Delete(e.UserLoginErr + loginName)
}

// MenuServiceV2 构造菜单
func (s *SysBase) MenuServiceV2(user *model.SysUser) (cacheMenu dto.CacheMenuV2) {
	userSvice := SysUser{}
	hasMenu, found := cache.Instance().Get(e.MenuCache + assertion.AnyToString(user.ID))
	if found && hasMenu != nil {
		cacheMenu = hasMenu.(dto.CacheMenuV2)
	} else {
		var authId []string
		result := userSvice.GetAuth(user)
		allowUrl := ""
		for _, v := range result {
			if !str.IsContain([]string{"", "/"}, v.AuthUrl) {
				if allowUrl == "" {
					allowUrl += v.AuthUrl
				} else {
					allowUrl += "," + v.AuthUrl
				}
			}
			authId = append(authId, strconv.Itoa(int(v.ID)))
		}

		authService := SysAuth{}
		auths, _, _ := authService.ListAuth(&dto.SysAuthListForm{PowerType: "0"})
		var menu []dto.MenuResp
		for _, auth := range auths {
			if userSvice.IsAdmin(user) == false && str.IsContain(authId, strconv.Itoa(int(auth.ID))) == false { // 管理员不限制
				continue
			}

			pid := assertion.AnyToString(auth.ID)
			childs, _, _ := authService.ListAuth(&dto.SysAuthListForm{PowerType: "1", Pid: pid})
			var childsResp []dto.MenuChildrenResp
			for _, child := range childs {
				if userSvice.IsAdmin(user) == false && str.IsContain(authId, strconv.Itoa(int(child.ID))) == false { // 管理员不限制
					continue
				}
				childsResp = append(childsResp, dto.MenuChildrenResp{
					ID:       int(child.ID),
					Title:    child.AuthName,
					Type:     child.PowerType,
					Icon:     child.Icon,
					Href:     child.AuthUrl,
					OpenType: "_iframe",
				})
			}
			menu = append(menu, dto.MenuResp{
				ID:       int(auth.ID),
				Title:    auth.AuthName,
				Type:     auth.PowerType,
				Icon:     auth.Icon,
				Href:     auth.AuthUrl,
				Children: childsResp,
			})
		}

		cacheMenu.AllowUrl = allowUrl
		cacheMenu.MenuResp = menu
		cache.Instance().Set(e.MenuCache+assertion.AnyToString(user.ID), cacheMenu, time.Hour)
	}
	return cacheMenu
}

// -------------------------------------------------------------------

func lock(loginName string) {
	cache.Instance().Set(e.UserLock+loginName, true, time.Minute*5)
}

func unLock(loginName string) {
	cache.Instance().Delete(e.UserLock + loginName)
}
