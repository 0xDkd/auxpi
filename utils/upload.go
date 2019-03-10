package utils

import (
	"auxpi/auxpiAll"
	"auxpi/bootstrap"
	"auxpi/models"
	"auxpi/server"
	"auxpi/tools"
	"mime/multipart"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
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
	server.VimCN
	server.Local
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}
//返回不同图床的 URL
func (this *UpLoadTools) HandleUrl(userID int, apiSelect string, f multipart.File, h *multipart.FileHeader) string {
	imgMime := h.Header.Get("Content-Type")
	imgInfo := h.Header.Get("Content-Disposition")
	//读取文件
	imgName := tools.GetFileNameByMimeType(imgInfo)
	beego.Alert(imgMime, imgInfo, imgName)
	size := h.Size
	fileContent := make([]byte, size)
	f.Read(fileContent)
	var url string
	switch apiSelect {
	case "Local":
		showUrl, backUrl, del, path := this.UpLoadToLocal(imgName, fileContent)
		url = showUrl
		storeImage(backUrl, imgName, userID, 12, del, path)
	case "SouGou":
		url = this.UpLoadToSouGou(fileContent)
		//this.UploadToVimCN(fileContent, imgInfo, imgMime)
		proxyUrl := bootstrap.SiteConfig.SiteUrl + "api/proxy?url=" + url
		storeImage(proxyUrl, imgName, userID, 1)
	case "Sina":
		if bootstrap.SiteConfig.SiteUploadWay.OpenSinaPicStore == false {
			url = ""
		} else {
			url = this.UpLoadToSina(fileContent, imgMime)

			storeImage(url, imgName, userID, 2)
		}
	case "Smms":
		url = this.UpLoatToSmms(fileContent, imgInfo)

		storeImage(url, imgName, userID, 3)
	case "CC":
		durl, del := this.UploadToCC(fileContent, imgInfo, imgMime)
		url = durl
		storeImage(durl, imgName, userID, 4, del)
	case "Flickr":
		if bootstrap.SiteConfig.SiteUploadWay.OpenFlickrStore == false {
			url = ""
		} else {
			file, err := h.Open()
			if err != nil {
				logs.Alert(err)
			}
			url = this.UploadToFlickr(file, h.Filename)

			storeImage(url, imgName, userID, 5)
		}
	case "Baidu":
		url = this.UploadToBaidu(fileContent, imgInfo, imgMime)

		storeImage(url, imgName, userID, 6)
	case "Qihoo":
		url = this.UploadToQihoo(fileContent, imgInfo, imgMime)

		storeImage(url, imgName, userID, 7)
	case "NetEasy":
		url = this.UploadToNetEasy(fileContent, imgInfo, imgMime)

		storeImage(url, imgName, userID, 8)
	case "Jd":
		url = this.UploadToJd(fileContent, imgInfo, imgMime)
		logs.Alert(url)
		storeImage(url, imgName, userID, 9)

	case "JueJin":
		url = this.UploadToJueJin(fileContent, imgInfo, imgMime)

		storeImage(url, imgName, userID, 10)
	case "Ali":
		url = this.UploadToAli(fileContent, imgInfo, imgMime)

		storeImage(url, imgName, userID, 11)
	case "Open":
		url = this.UpLoadToPublicSina(fileContent, imgInfo, imgMime)

		storeImage(url, imgName, userID, 2)

		//默认 API 上传
	default:
		switch bootstrap.SiteConfig.ApiDefault {
		case "Local":
			showUrl, backUrl, del, path := this.UpLoadToLocal(imgName, fileContent)
			url = showUrl
			storeImage(backUrl, imgName, userID, 12, del, path)
		case "SouGou":
			url = this.UpLoadToSouGou(fileContent)
			proxyUrl := bootstrap.SiteConfig.SiteUrl + "api/proxy?url=" + url
			storeImage(proxyUrl, imgName, userID, 1)
		case "Sina":
			if bootstrap.SiteConfig.SiteUploadWay.OpenSinaPicStore == false {
				url = ""
			} else {
				url = this.UpLoadToSina(fileContent, imgMime)

				storeImage(url, imgName, userID, 2)
			}
		case "Smms":
			url = this.UpLoatToSmms(fileContent, imgInfo)

			storeImage(url, imgName, userID, 3)
		case "CC":
			durl, del := this.UploadToCC(fileContent, imgInfo, imgMime)
			url = durl
			storeImage(durl, imgName, userID, 4, del)
		case "Flickr":
			if bootstrap.SiteConfig.SiteUploadWay.OpenFlickrStore == false {
				url = ""
			} else {
				file, err := h.Open()
				if err != nil {
					logs.Alert(err)
				}
				url = this.UploadToFlickr(file, h.Filename)

				storeImage(url, imgName, userID, 5)
			}
		case "Baidu":
			url = this.UploadToBaidu(fileContent, imgInfo, imgMime)

			storeImage(url, imgName, userID, 6)
		case "Qihoo":
			url = this.UploadToQihoo(fileContent, imgInfo, imgMime)

			storeImage(url, imgName, userID, 7)
		case "NetEasy":
			url = this.UploadToNetEasy(fileContent, imgInfo, imgMime)

			storeImage(url, imgName, userID, 8)
		case "Jd":
			url = this.UploadToJd(fileContent, imgInfo, imgMime)
			logs.Alert(url)
			storeImage(url, imgName, userID, 9)

		case "JueJin":
			url = this.UploadToJueJin(fileContent, imgInfo, imgMime)

			storeImage(url, imgName, userID, 10)
		case "Ali":
			url = this.UploadToAli(fileContent, imgInfo, imgMime)

			storeImage(url, imgName, userID, 11)
		case "Open":
			url = this.UpLoadToPublicSina(fileContent, imgInfo, imgMime)

			storeImage(url, imgName, userID, 2)
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

func storeImage(url, Name string, uID, sID int, v ...string) {

	if isUrl, isLocal := checkUrl(url); isLocal || isUrl {
		logs.Alert("yes")
		image := auxpi.ImageJson{}
		image.StoreID = sID
		image.UserID = uID
		image.Name = Name
		image.Url = url
		if len(v) != 0 {
			for key, value := range v {
				if key == 0 {
					//添加 delete
					image.Delete = value
				} else {
					//添加 Path
					image.Path = value
				}
			}
		}
		models.AddImage(&image)
	} else {
		name := models.GetStoreNameByImageID(sID)
		if url == "" {
			logs.Alert(name + "返回的 URL 无法插入数据库, URL为空")
			return
		}
		logs.Alert(name + "返回的 URL 无法插入数据库, URL为:" + url)

	}

}

func checkUrl(url string) (bool, bool) {
	var re = regexp.MustCompile(`(?m)https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&\/\/=]*)`)
	result := re.MatchString(url)
	//logs.Alert(result, url)
	isLocal := false
	var rex = regexp.MustCompile(`(?m)localhost`)
	isLocal = rex.MatchString(url)

	return result, isLocal
}
