package auxpi

import "github.com/astaxie/beego/logs"

func init()  {
	logs.SetLogger(logs.AdapterFile,`{"filename":"auxpiRun.log","level":6,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Async(1e3)
	//logs.EnableFuncCallDepth(true)

}
