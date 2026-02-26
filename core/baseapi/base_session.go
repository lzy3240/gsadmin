package baseapi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gsadmin/app/model"
	"gsadmin/core/log"
	"gsadmin/global/e"
	"gsadmin/utils/assertion"
	"gsadmin/utils/session"
	"sync"
)

var sessionMap sync.Map

// GetUserFromSession 获得用户信息详情
func (a *Api) GetUserFromSession(c *gin.Context) *model.SysUser {
	tmp := session.Get(c, e.UserInfo)
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

// GetUidFromSession 获得当前用户id
func (a *Api) GetUidFromSession(c *gin.Context) int {
	uid := session.Get(c, e.SysAuth)
	if uid != nil {
		//双重验证
		_, ok := sessionMap.Load(uid)
		if !ok {
			return 0
		}
		return assertion.AnyToInt(uid)
	}

	return 0
}

// SetUserToSession 设置用户信息
func (a *Api) SetUserToSession(c *gin.Context, user model.SysUser) (string, error) {
	err := session.Del(c, e.SysAuth)
	if err != nil {
		return "", err
	}
	err = session.Set(c, e.SysAuth, user.ID)
	if err != nil {
		log.Instance().Warn(err.Error())
		return "", err
	}
	userInfo, _ := json.Marshal(user)

	err = session.Set(c, e.UserInfo, string(userInfo))
	if err != nil {
		log.Instance().Warn(err.Error())
		return "", err
	}
	sessionMap.Store(user.ID, c)
	return assertion.AnyToString(user.ID), nil
}

// DelUserFromSession 删除用户信息
func (a *Api) DelUserFromSession(c *gin.Context, user model.SysUser) error {
	sessionMap.Delete(user.ID)
	err := session.Del(c, e.SysAuth)
	if err != nil {
		return err
	}
	err = session.Del(c, e.UserInfo)
	if err != nil {
		return err
	}
	return nil
}
