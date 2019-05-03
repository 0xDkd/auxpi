package v1

import (
	"log"

	"github.com/auxpi/auxpiAll/e"
	"github.com/auxpi/bootstrap"
	"github.com/auxpi/controllers/api/base"
	"github.com/auxpi/models"
)

type ApiUploadController struct {
	base.ApiController
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

func (a *ApiUploadController) UpLoadHandle() {
	//默认游客
	userID := 0
	ip := a.Ctx.Input.IP()

	user := a.Ctx.Input.GetData("user_info")
	u, _ := user.(models.User)
	userID = u.ID
	//}
	//获取上传类型
	apiSelect := a.GetString("apiSelect")
	f, h, err := a.GetFile("image")
	if f == nil {
		a.ErrorResp(e.ERROR_FILE_IS_EMPTY)
	}
	if h.Size > bootstrap.SiteConfig.SiteUploadMaxSize<<20 {
		a.ErrorResp(e.ERROR_FILE_IS_TOO_LARGE)
	}
	defer f.Close()
	if err != nil {
		log.Fatal("File Upload Err", err)
	}
	//验证
	validate := a.Validate(h.Header.Get("Content-Type"), h.Filename)
	if validate {
		resp,_ := a.UploadHandle(userID, apiSelect, h, ip, true)
		a.Data["json"] = resp
		a.ServeJSON()
	}
	//返回失败 json
	a.Data["json"] = a.ErrorResp(e.ERROR_FILE_TYPE)
	a.ServeJSON()
	return
}
