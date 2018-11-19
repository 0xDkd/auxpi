package auth

import (
	"auxpi/controllers/api/v1/auth"
	"github.com/astaxie/beego"
)

func RegisterAuth() {
	ns :=
		beego.NewNamespace("/api/v1",
			beego.NSRouter("/auth/login", &api.Auth{}, "post:GetAuthByUserName"),
		)
	beego.AddNamespace(ns)

}
