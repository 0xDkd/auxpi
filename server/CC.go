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

//upload.cc
type CC struct {
	FileLimit []string
	MaxSize   int
}

func (s *CC) Upload(image *ImageParam) (ImageReturn, error) {
	url := "https://upload.cc/image_upload"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "uploaded_file[]",
		Value: *image.Content,
		Type:  image.Type,
	}

	j := auxpi.CCResp{}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	err := j.UnmarshalJSON([]byte(data))
	if err != nil {
		return ImageReturn{}, err
	}
	url = "https://upload.cc/" + j.SuccessImage[0].URL
	deleteJson := `[{"path":"` + j.SuccessImage[0].URL + `",key":"` + j.SuccessImage[0].Delete + `"}]`
	return ImageReturn{
			Url:    url,
			Delete: deleteJson,
			ID:     4,
		},
		nil
}

func (s *CC) UploadToCC(img []byte, imgInfo string, imgType string) (string, string) {
	url := "https://upload.cc/image_upload"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "uploaded_file[]",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	j := auxpi.CCResp{}
	j.UnmarshalJSON([]byte(data))

	url = "https://upload.cc/" + j.SuccessImage[0].URL

	deleteJson := `[{"path":"` + j.SuccessImage[0].URL + `",key":"` + j.SuccessImage[0].Delete + `"}]`
	return url, deleteJson
}
