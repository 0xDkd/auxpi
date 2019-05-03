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

package middleware

import (
	"encoding/json"
	"os"
	"time"

	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/bootstrap"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

type IpData struct {
	//创建时间 会刷新
	Create int64 `json:"create"`
	//提交次数
	Num int `json:"num"`
	//block 次数
	BlockNum int `json:"block_num"`
	//当前状态
	Status int `json:"status"`
}

var UploadLimit = func(ctx *context.Context) {
	var config = auxpi.LimitConfig{}
	err := config.UnmarshalJSON([]byte(models.GetOption("ip_limit", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("MIDDLEWARE", err)
		return
	}
	//如果没开直接返回
	if !config.Status {
		return
	}

	//获取请求的 ip
	ip := ctx.Input.IP()
	logs.Alert(ip)
	if ip == "127.0.0.1" || ip == "::1" {
		return
	}

	cip := bootstrap.Cache.Get(ip)
	//如果已经有这个 ip 的信息了
	if cip != nil {
		ipStatus, ok := cip.(IpData)
		if !ok {
			auxpiLog.SetAWarningLog("MIDDLEWARE", err)
			return
		}

		//如果这个 ip 已经被 block 了，
		if isBlock(ipStatus, ctx, ip, config) {
			return
		}

		//重新获取一遍数据
		ipStatus, _ = bootstrap.Cache.Get(ip).(IpData)

		//如果次数超过了
		if ipStatus.Num >= config.AllowNum &&
			time.Now().Unix()-ipStatus.Create >= config.AllowTime &&
			ipStatus.Status == 1 {
			//更新数据
			ipStatus.BlockNum++
			ipStatus.Status = 0
			ipStatus.Create = time.Now().Unix()

			if ipStatus.BlockNum >= config.DeadLine {
				blockForever(ipStatus, ip)
			}

			//删除原来的缓存并且重新生成
			reFresh(ipStatus, ip)
			isBlock(ipStatus, ctx, ip, config)
			return
		}

		//更新数据
		ipStatus.Num++
		reFresh(ipStatus, ip)

		return
	}
	var info = IpData{
		Create:   time.Now().Unix(),
		BlockNum: 0,
		Num:      1,
		Status:   1,
	}
	err = bootstrap.Cache.Put(ip, info, 24*time.Hour)
	if err != nil {
		logs.Alert(err)
	}
	return

}

//没有 block false
//block true
func isBlock(data IpData, ctx *context.Context, key string, config auxpi.LimitConfig) bool {
	_, status := os.Stat("block/" + key)
	errorInfo := auxpi.RespJson{
		Code: 503,
		Msg:  "Service Unavailable",
		Data: "",
	}
	info, _ := errorInfo.MarshalJSON()
	//现在 - 创建时间 >= 封号时长  解封
	if time.Now().Unix()-data.Create >= config.BlockTime &&
		data.Status == 0 &&
		data.BlockNum < config.DeadLine {
		data.Create = time.Now().Unix()
		data.Num = 1
		data.Status = 1
		//重置缓存中的 ip 信息
		reFresh(data, key)
		return false
	}

	if status == nil {
		ctx.Output.Header("Content-Type", "application/json; charset=UTF-8")
		ctx.ResponseWriter.Write(info)
		return true
	}

	if data.Status == 0 {
		ctx.Output.Header("Content-Type", "application/json; charset=UTF-8")
		ctx.ResponseWriter.Write(info)
		return true
	}
	return false
}

func reFresh(data IpData, key string) {
	v := bootstrap.Cache.Get(key)
	if v != nil {
		bootstrap.Cache.Delete(key)
		e := bootstrap.Cache.Put(key, data, 24*time.Hour)
		logs.Alert("ReFresh")
		if e != nil {
			logs.Error("[MiddleWare] :", e)
		}
		return
	}
}

func blockForever(data IpData, key string) {
	bootstrap.CheckPath("block/")
	_, status := os.Stat("block/" + key)
	if status != nil {
		ip, _ := os.OpenFile("block/"+key, os.O_CREATE|os.O_RDWR, 0755)
		info, err := json.Marshal(data)
		if err != nil {
			return
		}
		ip.Write(info)
	}
}
