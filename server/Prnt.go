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
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type Prnt struct {
	FileLimit []string
	MaxSize   int
}

func (s *Prnt) Upload(image *ImageParam) (ImageReturn, error) {
	url := "https://prntscr.com/upload.php"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "image",
		Value: *image.Content,
		Type:  image.Type,
	}
	var header = make(map[string]string)
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:65.0) Gecko/20100101 Firefox/65.0"
	data := tools.FormPost(file, url, header)
	var d = map[string]string{}
	json.Unmarshal([]byte(data), &d)

	req, err := http.NewRequest("GET", d["data"], nil)
	if err != nil {
		return ImageReturn{}, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:65.0) Gecko/20100101 Firefox/65.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ImageReturn{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ImageReturn{}, err
	}

	var re = regexp.MustCompile(`(?m)<meta property="og:image" content="(.*?)"/>`)
	url = re.FindAllStringSubmatch(string(body), -1)[0][1]
	beego.Alert(url)
	return ImageReturn{
		Url: url,
		ID:  7,
	}, nil
}

func (s *Prnt) UploadToPrnt(img []byte, imgInfo string, imgType string) string {
	url := "https://prntscr.com/upload.php"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "image",
		Value: img,
		Type:  imgType,
	}
	var header = make(map[string]string)
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:65.0) Gecko/20100101 Firefox/65.0"
	data := tools.FormPost(file, url, header)
	var d = map[string]string{}
	json.Unmarshal([]byte(data), &d)

	req, _ := http.NewRequest("GET", d["data"], nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:65.0) Gecko/20100101 Firefox/65.0")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var re = regexp.MustCompile(`(?m)<meta property="og:image" content="(.*?)"/>`)
	url = re.FindAllStringSubmatch(string(body), -1)[0][1]
	beego.Alert(url)
	return url
}
