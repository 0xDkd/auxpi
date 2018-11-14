package base

import (
	"auxpi/auxpiAll"
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

//所有的 APi 不需要开启 CSRF
func (this *ApiController) Prepare() {
	this.EnableXSRF = false
}

//调试 APi 只有 dev 模式下才能使用
func (this *ApiController) Test() {
	resp := &auxpi.RespJson{
		200,
		"你好世界级",
		make(map[string]interface{}),
	}
	this.Data["json"] = resp
	this.ServeJSON()
}
