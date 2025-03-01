package service

import (
	"encoding/json"
	"gsadmin/app/model"
	"gsadmin/app/service/dto"
	"gsadmin/core/baseservice"
	"gsadmin/core/cache"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/core/utils/assertion"
	"gsadmin/global/e"
	"strings"
)

type SysConf struct {
	baseservice.Service
}

func (s *SysConf) GetBaseConf() (model.Base, error) {
	var base model.Base
	data, err := cache.Instance().Get(e.SysBase, e.BaseConfig)
	if err == nil {
		err = json.Unmarshal([]byte(data), &base)
		if err != nil {
			log.Instance().Error("Unmarshal base conf failed: " + err.Error())
		}
		//d, ok := data.(model.Base)
		//if ok {
		//	base = d
		//}
	} else {
		var baseConf model.SysConf
		err = db.Instance().Model(model.SysConf{}).Where("status = 1 AND type = ?", e.BaseConfig).First(&baseConf).Error
		if err != nil {
			log.Instance().Error("SysConfService.GetBaseConf:" + err.Error())
			return base, err
		}

		err = json.Unmarshal([]byte(baseConf.Info), &base)
		if err != nil {
			log.Instance().Error("SysConfService.GetBaseConf:" + err.Error())
			return base, err
		}
		b, _ := json.Marshal(base)
		_ = cache.Instance().Set(e.SysBase, e.BaseConfig, string(b), e.ConfCacheTime)
	}
	return base, nil
}

func (s *SysConf) GetSiteConf() (model.Site, error) {
	var site model.Site
	data, err := cache.Instance().Get(e.SysBase, e.SiteConfig)
	if err == nil {
		err = json.Unmarshal([]byte(data), &site)
		if err != nil {
			log.Instance().Error("Unmarshal site conf failed: " + err.Error())
		}
	} else {
		var baseConf model.SysConf
		err = db.Instance().Model(model.SysConf{}).Where("status = 1 AND type = ?", e.SiteConfig).First(&baseConf).Error
		if err != nil {
			log.Instance().Error("SysConfService.GetSiteConf:" + err.Error())
			return site, err
		}

		err = json.Unmarshal([]byte(baseConf.Info), &site)
		if err != nil {
			log.Instance().Error("SysConfService.GetSiteConf:" + err.Error())
			return site, err
		}

		b, _ := json.Marshal(site)
		_ = cache.Instance().Set(e.SysBase, e.SiteConfig, string(b), e.ConfCacheTime)
	}
	return site, nil
}

func (s *SysConf) EditSiteConf(req *dto.SiteConfForm) error {
	if strings.HasPrefix(req.WebUrl, "http://") == false { // http前缀校验
		req.WebUrl = "http://" + req.WebUrl
	}
	if strings.HasSuffix(req.WebUrl, "/") == false { // 结尾校验
		req.WebUrl = req.WebUrl + "/"
	}
	if req.WebUrl != "" && strings.HasPrefix(req.LogoUrl, req.WebUrl) == false { // logo地址校验
		req.LogoUrl = req.WebUrl + req.LogoUrl
	}
	str, err := json.Marshal(req)
	if err != nil {
		log.Instance().Error("SysConfService.EditSiteConf:" + err.Error())
		return err
	}
	// 查询是否存在
	var tmp model.SysConf
	_ = db.Instance().Model(model.SysConf{}).Where("type = ?", e.SiteConfig).First(&tmp).Error

	if tmp.ID == 0 { //不存在就新增
		err = db.Instance().Model(model.SysConf{}).
			Where("type = ?", e.SiteConfig).
			FirstOrCreate(&model.SysConf{Info: assertion.AnyToString(str), Status: 1}).
			Error
	} else { //存在就更新
		err = db.Instance().Model(model.SysConf{}).
			Where("type = ?", e.SiteConfig).
			Updates(&model.SysConf{Info: assertion.AnyToString(str), Status: 1}).
			Error
		if err != nil {
			log.Instance().Error("SysConfService.EditSiteConf:" + err.Error())
			return err
		}
	}

	// 修改后刷新一下cache
	var site = model.Site{
		WebName:     req.WebName,
		WebUrl:      req.WebUrl,
		LogoUrl:     req.LogoUrl,
		KeyWords:    req.KeyWords,
		Description: req.Description,
		Copyright:   req.Copyright,
		Icp:         req.Icp,
		SiteStatus:  req.SiteStatus,
	}
	b, _ := json.Marshal(site)
	_ = cache.Instance().Set(e.SysBase, e.SiteConfig, string(b), e.ConfCacheTime)
	return nil
}

func (s *SysConf) EditBaseConf() {

}
