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

	//是否允许游客上传
	AllowTourists bool `json:"allow_tourists"`

	//是否开启注册
	AllowRegister bool `json:"allow_register"`
	
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
	//是否开启 API
	OpenApiUpLoad bool `json:"open_api_up_load"`
	//Api token 空为 不设置token
	ApiToken string `json:"api_token"`
	//Api 默认上传图床 默认为搜狗 可选 SM 图床
	ApiDefault string `json:"api_default"`
	//是否对配置进行缓存
	CacheConfig bool `json:"cache_config"`
	//图床储存的一些配置
	SiteUploadWay UploadConfig `json:"site_upload_way"`

	AuxpiInfo Auxpi `json:"auxpi_info"`
}

type UploadConfig struct {
	//TODO:是否开启本地上传
	LocalStore LocalStore `json:"local_store"`
	//是否开启微博图床
	OpenSinaPicStore bool `json:"open_sina_pic_store"`
	//Sina Account
	SinaAccount SinaAccount `json:"sina_account"`
	//是否开启 flickr 图床 (此功能该可以在后台开启)
	OpenFlickrStore bool `json:"open_flickr_store"`
	//Flickr 配置
	FlickrAccount FlickrAccount `json:"flickr_account"`
}

type SinaAccount struct {
	//用户名
	UserName string `json:"user_name"`
	//密码
	PassWord string `json:"pass_word"`
	//新浪 Cookie 更新的频率,默认为3600s ,单位 s
	ResetSinaCookieTime int `json:"reset_sina_cookie_time"`
	//新浪图床默认使用的尺寸大小 square,thumb150,orj360,orj480,mw690,mw1024,mw2048,small,bmiddle,large 、默认为large
	DefultPicSize string `json:"defult_pic_size"`
}

//新浪公共接口，只需要提供 api 地址即可
//{"code":1,"msg":"操作成功","data":{"code":"200","width":176,"height":254,"size":13476,"pid":"005BYqpgly1fz9xxss19rj372jrq","url":"https:\/\/ws3.sinaimg.cn\/large\/005BYqpgly1fz9xxss19rj304w072jrq.jpg"},"runtime":"0.311697s"} 
type SinaPublicResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type FlickrAccount struct {
	//default size
	DefaultSize string `json:"default_size"`
	//api_key
	Id                 string `json:"id"`
	Api_key            string `json:"api_key"`
	Api_secret         string `json:"api_secret"`
	Oauth_token        string `json:"oauth_token"`
	Oauth_token_secret string `json:"oauth_token_secret"`
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
	//S string `json:"s"`
	//M string `json:"m"`
	D interface{} `json:"d"`
}

//网易图片
type NetEasyResp struct {
	Code string   `json:"code"`
	Data []string `json:"data"`
}

//Upload.cc
type CCResp struct {
	Code         int           `json:"code"`
	SuccessImage []interface{} `json:"success_image"`
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
	Open            bool   `json:"open"`
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

//Mail Config
type MailConfig struct {
	Status bool `json:"status"`
	Host string `json:"host"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Port string `json:"port"`
	From string `json:"from"`
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
}

//Role
type RoleJson struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	PIDs        []uint `json:"pids"`
}
