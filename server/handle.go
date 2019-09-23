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

//
//import (
//	"errors"
//)
//
//type UploadHandle struct {
//	List   map[string]Uploader
//	Before func(...interface{})
//	After  func(...interface{})
//	Resp   []ImageReturn
//}
//
//func NewUploadHandle(name string) (*UploadHandle, error) {
//	all := initAllUploadHandle()
//	for key := range all.List {
//		if name == key {
//			a := initUploadHandle()
//			a.List[key] = all.List[key]
//
//			return a, nil
//		}
//	}
//	err := errors.New("Can not make a handle named " + name)
//	return &UploadHandle{}, err
//
//}
//
////func Bnew() (*UploadHandle, error) {
////	return initAllUploadHandle(), nil
////}
//
//func initUploadHandle() *UploadHandle {
//	u := make(map[string]Uploader, 1)
//	var h = UploadHandle{
//		List: u,
//		Resp: make([]ImageReturn, len(u), 10),
//	}
//	return &h
//}
//
//func initAllUploadHandle() *UploadHandle {
//	u := make(map[string]Uploader, 200)
//	u["ali"] = &Ali{}
//	u["cc"] = &CC{}
//	u["flickr"] = &Flickr{}
//	u["imgur"] = &Imgur{}
//	u["jd"] = &Jd{}
//	u["juejin"] = &JueJin{}
//	u["local"] = &Local{}
//	u["netEasy"] = &NetEasy{}
//	u["ooxx"] = &OOXX{}
//	u["prnt"] = &Prnt{}
//	u["sina"] = &Sina{}
//	u["smms"] = &Smms{}
//	u["souhu"] = &SouHu{}
//	u["suning"] = &SuNing{}
//	u["toutiao"] = &TouTiao{}
//	u["vim"] = &VimCN{}
//	u["xiaomi"] = &XiaoMi{}
//	var h = UploadHandle{
//		List: u,
//	}
//	return &h
//}
//
//func (c *UploadHandle) AddUploader(uploaders map[string]Uploader) *UploadHandle {
//	if len(uploaders) == 0 {
//		return c
//	}
//	for key, uploader := range uploaders {
//		if c.List[key] != nil {
//			c.List[key] = uploader
//		}
//	}
//	return c
//}
//
//func (c *UploadHandle) Do(image *ImageParam) (ImageReturn, error) {
//	for key := range c.List {
//		return c.List[key].Upload(image)
//	}
//	return ImageReturn{}, errors.New("123")
//}
//
////func (c *UploadHandle) Bdo(image *ImageParam) {
////	var ch = make(chan ImageReturn)
////	var e = make(chan error)
////	var who = make(chan string)
////	for key := range c.List {
////		go func() {
////			content, err := c.List[key].Upload(image)
////			who <- key
////			ch <- content
////			e <- err
////		}()
////	}
////	for range c.List {
////		beego.Alert(<-who)
////		beego.Alert(<-ch)
////		if <-e != nil {
////			beego.Alert(<-e)
////		}
////	}
////
////}
