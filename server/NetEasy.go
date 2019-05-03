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
	"github.com/astaxie/beego"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type NetEasy struct {
	FileLimit []string
	MaxSize   int
}

func (s *NetEasy) Upload(image *ImageParam) (ImageReturn, error) {
	url := "http://you.163.com/xhr/file/upload.json"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "file",
		Value: *image.Content,
		Type:  image.Type,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	j := auxpi.NetEasyResp{}
	err := j.UnmarshalJSON([]byte(data))
	if err != nil {
		return ImageReturn{}, err
	}
	return ImageReturn{
		Url: j.Data[0],
		ID:  8,
	}, nil
}

func (s *NetEasy) UploadToNetEasy(img []byte, imgInfo string, imgType string) string {
	url := "http://you.163.com/xhr/file/upload.json"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	j := auxpi.NetEasyResp{}
	j.UnmarshalJSON([]byte(data))
	beego.Alert(j.Data[0])
	return j.Data[0]
}
