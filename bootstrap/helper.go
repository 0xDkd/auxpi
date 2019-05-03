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

package bootstrap

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gofrs/uuid"
)

func GetPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func Decode(enc *base64.Encoding, str string) string {
	data, err := enc.DecodeString(str)

	if err != nil {
		panic(err)
	}
	return string(data)
}

func Encode(enc *base64.Encoding, str string) string {
	bData := []byte(str)
	data := enc.EncodeToString(bData)
	return string(data)
}

func FormatSoftLink(url *string) {
	n := len(*url)
	rs := []rune(*url)
	s := string(rs[n-1 : n])
	if s == "/" {
		*url = string(rs[0 : n-1])
	}
	s = string(rs[0:1])
	if s != "/" {
		*url = "/" + string(rs[0:n])
	}
}

func FormatStoreLocation(location *string) {
	n := len(*location)
	rs := []rune(*location)
	s := string(rs[n-1 : n])
	if s != "/" {
		*location += "/"
	}
	s = string(rs[0:1])
	if s == "/" {
		*location = string(rs[1:n])
	}
}

func GetRandomString(l int, str string) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//检查路径并且创建
func CheckPath(path string) {
	//base := bootstrap.GetPath()
	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, 0775)
		if err != nil {
			logs.Alert("Create Images store unsuccessful:", err)
			return
		}
	}
}

//格式化 url
func FormatUrl(url *string) {
	n := len(*url)
	rs := []rune(*url)
	s := string(rs[n-1 : n])
	if s != "/" {
		*url += "/"
	}
	s = string(rs[0:1])
	if s == "/" {
		*url = string(rs[1:n])
	}
}

//获取图片后G缀
func GetImageSuffix(name string) (suffix string) {
	ss := strings.Split(name, ".")
	l := len(ss)
	//选取最后一个
	return ss[l-1]
}

//生成带有年月日的图片名称 2019/4/29/{hash}.suffix
func GenerateImageName(name string) string {
	nowTime := beego.Date(time.Now(), "Y/m/d/")
	suffix := GenerateNormalString(16) + "." + GetImageSuffix(name)
	return nowTime + suffix
}

//生成 a-zA-Z0-9 的指定长度的随机字符串
func GenerateNormalString(l int) string {
	return GetRandomString(l, "zxcvbnmasdfghjklqwertyuiop1234567890ZXCVBNMLKJHGFDSAQWERTYUIOP")
}

//生成 md5 加密以后的 uuid（唯一字符串）
func GenerateUniqueString() string {
	u, _ := uuid.NewV4()
	id := u.Bytes()
	h := md5.New()
	h.Write(id)
	return hex.EncodeToString(h.Sum(nil))
}
