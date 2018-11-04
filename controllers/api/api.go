package api

import (
	"auxpi/bootstrap"
	"auxpi/server"
	"github.com/astaxie/beego"
	"log"
	"strings"
)

//easyjson:json
type ApiUpLoadController struct {
	beego.Controller
	server.Sina
	server.SouGou
}

//easyjson:json
type ResultJson struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data fileData `json:"data"`
}

//easyjson:json
type fileData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

//easyjson:json
type ErrorJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//需要关闭 xsrf
func (this *ApiUpLoadController) Prepare() {
	this.EnableXSRF = false
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}
var siteConfig = bootstrap.Config()

func (c *ApiUpLoadController) URLMapping() {
	c.Mapping("UpLoad", c.AuthUpLoadHandle)
}

// @router /api/v1/upload/ [post]
func (this *ApiUpLoadController) AuthUpLoadHandle() {
	//获取上传类型
	apiSelect := this.GetString("apiSelect")
	f, h, err := this.GetFile("image")
	defer f.Close()
	if err != nil {
		log.Fatal("File Upload Err", err)
	}
	imgMime := h.Header.Get("Content-Type")
	//验证
	validate := this.validate(imgMime, h.Filename)
	if validate {
		//读取文件
		size := h.Size
		fileContent := make([]byte, size)
		f.Read(fileContent)
		url := ""
		if siteConfig.SiteUploadWay.OpenSinaPicStore == false {
			url = this.UpLoadToSouGou(fileContent)
		} else {
			switch apiSelect {
			case "SouGou":
				url = this.UpLoadToSouGou(fileContent)
			case "Sina":
				url = this.UpLoadToSina(fileContent, imgMime)
			default:
				url = ""
			}
		}
		//url := this.UpLoadToSina(fileContent,imgMime)
		//如果有返回值
		if strings.HasPrefix(url, "http") {
			//配置 json
			result := &ResultJson{}
			result.Code = 200
			result.Msg = "上传成功"
			result.Data.Url = url
			result.Data.Name = h.Filename
			beego.Alert(result)
			this.Data["json"] = result
			this.ServeJSON()
			return
		}

	}
	//返回失败 json
	result := &ErrorJson{}
	result.Code = 500
	result.Msg = "上传失败"
	this.Data["json"] = result
	this.ServeJSON()
	return
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
