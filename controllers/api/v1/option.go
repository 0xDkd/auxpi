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

package v1

import (
	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/auxpiAll/e"
	"github.com/auxpi/controllers/api/base"
	"github.com/auxpi/models"
)

type OptionController struct {
	base.ApiController
}

//PUT api/v1/options/update&group=xx&key=xxx
func (o *OptionController) Update() {
	//不经过验证直接入库，感觉很危险，但是暂不考虑，毕竟这是后台，只有管理员才可以使用
	//TODO:validate
	group := o.Input().Get("group")
	key := o.Input().Get("key")

	if group == "" || key == "" {
		o.Data["json"] = auxpi.RespJson{
			Code: e.INVALID_PARAMS,
			Msg:  e.GetMsg(e.INVALID_PARAMS),
		}
		o.ServeJSON()
		return
	}
	models.UpdateOption(key, string(o.Ctx.Input.RequestBody), group)
	o.Data["json"] = auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
	}
	o.ServeJSON()
	return
}

//POST api/v1/options/add
func (o *OptionController) Add() {

}

//Delete api/v1/options/delete
func (o *OptionController) Destroy() {

}

//get api/v1/options/info?group=xx&key=xx
func (o *OptionController) Get() {
	group := o.Input().Get("group")
	key := o.Input().Get("key")
	if group == "" || key == "" {
		o.Data["json"] = auxpi.RespJson{
			Code: e.INVALID_PARAMS,
			Msg:  e.GetMsg(e.INVALID_PARAMS),
		}
		o.ServeJSON()
		return
	}
	//
	data := models.GetOption(key, group)
	o.Data["json"] = &auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  "ok",
		Data: data,
	}
	o.ServeJSON()
}

//get api/v1/options/stores?group=xx&key=xx
func (o *OptionController) GetStoreOptions() {
	group := o.Input().Get("group")
	key := o.Input().Get("key")
	if group == "" || key == "" {
		o.Data["json"] = auxpi.RespJson{
			Code: e.INVALID_PARAMS,
			Msg:  e.GetMsg(e.INVALID_PARAMS),
		}
		o.ServeJSON()
		return
	}
	//
	data := models.GetOption(key, group)
	status := models.GetStoreInfoByApi(key)
	var maps = make(map[string]interface{})
	maps["account"] = data
	maps["store"] = status
	o.Data["json"] = &auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  "ok",
		Data: maps,
	}
	o.ServeJSON()
}
