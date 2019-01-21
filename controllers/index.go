package page

import (
	"auxpi/bootstrap"
	"github.com/astaxie/beego"
)

type PagesController struct {
	beego.Controller
}


func (this *PagesController) commonStyle()  {
	this.Data["siteName"] = bootstrap.SiteConfig.SiteName
	this.Data["siteUrl"] = bootstrap.SiteConfig.SiteUrl
	this.Data["siteFooterText"] = bootstrap.SiteConfig.SiteFooter
	this.Data["maxNumber"] = bootstrap.SiteConfig.SiteUploadMaxNumber
	//单位为Mb 5mb ==> 5*1024 kb
	this.Data["maxPicSize"] = bootstrap.SiteConfig.SiteUpLoadMaxSize << 10
	this.Data["xsrf_token"] = this.XSRFToken()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "webUpload/uploadScript.tpl"
	this.LayoutSections["Header"] = "layouts/header.tpl"
	this.LayoutSections["Footer"] = "layouts/footer.tpl"
	this.LayoutSections["Left"] = "layouts/left.tpl"
}

func (this *PagesController) IndexShow() {
	this.commonStyle()
	this.Data["apiSelect"] = "SouGou"
	this.Data["iconStyle"] = "sougou"
	this.Data["iconColor"] = "orange"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}


func (this *PagesController) SinaShow() {
	this.commonStyle()
	this.Data["apiSelect"] = "Sina"
	this.Data["iconStyle"] = "xinlang"
	this.Data["iconColor"] = "red"
	this.Layout = "layouts/app.tpl"
	//检测是否开启新浪图床
	if bootstrap.SiteConfig.SiteUploadWay.OpenSinaPicStore {
		this.TplName = "webUpload/box.tpl"
		return
	}
	this.TplName = "webUpload/ban.tpl"
}


func (this *PagesController) SmmsShow() {
	this.commonStyle()
	this.Data["apiSelect"] = "Smms"
	this.Data["iconStyle"] = "sm"
	this.Data["iconColor"] = "blue"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) CCShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "CC"
	this.Data["iconStyle"] = "sm"
	this.Data["iconColor"] = "blue"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) FlickrShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "Flickr"
	this.Data["iconStyle"] = "flickr"
	this.Data["iconColor"] = "pink darken-1"
	this.Layout = "layouts/app.tpl"
	//检测是否开启新浪图床
	if bootstrap.SiteConfig.SiteUploadWay.OpenFlickrStore {
		this.TplName = "webUpload/box.tpl"
		return
	}
	this.TplName = "webUpload/ban.tpl"
}

func (this *PagesController) BaiduShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "Baidu"
	this.Data["iconStyle"] = "baidu1"
	this.Data["iconColor"] = "blue-300"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) QihooShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "Qihoo"
	this.Data["iconStyle"] = "logo-"
	this.Data["iconColor"] = "green"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) NetEasyShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "NetEasy"
	this.Data["iconStyle"] = "wangyi"
	this.Data["iconColor"] = "red"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) JdShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "Jd"
	this.Data["iconStyle"] = "jingdong-"
	this.Data["iconColor"] = "orange-800"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) JueJinShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "JueJin"
	this.Data["iconStyle"] = "juejin"
	this.Data["iconColor"] = "blue"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) AliShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "Ali"
	this.Data["iconStyle"] = "ali"
	this.Data["iconColor"] = "orange"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) OpenShow()  {
	this.commonStyle()
	this.Data["apiSelect"] = "Open"
	this.Data["iconStyle"] = "gonggong"
	this.Data["iconColor"] = "red"
	this.Layout = "layouts/app.tpl"
	this.TplName = "webUpload/box.tpl"
}

func (this *PagesController) AboutShow() {
	this.commonStyle()
	this.Data["siteName"] = bootstrap.SiteConfig.SiteName
	this.Data["siteUrl"] = bootstrap.SiteConfig.SiteUrl
	this.Data["siteFooterText"] = bootstrap.SiteConfig.SiteFooter
	this.Data["title"] = "关于"+bootstrap.SiteConfig.SiteName
	this.Layout = "layouts/app.tpl"
	this.TplName = "about/about-me.tpl"
}


