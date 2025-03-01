package service

import (
	"encoding/json"
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/cache"
	"gsadmin/core/log"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/str"
	"gsadmin/global/e"
	"strconv"
)

type SysBase struct {
	baseservice.Service
}

func (s *SysBase) CheckLock(loginName string) bool {
	res, err := cache.Instance().Get(e.SysLogin, e.UserLock+loginName)
	if err == nil && res == "true" {
		return true
	}
	return false
}

func (s *SysBase) SetPwdErrNum(loginName string) int {
	countNum := 0
	errNum, err := cache.Instance().Get(e.SysLogin, e.UserPwdErr+loginName)
	if err == nil {
		countNum = assertion.AnyToInt(errNum)
	}
	countNum = countNum + 1
	_ = cache.Instance().Set(e.SysLogin, e.UserPwdErr+loginName, assertion.AnyToString(countNum), 60) //60s
	if countNum >= 5 {
		lock(loginName)
	}
	return countNum
}

func (s *SysBase) RemovePwdErrNum(loginName string) {
	_ = cache.Instance().Del(e.SysLogin, e.UserPwdErr+loginName)
}

// MenuServiceV2 构造菜单
func (s *SysBase) MenuServiceV2(user *model.SysUser) (cacheMenu dto.CacheMenuV2) {
	userSvice := SysUser{}
	menuCache, err := cache.Instance().Get(e.UserMenu, e.MenuCache+assertion.AnyToString(user.ID))
	if err == nil { //found && menuCache
		err = json.Unmarshal([]byte(menuCache), &cacheMenu)
		if err != nil {
			log.Instance().Error("Unmarshal menu cache failed: " + err.Error())
		}
		//cacheMenu = menuCache.(dto.CacheMenuV2)
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
		b, _ := json.Marshal(cacheMenu)
		_ = cache.Instance().Set(e.UserMenu, e.MenuCache+assertion.AnyToString(user.ID), string(b), e.MenuCacheTime) //用户菜单缓存1小时, 可考虑与token有效期时间一致
	}
	return cacheMenu
}

// -------------------------------------------------------------------

func lock(loginName string) {
	_ = cache.Instance().Set(e.SysLogin, e.UserLock+loginName, "true", e.UserLockTime) //5次失败之后, 5分钟内禁止登录
}

func unLock(loginName string) {
	_ = cache.Instance().Del(e.SysLogin, e.UserLock+loginName)
}
