package page

import (
	"auxpi/bootstrap"
	"github.com/astaxie/beego"
)

type PagesController struct {
	beego.Controller
}

//获取 config 的配置
var siteConfig = bootstrap.Config()

func (c *PagesController) URLMapping() {
	c.Mapping("Show", c.IndexShow)
	c.Mapping("AboutShow", c.AboutShow)
	c.Mapping("SinaShow", c.SinaShow)
	c.Mapping("TestShow", c.TestShow)
}

// @router / [get]
func (this *PagesController) IndexShow() {
	this.Data["siteName"] = siteConfig.SiteName
	this.Data["siteUrl"] = siteConfig.SiteUrl
	this.Data["siteFooterText"] = siteConfig.SiteFooter
	this.Data["apiUrl"] = "/api/v1/auth/upload"
	this.Data["maxNumber"] = siteConfig.SiteUploadMaxNumber
	//单位为Mb 5mb ==> 5*1024 kb
	this.Data["maxPicSize"] = siteConfig.SiteUpLoadMaxSize << 10
	this.Data["apiSelect"] = "SouGou"
	this.Data["iconStyle"] ="sougou"
	this.Data["iconColor"] ="orange"
	this.Data["xsrf_token"] = this.XSRFToken()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "upload/uploadScript.tpl"
	this.LayoutSections["Header"] = "layouts/header.tpl"
	this.LayoutSections["Footer"] = "layouts/footer.tpl"
	this.LayoutSections["Left"] = "layouts/left.tpl"
	this.Layout = "layouts/app.tpl"
	this.TplName = "upload/box.tpl"
}

// @router /Sina [get]
func (this *PagesController) SinaShow() {
	this.Data["siteName"] = siteConfig.SiteName
	this.Data["siteUrl"] = siteConfig.SiteUrl
	this.Data["siteFooterText"] = siteConfig.SiteFooter
	this.Data["apiUrl"] = "/api/v1/auth/upload"
	this.Data["maxNumber"] = siteConfig.SiteUploadMaxNumber
	//单位为Mb 5mb ==> 5*1024 kb
	this.Data["maxPicSize"] = siteConfig.SiteUpLoadMaxSize << 10
	this.Data["xsrf_token"] = this.XSRFToken()
	this.LayoutSections = make(map[string]string)
	this.Data["apiSelect"] = "Sina"
	this.Data["iconStyle"] ="xinlang"
	this.Data["iconColor"] ="red"
	this.LayoutSections["Scripts"] = "upload/uploadScript.tpl"
	this.LayoutSections["Header"] = "layouts/header.tpl"
	this.LayoutSections["Footer"] = "layouts/footer.tpl"
	this.LayoutSections["Left"] = "layouts/left.tpl"
	this.Layout = "layouts/app.tpl"
	//检测是否开启新浪图床
	if siteConfig.SiteUploadWay.OpenSinaPicStore {
		this.TplName = "upload/box.tpl"
		return
	}
	this.TplName = "upload/ban.tpl"
}

// @router /about/ [get]
func (this *PagesController) AboutShow() {
	this.Data["siteName"] = siteConfig.SiteName
	this.Data["siteUrl"] = siteConfig.SiteUrl
	this.Data["siteFooterText"] = siteConfig.SiteFooter
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "upload/uploadScript.tpl"
	this.LayoutSections["Header"] = "layouts/header.tpl"
	this.LayoutSections["Footer"] = "layouts/footer.tpl"
	this.LayoutSections["Left"] = "layouts/left.tpl"
	this.Data["title"] = "关于 Buster API 图床"
	this.Layout = "layouts/app.tpl"
	this.TplName = "about/about-me.tpl"
}

// @router /app [get]
func (this *PagesController) InstallShow() {
	this.Data["siteName"] = siteConfig.SiteName
	this.Data["siteUrl"] = siteConfig.SiteUrl
	this.Data["siteFooterText"] = siteConfig.SiteFooter
	this.TplName = "about-me.tpl"
}

// @router /test [get]
func (this *PagesController) TestShow() {
	this.Data["siteName"] = siteConfig.SiteName
	this.Data["siteUrl"] = siteConfig.SiteUrl
	this.Data["siteFooterText"] = siteConfig.SiteFooter
	this.Data["apiUrl"] = "/api/v1/auth/upload"
	this.Data["maxNumber"] = siteConfig.SiteUploadMaxNumber
	//单位为Mb 5mb ==> 5*1024 kb
	this.Data["maxPicSize"] = siteConfig.SiteUpLoadMaxSize << 10
	this.Data["xsrf_token"] = this.XSRFToken()
	this.LayoutSections = make(map[string]string)
	this.Data["apiSelect"] = "Sina"
	this.Data["iconStyle"] ="xinlang"
	this.Data["iconColor"] ="red"
	//this.LayoutSections["Scripts"] = "upload/uploadScript.tpl"
	this.LayoutSections["Header"] = "layouts/header.tpl"
	this.LayoutSections["Footer"] = "layouts/footer.tpl"
	this.LayoutSections["Left"] = "layouts/left.tpl"
	this.Layout = "layouts/app.tpl"
	this.TplName = "upload/ban.tpl"

}
