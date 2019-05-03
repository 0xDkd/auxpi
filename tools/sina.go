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
	"fmt"
	"hash/crc32"
)

func CheckPid(pid string, imgType string, size string) string {
	if pid == "" {
		return ""
	}
	sinaNumber := fmt.Sprint((crc32.ChecksumIEEE([]byte(pid)) & 3) + 1)
	n := len(imgType)
	rs := []rune(imgType)
	suffix := string(rs[6:n])
	if suffix != "gif" {
		suffix = "jpg"
	}
	sinaUrl := "https://ww" + sinaNumber + ".sinaimg.cn/" + size + "/" + pid + "." + suffix
	return sinaUrl

}
