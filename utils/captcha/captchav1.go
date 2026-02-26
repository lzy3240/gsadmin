package captcha

import (
	"github.com/mojocn/base64Captcha"
	"gsadmin/core/config"
	"image/color"
)

// 设置自带的store
var store = base64Captcha.DefaultMemStore

// CaptMake 生成验证码
func CaptMake() (id, b64s string, err error) {
	driver := mathCaptcha()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, _, lerr := captcha.Generate()
	return lid, lb64s, lerr
}

// 数字运算验证码
func mathCaptcha() base64Captcha.Driver {
	return base64Captcha.NewDriverMath(
		config.Instance().CaptChar.ImgHeight,
		config.Instance().CaptChar.ImgWidth,
		//e.ImgHeight,
		//e.ImgWidth,
		0,
		0,
		&color.RGBA{0, 0, 0, 0},
		nil,
		[]string{"RitaSmith.ttf"},
	)
}

// 字符验证码
func stringCaptcha() base64Captcha.Driver {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	captchaConfig := base64Captcha.DriverString{
		Height: config.Instance().CaptChar.ImgHeight,
		Width:  config.Instance().CaptChar.ImgWidth,
		Length: config.Instance().CaptChar.ImgKeyLength,
		//Height:          e.ImgHeight,
		//Width:           e.ImgWidth,
		//Length:          e.ImgKeyLength,
		NoiseCount:      0,     // 干扰字母
		ShowLineOptions: 1 | 3, // 干扰线
		Source:          "qwertyuioplkjhgfdsazxcvbnm",
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	return driver
}

// CaptVerify 验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
