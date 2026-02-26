package baseapi

import (
	"fmt"
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/core/baseapi/constructor"
	"gsadmin/core/baseapi/response"
	"gsadmin/global/e"
	"gsadmin/utils/translator"
	"net/http"
)

type Api struct {
	c *gin.Context
	r *response.CommonResp
}

// Bind 参数校验
func (a *Api) Bind(c *gin.Context, d interface{}, bindings ...binding.Binding) error {
	if a.c == nil {
		a.c = c
	}
	var err error
	if len(bindings) == 0 {
		bindings = constructor.Constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = c.ShouldBindUri(d)
		} else {
			err = c.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			err = nil
			continue
		}
		if err != nil {
			return err
		}
	}

	if err1 := vd.Validate(d); err1 != nil {
		return err1
	}

	return nil
}

// Success 返回成功消息体
func (a *Api) Success(c *gin.Context, msg string) *Api {
	if msg == "" {
		msg = "操作成功"
	}
	a.r = &response.CommonResp{
		Code: e.SUCCESS,
		Msg:  msg,
	}
	if a.c == nil {
		a.c = c
	}
	return a
}

// Error 返回错误消息体
func (a *Api) Error(c *gin.Context, msg string, err error) *Api {
	if msg == "" {
		msg = "操作失败"
	}
	if err != nil {
		errStr := translator.GetValidateError(err)
		tmp := fmt.Sprintf("%s --> %s", msg, errStr)
		msg = tmp
	}
	a.r = &response.CommonResp{
		Code: e.ERROR,
		Msg:  msg,
	}
	if a.c == nil {
		a.c = c
	}
	return a
}

// Forbidden 返回拒绝消息体
func (a *Api) Forbidden(c *gin.Context, msg string) *Api {
	if msg == "" {
		msg = "无操作权限"
	}
	a.r = &response.CommonResp{
		Code: e.FORBIDDEN,
		Msg:  msg,
	}
	if a.c == nil {
		a.c = c
	}
	return a
}

// Unauthorized 返回认证失败消息体
func (a *Api) Unauthorized(c *gin.Context, msg string) *Api {
	if msg == "" {
		msg = "鉴权失败"
	}
	a.r = &response.CommonResp{
		Code: e.UNAUTHORIZED,
		Msg:  msg,
	}
	if a.c == nil {
		a.c = c
	}
	return a
}

// Custom 返回自定义消息体
func (a *Api) Custom(c *gin.Context, code int, msg string) *Api {
	a.r = &response.CommonResp{
		Code: code,
		Msg:  msg,
	}
	if a.c == nil {
		a.c = c
	}
	return a
}

// SetMsg 设置消息体的内容
//func (a *Api) SetMsg(msg string) *Api {
//	a.r.Msg = msg
//	return a
//}

// SetCode 设置消息体的编码
//func (a *Api) SetCode(code int) *Api {
//	a.r.Code = code
//	return a
//}

// SetData 设置数据
func (a *Api) SetData(data interface{}) *Api {
	a.r.Data = data
	return a
}

// SetPageData 设置分页数据
func (a *Api) SetPageData(count int, list interface{}) *Api {
	a.r.Count = count
	a.r.Data = list
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
	a.c.Set("result", a.r)
	a.c.HTML(http.StatusOK, page, data)
	a.c.Abort()
}

// WriteStringExit 输出String
func (a *Api) WriteStringExit(format string, value ...any) {
	a.r.Data = fmt.Sprint(format, value)
	a.c.Set("result", a.r)
	a.c.String(http.StatusOK, format, value...)
	a.c.Abort()
}

// WriteJsonExit 输出json到客户端
func (a *Api) WriteJsonExit() {
	a.c.Set("result", a.r)
	a.c.JSON(http.StatusOK, a.r)
	a.c.Abort()
}

// WriteCustomJsonExit 兼容个性化json写法
func (a *Api) WriteCustomJsonExit(data any) {
	a.c.Set("result", a.r)
	a.c.JSON(http.StatusOK, data)
	a.c.Abort()
}

// WriteRedirect 重定向
func (a *Api) WriteRedirect(path string) {
	a.c.Redirect(http.StatusFound, path)
	a.c.Abort()
}
