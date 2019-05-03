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

package tools

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego/logs"
)

func CheckStatus(url string) bool {
	if url == "" {
		return false
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 8.1; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.143 Crosswalk/24.53.595.0 XWEB/358 MMWEBSDK/23 Mobile Safari/537.36 MicroMessenger/6.7.2.1340(0x2607023A) NetType/4G Language/zh_CN")

	resp, err := client.Do(req)
	if err != nil {
		logs.Error("[CHECK PIC RESP ERROR]: ", err)
		return false
	}

	err = resp.Body.Close()
	if err != nil {
		logs.Error("[CHECK PIC RESP ERROR]: ", err)
		return false
	}
	//检测请求头
	if resp.StatusCode != 200 {
		return false
	}
	//解析响应头是否是png/jpg/gif/bmp 等资源
	//是否返回的是图片资源
	info := resp.Header.Get("Content-Type")
	if strings.HasPrefix(info, "image") || strings.HasSuffix(info, "octet-stream") {
		return true
	}
	return false
}
