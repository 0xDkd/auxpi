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

//跨域Option
func (c *ApiController) Options() {
	c.AllowCross() //允许跨域
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

//跨域Allow
func (c *ApiController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:9527")       //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

//调试 APi 只有 dev 模式下才能使用
func (this *ApiController) Test() {
	data := make(map[string]interface{})
	data["_xsrf"] = this.XSRFToken()
	resp := &auxpi.RespJson{
		200,
		"你好世界级",
		data,
	}
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *ApiController) LoginTest() {
	this.AllowCross()
	data := make(map[string]interface{})
	//data["_xsrf"] = this.XSRFToken()
	token := "oasu09w4rsdlkfjasod9fuwer"
	data["token"] = token
	resp := &auxpi.RespJson{
		200,
		"login",
		data,
	}
	this.Data["json"] = resp
	this.ServeJSON()
}
