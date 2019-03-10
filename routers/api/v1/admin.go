package v1Router

import (
	"auxpi/controllers"
	"auxpi/controllers/api/v1"
	"auxpi/middleware"

	"github.com/astaxie/beego"
)

////管理员登出
//	beego.Router("/api/test/logout", &v1.User{}, "post:GetFakerUserInfo")
//
//	//查询表格
//	beego.Router("/api/test/table/list", &v1.Admin{}, "get:GetFakerTable")
//
//	//查询图片
//	beego.Router("/api/test/images/list", &v1.Admin{}, "get:GetImages")
//
//	//代理图片
//	beego.Router("/api/proxy", &v1.Admin{}, "get:ProxyImages")
//
//	//查询储存方式
//	beego.Router("/api/test/store/list", &v1.Admin{}, "get:GetStoreList")
//
//	//删除图片
//	beego.Router("/api/test/images/delete", &v1.Admin{}, "post:DelImages")
//
//	//同步图片
//	beego.Router("/api/test/images/sync", &v1.Admin{}, "post:SyncImages")

//TODO :改到用户 api 里面去


//注册后台所需的中间件
func RegisterAdminMiddleWare() {
	//只要是 api/v1/admin/ 的 路由，都必须使用 jwt 认证
	beego.InsertFilter("/api/v1/admin/*", beego.BeforeExec, middleware.JWT)
}

//注意，admin 无模板引擎，所有数据全部通过 api 获取
func RegisterAdminApi() {
	ns := beego.NewNamespace("/api/v1/admin",
		//获取站点配置
		beego.NSRouter("/get_site_conf", &v1.Admin{}, "get:ShowConf"),
		//获取所有图片列表
		beego.NSRouter("/get_images_list", &v1.Admin{}, "get:GetImages"),
		//获取所有储存方式列表
		beego.NSRouter("/get_store_list", &v1.Admin{}, "get:GetStoreList"),
		//批量删除图片
		beego.NSRouter("/del_images", &v1.Admin{}, "post:DelImages"),
		//批量同步图片
		beego.NSRouter("/sync_images", &v1.Admin{}, "post:SyncImages"),
		//首页图床各状况统计
		beego.NSRouter("/get_store_report", &v1.Admin{}, "get:HomeImageStore"),
		//首页七天人数统计
		beego.NSRouter("/get_user_report", &v1.Admin{}, "get:HomeUserReport"),
		//首页统计七天 api 调用情况
		beego.NSRouter("/get_api_report", &v1.Admin{}, "get:HomeApiReport"),
		//首页七天所有图片链接统计
		beego.NSRouter("/get_all_images_report", &v1.Admin{}, "get:HomeAllImageReport"),
		//首页七天 本地图片增加统计
		beego.NSRouter("/get_local_images_report", &v1.Admin{}, "get:HomeLocalImageReport"),
		//获取所有用户
		beego.NSRouter("/get_users_list", &v1.Admin{}, "get:GetUserList"),
		//获取所有日志
		beego.NSRouter("/get_logs_list", &v1.Admin{}, "get:LogList"),
		//获取程序运行信息
		beego.NSRouter("/get_auxpi_info",&v1.AuxpiInfo{},"get:GetAuxpiSystemInfo"),
		//获取 RSA key
		beego.NSRouter("/get_rsa_key",&v1.AuxpiInfo{},"get:GetRsaKey"),
		//获取配置
		beego.NSRouter("/get_site_config",&v1.AuxpiInfo{},"get:GetSiteConf"),
		//更新配置
		beego.NSRouter("/update_site_config",&v1.AuxpiInfo{},"post:ResetSiteConf"),
		//重置 RSA
		//beego.NSRouter("/reset_rsa_key",&v1.AuxpiInfo{},"post:RsaReset"),
		//删除用户
		beego.NSRouter("/delete_user",&v1.Admin{},"post:DeleteUser"),
		beego.NSRouter("/get_sync_images",&v1.Admin{},"get:GetSyncImages"),
		beego.NSRouter("/del_sync_images",&v1.Admin{},"post:DelSyncImages"),
	)
	//管理员登入
	beego.AddNamespace(ns)
}

func RegisterAdminRouter() {
	//将打包好的 index.html 用模板引擎渲染
	beego.Router("/admin", &controllers.AdminController{}, "get:Index")
}
