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
	a.MountCtx(c)
	svice := service.SysConf{}
	site, err := svice.GetSiteConf()
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.SiteEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperEdit, e.SiteEdit).WriteHtmlExit("site_config.html", gin.H{"site": site})
}

func (a SysConf) SiteEdit(c *gin.Context) {
	a.MountCtx(c)
	svice := service.SysConf{}

	var req dto.SiteConfForm
	err := a.Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.SiteEdit).WriteJsonExit()
		return
	}

	if err = svice.EditSiteConf(&req); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperEdit, e.SiteEdit).WriteJsonExit()
		return
	}
	a.SuccessResp().SetLogTag(e.OperEdit, e.SiteEdit).WriteJsonExit()
}

func (a SysConf) BaseEditPage(c *gin.Context) {

}

func (a SysConf) BaseEdit(c *gin.Context) {

}
