package ip

import (
	"encoding/json"
	"github.com/axgle/mahonia"
	"github.com/mssola/user_agent"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func GetCityByIp(ip string) string {
	//https://ip9.com.cn/get?ip=60.173.235.154
	//http://ip-api.com/json/60.173.235.154?lang=zh-CN
	//http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=60.173.235.154
	if ip == "" {
		return ""
	}
	if ip == "::1" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		tmp := ConvertToString(bodystr, "gbk", "utf-8")
		p := make(map[string]interface{}, 0)
		if err := json.Unmarshal([]byte(tmp), &p); err == nil {
			var country = ""
			var prov = ""
			var city = ""
			if p["country"] != nil && p["country"].(string) != "" {
				country = p["country"].(string)
			}

			if p["prov"] != nil && p["prov"].(string) != "" {
				prov = p["prov"].(string)
			}

			if p["pro"] != nil && p["pro"].(string) != "" {
				prov = p["pro"].(string)
			}

			if p["regionName"] != nil && p["regionName"].(string) != "" {
				prov = p["regionName"].(string)
			}

			if p["city"] != nil && p["city"].(string) != "" {
				city = p["city"].(string)
			}
			//return p["country"].(string) + "-" + p["prov"].(string) + "-" + p["city"].(string)
			return country + " " + prov + " " + city
		}
	}
	return ""
}

// ConvertToString
// src 字符串
// srcCode 字符串当前编码
// tagCode 要转换的编码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// GetClientIp 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func GetClientIp(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// GetDevice 判断请求类型 0-爬虫 1-手机 2-网页
func GetDevice(r *http.Request) int {
	ua := r.Header.Get("User-Agent")
	parsedUA := user_agent.New(ua)
	if parsedUA.Bot() {
		return 0
	} else if parsedUA.Mobile() {
		return 1
	} else {
		return 2
	}
}
