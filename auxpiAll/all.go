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

package auxpi

import (
	"encoding/xml"
	"time"
)

//Config 配置
type SiteConfig struct {
	//站点名称
	SiteName string `json:"site_name"`
	//底部信息
	SiteFooter string `json:"site_footer"`
	//网站链接
	SiteUrl string `json:"site_url"`
	//Logo 地址
	Logo string `json:"logo"`
	//最大上传的图片个数
	SiteUploadMaxNumber int `json:"site_upload_max_number"`
	//最大图片规格 MB
	SiteUploadMaxSize int64 `json:"site_upload_max_size"`

	Api ApiOptions `json:"api"`

	//是否允许游客上传
	AllowTourists bool `json:"allow_tourists"`

	//是否开启注册
	AllowRegister bool `json:"allow_register"`

	//代理配置
	ProxyStatus bool   `json:"proxy_status"`
	ProxyNode   string `json:"proxy_node"`

	//IP 配置
	IpConfig LimitConfig `json:"ip_config"`

	//分发配置
	DispatchOption Dispatch `json:"dispatch"`

	//邮件配置
	MailConfig MailConfig `json:"mail_config"`

	//是否使用 Mysql,使用 Mysql 后就不会再使用 json 进行配置
	DbOption DbOption `json:"db_option"`
	//JWT Token
	JwtSecret string `json:"jwt_secret"`
	//Jwt 认证时间
	JwtDueTime time.Duration `json:"jwt_due_time"`
	//加密所需 Salt
	AuxpiSalt string `json:"auxpi_salt"`
	ApiOptions
	//是否对配置进行缓存
	CacheConfig bool `json:"cache_config"`
	//图床储存的一些配置
	SiteUploadWay UploadConfig `json:"site_upload_way"`

	AuxpiInfo Auxpi `json:"auxpi_info"`
}

type UploadConfig struct {
	//TODO:是否开启本地上传
	LocalStore LocalStore `json:"local_store"`
	//Sina Account
	SinaAccount SinaAccount `json:"sina_account"`
	//Flickr 配置
	FlickrAccount FlickrAccount `json:"flickr_account"`
	//是否开启 Imgur 图床
	ImgurAccount ImgurAccount `json:"imgur_account"`
}

type SinaAccount struct {
	Status bool `json:"status"`
	//用户名
	UserName string `json:"user_name"`
	//密码
	PassWord string `json:"pass_word"`
	//新浪 Cookie 更新的频率,默认为3600s ,单位 s
	ResetSinaCookieTime int `json:"reset_sina_cookie_time"`
	//新浪图床默认使用的尺寸大小 square,thumb150,orj360,orj480,mw690,mw1024,mw2048,small,bmiddle,large 、默认为large
	DefultPicSize string `json:"defult_pic_size"`

	Proxy ProxyConf `json:"proxy"`
}

type SinaPublicResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type FlickrAccount struct {
	Status bool `json:"status"`
	//default size
	DefaultSize string `json:"default_size"`
	//api_key
	Id                 string `json:"id"`
	Api_key            string `json:"api_key"`
	Api_secret         string `json:"api_secret"`
	Oauth_token        string `json:"oauth_token"`
	Oauth_token_secret string `json:"oauth_token_secret"`
}

type ImgurAccount struct {
	Status  bool      `json:"status"`
	Proxy   ProxyConf `json:"proxy"`
	ImgurID string    `json:"imgur_id"`
}

type FlickrGetPicResp struct {
	XMLName        xml.Name `xml:"photo"`
	Id             string   `xml:"id,attr"`
	Secret         string   `xml:"secret,attr"`
	Server         string   `xml:"server,attr"`
	Farm           string   `xml:"farm,attr"`
	Dateuploaded   string   `xml:"dateuploaded,attr"`
	Originalsecret string   `xml:"originalsecret,attr"`
	Originalformat string   `xml:"originalformat,attr"`
}

//SM 图床 json
type SmResponse struct {
	Code string `json:"code"`
	Data SmData `json:"data"`
}

type SmData struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	Storename string `json:"storename"`
	Size      int    `json:"size"`
	Path      string `json:"path"`
	Hash      string `json:"hash"`
	Timestamp int    `json:"timestamp"`
	Ip        string `json:"ip"`
	Url       string `json:"url"`
	Delete    string `json:"delete"`
}

