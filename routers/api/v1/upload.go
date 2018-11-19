package v1

import (
	"auxpi/bootstrap"
	"auxpi/controllers/api"
	"auxpi/controllers/webUpload"

	"github.com/astaxie/beego"
)

//不需要控制的 API
func RegisterOpenApiV1() {
	//不需要控制的 api 放到这里
	ns :=
		beego.NewNamespace("/api/v1",
			beego.NSRouter("/web_upload/", &controllers.WebUpLoadController{}, "post:UpLoadHandle"),
		)
	beego.AddNamespace(ns)
}


//需要控制的 API
func RegisterControlApiV1()  {
	if bootstrap.SiteConfig.OpenApiUpLoad {
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "post:UpLoadHandle")
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "get,put,patch,delete,options,head:ErrorCapture")
	} else {
		beego.Router("api/v1/upload", &api.ApiUpLoadController{}, "*:ErrorCapture")
	}
}
