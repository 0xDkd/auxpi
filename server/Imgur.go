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
	"github.com/auxpi/bootstrap"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
	"github.com/auxpi/tools"
)

type Imgur struct {
	FileLimit []string
	MaxSize   int
}

func (s *Imgur) UploadToImgur(img []byte, imgInfo string, imgType string) (string, string) {
	url := "https://api.imgur.com/3/image"
	name := tools.GetFileNameByMimeType(imgInfo)
	file := &auxpi.FormFile{
		Name:  name,
		Key:   "image",
		Value: img,
		Type:  imgType,
	}

	var header = make(map[string]string)
	header["Authorization"] = "Client-ID " + bootstrap.SiteConfig.SiteUploadWay.ImgurAccount.ImgurID
	data := tools.FormPost(file, url, header)
	var jData = auxpi.ImgurResp{}
	jData.UnmarshalJSON([]byte(data))

	return "https://images.weserv.nl/?url=" + jData.Data.Link, jData.Data.Deletehash
}

func (s *Imgur) Upload(image *ImageParam) (ImageReturn, error) {

	var imgurAccount = auxpi.ImgurAccount{}

	err := imgurAccount.UnmarshalJSON([]byte(models.GetOption("imgur", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("CONTROLLER_BASE", err)
	}

	url := "https://api.imgur.com/3/image"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "image",
		Value: *image.Content,
		Type:  image.Type,
	}
	var header = make(map[string]string)
	header["Authorization"] = "Client-ID " + imgurAccount.ImgurID
	data := tools.FormPost(file, url, header)
	var jData = auxpi.ImgurResp{}
	err = jData.UnmarshalJSON([]byte(data))
	if err != nil {
		return ImageReturn{}, err
	}
	if imgurAccount.Proxy.Status {
		jData.Data.Link = imgurAccount.Proxy.Node + jData.Data.Link
	}
	return ImageReturn{
		Url:    jData.Data.Link,
		Delete: jData.Data.Deletehash,
		ID:     6,
	}, nil
}
