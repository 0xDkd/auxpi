package controllers

import (
	"auxpi/bootstrap"

	"github.com/astaxie/beego"
)

type PagesController struct {
	beego.Controller
}

func (i *PagesController) commonStyle() {
	un, _ := i.GetSecureCookie(bootstrap.SiteConfig.AuxpiSalt, "uname")
	if un != "" {
		i.Data["IsLogin"] = true
		i.Data["UserName"] = un
	}
	i.Data["siteName"] = bootstrap.SiteConfig.SiteName
	i.Data["siteUrl"] = bootstrap.SiteConfig.SiteUrl
	i.Data["siteFooterText"] = bootstrap.SiteConfig.SiteFooter
	i.Data["maxNumber"] = bootstrap.SiteConfig.SiteUploadMaxNumber
	//单位为Mb 5mb ==> 5*1024 kb
	i.Data["maxPicSize"] = bootstrap.SiteConfig.SiteUploadMaxSize << 10
	i.Data["xsrf_token"] = i.XSRFToken()
	i.LayoutSections = make(map[string]string)
	i.LayoutSections["Scripts"] = "webUpload/uploadScript.tpl"
	i.LayoutSections["Header"] = "layouts/header.tpl"
	i.LayoutSections["Footer"] = "layouts/footer.tpl"
	i.LayoutSections["Left"] = "layouts/left.tpl"

	
}

func (i *PagesController) IndexShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "SouGou"
	i.Data["iconStyle"] = "sougou"
	i.Data["iconColor"] = "orange"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) LocalShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Local"
	i.Data["iconStyle"] = "auxly"
	i.Data["iconColor"] = "purple"
	i.Layout = "layouts/app.tpl"
	//检测是否开了本地图床
	if bootstrap.SiteConfig.SiteUploadWay.LocalStore.Open {
		i.TplName = "webUpload/box.tpl"
		return
	}
	i.TplName = "webUpload/ban.tpl"
}

func (i *PagesController) SinaShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Sina"
	i.Data["iconStyle"] = "xinlang"
	i.Data["iconColor"] = "red"
	i.Layout = "layouts/app.tpl"
	//检测是否开启新浪图床
	if bootstrap.SiteConfig.SiteUploadWay.OpenSinaPicStore {
		i.TplName = "webUpload/box.tpl"
		return
	}
	i.TplName = "webUpload/ban.tpl"
}

func (i *PagesController) SmmsShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Smms"
	i.Data["iconStyle"] = "sm"
	i.Data["iconColor"] = "blue"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) CCShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "CC"
	i.Data["iconStyle"] = "cc"
	i.Data["iconColor"] = "green"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) FlickrShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Flickr"
	i.Data["iconStyle"] = "flickr"
	i.Data["iconColor"] = "pink darken-1"
	i.Layout = "layouts/app.tpl"
	//检测是否开启新浪图床
	if bootstrap.SiteConfig.SiteUploadWay.OpenFlickrStore {
		i.TplName = "webUpload/box.tpl"
		return
	}
	i.TplName = "webUpload/ban.tpl"
}

func (i *PagesController) BaiduShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Baidu"
	i.Data["iconStyle"] = "baidu1"
	i.Data["iconColor"] = "blue-300"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) QihooShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Qihoo"
	i.Data["iconStyle"] = "logo-"
	i.Data["iconColor"] = "green"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) NetEasyShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "NetEasy"
	i.Data["iconStyle"] = "wangyi"
	i.Data["iconColor"] = "red"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) JdShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Jd"
	i.Data["iconStyle"] = "jingdong-"
	i.Data["iconColor"] = "orange-800"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) JueJinShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "JueJin"
	i.Data["iconStyle"] = "juejin"
	i.Data["iconColor"] = "blue"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) AliShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Ali"
	i.Data["iconStyle"] = "ali"
	i.Data["iconColor"] = "orange"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
}

func (i *PagesController) OpenShow() {
	i.commonStyle()
	i.Data["apiSelect"] = "Open"
	i.Data["iconStyle"] = "gonggong"
	i.Data["iconColor"] = "red"
	i.Layout = "layouts/app.tpl"
	i.TplName = "webUpload/box.tpl"
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
