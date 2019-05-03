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
	"fmt"

	"github.com/astaxie/beego/context"
	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/auxpiAll/e"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
)

var Upload = func(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
	var api = auxpi.ApiOptions{}
	err := api.UnmarshalJSON([]byte(models.GetOption("api_option", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("MIDDLEWARE", err)
	}
	fmt.Println(api)
	if token == "" {
		//API 是否需要认证
		if api.Auth {
			JWT(ctx)
			return
		}
		//如果开放对外的 api 接口，则无需验证
		return
	}
	//解析 token
	user, status := models.GetUserInfoByToken(token)
	if status {
		//将用户信息写入 ctx
		fmt.Println("sldkjlsakjf======>>>>>")
		ctx.Input.SetData("user_info", user)
		return
	}
	//返回错误
	errorInfo := auxpi.RespJson{
		Code: e.ERROR_UPLOAD_TOKEN_ERROR,
		Msg:  e.GetMsg(e.ERROR_UPLOAD_TOKEN_ERROR),
	}
	info, _ := errorInfo.MarshalJSON()
	ctx.Output.Header("Content-Type", "application/json; charset=UTF-8")
	ctx.ResponseWriter.Write(info)
	return
}
