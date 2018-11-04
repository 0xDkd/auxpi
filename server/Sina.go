package server

import (
	"auxpi/bootstrap"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type Sina struct {
}

//easyjson:json
type SinaMsg struct {
	Code string   `json:"code"`
	Data SinaData `json:"data"`
}

//easyjson:json
type SinaData struct {
	Count int      `json:"count"`
	Data  string   `json:"data"`
	Pics  SinaPics `json:"pics"`
}

//easyjson:json
type SinaPics struct {
	Pic_1 picInfo `json:"pic_1"`
}

//easyjson:json
type picInfo struct {
	Width  int    `json:"width"`
	Size   int    `json:"size"`
	Ret    int    `json:"ret"`
	Height int    `json:"height"`
	Name   string `json:"name"`
	Pid    string `json:"pid"`
}

var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}
var memoryCache, _ = cache.NewCache("memory", `{"interval":3600}`)
//获取 config 的配置
var siteConfig = bootstrap.Config()

func (this *Sina) Login(name string, pass string) interface{} {
	url := "https://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.15)&_=1403138799543"
	userInfo := make(map[string]string)
	userInfo["UserName"] = bootstrap.Encode(base64.StdEncoding, name)
	userInfo["PassWord"] = pass
	cookie := this.getCookies(url, userInfo)
	return cookie
}

//返回综合 []*httpCookie
func (this *Sina) getCookies(durl string, data map[string]string) (interface{}) {
	//尝试从缓存里面获取 Cookie
	if memoryCache.Get("SinaCookies") != nil {
		//beego.Alert("cached")
		return memoryCache.Get("SinaCookies")
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
	defer resp.Body.Close()
	cookie := resp.Cookies()
	//缓存 Cookie 缓存一个小时
	memoryCache.Put("SinaCookies", cookie, time.Second*3600)
	return cookie
}

//上传图片
func (this *Sina) UpLoadToSina(img []byte, imgType string) string {
	if siteConfig.SiteUploadWay.OpenSinaPicStore == false {
		return ""
	}
	durl := "http://picupload.service.weibo.com/interface/pic_upload.php" +
		"?mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog"
	imgStr := base64.StdEncoding.EncodeToString(img)
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
	uncooikes := this.Login(siteConfig.SiteUploadWay.SinaAccount.UserName, siteConfig.SiteUploadWay.SinaAccount.PassWord)
	//需要进行断言转换
	cookies, ok := uncooikes.([]*http.Cookie)
	if !ok {
		panic(ok)
	}
	for _, value := range cookies {
		request.AddCookie(value)
	}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return this.getSinaUrl(body, imgType)
}

func (this *Sina) getSinaUrl(body []byte, imgType string) string {
	str := string(body)
	//正则获取
	pat := "({.*)"
	check := "[a-zA-Z0-9]{32}"
	res := regexp.MustCompile(pat)
	rule := regexp.MustCompile(check)
	jsons := res.FindAllStringSubmatch(str, -1)
	//beego.Alert(rule, jsons)
	msg := SinaMsg{}
	//解析 json 到 struct
	msg.UnmarshalJSON([]byte(jsons[0][1]))
	//验证 pid 的合法性
	pid := msg.Data.Pics.Pic_1.Pid
	if rule.MatchString(pid) {
		sinaNumber := fmt.Sprint((crc32.ChecksumIEEE([]byte(pid)) & 3) + 1)
		//从配置文件中获取
		size := siteConfig.SiteUploadWay.SinaAccount.DefultPicSize
		n := len(imgType)
		rs := []rune(imgType)
		suffix := string(rs[6:n])
		if suffix != "gif" {
			suffix = "jpg"
		}
		sinaUrl := "https://ws" + sinaNumber + ".sinaimg.cn/" + size + "/" + pid + "." + suffix
		//转成 rune 直接截取即可 "image/png"
		beego.Alert(sinaUrl)
		return sinaUrl
	}
	return ""
}
