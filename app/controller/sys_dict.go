package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/global/e"
)

type SysDict struct {
	baseapi.Api
}

//------------dict type operation-------------

func (a SysDict) DictTypeListPage(c *gin.Context) {
	a.Success(c, "操作成功").WriteHtmlExit("dict_type_list.html", gin.H{})
}

func (a SysDict) DictTypeJson(c *gin.Context) {
	dictSvice := service.SysDict{}
	var req dto.DictTypeListForm
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
		return
	}

	list, count, err := dictSvice.DictTypeListJson(&req)
	if err != nil {
		a.Error(c, "查询失败", err).WriteJsonExit()
		return
	}
	a.Custom(c, e.ZERO, "查询成功").SetPageData(count, list).WriteJsonExit()
}

func (a SysDict) DictTypeStatus(c *gin.Context) {
	dictSvice := service.SysDict{}
	req := dto.DictTypeStatusForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.DictTypeEdit).WriteJsonExit()
		return
	}

	err = dictSvice.UpdateDictTypeStatus(&req)
	if err != nil {
		a.Error(c, "更新失败", err).SetLogTag(e.OperEdit, e.DictTypeEdit).WriteJsonExit()
		return
	}
	a.Success(c, "更新成功").SetLogTag(e.OperEdit, e.DictTypeEdit).WriteJsonExit()
}

func (a SysDict) DictTypeAddPage(c *gin.Context) {
	a.Success(c, "操作成功").SetLogTag(e.OperAdd, e.DictTypeAdd).WriteHtmlExit("dict_type_add.html", gin.H{})
}

func (a SysDict) DictTypeEditPage(c *gin.Context) {
	dictSvice := service.SysDict{}
	//修改页面
	req := dto.DictTypeIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.DictTypeEdit).WriteStringExit("%s", err.Error())
		return
	}

	data, err := dictSvice.GetDictTypeByID(&req)
	if err != nil {
		a.Error(c, "查询失败", err).SetLogTag(e.OperEdit, e.DictTypeEdit).WriteStringExit("%s", err.Error())
		return
	}
	a.Success(c, "查询成功").SetLogTag(e.OperEdit, e.DictTypeEdit).WriteHtmlExit("dict_type_edit.html", gin.H{"dictType": data})
}

// DictTypeSave 兼容新增和修改字典类型
func (a SysDict) DictTypeSave(c *gin.Context) {
	dictSvice := service.SysDict{}
	user := a.GetUserFromSession(c)

	//提交数据接口
	var req dto.DictTypeEditForm
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperOther, e.DictTypeEdit).WriteJsonExit()
		return
	}

	if req.Status == "" {
		req.Status = "0"
	} else if req.Status == "on" {
		req.Status = "1"
	}

	if req.ID == "" {
		//新增
		req.SetCreate(user.ID)
		err = dictSvice.InsertDictType(&req)
		if err != nil {
			a.Error(c, "新增失败", err).SetLogTag(e.OperAdd, e.DictTypeAdd).WriteJsonExit()
			return
		}
		a.Success(c, "新增成功").SetLogTag(e.OperAdd, e.DictTypeAdd).WriteJsonExit()
	} else {
		//更新
		req.SetUpdate(user.ID)
		err = dictSvice.UpdateDictTypeAttr(&req)
		if err != nil {
			a.Error(c, "更新失败", err).SetLogTag(e.OperEdit, e.DictTypeEdit).WriteJsonExit()
			return
		}
		a.Success(c, "更新成功").SetLogTag(e.OperEdit, e.DictTypeEdit).WriteJsonExit()
	}
}

func (a SysDict) DictTypeDelete(c *gin.Context) {
	dictSvice := service.SysDict{}
	req := dto.DictTypeIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperDel, e.DictTypeDelete).WriteJsonExit()
		return
	}

	err = dictSvice.DeleteDictType(&req)
	if err != nil {
		a.Error(c, "删除失败", err).SetLogTag(e.OperDel, e.DictTypeDelete).WriteJsonExit()
		return
	}
	a.Success(c, "删除成功").SetLogTag(e.OperDel, e.DictTypeDelete).WriteJsonExit()
}

