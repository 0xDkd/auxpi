package v1Router

import (
	"auxpi/controllers"
	"auxpi/controllers/api/v1"
	"auxpi/middleware"

	"github.com/astaxie/beego"
)

//注册路由所需的中间件

func RegisterUserMiddleWare() {
	//所有用户开头的路由全部需要登录
	beego.InsertFilter("/users/*", beego.BeforeExec, middleware.CookieAuthCheck)
	//用户信息需要 jwt 认证
	beego.InsertFilter("/api/v1/users/*", beego.BeforeExec, middleware.JWT)

}

//不需要控制的 API
func RegisterUserAPi() {
	//不需要控制的 api 放到这里
	ns :=
		beego.NewNamespace("/api/v1/users",
			//用户图片
			beego.NSRouter("/:id/images", &v1.User{}, "get:UserImages"),
			//用户信息
			beego.NSRouter("/:id/info", &v1.User{}, "get:UserInfo"),
		)
	beego.AddNamespace(ns)
}

//非 api 路由
func RegisterUserRouter() {
	ns :=
		beego.NewNamespace("/users",
			//用户后台首页
			//TODO:使用 API
			beego.NSRouter("/index", &controllers.UsersController{}, "get:Show"),
			//用户后台首页重定向
			beego.NSRouter("", &controllers.UsersController{}, "get:Show"),
			//用户个人信息+修改
			//TODO:使用 API
			beego.NSRouter("/edit", &controllers.UsersController{}, "get:Edit"),
			//修改用户信息
			beego.NSRouter("/:id/edit", &controllers.UsersController{}, "post:ResetPass"),
		)
	beego.AddNamespace(ns)
}
