package controllers

import (
	"auxpi/auxpiAll"
	"auxpi/auxpiAll/e"
	"auxpi/bootstrap"
	"auxpi/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"log"
	"strings"
)

type WebUpLoadController struct {
	beego.Controller
	utils.UpLoadTools
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}


//代码冗余，但是使用 API 会造成不必要的消耗
func (this *WebUpLoadController) UpLoadHandle() {
	//获取上传类型
	apiSelect := this.GetString("apiSelect")
	f, h, err := this.GetFile("image")
	defer f.Close()
	if err != nil {
		log.Fatal("File Upload Err", err)
	}
	//是否为空文件
	if f == nil {
		this.errorResp(e.ERROR_FILE_IS_EMPTY)
		return
	}
	//检测是否超出大小限制
	if h.Size > bootstrap.SiteConfig.SiteUpLoadMaxSize<<20 {
		this.errorResp(e.ERROR_FILE_IS_TOO_LARGE)
		return
	}
	//验证
	validate := this.Validate(h.Header.Get("Content-Type"), h.Filename)
	if validate {
		url := this.HandleUrl(apiSelect, f, h)
		//如果有返回值
		if strings.HasPrefix(url, "http") {
			this.succResp(200, url, h.Filename)
			return
		}
		logs.Notice(h.Filename + "上传" + apiSelect + "失败")
		this.errorResp(e.ERROR_CAN_NOT_GET_IMG_URL)
		return
	}
	//返回失败 json
	this.errorResp(e.ERROR_FILE_TYPE)
	return
}

//错误resp
func (this *WebUpLoadController) errorResp(code int) {
	result := &auxpi.RespJson{}
	result.Code = code
	result.Msg = e.GetMsg(code)
	this.Data["json"] = result
	this.ServeJSON()
}

//成功 resp
func (this *WebUpLoadController) succResp(code int, url string, name string) {
	result := &auxpi.ResultJson{}
	result.Code = code
	result.Msg = e.GetMsg(code)
	result.Data.Url = url
	result.Data.Name = name
	this.Data["json"] = result
	this.ServeJSON()
}
