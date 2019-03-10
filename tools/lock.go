package tools

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var bm, err = cache.NewCache("memory", `{"interval":60}`)

func Lock() {
	if err != nil {
		beego.Alert("[Lock memory error]:", err)
		return
	}
	err = bm.Put("isLock", true, 100*time.Second)
	if err != nil {
		beego.Alert("[Lock memory error]:", err)
		return
	}
}

func UnLock() {
	if err != nil {
		beego.Alert("[Lock memory error]:", err)
		return
	}
	if IsLock() {
		err := bm.Delete("isLock")
		if err != nil {
			beego.Alert("[Lock memory error]:", err)
			return
		}
	}
}

func IsLock() bool {
	if bm.Get("isLock") != nil {
		return true
	}
	return false
}
