// Copyright (c) 2019 aimerforreimu. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//  GNU GENERAL PUBLIC LICENSE
//                        Version 3, 29 June 2007
//
//  Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
//  Everyone is permitted to copy and distribute verbatim copies
// of this license document, but changing it is not allowed.
//
// repo: https://github.com/aimerforreimu/auxpi

package controllers

import (
	"github.com/astaxie/beego"
	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/bootstrap"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
)

type PagesController struct {
	beego.Controller
}


func (i *PagesController) commonStyle() {
	un, _ := i.GetSecureCookie(bootstrap.SiteConfig.AuxpiSalt, "uname")

	var site = auxpi.SiteBase{}
	err := site.UnmarshalJSON([]byte(models.GetOption("site_base", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("CONTROLLER", err)
	}

	Stores := models.GetActiveStore()
	i.Data["Stores"] = Stores
	beego.Alert(Stores[0].Rank)
	if un != "" {
		i.Data["IsLogin"] = true
		i.Data["UserName"] = un
	}

	i.Data["siteName"] = site.SiteName
	i.Data["siteUrl"] = site.SiteUrl
	i.Data["siteFooterText"] = site.SiteFooter
	i.Data["maxNumber"] = site.SiteUploadMaxNumber
	//单位为Mb 5mb ==> 5*1024 kb
	i.Data["maxPicSize"] = site.SiteUploadMaxSize << 10
	i.Data["xsrf_token"] = i.XSRFToken()
	i.LayoutSections = make(map[string]string)
	i.LayoutSections["Scripts"] = "webUpload/uploadScript.tpl"
	i.LayoutSections["Header"] = "layouts/header.tpl"
	i.LayoutSections["Footer"] = "layouts/footer.tpl"
	i.LayoutSections["Left"] = "layouts/left.tpl"
}

func (i *PagesController) Show() {
	router := i.Ctx.Input.Param(":router")
	Store := models.GetStoreInfoByRouter(router)
	if Store.ID == 0 {
		i.Abort("404")
	}
	i.commonStyle()
	i.Data["apiSelect"] = Store.Api
	i.Data["iconStyle"] = Store.Icon
	i.Data["iconColor"] = Store.Color
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) LocalShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Local"
	i.Data["iconStyle"] = "auxpi"
	i.Data["iconColor"] = "purple"
	i.Layout = "layouts/app.tpl"
	var local = auxpi.LocalStore{}
	err := local.UnmarshalJSON([]byte(models.GetOption("local", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("CONTROLLER", err)
	}

	//检测是否开了本地图床
	if local.Status {
		i.TplName = "webUpload/box.tpl"
		return
	}
	i.TplName = "webUpload/ban.tpl"
}

func (i *PagesController) AboutShow() {
	i.commonStyle()
	i.Data["siteName"] = bootstrap.SiteConfig.SiteName
	i.Data["siteUrl"] = bootstrap.SiteConfig.SiteUrl
	i.Data["siteFooterText"] = bootstrap.SiteConfig.SiteFooter
	i.Data["title"] = "关于" + bootstrap.SiteConfig.SiteName
	i.Layout = "layouts/app.tpl"
	i.TplName = "about/about-me.tpl"
}
