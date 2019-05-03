// Copyright (c) 2019 aimerforreimu. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//  GNU GENERAL PUBLIC LICENSE
//                        Version 3, 29 June 2007
//
//  Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
//  Everyone is permitted to copy and distribute verbatim copies
// of this license document, but changing it is not allowed.
//
// repo: https://github.com/aimerforreimu/auxpi

package v2Router

import (
	"github.com/astaxie/beego"
	v2 "github.com/auxpi/controllers/api/v2"
	"github.com/auxpi/middleware"
)

//上传中间件
func RegisterUploadMiddleWare() {
	beego.InsertFilter("/api/v2/upload", beego.BeforeExec, middleware.Upload)

}

//不需要控制的 API
func RegisterApi() {
	//分发路由归属于 v2
	beego.Router(`dispatch/:hash(^[\w]+$)`, &v2.DispatchController{}, "get:Dispatch")

	// api 放到这里
	ns :=
		beego.NewNamespace("/api/v2",
			beego.NSRouter("/upload/", &v2.DispatchController{}, "post:UploadToRoot"),
		)
	beego.AddNamespace(ns)
}

////需要控制的 API
//func RegisterControlApi() {
//	if bootstrap.SiteConfig.OpenApiUpLoad {
//		beego.Router("api/v1/upload", &v1.ApiUploadController{}, "post:UpLoadHandle")
//	}
//}

