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
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/bootstrap"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
	"github.com/auxpi/tools"
	"github.com/pkg/errors"
)

type Sina struct {
	FileLimit []string
	MaxSize   int
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

func (s *Sina) Upload(image *ImageParam) (ImageReturn, error) {
	var sinaAccount = auxpi.SinaAccount{}
	err := sinaAccount.UnmarshalJSON([]byte(models.GetOption("sina", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("SERVER", err)
	}
	if sinaAccount.PassWord == "" || sinaAccount.UserName == "" {
		err = errors.New("Sina Account is null")
		return ImageReturn{}, err
	}

	durl := "http://picupload.service.weibo.com/interface/pic_upload.php" +
		"?ori=1&mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog"
	imgStr := base64.StdEncoding.EncodeToString(*image.Content)
	//构造 http 请求
	postData := make(url.Values)
	postData["b64_data"] = []string{imgStr}
	client := &http.Client{}
	request, err := http.NewRequest("POST", durl, strings.NewReader(postData.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//设置 cookie
	unCookies := s.Login(sinaAccount.UserName, sinaAccount.PassWord)
	//需要进行断言转换
	cookies, ok := unCookies.([]*http.Cookie)
	if !ok {
		panic(ok)
	}
	for _, value := range cookies {
		request.AddCookie(value)
	}
	resp, err := client.Do(request)

	if err != nil {
		return ImageReturn{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ImageReturn{}, err
	}
	var proxy = ""
	if sinaAccount.Proxy.Status {
		proxy = sinaAccount.Proxy.Node
	}
	sinaUrl := s.getSinaUrl(body, image.Type)
	if sinaUrl != "" {
		sinaUrl = proxy + sinaUrl
	}
	return ImageReturn{
		Url: sinaUrl,
		ID:  2,
	}, nil

}

//新浪图床登录
func (s *Sina) Login(name string, pass string) interface{} {
	url := "https://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.15)&_=1403138799543"
	userInfo := make(map[string]string)
	userInfo["UserName"] = bootstrap.Encode(base64.StdEncoding, name)
	userInfo["PassWord"] = pass
	cookie := s.getCookies(url, userInfo)
	return cookie
}

//获取新浪图床 Cookie
func (s *Sina) getCookies(durl string, data map[string]string) interface{} {
	//尝试从缓存里面获取 Cookie
	if bootstrap.Cache.Get("SinaCookies") != nil {
		//beego.Alert("cached")
		return bootstrap.Cache.Get("SinaCookies")
	}
	postData := make(url.Values)
	postData["entry"] = []string{"sso"}
	postData["gateway"] = []string{"1"}
	postData["from"] = []string{"null"}
	postData["savestate"] = []string{"30"}
	postData["uAddicket"] = []string{"0"}
	postData["pagerefer"] = []string{""}
	postData["vsnf"] = []string{"1"}
	postData["su"] = []string{data["UserName"]} //UserName
	postData["service"] = []string{"sso"}
	postData["sp"] = []string{data["PassWord"]} //PassWord
	postData["sr"] = []string{"1920*1080"}
	postData["encoding"] = []string{"UTF-8"}
	postData["cdult"] = []string{"3"}
	postData["domain"] = []string{"sina.com.cn"}
	postData["prelt"] = []string{"0"}
	postData["returntype"] = []string{"TEXT"}
	client := &http.Client{}
	request, err := http.NewRequest("POST", durl, strings.NewReader(postData.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)
	if err != nil {
		auxpiLog.SetAWarningLog("SERVER", err)
		return ""
	}
	body, _ := ioutil.ReadAll(resp.Body)
	sinaError := &auxpi.SinaError{}
	err = sinaError.UnmarshalJSON(body)
	if err != nil {
		auxpiLog.SetAWarningLog("SERVER", err)
		return ""
	}
	if sinaError.Retcode == "101" {
		logs.Alert("新浪图床上传错误:" + sinaError.Reason)
	}
	defer resp.Body.Close()
	cookie := resp.Cookies()
	//缓存 Cookie 缓存一个小时
	bootstrap.Cache.Put("SinaCookies", cookie, time.Second*3600)
	return cookie
}

//获取 Sina 图床 URL
func (s *Sina) getSinaUrl(body []byte, imgType string) string {
	var sinaAccount = auxpi.SinaAccount{}
	err := sinaAccount.UnmarshalJSON([]byte(models.GetOption("sina", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("SERVER", err)
	}

	str := string(body)
	//正则获取
	pat := "({.*)"
	res := regexp.MustCompile(pat)
	jsons := res.FindAllStringSubmatch(str, -1)
	msg := auxpi.SinaMsg{}
	//解析 json 到 struct
	msg.UnmarshalJSON([]byte(jsons[0][1]))
	//验证 pid 的合法性
	pid := msg.Data.Pics.Pic_1.Pid
	sinaUrl := tools.CheckPid(pid, imgType, sinaAccount.DefultPicSize)
	if sinaUrl == "" {
		return ""
	}
	return sinaUrl
}

////上传图片
//func (s *Sina) UpLoadToSina(img []byte, imgType string) string {
//	durl := "http://picupload.service.weibo.com/interface/pic_upload.php" +
//		"?ori=1&mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog"
//	imgStr := base64.StdEncoding.EncodeToString(img)
//	//构造 http 请求
//	postData := make(url.Values)
//	postData["b64_data"] = []string{imgStr}
//	client := &http.Client{}
//	request, err := http.NewRequest("POST", durl, strings.NewReader(postData.Encode()))
//	if err != nil {
//		fmt.Println(err)
//	}
//	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	//设置 cookie
//	uncooikes := s.Login(bootstrap.SiteConfig.SiteUploadWay.SinaAccount.UserName, bootstrap.SiteConfig.SiteUploadWay.SinaAccount.PassWord)
//	//需要进行断言转换
//	cookies, ok := uncooikes.([]*http.Cookie)
//	if !ok {
//		panic(ok)
//	}
//	for _, value := range cookies {
//		request.AddCookie(value)
//	}
//	resp, err := client.Do(request)
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	fmt.Println(string(body))
//	return s.getSinaUrl(body, imgType)
//}


