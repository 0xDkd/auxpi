package auth

import (
	"auxpi/controllers/api/auth"
	"github.com/astaxie/beego"
)

func RegisterAuth()  {
	beego.Router("/auth",&api.Auth{},"post:GetAuth")
}