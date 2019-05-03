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
	"regexp"
)

//通过 MimeType 信息获取文件名称
func GetFileNameByMimeType(info string) string {
	//beego.Alert(info)
	pat := `filename="(.*)"`
	res := regexp.MustCompile(pat)
	name := res.FindAllStringSubmatch(info, -1)
	return name[0][1]
}
