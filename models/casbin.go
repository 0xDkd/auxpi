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

package models

import (
	"fmt"

	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

func Create() {
	a := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/auxpi", true)
	e := casbin.NewEnforcer("conf/rbac_model.conf", a)
	err := e.LoadPolicy()
	if err != nil {
		fmt.Println(err)
	}

	e.AddPolicy("admin", "app", "/app/1", "GET")
	e.AddGroupingPolicy("alice", "admin", "app")

	err = e.SavePolicy()
	if err != nil {
		fmt.Println(err)
	}
}
