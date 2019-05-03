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
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type Smms struct {
	FileLimit []string
	MaxSize   int
}

func (s *Smms) Upload(image *ImageParam) (ImageReturn, error) {

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	contentType := w.FormDataContentType()
	file, _ := w.CreateFormFile("smfile", image.Name)
	_, err := file.Write(*image.Content)

	if err != nil {
		return ImageReturn{}, err
	}

	err = w.Close()
	if err != nil {
		return ImageReturn{}, err
	}
	req, err := http.NewRequest("POST", "https://sm.ms/api/upload", body)
	if err != nil {
		return ImageReturn{}, err
	}

	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ImageReturn{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ImageReturn{}, err
	}
	sm := auxpi.SmResponse{}
	err = sm.UnmarshalJSON([]byte(string(data)))

	if err != nil {
		return ImageReturn{}, err
	}

	return ImageReturn{
		Url:    sm.Data.Url,
		Delete: sm.Data.Delete,
		ID:     3,
	}, nil
}

//上传 SM 图床 返回图片 URL
func (s *Smms) UpLoadToSmms(img []byte, imgInfo string) string {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	contentType := w.FormDataContentType()
	beego.Alert(imgInfo)
	name := tools.GetFileNameByMimeType(imgInfo)
	file, _ := w.CreateFormFile("smfile", name)
	file.Write(img)
	w.Close()
	req, _ := http.NewRequest("POST", "https://sm.ms/api/upload", body)
	req.Header.Set("Content-Type", contentType)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	sm := auxpi.SmResponse{}
	beego.Alert(string(data))
	sm.UnmarshalJSON([]byte(string(data)))
	return sm.Data.Url
}

//上传 SM 图床 返回图片 URL
func (s *Smms) DispatchUpLoadToSmms(img []byte) string {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	content_type := w.FormDataContentType()
	name := "123.png"
	file, _ := w.CreateFormFile("smfile", name)
	file.Write(img)
	w.Close()
	req, _ := http.NewRequest("POST", "https://sm.ms/api/upload", body)
	req.Header.Set("Content-Type", content_type)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	sm := auxpi.SmResponse{}
	beego.Alert(string(data))
	sm.UnmarshalJSON([]byte(string(data)))
	return string(sm.Data.Url)
}
