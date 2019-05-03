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
	"time"

	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/auxpiAll/e"
	"github.com/auxpi/utils"

	"github.com/astaxie/beego/context"
)

var JWT = func(ctx *context.Context) {
	var code int
	var data interface{}
	var token string

	code = e.SUCCESS
	token = ctx.Request.Header.Get("X-Token")
	fmt.Println("JWJWJWJWJWJTTTTTTTTTTTTTTTTTT")
	if token == "" {
		code = e.INVALID_PARAMS
	} else {
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
	}
	if code != e.SUCCESS {
		errorInfo := auxpi.RespJson{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: data,
		}
		info, _ := errorInfo.MarshalJSON()
		ctx.Output.Header("Content-Type", "application/json; charset=UTF-8")
		ctx.ResponseWriter.Write(info)
		return
	}
}
