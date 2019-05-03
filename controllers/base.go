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
	"mime/multipart"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/auxpiAll/e"
	"github.com/auxpi/bootstrap"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
	"github.com/auxpi/server"
	"github.com/auxpi/utils"
)

type Base struct {
	beego.Controller
}

//公共上传方法
func (b *Base) UploadHandle(userID int, apiSelect string, h *multipart.FileHeader, ip string, w bool) (*auxpi.RespJson, server.ImageReturn) {
	if apiSelect == "" {
		apiSelect = bootstrap.SiteConfig.ApiDefault
	}
	api := strings.ToLower(apiSelect)
	//判断是否在开放的图床中
	list := models.GetStores()
	//加入缓存当中
	for _, value := range list {
		if value.Status == false && value.Api == api {
			return &auxpi.RespJson{
				Code: e.ERROR_UPLOAD_PARAM,
				Msg:  e.GetMsg(e.ERROR_UPLOAD_PARAM),
			}, server.ImageReturn{}
		}
	}
	//Proxy
	client := server.NewClient(api)
	if client.Error != nil {
		auxpiLog.SetAWarningLog("CONTROLLER_BASE", client.Error)
		//return client
		var resp = &auxpi.RespJson{
			Code: e.ERROR_UPLOAD_PARAM,
			Msg:  e.GetMsg(e.ERROR_UPLOAD_PARAM),
		}
		return resp, server.ImageReturn{}
	}
	h.Filename = strings.TrimSpace(h.Filename)

	var content = make([]byte, h.Size)
	file, err := h.Open()
	_, err = file.Read(content)
	if err != nil {
		auxpiLog.SetAnErrorLog("CONTROLLER_BASE", client.Error)
		var resp = &auxpi.RespJson{
			Code: e.ERROR_UPLOAD_PARAM,
			Msg:  e.GetMsg(e.ERROR_UPLOAD_PARAM),
		}
		return resp, server.ImageReturn{}
	}

	//构造传入参数
	image := server.ImageParam{
		Name:    h.Filename,
		Type:    h.Header.Get("Content-Type"),
		Info:    h.Header.Get("Content-Disposition"),
		Content: &content,
	}
	client.Do(&image)

	if client.Error != nil {
		auxpiLog.SetAWarningLog("CONTROLLER_BASE", client.Error)
		return &auxpi.RespJson{
			Code: e.ERROR_CAN_NOT_UPLOAD,
			Msg:  e.GetMsg(e.ERROR_CAN_NOT_UPLOAD),
		}, server.ImageReturn{}
	}
	var data = make(map[string]interface{})
	data["url"] = client.Resp.Url
	data["delete"] = client.Resp.Delete
	data["name"] = h.Filename
	//插入数据库
	if w {
		go b.StoreImage(userID, h.Filename, ip, client.Resp)
	}
	return &auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  "ok",
		Data: data,
	}, client.Resp

}

//公共检测方法
func (b *Base) Validate(contentType string, fileName string) bool {
	//首先检测文件的后缀
	isSuffix := false
	for _, pType := range picType {
		if strings.HasSuffix(fileName, pType) {
			isSuffix = true
			break
		}
	}
	//然后检测 MIME 类型
	//beego.Alert(contentType)
	if strings.HasPrefix(contentType, "image") && isSuffix {
		for _, pType := range picType {
			if strings.HasSuffix(contentType, pType) {
				return true
			}
		}

	}
	return false
}

//错误resp
func (b *Base) ErrorResp(code int) *auxpi.RespJson {
	result := &auxpi.RespJson{}
	result.Code = code
	result.Msg = e.GetMsg(code)
	return result
}

//插入数据库
func (b *Base) StoreImage(userID int, name, ip string, image server.ImageReturn) {

	if isUrl, isLocal := utils.CheckUrl(image.Url); isLocal || isUrl {
		i := auxpi.ImageJson{}
		i.StoreID = image.ID
		i.UserID = userID
		i.IP = ip
		i.Name = name
		i.Url = image.Url
		i.Delete = image.Delete
		i.Path = image.Path
		models.AddImage(&i)
	} else {
		name := models.GetStoreNameByImageID(image.ID)
		if image.Url == "" {
			logs.Alert(name + "返回的 URL 无法插入数据库, URL为空")
			return
		}
		logs.Alert(name + "返回的 URL 无法插入数据库, URL为:" + image.Url)

	}
}
