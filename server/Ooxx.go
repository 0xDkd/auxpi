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

	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type OOXX struct {
	FileLimit []string
	MaxSize   int
}

func (s *OOXX) Upload(image *ImageParam) (ImageReturn, error) {
	url := "https://ooxx.ooo/upload"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "files[]",
		Value: *image.Content,
		Type:  image.Type,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var m = []string{}
	json.Unmarshal([]byte(data), &m)

	return ImageReturn{
		Url: "https://i.ooxx.ooo" + m[0],
		ID:  16,
	}, nil

}

func (s *OOXX) UploadToOOXX(img []byte, imgInfo string, imgType string) string {
	url := "https://ooxx.ooo/upload"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "files[]",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var m = []string{}
	json.Unmarshal([]byte(data), &m)

	return "https://i.ooxx.ooo" + m[0]
}
