package base

import (
	"auxpi/auxpiAll"
	"auxpi/bootstrap"
	"strconv"

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
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
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

func (this *ApiController) ShowConf()  {
	data := &bootstrap.SiteConfig
	bootstrap.Reload()

	this.Data["json"] =  data
	this.ServeJSON()
}

func (this *ApiController) GetFakerTable() {
	table := &auxpi.FakerTable{}
	table.Code = 200
	mdata := make(map[string]string,6)
	data := make([]map[string]string,10)
	beego.Alert(data)
	for i:=0;i<10 ;i++  {
		mdata["ID"] = strconv.Itoa(i)
		mdata["title"] = "title"+strconv.Itoa(i)
		mdata["author"]= "author"+strconv.Itoa(i)
		mdata["pageviews"]= "pageviews"+strconv.Itoa(i)
		mdata["status"]= strconv.Itoa(i)
		mdata["display_time"]= "display_time"+strconv.Itoa(i)
		data[i] = mdata
	}
	table.Item = data
	this.Data["json"] = table

	this.ServeJSON()
}