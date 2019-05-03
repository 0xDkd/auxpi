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
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type Ali struct {
	FileLimit []string
	MaxSize   int
}

func (s *Ali) UploadToAli(img []byte, imgInfo string, imgType string) string {
	url := "https://kfupload.alibaba.com/mupload"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	//var header map[string]string
	data := tools.AliFormPost(file, url)
	j := auxpi.AliResp{}
	j.UnmarshalJSON([]byte(data))
	return j.Url
}

func (s *Ali) Upload(image *ImageParam) (ImageReturn, error) {
	url := "https://kfupload.alibaba.com/mupload"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "file",
		Value: *image.Content,
		Type:  image.Type,
	}
	//var header map[string]string
	data := tools.AliFormPost(file, url)
	j := auxpi.AliResp{}
	j.UnmarshalJSON([]byte(data))
	return ImageReturn{
		Url: j.Url,
		ID:  11,
	},
		nil
}
