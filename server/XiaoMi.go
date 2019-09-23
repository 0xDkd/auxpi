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

	"github.com/astaxie/beego"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type XiaoMi struct {
	FileLimit []string
	MaxSize   int
}

func (s *XiaoMi) Upload(image *ImageParam) (ImageReturn, error) {
	url := "https://shopapi.io.mi.com/homemanage/shop/uploadpic"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "pic",
		Value: *image.Content,
		Type:  image.Type,
	}
	data := tools.PayLoadFormPost(file, url)
	var xm = auxpi.XiaoMiResp{}
	err := xm.UnmarshalJSON([]byte(data))
	if err != nil {
		return ImageReturn{}, err
	}
	var re = regexp.MustCompile(`(?m)https:\/\/(.*?)&`)
	imgFix := re.FindAllStringSubmatch(xm.Result, -1)

	return ImageReturn{
		Url: "https://" + imgFix[0][1],
		ID:  14,
	}, nil
}

func (s *XiaoMi) UploadToXiaoMi(img []byte, imgInfo string, imgType string) string {
	url := "https://shopapi.io.mi.com/homemanage/shop/uploadpic"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "pic",
		Value: img,
		Type:  imgType,
	}
	data := tools.PayLoadFormPost(file, url)
	var xm = auxpi.XiaoMiResp{}
	beego.Alert(data)
	xm.UnmarshalJSON([]byte(data))
	//https:\/\/(.*?)&
	var re = regexp.MustCompile(`(?m)https:\/\/(.*?)&`)
	imgFix := re.FindAllStringSubmatch(xm.Result, -1)
	return "https://" + imgFix[0][1]
}
