package base

import (
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

//所有的 APi 不需要开启 CSRF
func (this *ApiController) Prepare() {
	this.EnableXSRF = false
}

//代理图片
func (this *ApiController) ProxyImages() {
	url := this.GetStrings("url")
	resp, _ := http.Get(url[0])
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	this.Ctx.Output.Header("Content-Type", resp.Header.Get("Content-Type"))
	this.Ctx.ResponseWriter.Write(data)
}