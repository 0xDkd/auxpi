package main

import (
	_ "auxpi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLogger(logs.AdapterFile,`{"filename":"auxpi.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//beego.Alert(os.Getpid())
	beego.Run()
}

