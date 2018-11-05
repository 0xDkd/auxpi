package api

import (
	"auxpi/auxpiAll"
	"auxpi/bootstrap"
	"auxpi/utils"
	"github.com/astaxie/beego"
	"log"
	"strings"
)

type ApiUpLoadController struct {
	beego.Controller
	utils.UpLoadTools
}

//需要关闭 xsrf
func (this *ApiUpLoadController) Prepare() {
	this.EnableXSRF = false
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}
var siteConfig = bootstrap.Config()

func (this *ApiUpLoadController) ApiUpLoadHandle() {
	//检测是否开启 token 认证
	if siteConfig.ApiToken != "" {
		//需要进行验证
		apiToken := this.GetString("token")
		if apiToken != siteConfig.ApiToken {
			this.errorResp(403, "Forbidden")
			return
		}
	}
	//获取上传类型
	apiSelect := this.GetString("apiSelect")
	f, h, err := this.GetFile("image")
	if f == nil {
		this.errorResp(500, "No files were uploaded.")
	}
	if h.Size > siteConfig.SiteUpLoadMaxSize<<20 {
		this.errorResp(500, "File is too large.")
	}
	defer f.Close()
	if err != nil {
		log.Fatal("File Upload Err", err)
	}
	//验证
	validate := this.validate(h.Header.Get("Content-Type"), h.Filename)
	if validate {
		url := this.HandleUrl(apiSelect, f, h)

		//如果有返回值
		if strings.HasPrefix(url, "http") {
			this.succResp(200, "上传成功", url, h.Filename)
			return
		}
	}
	//返回失败 json
	this.errorResp(500, "上传失败")
	return
}

func (this *ApiUpLoadController) ErrorCapture() {
	this.errorResp(405, "Method not allowed")
}

//验证文件后缀&文件MIME
func (this *ApiUpLoadController) validate(contentType string, fileName string) bool {
	//首先检测文件的后缀
	isSuffix := false
	for _, pType := range picType {
		if strings.HasSuffix(fileName, pType) {
			isSuffix = true
			break
		}
	}
	//然后检测 MIME 类型
	//beego.Alert(contentType)
	if strings.HasPrefix(contentType, "image") && isSuffix {
		for _, pType := range picType {
			if strings.HasSuffix(contentType, pType) {
				return true
			}
		}

	}
	return false
}

//错误resp
func (this *ApiUpLoadController) errorResp(code int, msg string) {
	result := &auxpi.ErrorJson{}
	result.Code = code
	result.Msg = msg
	this.Data["json"] = result
	this.ServeJSON()
}

//成功 resp
func (this *ApiUpLoadController) succResp(code int, msg string, url string, name string) {
	result := &auxpi.ResultJson{}
	result.Code = code
	result.Msg = msg
	result.Data.Url = url
	result.Data.Name = name
	//beego.Alert(result)
	this.Data["json"] = result
	this.ServeJSON()
}
