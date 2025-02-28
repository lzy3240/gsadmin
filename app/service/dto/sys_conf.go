package dto

type SiteConfForm struct {
	WebName     string `json:"web_name" form:"web_name"`
	WebUrl      string `json:"web_url" form:"web_url"`
	LogoUrl     string `json:"logo_url" form:"logo_url"`
	KeyWords    string `json:"key_words" form:"key_words"`
	Description string `json:"description" form:"description"`
	Copyright   string `json:"copyright" form:"copyright"`
	Icp         string `json:"icp" form:"icp"`
	SiteStatus  uint8  `json:"site_status" form:"site_status"`
}
