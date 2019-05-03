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
	"time"

	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (e *ErrorController) commonStyle() {
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["Header"] = "auth/header.tpl"
	e.LayoutSections["Footer"] = "auth/footer.tpl"
	e.Data["xsrf_token"] = e.XSRFToken()
	e.Layout = "auth/base.tpl"
	e.TplName = "auth/base.tpl"
	e.Data["SiteName"] = site.SiteName
	e.Data["Time"] = beego.Date(time.Now(), "Y")
	e.Data["SiteLink"] = site.SiteUrl
	e.Data["Logo"] = site.Logo
}

type UserMsg struct {
	AlertType     string
	AlertContent  string
	ButtonType    string
	ButtonContent string
	Link          string
}

func (e *ErrorController) Error404() {
	e.commonStyle()
	e.LayoutSections["Content"] = "auth/msg.tpl"
	e.Data["action"] = "register"
	e.Data["Msg"] = &UserMsg{
		AlertType:     "warning",
		AlertContent:  "404 NOT FOUND",
		ButtonType:    "primary",
		ButtonContent: "返回首页",
		Link:          site.SiteUrl,
	}
	e.Data["Part"] = "404"
}

func (e *ErrorController) Error403() {
	e.commonStyle()
	e.LayoutSections["Content"] = "auth/msg.tpl"
	e.Data["action"] = "register"
	e.Data["Msg"] = &UserMsg{
		AlertType:     "danger",
		AlertContent:  "403 Forbidden",
		ButtonType:    "primary",
		ButtonContent: "返回首页",
		Link:          site.SiteUrl,
	}
	e.Data["Part"] = "404"
}

func (e *ErrorController) Error501() {
	e.commonStyle()
	e.LayoutSections["Content"] = "auth/msg.tpl"
	e.Data["action"] = "register"
	e.Data["Msg"] = &UserMsg{
		AlertType:     "warning",
		AlertContent:  "501 SERVER ERROR",
		ButtonType:    "primary",
		ButtonContent: "返回首页",
		Link:          site.SiteUrl,
	}
	e.Data["Part"] = "404"
}
