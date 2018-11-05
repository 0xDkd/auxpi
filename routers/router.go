package routers

import (
	"auxpi/bootstrap"
	"auxpi/controllers"
	"auxpi/controllers/api"
	"auxpi/controllers/upload"
	"github.com/astaxie/beego"
)

var siteConfig = bootstrap.Config()

func init() {
	//注册控制器内的路由
	beego.Include(&controllers.UpLoadController{})
	beego.Include(&page.PagesController{})
	//API 路由单独控制
	if siteConfig.OpenApiUpLoad {
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "post:ApiUpLoadHandle")
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "get,put,patch,delete,options,head:ErrorCapture")
	} else {
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "*:ErrorCapture")
	}
}
