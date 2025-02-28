package baseapi

import (
	"encoding/json"
	"gsadmin/app/model"
	"gsadmin/core/log"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/session"
	"gsadmin/global/e"
	"sync"
)

var sessionMap sync.Map

// GetProfile 获得用户信息详情
func (a *Api) GetUserFromSession() *model.SysUser {
	tmp := session.Get(a.c, e.UserInfo)
	if tmp == nil {
		return nil
	}
	u := new(model.SysUser)
	if s, ok := tmp.(string); ok {
		err := json.Unmarshal([]byte(s), &u)
		if err != nil {
			return nil
		}
	}
	return u
}

// GetUid 获得当前用户id
func (a *Api) GetUidFromSession() int {
	uid := session.Get(a.c, e.SysAuth)
	if uid != nil {
		//双重验证
		//_, ok := sessionMap.Load(uid)
		//if !ok {
		//	return 0
		//}
		return assertion.AnyToInt(uid)
	}

	return 0
}

// SetUserToSession
func (a *Api) SetUserToSession(user model.SysUser) (string, error) {
	err := session.Del(a.c, e.SysAuth)
	if err != nil {
		return "", err
	}
	err = session.Set(a.c, e.SysAuth, user.ID)
	if err != nil {
		log.Instance().Warn(err.Error())
		return "", err
	}
	tmp, _ := json.Marshal(user)

	err = session.Set(a.c, e.UserInfo, string(tmp))
	if err != nil {
		log.Instance().Warn(err.Error())
		return "", err
	}
	sessionMap.Store(user.ID, a.c)
	return assertion.AnyToString(user.ID), nil
}

// DelUserFromSession
func (a *Api) DelUserFromSession(user model.SysUser) error {
	sessionMap.Delete(user.ID)
	err := session.Del(a.c, e.SysAuth)
	if err != nil {
		return err
	}
	err = session.Del(a.c, e.UserInfo)
	if err != nil {
		return err
	}
	return nil
}
