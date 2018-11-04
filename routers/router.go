package routers

import (
	"auxpi/controllers"
	"auxpi/controllers/api"
	"auxpi/controllers/upload"
	"github.com/astaxie/beego"
)

func init() {
	//注册控制器内的路由
	beego.Include(&controllers.UpLoadController{})
	beego.Include(&page.PagesController{})
	beego.Include(&api.ApiUpLoadController{})
}

