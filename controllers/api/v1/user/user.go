package userApi

import (
	"auxpi/auxpiAll"
	"auxpi/controllers/api/base"
)

type User struct {
	base.ApiController
}

//Need Use Jwt
//admin: {
//    roles: ['admin'],
//    token: 'admin',
//    introduction: '我是超级管理员',
//    avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
//    name: 'Super Admin'
//  }
func (this *User)GetFakerUserInfo()  {
	role := make([]string,1)
	role[0]="admin"
	user := &auxpi.UserInfo{}
	user.Token = "admin"
	user.Introduction= "I am Super Man"
	user.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	user.Name = "Super Man"
	user.Code = 200
	user.Roles = role

	this.Data["json"] = user
	this.ServeJSON()
}
