package api

import (
	"auxpi/auxpiAll"
	"auxpi/auxpiAll/e"
	"auxpi/controllers/api/base"
	"auxpi/models"
	"auxpi/utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

type Auth struct {
	base.ApiController
}
type authInfo struct {
	Username string `valid:"Required; MaxSize(32)" form:"username"`
	Password string `valid:"Required; MaxSize(32)" form:"password"`
}

func (this *Auth) GetAuthByUserName() {
	logs.Alert(this.GetString("username"))
	logs.Alert(this.GetString("password"))
	info := authInfo{}
	code := e.INVALID_PARAMS
	if err := this.ParseForm(&info); err != nil {
		//TODO :写入后台日志
		logs.Info("尝试鉴权失败")
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&info)
	logs.Alert(info)
	data := make(map[string]interface{})
	if ok {
		isExist := models.CheckAuth(info.Username, utils.GetSha256CodeWithSalt(info.Password))
		if isExist {
			token, err := utils.GenerateToken(info.Username, info.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}
		} else {
			logs.Alert("hello")
			code = e.ERROR_AUTH
		}
	} else {
		for _, errs := range valid.Errors {
			logs.Debug(errs.Value)
		}
	}
	resp := &auxpi.RespJson{
		code,
		e.GetMsg(code),
		data,
	}
	this.Data["json"] = resp
	this.ServeJSON()

}

func (this *Auth) GetAuthByUserEmail()  {
	
}
