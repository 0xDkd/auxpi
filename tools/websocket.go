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
