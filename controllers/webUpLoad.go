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

package controllers

import (
	"log"
	"strconv"

	"github.com/astaxie/beego"
	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/auxpiAll/e"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
)

type WebUpLoadController struct {
	beego.Controller
	Base
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

var site auxpi.SiteBase

func init() {
	err := site.UnmarshalJSON([]byte(models.GetOption("site_base", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("CONTROLLER", err)
	}
}

//代码冗余，但是使用 API 会造成不必要的消耗
func (w *WebUpLoadController) UpLoadHandle() {
	//获取用户信息
	userId, _ := strconv.Atoi(w.Ctx.GetCookie("id"))
	ip := w.Ctx.Input.IP()
	//获取上传类型
	apiSelect := w.GetString("apiSelect")
	f, h, err := w.GetFile("image")
	if err != nil {
		log.Fatal("File Upload Err", err)
	}
	defer f.Close()
	//是否为空文件
	if f == nil {
		w.ErrorResp(e.ERROR_FILE_IS_EMPTY)
		return
	}
	//检测是否超出大小限制
	if h.Size > site.SiteUploadMaxSize<<20 {
		w.ErrorResp(e.ERROR_FILE_IS_TOO_LARGE)
		return
	}
	//验证
	validate := w.Validate(h.Header.Get("Content-Type"), h.Filename)
	if validate {
		resp, _ := w.UploadHandle(userId, apiSelect, h, ip, true)
		w.Data["json"] = resp
		w.ServeJSON()
	}
	w.Data["json"] = w.ErrorResp(e.ERROR_FILE_TYPE)
	w.ServeJSON()
	return
}
