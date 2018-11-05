package controllers

import (
	"auxpi/auxpiAll"
	"auxpi/bootstrap"
	"auxpi/utils"
	"github.com/astaxie/beego"
	"log"
	"strings"
)

type UpLoadController struct {
	beego.Controller
	utils.UpLoadTools
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}
var siteConfig = bootstrap.Config()

func (c *UpLoadController) URLMapping() {
	c.Mapping("UpLoad", c.AuthUpLoadHandle)
}

// @router /api/v1/auth/upload/ [post]
func (this *UpLoadController) AuthUpLoadHandle() {
	//获取上传类型
	apiSelect := this.GetString("apiSelect")
	f, h, err := this.GetFile("image")
	defer f.Close()
	if err != nil {
		log.Fatal("File Upload Err", err)
	}
	//是否为空文件
	if f == nil {
		this.errorResp(500, "No files were uploaded.")
	}
	//检测是否超出大小限制
	if h.Size > siteConfig.SiteUpLoadMaxSize<<20 {
		this.errorResp(500, "File is too large.")
	}
	//验证
	validate := this.Validate(h.Header.Get("Content-Type"), h.Filename)
	if validate {
		url := this.HandleUrl(apiSelect, f, h)
		//如果有返回值
		if strings.HasPrefix(url, "http") {
			this.succResp(200, "上传成功", url, h.Filename)
		}

	}
	//返回失败 json
	this.errorResp(500, "上传失败")
	return
}

//错误resp
func (this *UpLoadController) errorResp(code int, msg string) {
	result := &auxpi.ErrorJson{}
	result.Code = code
	result.Msg = msg
	this.Data["json"] = result
	this.ServeJSON()
}

//成功 resp
func (this *UpLoadController) succResp(code int, msg string, url string, name string) {
	result := &auxpi.ResultJson{}
	result.Code = code
	result.Msg = msg
	result.Data.Url = url
	result.Data.Name = name
	//beego.Alert(result)
	this.Data["json"] = result
	this.ServeJSON()
}
