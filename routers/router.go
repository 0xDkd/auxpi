// @APIVersion 1.0.0
// @Title Auxpi Upload API
// @Description AuXpI 图床提供的 API 上传的方法
package routers

import (
	"github.com/auxpi/controllers/api/base"
	"github.com/auxpi/controllers/api/v1"
	"github.com/auxpi/middleware"
	"github.com/auxpi/routers/api/auth"
	"github.com/auxpi/routers/api/v1"
	v2Router "github.com/auxpi/routers/api/v2"

	"github.com/auxpi/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//正式环境不使用控制器内环境，调试时使用控制器内反射路由
	//找到对应的路由
	//首页
	beego.Router("/", &controllers.PagesController{}, "get:LocalShow")
	//从数据库中查询路由是否存在
	beego.Router("/:router([A-Za-z]+)", &controllers.PagesController{}, "get:Show")
	//关于页面
	beego.Router("about", &controllers.PagesController{}, "get:AboutShow")

	if beego.BConfig.RunMode == "dev" {
		//auth

		//Vue 调试的时候需要跨域 ()
		setCors()
		//部分需要调试的路由
		testRouter()
		//options
		beego.Router("/api/v1/options/stores", &v1.OptionController{}, "get:GetStoreOptions")
		beego.Router("/api/v1/options/info", &v1.OptionController{}, "get:Get")
		beego.Router("/api/v1/options/update", &v1.OptionController{}, "post:Update")
		beego.Router("/clear", &v1.Admin{}, "get:ClearCache")

	}

	//全局中间件
	beego.InsertFilter("*", beego.BeforeStatic, middleware.SSLRedirect)
	beego.InsertFilter("*", beego.BeforeRouter, middleware.CookieSignCheck)
	//v1 版本路由注册

	//上传路由
	v1Router.RegisterControlApi()
	v1Router.RegisterOpenApi()
	v1Router.RegisterUploadMiddleWare()

	//认证路由
	AuthRouter.RegisterMiddleWare()
	AuthRouter.RegisterRouter()
	AuthRouter.RegisterApi()

	//管理员路由
	v1Router.RegisterAdminMiddleWare()
	v1Router.RegisterAdminApi()
	v1Router.RegisterAdminRouter()
	//用户路由
	v1Router.RegisterUserMiddleWare()
	v1Router.RegisterUserAPi()
	v1Router.RegisterUserRouter()
	//websocket
	v1Router.RegisterWs()

	//UserController
	v1Router.RegisterUserRouter()

	//测试路由
	beego.Router("/pic", &v1.Admin{}, "get:GetSyncImages")
	//v2 版本路由注册
	v2Router.RegisterUploadMiddleWare()
	v2Router.RegisterApi()
}

//测试路由，不要随便开启
func testRouter() {
	//Goroutine 信息
	//beego.Router("/go/",&base.ApiController{},"get:CPUinfo")
	beego.Router("/get_auxpi_info", &v1.AuxpiInfo{}, "get:GetAuxpiSystemInfo")
	beego.Router("/get_site_config", &v1.AuxpiInfo{}, "get:GetSiteConf")
	//
	beego.Router("/m", &base.ApiController{}, "get:Migrate")
	//show info
	beego.Router("/conf", &base.ApiController{}, "get:ShowIt")
}

//跨域设置
func setCors() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:9528"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//Vue 跨域请求，需要允许跨域
	beego.Router("*", &base.ApiController{}, "options:Options")
}
