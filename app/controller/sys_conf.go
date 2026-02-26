package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/global/e"
)

type SysConf struct {
	baseapi.Api
}

// ----------------------------后台-----------------------------

func (a SysConf) SiteEditPage(c *gin.Context) {
	svice := service.SysConf{}
	site, err := svice.GetSiteConf()
	if err != nil {
		a.Error(c, "获取系统配置失败", err).SetLogTag(e.OperEdit, e.SiteEdit).WriteJsonExit()
		return
	}
	a.Success(c, "操作成功").SetLogTag(e.OperEdit, e.SiteEdit).WriteHtmlExit("site_config.html", gin.H{"site": site})
}

func (a SysConf) SiteEdit(c *gin.Context) {
	svice := service.SysConf{}

	var req dto.SiteConfForm
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).SetLogTag(e.OperEdit, e.SiteEdit).WriteJsonExit()
		return
	}

	if err = svice.EditSiteConf(&req); err != nil {
		a.Error(c, "修改系统配置失败", err).SetLogTag(e.OperEdit, e.SiteEdit).WriteJsonExit()
		return
	}
	a.Success(c, "操作成功").SetLogTag(e.OperEdit, e.SiteEdit).WriteJsonExit()
}

func (a SysConf) BaseEditPage(c *gin.Context) {

}

func (a SysConf) BaseEdit(c *gin.Context) {

}
