package captcha

//func Captcha(c *gin.Context) {
//	l := captcha.DefaultLen
//	w, h := e.ImgWidth, e.ImgHeight
//	captchaId := captcha.NewLen(l)
//	session := sessions.Default(c)
//	session.Set("captcha", captchaId)
//	_ = session.Save()
//	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
//}

//func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
//	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
//	w.Header().Set("Pragma", "no-cache")
//	w.Header().Set("Expires", "0")
//
//	var content bytes.Buffer
//	switch ext {
//	case ".png":
//		w.Header().Set("Content-Type", "image/png")
//		_ = captcha.WriteImage(&content, id, width, height)
//	case ".wav":
//		w.Header().Set("Content-Type", "audio/x-wav")
//		_ = captcha.WriteAudio(&content, id, lang)
//	default:
//		return captcha.ErrNotFound
//	}
//
//	if download {
//		w.Header().Set("Content-Type", "application/octet-stream")
//	}
//	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
//	return nil
//}

//func CaptchaVerify(c *gin.Context) {
//	code := c.PostForm("code")
//	session := sessions.Default(c)
//	if captchaId := session.Get("captcha"); captchaId != nil {
//		session.Delete("captcha")
//		_ = session.Save()
//		if captcha.VerifyString(captchaId.(string), code) {
//			resp.SuccessResp(c).WriteJsonExit()
//		} else {
//			resp.ErrorResp(c).WriteJsonExit()
//		}
//	} else {
//		resp.ErrorResp(c).WriteJsonExit()
//	}
//}
