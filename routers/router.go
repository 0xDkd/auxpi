// @APIVersion 1.0.0
// @Title Auxpi Upload API
// @Description AuXpI 图床提供的 API 上传的方法
package routers

import (
	"auxpi/controllers"
	"auxpi/controllers/api/base"
	"auxpi/controllers/api/v1"
	"auxpi/middleware"
	"auxpi/routers/api/auth"
	"auxpi/routers/api/v1"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//正式环境不使用控制器内环境，调试时使用控制器内反射路由
	beego.Router("/", &controllers.PagesController{}, "get:LocalShow")
	beego.Router("/SouGou", &controllers.PagesController{}, "get:IndexShow")
	beego.Router("/Sina", &controllers.PagesController{}, "get:SinaShow")
	beego.Router("/Smms", &controllers.PagesController{}, "get:SmmsShow")
	beego.Router("/cc", &controllers.PagesController{}, "get:CCShow")
	beego.Router("/Flickr", &controllers.PagesController{}, "get:FlickrShow")
	beego.Router("/Baidu", &controllers.PagesController{}, "get:BaiduShow")
	beego.Router("/360", &controllers.PagesController{}, "get:QihooShow")
	beego.Router("/NetEasy", &controllers.PagesController{}, "get:NetEasyShow")
	beego.Router("/Jd", &controllers.PagesController{}, "get:JdShow")
	beego.Router("/JueJin", &controllers.PagesController{}, "get:JueJinShow")
	beego.Router("/Ali", &controllers.PagesController{}, "get:AliShow")
	beego.Router("/Open", &controllers.PagesController{}, "get:OpenShow")
	beego.Router("/about", &controllers.PagesController{}, "get:AboutShow")

	if beego.BConfig.RunMode == "dev" {
		//auth

		//Vue 调试的时候需要跨域 ()
		setCors()
		//部分需要调试的路由
		testRouter()
	}

	//全局中间件
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
	beego.Router("/pic",&v1.Admin{},"get:GetSyncImages")
	//v2 版本路由注册

}

//测试路由，不要随便开启
func testRouter() {
	//Goroutine 信息
	//beego.Router("/go/",&base.ApiController{},"get:CPUinfo")
	beego.Router("/get_auxpi_info",&v1.AuxpiInfo{},"get:GetAuxpiSystemInfo")
	beego.Router("/get_site_config",&v1.AuxpiInfo{},"get:GetSiteConf")
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
