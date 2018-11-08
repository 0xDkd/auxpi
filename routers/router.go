package routers

import (
	"auxpi/bootstrap"
	"auxpi/controllers"
	"auxpi/controllers/api"
	"github.com/astaxie/beego"
)

var siteConfig = bootstrap.Config()

func init() {
	//正式环境不使用控制器内环境，调试时使用控制器内反射路由
	beego.Router("/",&page.PagesController{},"get:IndexShow")
	beego.Router("/Sina",&page.PagesController{},"get:SinaShow")
	beego.Router("/Smms",&page.PagesController{},"get:SmmsShow")
	beego.Router("/about",&page.PagesController{},"get:AboutShow")
	//注册无需控制的 API
	RegisterOpenApi()
	//部分 API 路由单独控制
	if siteConfig.OpenApiUpLoad {
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "post:UpLoadHandle")
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "get,put,patch,delete,options,head:ErrorCapture")
	} else {
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "*:ErrorCapture")
	}
}
