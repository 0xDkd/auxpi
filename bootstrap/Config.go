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
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/auxpi/auxpiAll"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
)

type AuxpiConfig struct {
}

var Cache = initCache()
//var Cache, _ = cache.NewCache("memory", `{"interval":60}`)

var SiteConfig *auxpi.SiteConfig
var InitConfig auxpi.SiteConfig

var Site auxpi.SiteBase

type jsonStruct struct {
}

func (jst *jsonStruct) load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

//返回 Config 里面的 数据
func config() *auxpi.SiteConfig {
	//如果开启了 Config 缓存则尝试从缓存中检索
	cacheConfig := Cache.Get("SiteConfig")
	if cacheConfig != nil {
		config, _ := cacheConfig.(*auxpi.SiteConfig)
		return config
	}
	reader := &jsonStruct{}
	config := &auxpi.SiteConfig{}
	configDir := "conf/siteConfig.json"
	reader.load(configDir, config)
	//缓存到内存
	if config.CacheConfig {
		Cache.Put("SiteConfig", config, time.Second*3600)
	}
	return config

}

//配置重新载入内存
func Reload() {
	Cache.Delete("SiteConfig")
	SiteConfig = config()
}

//在有锁的情况下强制更新配置文件
func ReGenerate() {
	configJson, err := SiteConfig.MarshalJSON()
	if err != nil {
		panic(err)
	}
	baseDir := "conf/"
	var f *os.File
	//创建 config 文件并且写入内容
	configDir := baseDir + "siteConfig.json"
	f, err = os.Create(configDir)
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(configJson))
	w.Flush()
	f.Close()
	Reload()
}

//传入并且重新生成
func ReGenerateByInput(siteConfig auxpi.SiteConfig) error {
	configJson, err := siteConfig.MarshalJSON()
	if err != nil {
		return err
	}
	baseDir := "conf/"
	var f *os.File
	//创建 config 文件并且写入内容
	configDir := baseDir + "siteConfig.json"
	f, err = os.Create(configDir)
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(configJson))
	if err != nil {
		beego.Alert(err)
	}
	w.Flush()
	f.Close()
	Reload()

	return nil
}

//初始化的时候检测是否进行安装，生成对应的 lock 文件,生成配置的 json
func init() {
	baseDir := "conf/"
	lockDir := baseDir + "install.lock"
	_, err := os.Stat(lockDir)
	if err == nil {
		SiteConfig = config()
		return
	}
	if os.IsNotExist(err) {
		var f *os.File
		ramdomString := `0123456789abcdef!@#$%^&*()__+ghijklmnop?></qrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
		//Site Init
		siteconfig := auxpi.SiteConfig{}
		siteconfig.SiteName = "AuXpI API 图床"
		siteconfig.SiteUrl = "http://yoursite.com/"
		siteconfig.SiteFooter = "新一代图床"
		siteconfig.Logo = "/static/app/images/logo.jpg"
		siteconfig.SiteUploadMaxSize = 5
		siteconfig.SiteUploadMaxNumber = 10

		//JWT
		siteconfig.JwtSecret = GetRandomString(16, ramdomString)
		//Jwt 过期时间 单位 小时
		siteconfig.JwtDueTime = 3
		siteconfig.AuxpiSalt = GetRandomString(16, ramdomString)

		siteconfig.ApiDefault = "SouGou"
		siteconfig.CacheConfig = false

		//upload way Init
		//本地储存
		siteconfig.SiteUploadWay.LocalStore.Status = true
		siteconfig.SiteUploadWay.LocalStore.Link = "/images"
		siteconfig.SiteUploadWay.LocalStore.StorageLocation = "public/upload"

		//新浪图床配置
		siteconfig.SiteUploadWay.SinaAccount.Status = false
		siteconfig.SiteUploadWay.SinaAccount.UserName = ""
		siteconfig.SiteUploadWay.SinaAccount.PassWord = ""
		siteconfig.SiteUploadWay.SinaAccount.ResetSinaCookieTime = 3600
		siteconfig.SiteUploadWay.SinaAccount.DefultPicSize = "large"

		//Flickr Init
		siteconfig.SiteUploadWay.FlickrAccount.Api_key = ""
		siteconfig.SiteUploadWay.FlickrAccount.Api_secret = ""
		siteconfig.SiteUploadWay.FlickrAccount.Oauth_token = ""
		siteconfig.SiteUploadWay.FlickrAccount.Oauth_token_secret = ""
		siteconfig.SiteUploadWay.FlickrAccount.Id = ""
		siteconfig.SiteUploadWay.FlickrAccount.DefaultSize = "h"

		//Db Init
		siteconfig.DbOption.UseDb = true
		siteconfig.DbOption.DbType = "mysql"
		siteconfig.DbOption.DbHost = "127.0.0.1:3306"
		siteconfig.DbOption.DbName = "auxpi"
		siteconfig.DbOption.DbUser = "root"
		siteconfig.DbOption.DbPass = "root"
		siteconfig.DbOption.TablePrefix = "auxpi_"

		//Auxpi Info
		siteconfig.AuxpiInfo.Author = "aimerforreimu"
		siteconfig.AuxpiInfo.Branch = "dev"
		siteconfig.AuxpiInfo.Version = "2.0"
		siteconfig.AuxpiInfo.Repositories = "https://github.com/aimerforreimu/AUXPI"

		configJson, err := siteconfig.MarshalJSON()
		if err != nil {
			panic(err)
		}
		//创建lock文件
		f, err = os.Create(lockDir)
		//创建 config 文件并且写入内容
		configDir := baseDir + "siteConfig.json"
		f, err = os.Create(configDir)
		w := bufio.NewWriter(f)
		_, err = w.WriteString(string(configJson))
		w.Flush()
		f.Close()
		SiteConfig = config()
	}

}

func initCache() cache.Cache {
	driver := beego.AppConfig.String("drive")
	driver = strings.ToLower(driver)
	switch driver {
	case "redis":
		c, err := cache.NewCache("redis", `{"key":"`+
			beego.AppConfig.String("redisCollection")+`","conn":":`+
			beego.AppConfig.String("redisPort")+`","dbNum":"0","password":"`+
			beego.AppConfig.String("redisPassword")+`"}`)
		if err != nil {
			panic(err)
		}
		return c
	case "memcache":
		c, err := cache.NewCache("memcache", `{"conn":"`+
			beego.AppConfig.String("memcacheConn")+`"}`)
		if err != nil {
			panic(err)
		}
		return c
	case "memory":
		c, err := cache.NewCache("memory", `{"interval":60}`)
		if err != nil {
			panic(err)
		}
		return c
	case "file":
		c, err := cache.NewCache("file", `{"CachePath":"`+
			beego.AppConfig.String("cachePath")+`","FileSuffix":"`+
			beego.AppConfig.String("fileSuffix")+`","DirectoryLevel":"`+
			beego.AppConfig.String("directoryLevel")+`","EmbedExpiry":"`+
			beego.AppConfig.String("EmbedExpiry")+`"}`)
		if err != nil {
			panic(err)
		}
		return c
	default:
		c, err := cache.NewCache("memory", `{"interval":60}`)
		if err != nil {
			panic(err)
		}
		return c
		
	}
}
