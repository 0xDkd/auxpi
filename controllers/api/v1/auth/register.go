package api

import "github.com/astaxie/beego"

type Register struct {
	//base.ApiController
	beego.Controller
}


type registerInfo struct {
	//用户名只允许使用字母和数字
	Username string `valid:Required;MinSize(6);MaxSize(32);Match(/[a-zA-Z0-9\-_]+/) form:"username"`
	Password string `valid:Required;MinSize(6);MaxSize(32) form:"password"`
	Email    string `valid:Required;Email form"email"`
	//TODO :验证码暂时不考虑
	//Captcha  string `valid:Required;MinSize(4);MaxSize(6)`
}

func register()  {

}
