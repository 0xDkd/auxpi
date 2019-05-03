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

package auxpiLog

import (
	"fmt"

	"github.com/auxpi/models"

	"github.com/astaxie/beego/logs"
)

func SetAnErrorLog(t string, err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Error(err)
		models.AddLog(t, str, "SYSTEM", "ERROR")
		return true
	}
	return false
}

func SetAWarningLog(t string, err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Warning(err)
		models.AddLog(t, str, "SYSTEM", "WARNING")
		return true
	}
	return false

}

func SetADebugLog(t string, err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Debug(err)
		models.AddLog(t, str, "SYSTEM", "DEBUG")
		return true
	}
	return false

}

func SetAImageLog(err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Error("[IMAGE ERROR]: ", str)
		models.AddLog("IMAGE", str, "SYSTEM", "ERROR")
		return true
	}
	return false
}

func SetAOptionLog(err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Error(err)
		models.AddLog("OPTION", str, "SYSTEM", "ERROR")
		return true
	}
	return false
}

func SetALog(t string, err error, part string) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Error(err)
		models.AddLog(t, str, part, "ERROR")
		return true
	}
	return false
}

func SetUserLogin(str string) {
	models.AddLog("USER_LOGIN", str, "SYSTEM", "INFO")
}
