package v1Router

import (
	"auxpi/bootstrap"
	"auxpi/controllers"
	"auxpi/controllers/api/v1"
	"auxpi/middleware"

	"github.com/astaxie/beego"
)

//上传中间件
func RegisterUploadMiddleWare()  {
	beego.InsertFilter("/api/v1/web_upload/",beego.BeforeExec,middleware.CookieUploadControl)
}


//不需要控制的 API
func RegisterOpenApi() {
	//不需要控制的 api 放到这里
	ns :=
		beego.NewNamespace("/api/v1",
			beego.NSRouter("/web_upload/", &controllers.WebUpLoadController{}, "post:UpLoadHandle"),
		)
	beego.AddNamespace(ns)
}

//需要控制的 API
func RegisterControlApi() {
	if bootstrap.SiteConfig.OpenApiUpLoad {
		beego.Router("api/v1/upload", &v1.ApiUpLoadController{}, "post:UpLoadHandle")
		beego.Router("api/v1/upload", &v1.ApiUpLoadController{}, "get,put,patch,delete,options,head:ErrorCapture")
	} else {
		beego.Router("api/v1/upload", &v1.ApiUpLoadController{}, "*:ErrorCapture")
	}
}
