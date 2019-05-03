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

package tools

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var (
	clients = make(map[*websocket.Conn]bool)
)

type Message struct {
	Title  string      `json:"title"`
	Msg    string      `json:"msg"`
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

var MsgChan = make(chan *Message)

//当前页面用户加入 websocket 列表中
func AddClient(ws *websocket.Conn) map[*websocket.Conn]bool {
	clients[ws] = true
	beego.Alert("新来了一个弟弟")
	return clients
}

//主动推送消息
func Send(msg *Message) {
	if len(clients) == 0 {
		beego.Alert("空的，不发")
		return
	}
	beego.Alert("准备发送，别堵住老子……")
	beego.Alert(msg)
	MsgChan <- msg
	beego.Alert("sent!")
	beego.Alert(msg)
}
