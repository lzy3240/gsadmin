package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mssola/user_agent"
	"gsadmin/app/model"
	"gsadmin/app/service"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseapi"
	"gsadmin/core/baseapi/response"
	"gsadmin/core/config"
	"gsadmin/core/store"
	"gsadmin/global/e"
	"gsadmin/utils/assertion"
	"gsadmin/utils/captcha"
	f "gsadmin/utils/file"
	"gsadmin/utils/ip"
	"gsadmin/utils/str"
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

// ----------------layout------------------

//-----------------------------前台----------------------------

func (a SysBase) Welcome(c *gin.Context) {
	//获取ip
	cip := ip.GetClientIp(c.Request)
	//获取归属地
	addr := ip.GetCityByIp(cip)
	//获取site conf
	svice := service.SysConf{}
	site, err := svice.GetSiteConf()
	if err != nil {
		a.Error(c, "获取系统配置失败", err).WriteJsonExit()
		return
	}

	a.Success(c, "操作成功").SetData(gin.H{"site": site, "local": gin.H{"ip": cip, "addr": addr}}).WriteJsonExit()
}

// ----------------system-------------------

// Base 基础样式
func (a SysBase) Base(c *gin.Context) {
	baseConf := service.SysConf{}
	data, err := baseConf.GetBaseConf()
	if err != nil {
		a.Error(c, "获取基础配置失败", err).WriteJsonExit()
		return
	}

	a.Success(c, "操作成功").WriteCustomJsonExit(data)
}

// Menu 基础菜单
func (a SysBase) Menu(c *gin.Context) {
	user := a.GetUserFromSession(c)
	menu := service.SysBase{}
	menuResp := menu.MenuServiceV2(user)
	a.Success(c, "操作成功").WriteCustomJsonExit(menuResp.MenuResp)
}

// UploadFile 文件上传页面
func (a SysBase) UploadFile(c *gin.Context) {
	a.Success(c, "操作成功").WriteHtmlExit("upload_file.html", nil)
}

// Upload 文件上传
func (a SysBase) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		a.Error(c, "文件上传失败", err).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}
	if file.Size > e.DefUploadSize {
		a.Error(c, "文件大小超限", nil).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}
	//处理文件
	day := time.Now().Format(e.TimeFormatDay)
	savePath := filepath.Join(config.Instance().App.FileSavePath, day) // 按年月日归档保存
	err = f.IsNotExistMkDir(savePath)
	if err != nil {
		a.Error(c, "文件路径失败", err).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}
	//保存文件, 多种存储方式
	//先保存到本地
	saveName := filepath.Join(savePath, file.Filename)
	if err = c.SaveUploadedFile(file, saveName); err != nil {
		a.Error(c, "文件保存失败", err).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}

	//再根据存储配置处理文件
	backFilePath, err1 := store.Instance().UploadFile(file.Filename, saveName)
	if err1 != nil {
		a.Error(c, "文件处理失败", err1).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
		return
	}

	a.Success(c, "操作成功").SetData(backFilePath).SetLogTag(e.OperAdd, e.DefaultUpload).WriteJsonExit()
}

// IndexPage 主页
func (a SysBase) IndexPage(c *gin.Context) {
	user := a.GetUserFromSession(c)

	sysConf := service.SysConf{}
	site, _ := sysConf.GetSiteConf()

	a.Success(c, "操作成功").WriteHtmlExit("index.html", gin.H{
		"site":      site,
		"user":      user,
		"copyright": template.HTML(site.Copyright), // 防止转义
	})
}

// MainPage 引用主页
func (a SysBase) MainPage(c *gin.Context) {
	a.Success(c, "操作成功").WriteHtmlExit("main.html", gin.H{})
}

// ----------------common-------------------

func (a SysBase) ServerErrPage(c *gin.Context) {
	a.Success(c, "服务器错误").WriteHtmlExit("server_err.html", nil)
}

func (a SysBase) IconShow(c *gin.Context) {
	a.Success(c, "操作成功").WriteHtmlExit("icon.html", nil)
}

func (a SysBase) ShowFile(c *gin.Context) {
	fp := c.Query("filePath")
	fi, err := os.Open(fp)
	if err != nil {
		a.Error(c, "文件不存在", err).WriteRedirect("/not_found")
		return
	}
	b, err := ioutil.ReadAll(fi)
	if err != nil {
		a.Error(c, "文件读取失败", err).WriteRedirect("/not_found")
		return
	}
	a.Success(c, "操作成功").WriteStringExit("%s", string(b))
}

func (a SysBase) GetCaptcha(c *gin.Context) {
	id, b64s, err := captcha.CaptMake()
	if err != nil {
		a.Error(c, "验证码生成失败", err).WriteJsonExit()
		return
	}
	a.Success(c, "操作成功").SetData(response.CaptchaResponse{CaptchaId: id, PicPath: b64s}).WriteJsonExit()
}

func (a SysBase) CaptchaVerify(c *gin.Context) {
	req := dto.CaptchaVerifyForm{}
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
		return
	}

	if captcha.CaptVerify(req.ID, strings.ToLower(req.Capt)) == true {
		a.Success(c, "验证码正确").WriteJsonExit()
	} else {
		a.Error(c, "验证码错误", nil).WriteJsonExit()
	}
	return
}

func (a SysBase) LoginPage(c *gin.Context) {
	a.Success(c, "操作成功").WriteHtmlExit("login.html", gin.H{})
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
	err := a.Bind(c, &req, binding.Form)
	if err != nil {
		a.Error(c, "参数校验失败", err).WriteJsonExit()
		return
	}

	isLock := base.CheckLock(req.UserName)
	if isLock {
		a.Error(c, "账号已锁定, 请稍后再试", nil).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
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
			a.Error(c, "登录日志写入失败", err).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
		if having > 0 {
			a.Error(c, "密码错误,剩余错误次数为:"+assertion.AnyToString(having), nil).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		} else {
			a.Error(c, "账户已锁定, 请稍后再试", nil).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
	} else {
		userID, err := a.SetUserToSession(c, user)
		if err != nil {
			a.Error(c, "用户会话处理失败", err).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
		var online model.SysUserOnline
		err = str.CopyFields(&online, info)
		if err != nil {
			a.Error(c, "用户信息处理失败", err).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
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
			a.Error(c, "用户登录日志失败", err).SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
			return
		}
		a.Success(c, "登陆成功").SetLogTag(e.OperOther, e.UserLogin).WriteJsonExit()
	}
}

func (a SysBase) Logout(c *gin.Context) {
	userService := service.SysUser{}
	uid := a.GetUidFromSession(c) //确定已登录
	if uid != 0 {
		//删除online表
		operUser := a.GetUserFromSession(c)
		_ = userService.SignOut(operUser)

		//删除session
		_ = a.DelUserFromSession(c, *operUser)
	}
	a.Success(c, "退出成功").WriteRedirect("/login")
}

func (a SysBase) NotFoundPage(c *gin.Context) {
	a.Success(c, "页面不存在").WriteHtmlExit("not_found.html", gin.H{})
}
