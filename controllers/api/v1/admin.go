package v1

import (
	"auxpi/auxpiAll"
	"auxpi/auxpiAll/e"
	"auxpi/bootstrap"
	"auxpi/controllers/api/base"
	"auxpi/models"
	"auxpi/tools"
	"auxpi/utils"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type Admin struct {
	base.ApiController
}

//管理员信息
func (u *User) GetInfo() {
	token := u.Ctx.Request.Header.Get("X-Token")
	claims, err := utils.ParseToken(token)
	if err != nil {
		beego.Alert("Token parsing unsuccessful")
		return
	}
	re := models.GetUserInfo(claims.Username)

	u.Data["json"] = &auxpi.RespJson{
		Code: 200,
		Msg:  e.GetMsg(200),
		Data: re,
	}
	u.ServeJSON()
}

//重置 & 显示配置
func (a *Admin) ShowConf() {
	data := &bootstrap.SiteConfig
	bootstrap.Reload()
	a.Data["json"] = data
	a.ServeJSON()
}

//获取图片列表
func (a *Admin) GetImages() {
	page := a.Input().Get("page")
	limit := a.Input().Get("limit")
	storeID := a.Input().Get("type")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		logs.Alert("The Type of page is not correct")
	}
	err = nil
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		logs.Alert("The type of limit is not correct")
	}
	intStoreID, _ := strconv.Atoi(storeID)
	intPage, intLimit = utils.GetPage(intPage, intLimit)
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if intStoreID != 0 {
		maps["store_id"] = intStoreID
	}
	data["list"], data["total"] = models.GetImages(intPage, intLimit, maps)
	data["msg"] = "数据获取成功"
	data["code"] = 200
	a.Data["json"] = data
	a.ServeJSON()
}

//获取储存列表
func (a *Admin) GetStoreList() {
	data := make(map[string]interface{})
	data["list"] = models.GetStores()
	data["code"] = 200
	data["msg"] = e.GetMsg(200)
	a.Data["json"] = data
	a.ServeJSON()
}

//获取备份图片
func (a *Admin) GetSyncImages() {
	page := a.Input().Get("page")
	limit := a.Input().Get("limit")
	storeID := a.Input().Get("type")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		logs.Alert("The Type of page is not correct")
	}
	err = nil
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		logs.Alert("The type of limit is not correct")
	}
	intStoreID, _ := strconv.Atoi(storeID)
	intPage, intLimit = utils.GetPage(intPage, intLimit)
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if intStoreID != 0 {
		maps["store_id"] = intStoreID
	}
	data["list"], data["total"] = models.GetSyncImages(intPage, intLimit, maps)
	data["msg"] = "数据获取成功"
	data["code"] = 200
	a.Data["json"] = data
	a.ServeJSON()
}

//删除图片
func (a *Admin) DelImages() {
	ids := make(map[string][]int)
	data := make(map[string]interface{})
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &ids)

	if err != nil {
		beego.Alert("*** Delete images error ***")
		beego.Alert(err)
		return
	}
	err = models.DelImages(ids["id"])
	if err != nil {
		beego.Alert("*** Model Delete images error ***")
		beego.Alert(err)
		return
	}
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	a.Data["json"] = data
	a.ServeJSON()

}

//删除同步的图片
func (a *Admin) DelSyncImages() {
	ids := make(map[string][]int)
	data := make(map[string]interface{})
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &ids)

	if err != nil {
		beego.Alert("*** Delete images error ***")
		beego.Alert(err)
		return
	}
	err = models.DelSyncImage(ids["id"])

	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	a.Data["json"] = data
	a.ServeJSON()

}

