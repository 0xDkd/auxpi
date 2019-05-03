package base

import (
	"github.com/astaxie/beego"
	"github.com/auxpi/bootstrap"
	"github.com/auxpi/controllers"
	"github.com/auxpi/models"
)

type ApiController struct {
	beego.Controller
	controllers.Base
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

//所有的 APi 不需要开启 CSRF
func (a *ApiController) Prepare() {
	a.EnableXSRF = false
}

func (a *ApiController) Migrate() {
	models.CreateAdminRole()
	a.ServeJSON()
}

func (a *ApiController) ShowIt() {
	a.Data["json"] = bootstrap.SiteConfig
	a.ServeJSON()
}
