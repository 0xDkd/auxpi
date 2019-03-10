package controllers

import (
	"auxpi/auxpiAll"
	"auxpi/auxpiAll/e"
	"auxpi/bootstrap"
	"auxpi/log"
	"auxpi/models"
	"auxpi/utils"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UsersController struct {
	beego.Controller
}

func (u *UsersController) commonStyle() {
	u.LayoutSections = make(map[string]string)
	u.LayoutSections["Script"] = "user/user_script.tpl"
	u.LayoutSections["Header"] = "user/user_header.tpl"
	u.LayoutSections["SiderBar"] = "user/user_sider_bar.tpl"
	u.LayoutSections["Content"] = "user/content/images.tpl"
	u.Data["xsrf_token"] = u.XSRFToken()
	r, _ := u.Ctx.GetSecureCookie(bootstrap.SiteConfig.AuxpiSalt, "r")
	if r == "admin" {
		u.Data["IsAdmin"] = true
	}
}

func (u *UsersController) Show() {
	u.commonStyle()
	uname, _ := u.Ctx.GetSecureCookie(bootstrap.SiteConfig.AuxpiSalt, "uname")
	page, size := utils.GetStringPage(u.Input().Get("page"), u.Input().Get("limit"))
	user, images, count := models.GetUserImagesByUserName(uname, size, page)
	tPage, _ := strconv.Atoi(u.Input().Get("page"))

	tplPage := utils.PageHtml(tPage, count, size)
	u.Data["User"] = &user
	u.Data["Images"] = &images
	u.Data["Page"] = tplPage

	u.Layout = "user/user_app.tpl"
	u.TplName = "user/user_app.tpl"
}

func (u *UsersController) Edit() {
	u.commonStyle()
	u.LayoutSections["Content"] = "user/content/edit.tpl"
	uname, _ := u.Ctx.GetSecureCookie(bootstrap.SiteConfig.AuxpiSalt, "uname")
	user := models.GetUserInfo(uname)

	u.Data["User"] = user

	u.Layout = "user/user_app.tpl"
	u.TplName = "user/user_app.tpl"
}

type UserResetPass struct {
	OldPass string `form:"old_password" valid:"Required;MinSize(6);MaxSize(32)"`
	NewPass string `form:"new_password" valid:"Required;MinSize(6);MaxSize(32)"`
	RePass  string `form:"re_password" valid:"Required;MinSize(6);MaxSize(32)"`
}

//重置密码
func (u *UsersController) ResetPass() {
	uid, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))
	passInfo := &UserResetPass{}
	if err := u.ParseForm(passInfo); err != nil {
		auxpiLog.SetAWarningLog("USER_LOGIN", err)
		u.ajaxErrorResp()
		return
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(passInfo)
	if !ok {
		u.ajaxErrorResp()
		return
	}

	if passInfo.RePass != passInfo.NewPass {
		u.Data["json"] = &auxpi.RespJson{
			Code: 200,
			Msg:  "两次密码不一致",
		}
		u.ServeJSON()
		return
	}

	ok = models.ResetUserPassWithOld(uid,
		utils.GetSha256CodeWithSalt(passInfo.OldPass),
		utils.GetSha256CodeWithSalt(passInfo.NewPass))

	if !ok {
		u.Data["json"] = &auxpi.RespJson{
			Code: 500,
			Msg:  e.GetMsg(500),
		}
		u.ServeJSON()
		return
	}

	u.Data["json"] = &auxpi.RespJson{
		Code: 200,
		Msg:  "修改成功",
	}
	u.ServeJSON()
	return

}

//ajax 错误相应封装
func (u *UsersController) ajaxErrorResp() bool {
	if u.IsAjax() {
		u.Data["json"] = &auxpi.RespJson{
			Code: e.INVALID_PARAMS,
			Msg:  e.GetMsg(e.INVALID_PARAMS),
		}
		u.ServeJSON()
		return false
	}
	return true
}