//协程同步图片
func (a *Admin) SyncImages() {
	var data = make(map[string]interface{})
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &data)
	if err != nil {
		beego.Alert("[Sync API Error] :", err)
		return
	}
	list, ok := data["list"].([]interface{})
	if !ok {
		beego.Alert("[Sync API Error] :", err)
		return
	}
	l := len(list)
	//----------------
	var maps = make(map[uint]string)
	for _, value := range list {
		image := value.(map[string]interface{})
		id, _ := image["id"].(float64)
		intID := uint(id)
		url, _ := image["url"].(string)
		maps[intID] = url
	}
	//构造 models
	var images = make([]models.SyncImage, l)
	i := 0
	for key, value := range maps {
		images[i].ImageID = key
		images[i].External = value
		i++
	}
	resp := &auxpi.RespJson{}
	resp.Code = e.SUCCESS
	resp.Msg = e.GetMsg(e.SUCCESS)
	a.Data["json"] = resp
	if !tools.IsLock() {
		tools.Lock()
		go models.TestSyncImages(images, l-1)
	} else {
		resp.Code = e.ERROR_TASK_REPEAT
		resp.Msg = e.GetMsg(e.ERROR_TASK_REPEAT)
		a.Data["json"] = resp
	}

	a.ServeJSON()

}

//首页 统计
func (a *Admin) HomeImageStore() {
	data := make(map[string]interface{})
	list := models.GetAllImagesStoreNumber()
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	data["list"] = list
	a.Data["json"] = data
	a.ServeJSON()
}
func (a *Admin) HomeUserReport() {
	data := make(map[string]interface{})
	list := models.GetUserRegisterSevenDayReport()
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	data["list"] = list
	a.Data["json"] = data
	a.ServeJSON()
}
func (a *Admin) HomeApiReport() {
	data := make(map[string]interface{})
	list := models.GetApiInfo()
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	data["list"] = list

	a.Data["json"] = data
	a.ServeJSON()
}
func (a *Admin) HomeAllImageReport() {
	data := make(map[string]interface{})
	list := models.GetAllImagesReport()
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	data["list"] = list

	a.Data["json"] = data
	a.ServeJSON()
}
func (a *Admin) HomeLocalImageReport() {
	data := make(map[string]interface{})
	list := models.GetLocalImageReport()
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	data["list"] = list

	a.Data["json"] = data
	a.ServeJSON()
}

//获取用户列表
func (a *Admin) GetUserList() {
	page, size := utils.GetStringPage(a.Input().Get("page"), a.Input().Get("limit"))
	users, total := models.GetUsers(page, size)
	data := make(map[string]interface{})
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	data["total"] = total
	data["list"] = users

	a.Data["json"] = data
	a.ServeJSON()

}

//获取日志列表
func (a *Admin) LogList() {
	page, size := utils.GetStringPage(a.Input().Get("page"), a.Input().Get("limit"))
	logs, count := models.GetLogs(page, size, make(map[string]interface{}))
	data := make(map[string]interface{})
	data["code"] = e.SUCCESS
	data["msg"] = e.GetMsg(e.SUCCESS)
	data["list"] = logs
	data["total"] = count

	a.Data["json"] = data
	a.ServeJSON()

	a.ServeJSON()

}

//删除用户
func (a *Admin) DeleteUser() {
	data := make(map[string]int)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &data)
	if err != nil {
		a.Data["json"] = auxpi.RespJson{
			Code: e.INVALID_PARAMS,
			Msg:  e.GetMsg(e.INVALID_PARAMS),
		}
		a.ServeJSON()
		return
	}

	if data["id"] <= 0 {
		a.Data["json"] = &auxpi.RespJson{
			Code: e.ERROR_USER_NOT_EXIST,
			Msg:  e.GetMsg(e.ERROR_USER_NOT_EXIST),
		}
		a.ServeJSON()
		return
	}

	status := models.DeleteUserById(data["id"])
	if status {
		a.Data["json"] = &auxpi.RespJson{
			Code: e.SUCCESS,
			Msg:  e.GetMsg(e.SUCCESS),
		}
		a.ServeJSON()
		return
	}

	a.Data["json"] = &auxpi.RespJson{
		Code: e.ERROR_USER_NOT_EXIST,
		Msg:  e.GetMsg(e.ERROR_USER_NOT_EXIST),
	}
	a.ServeJSON()
	return
}
