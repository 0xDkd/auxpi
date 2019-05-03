// Copyright (c) 2019 aimerforreimu. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//  GNU GENERAL PUBLIC LICENSE
//                        Version 3, 29 June 2007
//
//  Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
//  Everyone is permitted to copy and distribute verbatim copies
// of this license document, but changing it is not allowed.
//
// repo: https://github.com/aimerforreimu/auxpi

package v2

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/logs"
	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/auxpiAll/e"
	"github.com/auxpi/bootstrap"
	"github.com/auxpi/controllers/api/base"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
	"github.com/auxpi/server"
	"github.com/auxpi/tools"
	"github.com/auxpi/utils"
)

type DispatchController struct {
	base.ApiController
}

var site auxpi.SiteBase

func init() {
	err := site.UnmarshalJSON([]byte(models.GetOption("site_base", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("CONTROLLER", err)
	}
}

func (a *DispatchController) UploadToRoot() {
	r := models.GetOption("dispatch", "conf")
	dispatch := auxpi.Dispatch{}
	err := dispatch.UnmarshalJSON([]byte(r))
	if err != nil {

	}
	//默认游客
	userID := 0
	ip := a.Ctx.Input.IP()
	//文件验证
	//直接获取文件ctx 中的信息
	u := a.Ctx.Input.GetData("user_info")
	user, _ := u.(models.User)

	userID = user.ID
	f, h, err := a.GetFile("image")
	if f == nil {
		a.Data["json"] = a.ErrorResp(e.ERROR_FILE_IS_EMPTY)
		a.ServeJSON()
		return
	}
	if h.Size > site.SiteUploadMaxSize<<20 {
		a.Data["json"] = a.ErrorResp(e.ERROR_FILE_IS_TOO_LARGE)
		a.ServeJSON()
		return
	}
	defer f.Close()
	if err != nil {
		logs.Error("File Upload Err", err)
	}
	//验证
	validate := a.Validate(h.Header.Get("Content-Type"), h.Filename)
	if validate {
		resp, re := a.UploadHandle(userID, dispatch.Root, h, ip, false)
		if resp.Code == 200 {
			hash := bootstrap.GenerateUniqueString()
			//转 short
			r, _ := tools.ToShort(hash)
			dispatchUrl := bootstrap.SiteConfig.SiteUrl + "dispatch/" + hash
			//保存到数据库
			d := models.Distribution{
				Hash:        hash,
				Url:         "",
				StoreID:     0,
				UserID:      userID,
				IP:          ip,
				RootID:      re.ID,
				Short:       r[0],
				RootUrl:     re.Url,
				DispatchUrl: dispatchUrl,
			}

			models.CreateDistribution(d)

			//协程更新数据库
			var content = make([]byte, h.Size)
			_, err = f.Read(content)
			go a.RefreshUrl(d, content)
			var maps = make(map[string]string)
			maps["url"] = dispatchUrl
			maps["name"] = h.Filename
			a.Data["json"] = &auxpi.RespJson{
				Code: e.SUCCESS,
				Msg:  "ok",
				Data: maps,
			}
			a.ServeJSON()
			return
		}
		a.Data["json"] = &resp
		a.ServeJSON()
		return

	}
	a.Data["json"] = a.ErrorResp(e.ERROR_FILE_TYPE)
	a.ServeJSON()
}

func (a *DispatchController) RefreshUrl(distribution models.Distribution, content ...[]byte) string {
	//按照排名上传
	var auth = make(map[string]string)
	speedRank := server.NewClientsOrderBySpeed(auth)
	if distribution.Number > len(speedRank) {
		//说明图床全部挂掉了，可以进行新的一轮或者选择回源，目前仅支持回源
		// TODO:OOS,New Round
		if models.BackToSource(distribution.Hash) {
			//回源成功
			//TODO:失败处理
			return ""
		}
		return ""
	}
	//根据排名获取一个 client
	client := speedRank[distribution.Number]
	//第一次上传使用本地留下来的 byte
	imgMime := "image/png"
	imgInfo := ""
	data := content[0]
	if len(data) < 1 {
		resp, err := http.Get(distribution.RootUrl)
		if err != nil {
			auxpiLog.SetAnErrorLog("DISPATCH", err)
			return ""
		}
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			auxpiLog.SetAnErrorLog("DISPATCH", err)
			return ""
		}
		err = resp.Body.Close()
		if err != nil {
			auxpiLog.SetAnErrorLog("DISPATCH", err)
			return ""
		}
		imgMime = resp.Header.Get("Content-Type")
		imgInfo = resp.Header.Get("Content-Disposition")
	}

	var image = &server.ImageParam{
		Name:    distribution.Hash + ".png",
		Info:    imgInfo,
		Type:    imgMime,
		Content: &data,
	}
	r, err := client.Uploader.Upload(image)
	if err != nil {
		auxpiLog.SetAnErrorLog("DISPATCH", err)
		return ""
	}
	//入库更新+缓存
	models.ChangeUrl(distribution.Hash, r.Url, r.ID)
	//协程执行上传任务
	err = bootstrap.Cache.Put("url_"+distribution.Hash, r.Url, 3600)

	if err != nil {
		auxpiLog.SetAnErrorLog("DISPATCH", err)
		return ""
	}

	return r.Url

}

//分发链接对应方法
func (a *DispatchController) Dispatch() {
	hash := a.Ctx.Input.Param(":hash")
	//查找缓存
	cacheUrl := bootstrap.Cache.Get("url_" + hash)
	if cacheUrl != nil {
		url, _ := cacheUrl.(string)
		a.Ctx.Redirect(http.StatusFound, url)
		return
	}
	//查库
	dis, status := models.GetUrlByHash(hash)
	if !status {
		a.Abort("404")
		return
	}
	//检测 URL 状态
	if tools.CheckStatus(dis.Url) {
		err := bootstrap.Cache.Put("url_"+hash, dis.Url, 3600)
		if err != nil {
			auxpiLog.SetAnErrorLog("DISPATCH", err)
		}
		//写日志
		a.Ctx.Redirect(http.StatusFound, dis.Url)
		return
	}
	//按照权重重新上传
	//err := errors.New("从 " + dis.Url + " 更换")
	//auxpiLog.SetADebugLog("DISPATCH", err)
	a.Ctx.Redirect(http.StatusFound, a.RefreshUrl(dis))
}

//显示列表
// api/v2/dispatch/list?
func (a *DispatchController) ShowList() {
	page := a.Input().Get("page")
	limit := a.Input().Get("limit")
	storeID := a.Input().Get("type")
	sort := a.Input().Get("sort")
	if sort == "+id" {
		sort = "ASC"
	} else {
		sort = "DESC"
	}
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
	data["list"], data["total"] = models.GetDistributionList(intPage, intLimit, maps, sort)
	data["msg"] = "数据获取成功"
	data["code"] = 200
	a.Data["json"] = data
	a.ServeJSON()
}
