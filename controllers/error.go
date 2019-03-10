package controllers

import (
	"auxpi/bootstrap"
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
	e.Data["SiteName"] = bootstrap.SiteConfig.SiteName
	e.Data["Time"] = beego.Date(time.Now(), "Y")
	e.Data["SiteLink"] = bootstrap.SiteConfig.SiteUrl
	e.Data["Logo"] = bootstrap.SiteConfig.Logo
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
		Link:          bootstrap.SiteConfig.SiteUrl,
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
		Link:          bootstrap.SiteConfig.SiteUrl,
	}
	e.Data["Part"] = "404"
}
