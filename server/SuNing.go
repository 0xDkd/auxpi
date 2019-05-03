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
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type SuNing struct {
	FileLimit []string
	MaxSize   int
}

func (this *SuNing) Upload(image *ImageParam) (ImageReturn, error) {
	url := "http://review.suning.com/imageload/uploadImg.do"



	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "Filedata",
		Value: *image.Content,
		Type:  image.Type,
	}

	data := tools.SuNingFormPost(file, url)
	s := auxpi.SuNingResp{}
	err := s.UnmarshalJSON([]byte(data))

	if err != nil {
		return ImageReturn{}, err
	}

	url = "https:" + s.Src + ".jpg"

	return ImageReturn{
		Url: url,
		ID:  13,
	}, nil
}

func (this *SuNing) UploadToSuNing(img []byte, imgInfo string, imgType string) string {
	url := "http://review.suning.com/imageload/uploadImg.do"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "Filedata",
		Value: img,
		Type:  imgType,
	}

	data := tools.SuNingFormPost(file, url)
	var m = make(map[string]string)
	json.Unmarshal([]byte(data), &m)
	beego.Alert(data)
	url = "https:" + m["src"] + ".jpg"

	return url
}
