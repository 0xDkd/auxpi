package routers

import (
	"auxpi/controllers/api"
	"github.com/astaxie/beego"
)

func RegisterOpenApi() {
	//不需要控制的 api 放到这里
	ns :=
		beego.NewNamespace("/api/v1",
			beego.NSRouter("/web_upload/", &api.ApiUpLoadController{}, "post:UpLoadHandle"),

		)
	beego.AddNamespace(ns)
}
