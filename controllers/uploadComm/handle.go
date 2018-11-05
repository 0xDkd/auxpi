package uploadComm

import (
	"auxpi/auxpiAll"
	"auxpi/server"
	"github.com/astaxie/beego"
	"log"
	"strings"
)


type HandleController struct {
	beego.Controller
	server.Sina
	server.SouGou
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

//验证文件后缀&文件MIME
func (this *HandleController) Validate(contentType string, fileName string) bool {
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

func (this *HandleController) ErrorResponse(code int, msg string) {
	result := &auxpi.ErrorJson{}
	result.Code = code
	result.Msg = msg
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *HandleController) Check() {
	//获取上传类型
	apiSelect := this.GetString("apiSelect")
	f, h, err := this.GetFile("image")
	defer f.Close()
	if err != nil {
		log.Fatal("File Upload Err", err)
	}
	imgMime := h.Header.Get("Content-Type")
	//验证
	validate := this.Validate(imgMime, h.Filename)
	if validate {
		//读取文件
		size := h.Size
		fileContent := make([]byte, size)
		f.Read(fileContent)
		url := ""
		switch apiSelect {
		case "SouGou":
			url = this.UpLoadToSouGou(fileContent)
		case "Sina":
			url = this.UpLoadToSina(fileContent, imgMime)
		default:
			url = ""
		}
		//如果有返回值
		if strings.HasPrefix(url, "http") {
			//配置 json
			result := &auxpi.ResultJson{}
			result.Code = 200
			result.Msg = "上传成功"
			result.Data.Url = url
			result.Data.Name = h.Filename
			//beego.Alert(result)
			this.Data["json"] = result
			this.ServeJSON()
			return
		}

	}
	//返回失败 json
	result := &auxpi.ErrorJson{}
	result.Code = 500
	result.Msg = "上传失败"
	this.Data["json"] = result
	this.ServeJSON()
	return
}