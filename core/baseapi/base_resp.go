package baseapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gsadmin/global/e"
	"net/http"
)

type CommonResp struct {
	Code  int         `json:"code"`  //响应编码: 200 成功 500 错误 403 无操作权限 401 鉴权失败  -1  失败
	Msg   string      `json:"msg"`   //消息内容
	Data  interface{} `json:"data"`  //数据内容
	Count int         `json:"count"` //数据数量
	Tag   bool        `json:"tag"`   //日志标记
	Type  int         `json:"type"`  //业务类型
	Title string      `json:"title"` //操作名称
}

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}

// SuccessResp 返回一个成功的消息体
func (a *Api) SuccessResp() *Api {
	resp := CommonResp{
		Code: e.SUCCESS,
		Type: e.OperOther,
		Msg:  "操作成功",
	}
	a.r = &resp
	return a
}

// ErrorResp 返回一个错误的消息体
func (a *Api) ErrorResp() *Api {
	resp := CommonResp{
		Code: e.ERROR,
		Type: e.OperOther,
		Msg:  "操作失败",
	}
	a.r = &resp
	return a
}

// ForbiddenResp 返回一个拒绝访问的消息体
func (a *Api) ForbiddenResp() *Api {
	resp := CommonResp{
		Code: e.FORBIDDEN,
		Type: e.OperOther,
		Msg:  "无操作权限",
	}
	a.r = &resp
	return a
}

// UnauthorizedResp 认证失败
func (a *Api) UnauthorizedResp() *Api {
	resp := CommonResp{
		Code: e.UNAUTHORIZED,
		Type: e.OperOther,
		Msg:  "鉴权失败",
	}
	a.r = &resp
	return a
}

// SetMsg 设置消息体的内容
func (a *Api) SetMsg(msg string) *Api {
	a.r.Msg = msg
	return a
}

// SetCode 设置消息体的编码
func (a *Api) SetCode(code int) *Api {
	a.r.Code = code
	return a
}

// SetData 设置消息体的数据
func (a *Api) SetData(data interface{}) *Api {
	a.r.Data = data
	return a
}

// SetCount 设置消息体的业务类型
func (a *Api) SetCount(count int) *Api {
	a.r.Count = count
	return a
}

// SetLogTag 设置日志标识信息
func (a *Api) SetLogTag(buType int, opTitle string) *Api {
	a.r.Tag = true
	a.r.Type = buType
	a.r.Title = opTitle
	return a
}

// WriteHtmlExit 输出Html页面
func (a *Api) WriteHtmlExit(page string, data gin.H) {
	a.r.Data = data
	a.C.Set("result", a.r)
	a.C.HTML(http.StatusOK, page, data)
	a.C.Abort()
}

// WriteStringExit 输出String
func (a *Api) WriteStringExit(format string, value ...any) {
	a.r.Data = fmt.Sprint(format, value)
	a.C.Set("result", a.r)
	a.C.String(http.StatusOK, format, value...)
	a.C.Abort()
}

// WriteJsonExit 输出json到客户端
func (a *Api) WriteJsonExit() {
	a.C.Set("result", a.r)
	a.C.JSON(http.StatusOK, a.r)
	a.C.Abort()
}

// WriteCustomJsonExit 兼容个性化json写法
func (a *Api) WriteCustomJsonExit(data any) {
	a.C.Set("result", a.r)
	a.C.JSON(http.StatusOK, data)
	a.C.Abort()
}

// WriteRedirect 重定向
func (a *Api) WriteRedirect(path string) {
	a.C.Redirect(http.StatusFound, path)
	a.C.Abort()
}
