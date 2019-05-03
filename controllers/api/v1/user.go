package v1

import (
	"strconv"

	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/auxpiAll/e"
	"github.com/auxpi/controllers/api/base"
	"github.com/auxpi/models"
	"github.com/auxpi/utils"
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
	storeID := u.Input().Get("type")
	intStoreID, _ := strconv.Atoi(storeID)
	sort := u.Input().Get("sort")
	if sort == "+id" {
		sort = "ASC"
	} else {
		sort = "DESC"
	}

	num, size = utils.GetPage(num, size)
	maps := make(map[string]interface{})
	maps["user_id"] = uidx
	if intStoreID != 0 {
		maps["store_id"] = intStoreID
	}

	r, t := models.GetImagesByUserId(num, size, maps, sort)
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
