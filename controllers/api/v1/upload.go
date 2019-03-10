package v1

import (
	"auxpi/auxpiAll"
	"auxpi/auxpiAll/e"
	"auxpi/bootstrap"
	"auxpi/controllers/api/base"
	"auxpi/models"
	"auxpi/utils"
	"log"
	"strings"

	"github.com/astaxie/beego"
)

type ApiUpLoadController struct {
	base.ApiController
	utils.UpLoadTools
}

//需要关闭 xsrf
func (this *ApiUpLoadController) Prepare() {
	this.EnableXSRF = false
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

func (this *ApiUpLoadController) UpLoadHandle() {


	//默认游客
	userID := 0
	//检测是否开启 token 认证
	if bootstrap.SiteConfig.ApiToken != "" {
		//需要进行验证
		apiToken := this.GetString("token")
		//通过 token 获取用户信息
		user,status := models.GetUserInfoByToken(apiToken)
		userID = user.ID
		if status {
			this.errorResp(e.ERROR_AUTH_CHECK_TOKEN_FAIL)
			return
		}
	}
	//获取上传类型
	apiSelect := this.GetString("apiSelect")
	f, h, err := this.GetFile("image")
	if f == nil {
		this.errorResp(e.ERROR_FILE_IS_EMPTY)
	}
	if h.Size > bootstrap.SiteConfig.SiteUploadMaxSize<<20 {
		this.errorResp(e.ERROR_FILE_IS_TOO_LARGE)
	}
	defer f.Close()
	if err != nil {
		log.Fatal("File Upload Err", err)
	}
	//验证
	validate := this.validate(h.Header.Get("Content-Type"), h.Filename)
	if validate {
		url := this.HandleUrl(userID,apiSelect, f, h)
		//如果有返回值
		if strings.HasPrefix(url, "http") {
			models.AddApiLog(this.Ctx.Input.IP())
			beego.Alert("xxxxxxxxxxxx")
			this.succResp(200, url, h.Filename)
			return
		}
		this.errorResp(e.ERROR_CAN_NOT_GET_IMG_URL)
	}
	//返回失败 json
	this.errorResp(e.ERROR_FILE_TYPE)
	return
}

func (this *ApiUpLoadController) ErrorCapture() {
	this.errorResp(e.METHOD_NOT_ALLOWED)
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
func (this *ApiUpLoadController) errorResp(code int) {
	result := &auxpi.RespJson{}
	result.Code = code
	result.Msg = e.GetMsg(code)
	this.Data["json"] = result
	this.ServeJSON()
}

//成功 resp
func (this *ApiUpLoadController) succResp(code int, url string, name string) {
	result := &auxpi.ResultJson{}
	result.Code = code
	result.Msg = e.GetMsg(code)
	result.Data.Url = url
	result.Data.Name = name
	this.Data["json"] = result
	this.ServeJSON()
}
