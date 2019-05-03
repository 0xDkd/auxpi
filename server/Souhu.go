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
	"strings"

	"github.com/astaxie/beego"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type SouHu struct {
	FileLimit []string
	MaxSize   int
}

func (s *SouHu) Upload(image *ImageParam) (ImageReturn, error) {
	url := "http://changyan.sohu.com/api/2/comment/attachment"



	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "file",
		Value: *image.Content,
		Type:  image.Type,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var re = regexp.MustCompile(`(?m)"{\\"url\\":\\"(.*)\\"}"`)
	imgFix := re.FindAllStringSubmatch(data, -1)[0][1]
	url = strings.Replace(imgFix, "http", "https", 4)

	return ImageReturn{
		Url: url,
		ID:  17,
	}, nil
}

func (s *SouHu) UploadToSouHu(img []byte, imgInfo string, imgType string) string {
	url := "http://changyan.sohu.com/api/2/comment/attachment"
	name := tools.GetFileNameByMimeType(imgInfo)
	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var re = regexp.MustCompile(`(?m)"{\\"url\\":\\"(.*)\\"}"`)
	imgFix := re.FindAllStringSubmatch(data, -1)[0][1]
	url = strings.Replace(imgFix, "http", "https", 4)
	beego.Alert(url)
	return url
}
