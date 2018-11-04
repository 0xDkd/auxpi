package bootstrap

import (
	"bufio"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"io/ioutil"
	"os"
	"time"
)

type AuxpiConfig struct {
	
} 
type SiteConfig struct {
	//站点名称
	SiteName string `json:"site_name"`
	//底部信息
	SiteFooter string `json:"site_footer"`
	//网站链接
	SiteUrl string `json:"site_url"`
	//最大上传的图片个数
	SiteUploadMaxNumber int `json:"site_upload_max_number"`
	//最大图片规格 MB
	SiteUpLoadMaxSize int `json:"site_up_load_max_size"`
	//图床储存的一些配置
	SiteUploadWay UploadConfig `json:"site_upload_way"`
}

type UploadConfig struct {
	//是否开启本地上传
	LocalStore bool `json:"local_store"`
	//是否开启微博图床
	OpenSinaPicStore bool `json:"open_sina_pic_store"`
	//Sina Account
	SinaAccount Account `json:"sina_account"`
	//
}
type Account struct {
	//用户名
	UserName string `json:"user_name"`
	//密码
	PassWord string `json:"pass_word"`
	//新浪 Cookie 更新的频率,默认为3600s ,单位 s
	ResetSinaCookieTime int `json:"reset_sina_cookie_time"`
	//新浪图床默认使用的尺寸大小 square,thumb150,orj360,orj480,mw690,mw1024,mw2048,small,bmiddle,large 、默认为large
	DefultPicSize string `json:"defult_pic_size"`
}

var cCache, _ = cache.NewCache("memory", `{"interval":3600}`)

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
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
func Config() *SiteConfig {
	//尝试从缓存中检索
	cacheConfig := cCache.Get("SiteConfig")
	if cacheConfig != nil {
		config, _ := cacheConfig.(*SiteConfig)
		return config
	}
	reader := &JsonStruct{}
	config := &SiteConfig{}
	configDir := GetPath() + "/conf/siteConfig.json"
	reader.Load(configDir, config)
	//缓存到内存
	cCache.Put("SiteConfig", config, time.Second*3600)
	return config
}
//func (this *AuxpiConfig)GetConfig(key string)  {
//	config :=Config()
//}

//初始化的时候检测是否进行安装，生成对应的 lock 文件,生成配置的 json

func init() {
	baseDir := GetPath() + "/conf/"
	lockDir := baseDir + "install.lock"
	_, err := os.Stat(lockDir)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		var f *os.File
		siteconfig := SiteConfig{}
		siteconfig.SiteName = "BusterApi 图床"
		siteconfig.SiteUrl = "/"
		siteconfig.SiteFooter = "你好世界"
		siteconfig.SiteUpLoadMaxSize = 5
		siteconfig.SiteUploadMaxNumber = 10
		siteconfig.SiteUploadWay.OpenSinaPicStore = false
		siteconfig.SiteUploadWay.LocalStore = false
		siteconfig.SiteUploadWay.SinaAccount.UserName = ""
		siteconfig.SiteUploadWay.SinaAccount.PassWord = ""
		siteconfig.SiteUploadWay.SinaAccount.ResetSinaCookieTime = 3600
		siteconfig.SiteUploadWay.SinaAccount.DefultPicSize = "large"
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
	}

}
