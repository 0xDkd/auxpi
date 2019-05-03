package AuthRouter

import (
	"github.com/auxpi/controllers/api/v1"
	"github.com/auxpi/middleware"

	"github.com/astaxie/beego"
)

func RegisterMiddleWare() {
	//登录用户中间件 =>已登录重定向
	beego.InsertFilter("/login", beego.BeforeExec, middleware.CookieAuthedCheck)
	//重置密码中间件 =>已登录重定向
	beego.InsertFilter("/reset/*", beego.BeforeExec, middleware.CookieAuthedCheck)
	//忘记密码中间件 =>已登录重定向
	beego.InsertFilter("/forgot", beego.BeforeExec, middleware.CookieAuthedCheck)
	//注册中间件 =>已登录重定向
	beego.InsertFilter("/register/*", beego.BeforeExec, middleware.CookieAuthedCheck)
	//登出中间件  =>未登录重定向
	beego.InsertFilter("/logout", beego.BeforeExec, middleware.CookieAuthCheck)
}

//登录所使用的 api ，所有 api 都是无状态的
func RegisterApi() {
	ns :=
		beego.NewNamespace("/api/v1/auth",
			beego.NSRouter("/login", &v1.Auth{}, "post:GetAuthByUserName"),
			beego.NSRouter("/info", &v1.Admin{}, "get:GetInfo"),
			beego.NSRouter("/logout", &v1.Auth{}, "post:Destroy"),
		)
	beego.AddNamespace(ns)
}

//登录所使用的路由
func RegisterRouter() {
	//登录
	beego.Router("/login", &v1.Auth{}, "get:Show")
	beego.Router("/login", &v1.Auth{}, "post:Store")
	//注册
	beego.Router("/register", &v1.Auth{}, "get:Register")
	beego.Router("/register", &v1.Auth{}, "post:DoRegister")
	beego.Router(`/register/active/:token(^[\w|-]+$)`, &v1.Auth{}, "get:Register")
	//忘记密码
	beego.Router("/forgot", &v1.Auth{}, "get:Forgot")
	beego.Router("/forgot", &v1.Auth{}, "post:DoForgot")
	//密码找回
	beego.Router(`/reset/:token(^[\w|-]+$)`, &v1.Auth{}, "get:Reset")
	beego.Router(`/reset`, &v1.Auth{}, "post:DoReset")
	//信息页
	beego.Router("/msg", &v1.Auth{}, "get:Msg")
	//登出
	beego.Router("/logout", &v1.Auth{}, "get:Destroy")
}
