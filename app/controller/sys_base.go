package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mssola/user_agent"
	"gsadmin/app/model"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/core/config"
	"gsadmin/core/store"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/captcha"
	f "gsadmin/core/utils/file"
	"gsadmin/core/utils/ip"
	"gsadmin/core/utils/str"
	"gsadmin/global/e"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SysBase struct {
	baseapi.Api
}

// ----------------system-------------------

// Base 基础样式
func (a SysBase) Base(c *gin.Context) {
	a.MountCtx(c)
	baseConf := service.SysConf{}
	data, err := baseConf.GetBaseConf()
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}

	a.SuccessResp().WriteCustomJsonExit(data)
}

// Menu 基础菜单
func (a SysBase) Menu(c *gin.Context) {
	a.MountCtx(c)
	user := a.GetUserFromSession()
	menu := service.SysBase{}
	menuResp := menu.MenuServiceV2(user)
	a.SuccessResp().WriteCustomJsonExit(menuResp.MenuResp)
}

// Upload 文件上传
func (a SysBase) Upload(c *gin.Context) {
	a.MountCtx(c)
	file, err := c.FormFile("file")
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}
	if file.Size > e.DefUploadSize {
		a.ErrorResp().SetMsg("文件大小超限").SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}
	//处理文件
	day := time.Now().Format(e.TimeFormatDay)
	savePath := filepath.Join(config.Instance().App.FileSavePath, day) // 按年月日归档保存
	err = f.IsNotExistMkDir(savePath)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}
	//保存文件, 多种存储方式
	//先保存到本地
	saveName := filepath.Join(savePath, file.Filename)
	if err = c.SaveUploadedFile(file, saveName); err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}

	//再根据存储配置处理文件
	backFilePath, err1 := store.Instance().UploadFile(file.Filename, saveName)
	if err1 != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}

	a.SuccessResp().SetData(backFilePath).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
}

// IndexPage 主页
func (a SysBase) IndexPage(c *gin.Context) {
	a.MountCtx(c)
	user := a.GetUserFromSession()

	sysConf := service.SysConf{}
	site, _ := sysConf.GetSiteConf()

	a.SuccessResp().WriteHtmlExit("index.html", gin.H{
		"site":      site,
		"user":      user,
		"copyright": template.HTML(site.Copyright), // 防止转义
	})
}

// FramePage 引用主页
func (a SysBase) FramePage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("main.html", gin.H{})
}

// ----------------common-------------------

func (a SysBase) ServerErrPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("server_err.html", nil)
}

func (a SysBase) IconShow(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("icon.html", nil)
}

func (a SysBase) ShowFile(c *gin.Context) {
	a.MountCtx(c)
	fp := c.Query("filePath")
	fi, err := os.Open(fp)
	if err != nil {
		a.ErrorResp().WriteRedirect("/not_found")
		return
	}
	b, err := ioutil.ReadAll(fi)
	if err != nil {
		a.ErrorResp().WriteRedirect("/not_found")
		return
	}
	a.SuccessResp().WriteStringExit("%s", string(b))
}

func (a SysBase) GetCaptcha(c *gin.Context) {
	a.MountCtx(c)
	id, b64s, err := captcha.CaptMake()
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}
	a.SuccessResp().SetData(baseapi.CaptchaResponse{CaptchaId: id, PicPath: b64s}).WriteJsonExit()
}

func (a SysBase) CaptchaVerify(c *gin.Context) {
	a.MountCtx(c)
	req := dto.CaptchaVerifyForm{}
	err := a.Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg("请填写完整信息").WriteJsonExit()
		return
	}

	if captcha.CaptVerify(req.ID, strings.ToLower(req.Capt)) == true {
		a.SuccessResp().WriteJsonExit()
	} else {
		a.ErrorResp().SetMsg("验证码有误，请重新输入").WriteJsonExit()
	}
	return
}

func (a SysBase) LoginPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("login.html", gin.H{})
}

func (a SysBase) LoginHandler(c *gin.Context) {
	//日志服务
	loginLog := service.SysLog{}
	//基础服务
	base := service.SysBase{}
	//在线用户服务
	onlineService := service.SysUserOnline{}
	//用户服务
	userService := service.SysUser{}

	var req dto.LoginForm
	err := a.MountCtx(c).Bind(&req, binding.Form)
	if err != nil {
		a.ErrorResp().SetMsg(a.TransErr(err)).WriteJsonExit()
		return
	}

	isLock := base.CheckLock(req.UserName)
	if isLock {
		a.ErrorResp().SetMsg("密码错误次数超限，账号已锁定,请稍后再试").SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
		return
	}
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	ub, _ := ua.Browser()

	var info model.SysLoginLog
	info.LoginName = req.UserName
	info.IpAddr = ip.GetClientIp(c.Request)
	info.Os = ua.OS()
	info.Browser = ub
	info.LoginTime = time.Now()
	info.LoginLocation = ip.GetCityByIp(ip.GetClientIp(c.Request))
	clientIp := c.ClientIP()

	user, err := userService.SignIn(req.UserName, req.Password, clientIp)
	if err != nil {
		errNums := base.SetPwdErrNum(req.UserName)
		having := e.UserErrTimes - errNums
		info.Msg = "账号或密码错误"
		info.Status = "0"

		//写入登录日志
		err = loginLog.InsertSysLoginLog(info)
		if err != nil {
			a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
		if having > 0 {
			a.ErrorResp().SetMsg("账号或密码不正确,还有"+assertion.AnyToString(having)+"次之后账号将锁定").SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		} else {
			a.ErrorResp().SetMsg("密码错误次数超限，账号已锁定,请稍后再试").SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
	} else {
		userID, err := a.SetUserToSession(user)
		if err != nil {
			a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
		var online model.SysUserOnline
		err = str.CopyFields(&online, info)
		if err != nil {
			a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
		online.SessionID = userID
		online.Status = "on_line"
		online.ExpireTime = 1440
		online.StartTimestamp = time.Now().Unix()
		online.LastAccessTime = time.Now().Format(e.TimeFormat)
		_ = onlineService.Delete(userID)
		_ = onlineService.Insert(online)
		base.RemovePwdErrNum(req.UserName)

		info.Msg = "登陆成功"
		info.Status = "1"

		//写入登录日志
		err = loginLog.InsertSysLoginLog(info)
		if err != nil {
			a.ErrorResp().SetMsg(a.TransErr(err)).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
		a.SuccessResp().SetMsg("登陆成功").SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
	}
}

func (a SysBase) Logout(c *gin.Context) {
	userService := service.SysUser{}
	a.MountCtx(c)
	uid := a.GetUidFromSession() //确定已登录
	if uid != 0 {
		//删除online表
		operUser := a.GetUserFromSession()
		_ = userService.SignOut(operUser)

		//删除session
		_ = a.DelUserFromSession(*operUser)
	}
	a.SuccessResp().WriteRedirect("/login")
}

func (a SysBase) NotFoundPage(c *gin.Context) {
	a.MountCtx(c).SuccessResp().WriteHtmlExit("not_found.html", gin.H{})
}
