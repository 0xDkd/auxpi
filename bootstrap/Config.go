package bootstrap

import (
	"auxpi/auxpiAll"
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/astaxie/beego/cache"
)

type AuxpiConfig struct {
}

var cCache, _ = cache.NewCache("memory", `{"interval":3600}`)

var SiteConfig *auxpi.SiteConfig

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
	cacheConfig := cCache.Get("SiteConfig")
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
		cCache.Put("SiteConfig", config, time.Second*3600)
	}
	return config

}

//配置重新载入内存
func Reload() {
	cCache.Delete("SiteConfig")
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
		siteconfig.OpenApiUpLoad = true

		//JWT
		siteconfig.JwtSecret = GetRandomString(16, ramdomString)
		//Jwt 过期时间 单位 小时
		siteconfig.JwtDueTime = 3
		siteconfig.AuxpiSalt = GetRandomString(16, ramdomString)

		siteconfig.ApiToken = ""
		siteconfig.ApiDefault = "SouGou"
		siteconfig.CacheConfig = false

		//upload way Init
		//本地储存
		siteconfig.SiteUploadWay.LocalStore.Open = true
		siteconfig.SiteUploadWay.LocalStore.Link = "/images"
		siteconfig.SiteUploadWay.LocalStore.StorageLocation = "public/upload"

		//新浪图床配置
		siteconfig.SiteUploadWay.OpenSinaPicStore = false
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