//Sina 图床 json
type SinaMsg struct {
	Code string   `json:"code"`
	Data SinaData `json:"data"`
}

type SinaData struct {
	Count int      `json:"count"`
	Data  string   `json:"data"`
	Pics  SinaPics `json:"pics"`
}

type SinaPics struct {
	Pic_1 picInfo `json:"pic_1"`
}

type picInfo struct {
	Width  int    `json:"width"`
	Size   int    `json:"size"`
	Ret    int    `json:"ret"`
	Height int    `json:"height"`
	Name   string `json:"name"`
	Pid    string `json:"pid"`
}

type SinaError struct {
	Retcode string `json:"retcode"`
	Reason  string `json:"reason"`
}

//Api & upload Json
type ResultJson struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data fileData `json:"data"`
}

type fileData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type RespJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ImageRespJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Delete string `json:"delete"`
		Name   string `json:"name"`
		URL    string `json:"url"`
	} `json:"data"`
}

//Db Option

type DbOption struct {
	UseDb       bool   `json:"use_db"`
	DbType      string `json:"db_type"`
	DbHost      string `json:"db_host"`
	DbName      string `json:"db_name"`
	DbUser      string `json:"db_user"`
	DbPass      string `json:"db_pass"`
	TablePrefix string `json:"table_prefix"`
}

//User Info Struct

type UserInfo struct {
	User         string      `json:"user"`
	Status       string      `json:"status"`
	Code         int         `json:"code"`
	Token        string      `json:"token"`
	Name         string      `json:"name"`
	Avatar       string      `json:"avatar"`
	Introduction string      `json:"introduction"`
	Roles        []string    `json:"roles"`
	Setting      interface{} `json:"setting"`
}

//文件类型结构体
type FormFile struct {
	Name  string //File Name
	Key   string //File upload Name
	Value []byte //File Value
	Type  string //File MIME Type
}

//百度图片
type BaiduResp struct {
	Errorn    int    `json:"errorn"`
	Url       string `json:"url"`
	QuerySign string `json:"querySign"`
	Simid     string `json:"simid"`
}

//掘金图片
type JueJinResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Key    string `json:"key"`
		Domain string `json:"domain"`
		URL    struct {
			HTTP  string `json:"http"`
			HTTPS string `json:"https"`
		} `json:"url"`
		ImageInfo struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Format string `json:"format"`
			Size   int    `json:"size"`
		} `json:"imageInfo"`
	} `json:"d"`
}

//网易图片
type NetEasyResp struct {
	Code string   `json:"code"`
	Data []string `json:"data"`
}

//Upload.cc
type CCResp struct {
	Code         int `json:"code"`
	TotalSuccess int `json:"total_success"`
	TotalError   int `json:"total_error"`
	SuccessImage []struct {
		Name      string `json:"name"`
		URL       string `json:"url"`
		Thumbnail string `json:"thumbnail"`
		Delete    string `json:"delete"`
	} `json:"success_image"`
}

//阿里
type AliResp struct {
	FsUrl string `json:"fs_url"`
	Code  string `json:"code"`
	Size  string `json:"size"`
	Width string `json:"width"`
	Url   string `json:"url"`
	Hash  string `json:"hash"`
}

//苏宁
type SuNingResp struct {
	ImgID        string `json:"imgId"`
	OriginalSize string `json:"originalSize"`
	Src          string `json:"src"`
	Errorcode    string `json:"errorcode"`
}

type FakerTable struct {
	Code int                 `json:"code"`
	Item []map[string]string `json:"item"`
}

//Faker Data
type FakerData struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PageViews   int    `json:"page_views"`
	DisPlayTime string `json:"dis_play_time"`
}

//Faker Images Info

type FakerImage struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data []interface{} `json:"data"`
}

type Auxpi struct {
	Author       string `json:"author"`
	Version      string `json:"version"`
	Branch       string `json:"branch"`
	Repositories string `json:"repositories"`
}

type LocalStore struct {
	Status          bool   `json:"status"`
	StorageLocation string `json:"storage_location"`
	Link            string `json:"link"`
}

type AuxpiCookie struct {
	UName      string `json:"u_name" valid:"Required"`
	Email      string `json:"email" valid:"Required;Email;MaxSize(100)"`
	ID         int    `json:"id" valid:"Required"`
	Version    string `json:"version" valid:"Required"`
	AuxpiToken string `json:"auxpi_token" valid:"Required;"`
}

