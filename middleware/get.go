package middleware

//var Get := func() {
//
//}

//自定义 CSRF 错误码返回
//var XsrfError = func(ctx *context.Context) {
//	resp := &auxpi.RespJson{}
//	code := e.ERROR_ACCESS_DENIED
//	_xsrf := ctx.Request.Form.Get("_xsrf")
//	if _xsrf != getXSRFToken(beego.BConfig.WebConfig.XSRFKey, int64(beego.BConfig.WebConfig.XSRFExpire)) {
//		resp.Code = code
//		resp.Msg = e.GetMsg(code)
//		resp.Data = ""
//	}
//	ctx.Output.Header("Content-Type", "application/json; charset=UTF-8")
//	info, _ := resp.MarshalJSON()
//	ctx.ResponseWriter.Write(info)
//	return
//}
