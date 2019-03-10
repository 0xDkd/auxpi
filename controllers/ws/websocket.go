package ws

import (
	"auxpi/tools"
	"auxpi/utils"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

func (this *WebSocketController) Prepare() {
	this.EnableXSRF = false
}

var clients = make(map[*websocket.Conn]bool)


func (this *WebSocketController) Join() {
	//update
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//需要验证 cookie
	token := this.Ctx.Input.Cookie("Admin-Token")

	_, err := utils.ParseToken(token)
	if err != nil {
		beego.Alert("Token parsing unsuccessful")
		return
	}

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	clients = tools.AddClient(ws)

	if err != nil {
		beego.Alert("[Upgrade To WebSocket Error] :", err)
	}

	go handleMessage()
	this.ServeJSON()

	defer ws.Close()


	// Message receive loop.
	var read = make(map[string]string)

	for {
		err = ws.ReadJSON(&read)
		if err != nil {
			beego.Alert("[Json Read Error] :", err,clients)
			delete(clients, ws)
			beego.Alert("Now Client:",clients)
			err =ws.Close()
			if err!=nil {
				beego.Alert(err)
			}
			break
		}

	}

}

func handleMessage() {
	for client := range clients {
		beego.Alert("当前存在的客户端:",clients)
		for {
			select {
			case msg := <-tools.MsgChan:
				beego.Alert(msg)
				err := client.WriteJSON(msg)
				beego.Alert(clients)
				if err != nil {
					beego.Alert(clients)
					beego.Alert("[Write Error]:",err)
					//client.Close()
					delete(clients, client)
					return
				}
			}

		}

	}
}
