package cron

import (
	"errors"
	"regexp"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/gocolly/colly"
)

type Cron struct{}

type Charset string

const (
	//charset
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	UNKNOWN = Charset("UNKNOWN")
	// hostset
	CNBLOGS = Charset("CNBLOGS")
	CSDN    = Charset("CSDN")
	JB51    = Charset("JB51")
	CTO51   = Charset("CTO51")
	JIANSHU = Charset("JIANSHU")
	ZHIHU   = Charset("ZHIHU")
	WEIXIN  = Charset("WEIXIN")
)

var HostMap = map[Charset]string{
	CNBLOGS: "cnblogs.com",
	CSDN:    "blog.csdn.net",
	JB51:    "jb51.net",
	CTO51:   "51cto.com",
	JIANSHU: "jianshu.com",
	ZHIHU:   "zhihu.com",
	WEIXIN:  "mp.weixin.qq.com",
}

// Get 访问URL
func (cron Cron) Get(url, host string) (title string, body string, err error) {
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188"))
	switch getHostCode(strings.Split(host, ":")[0]) {
	case CNBLOGS:
		//c := colly.NewCollector()
		c.OnHTML("#cb_post_title_url", func(e *colly.HTMLElement) {
			title = e.Text
		})
		c.OnHTML("#cnblogs_post_body", func(e *colly.HTMLElement) {
			body, _ = e.DOM.Html()
		})
		c.Visit(url)
	case CSDN:
		//c := colly.NewCollector()
		c.OnHTML(".title-article", func(e *colly.HTMLElement) {
			title = e.Text
		})
		c.OnHTML("#content_views", func(e *colly.HTMLElement) {
			body, _ = e.DOM.Html()
		})
		c.Visit(url)
	case JB51:
		//c := colly.NewCollector()
		c.OnHTML("#article", func(e *colly.HTMLElement) {
			t := e.DOM.Children().Filter(".title").Text()
			switch getStrCode([]byte(t)) {
			case "UTF-8":
				t = strings.Replace(t, "&#34;", "\"", -1)
				t = strings.Replace(t, "&lt;", "<", -1)
				t = strings.Replace(t, "&gt;", ">", -1)

			case "GB18030":
				t = strings.Replace(convertToString(t, "gbk", "utf-8"), "聽", " ", -1)
				t = strings.Replace(t, "&#34;", "\"", -1)
				t = strings.Replace(t, "&gt;", ";", -1)

			default:
			}
			title = t
		})
		c.OnHTML("#content", func(e *colly.HTMLElement) {
			e.DOM.Children().Filter(".art_xg").Remove()
			e.DOM.Children().Filter(".art_xg ul").Remove()
			e.DOM.Children().Filter(".tip").Remove()
			e.DOM.Children().Filter(".xgcomm clearfix").Remove()

			s, _ := e.DOM.Html()
			switch getStrCode([]byte(s)) {
			case "UTF-8":
				s = strings.Replace(s, "&#34;", "\"", -1)
				s = strings.Replace(s, "&lt;", "<", -1)
				s = strings.Replace(s, "&gt;", ">", -1)
				s = strings.Replace(s, "&amp;", "&", -1)
				s = strings.Replace(s, "&#39;", "'", -1)

			case "GB18030":
				s = strings.Replace(convertToString(s, "gbk", "utf-8"), "聽", " ", -1)
				s = strings.Replace(s, "&#34;", "\"", -1)
				s = strings.Replace(s, "&gt;", ";", -1)
				s = strings.Replace(s, "&amp;", "&", -1)
				s = strings.Replace(s, "&#39;", "'", -1)

			default:
			}
			body = s
		})
		c.Visit(url)
	case ZHIHU:
		//c := colly.NewCollector()
		c.OnHTML(".Post-Header", func(e *colly.HTMLElement) {
			title = e.DOM.Children().Filter(".Post-Title").Text()
		})
		if strings.Contains(url, "zhuanlan") {
			c.OnHTML(".Post-RichTextContainer", func(e *colly.HTMLElement) {
				//e.DOM.Children().Filter(".ztext-empty-paragraph").Remove()
				s, _ := e.DOM.Html()
				body = s[strings.Index(s, "[object Object]")+17:]
				re, _ := regexp.Compile("( data-(pid|size|caption|rawheight|original|actualsrc|rawwidth)=\"[a-zA-Z0-9-_]*\"| class=\"[a-zA-Z0-9-_/ ]*\"|<noscript>[<a-zA-z0-9.\\-/:=\">? ]*</noscript>|<p></p>)")
				re1, _ := regexp.Compile("(<figure>|</figure>|data-actualsrc=\"[\\S\\s]+?\"|data:image[\\S\\s]+?data-original=\"|itemprop=\"text\">|\\?source=[a-z0-9]{8})") //itemprop="text">|\?source=[a-z0-9]{8} queston类型内容

				body = re.ReplaceAllString(body, "")
				body = re1.ReplaceAllString(body, "")
			})
		} else if strings.Contains(url, "question") {
			c.OnHTML(".RichContent-inner", func(e *colly.HTMLElement) {
				//e.DOM.Children().Filter(".ztext-empty-paragraph").Remove()
				s, _ := e.DOM.Html()
				body = s[strings.Index(s, "[object Object]")+17:]
				re, _ := regexp.Compile("( data-(pid|size|caption|rawheight|original|actualsrc|rawwidth)=\"[a-zA-Z0-9-_]*\"| class=\"[a-zA-Z0-9-_/ ]*\"|<noscript>[<a-zA-z0-9.\\-/:=\">? ]*</noscript>|<p></p>)")
				re1, _ := regexp.Compile("(<figure>|</figure>|data-actualsrc=\"[\\S\\s]+?\"|data:image[\\S\\s]+?data-original=\"|itemprop=\"text\">|\\?source=[a-z0-9]{8})") //itemprop="text">|\?source=[a-z0-9]{8} queston类型内容

				body = re.ReplaceAllString(body, "")
				body = re1.ReplaceAllString(body, "")
			})
		} else {
			return "", "", errors.New("type not zhuanlan or question")
		}

		c.Visit(url)
	case CTO51:
		//c := colly.NewCollector()
		c.OnHTML("article", func(e *colly.HTMLElement) {
			//e.DOM.Children().Filter(".original").Remove()
			title = e.DOM.Children().Filter(".title").Text()
		})
		c.OnHTML(".article-content-wrap", func(e *colly.HTMLElement) {
			s, _ := e.DOM.Html()
			body = s
		})
		c.Visit(url)
	case JIANSHU:
		//c := colly.NewCollector()
		c.OnHTML("._1RuRku", func(e *colly.HTMLElement) {
			title = e.Text
		})
		c.OnHTML("._2rhmJa", func(e *colly.HTMLElement) {
			body, _ = e.DOM.Html()
		})
		c.Visit(url)
	case WEIXIN:
		c.OnHTML(".rich_media_title", func(e *colly.HTMLElement) {
			title = e.Text
		})
		c.OnHTML(".rich_media_content", func(e *colly.HTMLElement) {
			body, _ = e.DOM.Html()
			re, _ := regexp.Compile("data-src=")
			body = re.ReplaceAllString(body, "src=")
		})
		c.Visit(url)
	case UNKNOWN:
		return "", "", errors.New("host not in list")
	default:
		return "", "", errors.New("host not in list")
	}
	return
}

// gbk2utf 编码转换
func convertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// 判断内容编码格式
func getStrCode(data []byte) Charset {
	if isUtf8(data) == true {
		return UTF8
	} else if isGBK(data) == true {
		return GB18030
	} else {
		return UNKNOWN
	}
}

// 判断是否为中文编码
func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// 判断是否为utf-8编码
func isUtf8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

// preNum
func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}

// 判断Host编码
func getHostCode(host string) Charset {
	for k, v := range HostMap {
		if strings.Contains(host, v) {
			return k
		}
	}
	return UNKNOWN
}
