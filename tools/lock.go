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
