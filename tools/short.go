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
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

const (
	VAL   = 0x3FFFFFFF
	INDEX = 0x0000003D
)

var (
	alphabet = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func ToShort(longURL string) ([4]string, error) {
	md5Str := getMd5Str(longURL)
	//var hexVal int64
	var tempVal int64
	var result [4]string
	var tempUri []byte
	for i := 0; i < 4; i++ {
		tempSubStr := md5Str[i*8 : (i+1)*8]
		hexVal, err := strconv.ParseInt(tempSubStr, 16, 64)
		if err != nil {
			return result, nil
		}
		tempVal = int64(VAL) & hexVal
		var index int64
		tempUri = []byte{}
		for i := 0; i < 6; i++ {
			index = INDEX & tempVal
			tempUri = append(tempUri, alphabet[index])
			tempVal = tempVal >> 5
		}
		result[i] = string(tempUri)
	}
	return result, nil
}

/** generate md5 checksum of URL in hex format **/
func getMd5Str(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	c := m.Sum(nil)
	return hex.EncodeToString(c)
}
