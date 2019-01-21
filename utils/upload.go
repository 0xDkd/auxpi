package utils

import (
	"auxpi/bootstrap"
	"auxpi/server"
	"mime/multipart"
	"strings"

	"github.com/astaxie/beego/logs"
)

type UpLoadTools struct {
	server.Smms
	server.SouGou
	server.Sina
	server.Flickr
	server.OpenApi
	server.Baidu
	server.Jd
	server.Qihoo
	server.JueJin
	server.NetEasy
	server.CC
	server.Ali
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}
//返回不同图床的 URL
func (this *UpLoadTools) HandleUrl(apiSelect string, f multipart.File, h *multipart.FileHeader) string {
	imgMime := h.Header.Get("Content-Type")
	imgInfo := h.Header.Get("Content-Disposition")
	//读取文件
	size := h.Size
	fileContent := make([]byte, size)
	f.Read(fileContent)
	var url string
	switch apiSelect {
	case "SouGou":
		url = this.UpLoadToSouGou(fileContent)
	case "Sina":
		if bootstrap.SiteConfig.SiteUploadWay.OpenSinaPicStore == false {
			url = ""
		} else {
			url = this.UpLoadToSina(fileContent, imgMime)
		}
	case "Smms":
		url = this.UpLoatToSmms(fileContent, imgInfo)
	case "Flickr":
		if bootstrap.SiteConfig.SiteUploadWay.OpenFlickrStore == false {
			url = ""
		} else {
			file, err := h.Open()
			if err != nil {
				logs.Alert(err)
			}
			url = this.UploadToFlickr(file, h.Filename)
		}
	case "Baidu":
		url = this.UploadToBaidu(fileContent, imgInfo, imgMime)
	case "Qihoo":
		url = this.UploadToQihoo(fileContent, imgInfo, imgMime)
	case "Jd":
		url = this.UploadToJd(fileContent, imgInfo, imgMime)
	case "NetEasy":
		url = this.UploadToNetEasy(fileContent, imgInfo, imgMime)
	case "JueJin":
		url = this.UploadToJueJin(fileContent, imgInfo, imgMime)
	case "CC":
		url = this.UploadToCC(fileContent, imgInfo, imgMime)
	case "Ali":
		url = this.UploadToAli(fileContent, imgInfo, imgMime)
	case "Open":
		url = this.UpLoadToPublicSina(fileContent, imgInfo, imgMime)
	default:
		switch bootstrap.SiteConfig.ApiDefault {
		case "SouGou":
			url = this.UpLoadToSouGou(fileContent)
		case "Smms":
			url = this.UpLoatToSmms(fileContent, imgInfo)
		}
	}
	return url
}

//返回所有图床的 URL
//
//func (this *UpLoadTools) ReturnAllUrl(f multipart.File, h *multipart.FileHeader)  {
//	var urls  = make(map[string]string)
//
//	imgMime := h.Header.Get("Content-Type")
//	imgInfo := h.Header.Get("Content-Disposition")
//	//读取文件
//	size := h.Size
//	fileContent := make([]byte, size)
//	f.Read(fileContent)
//
//}

//验证文件后缀&文件MIME
func (this *UpLoadTools) Validate(contentType string, fileName string) bool {
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
