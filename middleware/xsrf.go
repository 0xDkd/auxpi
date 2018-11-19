package middleware

import (
	"auxpi/auxpiAll"
	"auxpi/auxpiAll/e"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
	"strings"
)

var (
	mCtx *context.Context
)

//自定义 CSRF 错误码返回
var XsrfError = func(ctx *context.Context) {
	resp := &auxpi.RespJson{}
	code := e.ERROR_ACCESS_DENIED
	_xsrf := ctx.Request.Form.Get("_xsrf")
	if _xsrf != getXSRFToken(beego.BConfig.WebConfig.XSRFKey, int64(beego.BConfig.WebConfig.XSRFExpire)) {
		resp.Code = code
		resp.Msg = e.GetMsg(code)
		resp.Data = ""
	}
	ctx.Output.Header("Content-Type", "application/json; charset=UTF-8")
	info, _ := resp.MarshalJSON()
	ctx.ResponseWriter.Write(info)
	return
}

func getXSRFToken(key string, expire int64) string {
	beego.Alert(key)
	beego.Alert(expire)
	token, ok := GetSecureCookie(key, "_xsrf")
	if !ok {
		token = string(utils.RandomCreateBytes(32))
		mCtx.SetSecureCookie(key, "_xsrf", token, expire)
	}

	return token
}

// GetSecureCookie Get secure cookie from request by a given key.
func  GetSecureCookie(Secret, key string) (string, bool) {
	val := mCtx.Input.Cookie(key)
	if val == "" {
		return "", false
	}

	parts := strings.SplitN(val, "|", 3)

	if len(parts) != 3 {
		return "", false
	}

	vs := parts[0]
	timestamp := parts[1]
	sig := parts[2]

	h := hmac.New(sha1.New, []byte(Secret))
	fmt.Fprintf(h, "%s%s", vs, timestamp)

	if fmt.Sprintf("%02x", h.Sum(nil)) != sig {
		return "", false
	}
	res, _ := base64.URLEncoding.DecodeString(vs)
	return string(res), true
}

