package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
} 


//渲染前端单页面
func (a *AdminController) Index() {
	a.TplName = "admin/index.html"
}
