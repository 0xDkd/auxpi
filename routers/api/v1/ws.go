package v1Router

import (
	"auxpi/controllers/ws"

	"github.com/astaxie/beego"
)

//WebSocket 路由

func RegisterWs()  {
	beego.Router("/api/ws/sync", &ws.WebSocketController{}, "get:Join")
}
