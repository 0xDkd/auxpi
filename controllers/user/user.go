package user

import "github.com/astaxie/beego"

type UsersController struct {
	beego.Controller
}

func (this *UsersController)UserIndexShow()  {
	this.TplName = "user/user_app.tpl"
}