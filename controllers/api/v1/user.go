package v1

import (
	"auxpi/auxpiAll"
	"auxpi/auxpiAll/e"
	"auxpi/controllers/api/base"
	"auxpi/models"
	"auxpi/utils"
	"strconv"
)

type User struct {
	base.ApiController
}

//用户图片
func (u *User) UserImages() {
	uid, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))
	uidx := uint(uid)
	num, _ := strconv.Atoi(u.GetString("page"))
	size, _ := strconv.Atoi(u.GetString("limit"))

	num, size = utils.GetPage(num, size)
	r, t := models.GetImagesByUserId(num, size, uidx)
	data := make(map[string]interface{})
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	data["list"] = r
	data["total"] = t
	u.Data["json"] = data
	u.ServeJSON()
}

//获取单个用户的信息
func (u *User) UserInfo() {
	uid, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))
	user, status := models.GetUserInfoByID(uid)
	data := make(map[string]interface{})
	if status {
		data["code"] = e.SUCCESS
		data["msg"] = e.GetMsg(e.SUCCESS)
		data["list"] = user

		u.Data["json"] = data
		u.ServeJSON()
		return
	}

	data["code"] = e.ERROR_USER_NOT_EXIST
	data["msg"] = e.GetMsg(e.ERROR_USER_NOT_EXIST)
	data["list"] = ""

	u.Data["json"] = data
	u.ServeJSON()

}

//删除用户
func (u *User) UserDelete() {
	uid, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))
	ok := models.DeleteUserById(uid)
	if !ok {
		u.Data["json"] = &auxpi.RespJson{
			Code: e.ERROR_USER_DELTE,
			Msg:  e.GetMsg(e.ERROR_USER_DELTE),
		}
		u.ServeJSON()
		return
	}
	u.Data["json"] = &auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
	}
	u.ServeJSON()
}
