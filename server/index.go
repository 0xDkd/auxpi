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
)

//上传者
type Uploader interface {
	Upload(image *ImageParam) (ImageReturn, error)
}

//统一化参数结构体
//TODO: timeout set
type ImageParam struct {
	Name    string
	Info    string
	Type    string
	Content *[]byte
}

//统一化返回结构体
type ImageReturn struct {
	Url    string
	Delete string
	Path   string
	ID     int
	Other  interface{}
}

//uploader Manager
type Manager struct {
	List map[string]Uploader
}

//uploader Client
type Client struct {
	Uploader Uploader    //上传者实例
	Resp     ImageReturn //返回信息
	Error    error
}

var manager = initManager()
var rootManager = initRootManager()

////获取这张图片的 MineType [image/png]
//func (u *ImageParam) GetMimeType() string {
//	return u.FileHeader.Header.Get("Content-Type")
//}
//
////获取这张图片的总信息
//func (u *ImageParam) GetInfo() string {
//	return u.FileHeader.Header.Get("Content-Disposition")
//}

//获取文件的 []byte
//func (u *ImageParam) GetFileContent() ([]byte, error) {
//	b := make([]byte, u.FileHeader.Size)
//	file, err := u.FileHeader.Open()
//	_, err = file.Read(b)
//	if err != nil {
//		return []byte{}, err
//	}
//	return b, nil
//}

//初始化Manager
func initManager() *Manager {
	u := make(map[string]Uploader)
	var h = Manager{
		List: u,
	}
	doRegister(&h)
	return &h
}

//初始化根节点 Manager
func initRootManager() *Manager {
	u := make(map[string]Uploader)
	var h = Manager{
		List: u,
	}
	//Gitee
	h.RegisterUploader("gitee", &Gitee{})
	//GitHub
	h.RegisterUploader("github", &Github{})
	//Local
	h.RegisterUploader("local", &Local{})
	//TODO:OOS

	//Other :Flickr | Imgur
	h.RegisterUploader("flickr", &Flickr{})
	h.RegisterUploader("imgur", &Imgur{})

	return &h

}

//注册 Uploader 到 Manager 中
func doRegister(c *Manager) {
	c.RegisterUploader("ali", &Ali{})
	c.RegisterUploader("cc", &CC{})
	c.RegisterUploader("flickr", &Flickr{})
	c.RegisterUploader("imgur", &Imgur{})
	c.RegisterUploader("jd", &Jd{})
	c.RegisterUploader("juejin", &JueJin{})
	c.RegisterUploader("local", &Local{})
	c.RegisterUploader("ooxx", &OOXX{})
	c.RegisterUploader("prnt", &Prnt{})
	c.RegisterUploader("sina", &Sina{})
	c.RegisterUploader("smms", &Smms{})
	c.RegisterUploader("souhu", &SouHu{})
	c.RegisterUploader("suning", &SuNing{})
	c.RegisterUploader("toutiao", &TouTiao{})
	c.RegisterUploader("vim", &VimCN{})
	c.RegisterUploader("xiaomi", &XiaoMi{})
	c.RegisterUploader("sougou", &SouGou{})
	c.RegisterUploader("gitee", &Gitee{})
	c.RegisterUploader("github", &Github{})
	c.RegisterUploader("neteasy", &NetEasy{})
}

//生成一个 Upload Client
func NewClient(name string) *Client {
	var a = Client{}
	if manager.List[name] == nil {
		err := errors.New("[uploader]: Can not new uploader named " + name)
		a.Error = err
		return &a
	}
	a.Uploader = manager.List[name]
	return &a
}

//生成多个 Client 使用速度排名
func NewClientsOrderBySpeed(auth map[string]string, other ...interface{}) []Client {
	var a = make([]Client, 11)
	var rankList = [...]string{
		"ali",
		"souhu",
		"juejin",
		"sina",
		"suning",
		"xiaomi",
		"toutiao",
		"jd",
		"smms",
	}
	for key, value := range rankList {
		a[key].Uploader = manager.List[value]
	}
	return a
}

//Client Do
func (c *Client) Do(image *ImageParam) *Client {
	c.Resp, c.Error = c.Uploader.Upload(image)
	if c.Error != nil {
		return c
	}
	return c
}

//注册 Uploader 到 Manager 内
func (c *Manager) RegisterUploader(name string, uploader Uploader) *Manager {
	if uploader == nil {
		panic("[uploader]: Register uploader is nil")
	}
	if _, u := c.List[name]; u {
		panic("[uploader]: Register called twice for uploader " + name)
	}
	c.List[name] = uploader

	return c
}
