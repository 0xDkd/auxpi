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

type TouTiao struct {
	FileLimit []string
	MaxSize   int
}

func (s *TouTiao) Upload(image *ImageParam) (ImageReturn, error) {
	url := "https://mp.toutiao.com/upload_photo/?type=json"



	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "photo",
		Value: *image.Content,
		Type:  image.Type,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var t = auxpi.TouTiaoResp{}
	err := t.UnmarshalJSON([]byte(data))
	if err != nil {
		return ImageReturn{}, err
	}

	return ImageReturn{
		Url: t.WebURL,
		ID:  19,
	}, nil
}

func (s *TouTiao) UploadToTouTiao(img []byte, imgInfo string, imgType string) string {
	url := "https://mp.toutiao.com/upload_photo/?type=json"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "photo",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var t = auxpi.TouTiaoResp{}
	t.UnmarshalJSON([]byte(data))
	return t.WebURL
}
