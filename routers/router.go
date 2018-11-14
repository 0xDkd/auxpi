// @APIVersion 1.0.0
// @Title File Upload API
// @Description AuXpI 图床提供的 API 上传的方法
// @Contact aimerforreimu#gmail.com (#->@)
package routers

import (
	"auxpi/controllers"
	"auxpi/controllers/api/base"
	"auxpi/middleware"
	"auxpi/routers/api/auth"
	"auxpi/routers/api/v1"
	"github.com/astaxie/beego"
)

func init() {
	//正式环境不使用控制器内环境，调试时使用控制器内反射路由
	beego.Router("/", &page.PagesController{}, "get:IndexShow")
	beego.Router("/Sina", &page.PagesController{}, "get:SinaShow")
	beego.Router("/Smms", &page.PagesController{}, "get:SmmsShow")
	beego.Router("/about", &page.PagesController{}, "get:AboutShow")
	//auth
	auth.RegisterAuth()
	//dev 模式会打开测试路由方便调试
	if beego.BConfig.RunMode == "dev" {
		testRouter()
	}

	//v1 版本路由注册
	v1.RegisterControlApiV1()
	v1.RegisterOpenApiV1()
	//v2 版本路由注册
}

//测试路由，不要随便开启
func testRouter() {
	beego.InsertFilter("/test", beego.BeforeRouter, middleware.JWT)
	beego.Router("/test", &base.ApiController{}, "post:Test")

}
