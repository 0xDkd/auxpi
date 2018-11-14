package middleware

import (
	"auxpi/auxpiAll"
	"auxpi/auxpiAll/e"
	"auxpi/utils"
	"github.com/astaxie/beego/context"
	"time"
)

var JWT = func(ctx *context.Context) {
	var code int
	var data interface{}
	var token string

	code = e.SUCCESS
	token = ctx.Request.Header.Get("X-Token")
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
		ctx.Output.Header("Content-Type","application/json; charset=UTF-8")
		ctx.ResponseWriter.Write(info)
		return
	}
}