type Dispatch struct {
	Status    bool          `json:"status"`
	Root      string        `json:"root"`
	RootID    int           `json:"root_id"`
	CacheTime time.Duration `json:"cache_time"`
}

//Imgur Resp
type ImgurResp struct {
	Data struct {
		ID          string        `json:"id"`
		Title       interface{}   `json:"title"`
		Description interface{}   `json:"description"`
		Datetime    int           `json:"datetime"`
		Type        string        `json:"type"`
		Animated    bool          `json:"animated"`
		Width       int           `json:"width"`
		Height      int           `json:"height"`
		Size        int           `json:"size"`
		Views       int           `json:"views"`
		Bandwidth   int           `json:"bandwidth"`
		Vote        interface{}   `json:"vote"`
		Favorite    bool          `json:"favorite"`
		Nsfw        interface{}   `json:"nsfw"`
		Section     interface{}   `json:"section"`
		AccountURL  interface{}   `json:"account_url"`
		AccountID   int           `json:"account_id"`
		IsAd        bool          `json:"is_ad"`
		InMostViral bool          `json:"in_most_viral"`
		Tags        []interface{} `json:"tags"`
		AdType      int           `json:"ad_type"`
		AdURL       string        `json:"ad_url"`
		InGallery   bool          `json:"in_gallery"`
		Deletehash  string        `json:"deletehash"`
		Name        string        `json:"name"`
		Link        string        `json:"link"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

//TouTiao
type TouTiaoResp struct {
	Width   int    `json:"width"`
	Message string `json:"message"`
	WebURL  string `json:"web_url"`
	WebURI  string `json:"web_uri"`
	Height  int    `json:"height"`
}

//XiaoMI

type XiaoMiResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

//Gitee
type GiteeResp struct {
	Content struct {
		Name        string `json:"name"`
		Path        string `json:"path"`
		Size        int    `json:"size"`
		Sha         string `json:"sha"`
		Type        string `json:"type"`
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
		DownloadURL string `json:"download_url"`
		Links       struct {
			Self string `json:"self"`
			HTML string `json:"html"`
		} `json:"_links"`
	} `json:"content"`
	Commit struct {
		Sha    string `json:"sha"`
		Author struct {
			Name  string    `json:"name"`
			Date  time.Time `json:"date"`
			Email string    `json:"email"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Date  time.Time `json:"date"`
			Email string    `json:"email"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		Parents []struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"parents"`
	} `json:"commit"`
}

//Github account
type GiteeAccount struct {
	AccessToken string `json:"access_token"`
	Owner       string `json:"owner"`
	Repo        string `json:"repo"`
}

//Github Request
type GithubRequest struct {
	Message   string `json:"message"`
	Committer struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"committer"`
	Content string `json:"content"`
}

//Github resp
type GithubResp struct {
	Content struct {
		Name        string `json:"name"`
		Path        string `json:"path"`
		Sha         string `json:"sha"`
		Size        int    `json:"size"`
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
		GitURL      string `json:"git_url"`
		DownloadURL string `json:"download_url"`
		Type        string `json:"type"`
		Links       struct {
			Self string `json:"self"`
			Git  string `json:"git"`
			HTML string `json:"html"`
		} `json:"_links"`
	} `json:"content"`
	Commit struct {
		Sha     string `json:"sha"`
		NodeID  string `json:"node_id"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
		Author  struct {
			Date  time.Time `json:"date"`
			Name  string    `json:"name"`
			Email string    `json:"email"`
		} `json:"author"`
		Committer struct {
			Date  time.Time `json:"date"`
			Name  string    `json:"name"`
			Email string    `json:"email"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			URL string `json:"url"`
			Sha string `json:"sha"`
		} `json:"tree"`
		Parents []struct {
			URL     string `json:"url"`
			HTMLURL string `json:"html_url"`
			Sha     string `json:"sha"`
		} `json:"parents"`
		Verification struct {
			Verified  bool        `json:"verified"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Payload   interface{} `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
}

//Github account
type GithubAccount struct {
	AccessToken string    `json:"access_token"`
	Owner       string    `json:"owner"`
	Repo        string    `json:"repo"`
	Email       string    `json:"email"`
	Proxy       ProxyConf `json:"proxy"`
}

type GithubMsg struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

//Mail Config
type MailConfig struct {
	Status bool   `json:"status"`
	Host   string `json:"host"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Port   string `json:"port"`
	From   string `json:"from"`
}

//ip
type LimitConfig struct {
	Status bool `json:"status"`
	//暂时性封号时长
	BlockTime int64 `json:"block_time"`
	//最后底线，超过这个底线直接自动封 ip
	DeadLine int `json:"dead_line"`
	//允许在某段时间内上传的图片张数
	AllowNum int `json:"allow_num"`
	//允许的时间，例如 允许用户每小时最大上传 X 张，这里的 AllowTime 是指每小时
	AllowTime int64 `json:"allow_time"`
}

//models

//images
type ImageJson struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	UserID  int    `json:"user_id"`
	StoreID int    `json:"store_id"`
	Delete  string `json:"delete"`
	Path    string `json:"path"`
	IP      string `json:"ip"`
}

