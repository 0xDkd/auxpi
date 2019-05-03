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

package middleware

import (
	"github.com/astaxie/beego/context"
)

var SSLRedirect = func(ctx *context.Context) {
	//fmt.Println(ctx.Input.Scheme())
	//if ctx.Input.Scheme() == "http" {
	//	fmt.Println(ctx.Input.Domain())
	//	fmt.Println(ctx.Input.Port())
	//	fmt.Println(ctx.Input.URI())
	//	port := strconv.Itoa(ctx.Input.Port())
	//	url := "https://" + ctx.Input.Domain() + ":" + port + ctx.Input.URI()
	//	ctx.Redirect(302, url)
	//}
	//acme.sh  --issue  -d ssl.demo-1s.com   --webroot  /root/auxpi/static/
}
