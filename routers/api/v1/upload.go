package v1Router

import (
	"github.com/auxpi/controllers"
	"github.com/auxpi/controllers/api/v1"
	"github.com/auxpi/middleware"

	"github.com/astaxie/beego"
)

//上传中间件
func RegisterUploadMiddleWare() {
	beego.InsertFilter("/api/v1/web_upload/", beego.BeforeExec, middleware.CookieUploadControl)
	beego.InsertFilter("/api/v1/web_upload/", beego.BeforeExec, middleware.UploadLimit)
	beego.InsertFilter("/api/v1/upload", beego.BeforeExec, middleware.UploadLimit)
	beego.InsertFilter("/api/v1/upload", beego.BeforeExec, middleware.Upload)
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
	//if bootstrap.SiteConfig.OpenApiUpLoad {
		beego.Router("api/v1/upload", &v1.ApiUploadController{}, "post:UpLoadHandle")
	//}
}
