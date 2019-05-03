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
	"strings"
	"time"

	"github.com/auxpi/bootstrap"

	"github.com/astaxie/beego"
)

func LocalStoreInfo(n string, u string) (string, string, string) {
	nowTime := beego.Date(time.Now(), "Y/m/d/")
	str := `ZXCVBNMASDFGHJKLQWERTYUIOPzxcvbnmasdfghjklqwertyuiop1234567890`
	newName := bootstrap.GetRandomString(16, str)
	del := bootstrap.GetRandomString(16, str)
	sl := bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation
	sf := bootstrap.SiteConfig.SiteUploadWay.LocalStore.Link
	//beego.Alert(sf)
	bootstrap.FormatUrl(&sf)
	bootstrap.FormatUrl(&sl)
	//beego.Alert(sf)
	suffix := `.` + getImageSuffix(n, u)
	name := sl + nowTime + newName + suffix
	url := bootstrap.SiteConfig.SiteUrl + sf + "/" + nowTime + newName + suffix
	bootstrap.CheckPath(sl + nowTime)
	return url, name, del

}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

func getImageSuffix(name string, url string) (suffix string) {

	isSuffix := false
	for _, pType := range picType {
		if strings.HasSuffix(url, pType) {
			isSuffix = true
			break
		}
	}
	if isSuffix {
		return getSuffix(url)
	}
	return getSuffix(name)
}

func getSuffix(str string) (suffix string) {
	n := len(str)
	rs := []rune(str)
	suffix = string(rs[n-3 : n])
	//beego.Alert(suffix)
	if suffix == "peg" {
		suffix = "jpeg"
	}
	return suffix
}
