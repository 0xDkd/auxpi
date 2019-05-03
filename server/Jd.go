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
	"errors"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type Jd struct {
	FileLimit []string
	MaxSize   int
}

func (s *Jd) Upload(image *ImageParam) (ImageReturn, error) {
	url := "https://search.jd.com/image?op=upload"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "file",
		Value: *image.Content,
		Type:  image.Type,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var pre = regexp.MustCompile(`(?m)ERROR`)

	if !pre.MatchString(data) {
		var re = regexp.MustCompile(`(?m)\("(.*)"\)`)
		beego.Alert(data)
		imgFix := re.FindAllStringSubmatch(data, -1)[0][1]
		url = "https://img" + strconv.Itoa(rand.Intn(3)+11) + ".360buyimg.com/img/" + imgFix
		return ImageReturn{
			Url: url,
			ID:  9,
		}, nil
	} else {
		err := errors.New("Can not parse jingdong's response,please check response date ")
		return ImageReturn{}, err
	}
}

func (s *Jd) UploadToJd(img []byte, imgInfo string, imgType string) string {
	url := "https://search.jd.com/image?op=upload"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var pre = regexp.MustCompile(`(?m)ERROR`)

	if !pre.MatchString(data) {
		var re = regexp.MustCompile(`(?m)\("(.*)"\)`)
		imgFix := re.FindAllStringSubmatch(data, -1)[0][1]
		url = "https://img" + strconv.Itoa(rand.Intn(3)+11) + ".360buyimg.com/img/" + imgFix
		return url
	} else {
		return ""
	}

}
