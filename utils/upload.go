package utils

import (
	"auxpi/bootstrap"
	"auxpi/server"
	"mime/multipart"
	"strings"
)

type UpLoadTools struct {
	server.Smms
	server.SouGou
	server.Sina
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}
var siteConfig = bootstrap.Config()
//返回不同图床的 URL
func (this *UpLoadTools) HandleUrl(apiSelect string, f multipart.File, h *multipart.FileHeader) string {
	imgMime := h.Header.Get("Content-Type")
	imgInfo := h.Header.Get("Content-Disposition")
	//读取文件
	size := h.Size
	fileContent := make([]byte, size)
	f.Read(fileContent)
	url := ""
	switch apiSelect {
	case "SouGou":
		url = this.UpLoadToSouGou(fileContent)
	case "Sina":
		if siteConfig.SiteUploadWay.OpenSinaPicStore == false {
			url = ""
		}
		url = this.UpLoadToSina(fileContent, imgMime)
	case "Smms":
		url = this.UpLoatToSmms(fileContent, imgInfo)
	default:
		switch siteConfig.ApiDefault {
		case "SouGou":
			url = this.UpLoadToSouGou(fileContent)
		case "Smms":
			url = this.UpLoatToSmms(fileContent, imgInfo)
		}

	}
	return url
}

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

