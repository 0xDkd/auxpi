package bootstrap

import (
	"auxpi/auxpiAll"
	"bufio"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type AuxpiConfig struct {
}

var cCache, _ = cache.NewCache("memory", `{"interval":3600}`)

var SiteConfig *auxpi.SiteConfig

type JsonStruct struct {
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
func Config() *auxpi.SiteConfig {
	//如果开启了 Config 缓存则尝试从缓存中检索
	cacheConfig := cCache.Get("SiteConfig")
	if cacheConfig != nil {
		config, _ := cacheConfig.(*auxpi.SiteConfig)
		return config
	}
	reader := &JsonStruct{}
	config := &auxpi.SiteConfig{}
	configDir := GetPath() + "/conf/siteConfig.json"
	reader.Load(configDir, config)
	//缓存到内存
	if config.CacheConfig {
		cCache.Put("SiteConfig", config, time.Second*3600)
	}
	return config
}

//初始化的时候检测是否进行安装，生成对应的 lock 文件,生成配置的 json
func init() {
	baseDir := GetPath() + "/conf/"
	lockDir := baseDir + "install.lock"
	_, err := os.Stat(lockDir)
	if err == nil {
		SiteConfig = Config()
		return
	}
	if os.IsNotExist(err) {
		var f *os.File
		siteconfig := auxpi.SiteConfig{}
		siteconfig.SiteName = "AuXpI API 图床"
		siteconfig.SiteUrl = "/"
		siteconfig.SiteFooter = ""
		siteconfig.SiteUpLoadMaxSize = 5
		siteconfig.SiteUploadMaxNumber = 10
		siteconfig.OpenApiUpLoad = true
		siteconfig.JwtSecret = GetRandomString(16)
		siteconfig.AuxpiSalt = GetRandomString(16)
		siteconfig.ApiToken = ""
		siteconfig.ApiDefault = "SouGou"
		siteconfig.CacheConfig = false
		siteconfig.SiteUploadWay.OpenSinaPicStore = false
		siteconfig.SiteUploadWay.LocalStore = false
		siteconfig.SiteUploadWay.SinaAccount.UserName = ""
		siteconfig.SiteUploadWay.SinaAccount.PassWord = ""
		siteconfig.SiteUploadWay.SinaAccount.ResetSinaCookieTime = 3600
		siteconfig.SiteUploadWay.SinaAccount.DefultPicSize = "large"
		siteconfig.DbOption.UseDb = true
		siteconfig.DbOption.DbType = "mysql"
		siteconfig.DbOption.DbHost = "127.0.0.1:3306"
		siteconfig.DbOption.DbName = "auxpi"
		siteconfig.DbOption.DbUser = "root"
		siteconfig.DbOption.DblPass = "root"
		siteconfig.DbOption.TablePrefix = "auxpi_"
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
		SiteConfig = Config()
	}

}

func GetRandomString(l int) string {
	str := `0123456789abcdef!@#$%^&*()__+ghijklmnop?></qrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

