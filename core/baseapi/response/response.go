package response

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}

type CommonResp struct {
	Code  int         `json:"code"`  //响应编码: 200 成功 500 错误 403 无操作权限 401 鉴权失败  -1  失败
	Msg   string      `json:"msg"`   //消息内容
	Data  interface{} `json:"data"`  //数据内容
	Count int         `json:"count"` //数据数量
	Tag   bool        `json:"tag"`   //日志标记
	Type  int         `json:"type"`  //业务类型
	Title string      `json:"title"` //操作名称
}
