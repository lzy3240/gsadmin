package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gsadmin/app/model"
	"gsadmin/core/baseapi"
	"gsadmin/core/config"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/core/queue"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/ip"
	"gsadmin/core/utils/session"
	"gsadmin/global/e"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// LogTo 日志中间件
func LogTo() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 获取请求参数
		var param string

		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			bf := bytes.NewBuffer(nil)
			wt := bufio.NewWriter(bf)
			_, err := io.Copy(wt, c.Request.Body)
			if err != nil {
				log.Instance().Error("copy request body Failed." + err.Error())
				err = nil
			}
			rb, _ := ioutil.ReadAll(bf)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
			param = string(rb)
			//中文参数处理
			text, err1 := url.QueryUnescape(param)
			if err1 == nil {
				param = text
			} else {
				log.Instance().Error("parse request param failed." + err1.Error())
			}
		}
		// 继续执行
		c.Next()

		// 结束时间
		endTime := time.Now()

		// OPTION不记录
		if c.Request.Method == http.MethodOptions {
			return
		}

		// 处理返回结果
		outBody, _ := c.Get("result")
		respBody := outBody.(*baseapi.CommonResp)

		// Tag为false不记录
		if respBody.Tag == false {
			return
		}

		//获取用户信息
		tmp := session.Get(c, e.UserInfo)

		//类型转换
		user := new(model.SysUser)
		if s, ok := tmp.(string); ok {
			_ = json.Unmarshal([]byte(s), &user)
		}

		var status = 0
		if respBody.Code == 200 || respBody.Code == 302 || respBody.Code == 0 {
			status = 1
		}

		var errMsg string
		if respBody.Code == 500 {
			errMsg = respBody.Msg
		}

		var operIp = ip.GetClientIp(c.Request)

		// 暂不写完,仅写255
		respData := assertion.AnyToString(respBody.Data)
		if len(respData) > 256 {
			respData = respData[0:255]
		}

		var operLog = model.SysOperLog{
			Title:         respBody.Title,
			UserId:        assertion.AnyToInt(user.ID),
			OperName:      user.LoginName,
			DeptName:      "",
			OperParam:     param,
			Status:        status,
			JsonResult:    respData,
			ErrorMsg:      errMsg,
			BusinessType:  respBody.Type,
			RequestMethod: c.Request.Method,
			OperUrl:       c.Request.URL.Path,
			Method:        c.Request.Method,
			OperIp:        operIp,
			OperatorType:  ip.GetDevice(c.Request), // 操作类别（0爬虫 1手机用户 2网页用户）
			OperLocation:  ip.GetCityByIp(operIp),
			UserAgent:     c.Request.UserAgent(),
			LatencyTime:   int(endTime.Sub(startTime).Milliseconds()),
			OperTime:      time.Now().Format(e.TimeFormat),
		}

		// 写入队列, 异步持久化
		msg, err := json.Marshal(operLog)
		if err != nil {
			log.Instance().Error("Marshal OperForm Failed..." + err.Error())
		}
		err = queue.Instance().Publish(e.TopicOperLog, string(msg))
		if err != nil {
			log.Instance().Error("Publish Msg Failed..." + err.Error())
		}
	}
}

// WriteTo 队列日志写入
func WriteTo(tmp string) {
	var sysOperLog model.SysOperLog
	err := json.Unmarshal([]byte(tmp), &sysOperLog)
	if err != nil {
		log.Instance().Error("UnMarshal OperLog Failed..." + err.Error())
	}
	switch config.Instance().ZapLog.SaveMode {
	case "file":
		logToFile(sysOperLog)
	case "db":
		logToDB(sysOperLog)
	default:
		logToFile(sysOperLog)
	}
}

// 操作日志入库
func logToDB(operLog model.SysOperLog) {
	if err := db.Instance().Model(model.SysOperLog{}).Create(&operLog).Error; err != nil {
		log.Instance().Error("Insert OperLog Failed..." + err.Error())
	}
}

// 操作日志入文件
func logToFile(operLog model.SysOperLog) {
	log.Instance().Info("OperLog SuccessResp...", zap.Any("operLog", operLog))
}
