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
	"bufio"
	"errors"
	"os"
	"time"

	"github.com/astaxie/beego"
	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/bootstrap"
	"github.com/auxpi/models"
)

type Local struct {
	FileLimit []string
	MaxSize   int
}

var local = auxpi.LocalStore{}
var site = auxpi.SiteBase{}

func (s *Local) Upload(image *ImageParam) (ImageReturn, error) {
	var err = local.UnmarshalJSON([]byte(models.GetOption("local", "conf")))
	if err != nil {
		return ImageReturn{}, nil
	}
	err = site.UnmarshalJSON([]byte(models.GetOption("site_base", "conf")))
	if !local.Status {
		err := errors.New("LocalStorage is close by user,please open it ")
		return ImageReturn{}, err
	}

	host := &site.SiteUrl
	storeLocation := &local.StorageLocation
	softLink := &local.Link
	//修正URL
	bootstrap.FormatUrl(softLink)
	bootstrap.FormatUrl(host)
	bootstrap.FormatUrl(storeLocation)

	suffix := s.storeImage(*storeLocation, image.Name, *image.Content)
	url := *host + *softLink + suffix
	beego.Alert(url)
	backup := *host + "backup/" + suffix
	str := `ZXCVBNMASDFGHJKLQWERTYUIOPzxcvbnmasdfghjklqwertyuiop1234567890`
	randomStr := bootstrap.GetRandomString(16, str)
	return ImageReturn{
		Url:    url,
		Delete: randomStr,
		Path:   *storeLocation + suffix,
		Other:  backup,
		ID:     12,
	}, nil
}

//储存图片
func (s *Local) storeImage(path string, n string, fileContent []byte) string {
	nowTime := beego.Date(time.Now(), "Y/m/d/")
	str := `ZXCVBNMASDFGHJKLQWERTYUIOPzxcvbnmasdfghjklqwertyuiop1234567890`
	suffix := bootstrap.GetRandomString(16, str) + "." + bootstrap.GetImageSuffix(n)
	dir := path + nowTime
	file := dir + suffix
	bootstrap.CheckPath(dir)
	var f *os.File
	f, err := os.Create(file)
	if err != nil {
		beego.Alert("File Create Error:", err)
	}
	w := bufio.NewWriter(f)
	_, err = w.Write(fileContent)
	if err != nil {
		beego.Alert("File Create Error:", err)
	}
	w.Flush()
	f.Close()

	return nowTime + suffix
}
