// Copyright (c) 2019 aimerforreimu. All Rights Reserved.
// Use of this source code is governed by a GNU GENERAL PUBLIC
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

package utils

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/auxpi/auxpiAll"
)

func FormPost(file *auxpi.FormFile, url string, header map[string]string) string {
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"" + file.Key + "\"; filename=\"" + file.Name + "\"\r\nContent-Type: " + file.Type + "\r\n\r\n\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
