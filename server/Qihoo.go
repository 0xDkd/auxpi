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

package server

import (
	"regexp"

	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type Qihoo struct {
	FileLimit []string
	MaxSize int
}

func (this *Qihoo) UploadToQihoo(img []byte, imgInfo string, imgType string) string {
	url := "http://kuaichuan.360.cn/upload/img"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "upload",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var re = regexp.MustCompile(`(?m)data-imgkey="(.*)"`)
	imgKey := re.FindAllStringSubmatch(data, -1)[0][1]
	url = "https://ps.ssl.qhmsg.com/" + imgKey
	return url
}
