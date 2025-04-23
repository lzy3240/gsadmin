package baseapi

import (
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/core/baseapi/constructor"
	"gsadmin/core/utils/validate"
)

type Api struct {
	C *gin.Context
	r *CommonResp
	//Error error
}

// AddError 处理错误
//func (a *Api) AddError(err error) {
//	if a.Error == nil {
//		a.Error = err
//	} else if a.Error != nil {
//		log.Instance().Error(err.Error())
//		a.Error = fmt.Errorf("%v;%w", a.Error, err)
//	}
//}

// MountCtx 挂载上下文
func (a *Api) MountCtx(c *gin.Context) *Api {
	if a.C == nil {
		a.C = c
	}
	return a
}

// Bind 参数校验
func (a *Api) Bind(d interface{}, bindings ...binding.Binding) error {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.Constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = a.C.ShouldBindUri(d)
		} else {
			err = a.C.ShouldBindWith(d, bindings[i])
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

// TransErr 转换中文错误
func (a *Api) TransErr(err error) string {
	return validate.GetValidateError(err)
}
