package controllers

import (
	"auxpi/server"
	"encoding/base64"
	"github.com/astaxie/beego"
	"log"
	"strings"
)

//easyjson:json
type UpLoadController struct {
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

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

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
	imgMime := h.Header.Get("Content-Type")
	//验证
	validate := this.validate(imgMime, h.Filename)
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
			result := &ResultJson{}
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
	result := &ErrorJson{}
	result.Code = 500
	result.Msg = "上传失败"
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//验证文件后缀&文件MIME
func (this *UpLoadController) validate(contentType string, fileName string) bool {
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

//tools---
func Decode(enc *base64.Encoding, str string) string {
	data, err := enc.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(data)
}