//------------dict data operation-------------

func (a SysDict) DictDataListPage(c *gin.Context) {
	req := dto.DictDataJsonForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
		return
	}

	// 传递字典类型
	a.Success(c, "查询成功").WriteHtmlExit("dict_data_list.html", gin.H{"data_type": req.DictType})
}

func (a SysDict) DictDataJson(c *gin.Context) {
	dictSvice := service.SysDict{}
	req := dto.DictDataJsonForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
		return
	}

	list, count, err := dictSvice.DictDataListJson(&req)
	if err != nil {
		a.Error(c, "查询失败", err).WriteJsonExit()
		return
	}
	a.Custom(c, e.ZERO, "查询成功").SetPageData(count, list).WriteJsonExit()
}

func (a SysDict) DictDataAddPage(c *gin.Context) {
	req := dto.DictDataJsonForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperAdd, e.DictDataAdd).WriteJsonExit()
		return
	}
	// 传递字典类型
	a.Success(c, "查询成功").SetLogTag(e.OperAdd, e.DictDataAdd).WriteHtmlExit("dict_data_add.html", gin.H{"data_type": req.DictType})
}

func (a SysDict) DictDataEditPage(c *gin.Context) {
	req := dto.DictDataIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.DictDataEdit).WriteJsonExit()
		return
	}
	dictSvice := service.SysDict{}
	data, err := dictSvice.GetDictDataByID(&req)
	if err != nil {
		a.Error(c, "查询失败", err).SetLogTag(e.OperEdit, e.DictDataEdit).WriteJsonExit()
		return
	}
	a.Success(c, "查询成功").SetLogTag(e.OperEdit, e.DictDataEdit).WriteHtmlExit("dict_data_edit.html", gin.H{"dictData": data})
}

func (a SysDict) DictDataSave(c *gin.Context) {
	dictSvice := service.SysDict{}
	req := dto.DictDataEditForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperOther, e.DictDataEdit).WriteJsonExit()
		return
	}
	user := a.GetUserFromSession(c)

	if req.Status == "" {
		req.Status = "0"
	} else if req.Status == "on" {
		req.Status = "1"
	}

	if req.ID == "" {
		//新增
		req.SetCreate(user.ID)
		err = dictSvice.InsertDictData(&req)
		if err != nil {
			a.Error(c, "新增失败", err).SetLogTag(e.OperAdd, e.DictDataAdd).WriteJsonExit()
			return
		}
		a.Success(c, "新增成功").SetLogTag(e.OperAdd, e.DictDataAdd).WriteJsonExit()
	} else {
		//更新
		req.SetUpdate(user.ID)
		err = dictSvice.UpdateDictDataAttr(&req)
		if err != nil {
			a.Error(c, "更新失败", err).SetLogTag(e.OperEdit, e.DictDataEdit).WriteJsonExit()
			return
		}
		a.Success(c, "更新成功").SetLogTag(e.OperEdit, e.DictDataEdit).WriteJsonExit()
	}
}

func (a SysDict) DictDataDelete(c *gin.Context) {
	dictSvice := service.SysDict{}
	req := dto.DictDataIDForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperDel, e.DictDataDelete).WriteJsonExit()
		return
	}

	err = dictSvice.DeleteDictData(&req)
	if err != nil {
		a.Error(c, "删除失败", err).SetLogTag(e.OperDel, e.DictDataDelete).WriteJsonExit()
		return
	}
	a.Success(c, "删除成功").SetLogTag(e.OperDel, e.DictDataDelete).WriteJsonExit()
}

func (a SysDict) DictDataStatus(c *gin.Context) {
	dictSvice := service.SysDict{}
	req := dto.DictDataStatusForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.DictDataEdit).WriteJsonExit()
		return
	}

	err = dictSvice.UpdateDictDataStatus(&req)
	if err != nil {
		a.Error(c, "更新失败", err).SetLogTag(e.OperEdit, e.DictDataEdit).WriteJsonExit()
		return
	}
	a.Success(c, "更新成功").SetLogTag(e.OperEdit, e.DictDataEdit).WriteJsonExit()
}