//Role
type RoleJson struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	PIDs        []uint `json:"pids"`
}

type DistributionJson struct {
	Hash    string `json:"hash"`
	Url     string `json:"url"`
	IP      string `json:"ip"`
	StoreID int    `json:"store_id"`
	UserID  int    `json:"user_id"`
	RootID  int    `json:"root"`
	Number  int    `json:"number"`
	Short   string `json:"short"`
	RootUrl string `json:"root_url"`
	Delete  string `json:"delete"`
}

//菜单设计
type MenuReceive struct {
	Disable []struct {
		ID         int         `json:"id"`
		CreatedOn  int         `json:"created_on"`
		ModifiedOn int         `json:"modified_on"`
		DeletedOn  int         `json:"deleted_on"`
		CreatedDay string      `json:"created_day"`
		Name       string      `json:"name"`
		Desc       string      `json:"desc"`
		URL        string      `json:"url"`
		Icon       string      `json:"icon"`
		Color      string      `json:"color"`
		API        string      `json:"api"`
		Router     string      `json:"router"`
		Status     bool        `json:"Status"`
		Weight     int         `json:"weight"`
		Rank       int         `json:"rank"`
		Auth       int         `json:"Auth"`
		Images     interface{} `json:"Images"`
	} `json:"disable"`
	Enable []struct {
		ID         int         `json:"id"`
		CreatedOn  int         `json:"created_on"`
		ModifiedOn int         `json:"modified_on"`
		DeletedOn  int         `json:"deleted_on"`
		CreatedDay string      `json:"created_day"`
		Name       string      `json:"name"`
		Desc       string      `json:"desc"`
		URL        string      `json:"url"`
		Icon       string      `json:"icon"`
		Color      string      `json:"color"`
		API        string      `json:"api"`
		Router     string      `json:"router"`
		Status     bool        `json:"Status"`
		Weight     int         `json:"weight"`
		Rank       int         `json:"rank"`
		Auth       int         `json:"Auth"`
		Images     interface{} `json:"Images"`
	} `json:"enable"`
}

//Proxy
type ProxyConf struct {
	Status bool   `json:"status"`
	Node   string `json:"node"`
}

//Upload Auth
type UploadAuth struct {
	Github GithubAccount `json:"github"`
	Gitee  GiteeAccount  `json:"gitee"`
	Sina   SinaAccount   `json:"sina"`
	Imgur  ImgurAccount  `json:"imgur"`
}

//API options
type ApiOptions struct {
	Status     bool   `json:"status"`
	Auth       bool   `json:"auth"`
	ApiDefault string `json:"api_default"`
}

//Site Base Options
type SiteBase struct {
	//站点名称
	SiteName string `json:"site_name"`
	//底部信息
	SiteFooter string `json:"site_footer"`
	//网站链接
	SiteUrl string `json:"site_url"`
	//Logo 地址
	Logo string `json:"logo"`
	//最大上传的图片个数
	SiteUploadMaxNumber int `json:"site_upload_max_number"`
	//最大图片规格 MB
	SiteUploadMaxSize int64 `json:"site_upload_max_size"`
	//是否允许游客上传
	AllowTourists bool `json:"allow_tourists"`
	//是否开启注册
	AllowRegister bool `json:"allow_register"`
	//SMTP
	JwtSecret string `json:"jwt_secret"`
	//Jwt 认证时间
	JwtDueTime time.Duration `json:"jwt_due_time"`

	MailConfig MailConfig `json:"mail_config"`
}

//Dispatch Options
type DisPatchOptions struct {
	Node string `json:"node"`
}
